#!/bin/sh
###################
# 镜像入口：
# 1) 初始化docker环境
# 2) 启动imagebuild进程提供服务
###################

# 判断是否已经替换Harbor地址
harbor=`grep "${HARBOR_ADDRESS}" /usr/local/bin/dockerd-entrypoint.sh|wc -l`
if [ $harbor -eq 0 ]; then
    # 更新docker daemon启动参数, 将harbor地址添加到insecure-registry中
    UPDATE_DOCKER_DAEMON_PARAM_COMMAND="sed -i 's/\(--storage-driver=vfs \\\\\\)/\1\n                --insecure-registry=${HARBOR_ADDRESS} \\\\\n/g' /usr/local/bin/dockerd-entrypoint.sh"
    # 更新
    /bin/sh -c "${UPDATE_DOCKER_DAEMON_PARAM_COMMAND}"
fi

# 启动docker daemon
rm -rf /var/run/docker.pid
nohup dockerd-entrypoint.sh > /tmp/imagebuild-console.log 2>&1 &
# 登陆harbor
docker login -u ${HARBOR_USER} -p ${HARBOR_PASSWORD} ${HARBOR_ADDRESS}
# 启动服务器
/imagebuild/code/web/main /imagebuild
