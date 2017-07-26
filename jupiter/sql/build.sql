# Dump of table bill
# ------------------------------------------------------------

CREATE TABLE `bill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cluster_id` bigint(20) DEFAULT NULL,
  `costs` int(11) NOT NULL DEFAULT '0',
  `credit` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table cluster
# ------------------------------------------------------------

CREATE TABLE `cluster` (
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
  'falvor_id' varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table instance_organization
# ------------------------------------------------------------

CREATE TABLE `instance_organization` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `organization_id` int(11) NOT NULL,
  `instance_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_organization_id` (`organization_id`),
  KEY `FK_instance_id` (`instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table instance
# ------------------------------------------------------------

CREATE TABLE `instance` (
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
  'tenant_id' varchar(255) NOT NULL DEFAULT '',
  'user_id'   varchar(255) NOT NULL DEFAULT '',
  'name'    varchar(255) NOT NULL DEFAULT '',
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

CREATE TABLE `organization` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `desc` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table replication
# ------------------------------------------------------------

CREATE TABLE `replication` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `part_num` int(11) NOT NULL DEFAULT '0',
  `cluster_id` int(11) NOT NULL,
  `instance_id` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table zone
# ------------------------------------------------------------

CREATE TABLE `zone` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `region_name` varchar(255) NOT NULL DEFAULT '',
  `zone_name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `region_name` (`region_name`,`zone_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table log
# ------------------------------------------------------------

CREATE TABLE `log` (
  `instance_id` varchar(255) NOT NULL,
  `correlation_id` varchar(255) NOT NULL DEFAULT '',
  `ip` varchar(255) NOT NULL DEFAULT '',
  `message` longtext NOT NULL,
  PRIMARY KEY (`instance_id`,`correlation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
