# k8s-node-update-scheduler

a cli tool to schedule update/terminations of nodes in a kubernetes cluster designed to work together with [node-terminator](https://github.com/mad01/k8s-node-terminator) that looks for annotations. This tool implements adding those annotations. The result is that you can tag N instances for update/termination

the time window flags support `hh:mm AM/PM` as time format

## example

the cli will annotate all nodes in the cluster will skip masters since selector won't match. The annotations will set the set terminations to be allowed between 2-5 in the morning and the terminate flag till be set to true
```
k8s-node-update-scheduler schedule \
    --selector kubernetes.io/role=node \
    --kube.config ~/.kube/config  \
    --schedule.fromWindow="02:01 AM" \
    --schedule.toWindow="05:01 AM" \
    --terminate
```
