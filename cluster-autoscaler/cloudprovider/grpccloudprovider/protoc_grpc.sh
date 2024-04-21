#/bin/bash
CURDIR=$(dirname $0)
PB_RELEASE="25.1"
PB_REL="https://github.com/protocolbuffers/protobuf/releases"

export PROTOC_DIR="/tmp/protoc-${PB_RELEASE}"
export GOPATH=$PROTOC_DIR
export GO111MODULE=on
export PATH=$PROTOC_DIR/bin:$PATH

mkdir -p $PROTOC_DIR

pushd $PROTOC_DIR

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

PB_ARCH=$([[ "$(uname -m)" =~ arm64|aarch64 ]] && echo -n aarch_64 || echo -n x86_64)
OS=$([ "$(uname)" = "Darwin" ] && echo -n osx || echo -n linux)

curl -sLO ${PB_REL}/download/v${PB_RELEASE}/protoc-${PB_RELEASE}-${OS}-${PB_ARCH}.zip
unzip protoc-${PB_RELEASE}-${OS}-${PB_ARCH}.zip

popd

$PROTOC_DIR/bin/protoc -I . -I vendor --proto_path=cloudprovider/grpccloudprovider --go_out=cloudprovider/grpccloudprovider --go-grpc_out=cloudprovider/grpccloudprovider grpc.proto

pushd $CURDIR
cp ./grpc/grpccloudprovider/*.go .
rm -rf ./grpc
popd

sudo rm -rf $PROTOC_DIR
