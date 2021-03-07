#!/bin/bash
sudo rm -rf out
make -e REGISTRY=fred78290 -e TAG=v1.21.0-beta.0 container