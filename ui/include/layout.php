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
