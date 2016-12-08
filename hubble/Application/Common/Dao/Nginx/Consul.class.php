<?php
/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 2016/10/24
 * Time: 下午5:19
 */

namespace Common\Dao\Nginx;

class Consul {

    private $url;

    function __construct(){
        $this->url = C('NGINX_CONSUL_ADDRESS');
    }

    public function addNode($name, $groupId, $ips, $port, $weight){
        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        return $return;
    }

    public function delNode($name, $groupId, $ips){
        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        return $return;
    }
}