# Deploy with kustomize

For example:

```shell
kubectl apply -k deploy/kustomize/overlays/dev
```

or

```shell
kubectl apply -k --enable-helm deploy/kustomize/overlays/k3d-k3s
```
