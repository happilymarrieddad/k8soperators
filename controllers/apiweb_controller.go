/*
Copyright 2020 Nick Kotenberg.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hyperledgerv1alpha1 "k8soperators/api/v1alpha1"
)

// ApiWebReconciler reconciles a ApiWeb object
type ApiWebReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=hyperledger.example.org,resources=apiwebs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=hyperledger.example.org,resources=apiwebs/status,verbs=get;update;patch

func (r *ApiWebReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("apiweb", req.NamespacedName)

	instance := &hyperledgerv1alpha1.ApiWeb{}
	if err := r.Get(ctx, req.NamespacedName, instance); err != nil {
		if errors.IsNotFound(err) {
			log.Info("ApiWeb resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}

		log.Error(err, "Failed to get ApiWeb")
		return ctrl.Result{}, err
	}

	// If deployment is not found or there is an error
	found := &appsv1.Deployment{}
	if err := r.Get(ctx, types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, found); err != nil && errors.IsNotFound(err) {
		srv, dpl := createApiServiceAndDeployment(instance)

		if err3 := r.Get(ctx, types.NamespacedName{Name: "service-" + instance.Name, Namespace: instance.Namespace}, found); err3 != nil && errors.IsNotFound(err3) {
			log.Info("Creating a new Service", "Service.Namespace", srv.Namespace, "Service.Name", "service-"+srv.Name)
			if err2 := r.Create(ctx, srv); err2 != nil {
				log.Error(err2, "Failed to create new Service", "Service.Namespace", srv.Namespace, "Service.Name", "service-"+srv.Name)
				return ctrl.Result{}, err2
			}
		}

		log.Info("Creating a new Deployment", "Deployment.Namespace", dpl.Namespace, "Deployment.Name", dpl.Name)
		if err2 := r.Create(ctx, dpl); err2 != nil {
			log.Error(err2, "Failed to create new Deployment", "Deployment.Namespace", dpl.Namespace, "Deployment.Name", dpl.Name)
			return ctrl.Result{}, err2
		}

		return ctrl.Result{}, err
	} else if err != nil {
		log.Error(err, "Failed to get deployment")
		return ctrl.Result{}, err
	}

	// Update current deployment
	apiImage := instance.Spec.ApiImage
	if len(found.Spec.Template.Spec.Containers) > 0 {
		con := found.Spec.Template.Spec.Containers[0]

		strs := strings.Split(con.Image, ":")
		if len(strs) > 1 && strs[1] != apiImage {
			strs[1] = apiImage
			con.Image = strings.Join(strs, ":")

			found.Spec.Template.Spec.Containers[0] = con

			if err := r.Update(ctx, found); err != nil {
				log.Error(err, "Failed to update Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
				return ctrl.Result{}, err
			}
			return ctrl.Result{Requeue: true}, nil
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ApiWebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hyperledgerv1alpha1.ApiWeb{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}

func createApiServiceAndDeployment(cr *hyperledgerv1alpha1.ApiWeb) (*corev1.Service, *appsv1.Deployment) {
	port := int32(8000)
	label := map[string]string{
		"app": "api",
	}

	if len(cr.Spec.ApiImage) == 0 {
		cr.Spec.ApiImage = "latest"
	}

	return &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "service-" + cr.ObjectMeta.Name,
				Namespace: cr.ObjectMeta.Namespace,
				Labels:    label,
			},
			Spec: corev1.ServiceSpec{
				Type:     corev1.ServiceTypeClusterIP,
				Selector: label,
				Ports:    []corev1.ServicePort{{Port: port, TargetPort: intstr.FromInt(int(port))}},
			},
		}, &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      cr.ObjectMeta.Name,
				Namespace: cr.ObjectMeta.Namespace,
				Labels:    label,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: func(a int32) *int32 { return &a }(3),
				Selector: &metav1.LabelSelector{
					MatchLabels: label,
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: label,
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "api",
								Image: "happilymarrieddadudemy/node-web-server:" + cr.Spec.ApiImage,
								Ports: []corev1.ContainerPort{{ContainerPort: port}},
							},
						},
					},
				},
			},
		}
}
