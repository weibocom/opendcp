<?php
//站点配置
define("MY_SITE_ROOT_PATH", "/");//站点根路径
define("MY_SITE_TITLE", "DCP Open");//站点全局title
define("MY_SITE_ALIAS", "OpenDCP混合云");//站点名称
define("MY_SITE_AUTHOR", "OpenDCP平台研发");//Author

//数据库配置信息
define('DB_NAME', 'web');
define('DB_CHARSET', 'utf8');
define('DB_HOST', 'db');
define('DB_PORT', 3306);
define('DB_USER', 'root');
define('DB_PW', '12345');

//LDAP配置信息
define('LDAP_HOST','');
define('LDAP_PORT',389);
define('LDAP_USER','');
define('LDAP_PASS','');
define('LDAP_BIND','');
define('LDAP_SEARCH','');

//Cookie配置   
define('COOKIE_DOMAIN', ''); //Cookie 作用域
define('COOKIE_PATH', '/'); //Cookie 作用路径
define('COOKIE_PRE', 'Open'); //Cookie 前缀，同一域名下安装多套Phpcms时，请修改Cookie前缀
define('COOKIE_TTL', 0); //Cookie 生命周期，0 表示随浏览器进程

//多云对接
define('CLOUD_DOMAIN', 'http://jupiter:8080');

//镜像仓库
define('REPOS_DOMAIN', 'http://harbor_ip:12380');
define('REOPS_AUTH', 'admin:Harbor12345');
//打包系统
define('PACKAGE_DOMAIN', 'http://imagebuild:8080');

//服务编排
define('LAYOUT_DOMAIN', 'http://orion:8080');

//服务发现
define('HUBBLE_DOMAIN', 'http://localhost:5555');
define('HUBBLE_APPKEY', '6741bc42-9e21-4763-977c-ac3a1fc0bdd8');

$_config=array();
//超级管理员
$_config['super']=array(
  'root'=>"管理员",
  'yingeng'=>"张隐耕",
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
$myFooter.='    '."\n";
$myFooter.='  </div>'."\n";
$myFooter.='  <div class="clearfix"></div>'."\n";
$myFooter.='</div>'."\n";
?>
