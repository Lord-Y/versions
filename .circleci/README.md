# developments

## Creation
For our development we will use kind like this:

```bash
sudo kind create cluster --name versions --image kindest/node:v1.22.0 --wait 5m --config kind-config.yaml
sudo cp -r /root/.kube/ ~/ && sudo chown -R $USER:$USER ~/.kube
sudo kubectl cluster-info --context versions

# to fix dns issues
kubectl -n kube-system apply -f configmap.yaml
kubectl -n kube-system rollout restart deploy coredns
```

Kind documentation can be found [here](https://kind.sigs.k8s.io/docs/user/quick-start/)

## Cleaning

```bash
sudo kind delete cluster --name versions
```