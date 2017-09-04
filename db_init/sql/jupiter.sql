CREATE DATABASE IF NOT EXISTS jupiter CHARACTER SET utf8 COLLATE utf8_general_ci;
USE jupiter;

# Dump of table bill
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `bill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cluster_id` bigint(20) DEFAULT NULL,
  `costs` int(11) NOT NULL DEFAULT '0',
  `credit` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table cluster
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `cluster` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `provider` varchar(255) NOT NULL DEFAULT '',
  `lastest_part_num` int(11) NOT NULL DEFAULT '0',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `delete_time` datetime DEFAULT NULL,
  `cpu` int(11) NOT NULL DEFAULT '0',
  `ram` int(11) NOT NULL DEFAULT '0',
  `instance_type` varchar(255) NOT NULL DEFAULT '',
  `image_id` varchar(255) NOT NULL DEFAULT '',
  `post_script` varchar(255) NOT NULL DEFAULT '',
  `key_name` varchar(255) NOT NULL DEFAULT '',
  `network_id` bigint(20) NOT NULL,
  `zone_id` bigint(20) NOT NULL,
  `system_disk_category` varchar(255) NOT NULL DEFAULT '',
  `data_disk_size` int(11) NOT NULL DEFAULT '0',
  `data_disk_num` int(11) NOT NULL DEFAULT '0',
  `data_disk_category` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table instance_organization
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `instance_organization` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `organization_id` int(11) NOT NULL,
  `instance_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_organization_id` (`organization_id`),
  KEY `FK_instance_id` (`instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table instance
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `instance` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cluster_id` bigint(20) NOT NULL,
  `provider` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `cpu` int(11) NOT NULL DEFAULT '0',
  `ram` int(11) NOT NULL DEFAULT '0',
  `instance_id` varchar(255) NOT NULL DEFAULT '',
  `image_id` varchar(255) NOT NULL DEFAULT '',
  `instance_type` varchar(255) NOT NULL DEFAULT '',
  `key_name` varchar(255) NOT NULL DEFAULT '',
  `vpc_id` varchar(255) NOT NULL DEFAULT '',
  `subnet_id` varchar(255) NOT NULL DEFAULT '',
  `security_group_id` varchar(255) NOT NULL DEFAULT '',
  `region_id` varchar(255) NOT NULL DEFAULT '',
  `zone_id` varchar(255) NOT NULL DEFAULT '',
  `data_disk_num` int(11) NOT NULL DEFAULT '0',
  `data_disk_size` int(11) NOT NULL DEFAULT '0',
  `data_disk_category` varchar(255) NOT NULL DEFAULT '',
  `system_disk_category` varchar(255) NOT NULL DEFAULT '',
  `cost_way` varchar(255) NOT NULL DEFAULT '',
  `pre_buy_month` int(11) NOT NULL DEFAULT '0',
  `private_ip_address` varchar(255) NOT NULL DEFAULT '',
  `public_ip_address` varchar(255) NOT NULL DEFAULT '',
  `nat_ip_address` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `public_key` longtext,
  `private_key` longtext,
  `label` varchar(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table network
# ------------------------------------------------------------

CREATE TABLE `network` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vpc_id` varchar(255) DEFAULT '',
  `subnet_id` varchar(255) DEFAULT '',
  `security_group` varchar(255) DEFAULT NULL,
  `internet_charge_type` varchar(255) NOT NULL DEFAULT '',
  `internet_max_bandwidth_out` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `network` (`vpc_id`,`subnet_id`,`security_group`,`internet_charge_type`,`internet_max_bandwidth_out`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


# Dump of table organization
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `organization` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `desc` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table zone
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `zone` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `region_name` varchar(255) NOT NULL DEFAULT '',
  `zone_name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `region_name` (`region_name`,`zone_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# Dump of table log
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `log` (
  `instance_id` varchar(255) NOT NULL,
  `correlation_id` varchar(255) NOT NULL DEFAULT '',
  `ip` varchar(255) NOT NULL DEFAULT '',
  `message` longtext NOT NULL,
  PRIMARY KEY (`instance_id`,`correlation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# ------------------------------------------------------------
# DATA
# ------------------------------------------------------------
LOCK TABLES `network` WRITE;
INSERT INTO `network` VALUES 
    (1,'','',NULL,'PayByBandwidth',5),
    (2,'vpc-62a57d06','subnet-f303e897',NULL,'PayByBandwidth',5);
UNLOCK TABLES;

LOCK TABLES `zone` WRITE;
INSERT INTO `zone` VALUES (1,'cn-beijing','cn-beijing-c'),
                          (2, 'cn-beijing','cn-beijing-a'),
                          (3, 'cn-north','cn-north-1');
UNLOCK TABLES;

LOCK TABLES `cluster` WRITE;
INSERT INTO `cluster` VALUES 
    (1,'16Core16G经典网','aliyun',0,'',NOW(),NULL,16,16,'ecs.c2.medium','centos7u2_64_40G_cloudinit_20160728.raw','','key',1,1,'cloud_efficiency',100,1,'cloud_efficiency'),
    (2,'4Core8G经典网',  'aliyun',0,'',NOW(),NULL,4, 8, 'ecs.n2.large', 'centos7u2_64_40G_cloudinit_20160728.raw','','key',1,1,'cloud_efficiency',100,1,'cloud_efficiency'),
    (3,'1Core1G经典网',  'aliyun',0,'',NOW(),NULL,1, 1, 'ecs.n1.tiny',  'centos7u2_64_40G_cloudinit_20160728.raw','','key',1,1,'cloud_efficiency',100,1,'cloud_efficiency'),
    (4,'1Core-1Gib',    'aws',   0,'',NOW(),NULL,1, 1, 't2.micro',     'ami-3965b454',                           '','zhaowei9',2,3,'standard'     ,100,1,'standard'        )
    ;
UNLOCK TABLES;

LOCK TABLES `bill` WRITE;
INSERT INTO `bill` VALUES 
    (1,1,0,10),
    (2,2,0,0),
    (3,3,0,10),
    (4,4,0,10);
UNLOCK TABLES;


# Dump of table detail
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `instance_number` longtext NOT NULL,
  `running_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `instance_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_id` VARCHAR(255) NOT NULL DEFAULT '',
  `correlation_id` varchar(255) NOT NULL DEFAULT '',
  `cluster_id` bigint(20) NOT NULL,
  `instance_id` VARCHAR(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `status` INT(11) NOT NULL DEFAULT '0',
  `err_log`  VARCHAR(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;