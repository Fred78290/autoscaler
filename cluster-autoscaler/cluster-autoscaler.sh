#!/bin/bash

# Copyright 2021 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
