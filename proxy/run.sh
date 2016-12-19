#!/bin/sh

cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#apt-get -y install wget

service php5-fpm restart
mkdir -p /data1/hubble/Application/Runtime/Hubble/Nginx/upstream
rm -rf /data1/hubble/Application/Runtime/*
chmod 777 -R /data1/hubble/Application/Runtime
rsync --daemon --config=/etc/rsyncd.conf &
nginx -g "daemon off;"
