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
include_once('../../include/reg.php');
$thisClass=$reg;

class myself{

  function getList(){
    global $thisClass,$myPage,$myPageSize,$fIdx;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    $arrList = $thisClass->get('', $fIdx, $myPage, $myPageSize);
    if($arrList!==false){
      $ret = array(
        'code' => 0,
        'msg' => 'success',
        'page'=>$myPage,
        'pageSize'=>$myPageSize,
        'pageCount'=>0,
        'count'=>0,
        'filter'=>$fIdx,
        'title' => array(
          '#',
          '账号',
          '姓名',
          '手机',
          '邮箱',
          '业务方名称',
          '申请时间',
          '审核状态',
          '审核时间',
          '#',
        ),
        'content' => array(),
      );
      $ret['count'] = (int)$thisClass->getCount($fIdx);
      $ret['pageCount']=($ret['count']>0)?ceil($ret['count']/$ret['pageSize']):1;
      if($ret['page']>$ret['pageCount']) $ret['page']=1;
      $i=0;
      foreach($arrList as $k => $v){
        $i++;
        $tArr = array();
        $tArr['i'] = $i;
        foreach($v as $key => $value){
          $tArr[$key] = $value;
        }
        $ret['content'][] = $tArr;
      }
    }else{
      $ret['msg'] = 'db failed';
    }
    $ret['ret'] = $arrList;
    return $ret;
  }

