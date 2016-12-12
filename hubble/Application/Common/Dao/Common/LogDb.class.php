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
 * Date: 2016/11/21
 * Time: 下午7:12
 */

namespace Common\Dao\Common;

class LogDb {

    private $log;

    function __construct(){

        $this->log = M('log');
    }

    public function insert($gid, $module, $log_info, $level){

        $data = [
            'global_id' => $gid,
            'module'    => $module,
            'log_info'  => $log_info,
            'level'     => $level,
            'url'       => I('server.REQUEST_URI'),
            'create_time' => date('Y-m-d H:i:s'),
        ];

        $this->log->add($data);
        return true;
    }

    public function getOctanLog($gid, $ip = ''){

        $url = C('HUBBLE_ANSIBLE_HTTP') .'/api/getlog';

        $header = [
            "X-CORRELATION-ID: $gid",
            "X-SOURCE: hubble",
            "Content-Type: application/json"
        ];

        $data= ['source' => 'hubble'];
        if(!empty($ip))
            $data['host'] = $ip;

        $ret = http($url, json_encode($data), 'POST', 5, $header);

        if($ret['code'] != 0)
            return ["get log from ansible failed: http request failed, code {$ret['code']}"];

        $response = json_decode($ret['data'], true);
        if($response === null)
            return ["get log from ansible failed: json decode fail", $ret['http']['url'], $ret['data']];

        if($response['code'] != 0)
            return ["get log from ansible failed: ".$response['message']];
        return empty($response['content']['log'])? ['ansible return empty log']:$response['content']['log'];
    }

    public function getAllLog($gid){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->log->field('log_info')
            ->where(['global_id' => $gid])
            ->where(['module' => ['NOT IN', 'GET,POST,PUT,DELETE']])
            ->select();

        if($ret === false) {
            $return['code'] = 1;
            $return['msg'] = '读取数据库失败';
        }
        if(empty($ret)) return $return;

        $return['content'] = array_map(function($i){return $i['log_info']; }, $ret);
        $return['content'][] = '----------------------reload server process log ------------------------';
        $return['content'] = array_merge($return['content'], $this->getOctanLog($gid));
        $return['content'] = implode("\n", $return['content']);

        return $return;
    }
}
