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



function isLogin(){
  if(isset($_SESSION['open_user']))  return true;
  return false;
}

function isSuper(){
  global $mySuper,$myUser;
  if(isset($mySuper[$myUser])) return true;
  return false;
}

function logRecord($arr){
  global $db;
  if($arr){
	$sqlKey='';$sqlValue='';
	foreach($arr as $k=>$v){
		$sqlKey.=($sqlKey)?','.$k:$k;
    if($k=='t_code') $v=addslashes($v);
		$sqlValue.=($sqlValue)?",'".$v."'":"'".$v."'";
	}
	$sql='INSERT INTO user_log ('.$sqlKey.') VALUES ('.$sqlValue.');';
	$query=$db->query($sql);
  }
  return false;
}

/* 时间格式化 */
function format_date($time){
  $t=time()-$time;
  $f=array(
    '31536000'=>'年',
    '2592000'=>'个月',
    '604800'=>'星期',
    '86400'=>'天',
    '3600'=>'小时',
    '60'=>'分钟',
    '1'=>'秒'
  );
  foreach ($f as $k=>$v)    {
    if (0 !=$c=floor($t/(int)$k)) {
      return $c.$v.'前';
    }
  }
}

function date_inteval($begin,$end){
  $begin=($begin)?strtotime($begin):date('U');
  $end=($end)?strtotime($end):date('U');
  $t=$end-$begin;
  $f=array(
    '31536000'=>'年',
    '2592000'=>'个月',
    '604800'=>'星期',
    '86400'=>'天',
    '3600'=>'小时',
    '60'=>'分钟',
    '1'=>'秒'
  );
  foreach ($f as $k=>$v)    {
    if (0 !=$c=floor($t/(int)$k)) {
      return ' '.$c.$v;
    }
  }
}

$db = new mysqli(DB_HOST, DB_USER, DB_PW, DB_NAME);
/* check connection */
if (mysqli_connect_errno()) {
  printf("Connect failed: %s\n", mysqli_connect_error());
  exit();
}
$db->set_charset(DB_CHARSET);

$mySite=(strrpos($mySite,'/')===strlen($mySite)-1)?substr($mySite,0,strlen($mySite)-1):$mySite;
?>
