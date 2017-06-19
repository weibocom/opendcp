# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 10.236.67.183 (MySQL 5.5.52-MariaDB)
# Database: hubble
# Generation Time: 2017-06-15 05:48:42 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE IF NOT EXISTS hubble CHARACTER SET utf8 COLLATE utf8_general_ci;
USE hubble;

# Dump of table tbl_hubble_alteration_history
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_alteration_history`;

CREATE TABLE `tbl_hubble_alteration_history` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `global_id` varchar(60) NOT NULL,
  `type` varchar(11) NOT NULL DEFAULT 'sync',
  `task_id` int(11) NOT NULL,
  `task_name` varchar(60) NOT NULL DEFAULT '',
  `channel` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_alteration_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_alteration_type`;

CREATE TABLE `tbl_hubble_alteration_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) NOT NULL DEFAULT 'NGINX',
  `content` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_alteration_type` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_alteration_type` DISABLE KEYS */;

INSERT INTO `tbl_hubble_alteration_type` (`id`, `name`, `type`, `content`, `create_time`, `update_time`, `opr_user`, `biz_id`)
VALUES
	(1,'default_service_name','NGINX','{"group_id":"1","name":"default.upstream","port":"8080","weight":"20","script_id":"2"}','2016-11-15 22:16:50','2016-11-15 22:16:50','system',0);

/*!40000 ALTER TABLE `tbl_hubble_alteration_type` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_log`;

CREATE TABLE `tbl_hubble_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `global_id` varchar(60) NOT NULL DEFAULT '',
  `module` varchar(20) NOT NULL,
  `log_info` text NOT NULL,
  `level` varchar(10) NOT NULL DEFAULT '',
  `url` varchar(512) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_conf_main
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_nginx_conf_main`;

CREATE TABLE `tbl_hubble_nginx_conf_main` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `version` int(11) NOT NULL,
  `content` text NOT NULL,
  `unit_id` int(11) NOT NULL,
  `deprecated` tinyint(1) NOT NULL DEFAULT '1',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_nginx_conf_main` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_nginx_conf_main` DISABLE KEYS */;