  function getInfo($id = 0){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($id){
      if($arrList = $thisClass->get($id)){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'content' => $arrList[$id],
        );
      }
      $ret['ret'] = $arrList;
    }else{
      $ret['msg'] = 'id null';
    }
    return $ret;
  }

  function update($action = 'add', $param = array()){
    global $thisClass;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if(!empty($param)){
      switch($action) {
        case 'add':
          $param['status'] = 99;
          $param['reg_time'] = date('Y-m-d H:i:s');
          $param['audit_time'] = '';
          $ret = $thisClass->add($param);
          break;
        case 'update':
          $ret = $thisClass->update($param);
          break;
        case 'delete':
          $ret = $thisClass->delete($param);
          break;
      }
    }
    return $ret;
  }

  function audit($param = array()){
    global $thisClass,$biz,$user;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if(!empty($param)){
      $regId = (isset($param['id'])&&!empty($param['id'])) ? $param['id'] : 0;
      if(empty($regId)){
        $ret['msg'] = '无效申请ID';
        return $ret;
      }
      $param['audit_time'] = date('Y-m-d H:i:s');
      //更新审核状态
      $retAudit = $thisClass->update($param);
      //获取申请详情
      $retReg = $thisClass->get($regId);
      $regStatus = (isset($retReg[$regId]['status'])) ? $retReg[$regId]['status'] : 99;
      if($regStatus!==0&&$regStatus!==1){
        $ret['msg'] = '审核操作失败';
        return $ret;
      }
      if($regStatus===1){
        $ret['code'] = 0;
        $ret['msg'] = '审核操作成功';
        return $ret;
      }
      //注册公司名称
      $retBiz = $biz->add(['name' => $retReg[$regId]['biz'], 'status' => 0]);
      $bizId = (isset($retBiz['content'])&&!empty($retBiz['content'])) ? $retBiz['content'] : 0;
      if(empty($bizId)){
        $ret['msg'] = '公司名称注册失败';
        $param['status'] = 99;
        unset($param['audit_time']);
        $retAudit = $thisClass->update($param); //回滚审批状态
        return $ret;
      }
      //注册用户信息
      $paramUser = [
        'en' => $retReg[$regId]['en'],
        'cn' => $retReg[$regId]['cn'],
        'type' => 'local',
        'mobile' => $retReg[$regId]['mobile'],
        'mail' => $retReg[$regId]['mail'],
        'status' => 0,
        'pw' => $retReg[$regId]['pw'],
        'biz_id' => $bizId,
      ];
      $retUser = $user->add($paramUser, false);
      $userId = (isset($retUser['content'])&&!empty($retUser['content'])) ? $retUser['content'] : 0;
      if(empty($userId)){
        $ret['msg'] = '用户信息写入失败';
        unset($param['audit_time']);
        $param['status'] = 99;
        $retAudit = $thisClass->update($param); //回滚审批状态
        $retBiz = $biz->delete(['id' => $bizId]); //删除公司信息
        return $ret;
      }
      //记录初始化结果
      $reInit = '初始化结果: <br>';
      //通知多云对接模块初始化
      $retCloud = $this->initCloud($bizId, $retReg[$regId]['biz'], 0);
      $reInit .= ($retCloud['code'] === 0) ? '1.多云对接模块: 成功; <br>' : '1.多云对接模块: <span class="text-danger">失败</span>; <br>';
      $ret['init']['cloud'] = $retCloud;
      //通知镜像市场模块初始化
      $retRepos = $this->initRepos($bizId, $retReg[$regId]['biz'], 0);
      $reInit .= ($retRepos['code'] === 0) ? '2.镜像仓库模块: 成功; <br>' : '2.镜像仓库模块: <span class="text-danger">失败</span>; <br>';
      $ret['init']['repos'] = $retRepos;
      //通知服务发现模块初始化
      $retHubble = $this->initHubble($bizId, $retReg[$regId]['biz'], 0);
      $reInit .= ($retHubble['code'] === 0) ? '3.服务发现模块: 成功; <br>' : '3.服务发现模块: <span class="text-danger">失败</span>; <br>';
      $ret['init']['hubble'] = $retHubble;
      //通知服务编排模块初始化
      $retLayout = $this->initLayout($bizId, $retReg[$regId]['biz'], 0);
      $reInit .= ($retLayout['code'] === 0) ? '4.服务编排模块: 成功; <br>' : '4.服务编排模块: <span class="text-danger">失败</span>; <br>';
      $ret['init']['layout'] = $retLayout;

      $ret['code'] = 0;
      $ret['msg'] = '审批操作成功<br>' . $reInit;
    }
    return $ret;
  }

  function initCloud($id, $name, $status){
    global $myUser, $arrRecodeLog;
    $ret = [ 'code' => 1, 'msg' => 'Param Error' ];
    $header = array(
      'accept: application/json',
      'Content-Type: application/json',
      'X-HTTP-Method-Override: POST',
      'Authorization: '.$myUser,
      'X-CORRELATION-ID: ' . str_replace(array('0.',' '),'',microtime()),
      'X-Biz-ID: ' . $id,
      'X-Biz-Name: ' . $name,
      'X-Biz-Status: ' . $status,
    );
    $url = CLOUD_DOMAIN . '/v1/init';

    $handle = curl_init();
    curl_setopt($handle, CURLOPT_URL, $url);
    curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
    curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($handle, CURLOPT_TIMEOUT, 20);
    curl_setopt($handle, CURLOPT_CUSTOMREQUEST, 'POST');
    curl_setopt($handle, CURLOPT_POST, 1);
    $result = curl_exec($handle);
    $arrRecodeLog['t_code'] .= '多云对接接口：' . $url . "\n";
    $arrRecodeLog['t_code'] .= '多云对接返回：' . str_replace(array("\n", "\r"), '', $result) . "\n\n";
    $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
    if($http_code < 200 || $http_code >= 300){
      if($http_code == 0) $result = 'timeout';
      if($aRe=json_decode($result,true)){
        if(isset($aRe['msg'])){
          $ret['msg'] = $aRe['msg'];
        }else{
          $ret['msg'] = $result;
          $ret['http_code'] = $http_code;
          $ret['url'] = addslashes($url);
        }
      }else{
        $ret['msg'] = preg_replace('/\s+/',' ',$result);
        $ret['http_code'] = $http_code;
        $ret['url'] = addslashes($url);
      }
    }else{
      if($result = json_decode($result, true)){
        $ret['code'] = (isset($result['code'])) ? $result['code'] : 1;
        $ret['msg'] = (isset($result['msg'])) ? $result['msg'] : json_encode($result, JSON_UNESCAPED_UNICODE);
      }else{
        $ret['msg'] = $result;
      }
    }
    return $ret;
  }

  function initRepos($id, $name, $status){
    global $arrRecodeLog;
    $ret = [ 'code' => 1, 'msg' => 'Param Error' ];
    $header = array(
      'accept: application/json',
      'Content-Type: application/json',
      'X-HTTP-Method-Override: POST',
      'X-CORRELATION-ID: ' . str_replace(array('0.',' '),'',microtime()),
      'X-Biz-ID: ' . $id,
      'X-Biz-Name: ' . $name,
      'X-Biz-Status: ' . $status,
    );
    $url = REPOS_DOMAIN . '/api/projects';
    $handle = curl_init();
    curl_setopt($handle, CURLOPT_URL, $url);
    curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
    curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($handle, CURLOPT_TIMEOUT, 10);
    curl_setopt($handle, CURLOPT_USERPWD, REOPS_AUTH);
    curl_setopt($handle, CURLOPT_HTTPAUTH, CURLAUTH_BASIC);
    curl_setopt($handle, CURLOPT_HEADER, 1);
    curl_setopt($handle, CURLOPT_CUSTOMREQUEST, 'POST');
    curl_setopt($handle, CURLOPT_POST, 1);
    //传入的镜像名称改变为业务方ID
    $the_projec_name = (string)$id;
    while(!isset($the_projec_name{3})){
        $the_projec_name = '0'.$the_projec_name;
    }
    $data = [ 'public' => 1, 'project_name' => $the_projec_name ];
    curl_setopt($handle, CURLOPT_POSTFIELDS, json_encode($data, JSON_UNESCAPED_UNICODE));
    $result = curl_exec($handle);
    $arrRecodeLog['t_code'] .= '镜像仓库接口：' . $url . "\n";
    $arrRecodeLog['t_code'] .= '镜像仓库传入：' . json_encode($data, JSON_UNESCAPED_UNICODE) . "\n";
    $arrRecodeLog['t_code'] .= '镜像仓库返回HTTPCODE：' . str_replace(array("\n", "\r"), '', $result) . "\n\n";
    $header_size = curl_getinfo($handle, CURLINFO_HEADER_SIZE);
    $header = explode("\r\n",substr($result,0,$header_size));
    $arrHeader = array();
    foreach($header as $v){
      if(trim($v)==='') continue;
      if(strpos($v,'HTTP/')!==false){
        $arrHeader['http_code']=$v;
        continue;
      }
      $kv=strpos($v,':');
      $arrHeader[substr($v,0,$kv)]=trim(substr($v,$kv+1));
    }
    $result = substr($result, $header_size);
    $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
    $arrRecodeLog['t_code'] .= '镜像仓库返回：' . $http_code . "\n\n";
    if($http_code < 200 || $http_code >= 300){
      if($http_code == 0) $result = 'timeout';
      if($aRe=json_decode($result,true)){
        if(isset($aRe['msg'])){
          $ret['msg'] = $aRe['msg'];
        }else{
          $ret['msg'] = $result;
          $ret['http_code'] = $http_code;
          $ret['url'] = addslashes($url);
        }
      }else{
        $ret['msg'] = preg_replace('/\s+/',' ',$result);
        $ret['http_code'] = $http_code;
        $ret['url'] = addslashes($url);
      }
    }else{
      $ret['code'] = 0;
      $ret['msg'] = 'http_code: ' . $http_code;
    }
    $ret['header'] = $arrHeader;
    return $ret;
  }

  function initLayout($id, $name, $status){
    global $myUser, $arrRecodeLog;
    $ret = [ 'code' => 1, 'msg' => 'Param Error' ];
    $header = array(
      'accept: application/json',
      'Content-Type: application/json',
      'X-HTTP-Method-Override: GET',
      'Authorization: '.$myUser,
      'X-CORRELATION-ID: ' . str_replace(array('0.',' '),'',microtime()),
      'X-Biz-ID: ' . $id,
      'X-Biz-Name: ' . $name,
      'X-Biz-Status: ' . $status,
    );
    $url = LAYOUT_DOMAIN . '/init/db';
    $handle = curl_init();
    curl_setopt($handle, CURLOPT_URL, $url);
    curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
    curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($handle, CURLOPT_TIMEOUT, 10);
    curl_setopt($handle, CURLOPT_CUSTOMREQUEST, 'GET');
    $result = curl_exec($handle);
    $arrRecodeLog['t_code'] .= '服务编排接口：' . $url . "\n";
    $arrRecodeLog['t_code'] .= '服务编排返回：' . str_replace(array("\n", "\r"), '', $result) . "\n\n";
    $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
    if($http_code < 200 || $http_code >= 300){
      if($http_code == 0) $result = 'timeout';
      if($aRe=json_decode($result,true)){
        if(isset($aRe['msg'])){
          $ret['msg'] = $aRe['msg'];
        }else{
          $ret['msg'] = $result;
          $ret['http_code'] = $http_code;
          $ret['url'] = addslashes($url);
        }
      }else{
        $ret['msg'] = preg_replace('/\s+/',' ',$result);
        $ret['http_code'] = $http_code;
        $ret['url'] = addslashes($url);
      }
    }else{
      if($result = json_decode($result, true)){
        $ret['code'] = (isset($result['code'])) ? $result['code'] : 1;
        $ret['msg'] = (isset($result['msg'])) ? $result['msg'] : json_encode($result, JSON_UNESCAPED_UNICODE);
      }else{
        $ret['msg'] = $result;
      }
    }
    return $ret;
  }

  function initHubble($id, $name, $status){
    global $myUser, $arrRecodeLog;
    $ret = [ 'code' => 1, 'msg' => 'Param Error' ];
    $header = array(
      'accept: application/json',
      'Content-Type: application/json',
      'X-HTTP-Method-Override: POST',
      'appkey: ' . HUBBLE_APPKEY,
      'Authorization: '.$myUser,
      'X-CORRELATION-ID: ' . str_replace(array('0.',' '),'',microtime()),
      'X-Biz-ID: ' . $id,
      'X-Biz-Name: ' . $name,
      'X-Biz-Status: ' . $status,
    );
    $url = HUBBLE_DOMAIN . '/v1/tools/tool/nginx_init/';
    $handle = curl_init();
    curl_setopt($handle, CURLOPT_URL, $url);
    curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
    curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($handle, CURLOPT_TIMEOUT, 10);
    curl_setopt($handle, CURLOPT_CUSTOMREQUEST, 'POST');
    curl_setopt($handle, CURLOPT_POST, 1);
    $result = curl_exec($handle);
    $arrRecodeLog['t_code'] .= '服务发现接口：' . $url . "\n";
    $arrRecodeLog['t_code'] .= '服务发现返回：' . str_replace(array("\n", "\r"), '', $result) . "\n\n";
    $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
    if($http_code < 200 || $http_code >= 300){
      if($http_code == 0) $result = 'timeout';
      if($aRe=json_decode($result,true)){
        if(isset($aRe['msg'])){
          $ret['msg'] = $aRe['msg'];
        }else{
          $ret['msg'] = $result;
          $ret['http_code'] = $http_code;
          $ret['url'] = addslashes($url);
        }
      }else{
        $ret['msg'] = preg_replace('/\s+/',' ',$result);
        $ret['http_code'] = $http_code;
        $ret['url'] = addslashes($url);
      }
    }else{
      if($result = json_decode($result, true)){
        $ret['code'] = (isset($result['code'])) ? $result['code'] : 1;
        $ret['msg'] = (isset($result['msg'])) ? $result['msg'] : json_encode($result, JSON_UNESCAPED_UNICODE);
      }else{
        $ret['msg'] = $result;
      }
    }
    return $ret;
  }
}
$mySelf=new myself();

