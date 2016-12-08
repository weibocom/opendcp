#!/bin/bash

IMAGE_TAG="octans"

iid=`docker images | grep $IMAGE_TAG | awk '{printf $3" "}'`

if [ "$iid" !=  "" ];then

docker rmi $iid

echo "delete image $iid"

fi
# xxx is registry add
docker build -t xxx/$IMAGE_TAG .

docker push xxx/yinan/$IMAGE_TAG

echo "build docker image success"
