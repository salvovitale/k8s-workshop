# Set-up the local development setup (LDS)

The following LDS setup instructions are created only for MacOS users.

## Required tooling

### Install required tooling for the workshop
- docker
- [kind](https://kind.sigs.k8s.io/)
- kubectl
- [go-task]( https://taskfile.dev/)
- k9s (optional)

Install kind
```shell
brew install kind
```
Install Kubectl
```shell
brew install kubectl
```
Install go-task
```shell
brew install go-task
```
Install k9s
```shell
brew install k9s
```

## Create/Destroy the k8s cluster

Create the cluster

```shell
task lds:create-k8s-kind-cluster
```

Verify that the cluster is up and running
```shell
kubectl get nodes
```
you should get an output like this

```shell
NAME                   STATUS   ROLES           AGE   VERSION
k8s-ws-control-plane   Ready    control-plane   33m   v1.32.0
k8s-ws-worker          Ready    <none>          33m   v1.32.0
k8s-ws-worker2         Ready    <none>          33m   v1.32.0
```

Destroy the cluster
```shell
task lds:destroy-k8s-kind-cluster
```
