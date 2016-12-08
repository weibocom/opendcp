<?php
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