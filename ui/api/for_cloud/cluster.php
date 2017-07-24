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
  private $module = 'cluster';

  function getList($myUser = '', $param = array(), $id=''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module, 'GET', $param ,$id)){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['content'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'title' => array(
            '#',
            '名称',
            '使用云',
            '配额余量',
            '创建时间',
            '#',
          ),
          'content' => array(),
        );
        $i=0;
        foreach($arrList['content'] as $k => $v){

          //机型模板不再对外显示物理机
          if($v['Provider'] == 'phydev') continue;

          $i++;
          $tArr = array();
          $tArr['i'] = $i;
          foreach($v as $key => $value){
            $tArr[$key] = $value;
          }
          $ret['content'][] = $tArr;
        }
        $ret['count'] = (isset($arrList['totalCount'])) ? $arrList['totalCount'] : 0;
        if($i>$ret['count']) $ret['count']=count($ret['content']);
        $ret['pageCount'] = ($ret['count']>0&&isset($arrList['pageSize'])&&$arrList['pageSize']>0) ? ceil($ret['count']/$arrList['pageSize']) : 1;
        $ret['page'] = (isset($arrList['pageNumber'])) ? $arrList['pageNumber'] : 1;
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

  function getMachines($myUser = '', $namespace = '', $method = 'GET', $id = ''){
      global $thisClass;
      $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
      if($strList = $thisClass->get($myUser, $this->module.'/'.$namespace, $method, '' ,$id)){
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

  function getInfo($myUser = '', $idx = ''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module, 'GET', '', $idx)){
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
      if($strList = $thisClass->get($myUser, $this->module, $method, $param , $id)){
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

}
$mySelf=new myself();

/*权限检查*/
$pageForSuper = false;//当前页面是否需要管理员权限
$hasLimit = ($pageForSuper)?isSuper($myUser):true;
$myAction = (isset($_POST['action'])&&!empty($_POST['action']))?trim($_POST['action']):((isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'');
$myPage = (isset($_POST['page'])&&intval($_POST['page'])>0)?intval($_POST['page']):((isset($_GET['page'])&&intval($_GET['page'])>0)?intval($_GET['page']):1);
$myPageSize = (isset($_POST['pagesize'])&&intval($_POST['pagesize'])>0)?intval($_POST['pagesize']):((isset($_GET['pagesize'])&&intval($_GET['pagesize'])>0)?intval($_GET['pagesize']):$myPageSize);

$fOrg=(isset($_POST['fOrg'])&&!empty($_POST['fOrg']))?trim($_POST['fOrg']):((isset($_GET['fOrg'])&&!empty($_GET['fOrg']))?trim($_GET['fOrg']):'');
$fIdx=(isset($_POST['fIdx'])&&!empty($_POST['fIdx']))?trim($_POST['fIdx']):((isset($_GET['fIdx'])&&!empty($_GET['fIdx']))?trim($_GET['fIdx']):'');

$myJson=(isset($_POST['data'])&&!empty($_POST['data']))?trim($_POST['data']):((isset($_GET['data'])&&!empty($_GET['data']))?trim($_GET['data']):'');
$arrJson=($myJson)?json_decode($myJson,true):array();

//记录操作日志
$logFlag = true;
$logDesc = 'FAILED';
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => $myUser,
  't_module' => '类型管理',
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
        'pagesize' => $myPageSize,
      );
      $retArr = $mySelf->getList($myUser, $arrJson);
      $retArr['page'] = $myPage;
      $retArr['pageSize'] = $myPageSize;
      if(!isset($retArr['pageCount']) || $retArr['page'] > $retArr['pageCount']) $retArr['page'] = 1;
      break;
    case 'info':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getInfo($myUser,$fIdx);
      break;
    case 'insert':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '创建';
      if(isset($arrJson) && !empty($arrJson)){
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
    case 'machine':
        $logFlag = false;//本操作不记录日志
        if(isset($arrJson) && !empty($arrJson)){
            $retArr=$mySelf->getMachines($myUser, $arrJson["action"],'GET', $arrJson["hour"]);
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
