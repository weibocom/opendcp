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
$fMethod = 'GET';
$fUri = '/view/plugin';
$fParam = http_build_query($_GET);

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
  )
);
if($hasLimit){
  $retArr['msg'] = 'Param Error!';
  if($fUri){
    $retArr = $mySelf->get($myUser, $fMethod, $fUri , $fParam);
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
