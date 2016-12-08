#!/bin/sh

DIRS="orion jupiter octans db_init proxy hubble dcp_open imagebuild"
REG=weibo.com
LOC=opendcp
VER=latest

if [ "" != "$1" ] ;then
    VER=$1
fi

for DIR in $DIRS; do
    TAG=${REG}/${LOC}/${DIR}:${VER}
    TAG_PUB=${LOC}/${DIR}:${VER}
    docker tag $TAG $TAG_PUB
    echo "Pushing $TAG ..."
    docker push $TAG #_PUB
done