/*权限检查*/
$pageForSuper = true;//当前页面是否需要管理员权限
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

//记录操作日志
$logFlag = true;
$logDesc = '';
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => $myUser,
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
    case 'list':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getList();
      if(count($retArr['content'])>$myPageSize) $myPageSize=count($retArr['content']);
      $retArr['page'] = $myPage;
      $retArr['pageSize'] = $myPageSize;
      if(!isset($retArr['pageCount'])||$retArr['pageCount']<1) $retArr['pageCount']=1;
      if(!isset($retArr['count'])) $retArr['count']=count($retArr['content']);
      if($retArr['page'] > $retArr['pageCount']) $retArr['page'] = 1;
      break;
    case 'info':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getInfo($fIdx);
      break;
    case 'insert':
      if(!$pageForSuper && $myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '添加';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) unset($arrJson['id']);
        $retArr = $mySelf->update('add', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'update':
      if(!$pageForSuper && $myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '修改';
      if(isset($arrJson) && !empty($arrJson)){
        $retArr = $mySelf->update('update', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'delete':
      if(!$pageForSuper && $myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '删除';
      if(isset($arrJson) && !empty($arrJson)){
        $retArr=$mySelf->update('delete', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'audit':
      if(!$pageForSuper && $myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '审核申请';
      if(isset($arrJson) && !empty($arrJson)){
        include_once('../../include/biz.php');
        include_once('../../include/user.php');
        $retArr = $mySelf->audit($arrJson);
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
