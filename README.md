# k8s-node-update-scheduler

a cli tool to schedule update/terminations of nodes in a kubernetes cluster designed to work together with [node-terminator](https://github.com/mad01/k8s-node-terminator) that looks for annotations. This tool implements adding those annotations. The result is that you can tag N instances for update/termination

the time window flags support `hh:mm AM/PM` as time format

## example

will annotate all nodes in the cluster will skip masters since selector won't match with a crom time window between 2 and 5 in the morning
```
k8s-node-update-scheduler schedule --selector kubernetes.io/role=node --kube.config ~/.kube/config  --schedule.fromWindow="02:01 AM" --schedule.toWindow="05:01 AM" --reboot
```
