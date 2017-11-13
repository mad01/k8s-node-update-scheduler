# k8s-node-update-scheduler

service to check the current state of all nodes in a cluster (node masters)
    - annotate nodes with incorrect k8s version for termination (i.e update with kops as long as remote state is updated)


## example
```
k8s-node-update-scheduler schedule --selector kubernetes.io/role=node --kube.config ~/.kube/config  --schedule.fromWindow="* 12 * * *" --schedule.toWindow="* 17 * * *"
```
