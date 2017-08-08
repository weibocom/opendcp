# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Database: hubble
# Generation Time: 2016-10-25 09:00:39 +0000
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

# Dump of table tbl_hubble_log
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `global_id` varchar(60) NOT NULL DEFAULT '',
  `module` varchar(20) NOT NULL,
  `log_info` text NOT NULL,
  `level` varchar(10) NOT NULL DEFAULT '',
  `url` varchar(512) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# Dump of table tbl_hubble_alteration_history
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_alteration_history` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `global_id` varchar(60) NOT NULL,
  `type` varchar(11) NOT NULL DEFAULT 'sync',
  `task_id` int(11) NOT NULL,
  `task_name` varchar(60) NOT NULL DEFAULT '',
  `channel` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


# Dump of table tbl_hubble_alteration_type
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_alteration_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) NOT NULL DEFAULT 'NGINX',
  `content` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_conf_main
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_nginx_conf_main` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `version` int(11) NOT NULL,
  `content` text NOT NULL,
  `unit_id` int(11) NOT NULL,
  `deprecated` tinyint(1) NOT NULL DEFAULT '1',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_conf_upstream
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_nginx_conf_upstream` (
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# Dump of table tbl_hubble_nginx_group
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_nginx_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_node
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_nginx_node` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(15) NOT NULL DEFAULT '',
  `unit_id` int(11) NOT NULL,
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  KEY `unit_id_idx` (`unit_id`),
  CONSTRAINT `unit_id` FOREIGN KEY (`unit_id`) REFERENCES `tbl_hubble_nginx_unit` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_shell
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_nginx_shell` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_unit
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_nginx_unit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `group_id` int(11) NOT NULL,
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  KEY `group_id_idx` (`group_id`),
  CONSTRAINT `group_id` FOREIGN KEY (`group_id`) REFERENCES `tbl_hubble_nginx_group` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_nginx_version
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_nginx_version` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `version` int(11) NOT NULL,
  `deprecated` tinyint(1) NOT NULL,
  `unit_id` int(11) NOT NULL,
  `files` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `opr_user` varchar(255) NOT NULL DEFAULT '',
  `type` varchar(255) DEFAULT NULL,
  `is_release` tinyint(1) NOT NULL,
  `release_id` varchar(255) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_secure_appkey
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_secure_appkey` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `appkey` varchar(256) NOT NULL DEFAULT '',
  `name` varchar(60) NOT NULL DEFAULT '',
  `describe` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_secure_interface
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_secure_interface` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `addr` varchar(1024) NOT NULL,
  `desc` varchar(256) NOT NULL DEFAULT '',
  `method` varchar(10) DEFAULT 'GET',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_secure_oprlog
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_secure_oprlog` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `module` varchar(256) NOT NULL DEFAULT '',
  `operation` varchar(256) NOT NULL DEFAULT '',
  `opr_time` datetime NOT NULL,
  `appkey` varchar(256) NOT NULL DEFAULT '',
  `user` varchar(60) NOT NULL DEFAULT '',
  `args` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tbl_hubble_secure_privileges
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `tbl_hubble_secure_privileges` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `appkey_id` int(11) unsigned NOT NULL,
  `interface_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


