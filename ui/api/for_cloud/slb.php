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
include_once('../../include/cloud.php');
$thisClass = $cloud;

class myself{
  private $module = 'slb';
  private $sub_module = 'loadbalancer';

  function getList($myUser = '', $type='list', $param = array(), $id=''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module.'/'.$type, 'GET', $param ,$id)){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['content'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'title' => array(
            '#',
            '名称',
            '服务地址',
            '服务端口',
            '后端端口',
            'Backend',
            'Listener',
            '实例状态',
            '类型',
            '创建时间',
            '#',
          ),
          'content' => array(),
        );
        $ret['count'] = (isset($arrList['totalCount'])) ? $arrList['totalCount'] : 0;
        $ret['pageCount'] = ($ret['count']>0&&isset($arrList['pageSize'])&&$arrList['pageSize']>0) ? ceil($ret['count']/$arrList['pageSize']) : 1;
        $ret['page'] = (isset($arrList['pageNumber'])) ? $arrList['pageNumber'] : 1;
        $i=0;
        foreach($arrList['content'] as $k => $v){
          $i++;
          $tArr = array();
          $tArr['i'] = $i;
          foreach($v as $key => $value){
            $tArr[$key] = $value;
          }
          $ret['content'][] = $tArr;
        }
      }else{
        $ret['code'] = 1;
        $arrList = json_decode($strList,true);
        $ret['msg'] = (isset($arrList['msg']))?$arrList['msg']:$strList;
        $ret['remote'] = $strList;
      }
    }
    $ret['ret'] = $strList;
    return $ret;
  }

  function getInfo($myUser = '', $idx = '', $fKeyId = ''){
    global $thisClass;

    if(!empty($fKeyId)) {
      $data = array(
          'keyId' => $fKeyId,
      );
    } else {
      $data = '';
    }

    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module.'/'.$this->sub_module, 'GET', $data, $idx)){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['content'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'content' => array(),
        );
        $ret['content']=$arrList['content'];
      }else{
        $ret['code'] = 1;
        $arrList = json_decode($strList,true);
        $ret['msg'] = (isset($arrList['msg']))?$arrList['msg']:$strList;
        $ret['remote'] = $strList;
      }
    }
    $ret['ret'] = $strList;
    return $ret;
  }

  function update($myUser = '', $method = 'POST', $param = array(), $id = ''){
    global $thisClass;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($param){
      if($strList = $thisClass->get($myUser, $this->module.'/'.$this->sub_module, $method, $param , $id)){
        $arrList = json_decode($strList,true);
        if(isset($arrList['code']) && $arrList['code'] == 0){
          $ret = array(
            'code' => 0,
            'msg' => 'success',
          );
        }else{
          $ret['code'] = 1;
          $arrList = json_decode($strList,true);
          $ret['msg'] = (isset($arrList['msg']))?$arrList['msg']:$strList;
          $ret['remote'] = $strList;
        }
      }
      $ret['ret'] = $strList;
    }
    return $ret;
  }

  function active($myUser = '', $method = 'POST', $param = array(), $id = ''){
    global $thisClass;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($param){
      if($strList = $thisClass->get($myUser, $this->module.'/'.$this->sub_module.'/status', $method, $param , $id)){
        $arrList = json_decode($strList,true);
        if(isset($arrList['code']) && $arrList['code'] == 0){
          $ret = array(
            'code' => 0,
            'msg' => 'success',
          );
        }else{
          $ret['code'] = 1;
          $arrList = json_decode($strList,true);
          $ret['msg'] = (isset($arrList['msg']))?$arrList['msg']:$strList;
          $ret['remote'] = $strList;
        }
      }
      $ret['ret'] = $strList;
    }
    return $ret;
  }

  function format($arr = array()){
    $ret = array();
    if(is_array($arr)&&!empty($arr)){
      foreach($arr as $k=>$v){
        switch($k){
          case 'RegionName':  $ret['Zone'][$k]=$v; break;
          case 'ZoneName':    $ret['Zone'][$k]=$v; break;
          case 'VpcId':       $ret['Network'][$k]=$v; break;
          case 'SubnetId':    $ret['Network'][$k]=$v; break;
          case 'OrganizationId':  $ret[$k]=(int)$v; break;
          case 'DataDiskSize':    $ret[$k]=(int)$v; break;
          case 'DataDiskNum':     $ret[$k]=(int)$v; break;
          default :           $ret[$k]=$v; break;
        }
      }
    }
    if(!empty($ret)&&(!isset($ret['Name']))) $ret['Name']='Test_'.date('U');
    return $ret;
  }

}
$mySelf=new myself();

