<?php
/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
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
