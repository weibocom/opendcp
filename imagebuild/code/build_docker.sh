#!/bin/bash

# #############################
# 构建docker镜像
# #############################

# 定位到脚本所在目录
cd `dirname $0`
echo "current pwd:"`pwd`

CLOUD=
if [ "" != "$1" ] ;then
    CLOUD=$1
fi
echo "CLOUD:$CLOUD"

if [ "$CLOUD" = "aliyun" ]; then
    \cp -f Dockerfile-Golang-Alpine_Aliyun  Dockerfile-Golang-Alpine
fi

if [ "$CLOUD" = "dockerio" ]; then
    \cp -f Dockerfile-Golang-Alpine_Dockerio  Dockerfile-Golang-Alpine
fi


# 定位到上层目录也就是项目目录
SERVER_BASEDIR=$(dirname $(pwd))

echo "start build dcp/golang-alpine:latest.."
# 构建alpine-golang镜像，这个镜像用于编译当前项目
docker build -t dcp/golang-alpine:latest -f Dockerfile-Golang-Alpine .
echo "finish build dcp/golang-alpine:latest.."

# 编译当前项目
BASEDIR=$(dirname ${SERVER_BASEDIR})
echo "start build current project..."
docker run --rm --net=host --name go-build -v ${SERVER_BASEDIR}:/go/src/weibo.com/opendcp/imagebuild -v ${SERVER_BASEDIR}/vendor/github.com:/go/src/github.com -w /go/src/weibo.com/opendcp/imagebuild dcp/golang-alpine:latest sh -c " /bin/bash code/go_build.sh "
echo "finish build current project..."

echo "start build dcp/imagebuild:latest..."
docker build -t dcp/imagebuild:latest ${SERVER_BASEDIR}
echo "finish build dcp/imagebuild:latest..."
