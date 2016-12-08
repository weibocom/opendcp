<?php
/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 2016/11/4
 * Time: 下午2:20
 */
namespace Common\Dao\Slb;


class Slb {

    private $address;

    function __construct(){

        $this->address =  C('HUBBLE_SLB_HTTP');
    }

    public function addNode($id, $ips, $weight){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $url = $this->address . '/v1/slb/backendservers/by_ip/';

        $paramIp = [];
        foreach($ips as $ip){
            $paramIp[] = [ "Address" =>$ip, "Weight"=> (int)$weight];
        }

        $param['LoadBalancerId'] = $id;
        $param['BackendServerList'] = $paramIp;

        $success = false;
        for($i = 1; $i <= 3; $i++){
            $ret = http_for_slb($url, json_encode($param), 'POST');
            hubble_log(HUBBLE_INFO, "request to slb server[$url]-[".json_encode($param)."]");
            hubble_log(HUBBLE_INFO, "request to slb server response: ".json_encode($ret));
            if($ret['code'] != 0){
                hubble_log(HUBBLE_WARN, "request slb failed...");
                continue;
            }
            $tmp = json_decode($ret['data'], true);
            if(empty($tmp)){
                hubble_log(HUBBLE_WARN, "request slb failed...content not vaild json");
                $ret['error'] = 'request slb failed...content not vaild json';
                continue;
            }
            if($tmp['code'] != 0){
                hubble_log(HUBBLE_WARN, "request slb failed... error:{$tmp['msg']}");
                $ret['error'] = $tmp['msg'];
                continue;
            }

            $success = true;
            break;
        }

        if(!$success){
            $return['code'] = 1;
            $return ['msg'] = $ret['error'];
        }
        return $return;
    }

    public function _delNode($id, $ips){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if(count($ips) > 20){
            $return['code'] = 1;
            $return['msg'] = 'count of ip must less than or equal to 20.';
            return $return;
        }

        $url = $this->address . '/v1/slb/backendservers/by_ip/';
        $param = [ "BackendServers" =>$ips, "LoadBalancerId"=> $id];

        $success = false;
        for($i = 1; $i <= 3; $i++){
            $ret = http_for_slb($url, json_encode($param), 'DELETE');
            hubble_log(HUBBLE_INFO, "request to slb server[$url]-[".json_encode($param)."]");
            hubble_log(HUBBLE_INFO, "request to slb server response: ".json_encode($ret));
            if($ret['code'] != 0){
                hubble_log(HUBBLE_WARN, "request slb failed...");
                continue;
            }
            $tmp = json_decode($ret['data'], true);
            if(empty($tmp)){
                hubble_log(HUBBLE_WARN, "request slb failed...content not vaild json");
                $ret['error'] = 'request slb failed...content not vaild json';
                continue;
            }
            if($tmp['code'] != 0){
                hubble_log(HUBBLE_WARN, "request slb failed... error:{$tmp['msg']}");
                $ret['error'] = $tmp['msg'];
                continue;
            }

            $success = true;
            break;
        }

        if(!$success){
            $return['code'] = 1;
            $return ['msg'] = $ret['error'];
        }
        return $return;
    }

    public function delNode($id, $ips){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $success = [];
        for($i = count($ips), $j = 0; $i > 0; $i -= 20,$j += 20){
            $tmp = array_slice($ips, $j, 20);
            $ret = $this->_delNode($id, $tmp);
            if($ret['code'] != 0){
                $return['code'] = 1;
                $return['msg'] = $ret['msg'];
                $return['content']['success'] = $success;
                $return['content']['fail'] = array_diff($success, $ips);
                return $return;
            }
            $success[] = $tmp;
        }

        $return['content']['success'] = $success;
        $return['content']['fail'] = [];
        return $return;
    }

}