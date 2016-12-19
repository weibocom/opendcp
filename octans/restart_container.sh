#!/bin/bash

TAG_NAME="octans"
REGISTRY_PATH="xxx/yinan/octans"

cid=`docker ps -a |grep $TAG_NAME | awk '{printf $1" "}'`

if [ "$cid" !=  "" ];then

docker rm -f $cid

echo "delete container $cid"

fi

iid=`docker images | grep $TAG_NAME | awk '{printf $3" "}'`

if [ "$iid" !=  "" ];then

docker rmi $iid

echo "delete image $iid"

fi

echo "pull image from daocloud"

docker pull $REGISTRY_PATH

niid=`docker images | grep $TAG_NAME | awk '{printf $3" "}'`

if [ "$niid" =  "" ];then

echo 'pull image failed....'

return 1

fi

docker run --name $TAG_NAME --net=host -v /data0/$TAG_NAME/config.yml:/data/octans/config.yml -d $REGISTRY_PATH

echo "start container success"

