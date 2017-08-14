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

    public function doAddNode($id, $args, $user){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        // 获取变更的配置信息
        $ret = $this->alterationTypeTbl->field('type,content')
            ->where("id = '$id'")
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
                    $content['port'], $content['weight']);

                if($ret['code'] != 0) return $ret;

                // -------- 对 consul 的处理
                if($ret['content']['is_consul']){
                    $return['content'] = [
                        'type' => 'sync',
                        'task_id' => 0,
                    ];
                    return $return;
                }
                
                
                //数据转存;进入定时任务
                $sid=$content['script_id'];
                $name=$content['name'];
                $gid=$content['group_id'];
                $correlation=I('server.HTTP_X_CORRELATION_ID');
                //$url='http://101.201.76.175:5555/v1/nginx/timing/beginTimingReload?id='.$id.'&sid='.$sid.'&name='.$name.'&user='.$user.'&gid='.$gid.'&correlation='.$correlation;
                $url  = 'http://'.C('HUBBLE_HOST').':'.C('HUBBLE_PORT').'/v1/nginx/timing/beginTimingReload?id='.$id.'&sid='.$sid.'&name='.$name.'&user='.$user.'&gid='.$gid.'&correlation='.$correlation;
                $host = parse_url($url,PHP_URL_HOST);
                $port = parse_url($url,PHP_URL_PORT);
                $port = $port ? $port : 80;
                $scheme = parse_url($url,PHP_URL_SCHEME);
                $path = parse_url($url,PHP_URL_PATH);
                $query = parse_url($url,PHP_URL_QUERY);
                if($query) $path .= '?'.$query;
                if($scheme == 'https') {
                    $host = 'ssl://'.$host;
                }
                $fp = fsockopen($host,$port,$error_code,$error_msg,1);
                stream_set_blocking($fp,true);//开启了手册上说的非阻塞模式
                stream_set_timeout($fp,1);//设置超时
                $header = "GET $path HTTP/1.1\r\n";
                $header.="Host: $host\r\n";
                $header.="Connection: close\r\n\r\n";//长连接关闭
                fwrite($fp, $header);
                usleep(1000); // 如果没有这延时，可能在nginx服务器上就无法执行成功
                fclose($fp);
                //返回下发任务完成，后续操作交给定时任务
                $return['content'] = [
                    'type' => 'async',
                    'task_id' => '0'
                ];
                return $return;
                

                $task = $upstream->callTunnel(
                    $content['script_id'], $content['name'], $user, true, $content['group_id']);

                if($task['code'] != 0) return $task;
                $task = $task['content'];

                // 记录变更表
                $history = new AlterationHistory();
                $ret = $history->addRecord('async', $task['ansible_id'], $task['ansible_name'], 'ansible', $user);
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
                $ret = $history->addRecord('sync', 0, 'add slb', 'slb', $user);
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

    public function doDelNode($id, $args, $user){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->alterationTypeTbl->field('type,content')
            ->where("id = '$id'")
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

                $ret = $upstream->delNode($content['name'], $content['group_id'], $content['ips']);
                if($ret['code'] != 0) return $ret;

                // -------- 对 consul 的处理
                if($ret['content']['is_consul']){
                    $return['content'] = [
                        'type' => 'sync',
                        'task_id' => 0,
                    ];
                    return $return;
                }
                
                //数据转存;进入定时任务
                $sid=$content['script_id'];
                $name=$content['name'];
                $gid=$content['group_id'];
                $correlation=I('server.HTTP_X_CORRELATION_ID');
                //$url='http://101.201.76.175:5555/v1/nginx/timing/beginTimingReload?id='.$id.'&sid='.$sid.'&name='.$name.'&user='.$user.'&gid='.$gid.'&correlation='.$correlation;
                $url  = 'http://'.C('HUBBLE_HOST').':'.C('HUBBLE_PORT').'/v1/nginx/timing/beginTimingReload?id='.$id.'&sid='.$sid.'&name='.$name.'&user='.$user.'&gid='.$gid.'&correlation='.$correlation;
                $host = parse_url($url,PHP_URL_HOST);
                $port = parse_url($url,PHP_URL_PORT);
                $port = $port ? $port : 80;
                $scheme = parse_url($url,PHP_URL_SCHEME);
                $path = parse_url($url,PHP_URL_PATH);
                $query = parse_url($url,PHP_URL_QUERY);
                if($query) $path .= '?'.$query;
                if($scheme == 'https') {
                    $host = 'ssl://'.$host;
                }
                $fp = fsockopen($host,$port,$error_code,$error_msg,1);
                stream_set_blocking($fp,true);//开启了手册上说的非阻塞模式
                stream_set_timeout($fp,1);//设置超时
                $header = "GET $path HTTP/1.1\r\n";
                $header.="Host: $host\r\n";
                $header.="Connection: close\r\n\r\n";//长连接关闭
                fwrite($fp, $header);
                usleep(1000); // 如果没有这延时，可能在nginx服务器上就无法执行成功
                fclose($fp);
                //返回下发任务完成，后续操作交给定时任务
                $return['content'] = [
                    'type' => 'async',
                    'task_id' => '0'
                ];
                return $return;
                
                $task = $upstream->callTunnel(
                    $content['script_id'], $content['name'], $user, true, $content['group_id']);

                if($task['code'] != 0) return $task;
                $task = $task['content'];

                // 记录变更表
                $history = new AlterationHistory();
                $ret = $history->addRecord('async', $task['ansible_id'], $task['ansible_name'], 'ansible', $user);
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
                $ret = $history->addRecord('sync', 0, 'delete slb', 'slb', $user);
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
