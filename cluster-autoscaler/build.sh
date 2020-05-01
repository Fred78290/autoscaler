#!/bin/bash
sudo rm -rf out
make -e REGISTRY=fred78290 -e TAG=v1.17.5 container