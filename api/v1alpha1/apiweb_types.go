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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApiWebSpec defines the desired state of ApiWeb
type ApiWebSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ApiImage string `json:"apiImage,omitempty"`
	WebImage string `json:"webImage,omitempty"`
}

// ApiWebStatus defines the observed state of ApiWeb
type ApiWebStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ApiStatus string `json:"apiStatus"`
	WebStatus string `json:"webStatus"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ApiWeb is the Schema for the apiwebs API
type ApiWeb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiWebSpec   `json:"spec,omitempty"`
	Status ApiWebStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiWebList contains a list of ApiWeb
type ApiWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiWeb `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApiWeb{}, &ApiWebList{})
}
