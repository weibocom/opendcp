#!/bin/sh

CLOUD=$1
VER=$2

code/build_docker.sh $CLOUD
docker tag dcp/imagebuild:latest $VER
