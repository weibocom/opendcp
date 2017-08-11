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
#include_once('../../include/func_session.php');
include_once('../../include/layout.php');
$thisClass = $layout;

class myself{
  private $module = 'task';

  function getList($myUser = '', $param = array()){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module, 'list', $param)){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'title' => array(
            '#',
            '名称',
            '状态',
            //'模板',
            '用户',
            '创建时间',
            '#',
            ),
          'content' => array(),
        );
        $ret['count'] = (isset($arrList['query_count'])) ? $arrList['query_count'] : count($arrList['data']);
        $ret['pageCount'] = (isset($arrList['page_size'])) ? ceil($ret['count'] / $arrList['page_size']) : 1;
        $ret['page'] = (isset($arrList['page'])) ? $arrList['page'] : 1;
        $i=0;
        foreach($arrList['data'] as $k => $v){
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

  function getInfo($myUser = '', $idx=''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module, $idx)){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'content' => array(),
        );
        $ret['content']=$arrList['data'];
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

  function getResult($myUser = '', $idx=''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module.'/node/'.$idx,'log')){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'content' => array(),
        );
        $ret['content']=$arrList['data'];
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

  function getTaskNode($myUser = '', $idx=''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module.'/'.$idx, 'detail')){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'content' => array(),
        );
        $ret['content']=$arrList['data'];
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

  function update($myUser = '', $action = '', $param = array(), $id=''){
    global $thisClass;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($action){
      if($strList = $thisClass->get($myUser, $this->module, $action, $param, $id)){
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

  function getTaskLog($myUser = '', $idx=''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module.'/flow/'.$idx,'log')){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data'])){
        $ret = array(
            'code' => 0,
            'msg' => 'success',
            'content' => array(),
        );
        $ret['content']=$arrList['data'];
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
}
$mySelf=new myself();

/*权限检查*/
$pageForSuper = false;//当前页面是否需要管理员权限
$hasLimit = ($pageForSuper)?isSuper($myUser):true;
$myAction = (isset($_POST['action'])&&!empty($_POST['action']))?trim($_POST['action']):((isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'');
$myPage = (isset($_POST['page'])&&intval($_POST['page'])>0)?intval($_POST['page']):((isset($_GET['page'])&&intval($_GET['page'])>0)?intval($_GET['page']):1);
$myPageSize = (isset($_POST['pagesize'])&&intval($_POST['pagesize'])>0)?intval($_POST['pagesize']):((isset($_GET['pagesize'])&&intval($_GET['pagesize'])>0)?intval($_GET['pagesize']):$myPageSize);

$fIdx=(isset($_POST['fIdx'])&&!empty($_POST['fIdx']))?trim($_POST['fIdx']):((isset($_GET['fIdx'])&&!empty($_GET['fIdx']))?trim($_GET['fIdx']):'');

$myJson=(isset($_POST['data'])&&!empty($_POST['data']))?trim($_POST['data']):((isset($_GET['data'])&&!empty($_GET['data']))?trim($_GET['data']):'');
$arrJson=($myJson)?json_decode($myJson,true):array();

$myUser = 'root';
//记录操作日志
$logFlag = true;
$logDesc = '';
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => $myUser,
  't_module' => '任务调度',
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
        'page_size' => $myPageSize,
        'name' => $fIdx,
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
    case 'result':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getResult($myUser,$fIdx);
      break;
    case 'nodes':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getTaskNode($myUser,$fIdx);
      break;
    case 'gettaskid':
      $arrRecodeLog['t_action'] = '获取taskid';
      $ip = $_REQUEST['ip'];
      include_once('../../include/node_init.php');
      $ret = node_init::getOneTaskByIp($ip);
      if(!empty($ret['id'])){
	echo $ret['id'];exit;
      } else {
	echo 0;exit;
      }
      break;
    case 'getdiskname':
        $arrRecodeLog['t_action'] = '获取diskname';
        $ip = $_REQUEST['ip'];
        include_once('../../include/node_init.php');
        $ret = node_init::getOneTaskByIp($ip);
        if (!empty($ret['disk_name'])) {
            echo $ret['disk_name'];
            exit;
        } else {
            echo 0;
            exit;
        }
        break;
    case 'addlog':
      $arrRecodeLog['t_action'] = '日志';
      include_once('../../include/node_init_log.php');
      include_once('../../include/node_init.php');
      $arr = array(
		'title'=>urldecode($_REQUEST['title']),
		'task_id'=>$_REQUEST['task_id'],
		'data'=>array('text'=>urldecode($_REQUEST['text'])),
		'status'=>$_REQUEST['status'],
      );
      $ret = node_init_log::insertLog($arr);
      if(!empty($_REQUEST['final'])){
		node_init::modifyOneNodeInit($_REQUEST['task_id'], array('status'=>$_REQUEST['final']));
      }
      $retArr['code'] = 0;
      $retArr['msg'] = '';
      $logDesc = 'SUCCESS';
      break;
    case 'addip':
      $arrRecodeLog['t_action'] = '创建';
      include_once('../../include/node_init.php');
      $arr = array(
		'ip'=>$_POST['ip'],
		'password'=>$_POST['password'],
		'type'=>$_POST['type'],
        'disk_name'=>$_POST['disk_name']
      );
      $ret = node_init::insertNodeInit($arr);
      $retArr['code'] = 0;
      $retArr['msg'] = '';
      $logDesc = 'SUCCESS';
      break;
    case 'getcontrollerip':
      include_once('../../include/keydata.php');
      $ip = keydata::getContentByKey('controller_ip');
      echo $ip;
      exit;

    case 'getcomputepowerbytime':
        include_once('../../include/computepower.php');
        $time = $_REQUEST['time'];
        $thetime = time() - $time*3600;
        $arr_ret = computepower::getListByTime($thetime);
        foreach($arr_ret as $k=>$v){
                $arr_ret[$k]['data'] = @json_decode($v['data'], true);
                $arr_ret[$k]['create_time'] = date('Y-m-d H:i:s', $v['create_time']);
        }
        echo json_encode(array('code'=>0, 'data'=>$arr_ret,));
        exit;

    case 'getcomputepower':
        include_once('../../include/openstack.php');
        openstack::needOpenstackLogin();
        $arr_hyper = openstack::getHypervisorList(array(), 0, 1000);
        $arr_ret = array(
                'vcpus'=>0,
                'vcpus_used'=>0,
                'memory_mb'=>0,
                'memory_mb_used'=>0,
                'memory_gb'=>0,
                'memory_gb_used'=>0,
                'local_gb'=>0,
                'local_gb_used'=>0,
                'machine_count'=>0,
        );
        if(!empty($arr_hyper['hypervisors'])){
                foreach($arr_hyper['hypervisors'] as $onehyper){ 
                        if($onehyper['state']=='up' && $onehyper['status']=='enabled'){
                                $arr_ret['machine_count'] += 1;
                                $arr_ret['vcpus'] += $onehyper['vcpus'];
                                $arr_ret['vcpus_used'] += $onehyper['vcpus_used'];
                                $arr_ret['memory_mb'] += $onehyper['memory_mb'];
                                $arr_ret['memory_mb_used'] += $onehyper['memory_mb_used'];
                                $arr_ret['memory_gb'] += sprintf("%.2f", $onehyper['memory_mb']/1024);
                                $arr_ret['memory_gb_used'] += sprintf("%.2f", $onehyper['memory_mb_used']/1024);
                                $arr_ret['local_gb'] += $onehyper['local_gb'];
                                $arr_ret['local_gb_used'] += $onehyper['local_gb_used'];                                                    
                        }
                }
        }

        echo json_encode(array('code'=>0, 'data'=>$arr_ret,));                                                                              
        exit;                                                                                                                               
    case 'start':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '启动';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) $arrJson['id']=(int)$arrJson['id'];
        $retArr = $mySelf->update($myUser, 'start', $arrJson, $arrJson['id']);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
    break;
    case 'pause':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '暂停';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) $arrJson['id']=(int)$arrJson['id'];
        $retArr = $mySelf->update($myUser, 'pause', $arrJson, $arrJson['id']);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'finish':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '暂停';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) $arrJson['id']=(int)$arrJson['id'];
        $retArr = $mySelf->update($myUser, 'stop', $arrJson, $arrJson['id']);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'stop_sub':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '停止子任务';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) $arrJson['id']=(int)$arrJson['id'];
        $retArr=$mySelf->update($myUser, 'stop_sub', $arrJson, $arrJson['id']);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'tasklog':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getTaskLog($myUser,$fIdx);
  }
}else{
  $retArr['msg'] = 'Permission Denied!';
}
//记录日志
if($logFlag){
  $arrRecodeLog['t_desc'] = $logDesc.', '.$arrRecodeLog['t_desc'];
  $arrRecodeLog['t_code'] .= '外部接口传入：' . json_encode($arrJson,JSON_UNESCAPED_UNICODE) . "\n\n";
  $arrRecodeLog['t_code'] .= '外部接口返回：' . str_replace(array("\n", "\r"), '', '') . "\n\n";
  $arrRecodeLog['t_code'] .= '返回：' . json_encode($retArr,JSON_UNESCAPED_UNICODE);
  if(empty($arrRecodeLog['t_action'])) $arrRecodeLog['t_action'] = $myAction;
  logRecord($arrRecodeLog);
}
//返回结果
if(isset($retArr['action']) && !empty($retArr['action'])) $retArr['action'] = $myAction;
if(isset($retArr['ret'])) unset($retArr['ret']);
echo json_encode($retArr, JSON_UNESCAPED_UNICODE);
?>
