#!/bin/bash

checkImage(){
    images=$(docker images|grep "$1"|awk '{print $2}')
    if [[ -z  "$images" ]];then
	echo "image $1:$2 does not exists"
	exit 1
    fi
    for image in $images
    do
       #echo $image
       if [[ "$2" == "$image"  ]];then
  	  echo "image $1:$2 exists..."
	  return
       fi
    done
    echo "image $1:$2 does not exists"
    exit 1
}
echo "###################begin:"`date +%Y%m%d" "%H":"%M":"%S`
#打印参数
echo "param:"$*
if [ -z "$1" ]; then
    echo "mysql address is  empty! you should be set like this:"
    echo "sh get.sh mysql://DBUSER:@DBADDRESS/DBNAME?charset=utf8"
    exit 1
fi
#sed -i "s/mysql:\/\/root:@127.0.0.1\/octans?charset=utf8/$1/g" config.yml
if [ -z "$2" ]; then
    echo "must set jupit ssh key addr ,you should type like:"
    echo "sh get.sh mysqladdress http://JUPITERADDRESS:JUPITERPORT/v1/instance/sshkey"
    exit 1
fi
if [ -z "$3" ]; then
    echo "must set push jupit ssh key addr ,you should type like:"
    echo "sh get.sh mysqladdress getsshkeyurl JUPITERADDRESS:JUPITERPORT"
    exit 1
fi

if [ -z "$4" ]; then
    echo "must set push instance id ,you should type like:"
    echo "sh get.sh mysqladdress getsshkeyurl jupitaddress&port INSTANCE_ID"
    exit 1
fi

if [ -z "$5" ]; then
   echo "must set local ip"
   exit 1
fi

if [ -z "$6" ]; then
    echo "must set private registry address ,you should type like ip or domain:"
    exit 1
fi

#1、安装docker
echo "1、安装docker"
yum install -y docker

#2、修改docker配置
echo "2、修改docker配置"
#文件 /etc/sysconfig/docker 追加如下两行
echo 'OPTIONS="-g=/data0/docker -s=devicemapper --label idc=aliyun"' >> /etc/sysconfig/docker
echo 'INSECURE_REGISTRY="--insecure-registry docker.io --insecure-registry '$6'"' >> /etc/sysconfig/docker

#3、重新启动docker
echo "3、重新启动docker"
service docker restart

echo "4、下载octans-agent镜像"
#4、下载octans-agent镜像

docker pull registry.cn-beijing.aliyuncs.com/opendcp/octans-agent:latest

echo "5、检查镜像是否下载成功"
#5、检查镜像是否下载成功
checkImage registry.cn-beijing.aliyuncs.com/opendcp/octans-agent  latest


#6、启动octans-agent容器，并且修改配置(通过环境变量设置到容器内部)
echo "6、启动octans-agent容器，并且修改配置(通过环境变量设置到容器内部)"

hn=`hostname`
echo "hostname=" $hn
wc=$(grep "$5 $hn" /etc/hosts|wc -l )
if [ $wc -eq 0 ]; then
    echo "$5 $hn" >> /etc/hosts
fi

#环境变量设置
#mysql://root:12345@10.85.41.168:3306/octans?charset=utf8
#http://10.85.41.168:8083/v1/instance/sshkey/
#10.85.41.168:8083
#i-2zeen6mal4s9qvpqb4iq
#47.93.162.228
docker run -d -e "mysql_url=$1" -e "get_key_url=$2" -e "report_url=$3" -e "instance_id=$4"  --net=host --name octans-agent registry.cn-beijing.aliyuncs.com/opendcp/octans-agent:latest

echo "[DONE] --------------"
echo "###################end:"`date +%Y%m%d" "%H":"%M":"%S`
