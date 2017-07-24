
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
--  Table Structure for `weibo.com/opendcp/orion/models.Node`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `node` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `ip` varchar(255),
    `vm_id` varchar(255),
    `status` integer NOT NULL DEFAULT 0 ,
    `pool_id` integer NOT NULL,
    `node_type` varchar(255) NOT NULL DEFAULT 'manual'
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
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.FlowBatch`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `flow_batch` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `flow_id` integer NOT NULL,
    `status` integer NOT NULL DEFAULT 0 ,
    `step` integer NOT NULL DEFAULT 0 ,
    `nodes` longtext NOT NULL,
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
    `corr_id` varchar(255),
    `node_id` integer NOT NULL DEFAULT 0,
    `pool_id` integer NOT NULL,
    `flow_id` integer NOT NULL,
    `batch_id` integer,
    `status` integer NOT NULL DEFAULT 0 ,
    `steps` longtext NOT NULL,
    `step_num` integer NOT NULL DEFAULT 0 ,
    `log` longtext NOT NULL,
    `last_op` varchar(255),
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL
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
    `params` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.RemoteActionImpl`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `remote_action_impl` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `type` varchar(50) NOT NULL DEFAULT '' ,
    `template` longtext NOT NULL,
    `action_id` integer NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------
--  Table Structure for `weibo.com/opendcp/orion/models.Logs`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `logs` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `fid` int(10) NOT NULL,
  `batch_id` int(10) NOT NULL DEFAULT '0',
  `correlation_id` varchar(20) NOT NULL DEFAULT '0' COMMENT '全局id',
  `message` text NOT NULL,
  `ctime` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志信息表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------
-- DATA
-- --------------------------------------------------
LOCK TABLES `cluster` WRITE;
INSERT INTO `cluster` VALUES (1,'default_cluster','默认集群','1');
UNLOCK TABLES;

LOCK TABLES `service` WRITE;
INSERT INTO `service` VALUES
    (1,'sd-nginx','服务发现-Nginx服务','nginx','-',1),
    (2,'my_server','my_server','Java','registry.cn-beijing.aliyuncs.com/opendcp/java-web:latest',1);
UNLOCK TABLES;

LOCK TABLES `pool` WRITE;
INSERT INTO `pool` VALUES
    (1,'sd-nginx-aliyun','服务发现nginx',3,1,'{\"deploy\":3,\"expand\":1,\"shrink\":2}',1),
    (2,'my_server_nginx','使用nginx服务发现',3,1,'{\"deploy\":6,\"expand\":4,\"shrink\":5}',2);
UNLOCK TABLES;

