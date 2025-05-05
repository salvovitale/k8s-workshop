# List of commands part 1

## Hands-on Pods
Lets create a pod for the hello-world app

```shell
kubectl run hello-world-1 --image=app-hello-world:latest --image-pull-policy=IfNotPresent --restart=Never
```

Lets check that the pod is running

```shell
kubectl get pods hello-world-1
```

Use custom-columns
```shell
kubectl get pod hello-world-1 --output \
custom-columns=NAME:metadata.name,NODE_IP:status.hostIP,POD_IP:status.podIP
```

Lets get the ContainerID field 

```shell
kubectl get pod hello-world-1 -o jsonpath='{.status.containerStatuses[0].containerID}'
```

Port Forward

```shell
kubectl port-forward pod/hello-world-1 8080:8080
```

## Hands-on Deployment

Lets create a deployment with the kubectl cli

```shell
 kubectl create deployment  hello-world-2 --image=salvovitale/k8s-ws-app-hello-world:latest
```

check the deployment
```shell
kubectl get deploy hello-world-2
```

```shell
kubectl describe deploy hello-world-2
```

Get the labels for the deployment

```shell
kubectl get deploy hello-world-2 -o jsonpath='{.spec.template.metadata.labels}'
```

we can use the labels to identify the pods

```shell
kubectl get pod -l app=hello-world-2
```

lets inspect the current labels situation in pods
```shell
kubectl get pods -o custom-columns=NAME:metadata.name,LABELS:metadata.labels
```

lets override the label of a pod linked to the deployment

```shell
 kubectl label pods -l app=hello-world-2 --overwrite app=hello-world-x
```

lets revert back the situation

```shell
 kubectl label pods -l app=hello-world-x --overwrite app=hello-world-2
```

## Hands-On Manifests

create an initial pod manifest with kubectl
```shell
kubectl run hello-world-3 \
--image=salvovitale/k8s-ws-app-hello-world:latest \
--image-pull-policy=IfNotPresent \
-o yaml \
--dry-run=client
```

create a deployment manifest
```shell
kubectl create deployment hello-world-4 \
--image="app-hello-world:latest" \
-o yaml \
--dry-run=client > deployment.yaml
```

### Interactive shell

interactive shell with pod

```shell
kubectl exec -it hello-world-3 -- sh
```

interactive shell with deployment

```shell
kubectl exec -it deploy/hello-world-4 -- sh
```

### Inspect Logs

pod logs
```shell
 kubectl logs pod/hello-world-3 -f
```

deployment logs

```shell
kubectl logs deploy/hello-world-4 -f
```

### Filesystem

copy a file from local filesystem to pod

```shell
 kubectl cp temp/text.txt hello-world-3:/app/text.txt
```


### Delete Resources

delete all the pods

```shell
kubectl delete pods --all
```

delete all the deployments

```shell
kubectl delete deploy --all
```

