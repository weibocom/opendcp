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


header('Content-type: application/json');
include_once('../../include/config.inc.php');
include_once('../../include/function.php');
include_once('../../include/func_session.php');
include_once('../../include/package.php');
$thisClass = $package;

class myself{

  function get($myUser = '', $method='GET', $uri='' , $param = ''){
    global $thisClass;
    $ret = $thisClass->proxyCurl($myUser, $method, $uri, $param);
    return $ret;
  }

}
$mySelf=new myself();

/*权限检查*/
$pageForSuper = false;//当前页面是否需要管理员权限
$hasLimit = ($pageForSuper)?isSuper($myUser):true;
$fMethod = 'POST';
$fUri = '/api/plugin/extension/interface';
$fParam = http_build_query($_POST);
$fIdx=(isset($_POST['fIdx'])&&!empty($_POST['fIdx']))?trim($_POST['fIdx']):((isset($_GET['fIdx'])&&!empty($_GET['fIdx']))?trim($_GET['fIdx']):'');

//记录操作日志
$logFlag = false;
$logDesc = 'FAILED';
$arrJson = array(
  'method' => $fMethod,
  'uri' => $fUri,
  'param' => $fParam,
);
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => $myUser,
  't_module' => '打包代理',
  't_action' => $fUri,
  't_desc' => 'Resource:' . $_SERVER['REMOTE_ADDR'] . '.',
  't_code' => '传入：' . json_encode($arrJson,JSON_UNESCAPED_UNICODE) . "\n\n",
);
//返回
$retArr = array(
  'code' => 1,
  'proxy' => array(
    'method' => $fMethod,
    'uri' => $fUri,
    'param' => $fParam,
  ),
  'filter' => $fIdx,
);
if($hasLimit){
  $retArr['msg'] = 'Param Error!';
  if($fUri){
    $retArr = $mySelf->get($myUser, $fMethod, $fUri , $fParam);
    if($fIdx){
      if(isset($retArr['data'])){
        $arrFilter=array();
        foreach($retArr['data'] as $v){
          if(strpos($v,'/'.$fIdx.'/')!==false) $arrFilter[]=$v;
        }
        $retArr['data']=$arrFilter;
      }
    }
  }
}else{
  $retArr['msg'] = 'Permission Denied!';
}
//记录日志
if($retArr['code']>0) $logFlag=true;
if($logFlag){
  $arrRecodeLog['t_desc'] = 'FAILED, '.$arrRecodeLog['t_desc'];
  $arrRecodeLog['t_code'] .= '返回：' . json_encode($retArr,JSON_UNESCAPED_UNICODE);
  logRecord($arrRecodeLog);
}
//返回结果
if($retArr['code']>0){
  echo json_encode($retArr, JSON_UNESCAPED_UNICODE);
}else{
  echo $retArr['content'];
}
?>
