<?php
/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */


//站点配置
define("MY_SITE_ROOT_PATH", "/");//站点根路径
define("MY_SITE_TITLE", "DCP Open");//站点全局title
define("MY_SITE_ALIAS", "OpenDCP混合云");//站点名称
define("MY_SITE_AUTHOR", "OpenDCP平台研发");//Author

//数据库配置信息
define('DB_NAME', 'open');
define('DB_CHARSET', 'utf8');
define('DB_HOST', '127.0.0.1');
define('DB_PORT', 3306);
define('DB_USER', 'xxxxxx');
define('DB_PW', 'xxxxxx');

//LDAP配置信息
define('LDAP_HOST','127.0.0.1');
define('LDAP_PORT',389);
define('LDAP_USER','xxxxxx');
define('LDAP_PASS','xxxxxx');
define('LDAP_BIND','OU=people,DC=sina,DC=com,DC=cn');
define('LDAP_SEARCH','OU=people,DC=sina,DC=com,DC=cn');

//Cookie配置
define('COOKIE_DOMAIN', ''); //Cookie 作用域
define('COOKIE_PATH', '/'); //Cookie 作用路径
define('COOKIE_PRE', 'Open'); //Cookie 前缀，同一域名下安装多套Phpcms时，请修改Cookie前缀
define('COOKIE_TTL', 0); //Cookie 生命周期，0 表示随浏览器进程

//多云对接
define('CLOUD_DOMAIN', 'http://127.0.0.1:8083');

//镜像仓库
define('REPOS_DOMAIN', 'http://127.0.0.1:12381');
define('REOPS_AUTH', 'xxxxxx:xxxxxx');

//打包系统
define('PACKAGE_DOMAIN', 'http://127.0.0.1:8084');

//服务编排
define('LAYOUT_DOMAIN', 'http://127.0.0.1:8081');

//服务发现
define('HUBBLE_DOMAIN', 'http://127.0.0.1:5555');
define('HUBBLE_APPKEY', '6741bc42-9e21-4763-977c-ac3a1fc0bdd8');

$_config=array();
//超级管理员
$_config['super']=array(
  'root'=>"管理员",
);

$mySite=MY_SITE_ROOT_PATH;
$mySiteTitle=MY_SITE_TITLE;
$mySiteAlias=MY_SITE_ALIAS;
$myAuthor=MY_SITE_AUTHOR;
$myLdapHost=LDAP_HOST;
$myLdapPort=LDAP_PORT;
$myLdapUser=LDAP_USER;
$myLdapPass=LDAP_PASS;
$myLdapBind=LDAP_BIND;
$myLdapSearch=LDAP_SEARCH;

//分页
$myPageSize=20;//每页数据量
$myPage=(isset($_GET['page'])&&!empty($_GET['page']))?intval($_GET['page']):1;//当前页码
$myPageCount=1;//总页数

//页脚
$myFooter='<footer>'."\n";
$myFooter.='  <div class="pull-right">'."\n";
$myFooter.='     '."\n";
$myFooter.='  </div>'."\n";
$myFooter.='  <div class="clearfix"></div>'."\n";
$myFooter.='</div>'."\n";
?>
