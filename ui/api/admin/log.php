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
include_once('../../include/user.php');
$thisClass=$user;

/*权限检查*/
$pageForSuper=true;//当前页面是否需要管理员权限
$hasLimit=($pageForSuper)?isSuper($myUser):false;
$myAction=(isset($_POST['action'])&&!empty($_POST['action']))?trim($_POST['action']):((isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'');
$myFilter=(isset($_POST['filter'])&&!empty($_POST['filter']))?trim($_POST['filter']):((isset($_GET['filter'])&&!empty($_GET['filter']))?trim($_GET['filter']):'');
$myIndex=(isset($_POST['index'])&&!empty($_POST['index']))?trim($_POST['index']):((isset($_GET['index'])&&!empty($_GET['index']))?trim($_GET['index']):'');
$myPage=(isset($_GET['page'])&&intval($_GET['page'])>0)?intval($_GET['page']):1;
$myPageSize=(isset($_GET['pagesize']))?intval($_GET['pagesize']):$myPageSize;

$myJson=(isset($_POST['data'])&&!empty($_POST['data']))?trim($_POST['data']):((isset($_GET['data'])&&!empty($_GET['data']))?trim($_GET['data']):'');
$arrJson=($myJson)?json_decode($myJson,true):array();

//记录操作日志
$logFlag=true;
$logDesc='';
$arrRecodeLog=array(
  't_time'=>date('Y-m-d H:i:s'),
  't_user'=>$myUser,
  't_module'=>'日志管理',
  't_action'=>'',
  't_desc'=>'Resource:'.$_SERVER['REMOTE_ADDR'].'.',
  't_code'=>'传入：'.$myJson."\n\n",
);
//返回
$retArr=array(
  'code'=>1,
  'action'=>$myAction,
);
if($hasLimit){
  $retArr['msg']='Param Error!';
  switch($myAction){
    case 'list':
      $logFlag=false;//本操作不记录日志
      $retArr=array(
        'code'=>0,
        'action'=>$myAction,
        'msg'=>'success',
        'page'=>$myPage,
        'pageSize'=>$myPageSize,
        'pageCount'=>0,
        'count'=>0,
        'filter'=>$myFilter,
        'title'=>array(),
        'content'=>array(),
      );
      $arrThead=array(
        '#',
        '时间',
        '用户',
        '模块',
        '行为',
        '明细',
        '#',
      );
      if($arrThead){
        foreach($arrThead as $t){
          $retArr['title'][]=$t;
        }
      }
      $retArr['count']=(int)$thisClass->getLogCount('',$myFilter);
      $retArr['pageCount']=($retArr['count']>0)?ceil($retArr['count']/$retArr['pageSize']):1;
      if($retArr['page']>$retArr['pageCount']) $retArr['page']=1;
      $arrList=$thisClass->getLog('',$myFilter,$retArr['page'],$retArr['pageSize']);
      if($arrList){
        $i=0;
        foreach($arrList as $k=>$v){
          $i++;
          $tArr=array();
          $tArr['i']=$i;
          foreach($v as $key=>$value){
            if(strpos($key,'t_')===0) $key=substr($key,2);
            $tArr[$key]=$value;
          }
          $retArr['content'][]=$tArr;
        }
      }
    break;
    case 'info':
      $logFlag=false;//本操作不记录日志
      if($myIndex){
        $retArr=array(
          'code'=>0,
          'action'=>$myAction,
          'msg'=>'success',
          'page'=>$myPage,
          'pageSize'=>$myPageSize,
          'pageCount'=>0,
          'count'=>0,
          'filter'=>$myFilter,
          'title'=>array(),
          'content'=>array(),
        );
        $arrThead=array(
          '#',
          '时间',
          '用户',
          '模块',
          '行为',
          '明细',
          'CODE',
        );
        if($arrThead){
          foreach($arrThead as $t){
            $retArr['title'][]=$t;
          }
        }
        $retArr['count']=0;
        $arrList=$thisClass->getLogById($myIndex);
        if($arrList){
          $i=0;
          foreach($arrList as $k=>$v){
            $i++;
            $tArr=array();
            $tArr['i']=$i;
            foreach($v as $key=>$value){
              if(strpos($key,'t_')===0) $key=substr($key,2);
              $tArr[$key]=$value;
            }
            $retArr['content'][]=$tArr;
            $retArr['count']++;
          }
        }
        $retArr['pageCount']=($retArr['count']>0)?ceil($retArr['count']/$retArr['pageSize']):1;
        if($retArr['page']>$retArr['pageCount']) $retArr['page']=1;
      }
    break;
    case 'dellog':
      $arrRecodeLog['t_action']='删除日志';
      if(isset($arrJson['id'])&&!empty($arrJson['id'])){
        if($ret=$thisClass->deleteLog($arrJson['id'])){
          $retArr['code']=0;
          $retArr['msg']='success';
          $logDesc.='SUCCESS';
        }else{
          $retArr['msg']='failed';
          $logDesc.='FAILED';
        }
        $arrRecodeLog['t_desc']=$logDesc.', '.$arrRecodeLog['t_desc'];
      }
      $arrRecodeLog['t_code'].='返回：'.json_encode($retArr,JSON_UNESCAPED_UNICODE);
    break;
  }
}else{
  $retArr['msg']='Permission Denied!';
}
//记录日志
if($logFlag){
  if(empty($arrRecodeLog['t_action'])) $arrRecodeLog['t_action']=$myAction;
  logRecord($arrRecodeLog);
}
//返回结果
echo json_encode($retArr,JSON_UNESCAPED_UNICODE);
?>
