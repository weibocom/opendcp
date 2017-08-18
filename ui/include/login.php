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


class login{
  var $useren;
  var $usercn;
  var $userpass;
  var $userid;
  var $usertype;
  var $usermail;
  var $userstatus;
  
  var $authtable='member';//验证用数据表
  
  function ldapAuth($useren,$userpass){
    global $myLdapHost,$myLdapPort,$myLdapUser,$myLdapPass,$myLdapBind,$myLdapSearch;
    if(empty($useren)||empty($userpass)) return false;
    $ds = ldap_connect($myLdapHost, $myLdapPort) ;
    if ($ds) {
      $r = ldap_bind($ds, "CN={$myLdapUser},{$myLdapBind}", $myLdapPass);
      $sr = ldap_search($ds, $myLdapSearch, "(sAMAccountName={$useren})") ;
      $user_arr = ldap_get_entries($ds, $sr);
      if($user_arr['count'] == "1"){
        $user_binddn = $user_arr[0]["dn"];
        $ub = ldap_bind($ds, $user_binddn, $userpass);
        if($ub){
          $uInfo=array(
            "samaccountname"=>$user_arr[0]['samaccountname'][0],
            "cn"=>mb_convert_encoding($user_arr[0]['cn'][0],"UTF-8","GBK"),
            "mail"=>$user_arr[0]['mail'][0],
          );
          return $uInfo;
        }
      }
    }
    return false;
  }
  
  function setSession(){
    @session_start();
    $_SESSION['open_uid'] = $this->userid;
    $_SESSION['open_user'] = $this->useren;
    $_SESSION['open_cnuser'] = $this->usercn;
    $_SESSION['open_usertype'] = $this->usertype;
    $_SESSION['open_email'] = $this->usermail;
    $_SESSION['open_status'] = (int)$this->userstatus;
  }
  
  function userLogout(){
    @session_start();
    unset($_SESSION['open_user']);
    session_unset();
    session_destroy();
  }
  
  function userAuth($arrJson){
    global $db;
    $useren=(isset($arrJson['user'])&&!empty($arrJson['user']))?$arrJson['user']:'';
    $userpass=(isset($arrJson['pass'])&&!empty($arrJson['pass']))?$arrJson['pass']:'';
    $usertype=(isset($arrJson['type'])&&!empty($arrJson['type']))?$arrJson['type']:'';
    switch($usertype){
      case 'ldap':
        $ldapArr=$this->ldapAuth($useren,$userpass);
        if($ldapArr){
          $sql='SELECT * FROM '.$this->authtable." WHERE en='{$useren}';";
          if($query=$db->query($sql)){
            if(!$arr=$query->fetch_array(MYSQL_ASSOC)){
              $ldapuser=$ldapArr['samaccountname'];
              $ldapcn=$ldapArr['cn'];
              $ldapmail=$ldapArr['mail'];
              $sql="INSERT INTO ".$this->authtable." (en,cn,type,mail,status) VALUES('{$ldapuser}','{$ldapcn}','ldap','{$ldapmail}',1);";
              $query=$db->query($sql);
              $sql='SELECT * FROM '.$this->authtable." WHERE en='{$useren}';";
              if($query=$db->query($sql)) $arr=$query->fetch_array(MYSQL_ASSOC);
            }
          }
          unset($arr['pw']);
          $this->userid=$arr['id'];
          $this->useren=$arr['en'];
          $this->usercn=$arr['cn'];
          $this->usertype=$arr['type'];
          $this->usermail=$arr['mail'];
          $this->userstatus=$arr['status'];
          $this->setSession();
          if(!empty($arr)) return $arr;
        }
        break;
      case 'local':
        $password=md5($userpass);
        $sql="SELECT * FROM ".$this->authtable." WHERE en='{$useren}' AND type='local' AND pw='{$password}';";
        if($query=$db->query($sql)){
          if($arr=$query->fetch_array(MYSQL_ASSOC)){
            $this->userid=$arr['id'];
            $this->useren=$arr['en'];
            $this->usercn=$arr['cn'];
            $this->usertype=$arr['type'];
            $this->usermail=$arr['mail'];
            $this->userstatus=$arr['status'];
            $this->setSession();
            unset($arr['pw']);
            if(!empty($arr)) return $arr;
          }
        }
        break;
      default:
        return false;
    }
    return false;
  }
}
$login=new login();
?>
