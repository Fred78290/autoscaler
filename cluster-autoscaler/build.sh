#!/bin/bash
sudo rm -rf out
VERSION=v1.19.9
REGISTRY=fred78290

make -e REGISTRY=$REGISTRY -e TAG=$VERSION container -e GOARCH=amd64
make -e REGISTRY=$REGISTRY -e TAG=$VERSION container -e GOARCH=arm64

docker buildx build --pull --platform linux/amd64,linux/arm64 --push -t ${REGISTRY}/cluster-autoscaler:${VERSION} .