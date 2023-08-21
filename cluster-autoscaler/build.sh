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

sudo rm -rf out vendor

VERSION=v1.25.12
REGISTRY=fred78290

make -e REGISTRY=$REGISTRY -e TAG=$VERSION -e DOCKER_RM=1 container -e GOARCH=amd64
make -e REGISTRY=$REGISTRY -e TAG=$VERSION -e DOCKER_RM=1 container -e GOARCH=arm64

mkdir -p out/linux/amd64 out/linux/arm64
cp cluster-autoscaler-amd64 out/linux/amd64/cluster-autoscaler
cp cluster-autoscaler-arm64 out/linux/arm64/cluster-autoscaler

docker buildx build --pull --platform linux/amd64,linux/arm64 --push -t ${REGISTRY}/cluster-autoscaler:${VERSION} .
