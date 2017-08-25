-- MySQL dump 10.13  Distrib 5.7.9, for osx10.9 (x86_64)
--
-- Host: 127.0.0.1    Database: open
-- ------------------------------------------------------
-- Server version	5.5.47

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `member`
--


CREATE DATABASE IF NOT EXISTS web CHARACTER SET utf8 COLLATE utf8_general_ci;
USE web;

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE IF NOT EXISTS `member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `en` varchar(50) NOT NULL DEFAULT '',
  `cn` varchar(20) NOT NULL DEFAULT '',
  `type` varchar(20) NOT NULL DEFAULT '',
  `mail` varchar(50) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL,
  `pw` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` VALUES (1,'root','管理员','local','root@weibo.com',0,'21232f297a57a5a743894a0e4a801fc3');
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nav_bar`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE IF NOT EXISTS `nav_bar` (
  `nb_id` int(11) NOT NULL,
  `nb_fid` int(11) NOT NULL,
  `nb_name` varchar(20) NOT NULL DEFAULT '',
  `nb_href` varchar(200) NOT NULL DEFAULT '',
  `nb_target` varchar(20) NOT NULL DEFAULT '',
  `nb_desc` varchar(50) NOT NULL DEFAULT '',
  `nb_sort` int(11) NOT NULL,
  `nb_icon` varchar(20) NOT NULL DEFAULT '',
  `nb_new` varchar(20) NOT NULL DEFAULT '',
  `nb_status` int(11) NOT NULL,
  PRIMARY KEY (`nb_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nav_bar`
--

LOCK TABLES `nav_bar` WRITE;
/*!40000 ALTER TABLE `nav_bar` DISABLE KEYS */;
INSERT INTO `nav_bar` VALUES (0,0,'DCP','#','_self','DCP',0,'','',1),(1,0,'多云对接','for_cloud/','_self','多云对接',1,'fa fa-cloud','',1),(2,0,'镜像市场','for_image/','_self','镜像市场',2,'fa fa-file-zip-o','',1),(3,0,'服务编排','for_layout/','_self','服务编排',3,'fa fa-clone','',1),(4,0,'服务发现','for_hubble/','_self','服务发现',4,'fa fa-eye','',1),(5,0,'虚拟化管理','for_openstack/','_self','虚拟化管理',5,'fa fa-file-zip-o','',1),(101,1,'机型模板','for_cloud/org.php','_self','模板和配额管理',101,'','',0),(102,1,'机器管理','for_cloud/ecs.php','_self','云端机器管理',102,'','',0),(201,2,'镜像仓库','for_repos/repos.php','_self','Docker镜像管理',201,'','',0),(202,2,'打包系统','for_repos/package.php','_self','打包配置管理和镜像构建',202,'','',0),(301,3,'集群管理','for_layout/cluster.php','_self','业务线管理',301,'','',0),(302,3,'服务管理','for_layout/service.php','_self','服务、服务池、节点管理',302,'','',0),(303,3,'任务管理','for_layout/task.php','_self','任务列表和模板',303,'','',0),(304,3,'远程命令','for_layout/action.php','_self','远程命令和命令组管理',304,'','',0),(305,3,'任务详情','for_layout/task_detail.php','_self','任务详细信息',305,'','',2),(401,4,'服务注册','for_hubble/balance.php','_self','服务发现类型管理',401,'','',0),(403,4,'七层Nginx','for_hubble/nginx_group.php','_self','七层负载均衡',403,'','',0),(404,4,'阿里云SLB','for_hubble/slb.php','_self','SLB控制台',404,'','',0),(410,4,'脚本管理','for_hubble/shell.php','_self','发布用脚本管理',410,'','',0),(411,4,'授权管理','for_hubble/appkey.php','_self','AppKey和API授权管理',411,'','',0),(412,4,'操作日志','for_hubble/oprlog.php','_self','操作日志和发布变更历史',412,'','',0),(40301,403,'Nginx分组','for_hubble/nginx_group.php','_self','Nginx集群管理',40301,'','',0),(40302,403,'Nginx单元','for_hubble/nginx_unit.php','_self','Upstream配置和Nginx服务池及节点管理',40302,'','',0),(40303,403,'Nginx配置','for_hubble/nginx_conf.php','_self','主配置和vHost配置',40303,'','',1),(50001,5,'物理节点管理','for_openstack/machine.php','_self','物理节点管理',50001,'','',0),(50002,5,'初始化任务','for_openstack/initlist.php','_self','初始化任务',50002,'','',0),(50003,5,'存储节点管理','for_openstack/storage.php','_self','存储节点管理',50003,'','',0),(90001,0,'系统管理','admin/','_self','系统管理',90001,'fa fa-cogs','',1),(90081,90001,'导航树管理','admin/navbar.php','_self','导航树管理',90081,'','',0),(90091,90001,'用户管理','admin/user.php','_self','用户增删改查',90091,'','',0),(90093,90001,'用户日志','admin/user_log.php','_self','查看用户日志',90093,'','',0);
/*!40000 ALTER TABLE `nav_bar` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `user_log`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE IF NOT EXISTS `user_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `t_time` varchar(20) NOT NULL DEFAULT '',
  `t_user` varchar(50) NOT NULL DEFAULT '',
  `t_module` varchar(20) NOT NULL DEFAULT '',
  `t_action` varchar(20) NOT NULL DEFAULT '',
  `t_desc` varchar(512) NOT NULL,
  `t_code` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_log`
--

LOCK TABLES `user_log` WRITE;
/*!40000 ALTER TABLE `user_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_log` ENABLE KEYS */;
UNLOCK TABLES;



CREATE TABLE `node_init` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL DEFAULT '',
  `type` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `disk_name` varchar(64) NOT NULL DEFAULT '',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `node_init_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL DEFAULT '',
  `task_id` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `data` text NOT NULL,
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_task` (`task_id`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `keydata` (
  `datakey` varchar(128) NOT NULL,
  `datacontent` longtext NOT NULL,
  PRIMARY KEY (`datakey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO keydata (datakey, datacontent) VALUES ('controller_ip', '""');

CREATE TABLE `compute_power` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `data` text NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-28 11:12:57


