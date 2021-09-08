#!/bin/bash

DIR=$(dirname $0)

$DIR/cluster-autoscaler \
--v=4 \
--stderrthreshold=info \
--cloud-provider=grpc \
--cloud-config=$HOME/go/src/github.com/Fred78290/kubernetes-vmware-autoscaler/masterkube/config/grpc-config.json \
--kubeconfig=$HOME/go/src/github.com/Fred78290/kubernetes-vmware-autoscaler/masterkube/cluster/config \
--nodes="0:3:true/afp-slyo-ca-k8s|monitor=true;database=true" \
--max-nodes-total=3 \
--cores-total=0:16 \
--memory-total=0:24 \
--node-autoprovisioning-enabled \
--max-autoprovisioned-node-group-count=1 \
--unremovable-node-recheck-timeout=2m \
--scale-down-enabled=true \
--scale-down-delay-after-add=1m \
--scale-down-delay-after-delete=1m \
--scale-down-delay-after-failure=1m \
--scale-down-unneeded-time=1m \
--scale-down-unready-time=1m
