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
 * Date: 16/9/6
 * Time: 下午2:13
 */

namespace Common\Dao\Adaptor;

class Channel {

    private $typeList = ['ANSIBLE', 'RSYNC'];

    public function isIllegal($type){
        return in_array($type, $this->typeList);

    }

    public function getTypeList(){
        return $this->typeList;
    }

    /*
     * @param $ips Array
     * @param $name 变更的名称
     */
    public function ansible($ips, $user, $tasks, $params, $fork_num,$cor_id=''){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if(!is_array($ips)|| empty($ips)){
            $return['code'] = 1;
            $return['msg'] = 'ansible: parameter ips is empty or not a array';
            return $return;
        }

        if(empty($tasks)){
            $return['code'] = 1;
            $return['msg'] = 'ansible: parameter tasks is empty';
            return $return;
        }
        $tasks_name = 'auto_reload_nginx'.date("Y-m-d-H:i:s").'_'.rand(10000,99999);
        $data = [
            'nodes'    => $ips,
            'user'     => $user,
            'name'     => $tasks_name,
            'tasks'    => $tasks,
            'params'   => $params,
            'fork_num' => $fork_num,
        ];

        $url = C('HUBBLE_ANSIBLE_HTTP').'/api/parallel_run';
        $data = json_encode($data);
        hubble_log(HUBBLE_DEBUG, "call ansible http: [$url]- [$data]");
        hubble_log(HUBBLE_INFO, "call ansible http: [$url]");

        $ret = http($url, $data, 'POST', 3,
            ['X-CORRELATION-ID:'.(I('server.HTTP_X_CORRELATION_ID')?I('server.HTTP_X_CORRELATION_ID'):$cor_id), 'X-SOURCE: hubble']);
        if($ret['code'] != 0){
            $return['code'] = $ret['errno'];
            $return['msg']  = 'ansible:'.$ret['error'];
            hubble_log(HUBBLE_WARN, 'ansible interface failed: '.$ret['error'].' exit!');
            return $return;
        }
        $content = $ret['data'];
        hubble_log(HUBBLE_INFO, "ansible http return : $content", 'ansible');

        $content = json_decode($content, true);
        if(empty($content)){
            $return['code'] = 1;
            $return['msg']  = 'ansible:'.'ansible http interface return null or not json format';
            return $return;
        }

        if($content['code'] != 0){
            $return['code'] = 1;
            $return['msg']  = 'ansible:'.'ansible http interface error: '. $content['message'];
            return $return;
        }
        $return['content'] = ['ansible_id' => $content['content']['id'], 'ansible_name' => $tasks_name];

        return $return;
    }

    public function ansibleCheck($taskName){

        $url = C('HUBBLE_ANSIBLE_HTTP') . '/api/check';

        $result = http($url,json_encode(['name' => $taskName]),
            'POST', 5,
            ['X-CORRELATION-ID:'.I('server.HTTP_X_CORRELATION_ID'),
             'X-SOURCE: hubble']);
        hubble_log(HUBBLE_INFO, 'call ansible interface check return: ' .$result['data']);

        return $result;
    }
}
