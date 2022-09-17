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

PB_RELEASE="3.11.1"
PB_REL="https://github.com/protocolbuffers/protobuf/releases"

export PROTOC_DIR="/tmp/protoc-${PB_RELEASE}"
export GOPATH=$PROTOC_DIR
export GO111MODULE=on
export PATH=$PROTOC_DIR/bin:$PATH

mkdir -p $PROTOC_DIR

pushd $PROTOC_DIR

#go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

go get -v google.golang.org/grpc@v1.29.1
#go get -v github.com/golang/protobuf@v1.3.2
#go get -v github.com/golang/protobuf/protoc-gen-go@v1.3.2

curl -LO ${PB_REL}/download/v${PB_RELEASE}/protoc-${PB_RELEASE}-linux-x86_64.zip
unzip protoc-${PB_RELEASE}-linux-x86_64.zip

popd

$PROTOC_DIR/bin/protoc -I . -I vendor cloudprovider/grpc/grpc.proto --go_out=plugins=grpc:.

sudo rm -rf $PROTOC_DIR