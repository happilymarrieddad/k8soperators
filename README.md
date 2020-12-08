Learning Kubernetes Operators
===============================

## Installation - Linux
[Operator SDK Install](https://github.com/operator-framework/operator-sdk)
[Kubebuilder Install](https://book-v1.book.kubebuilder.io/getting_started/installation_and_setup.html)

## Example - https://codeburst.io/kubernetes-operators-by-example-99a77ea4ac43
example.com as the Kubernetes API group
v1alpha1 as the API version
Hello as the API resource kind
```bash
apiVersion: example.com/v1alpha1
kind: Hello
metadata:
  name: example-hello
spec:
  favoriteDrink: tea
```

## Clean
```bash
rm PROJECT 
rm Makefile 
rm main.go 
rm go.mod
rm go.sum 
rm Dockerfile 
rm -rf controllers 
rm -rf config 
rm -rf api 
rm -rf hack 
```

## Commands
```bash
operator-sdk init --project-version=2 --domain example.org --license apache2 --owner "Nick Kotenberg" --repo=k8soperators
```

```bash
operator-sdk create api --group hyperledger --version v1alpha1 --kind CertificateAuthority
```

```bash
kubectl apply -f deploy/hyperledger/example.yaml
```

```bash
kubectl describe certificateauthorities certificate-authorities
```

### Api/Web example
```bash
operator-sdk create api --group hyperledger --version v1alpha1 --kind ApiWeb

make install

make deploy

kubectl describe apiwebs api-web
```

