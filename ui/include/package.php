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


class package{

  private $domain;
  private $reqid;

  public function __construct() {
    $this->domain = PACKAGE_DOMAIN;
    $this->reqid = str_replace(array('0.',' '),'',microtime());
  }

  //Action
  private $arrAction = array(
    'list'      =>  'GET',
    'info'      =>  'GET',
    'new'       =>  'POST',
    'save'     =>  'POST',
    'clone'     =>  'POST',
    'delete'    =>  'POST',
    'build'     =>  'POST',
  );

  function curl($token='', $module = '', $action = '', $data = '', $id = '') {
    if($module && $action ){
      $method = (isset($this -> arrAction[$action])) ? $this -> arrAction[$action] : 'GET';
      $header = array(
        'accept: application/json',
        'X-HTTP-Method-Override: ' . $method,
        'Authorization: '.$token,
        'X-CORRELATION-ID: ' . $this->reqid,
        'X-Biz-ID: ' . $_SESSION['open_biz_id'],
        'X-Biz-Name: ' . $_SESSION['open_biz_name'],
        'X-Biz-Status: ' . $_SESSION['open_biz_status'],
      );
      $url = $this -> domain . '/api/' . $module . '/' . $action;
      if($id) $url.='/' . $id;
      if($method == 'GET' || $method == 'DELETE') $url.=(is_array($data))?'?'.http_build_query($data):'?'.$data;
      $handle = curl_init();
      curl_setopt($handle, CURLOPT_URL, $url);
      curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
      curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
      curl_setopt($handle, CURLOPT_TIMEOUT, 10);
      curl_setopt($handle, CURLOPT_CUSTOMREQUEST, $method);
      if($method == 'POST') curl_setopt($handle, CURLOPT_POST, 1);
      if(is_array($data)){
        curl_setopt($handle, CURLOPT_POSTFIELDS, http_build_query($data));
      }else{
        curl_setopt($handle, CURLOPT_POSTFIELDS, $data);
      }
      $result = curl_exec($handle);
      if($t = json_decode($result,true)) $result = json_encode($t, JSON_UNESCAPED_UNICODE);
      $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
      if($http_code < 200 || $http_code >= 300){
        if($http_code == 0) $result = 'timeout';
        if($aRe=json_decode($result,true)){
          if(is_array($aRe['errMsg'])){
            $result=$aRe['errMsg'];
          }else{
            return '{"code":1,"http_code":' . $http_code . ',"url":"' . addslashes($url) . '","msg":' . $result . '}';
          }
        }
        return '{"code":1,"http_code":' . $http_code . ',"url":"' . addslashes($url) . '","msg":"' . preg_replace('/\s+/',' ',$result) . '"}';
      }else{
        if($tArr=json_decode($result,true)){
          $code=(isset($tArr['code']))?(($tArr['code']==10000)?0:$tArr['code']):0;
          $msg=(isset($tArr['errMsg']))?$tArr['errMsg']:'success';
          $tArr['code']=$code;
          $tArr['msg']=$msg;
          unset($tArr['errMsg']);
          return json_encode($tArr);
        }else{
          return '{"code":0,"msg":"success","content":"'.$result.'"}';
        }
      }
    }
    return false;
  }

  function getHtml($token='', $module = '', $action = '', $data = '') {
    if($module && $action ){
      $method = 'GET';
      $header = array(
        'accept: application/json',
        'X-HTTP-Method-Override: ' . $method,
        'Authorization: '.$token,
      );
      $url = $this -> domain . '/' . $module . '/' . $action;
      $url.=(is_array($data))?'?'.http_build_query($data):$data;
      $handle = curl_init();
      curl_setopt($handle, CURLOPT_URL, $url);
      curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
      curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
      curl_setopt($handle, CURLOPT_TIMEOUT, 10);
      curl_setopt($handle, CURLOPT_CUSTOMREQUEST, $method);
      $result = curl_exec($handle);
      $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
      if($http_code < 200 || $http_code >= 300){
        if($http_code == 0) $result = 'timeout';
        return '{"code":1,"http_code":' . $http_code . ',"url":"' . addslashes($url) . '","msg":"' . preg_replace('/\s+/',' ',$result) . '"}';
      }else{
        $result=array(
          'code' => 0,
          'msg' => 'success',
          'content' => $result,
        );
        return json_encode($result);
      }
    }
    return false;
  }

  function get($token='', $module = '', $action = '', $data = '', $id = '') {
    if($ret = $this -> curl($token, $module, $action, $data, $id)) return $ret;
    return false;
  }

  function proxyCurl($token='', $method='GET', $uri = '', $data = '') {
    if(empty($method)) $method='GET';
    $ret=array(
      'code' => 1,
      'proxy' => array(
        'method' => $method,
        'uri' => $uri,
        'param' => $data,
      ),
      'msg' => 'success',
      'content' => '',
    );
    if($uri){
      $header = array(
        'accept: application/json',
        'X-HTTP-Method-Override: ' . $method,
        'Authorization: '.$token,
      );
      $url = $this -> domain . $uri;
      if($method == 'GET' || $method == 'DELETE') $url.='?'.$data;
      $handle = curl_init();
      curl_setopt($handle, CURLOPT_URL, $url);
      curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
      curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
      curl_setopt($handle, CURLOPT_TIMEOUT, 10);
      curl_setopt($handle, CURLOPT_CUSTOMREQUEST, $method);
      if($method == 'POST'){
        curl_setopt($handle, CURLOPT_POST, 1);
        curl_setopt($handle, CURLOPT_POSTFIELDS, $data);
      }
      $ret['content'] = curl_exec($handle);
      $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
      if($http_code < 200 || $http_code >= 300){
        $ret['msg'] = ($http_code == 0) ? 'timeout' : 'http_code: '.$http_code;
      }else{
        $ret['code'] = 0;
      }
      return $ret;
    }else{
      $ret['msg'] = 'uri null';
    }
    return false;
  }

}

$package=new package();

?>
