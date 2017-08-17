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
        $data=M('AlterationHistory')->where(['global_id' => $gid])->find();
        $gid=$data['task_name'];
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