/*权限检查*/
$pageForSuper = false;//当前页面是否需要管理员权限
$hasLimit = ($pageForSuper)?isSuper($myUser):true;
$myAction = (isset($_POST['action'])&&!empty($_POST['action']))?trim($_POST['action']):((isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'');
$myPage = (isset($_POST['page'])&&intval($_POST['page'])>0)?intval($_POST['page']):((isset($_GET['page'])&&intval($_GET['page'])>0)?intval($_GET['page']):1);
$myPageSize = (isset($_POST['pagesize'])&&intval($_POST['pagesize'])>0)?intval($_POST['pagesize']):((isset($_GET['pagesize'])&&intval($_GET['pagesize'])>0)?intval($_GET['pagesize']):$myPageSize);

$fRegion=(isset($_POST['fRegion'])&&!empty($_POST['fRegion']))?trim($_POST['fRegion']):((isset($_GET['fRegion'])&&!empty($_GET['fRegion']))?trim($_GET['fRegion']):'');
$fIdx=(isset($_POST['fIdx'])&&!empty($_POST['fIdx']))?trim($_POST['fIdx']):((isset($_GET['fIdx'])&&!empty($_GET['fIdx']))?trim($_GET['fIdx']):'');

$myJson=(isset($_POST['data'])&&!empty($_POST['data']))?trim($_POST['data']):((isset($_GET['data'])&&!empty($_GET['data']))?trim($_GET['data']):'');
$arrJson=($myJson)?json_decode($myJson,true):array();

$fKeyId=(isset($_POST['fKeyId'])&&!empty($_POST['fKeyId']))?trim($_POST['fKeyId']):((isset($_GET['fKeyId'])&&!empty($_GET['fKeyId']))?trim($_GET['fKeyId']):'');
$arrJson['keyId'] = $fKeyId;

//记录操作日志
$logFlag = true;
$logDesc = 'FAILED';
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => $myUser,
  't_module' => '阿里云SLB管理',
  't_action' => '',
  't_desc' => 'Resource:' . $_SERVER['REMOTE_ADDR'] . '.',
  't_code' => '传入：' . $myJson . "\n\n",
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
      $arrJson = array(
        'page' => $myPage,
        'pageSize' => $myPageSize,
        'keyId' => $fKeyId,
      );
      $retArr = $mySelf->getList($myUser, 'list', $arrJson , $fRegion);
      $retArr['page'] = $myPage;
      $retArr['pageSize'] = $myPageSize;
      if(!isset($retArr['pageCount']) || $retArr['page'] > $retArr['pageCount']) $retArr['page'] = 1;
      break;
    case 'info':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getInfo($myUser,$fIdx,$fKeyId);
      break;
    case 'insert':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '创建';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) unset($arrJson['id']);
        $arrJson=$mySelf->format($arrJson);
        $retArr = $mySelf->update($myUser,'POST', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'delete':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '删除';
      if(isset($arrJson) && !empty($arrJson)){
        $retArr=$mySelf->update($myUser, 'DELETE', $arrJson, $arrJson['id']);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'active':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '启用';
      if(!empty($arrJson)){
        $arrJson['LoadBalancerStatus']='active';
        $retArr = $mySelf->active($myUser,'POST', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'inactive':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '停用';
      if(!empty($arrJson)){
        $arrJson['LoadBalancerStatus']='inactive';
        $retArr = $mySelf->active($myUser,'POST', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
  }
}else{
  $retArr['msg'] = 'Permission Denied!';
}
//记录日志
if($logFlag){
  $arrRecodeLog['t_desc'] = $logDesc.', '.$arrRecodeLog['t_desc'];
  $arrRecodeLog['t_code'] .= '外部接口传入：' . json_encode($arrJson,JSON_UNESCAPED_UNICODE) . "\n\n";
  $arrRecodeLog['t_code'] .= '外部接口返回：' . str_replace(array("\n", "\r"), '', $retArr['ret']) . "\n\n";
  $arrRecodeLog['t_code'] .= '返回：' . json_encode($retArr,JSON_UNESCAPED_UNICODE);
  if(empty($arrRecodeLog['t_action'])) $arrRecodeLog['t_action'] = $myAction;
  logRecord($arrRecodeLog);
}
//返回结果
if(isset($retArr['action']) && !empty($retArr['action'])) $retArr['action'] = $myAction;
if(isset($retArr['ret'])) unset($retArr['ret']);
echo json_encode($retArr, JSON_UNESCAPED_UNICODE);
?>
