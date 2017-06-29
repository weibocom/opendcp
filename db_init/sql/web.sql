CREATE DATABASE  IF NOT EXISTS `web` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `web`;
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
-- Table structure for table `biz`
--

DROP TABLE IF EXISTS `biz`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `biz` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `status` int(11) NOT NULL,
  PRIMARY KEY (`id`)
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `biz`
--

LOCK TABLES `biz` WRITE;
/*!40000 ALTER TABLE `biz` DISABLE KEYS */;
/*!40000 ALTER TABLE `biz` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member`
--

DROP TABLE IF EXISTS `member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `en` varchar(50) NOT NULL DEFAULT '',
  `cn` varchar(20) NOT NULL DEFAULT '',
  `type` varchar(20) NOT NULL DEFAULT '',
  `mobile` varchar(45) NOT NULL,
  `mail` varchar(50) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL,
  `pw` varchar(32) NOT NULL DEFAULT '',
  `biz_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `en_UNIQUE` (`en`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` VALUES (1,'root','管理员','local','','root@sina.com',0,'21232f297a57a5a743894a0e4a801fc3',0);
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nav_bar`
--

DROP TABLE IF EXISTS `nav_bar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `nav_bar` (
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
/*!INSERT INTO `nav_bar` VALUES (0,0,'DCP','#','_self','DCP',0,'','',1),(1,0,'多云对接','for_cloud/','_self','多云对接',1,'fa fa-cloud','',1),(2,0,'镜像市场','for_image/','_self','镜像市场',2,'fa fa-file-zip-o','',1),(3,0,'服务编排','for_layout/','_self','服务编排',3,'fa fa-clone','',1),(4,0,'服务发现','for_hubble/','_self','服务发现',4,'fa fa-eye','',1),(101,1,'机型模板','for_cloud/org.php','_self','模板和配额管理',101,'','',0),(102,1,'机器管理','for_cloud/ecs.php','_self','云端机器管理',102,'','',0),(103,1,'账号管理','for_cloud/account.php','_self','云厂商账号管理',103,'','',0),(201,2,'打包系统','for_repos/package.php','_self','打包配置管理和镜像构建',201,'','',0),(202,2,'镜像仓库','for_repos/repos.php','_self','Docker镜像管理',202,'','',0),(301,3,'集群管理','for_layout/cluster.php','_self','业务线管理',301,'','',0),(302,3,'服务管理','for_layout/service.php','_self','服务、服务池、节点管理',302,'','',0),(303,3,'任务管理','for_layout/task.php','_self','任务列表和模板',303,'','',0),(304,3,'远程命令','for_layout/action.php','_self','远程命令和命令组管理',304,'','',0),(305,3,'任务详情','for_layout/task_detail.php','_self','任务详细信息',305,'','',2),(401,4,'服务注册','for_hubble/balance.php','_self','服务发现类型管理',401,'','',0),(403,4,'七层Nginx','for_hubble/nginx_group.php','_self','七层负载均衡',403,'','',0),(404,4,'阿里云SLB','for_hubble/slb.php','_self','SLB控制台',404,'','',1),(410,4,'脚本管理','for_hubble/shell.php','_self','发布用脚本管理',410,'','',0),(411,4,'授权管理','for_hubble/appkey.php','_self','AppKey和API授权管理',411,'','',0),(412,4,'操作日志','for_hubble/oprlog.php','_self','操作日志和发布变更历史',412,'','',0),(40301,403,'Nginx分组','for_hubble/nginx_group.php','_self','Nginx集群管理',40301,'','',0),(40302,403,'Nginx单元','for_hubble/nginx_unit.php','_self','Upstream配置和Nginx服务池及节点管理',40302,'','',0),(40303,403,'Nginx配置','for_hubble/nginx_conf.php','_self','主配置和vHost配置',40303,'','',1),(90001,0,'系统管理','admin/','_self','系统管理',90001,'fa fa-cogs','',1),(90061,90001,'体验申请管理','admin/reg.php','_self','体验申请管理',90061,'','',0),(90071,90001,'业务方管理','admin/biz.php','_self','业务方管理',90071,'','',0),(90081,90001,'导航树管理','admin/navbar.php','_self','导航树管理',90081,'','',0),(90091,90001,'用户管理','admin/user.php','_self','用户增删改查',90091,'','',0),(90093,90001,'用户日志','admin/user_log.php','_self','查看用户日志',90093,'','',0);*/;
INSERT INTO `nav_bar` VALUES (0,0,'DCP','#','_self','DCP',0,'','',1),(1,0,'多云对接','for_cloud/','_self','多云对接',1,'fa fa-cloud','',1),(2,0,'镜像市场','for_image/','_self','镜像市场',2,'fa fa-file-zip-o','',1),(3,0,'服务编排','for_layout/','_self','服务编排',3,'fa fa-clone','',1),(4,0,'服务发现','for_hubble/','_self','服务发现',4,'fa fa-eye','',1),(101,1,'机型模板','for_cloud/org.php','_self','模板和配额管理',101,'','',0),(102,1,'机器管理','for_cloud/ecs.php','_self','云端机器管理',102,'','',0),(103,1,'账号管理','for_cloud/account.php','_self','云厂商账号管理',103,'','',0),(201,2,'打包系统','for_repos/package.php','_self','打包配置管理和镜像构建',201,'','',0),(202,2,'镜像仓库','for_repos/repos.php','_self','Docker镜像管理',202,'','',0),(301,3,'集群管理','for_layout/cluster.php','_self','业务线管理',301,'','',0),(302,3,'服务管理','for_layout/service.php','_self','服务、服务池、节点管理',302,'','',0),(303,3,'任务管理','for_layout/task.php','_self','任务列表和模板',303,'','',0),(304,3,'远程命令','for_layout/action.php','_self','远程命令和命令组管理',304,'','',0),(305,3,'任务详情','for_layout/task_detail.php','_self','任务详细信息',305,'','',2),(401,4,'服务注册','for_hubble/balance.php','_self','服务发现类型管理',401,'','',0),(403,4,'七层Nginx','for_hubble/nginx_group.php','_self','七层负载均衡',403,'','',0),(410,4,'脚本管理','for_hubble/shell.php','_self','发布用脚本管理',410,'','',0),(411,4,'授权管理','for_hubble/appkey.php','_self','AppKey和API授权管理',411,'','',0),(412,4,'操作日志','for_hubble/oprlog.php','_self','操作日志和发布变更历史',412,'','',0),(40301,403,'Nginx分组','for_hubble/nginx_group.php','_self','Nginx集群管理',40301,'','',0),(40302,403,'Nginx单元','for_hubble/nginx_unit.php','_self','Upstream配置和Nginx服务池及节点管理',40302,'','',0),(40303,403,'Nginx配置','for_hubble/nginx_conf.php','_self','主配置和vHost配置',40303,'','',1),(90001,0,'系统管理','admin/','_self','系统管理',90001,'fa fa-cogs','',1),(90061,90001,'体验申请管理','admin/reg.php','_self','体验申请管理',90061,'','',0),(90071,90001,'业务方管理','admin/biz.php','_self','业务方管理',90071,'','',0),(90081,90001,'导航树管理','admin/navbar.php','_self','导航树管理',90081,'','',0),(90091,90001,'用户管理','admin/user.php','_self','用户增删改查',90091,'','',0),(90093,90001,'用户日志','admin/user_log.php','_self','查看用户日志',90093,'','',0);
/*!40000 ALTER TABLE `nav_bar` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reg`
--

DROP TABLE IF EXISTS `reg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reg` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `en` varchar(45) NOT NULL,
  `cn` varchar(45) NOT NULL,
  `mobile` varchar(45) NOT NULL,
  `mail` varchar(45) NOT NULL,
  `pw` varchar(45) NOT NULL,
  `biz` varchar(45) NOT NULL,
  `status` int(11) NOT NULL,
  `reg_time` varchar(45) NOT NULL,
  `audit_time` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `en_UNIQUE` (`en`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reg`
--

LOCK TABLES `reg` WRITE;
/*!40000 ALTER TABLE `reg` DISABLE KEYS */;
/*!40000 ALTER TABLE `reg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_log`
--

DROP TABLE IF EXISTS `user_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_log` (
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
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-06-16 10:45:09
