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
