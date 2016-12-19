#!/bin/bash

# #############################
# 构建并且本地直接运行
# #############################


# 定位到脚本所在目录
cd `dirname $0`
# 定位到上层目录也就是项目目录
SERVER_BASEDIR=$(dirname $(pwd))

#　是否使用docker build，用于没有安装golang的环境
USE_DOCKER_BUILD=false
RUN_BACKGROUND=false

while getopts "bdl:" opt; do
    case ${opt} in
        d)
            USE_DOCKER_BUILD=true
        ;;
        b)
            RUN_BACKGROUND=true
        ;;
        \?)
        ;;
    esac
done

echo "current work path: "${SERVER_BASEDIR}
echo "whether use docker to build project: "${USE_DOCKER_BUILD}
echo "whether run background: "${RUN_BACKGROUND}

# 编译当前项目
echo "begin build current project..."
if test ${USE_DOCKER_BUILD} == true; then
    #　使用镜像
    docker run --rm --net=host --name go-build -v ${SERVER_BASEDIR}:/go/src/weibo.com/opendcp/imagebuild -w /go/src/weibo.com/opendcp/imagebuild golang sh -c " go get ./... ; /bin/bash code/go_build.sh "
else
    # 本地编译
    /bin/bash ./go_build.sh
fi
echo "finish build current project..."


# 运行程序
cd ${SERVER_BASEDIR}/code/web

if test ${RUN_BACKGROUND} == true; then
    nohup ./main ${SERVER_BASEDIR} > /tmp/imagebuild_console.log 2>&1 &
else
    ./main ${SERVER_BASEDIR}
fi
