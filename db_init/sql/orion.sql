CREATE DATABASE IF NOT EXISTS orion CHARACTER SET utf8 COLLATE utf8_general_ci;
USE orion;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.Cluster`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `cluster` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(50) NOT NULL DEFAULT '' UNIQUE ,
    `desc` varchar(255),
    `biz` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.Service`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `service` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(50) NOT NULL UNIQUE DEFAULT '' ,
    `desc` varchar(255),
    `service_type` varchar(255) NOT NULL DEFAULT '' ,
    `docker_image` varchar(255) NOT NULL DEFAULT '' ,
    `cluster_id` integer NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.Pool`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `pool` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(50) NOT NULL UNIQUE DEFAULT '' ,
    `desc` varchar(255),
    `vm_type` integer NOT NULL DEFAULT 0 ,
    `sd_id` integer NOT NULL DEFAULT 0 ,
    `tasks` varchar(255) NOT NULL DEFAULT '' ,
    `service_id` integer NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.ExecTask`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `exec_task` (
  `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `pool_id` integer NOT NULL,
  `type` varchar(50) NOT NULL DEFAULT 'expand',
  `exec_type` VARCHAR(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.CronItem`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `cron_item` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `exec_task_id` integer NOT NULL,
    `instance_num` integer ,
    `concurr_ratio` integer ,
    `concurr_num` integer ,
    `week_day` integer NOT NULL DEFAULT 0,
    `time` VARCHAR(255) NOT NULL,
    `ignore` tinyint(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.DependItem`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `depend_item` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `exec_task_id` integer NOT NULL,
    `pool_id` integer NOT NULL,
    `ratio` DOUBLE NOT NULL,
    `elastic_count` integer NOT NULL,
    `step_name` VARCHAR(255) NOT NULL,
    `ignore` tinyint(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.FlowImpl`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `flow_impl` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(50) NOT NULL DEFAULT ''  UNIQUE,
    `desc` varchar(255) NOT NULL DEFAULT '' ,
    `steps` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.Flow`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `flow` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(50) NOT NULL DEFAULT '',
    `status` integer NOT NULL DEFAULT 0 ,
    `pool_id` integer,
    `options` longtext NOT NULL,
    `impl_id` integer NOT NULL,
    `step_len` integer NOT NULL DEFAULT 0 ,
    `op_user` varchar(255) NOT NULL DEFAULT '' ,
    `flow_type` varchar(50) NOT NULL DEFAULT 'manual',
    `run_time` DOUBLE  NOT NULL DEFAULT 0.0 ,
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.NodeState`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `node_state` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `ip` varchar(255) NOT NULL DEFAULT '' ,
    `vm_id` varchar(255) NOT NULL DEFAULT '' ,
    `pool_id` integer NOT NULL,
    `flow_id` integer NOT NULL,
    `status` integer NOT NULL DEFAULT 0 ,
    `steps` longtext NOT NULL,
    `step_num` integer NOT NULL DEFAULT 0 ,
    `log` longtext NOT NULL,
    `last_op` varchar(255),
    `label` varchar(255),
    `step_run_time`  longtext NOT NULL,
    `run_time` DOUBLE  NOT NULL DEFAULT 0.0 ,
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL,
    `node_type` varchar(255) NOT NULL DEFAULT 'manual',
    `deleted`   tinyint(1) DEFAULT 0

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.RemoteStep`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `remote_step` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(50) NOT NULL DEFAULT ''  UNIQUE,
    `desc` varchar(255),
    `actions` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.RemoteAction`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `remote_action` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(50) NOT NULL DEFAULT ''  UNIQUE,
    `desc` varchar(255),
    `task_type` varchar(50) NOT NULL DEFAULT '',
    `params` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.RemoteActionImpl`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `remote_action_impl` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `type` varchar(50) NOT NULL DEFAULT '' ,
    `task_type` varchar(50) NOT NULL DEFAULT '',
    `template` longtext NOT NULL,
    `action_id` integer NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.Logs`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `logs` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `fid` int(10) NOT NULL,
  `message` text NOT NULL,
  `ctime` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志信息表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.RoleResource`
-- --------------------------------------------------
CREATE TABLE `role_resource` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL UNIQUE,
    `desc` varchar(200) NOT NULL DEFAULT '',
    `resource_type` varchar(200) NOT NULL COMMENT 'file var template meta task handle',
    `resource_content` longtext NOT NULL,
    `user` varchar(200) NOT NULL,
    `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 初始状态 1 变更中 2 变更完成',
    `type` tinyint(3) unsigned NOT NULL DEFAULT '1',
    `hidden` tinyint(3) unsigned NOT NULL DEFAULT '0',
    `create_time` datetime,
    `update_time` datetime,
    `template_file_path` varchar(512) NOT NULL DEFAULT '' COMMENT 'only tpl need',
    `template_file_perm` varchar(32) NOT NULL DEFAULT '',
    `template_file_owner` varchar(32) NOT NULL DEFAULT '',
    `associate_role` varchar(512) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.Role`
-- --------------------------------------------------
CREATE TABLE `role` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL UNIQUE,
    `desc` varchar(200) NOT NULL DEFAULT '',
    `role_file_path` varchar(512) NOT NULL DEFAULT '',
    `files` varchar(512) NOT NULL DEFAULT '',
    `handles` varchar(512) NOT NULL DEFAULT '',
    `meta` varchar(512) NOT NULL DEFAULT '',
    `tasks` varchar(512) NOT NULL DEFAULT '',
    `templates` varchar(512) NOT NULL DEFAULT '',
    `vars` varchar(512) NOT NULL DEFAULT '',
    `user` varchar(200) NOT NULL DEFAULT 'root',
    `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 初始状态 1 变更中 2 变更完成',
    `create_time` DATETIME,
    `update_time` DATETIME COMMENT '变更时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
-- DATA
-- --------------------------------------------------
LOCK TABLES `cluster` WRITE;
INSERT INTO `cluster` VALUES (1,'default_cluster','默认集群','1');
INSERT INTO `cluster` VALUES (2,'openstack_cluster','虚拟化集群','1');
UNLOCK TABLES;

LOCK TABLES `service` WRITE;
INSERT INTO `service` VALUES
    (1,'sd-nginx','服务发现-Nginx服务','nginx','-',1),
    (2,'my_server','my_server','Java','registry.cn-beijing.aliyuncs.com/opendcp/java-web:latest',1),
    (3,'controller_service','虚拟化控制节点服务','openstack','-',2),
    (4,'compute_service','虚拟化计算节点服务','openstack','-',2),
    (5,'storage_service','虚拟化存储节点服务','openstack','-',2);
UNLOCK TABLES;

LOCK TABLES `pool` WRITE;
INSERT INTO `pool` VALUES
    (1,'sd-nginx-aliyun','服务发现nginx',3,1,'{\"deploy\":3,\"expand\":1,\"shrink\":2}',1),
    (2,'my_server_nginx','使用nginx服务发现',3,1,'{\"deploy\":6,\"expand\":4,\"shrink\":5}',2),
    (3,'controller_pool','控制节点服务池',3,1,'{\"deploy\":6,\"expand\":4,\"shrink\":5}',3),
    (4,'compute_pool','计算节点服务池',3,1,'{\"deploy\":6,\"expand\":4,\"shrink\":5}',4),
    (5,'storage_pool','存储节点服务池',3,1,'{\"deploy\":6,\"expand\":4,\"shrink\":5}',5);
UNLOCK TABLES;

LOCK TABLES `flow_impl` WRITE;
INSERT INTO `flow_impl` VALUES
    (1,'expand_nginx','扩容nginx服务','[{"name":"create_vm","param_values":{"vm_type_id":3},"retry":{"retry_times":0,"ignore_error":false}},{"name":"install_nginx","param_values":{},"retry":{"retry_times":0,"ignore_error":false}},{"name":"check_port","param_values":{"check_port":80,"check_times":30},"retry":{"retry_times":0,"ignore_error":false}}]'),
    (2,'undeploy_nginx','缩容nginx服务','[{\"name\":\"return_vm\",\"param_values\":{\"vm_type_id\":1},\"retry\":{\"retry_times\":2,\"ignore_error\":false}}]'),
    (3,'noop','No op','[{\"name\":\"echo\",\"param_values\":{\"echo_word\":\"noop\"},\"retry\":{\"retry_times\":0,\"ignore_error\":true}}]'),
    (4,'expand_my_server','扩容my_server','[{\"name\":\"create_vm\",\"param_values\":{\"vm_type_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":false}},{\"name\":\"start_service\",\"param_values\":{\"host\":\"host\",\"name\":\"my_server\",\"tag\":\"registry.cn-beijing.aliyuncs.com/opendcp/java-web:latest \"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}},{\"name\":\"register\",\"param_values\":{\"service_discovery_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (5,'unexpand_my_server','缩容my_server','[{\"name\":\"unregister\",\"param_values\":{\"service_discovery_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":true}},{\"name\":\"stop_service\",\"param_values\":{\"name\":\"my_server\"},\"retry\":{\"retry_times\":0,\"ignore_error\":true}},{\"name\":\"return_vm\",\"param_values\":{\"vm_type_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (6,'upgrade_my_server','上线my_server','[{\"name\":\"stop_service\",\"param_values\":{\"name\":\"my_server\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}},{\"name\":\"start_service\",\"param_values\":{\"host\":\"host\",\"name\":\"my_server\",\"tag\":\"registry.cn-beijing.aliyuncs.com/opendcp/java-web:latest\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (7,'init_controller','controller初始化','[{\"name\":\"init_controller\",\"param_values\":{\"opendcp_host\":\"host_ip\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (8,'init_compute','compute初始化','[{\"name\":\"init_compute\",\"param_values\":{\"opendcp_host\":\"host_ip\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (9,'add-openstack-default-image','添加openstack缺省镜像','[{\"name\":\"add-default-image\",\"param_values\":{},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (10,'init_storage','storage初始化','[{\"name\":\"init_storage\",\"param_values\":{\"opendcp_host\":\"host_ip\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]');
UNLOCK TABLES;

LOCK TABLES `remote_action` WRITE;
INSERT INTO `remote_action` VALUES
    (1,'start_docker','启动docker','ansible_task','{\"host\":\"string\",\"name\":\"string\",\"tag\":\"string\"}'),
    (2,'check_port','检查端口','ansible_task','{\"check_port\":\"integer\",\"check_times\":\"integer\"}'),
    (3,'check_url','检测URL','ansible_task','{\"check_keyword\":\"string\",\"check_url\":\"string\"}'),
    (4,'stop_docker','停止Docker容器','ansible_task','{\"name\":\"string\"}'),
    (5,'echo','echo','ansible_task','{\"echo_word\":\"string\"}'),
    (6,'install_nginx','安装nginx','ansible_role','{}'),
    (7,'init_controller','初始化openstack控制节点','ansible_task','{\"opendcp_host\":\"string\"}'),
    (8,'init_compute','init_compute','ansible_task','{\"opendcp_host\":\"string\"}'),
    (9,'add-default-image','添加openstack Centos7缺省镜像','ansible_task','{}'),
    (10,'init_storage','init_storage','ansible_task','{\"opendcp_host\":\"string\"}');
UNLOCK TABLES;

LOCK TABLES `remote_action_impl` WRITE;
INSERT INTO `remote_action_impl` VALUES
    (1,'ansible','ansible_task','{\"action\":{\"content\":\"docker run -d --net=\\\"{{host}}\\\" --name {{name}} {{tag}} \",\"module\":\"longscript\"}}',1),
    (2,'ansible','ansible_task','{\"action\":{\"content\":\"# check port\\nTIMES={{check_times}}\\nPORT={{check_port}}\\nfor ((i=0;i\\u003c$TIMES;i++));\\ndo\\n\\techo \\\"check $PORT time $i ...\\\"\\n\\tres=`netstat -an | grep LISTEN | grep -e \\\"\\\\b$PORT\\\\b\\\"`\\n\\tif [ \\\"\\\" != \\\"$res\\\" ]; then\\n\\t\\techo \\\"OK\\\"\\n\\t\\texit 0\\n\\tfi\\n\\tsleep 5\\ndone\\necho \\\"error\\\" \\nexit 1\",\"module\":\"longscript\"}}',2),
    (3,'ansible','ansible_task','{\"action\":{\"content\":\"sleep 20\\nres=`curl -m 400 {{check_url}} | grep {{check_keyword}}`\\nif [ \\\"\\\" != \\\"$res\\\" ]; then\\n    echo \\\"OK\\\"\\n    exit 0\\nfi\\n\\necho \\\"check fails\\\"\\nexit 1\\n\",\"module\":\"longscript\"}}',3),
    (4,'ansible','ansible_task','{\"action\":{\"content\":\"cname={{name}}\\ncontainer=`docker ps|grep -w $cname`\\nif [ \\\"\\\" != \\\"$container\\\" ];then\\n    docker stop $cname\\nfi\\nsleep 5\\ncontainer=`docker ps -af status=exited|grep -w  $cname`\\nif [ \\\"\\\" != \\\"$container\\\" ];then\\n        docker rm $cname\\nfi\\nexit 0\",\"module\":\"longscript\"}}',4),
    (5,'ansible','ansible_task','{\"action\":{\"args\":\"echo {{echo_word}} \",\"module\":\"shell\"}}',5),
    (6,'ansible','ansible_role','{\"action\":{\"module\":\"role\"}}',6),
    (7,'ansible','ansible_task','{\"action\":{\"content\":\"docker rm -f oskfile\\ndocker pull registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\ndocker run --name=oskfile -tid registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\nrm -rf /tmp/oskfile\\nmkdir -p /tmp/oskfile\\ndocker cp oskfile:/data1/openstack /tmp/oskfile\\ncd /tmp/oskfile/openstack\\necho \'start\'\\nchmod +x init.sh\\nsh init.sh {{opendcp_host}} \\u003e /tmp/osk.log 2\\u003e\\u00261\\necho \'ok\'\\nrm -rf /tmp/oskfile\\ndocker rm -f oskfile\",\"module\":\"longscript\"}}',7),
    (8,'ansible','ansible_task','{\"action\":{\"content\":\"docker rm -f oskfile\\ndocker pull registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\ndocker run --name=oskfile -tid registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\nrm -rf /tmp/oskfile\\nmkdir -p /tmp/oskfile\\ndocker cp oskfile:/data1/openstack /tmp/oskfile\\ncd /tmp/oskfile/openstack\\necho \'start\'\\nchmod +x init_compute.sh\\nsh init_compute.sh {{opendcp_host}} \\u003e /tmp/osk.log 2\\u003e\\u00261\\necho \'ok\'\\nrm -rf /tmp/oskfile\\ndocker rm -f oskfile\",\"module\":\"longscript\"}}',8),
    (9,'ansible','ansible_task','{\"action\":{\"content\":\"docker rm -f oskfile\\ndocker pull registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\ndocker run --name=oskfile -tid registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\nrm -rf /tmp/oskfile\\nmkdir -p /tmp/oskfile\\ndocker cp oskfile:/data1/openstack /tmp/oskfile\\ncd /tmp/oskfile/openstack\\necho \'start\'\\nchmod +x add-default-image.sh\\nsh add-default-image.sh \\u003e /tmp/addimage.log 2\\u003e\\u00261\\necho \'ok\'\\nrm -rf /tmp/oskfile\\ndocker rm -f oskfile\",\"module\":\"longscript\"}}',9),
    (10,'ansible','ansible_task','{\"action\":{\"content\":\"docker rm -f oskfile\\ndocker pull registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\ndocker run --name=oskfile -tid registry.cn-beijing.aliyuncs.com/opendcp/openstack-scripts:latest\\nrm -rf /tmp/oskfile\\nmkdir -p /tmp/oskfile\\ndocker cp oskfile:/data1/openstack /tmp/oskfile\\ncd /tmp/oskfile/openstack\\necho \'start\'\\nchmod +x init_storage.sh\\nsh init_storage.sh {{opendcp_host}} \\u003e /tmp/osk.log 2\\u003e\\u00261\\necho \'ok\'\\nrm -rf /tmp/oskfile\\ndocker rm -f oskfile\",\"module\":\"longscript\"}}',10);
UNLOCK TABLES;

LOCK TABLES `remote_step` WRITE;
INSERT INTO `remote_step` VALUES
    (3,'echo','echo','[\"echo\"]'),
    (4,'install_nginx','安装nginx','[\"install_nginx\"]'),
    (8,'start_service','启动服务','[\"start_docker\"]'),
    (9,'stop_service','停止服务','[\"stop_docker\"]'),
    (10,'init_controller','controller初始化','[\"init_controller\"]'),
    (11,'init_compute','init_compute','[\"init_compute\"]'),
    (12,'add-default-image','添加openstack缺省镜像','[\"add-default-image\"]'),
    (13,'init_storage','init_storage','[\"init_storage\"]'),
    (14,'check_port','check_port','[\"check_port\"]');
UNLOCK TABLES;

LOCK TABLES  `role_resource` WRITE;
INSERT INTO `role_resource` VALUES
    (1,'task_install_nginx', '安装nginx', 'task', '---
- name: install rsync
  yum: name=rsync state=present

- name: create dest dir
  shell: ''mkdir -p /usr/local/nginx_conf/upstream''

- name: copy nginx main config file nginx.conf
  copy: src=/etc/ansible/roles/install_nginx_role/templates/nginx.conf.j2 dest=/usr/local/nginx_conf/nginx.conf

- name: copy nginx upstream config file default.upstream
  copy: src=/etc/ansible/roles/install_nginx_role/templates/default.upstream.j2 dest=/usr/local/nginx_conf/upstream/default.upstream

- name: start a docker container
  shell: ''docker run -d -v /usr/local/nginx_conf:/usr/local/nginx/conf --name opendcp_lb_ngx_ctn --net=host registry.cn-beijing.aliyuncs.com/opendcp/hubble-nginx:latest ./run.sh''
', '', 0, 0, 0, '2017-09-08 02:07:34', '2017-09-08 02:07:34', '', '', '', ''),
    (2, 'default.upstream', 'default.upstream', 'template', '#DEFAULT.UPSTREAM
upstream default_upstream{
		keepalive 60;
		server 127.0.0.1:8080 max_fails=0 fail_timeout=30s weight=20;
		check interval=1000 rise=3 fall=2 timeout=3000 type=http default_down=false;
		check_http_send "GET / HTTP/1.0\\r\\n\\r\\n";
		check_http_expect_alive http_2xx;
}
', '', 0, 0, 0, '2017-09-08 02:08:31', '2017-09-08 02:08:31', '', '', '', ''),
    (3, 'nginx.conf', 'nginx.conf', 'template', '#DEFAULT MAIN CONFIG FILE: nginx.conf
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
}
', '', 0, 0, 0, '2017-09-08 02:09:01', '2017-09-08 02:09:01', '', '', '', '');
UNLOCK TABLES;

LOCK TABLES `role` WRITE;
INSERT INTO `role` VALUES
    (1, 'install_nginx', 'role_task_nginx', '', '', '', '', '1', '2,3', '', '', 0, '2017-09-08 02:09:48', '2017-09-08 02:09:48');
UNLOCK TABLES;


