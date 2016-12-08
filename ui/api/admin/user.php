<?php
header('Content-type: application/json');
include_once('../../include/config.inc.php');
include_once('../../include/function.php');
include_once('../../include/func_session.php');
include_once('../../include/user.php');
$thisClass=$user;

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
          '类型',
          '邮箱',
          '状态',
          '#',
        ),
        'content' => array(),
      );
      $ret['count'] = (int)$thisClass->getCount('',$fIdx);
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
  't_module' => '用户管理',
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
        $arrJson['user'] = $myUser;
        $retArr=$mySelf->update('delete', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'on':
      if(!$pageForSuper && $myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '启用';
      if(isset($arrJson) && !empty($arrJson)){
        $arrJson['status'] = 0;
        $retArr = $mySelf->update('update', $arrJson);
        $logDesc = (isset($retArr['code']) && $retArr['code'] == 0) ? 'SUCCESS' : 'FAILED';
      }
      break;
    case 'off':
      if(!$pageForSuper && $myStatus > 0){ $retArr['msg'] = 'Permission Denied!'; break; }
      $arrRecodeLog['t_action'] = '停用';
      if(isset($arrJson) && !empty($arrJson)){
        $arrJson['status'] = 1;
        $retArr = $mySelf->update('update', $arrJson);
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