/* INSERT DEFAULT DATA */
INSERT INTO `tbl_hubble_alteration_type` VALUES (1,'default_service_name','NGINX','{\"group_id\":\"1\",\"name\":\"default.upstream\",\"port\":\"8080\",\"weight\":\"20\",\"script_id\":\"2\"}','2016-11-15 22:16:50','2016-11-15 22:16:50','system');
INSERT INTO `tbl_hubble_nginx_conf_main` VALUES (1,'nginx.conf',1,'#DEFAULT MAIN CONFIG FILE: nginx.conf\nerror_log /usr/local/nginx/logs/error.log notice;\npid /usr/local/nginx/logs/nginx.pid;\nevents {\n    worker_connections 1024;\n    use   epoll;\n}\nhttp {\n    default_type  application/octet-stream;\n    log_format  main  \'$remote_addr - $remote_user [$time_local] \"$request\" \'\n                      \'$status $body_bytes_sent \"$http_referer\" \'\n					  \'\"$http_user_agent\" \"$http_x_forwarded_for\" $request_time\';\n    include upstream/*.upstream;\n    #DEFAULT VHOST\n    server {\n        listen       80;\n        server_name  0.0.0.0:80;\n        location / {\n            proxy_pass http://default_upstream;\n        }\n        location /status {\n            check_status;\n            access_log on;\n        }\n        access_log  logs/default_vhost.log main;\n    }\n    access_log logs/access.log  main;\n}\n',1,0,'2016-11-15 22:10:25','system');
INSERT INTO `tbl_hubble_nginx_conf_upstream` VALUES (1,'default.upstream','upstream default_upstream{\n        keepalive 1;\n        server 127.0.0.1:8080 max_fails=0 fail_timeout=30s weight=20;\n        check interval=1000 rise=3 fall=2 timeout=3000 type=http default_down=false;\n        check_http_send \"GET / HTTP/1.0\\r\\n\\r\\n\";\n        check_http_expect_alive http_2xx;\n}\n',1,0,0,0,'2016-11-15 22:11:23','2016-11-15 22:11:23','system');
INSERT INTO `tbl_hubble_nginx_group` VALUES (1,'default_group','','1970-01-01 00:00:00');
INSERT INTO `tbl_hubble_nginx_shell` (`id`, `name`, `desc`, `content`, `create_time`, `update_time`, `opr_user`)
VALUES
  (1,'updateMainConf.sh','更新主配置脚本','#!/bin/bash\nHUBBLE_UNIT_ID=\"{{HUBBLE_UNIT_ID}}\"\nHUBBLE_GROUP_ID=\"{{HUBBLE_GROUP_ID}}\"\nHUBBLE_RSYNC_HOST=\"{{HUBBLE_RSYNC_HOST}}\"\n\nHUBBLE_OUTFILE=\"/tmp/alterationMainConf.out\"\ncat > $HUBBLE_OUTFILE<< EOF\nEOF\n\nexec 1> $HUBBLE_OUTFILE  2> $HUBBLE_OUTFILE\n\nconfdir=\"/usr/local/nginx_conf\"\ntime_echo(){\n    echo `date +%F\"-\"%T`\" \"$*\n}\n#RSYNC MAIN CONFIG FILE nginx.conf FROM HUBBLE SERVER\n\ntime_echo rsync -argtv \"${HUBBLE_RSYNC_HOST}/group_${HUBBLE_GROUP_ID}/unit_${HUBBLE_UNIT_ID}/current/main/\" $confdir/\nrsync -argtv \"${HUBBLE_RSYNC_HOST}/group_${HUBBLE_GROUP_ID}/unit_${HUBBLE_UNIT_ID}/current/main/\" $confdir/ \n\nif [ $? -ne 0 ]; then\n    time_echo \"rsync file fail\" && exit 1\nfi\n\ntime_echo rsync nginx_conf successful...\n#CONFIGRATION CHECK AND RELOAD\ndocker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t\ntime_echo docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t\nif [ $? -eq 0 ];then\n    docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload\n    if [ $? -ne 0 ]; then\n        time_echo \"reload nginx failed\" && exit 1\n    fi\n    time_echo docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload\n    exit 0\nfi\ntime_echo \"check nginx_conf failed…\" && exit 1','2016-11-28 12:52:17','2016-11-28 12:52:17','system'),
  (2,'updateUpstreamConf.sh','更新upstream配置脚本','#!/bin/bash\n\nHUBBLE_OUTFILE=\"/tmp/alterationUpstream.out\"\n\nexec 1> $HUBBLE_OUTFILE  2> $HUBBLE_OUTFILE\n\nHUBBLE_FILE_COUNT=\"{{HUBBLE_FILE_COUNT}}\"\nHUBBLE_FILE_NAMES=\"{{HUBBLE_FILE_NAMES}}\"\nHUBBLE_GROUP_ID=\"{{HUBBLE_GROUP_ID}}\"\nHUBBLE_HOST=\"{{HUBBLE_HOST}}\"\ncat > $HUBBLE_OUTFILE<< EOF\nEOF\ntime_echo(){\n    echo `date +%F\"-\"%T`\" \"$*\n}\n\ntime_echo \"HUBBLE_FILE_COUNT: $HUBBLE_FILE_COUNT\"\ntime_echo \"HUBBLE_FILE_NAMES: $HUBBLE_FILE_NAMES\"\ntime_echo \"HUBBLE_GROUP_ID: $HUBBLE_GROUP_ID\"\ntime_echo \"HUBBLE_HOST: $HUBBLE_HOST\"\n\n\n# nginx upstream dir\nconf_dir=\"/usr/local/nginx_conf/upstream\"\npgrep nginx\nrs=`echo $?`\nif [ $rs -ne 0 ];then\n  time_echo \"The nginx process is not exist!\"\n  exit 1\nfi\nif [ ! -d $conf_dir ];then\n  time_echo \"This server is not nginx server!\"\n  exit 1\nfi\ncd $conf_dir\npwd\nfor i in `echo \"$HUBBLE_FILE_NAMES\" | tr \',\' \'\\n\'`; do\n           time_echo \"deal with $i\"\n           url=\"${HUBBLE_HOST}&group_id=${HUBBLE_GROUP_ID}&name=$i\"\n           time_echo \"wget $url\" -O \"${i}.new\"\n           wget \"$url\" -O \"${i}.new\"\n           if [ $? -ne 0 ]; then\n               rm -f \"*.upstream.new\"\n               exit 1\n           fi\ndone\nfor i in `find . -name *.upstream.new`; do\n           mv \"$i\" \"`echo $i | sed \'s/\\(.*\\)\\.new/\\1/g\'`\"\ndone\n\nfor i in {1..5}\ndo\ntime_echo \"----Exec $i times----\"\nng_conf_ok=`docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t 2>&1 |grep successful |wc -l`\nif [[ $ng_conf_ok -eq 1 ]];then\n    time_echo \"--------begin to start Nginx————\"\n    docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload\n    result=`echo $?`\n    echo $result\n    time_echo \"--------Reload Nginx Finish--------\" && exit 0\nfi\nsleep 1s\ndone\ntime_echo \"----Nginx config file wrong----\"\ntime_echo \"$ng_conf_error\" && exit 1','2016-11-28 12:51:23','2016-11-28 12:51:23','root');

INSERT INTO `tbl_hubble_nginx_unit` VALUES (1,'default_unit',1,'system','2016-11-15 22:09:44');
INSERT INTO `tbl_hubble_secure_appkey` VALUES (1,'048EA63AFD38D629993819EF980DE5AE','default_appkey','默认配置，建议删除');
INSERT INTO `tbl_hubble_secure_interface` VALUES (1,'/v1/secure/appkey/list_interface/','接口列表','GET'),(2,'/v1/adaptor/auto_alteration/add/','注册','POST'),(3,'/v1/adaptor/auto_alteration/remove/','注销','POST'),(4,'/v1/adaptor/auto_alteration/check_state/','检测','GET');
INSERT INTO `tbl_hubble_secure_privileges` VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4);
