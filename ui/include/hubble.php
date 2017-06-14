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


class hubble{

  private $domain;
  private $appkey;
  private $reqid;

  public function __construct() {
    $this->domain = HUBBLE_DOMAIN;
    $this->appkey = HUBBLE_APPKEY;
    $this->reqid = str_replace(array('0.',' '),'',microtime());
  }

  //Action
  private $arrAction = array(
    'list'      =>  'GET',
    'detail'    =>  'GET',
    'content'   =>  'GET',
    'result'    =>  'GET',
    'history'   =>  'GET',
    'state'     =>  'GET',
    'modify'    =>  'PUT',
    'delete'    =>  'DELETE',
    'delarg'    =>  'DELETE',
    'deletew'   =>  'DELETE',
    'delete_interface'    =>  'DELETE',
    'delete_privilege'    =>  'DELETE',
    'list_interface'      =>  'GET',
    'list_privilege'      =>  'GET',
    'list_ver'            =>  'GET',
    'statistics'          =>  'GET',
    'white_list'          =>  'GET',
    'illegal_list'        =>  'GET',
    'unit_list'           =>  'GET',
    'upstream_list'       =>  'GET',
    'type_param'          =>  'GET',
    'check_state'         =>  'GET',
    'iplog'               =>  'GET',
  );

  function curl($module = '', $subModule='', $action = '', $data = '') {
    if($module && $action ){
      $method = (isset($this -> arrAction[$action])) ? $this -> arrAction[$action] : 'POST';
      $header = array(
        'X-HTTP-Method-Override: ' . $method,
        'appkey: ' . $this->appkey,
        'X-CORRELATION-ID: ' . $this->reqid,
        'X-Biz-ID: ' . $_SESSION['open_biz_id'],
        'X-Biz-Name: ' . $_SESSION['open_biz_name'],
        'X-Biz-Status: ' . $_SESSION['open_biz_status'],
      );
      if($module=='extension'){
        if($subModule){
          $url = $this -> domain . '/' . $module . '/v1/' . $subModule . '/' . $action;
        }else{
          $url = $this -> domain . '/' . $module . '/v1/' . $action;
        }
      }else{
        if($subModule){
          $url = $this -> domain . '/v1/' . $module . '/' . $subModule . '/' . $action;
        }else{
          $url = $this -> domain . '/v1/' . $module . '/' . $action;
        }
      }
      $jsonData=($action=='check_state')?$data:array();
      if(is_array($data)){
        foreach($data as $k=>$v){
          if(is_array($v)) $data[$k]=json_encode($v);
        }
        $data = http_build_query($data);
      }
      if($method == 'GET' || $method == 'DELETE') $url.='?'.$data;
      $handle = curl_init();
      curl_setopt($handle, CURLOPT_URL, $url);
      curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
      curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
      curl_setopt($handle, CURLOPT_TIMEOUT, 10);
      curl_setopt($handle, CURLOPT_CUSTOMREQUEST, $method);
      if($action=='check_state'){
        $data=json_encode($jsonData);
      }
      if($method == 'POST'){
        curl_setopt($handle, CURLOPT_POST, 1);
      }
      if(is_array($data)) $data=json_encode($data);
      curl_setopt($handle, CURLOPT_POSTFIELDS, $data);
      $result = curl_exec($handle);
      if($t = json_decode($result,true)) $result = json_encode($t, JSON_UNESCAPED_UNICODE);
      $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
      if($http_code < 200 || $http_code >= 300){
        if($http_code == 0) $result = 'timeout';
        if(json_decode($result)) return '{"code":1,"http_code":' . $http_code . ',"url":"' . addslashes($url) . '","msg":' . $result . '}';
        return '{"code":1,"http_code":' . $http_code . ',"url":"' . addslashes($url) . '","msg":"' . preg_replace('/\s+/',' ',$result) . '"}';
      }else{
        return $result;
      }
    }
    return false;
  }

  function get($module = '', $subModule='', $action = '', $data = '') {
    if($ret = $this -> curl($module, $subModule, $action, $data)) return $ret;
    return false;
  }

}

$hubble=new hubble();

?>
