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
include_once('../../include/layout.php');
$thisClass = $layout;

class myself{
  private $module = 'pool';

  function getList($myUser = '', $param = array()){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, 'service/'.$param['service_id'], 'list_pools', $param)){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'title' => array(
            '#',
            '服务池名称',
            '隶属服务',
            '描述',
            '节点数',
            '机型模板',
            '服务发现',
            '操作',
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

    function getAllList($myUser = '', $param = array()){
        global $thisClass;
        $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
        if($strList = $thisClass->get($myUser, $this->module, 'list', $param)){
            $arrList = json_decode($strList,true);
            if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data'])){
                $ret = array(
                    'code' => 0,
                    'msg' => 'success',
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

  function getInfo($myUser = '', $idx = ''){
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

  function getState($myUser = '', $idx = ''){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($myUser, $this->module.'/search_by_ip', $idx)){
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

  function update($myUser = '', $action = '', $param = array(), $id = ''){
    global $thisClass;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($action){
      if($strList = $thisClass->get($myUser, $this->module, $action, $param , $id)){
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

$fService=(isset($_POST['fService'])&&!empty($_POST['fService']))?trim($_POST['fService']):((isset($_GET['fService'])&&!empty($_GET['fService']))?trim($_GET['fService']):'');
$fIdx=(isset($_POST['fIdx'])&&!empty($_POST['fIdx']))?trim($_POST['fIdx']):((isset($_GET['fIdx'])&&!empty($_GET['fIdx']))?trim($_GET['fIdx']):'');

$myJson=(isset($_POST['data'])&&!empty($_POST['data']))?trim($_POST['data']):((isset($_GET['data'])&&!empty($_GET['data']))?trim($_GET['data']):'');
$arrJson=($myJson)?json_decode($myJson,true):array();

//记录操作日志
$logFlag = true;
$logDesc = 'FAILED';
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => $myUser,
  't_module' => '服务池管理',
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
        'service_id' => $fService,
        'name' => $fIdx,
      );
      $retArr = $mySelf->getList($myUser, $arrJson);
      $retArr['page'] = $myPage;
      $retArr['pageSize'] = $myPageSize;
      $retArr['pageCount'] = ceil($retArr['count']/$retArr['pageSize']);
      if($retArr['page'] > $retArr['pageCount']) $retArr['page'] = 1;
      break;
    case 'info':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getInfo($myUser,$fIdx);
      break;
    case 'state':
      $logFlag = false;//本操作不记录日志
      $retArr = $mySelf->getState($myUser,$fIdx);
      break;
    case 'insert':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '创建';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) unset($arrJson['id']);
        $tExpand=(isset($arrJson['tpl_expand']))?(int)$arrJson['tpl_expand']:0;
        $tShrink=(isset($arrJson['tpl_shrink']))?(int)$arrJson['tpl_shrink']:0;
        $tDeploy=(isset($arrJson['tpl_deploy']))?(int)$arrJson['tpl_deploy']:0;
        $arrJson['tasks']=array(
          'expand'=>$tExpand,
          'shrink'=>$tShrink,
          'deploy'=>$tDeploy,
        );
        if(isset($arrJson['tpl_expand'])) unset($arrJson['tpl_expand']);
        if(isset($arrJson['tpl_shrink'])) unset($arrJson['tpl_shrink']);
        if(isset($arrJson['tpl_deploy'])) unset($arrJson['tpl_deploy']);
        if(isset($arrJson['service_id'])) $arrJson['service_id']=(int)$arrJson['service_id'];
        if(isset($arrJson['vm_type'])) $arrJson['vm_type']=(int)$arrJson['vm_type'];
        if(isset($arrJson['sd_id'])) $arrJson['sd_id']=(int)$arrJson['sd_id'];
        $retArr = $mySelf->update($myUser,'create', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'update':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '修改';
      if(isset($arrJson) && !empty($arrJson)){
        $tExpand=(isset($arrJson['tpl_expand']))?(int)$arrJson['tpl_expand']:0;
        $tShrink=(isset($arrJson['tpl_shrink']))?(int)$arrJson['tpl_shrink']:0;
        $tDeploy=(isset($arrJson['tpl_deploy']))?(int)$arrJson['tpl_deploy']:0;
        $arrJson['tasks']=array(
          'expand'=>$tExpand,
          'shrink'=>$tShrink,
          'deploy'=>$tDeploy,
        );
        if(isset($arrJson['tpl_expand'])) unset($arrJson['tpl_expand']);
        if(isset($arrJson['tpl_shrink'])) unset($arrJson['tpl_shrink']);
        if(isset($arrJson['tpl_deploy'])) unset($arrJson['tpl_deploy']);
        if(isset($arrJson['service_id'])) $arrJson['service_id']=(int)$arrJson['service_id'];
        if(isset($arrJson['vm_type'])) $arrJson['vm_type']=(int)$arrJson['vm_type'];
        if(isset($arrJson['sd_id'])) $arrJson['sd_id']=(int)$arrJson['sd_id'];
        if(isset($arrJson['id'])) $arrJson['id']=(int)$arrJson['id'];
        $retArr=$mySelf->update($myUser, 'update', $arrJson, $arrJson['id']);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'delete':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '删除';
      if(isset($arrJson) && !empty($arrJson)){
        if(isset($arrJson['id'])) $arrJson['id']=(int)$arrJson['id'];
        $retArr=$mySelf->update($myUser, 'delete', $arrJson, $arrJson['id']);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'expand':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '扩容';
      if(isset($arrJson) && !empty($arrJson)){
        $fIdx=(isset($arrJson['pool']))?(int)$arrJson['pool']:0;
        if(isset($arrJson['pool'])) $arrJson['pool']=(int)$arrJson['pool'];
        if(isset($arrJson['num'])) $arrJson['num']=(int)$arrJson['num'];
        if(isset($arrJson['template_id'])) $arrJson['template_id']=(int)$arrJson['template_id'];
        $retArr = $mySelf->update($myUser,'expand', $arrJson, $fIdx);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'shrink':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '缩容';
      if(isset($arrJson) && !empty($arrJson)){
        $fIdx=(isset($arrJson['pool']))?(int)$arrJson['pool']:0;
        if(isset($arrJson['pool'])) $arrJson['pool']=(int)$arrJson['pool'];
        if(isset($arrJson['template_id'])) $arrJson['template_id']=(int)$arrJson['template_id'];
        if(isset($arrJson['check_all'])) unset($arrJson['check_all']);
        if(isset($arrJson['list'])) unset($arrJson['list']);
        $tArr=preg_split("/[\s,;]+/",$arrJson['ip']);
        $arrJson['nodes']=$tArr;
        if(isset($arrJson['ip'])) unset($arrJson['ip']);
        $retArr = $mySelf->update($myUser,'shrink', $arrJson, $fIdx);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'deploy':
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '上线';
      if(isset($arrJson) && !empty($arrJson)){
        $fIdx=(isset($arrJson['pool']))?(int)$arrJson['pool']:0;
        if(isset($arrJson['pool'])) $arrJson['pool']=(int)$arrJson['pool'];
        if(isset($arrJson['timeout'])) $arrJson['timeout']=(int)$arrJson['timeout'];
        if(isset($arrJson['auto'])) $arrJson['auto']=(int)$arrJson['auto'];
        if(isset($arrJson['max_num'])) $arrJson['max_num']=(int)$arrJson['max_num'];
        if(isset($arrJson['max_ratio'])) $arrJson['max_ratio']=(int)$arrJson['max_ratio'];
        $retArr = $mySelf->update($myUser,'deploy', $arrJson, $fIdx);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'poolList':
        $logFlag = false;//本操作不记录日志
        $arrJson = array(
            'page' => $myPage,
            'page_size' => $myPageSize,
        );
        $retArr = $mySelf->getAllList($myUser, $arrJson);
        $retArr['page'] = $myPage;
        $retArr['pageSize'] = $myPageSize;
        $retArr['pageCount'] = ceil($retArr['count']/$retArr['pageSize']);
        if($retArr['page'] > $retArr['pageCount']) $retArr['page'] = 1;
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
if(!isset($retArr['action']) || empty($retArr['action'])) $retArr['action'] = $myAction;
if(isset($retArr['ret'])) unset($retArr['ret']);
echo json_encode($retArr, JSON_UNESCAPED_UNICODE);
?>