LOCK TABLES `flow_impl` WRITE;
INSERT INTO `flow_impl` VALUES
    (1,'expand_nginx','扩容nginx服务','[{\"name\":\"create_vm\",\"param_values\":{\"vm_type_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":false}},{\"name\":\"install_nginx\",\"param_values\":{\"check_port\":80,\"check_times\":30,\"eth\":\"eth1\",\"octans_host\":\"host_ip\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (2,'undeploy_nginx','缩容nginx服务','[{\"name\":\"return_vm\",\"param_values\":{\"vm_type_id\":1},\"retry\":{\"retry_times\":2,\"ignore_error\":false}}]'),
    (3,'noop','No op','[{\"name\":\"echo\",\"param_values\":{\"echo_word\":\"noop\"},\"retry\":{\"retry_times\":0,\"ignore_error\":true}}]'),
    (4,'expand_my_server','扩容my_server','[{\"name\":\"create_vm\",\"param_values\":{\"vm_type_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":false}},{\"name\":\"start_service\",\"param_values\":{\"host\":\"host\",\"name\":\"my_server\",\"tag\":\"registry.cn-beijing.aliyuncs.com/opendcp/java-web:latest \"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}},{\"name\":\"register\",\"param_values\":{\"service_discovery_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (5,'unexpand_my_server','缩容my_server','[{\"name\":\"unregister\",\"param_values\":{\"service_discovery_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":true}},{\"name\":\"stop_service\",\"param_values\":{\"name\":\"my_server\"},\"retry\":{\"retry_times\":0,\"ignore_error\":true}},{\"name\":\"return_vm\",\"param_values\":{\"vm_type_id\":1},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]'),
    (6,'upgrade_my_server','上线my_server','[{\"name\":\"stop_service\",\"param_values\":{\"name\":\"my_server\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}},{\"name\":\"start_service\",\"param_values\":{\"host\":\"host\",\"name\":\"my_server\",\"tag\":\"registry.cn-beijing.aliyuncs.com/opendcp/java-web:latest\"},\"retry\":{\"retry_times\":0,\"ignore_error\":false}}]');
UNLOCK TABLES;

LOCK TABLES `remote_action` WRITE;
INSERT INTO `remote_action` VALUES
    (1,'start_docker','启动docker','{\"host\":\"string\",\"name\":\"string\",\"tag\":\"string\"}'),
    (2,'check_port','检查端口','{\"check_port\":\"integer\",\"check_times\":\"integer\"}'),
    (3,'check_url','检测URL','{\"check_keyword\":\"string\",\"check_url\":\"string\"}'),
    (4,'stop_docker','停止Docker容器','{\"name\":\"string\"}'),
    (5,'echo','echo','{\"echo_word\":\"string\"}'),
    (6,'install_nginx','安装nginx','{\"eth\":\"string\",\"octans_host\":\"string\"}');
UNLOCK TABLES;

LOCK TABLES `remote_action_impl` WRITE;
INSERT INTO `remote_action_impl` VALUES
    (1,'ansible','{\"action\":{\"content\":\"docker run -d --net=\\\"{{host}}\\\" --name {{name}} {{tag}} \",\"module\":\"longscript\"}}',1),
    (2,'ansible','{\"action\":{\"content\":\"# check port\\nTIMES={{check_times}}\\nPORT={{check_port}}\\nfor ((i=0;i\\u003c$TIMES;i++));\\ndo\\n\\techo \\\"check $PORT time $i ...\\\"\\n\\tres=`netstat -an | grep LISTEN | grep -e \\\"\\\\b$PORT\\\\b\\\"`\\n\\tif [ \\\"\\\" != \\\"$res\\\" ]; then\\n\\t\\techo \\\"OK\\\"\\n\\t\\texit 0\\n\\tfi\\n\\tsleep 5\\ndone\\necho \\\"error\\\" \\nexit 1\",\"module\":\"longscript\"}}',2),
    (3,'ansible','{\"action\":{\"content\":\"sleep 20\\nres=`curl -m 400 {{check_url}} | grep {{check_keyword}}`\\nif [ \\\"\\\" != \\\"$res\\\" ]; then\\n    echo \\\"OK\\\"\\n    exit 0\\nfi\\n\\necho \\\"check fails\\\"\\nexit 1\\n\",\"module\":\"longscript\"}}',3),
    (4,'ansible','{\"action\":{\"content\":\"cname={{name}}\\ncontainer=`docker ps|grep -w $cname`\\nif [ \\\"\\\" != \\\"$container\\\" ];then\\n    docker stop $cname\\nfi\\nsleep 5\\ncontainer=`docker ps -af status=exited|grep -w  $cname`\\nif [ \\\"\\\" != \\\"$container\\\" ];then\\n        docker rm $cname\\nfi\\nexit 0\",\"module\":\"longscript\"}}',4),
    (5,'ansible','{\"action\":{\"args\":\"echo {{echo_word}} \",\"module\":\"shell\"}}',5),
    (6,'ansible','{\"action\":{\"content\":\"#!/bin/sh\\n\\n# get ip address\\nIP=`ifconfig {{eth}} | grep -w inet | awk \'{print $2}\'`\\necho \\\"IP is $IP\\\"\\n\\n# run role\\necho \\\"Deploy nginx on $IP ...\\\"\\nNOW=`date +\\\"%Y%m%d-%H%M%S\\\"`\\ncurl -l -H \\\"Content-type: application/json\\\" -H \\\"X-CORRELATION-ID: $IP-$NOW\\\" -H \\\"X-SOURCE: orion\\\" -X POST \\\\\\n    -d  \\\"{\\\\\\\"tasks\\\\\\\": [\\\\\\\"hubble-nginx\\\\\\\"], \\\\\\\"name\\\\\\\": \\\\\\\"$IP-$NOW\\\\\\\", \\\\\\\"fork_num\\\\\\\":5, \\\\\\\"tasktype\\\\\\\": \\\\\\\"ansible_role\\\\\\\", \\\\\\\"nodes\\\\\\\": [\\\\\\\"$IP\\\\\\\"], \\\\\\\"user\\\\\\\": \\\\\\\"root\\\\\\\"}\\\" \\\\\\n    http://$IP:8000/api/parallel_run\\n \",\"module\":\"longscript\"}}',6)
    ;
UNLOCK TABLES;

LOCK TABLES `remote_step` WRITE;
INSERT INTO `remote_step` VALUES
    (3,'echo','echo','[\"echo\"]'),
    (4,'install_nginx','安装nginx','[\"install_nginx\",\"check_port\"]'),
    (8,'start_service','启动服务','[\"start_docker\"]'),
    (9,'stop_service','停止服务','[\"stop_docker\"]');
UNLOCK TABLES;