INSERT INTO `tbl_hubble_nginx_conf_main` (`id`, `name`, `version`, `content`, `unit_id`, `deprecated`, `create_time`, `opr_user`, `biz_id`)
VALUES
	(1,'nginx.conf',1,'#DEFAULT MAIN CONFIG FILE: nginx.conf
error_log /usr/local/nginx/logs/error.log notice;
pid /usr/local/nginx/logs/nginx.pid;
events {
    worker_connections 1024;
    use   epoll;
}
http {
    default_type  application/octet-stream;
    log_format  main  ''$remote_addr - $remote_user [$time_local] "$request" ''
                      ''$status $body_bytes_sent "$http_referer" ''
					  ''"$http_user_agent" "$http_x_forwarded_for" $request_time'';
    include upstream/*.upstream;
    #DEFAULT VHOST
    server {
        listen       80;
        server_name  0.0.0.0:80;
        location / {
            proxy_pass http://default_upstream;
        }
        location /status {
            check_status;
            access_log on;
        }
        access_log  logs/default_vhost.log main;
    }
    access_log logs/access.log  main;
}',1,0,'2016-11-15 22:10:25','system',0);

/*!40000 ALTER TABLE `tbl_hubble_nginx_conf_main` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_nginx_conf_upstream
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_nginx_conf_upstream`;

CREATE TABLE `tbl_hubble_nginx_conf_upstream` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `group_id` int(11) NOT NULL,
  `is_consul` tinyint(1) NOT NULL DEFAULT '0',
  `deprecated` tinyint(1) NOT NULL DEFAULT '1',
  `release_id` int(11) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_nginx_conf_upstream` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_nginx_conf_upstream` DISABLE KEYS */;

INSERT INTO `tbl_hubble_nginx_conf_upstream` (`id`, `name`, `content`, `group_id`, `is_consul`, `deprecated`, `release_id`, `create_time`, `update_time`, `opr_user`, `biz_id`)
VALUES
	(1,'default.upstream','upstream default_upstream{
	keepalive 4;
        server 127.0.0.1:8080 max_fails=0 fail_timeout=30s weight=20;
        check interval=1000 rise=3 fall=2 timeout=3000 type=http default_down=false;
        check_http_send "GET / HTTP/1.0";
        check_http_expect_alive http_2xx;
}',1,0,0,0,'2016-11-15 22:11:23','2016-11-15 22:11:23','system',0);

/*!40000 ALTER TABLE `tbl_hubble_nginx_conf_upstream` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_nginx_group
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_nginx_group`;

CREATE TABLE `tbl_hubble_nginx_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_nginx_group` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_nginx_group` DISABLE KEYS */;

INSERT INTO `tbl_hubble_nginx_group` (`id`, `name`, `opr_user`, `biz_id`, `create_time`)
VALUES
	(1,'default_group','',0,'1970-01-01 00:00:00');

/*!40000 ALTER TABLE `tbl_hubble_nginx_group` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_nginx_node
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_nginx_node`;

CREATE TABLE `tbl_hubble_nginx_node` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(15) NOT NULL DEFAULT '',
  `unit_id` int(11) NOT NULL,
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  KEY `unit_id_idx` (`unit_id`),
  CONSTRAINT `unit_id` FOREIGN KEY (`unit_id`) REFERENCES `tbl_hubble_nginx_unit` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_shell
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_nginx_shell`;

CREATE TABLE `tbl_hubble_nginx_shell` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_nginx_shell` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_nginx_shell` DISABLE KEYS */;

INSERT INTO `tbl_hubble_nginx_shell` (`id`, `name`, `desc`, `content`, `create_time`, `update_time`, `opr_user`, `biz_id`)
VALUES
	(1,'updateMainConf.sh','更新主配置脚本','#!/bin/bash
HUBBLE_UNIT_ID="{{HUBBLE_UNIT_ID}}"
HUBBLE_GROUP_ID="{{HUBBLE_GROUP_ID}}"
HUBBLE_RSYNC_HOST="{{HUBBLE_RSYNC_HOST}}"

HUBBLE_OUTFILE="/tmp/alterationMainConf.out"
cat > $HUBBLE_OUTFILE<< EOF
EOF

exec 1> $HUBBLE_OUTFILE  2> $HUBBLE_OUTFILE

confdir="/usr/local/nginx_conf"
time_echo(){
    echo `date +%F"-"%T`" "$*
}
#RSYNC MAIN CONFIG FILE nginx.conf FROM HUBBLE SERVER

time_echo rsync -argtv "${HUBBLE_RSYNC_HOST}/group_${HUBBLE_GROUP_ID}/unit_${HUBBLE_UNIT_ID}/current/main/" $confdir/
rsync -argtv "${HUBBLE_RSYNC_HOST}/group_${HUBBLE_GROUP_ID}/unit_${HUBBLE_UNIT_ID}/current/main/" $confdir/

if [ $? -ne 0 ]; then
    time_echo "rsync file fail" && exit 1
fi

time_echo rsync nginx_conf successful...
#CONFIGRATION CHECK AND RELOAD
docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t
time_echo docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t
if [ $? -eq 0 ];then
    docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload
    if [ $? -ne 0 ]; then
        time_echo "reload nginx failed" && exit 1
    fi
    time_echo docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload
    exit 0
fi
time_echo "check nginx_conf failed…" && exit 1','1970-01-01 00:00:00','1970-01-01 00:00:00','system',0),
	(2,'updateUpstreamConf.sh','更新upstream配置脚本','#!/bin/bash

HUBBLE_OUTFILE="/tmp/alterationUpstream.out"

exec 1> $HUBBLE_OUTFILE  2> $HUBBLE_OUTFILE

HUBBLE_FILE_COUNT="{{HUBBLE_FILE_COUNT}}"
HUBBLE_FILE_NAMES="{{HUBBLE_FILE_NAMES}}"
HUBBLE_GROUP_ID="{{HUBBLE_GROUP_ID}}"
HUBBLE_HOST="{{HUBBLE_HOST}}"
HUBBLE_BIZ="{{HUBBLE_BIZ}}"

cat > $HUBBLE_OUTFILE<< EOF
EOF
time_echo(){
    echo `date +%F"-"%T`" "$*
}

time_echo "HUBBLE_FILE_COUNT: $HUBBLE_FILE_COUNT"
time_echo "HUBBLE_FILE_NAMES: $HUBBLE_FILE_NAMES"
time_echo "HUBBLE_GROUP_ID: $HUBBLE_GROUP_ID"
time_echo "HUBBLE_HOST: $HUBBLE_HOST"
time_echo "HUBBLE_HOST:$HUBBLE_BIZ"


# nginx upstream dir
conf_dir="/usr/local/nginx_conf/upstream"
pgrep nginx
rs=`echo $?`
if [ $rs -ne 0 ];then
  time_echo "The nginx process is not exist!"
  exit 1
fi
if [ ! -d $conf_dir ];then
  time_echo "This server is not nginx server!"
  exit 1
fi
cd $conf_dir
pwd
for i in `echo "$HUBBLE_FILE_NAMES" | tr '','' ''\n''`; do
           time_echo "deal with $i"
           url="${HUBBLE_HOST}&group_id=${HUBBLE_GROUP_ID}&name=$i"
           header="x-biz-id:${HUBBLE_BIZ}"
           time_echo "wget --header=$header $url" -O "${i}.new"
           wget --header="$header" "$url" -O "${i}.new"
           if [ $? -ne 0 ]; then
               rm -f "*.upstream.new"
               exit 1
           fi
done
for i in `find . -name *.upstream.new`; do
           mv "$i" "`echo ${i%.*}`"
done
ng_conf_ok=`docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t 2>&1 |grep successful |wc -l`
ng_conf_error=`docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t 2>&1|grep failed -A 1`
system_version=`cat /etc/redhat-release | awk ''{print $(NF-1)}''`
if [[ $ng_conf_ok -eq 1 ]];then
    time_echo "--------begin to start Nginx————"
    docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload
    result=`echo $?`
    time_echo "--------Reload Nginx Finish--------" && exit 0
else
    time_echo "----Nginx config file wrong----"
    time_echo "$ng_conf_error" && exit 1
fi','1970-01-01 00:00:00','1970-01-01 00:00:00','system',0);

/*!40000 ALTER TABLE `tbl_hubble_nginx_shell` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_nginx_unit
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_nginx_unit`;

CREATE TABLE `tbl_hubble_nginx_unit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `group_id` int(11) NOT NULL,
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `group_id_idx` (`group_id`),
  CONSTRAINT `group_id` FOREIGN KEY (`group_id`) REFERENCES `tbl_hubble_nginx_group` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_nginx_unit` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_nginx_unit` DISABLE KEYS */;

INSERT INTO `tbl_hubble_nginx_unit` (`id`, `name`, `group_id`, `opr_user`, `create_time`, `biz_id`)
VALUES
	(1,'default_unit',1,'system','2016-11-15 22:09:44',0);

/*!40000 ALTER TABLE `tbl_hubble_nginx_unit` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_nginx_version
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_nginx_version`;

CREATE TABLE `tbl_hubble_nginx_version` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `version` int(11) NOT NULL,
  `deprecated` tinyint(1) NOT NULL,
  `unit_id` int(11) NOT NULL,
  `files` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `type` varchar(255) DEFAULT NULL,
  `is_release` tinyint(1) NOT NULL,
  `release_id` varchar(255) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_secure_appkey
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_secure_appkey`;

CREATE TABLE `tbl_hubble_secure_appkey` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `appkey` varchar(256) NOT NULL DEFAULT '',
  `name` varchar(60) NOT NULL DEFAULT '',
  `describe` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_secure_appkey` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_secure_appkey` DISABLE KEYS */;

INSERT INTO `tbl_hubble_secure_appkey` (`id`, `appkey`, `name`, `describe`)
VALUES
	(1,'048EA63AFD38D629993819EF980DE5AE','default_appkey','默认配置，建议删除');

/*!40000 ALTER TABLE `tbl_hubble_secure_appkey` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_secure_interface
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_secure_interface`;

CREATE TABLE `tbl_hubble_secure_interface` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `addr` varchar(1024) NOT NULL,
  `desc` varchar(256) NOT NULL DEFAULT '',
  `method` varchar(10) DEFAULT 'GET',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_secure_interface` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_secure_interface` DISABLE KEYS */;

INSERT INTO `tbl_hubble_secure_interface` (`id`, `addr`, `desc`, `method`)
VALUES
	(1,'/v1/secure/appkey/list_interface/','接口列表','GET'),
	(2,'/v1/adaptor/auto_alteration/add/','注册','POST'),
	(3,'/v1/adaptor/auto_alteration/remove/','注销','POST'),
	(4,'/v1/adaptor/auto_alteration/check_state/','检测','GET'),
	(5,'/v1/tools/tool/nginx_init/','初始化','POST');

/*!40000 ALTER TABLE `tbl_hubble_secure_interface` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tbl_hubble_secure_oprlog
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_secure_oprlog`;

CREATE TABLE `tbl_hubble_secure_oprlog` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `module` varchar(256) NOT NULL DEFAULT '',
  `operation` varchar(256) NOT NULL DEFAULT '',
  `opr_time` datetime NOT NULL,
  `appkey` varchar(256) NOT NULL DEFAULT '',
  `user` varchar(60) NOT NULL DEFAULT '',
  `biz_id` int(11) unsigned NOT NULL DEFAULT '0',
  `args` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_secure_privileges
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tbl_hubble_secure_privileges`;

CREATE TABLE `tbl_hubble_secure_privileges` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `appkey_id` int(11) unsigned NOT NULL,
  `interface_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tbl_hubble_secure_privileges` WRITE;
/*!40000 ALTER TABLE `tbl_hubble_secure_privileges` DISABLE KEYS */;

INSERT INTO `tbl_hubble_secure_privileges` (`id`, `appkey_id`, `interface_id`)
VALUES
	(1,1,1),
	(2,1,2),
	(3,1,3),
	(4,1,4),
	(5,1,5);

/*!40000 ALTER TABLE `tbl_hubble_secure_privileges` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
