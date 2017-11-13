# k8s-node-update-scheduler

a cli tool to schedule update/terminations of nodes in a kubernetes cluster designed to work together with [node-terminator](https://github.com/mad01/k8s-node-terminator) that looks for the annotations added my this cli tool.

t

## example

will annotate all nodes in the cluster will skip masters since selector won't match with a crom time window between 12-17
```
k8s-node-update-scheduler schedule --selector kubernetes.io/role=node --kube.config ~/.kube/config  --schedule.fromWindow="* 12 * * *" --schedule.toWindow="* 17 * * *"
```
