#!/bin/bash
sudo rm -rf out
VERSION=v1.20.5
make -e REGISTRY=fred78290 -e TAG=$VERSION container -e GOARCH=amd64
make -e REGISTRY=fred78290 -e TAG=$VERSION container -e GOARCH=arm64

docker push fred78290/cluster-autoscaler:$VERSION-amd64
docker push fred78290/cluster-autoscaler:$VERSION-arm64

docker manifest create fred78290/cluster-autoscaler:$VERSION --amend fred78290/cluster-autoscaler:$VERSION-amd64 --amend fred78290/cluster-autoscaler:$VERSION-arm64
docker manifest push fred78290/cluster-autoscaler:$VERSION