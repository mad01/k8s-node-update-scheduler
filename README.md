# k8s-node-update-scheduler

service to check the current state of all nodes in a cluster (node masters)
    - annotate nodes with incorrect k8s version for termination (i.e update with kops as long as remote state is updated)
