#!/bin/bash
sudo rm -rf out

make -e REGISTRY=fred78290 -e TAG=v1.21.0 container -e GOARCH=amd64
make -e REGISTRY=fred78290 -e TAG=v1.21.0 container -e GOARCH=arm64