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
 * Date: 16/9/5
 * Time: 下午1:23
 */

namespace Common\Dao\Adaptor;

use Common\Dao\Nginx\Consul;
use Common\Dao\Nginx\NodeModel;
use Common\Dao\Nginx\Shell;
use Common\Dao\Nginx\Upstream;
use Common\Dao\Slb\Slb;

class Adaptor {

    private $alterationTypeTbl;

    function __construct(){
        $this->alterationTypeTbl = M('AlterationType');
    }

    public function doAddNode($id, $args, $user, $bid){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        // 获取变更的配置信息
        $ret = $this->alterationTypeTbl->field('type,content')
            ->where(['id' => $id, 'biz_id' => $bid])
            ->find();

        if($ret === false){
            $ret['code'] = 1;
            $ret['msg'] = 'read alteration type failed, ID:'.$id.' ERROR:'. $this->alterationTypeTbl->getDbError();
            return $return;
        }
        if(is_null($ret)){
            $ret['code'] = 1;
            $ret['msg'] = 'read alteration type failed, no such type. ID:'.$id;
            return $return;
        }
        $type = $ret['type'];

        $content = json_decode($ret['content'], true);
        if(is_null($content)){
            $ret['code'] = 1;
            $ret['msg'] = 'alteration content decode failed: '. $ret['content'];
            return $return;
        }

        $content = array_merge($content, $args);

        switch($type){

            case 'NGINX':
                // 0... 准备变更文件
                // 1... 获取脚本的内容
                // 2... 合成ansible json
                // 3... 启动一个任务
                // 4... 记录ID进入变更表
                // 5... 返回记录ID,变更类型

                // 准备变更文件
                $upstream = new Upstream();

                $ret = $upstream->checkArgs($content);
                if($ret['code'] != 0) return $ret;

                $ret = $upstream->addNode(
                    $content['name'], $content['group_id'], $content['ips'],
                    $content['port'], $content['weight'], $bid);

                if($ret['code'] != 0) return $ret;

                // -------- 对 consul 的处理
                if($ret['content']['is_consul']){
                    $return['content'] = [
                        'type' => 'sync',
                        'task_id' => 0,
                    ];
                    return $return;
                }

                $task = $upstream->callTunnel(
                    $content['script_id'], $content['name'], $user, $bid, true, $content['group_id']);

                if($task['code'] != 0) return $task;
                $task = $task['content'];

                // 记录变更表
                $history = new AlterationHistory();
                $ret = $history->addRecord('async', $task['ansible_id'], $task['ansible_name'], 'ansible', $user, $bid);
                if($ret['code'] != 0) return $ret;

                $return['content'] = [
                    'type' => 'async',
                    'task_id' => $ret['content'],
                ];
                return $return;

                break;
            case 'SLB':

                // 记录变更表
                $history = new AlterationHistory();
                $ret = $history->addRecord('sync', 0, 'add slb', 'slb', $user, $bid);
                if($ret['code'] != 0) return $ret;

                hubble_log(HUBBLE_INFO, "adding slb: ID:{$content['slb_id']}, ips:". implode(',', $content['ips']));

                $slb = new Slb();
                $ret = $slb->addNode($content['slb_id'], $content['ips'], $content['weight']);
                if($ret['code'] != 0) return $ret;

                $return['content'] = ['type' => 'sync'];
                return $return;
                break;
            case 'ELB':
        }

        return $return;
    }

    public function doDelNode($id, $args, $user, $bid){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->alterationTypeTbl->field('type,content')
            ->where(['id' => $id, 'biz_id' => $bid])
            ->find();

        if($ret === false){
            $ret['code'] = 1;
            $ret['msg'] = 'read alteration type failed, ID:'.$id.' ERROR:'. $this->alterationTypeTbl->getDbError();
            return $return;
        }
        if(is_null($ret)){
            $ret['code'] = 1;
            $ret['msg'] = 'read alteration type failed, no such type. ID:'.$id;
            return $return;
        }
        $type = $ret['type'];

        $content = json_decode($ret['content'], true);
        if(is_null($content)){
            $ret['code'] = 1;
            $ret['msg'] = 'alteration content decode failed: '. $ret['content'];
            return $return;
        }

        $content = array_merge($content, $args);

        switch($type){
            case 'NGINX':
                $upstream = new Upstream();

                $ret = $upstream->checkArgs($content);
                if($ret['code'] != 0) return $ret;

                $ret = $upstream->delNode($content['name'], $content['group_id'], $content['ips'], $bid);
                if($ret['code'] != 0) return $ret;

                // -------- 对 consul 的处理
                if($ret['content']['is_consul']){
                    $return['content'] = [
                        'type' => 'sync',
                        'task_id' => 0,
                    ];
                    return $return;
                }
                
                $task = $upstream->callTunnel(
                    $content['script_id'], $content['name'], $user, $bid, true, $content['group_id']);

                if($task['code'] != 0) return $task;
                $task = $task['content'];

                // 记录变更表
                $history = new AlterationHistory();
                $ret = $history->addRecord('async', $task['ansible_id'], $task['ansible_name'], 'ansible', $user, $bid);
                if($ret['code'] != 0) return $ret;

                $return['content'] = [
                    'type' => 'async',
                    'task_id' => $ret['content'],
                ];
                return $return;

                break;
            case 'SLB':
                // 记录变更表
                $history = new AlterationHistory();
                $ret = $history->addRecord('sync', 0, 'delete slb', 'slb', $user, $bid);
                if($ret['code'] != 0) return $ret;

                hubble_log(HUBBLE_INFO, "delete slb: ID:{$content['slb_id']}, ips:". implode(',', $content['ips']));

                $slb = new Slb();
                $ret = $slb->delNode($content['slb_id'], $content['ips']);
                if($ret['code'] != 0) return $ret;

                $return['content'] = ['type' => 'sync'];
                return $return;
                break;
            case 'ELB':
        }

        return $return;
    }
}
