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
