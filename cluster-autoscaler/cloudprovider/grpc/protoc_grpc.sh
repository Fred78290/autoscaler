#/bin/bash

protoc -I . -I vendor cloudprovider/grpc/grpc.proto --go_out=plugins=grpc:.
