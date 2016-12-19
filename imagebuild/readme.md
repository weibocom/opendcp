# 说明
该项目用于docker镜像构建。

# 目录结构
```
.
├── code 源码目录
├── Dockerfile 构建镜像的dockerfile
├── global_config 全局配置
└── sql 数据库语句
```

# 使用说明

## 配置文件

* globle_config/env

```
SERVER_IP=127.0.0.1   # 服务器ip
SERVER_PORT=8080 # 服务器端口
MYSQL_HOST=127.0.0.1 # mysql ip
MYSQL_PORT=3306 # mysql port
MYSQL_USER=root # mysql user
MYSQL_PASSWORD=19890809 # mysql password
LOG_PATH=/tmp/imagebuild.log # 日志路径
PLUGIN_VIEW_URL=/view/plugin # 可配置接口路径
EXTENSION_INTERFACE_URL=/api/plugin/extension/interface　＃　可配置接口路径
```

* code/web/conf/app.conf

beego的配置文件

beego地址[http://beego.me/](http://beego.me/)

## 本地直接运行
cd imagebuild/code/
/bin/bash local-run.sh -d -b

参数说明

* -d 使用docker进行golang的构建，如果本地有go环境可以不使用该参数
* -b 后台运行

脚本说明

```

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
    #　从官方的docker hub下载golang镜像，并且将当前的目录挂在到golang镜像上，在镜像内部执行go get下载项目的依赖包，然后调用go-build.sh进行实际的编译
    docker run --rm --net=host --name go-build -v ${SERVER_BASEDIR}:/go/src/imagebuild -w /go/src/imagebuild golang sh -c " go get ./... ; /bin/bash code/go-build.sh "
else
    # 本地编译
    /bin/bash ./go-build.sh
fi
echo "finish build current project..."


# 运行程序
cd ${SERVER_BASEDIR}/code/web

if test ${RUN_BACKGROUND} == true; then
    nohup ./main ${SERVER_BASEDIR} > /tmp/imagebuild_console.log 2>&1 &
else
    ./main ${SERVER_BASEDIR}
fi
```

## 使用docker运行

### 运行

* 构建docker镜像

/bin/bash build-docker.sh
* 运行镜像

docker run -tid --name=imagebuild --privileged -p 8081:8080 -v /data1/workspace_liupeng/imagebuild/tmp:/tmp -v /data1/workspace_liupeng/imagebuild/project:/imagebuild/project -e SERVER_IP=xx.xx.xx.xx -e SERVER_PORT=8081 -e MYSQL_HOST=xx.xx.xx.xx -e MYSQL_PORT=3307 -e MYSQL_USER=root -e MYSQL_PASSWORD=pwd -e PLUGIN_VIEW_URL=/api/for_repos/proxy_view_plugin.php -e EXTENSION_INTERFACE_URL=/api/for_repos/proxy_interface.php -e HARBOR_ADDRESS=xxx.xxx.xxx.xxx:yyyy -e HARBOR_USER=admin -e HARBOR_PASSWORD=Harbor12345 dcp/imagebuild:latest

运行参数说明:
HARBOR_ADDRESS # harbor地址
HARBOR_USER # harbor用户
HARBOR_PASSWORD # harbor密码
其他参数和evn一致。

### 脚本说明

* code/Dockerfile-Golang-Alpine
用于编译go工程的镜像的dockerfile

```
FROM golang:alpine
# 安装必要的软件
RUN apk add --no-cache bash git openssh-client ca-certificates
```

* code/build-docker.sh
用于构建当前项目docker镜像

```
#!/bin/bash

# #############################
# 构建docker镜像
# #############################

# 定位到脚本所在目录
cd `dirname $0`
echo "current pwd:"`pwd`

# 定位到上层目录也就是项目目录
SERVER_BASEDIR=$(dirname $(pwd))

echo "start build dcp/golang-alpine:latest.."
# 构建alpine-golang镜像，这个镜像用于编译当前项目
docker build -t dcp/golang-alpine:latest -f Dockerfile-Golang-Alpine .
echo "finish build dcp/golang-alpine:latest.."

echo "start build current project..."
# 编译当前项目
docker run --rm --net=host --name go-build -v ${SERVER_BASEDIR}:/go/src/imagebuild -w /go/src/imagebuild dcp/golang-alpine:latest sh -c " go get ./... ; /bin/bash code/go-build.sh "
echo "finish build current project..."

echo "start build dcp/imagebuild:latest..."
docker build -t dcp/imagebuild:latest ${SERVER_BASEDIR}
echo "finish build dcp/imagebuild:latest..."
```

* code/go-build.sh

调用go进行项目编译

```
#!/bin/bash

##########################
# 编译所有插件以及主项目
##########################

# 定位到脚本所在目录
cd `dirname $0`

SRC_DIR=`pwd`

BUILD_PLUGINS_BASE=${SRC_DIR}/plugins/build/

DOCKERFILE_PLUGINS_BASE=${SRC_DIR}/plugins/dockerfile/

TO_BUILD_FOLDER=(${BUILD_PLUGINS_BASE} ${DOCKERFILE_PLUGINS_BASE})

for folder in ${TO_BUILD_FOLDER[@]}
do
    for file in ${folder}*
    do
        if test -d ${file}
        then
           echo "begin build plugin: "${file}
           plug=${file}"/"`basename ${file}`"Plugin"
           go build -o ${plug} ${plug}".go"
           echo "finish build plugin: "${file}
        else
           echo ${file}
        fi
    done
done

echo "begin build main"
go build -o ${SRC_DIR}/web/main ${SRC_DIR}/web/main.go
echo "finish build main"
```

* Dockerfile

当前工程的dockerfile，用于构建镜像

```
FROM docker:dind

# 安装可能会用到的工具
RUN apk add --no-cache bash git subversion openssh-client ca-certificates

# 程序目录
COPY code /imagebuild/code
COPY globle_config /imagebuild/globle_config
RUN mkdir -p /imagebuild/project

# 端口
EXPOSE 8080

# 工作目录
WORKDIR "/imagebuild/code/web"

# entrypoint
ENTRYPOINT ["/bin/sh", "/imagebuild/code/entrypoint.sh"]
```

* entrypoint.sh

```
#!/bin/sh
###################
# 镜像入口：
# 1) 初始化docker环境
# 2) 启动imagebuild进程提供服务
###################

# 更新docker daemon启动参数, 将harbor地址添加到insecure-registry中
UPDATE_DOCKER_DAEMON_PARAM_COMMAND="sed -i 's/\(--storage-driver=vfs \\\\\\)/\1\n                --insecure-registry=${HARBOR_ADDRESS} \\\\\n/g' /usr/local/bin/dockerd-entrypoint.sh"
# 更新
/bin/sh -c "${UPDATE_DOCKER_DAEMON_PARAM_COMMAND}"
# 启动docker daemon
nohup dockerd-entrypoint.sh > /tmp/imagebuild-console.log 2>&1 &
# 登陆harbor
docker login -u ${HARBOR_USER} -p ${HARBOR_PASSWORD} ${HARBOR_ADDRESS}
# 启动服务器
/imagebuild/code/web/main /imagebuild
```
