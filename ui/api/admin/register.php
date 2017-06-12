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

session_start();
header('Content-type: application/json');
include_once('../../include/config.inc.php');
include_once('../../include/function.php');
include_once('../../include/reg.php');
$thisClass=$reg;

class myself{

  function update($action = 'add', $param = array()){
    global $thisClass;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if(!empty($param)){
      switch($action) {
        case 'add':
          $param['status'] = 99;
          $param['reg_time'] = date('Y-m-d H:i:s');
          $ret = $thisClass->add($param);
          break;
      }
    }
    return $ret;
  }
}
$mySelf=new myself();

/*权限检查*/
$pageForSuper = false;//当前页面是否需要管理员权限
$hasLimit = ($pageForSuper)?isSuper($myUser):true;
$myAction = (isset($_POST['action'])&&!empty($_POST['action']))?trim($_POST['action']):((isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'');
$myIndex = (isset($_POST['index'])&&!empty($_POST['index']))?trim($_POST['index']):((isset($_GET['index'])&&!empty($_GET['index']))?trim($_GET['index']):'');
$myPage = (isset($_POST['page'])&&intval($_POST['page'])>0)?intval($_POST['page']):((isset($_GET['page'])&&intval($_GET['page'])>0)?intval($_GET['page']):1);
$myPageSize = (isset($_POST['pagesize'])&&intval($_POST['pagesize'])>0)?intval($_POST['pagesize']):((isset($_GET['pagesize'])&&intval($_GET['pagesize'])>0)?intval($_GET['pagesize']):$myPageSize);

$fIdx=(isset($_POST['fIdx'])&&!empty($_POST['fIdx']))?trim($_POST['fIdx']):((isset($_GET['fIdx'])&&!empty($_GET['fIdx']))?trim($_GET['fIdx']):'');

$myJson=(isset($_POST['data'])&&!empty($_POST['data']))?trim($_POST['data']):((isset($_GET['data'])&&!empty($_GET['data']))?trim($_GET['data']):'');
$arrJson=($myJson)?json_decode($myJson,true):array();
$logJson=$arrJson;
if(isset($logJson['pw'])) $logJson['pw']=md5('z_'+$logJson['pw']);

if(isset($arrJson) && !empty($arrJson)){
  //校验验证码
  if(!isset($arrJson['verification_code']) || $arrJson['verification_code']!==$_SESSION['verification_code']){
    $retArr = array(
      'code' => 1,
      'msg' => '验证码错误',
    );
    echo json_encode($retArr, JSON_UNESCAPED_UNICODE);
    exit;
  }
}else{
  $retArr = array(
    'code' => 1,
    'msg' => '非法请求'
  );
  echo json_encode($retArr, JSON_UNESCAPED_UNICODE);
  exit;
}

//记录操作日志
$logFlag = true;
$logDesc = '';
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => 'Guest',
  't_module' => '申请体验管理',
  't_action' => '',
  't_desc' => 'Resource:' . $_SERVER['REMOTE_ADDR'] . '.',
  't_code' => '传入：' . json_encode($logJson) . "\n\n",
);
//返回
$retArr = array(
  'code' => 1,
  'action' => $myAction,
);
if($hasLimit){
  $retArr['msg'] = 'Param Error!';
  switch($myAction){
    case 'insert':
      if(!$pageForSuper && $myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '添加';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) unset($arrJson['id']);
        if(isset($arrJson['verification_code'])) unset($arrJson['verification_code']);
        $retArr = $mySelf->update('add', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
  }
}else{
  $retArr['msg'] = 'Permission Denied!';
}
//记录日志
if($logFlag){
  $arrRecodeLog['t_desc'] = ($logDesc) ? $logDesc.', '.$arrRecodeLog['t_desc'] : $arrRecodeLog['t_desc'];
  $arrRecodeLog['t_code'] .= '返回：' . json_encode($retArr,JSON_UNESCAPED_UNICODE);
  if(empty($arrRecodeLog['t_action'])) $arrRecodeLog['t_action'] = $myAction;
  logRecord($arrRecodeLog);
}
//返回结果
if(isset($retArr['action']) && !empty($retArr['action'])) $retArr['action'] = $myAction;
if(isset($retArr['ret'])) unset($retArr['ret']);
echo json_encode($retArr, JSON_UNESCAPED_UNICODE);
?>
