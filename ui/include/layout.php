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


class layout{

  private $domain;
  private $reqid;

  public function __construct() {
    $this->domain = LAYOUT_DOMAIN;
    $this->reqid = str_replace(array('0.',' '),'',microtime());
  }

  //Action
  private $arrAction = array(
    'list'      =>  'GET',
    'create'    =>  'POST',
    'start'     =>  'POST',
    'stop_sub'  =>  'POST',
    'report'    =>  'POST',
    'update'    =>  'POST',
    'delete'    =>  'POST',
    'expand'    =>  'POST',
    'shrink'    =>  'POST',
    'deploy'    =>  'POST',
    'list_nodes'      =>  'GET',
    'add_nodes'       =>  'POST',
    'remove_nodes'    =>  'POST',
    'expandList'      =>  'POST',
    'uploadList'      =>  'POST',
    'saveTask'        =>  'POST',

  );

  function curl($token='', $module = '', $action = '', $data = '', $id = '') {
    if( $module && $action ){
      $method = (isset($this -> arrAction[$action])) ? $this -> arrAction[$action] : 'GET';
      $header = array(
        'accept: application/json',
        'Content-Type: application/json',
        'X-HTTP-Method-Override: ' . $method,
        'Authorization: '.$token,
        'X-CORRELATION-ID: ' . $this->reqid,
      );
      $url = $this -> domain . '/' . $module . '/' . $action;
      if($id) $url.='/' . $id;
      if($method == 'GET' || $method == 'DELETE') $url.=(is_array($data))?'?'.http_build_query($data):$data;
      $handle = curl_init();
      curl_setopt($handle, CURLOPT_URL, $url);
      curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
      curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
      curl_setopt($handle, CURLOPT_TIMEOUT, 10);
      curl_setopt($handle, CURLOPT_CUSTOMREQUEST, $method);
      if($method == 'POST'){
        curl_setopt($handle, CURLOPT_POST, 1);
        if(is_array($data)) $data=json_encode($data);
        curl_setopt($handle, CURLOPT_POSTFIELDS, $data);
      }
      $result = curl_exec($handle);
      if($t = json_decode($result)) $result = json_encode($t, JSON_UNESCAPED_UNICODE);
      $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
      if($http_code < 200 || $http_code >= 300){
        if($http_code == 0) $result = 'timeout';
        if($aRe=json_decode($result,true)){
          if(isset($aRe['msg'])){
            $result=$aRe['msg'];
          }else{
            return '{"code":1,"http_code":' . $http_code . ',"url":"' . addslashes($url) . '","msg":' . $result . '}';
          }
        }
        return '{"code":1,"http_code":' . $http_code . ',"url":"' . addslashes($url) . '","msg":"' . preg_replace('/\s+/',' ',$result) . '"}';
      }else{
        return $result;
      }
    }
    return false;
  }

  function get($token='', $module = '', $action = '', $data = '', $id = '') {
    if($ret = $this -> curl($token, $module, $action, $data, $id)) return $ret;
    return false;
  }

}

$layout=new layout();

?>
