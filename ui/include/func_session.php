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


$arrUri=explode('?',$_SERVER['REQUEST_URI']);
$pageHref=$arrUri[0];
@session_start();
$myUid    = @$_SESSION['open_uid'];
$myUser   = @$_SESSION['open_user'];
$myCn     = @$_SESSION['open_cnuser'];
$myType   = @$_SESSION['open_usertype'];
$myMail   = @$_SESSION['open_email'];
$myStatus = @$_SESSION['open_status'];
if(isset($pageAuth)&&$pageAuth===false){
  if(!$myUser) $myUser = 'AnyOne';
  if(!$myCn) $myCn = '访客';
  if($myStatus!==0) $myStatus = 1;
}else{
  //针对需要验证的页面，验证是否登录
  if(strpos($pageHref,'/api/')!==false){
    //API接口，返回JSON
    if(!isLogin()){
      $retArr=array(
        'code'=>1,
        'msg'=>'Illegal request, please login at first.',
      );
      echo json_encode($retArr,JSON_UNESCAPED_UNICODE);
      exit;
    }
  }else{
    //普通页面，跳转到登录页
    if(!isLogin()){
      header('location: /login.html');
      exit;
    }
  }
}
$mySuper=$_config['super'];
?>
