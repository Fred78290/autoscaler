#!/bin/bash
sudo rm -rf out

VERSION=v1.22.2
REGISTRY=fred78290

make -e REGISTRY=$REGISTRY -e TAG=$VERSION -e DOCKER_RM=1 container -e GOARCH=amd64
make -e REGISTRY=$REGISTRY -e TAG=$VERSION -e DOCKER_RM=1 container -e GOARCH=arm64

mkdir -p out/linux/amd64 out/linux/arm64
cp cluster-autoscaler-amd64 out/linux/amd64/cluster-autoscaler
cp cluster-autoscaler-arm64 out/linux/arm64/cluster-autoscaler

docker buildx build --pull --platform linux/amd64,linux/arm64 --push -t ${REGISTRY}/cluster-autoscaler:${VERSION} .
