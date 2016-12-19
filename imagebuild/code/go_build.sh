#!/bin/bash

##########################
# 编译所有插件以及主项目
##########################


# 定位到脚本所在目录
cd `dirname $0`

SRC_DIR=`pwd`

DOCKERFILE_PLUGINS_BASE=${SRC_DIR}/plugins/dockerfile/

TO_BUILD_FOLDER=(${DOCKERFILE_PLUGINS_BASE})

for folder in ${TO_BUILD_FOLDER[@]}
do
    for file in ${folder}*
    do 
        if test -d ${file}
        then
           echo "begin build plugin: "${file}
           plug=${file}"/"`basename ${file}`"_plugin"
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



