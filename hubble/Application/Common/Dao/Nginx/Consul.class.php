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
