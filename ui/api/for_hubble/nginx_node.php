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
include_once('../../include/hubble.php');
$thisClass = $hubble;

class myself{
  private $module = 'nginx';
  private $subModule = 'node';
  private $arrFormat = array();

  function getList($param = array()){
    global $thisClass;
    $ret=array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($strList = $thisClass->get($this->module, $this->subModule, 'list', $param)){
      $arrList = json_decode($strList,true);
      if(isset($arrList['code']) && $arrList['code'] == 0 && isset($arrList['data']['content'])){
        $ret = array(
          'code' => 0,
          'msg' => 'success',
          'title' => array(
            '<input type="checkbox" id="selectAll" onclick="checkAll(this)"/>',
            '#',
            '单元',
            'IP',
            '创建时间',
            '用户',
            '#',
            ),
          'content' => array(),
        );
        if(isset($arrList['data']['count'])) $ret['count'] = $arrList['data']['count'];
        if(isset($arrList['data']['total_page'])) $ret['pageCount'] = $arrList['data']['total_page'];
        if(isset($arrList['data']['page'])) $ret['page'] = $arrList['data']['page'];
        $i=0;
        foreach($arrList['data']['content'] as $k => $v){
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
      }
    }
    $ret['ret'] = $strList;
    return $ret;
  }

  function update($action = '', $param = array()){
    global $thisClass;
    $ret = array('code' => 1, 'msg' => 'Illegal Request', 'ret' => '');
    if($action){
      if($strList = $thisClass->get($this->module, $this->subModule, $action, $param)){
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
        }
      }
      $ret['ret'] = $strList;
    }
    return $ret;
  }

  function format($arr = array(),$field = 'data'){
    if(!is_array($arr) || empty($arr) || empty($this->arrFormat)) return $arr;
    $ret = array();
    foreach($arr as $k => $v){
      if(in_array($k, $this->arrFormat)){
        $ret[$k] = $v;
      }else{
        $ret[$field][$k] = $v;
      }
    }
    return $ret;
  }

  function checkParam($name='',$value=''){
    $ret='';
    switch($name){
      case 'ips':
        $value=preg_split("/[\s,;]+/",$value);
        if(is_array($value)){
          foreach($value as $v){
            if(filter_var($v, FILTER_VALIDATE_IP, FILTER_FLAG_IPV4)) $ret.=($ret)?','.$v:$v;
          }
        }else{
          $ret=(filter_var($value, FILTER_VALIDATE_IP, FILTER_FLAG_IPV4))?$value:false;
        }
        break;
      default:
        $ret=$value;
        break;
    }
    return (empty($ret)) ? false : $ret;
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

$fUnit=(isset($_POST['fUnit'])&&!empty($_POST['fUnit']))?trim($_POST['fUnit']):((isset($_GET['fUnit'])&&!empty($_GET['fUnit']))?trim($_GET['fUnit']):'');
$fIp=(isset($_POST['fIp'])&&!empty($_POST['fIp']))?trim($_POST['fIp']):((isset($_GET['fIp'])&&!empty($_GET['fIp']))?trim($_GET['fIp']):'');
$fIdx=(isset($_POST['fIdx'])&&!empty($_POST['fIdx']))?trim($_POST['fIdx']):((isset($_GET['fIdx'])&&!empty($_GET['fIdx']))?trim($_GET['fIdx']):'');

$myJson=(isset($_POST['data'])&&!empty($_POST['data']))?trim($_POST['data']):((isset($_GET['data'])&&!empty($_GET['data']))?trim($_GET['data']):'');
$arrJson=($myJson)?json_decode($myJson,true):array();

$sid=(isset($_POST['sid'])&&!empty($_POST['sid']))?trim($_POST['sid']):((isset($_GET['sid'])&&!empty($_GET['sid']))?trim($_GET['sid']):'');
$nodeips=(isset($_POST['nodeips'])&&!empty($_POST['nodeips']))?trim($_POST['nodeips']):((isset($_GET['nodeips'])&&!empty($_GET['nodeips']))?trim($_GET['nodeips']):'');


//记录操作日志
$logFlag = true;
$logDesc = '';
$arrRecodeLog=array(
  't_time' => date('Y-m-d H:i:s'),
  't_user' => $myUser,
  't_module' => '节点管理',
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
        'limit' => $myPageSize,
        'unit_id' => $fUnit,
        'ip' => $fIdx,
      );
      if(!$fUnit) unset($arrJson['unit_id']);
      if(!$fIdx) unset($arrJson['ip']);
      $retArr = $mySelf->getList($arrJson);
      $retArr['page'] = $myPage;
      $retArr['pageSize'] = $myPageSize;
      if($retArr['page'] > $retArr['pageCount']) $retArr['page'] = 1;
    break;
    case 'insert':
      //添加节点
      if(isset($nodeips) && !empty($nodeips)){
        $arrJson['ips']=$nodeips;
        $con = new mysqli(DB_HOST, DB_USER, DB_PW,'hubble');
        if (!mysqli_connect_errno()) {
          $sql='SELECT content FROM tbl_hubble_alteration_type WHERE id='.$sid.';';
          if ($result=$con->query($sql)){
            if($row=$result->fetch_row()){
              $data=json_decode($row[0], true);
              $sql='SELECT id FROM tbl_hubble_nginx_unit WHERE group_id='.$data['group_id'];
              if ($result=$con->query($sql)){
                if($row=$result->fetch_row()){
                  $arrJson['unit_id']=$row[0];
                }
              }
            }
          }
        }
        $con->close();
      }
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '添加';
      if(isset($arrJson) && !empty($arrJson)){
        if($arrJson['ips'] = $mySelf->checkParam('ips', $arrJson['ips'])){
          $arrJson['user'] = $myUser;
          $retArr = $mySelf->update('add', $arrJson);
          $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
        }else{
          $retArr['msg']='ip format error';
        }
      }
    break;
    case 'delete':
      //删除节点
      if(isset($nodeips) && !empty($nodeips)){
        $con = new mysqli(DB_HOST, DB_USER, DB_PW,'hubble');
        $arrJson['nodes']='';
        if (!mysqli_connect_errno()) {
          $arr=explode(",",$mySelf->checkParam('ips', $nodeips));
          for($i=0;$i<count($arr);$i++){
            $sql='SELECT id,unit_id FROM tbl_hubble_nginx_node WHERE ip="'.$arr[$i].'";';
            if ($result=$con->query($sql)){
              if($row=$result->fetch_row()){
                $arrJson['nodes'].=$row[0].',';
                $arrJson['unit_id']=$row[1];
              }
            }
          }
          $arrJson['nodes'] = substr($arrJson['nodes'],0,strlen($str)-1);
        }
        $con->close();
      }
      if($myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '删除';
      if(isset($arrJson) && !empty($arrJson)){
        $arrJson['user'] = $myUser;
        $retArr=$mySelf->update('delete', $arrJson);
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
