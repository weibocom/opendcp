<?php
class cloud{

  private $domain;
  private $reqid;

  public function __construct() {
    $this->domain = CLOUD_DOMAIN;
    $this->reqid = str_replace(array('0.',' '),'',microtime());
  }

  function curl($token='', $module = '', $method = '', $data = '', $id = '') {
    if($module && $method ){
      $header = array(
        'accept: application/json',
        'Content-Type: application/json',
        'X-HTTP-Method-Override: ' . $method,
        'Authorization: '.$token,
        'X-CORRELATION-ID: ' . $this->reqid,
      );
      $url = $this -> domain . '/v1/' . $module;
      if($id) $url.='/' . $id;
      if($method == 'GET' || $method == 'DELETE') $url.=(is_array($data) && !empty($data)) ? '?'.http_build_query($data) : $data;

      $handle = curl_init();
      curl_setopt($handle, CURLOPT_URL, $url);
      curl_setopt($handle, CURLOPT_HTTPHEADER, $header);
      curl_setopt($handle, CURLOPT_RETURNTRANSFER, 1);
      curl_setopt($handle, CURLOPT_TIMEOUT, 20);
      curl_setopt($handle, CURLOPT_CUSTOMREQUEST, $method);
      if($method == 'POST'){
        curl_setopt($handle, CURLOPT_POST, 1);
      }
      if(is_array($data)) $data=json_encode($data);
      curl_setopt($handle, CURLOPT_POSTFIELDS, $data);
      $result = curl_exec($handle);
      $http_code = curl_getinfo($handle, CURLINFO_HTTP_CODE);
      if($t = json_decode($result,true)){
        if(!isset($t['content'])) $t['content']=array();
        $result = json_encode($t, JSON_UNESCAPED_UNICODE);
      }
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

  function get($token='', $module = '', $method = '', $data = '', $id = '') {
    if($ret = $this -> curl($token, $module, $method, $data, $id)) return $ret;
    return false;
  }

}

$cloud=new cloud();

?>
