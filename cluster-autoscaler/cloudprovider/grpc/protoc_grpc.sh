#/bin/bash

PB_RELEASE="3.11.1"
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
PROTOC_DIR=$(mktemp -d /tmp/protoc-${PB_RELEASE}-XXXX)

go get -v github.com/golang/protobuf/protoc-gen-go@v1.3.2

pushd $PROTOC_DIR
curl -LO ${PB_REL}/download/v${PB_RELEASE}/protoc-${PB_RELEASE}-linux-x86_64.zip
unzip protoc-${PB_RELEASE}-linux-x86_64.zip
popd

$PROTOC_DIR/bin/protoc -I . -I vendor cloudprovider/grpc/grpc.proto --go_out=plugins=grpc:.

rm -rf $PROTOC_DIR