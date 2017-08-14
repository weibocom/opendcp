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
 * Date: 16/9/1
 * Time: 下午3:04
 */

namespace Common\Dao\Nginx;


use Common\Dao\Adaptor\AlterationHistory;
use Common\Dao\Adaptor\Channel;

class Upstream {

    private $upstreamTbl;
    private $upstreamArg;

    function __construct(){

        $this->upstreamTbl = M('NginxConfUpstream');
        $this->upstreamArg = ' max_fails=0 fail_timeout=30s weight=';
    }

    public function countUpstream($where, $like = true){

        foreach($where as $k => $v){

            if($k == 'name' && $like){
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->upstreamTbl
            ->where($where)
            ->count();

        if($ret === NULL){
            return 0;
        } elseif($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());
            return false;
        } else{
            return (int)$ret;
        }
    }


    public function getUpstreamList($where, $page, $limit, $like = true){

        foreach($where as $k => $v){

            if($k == 'name' && $like){
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->upstreamTbl->field('content', true)
            ->where($where)
            ->page($page, $limit)
            ->select();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if($ret === NULL){
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->upstreamTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;
    }

    public function getUpstreamDetail($id){

        $ret = $this->upstreamTbl
            ->where(['id' => $id])
            ->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if($ret === NULL){
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->upstreamTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;
    }

    public function getUpstreamContent($groupId, $name){

        $ret = $this->upstreamTbl->where([
            'group_id' => $groupId,
            'name'     => $name,
        ])->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->upstreamTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());

        } elseif(empty($ret)){
            $return['code'] = 1;
            $return['msg'] = 'no such file';
        } else{
            $return['content'] = $ret['content'];
        }
        return $return;
    }

    public function addUpstream($name, $content, $groupId, $isConsul, $user){

        $data = [
            'name'        => $name,
            'content'     => $content,
            'is_consul'   => $isConsul,
            'group_id'    => $groupId,
            'deprecated'  => 0,
            'create_time' => date("Y-m-d H:i:s"),
            'update_time' => date("Y-m-d H:i:s"),
            'opr_user'    => $user,
        ];

        $ret = $this->upstreamTbl->add($data);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->upstreamTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;


    }

    public function setDeprecated($id){

        $ret = $this->upstreamTbl
            ->where(['id' => $id])
            ->setField('deprecated', 1);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->upstreamTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;
    }

    public function deleteUpstream($id){

        $ret = $this->upstreamTbl
            ->where(['id' => $id])
            ->delete();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->upstreamTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;

    }

    public function modifyUpstream($id, $content){

        $data = [
            'content' => $content,
            'update_time' => date("Y-m-d H:i:s"),
        ];

        $ret = $this->upstreamTbl
            ->where(['id' => $id])
            ->save($data);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->upstreamTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->upstreamTbl->getLastSql().' ERROR: '. $this->upstreamTbl->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;
    }

    // --------------------------------------------
    // -- 下面的函数使用与自动变更 -------------------

    /* 将一个正常的upstream分解为三部分
     *
     * @param upstream Array 正常的upstream
     * @return array
     *          head
     *          server
     *              key: ip value: value
     *          tail
     */
    private function _splitUpstream($upstream){

        $newUpstream = [];
        foreach($upstream as $line){
            if(preg_match('/^\s+server (\d+\.\d+\.\d+\.\d+).*/', $line) == 1){
                break;
            }
            if(preg_match('/^\s+keepalive.*/',$line)){
                array_shift($upstream);
                continue;
            }
            if(preg_match('/^\s+#.*/', $line)){  // 所有注释行 全部干掉
                array_shift($upstream);
                continue;
            }
            $newUpstream[] = $line;
            array_shift($upstream);
        }

        $server = [];
        foreach($upstream as $line){
            $ret = preg_match('/^\s+server (\d+\.\d+\.\d+\.\d+).*/', $line, $matches);
            if($ret == 1){
                $server[$matches[1]] = $line;
            }else {
                break;
            }
            array_shift($upstream);
        }
        // upstream 剩下的即为 尾部

        return [
            "head"      => $newUpstream,
            "server"    => $server,
            "tail"      => $upstream,
            "valid"     => count($server) != 0 ? true: false
        ];

    }

    /*
     * 删除某些节点
     * @param $upstream upstream的内容
     * @param $ips 要删除的ip数组
     */
    private function _fileDelNode($upstream, $ips){

        // 提取出upstream的头,放入 newUpstream中

        $new = $this->_splitUpstream($upstream);
        if(!$new['valid']) return null;

        foreach($ips as $item){
            unset($new['server'][$item]);
        }
        if(count($new['server']) == 0) return null;

        $newUpstream = $new['head'];
        $newUpstream[] = "\t\tkeepalive ".count($new['server']).";";

        $newUpstream = array_merge($newUpstream, array_values($new['server']));
        return array_merge($newUpstream, $new['tail']);
    }

    /*
     * 添加一些节点到upstream中
     * @param $upstream upstream内容
     * @param $ips Array 要添加的ip数组
     * @param $port ip的端口
     * @param $weight 流量权重
     *
     * @return String 新的upstream内容
     */
    private function _fileAddNode($upstream, $ips, $port, $weight = 20){

        $new = $this->_splitUpstream($upstream);
        if(!$new['valid']) return null;

        $oldIps = array_keys($new['server']);
        $allIps = array_unique(array_merge($oldIps, $ips));

        hubble_log(HUBBLE_INFO, "all ips is ". implode(',', $allIps));

        $newUpstream = $new['head'];
        $newUpstream[] = "\tkeepalive ".count($allIps).";";
        foreach($allIps as $ip){
            if(empty($ip)) continue; // 防止有空ip

            if(in_array($ip,$oldIps))
                $newUpstream[] = $new['server'][$ip];
            else{
                $newUpstream[] = "\tserver $ip:$port" .$this->upstreamArg. $weight.";";
                hubble_log(HUBBLE_INFO, "add new node $ip:$port");
            }
        }
        return array_merge($newUpstream, $new['tail']);
    }

    public function checkArgs($args){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $key = ['name', 'group_id', 'ips', 'port', 'weight'];
        foreach($key as $item){
            if(!isset($args[$item]) || empty($args[$item])){
                $return['code'] = 1;
                $return['msg'] = "parameter [$item] is empty";
                return $return;
            }
        }

        if(!is_array($args['ips'])){
            $return['code'] = 1;
            $return['msg'] = "parameter [ips] is not a Array";
            return $return;
        }

        return $return;
    }


    /*
     * 自动变更的入口函数
     */
    public function addNode($name, $gid, $ips, $port, $weight){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ['is_consul' => false]];

        $ret = $this->upstreamTbl->field('content,is_consul')
            ->where(['name' => $name, 'group_id' => $gid])
            ->find();

        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = 'get upstream content from db failed: '. $this->upstreamTbl->getDbError();
            return $return;
        }

        if(empty($ret)){
            $return['code'] = 1;
            $return['msg'] = "there is no such upstream file [$name:$gid]";
            return $return;
        }

        $oldUpstreamTmp = $ret['content'];
        $oldUpstream = explode("\n", $ret['content']);
        $newUpstream = $this->_fileAddNode($oldUpstream, $ips, $port, $weight);


        $this->upstreamTbl->startTrans();
        $ret = $this->upstreamTbl
            ->where(['name' => $name, 'group_id' => $gid])
            ->save([
                'content' => implode("\n", $newUpstream),
                'update_time' => date("Y-m-d H:i:s"),
            ]);
        $success = true;
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = 'save new upstream to db failed: '. $this->upstreamTbl->getDbError();
            $success = false;
        }

        // ------- 对 consul变更的处理 -----------
        if($ret['is_consul'] == true){
            $return['content']['is_consul'] = true;

            $consul = new Consul();
            $tmp = $consul->addNode($name, $gid, $ips, $port, $weight);
            if($tmp['code'] != 0) $success = false;
        }
        if($success == false){
            $this->upstreamTbl->rollback();
            return $return;
        }

        $this->upstreamTbl->commit();
        /*
         * the code that write upstream content to file has been deprecated.
         * Because we don't need that anymore. Every time we take it from db for now.
         *          reposkeeper
         */
        return $return;

    }

    // 删除节点
    public function delNode($name, $gid, $ips){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ['is_consul' => false]];

        $ret = $this->upstreamTbl->field('content')
            ->where(['name' => $name, 'group_id' => $gid])
            ->find();

        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = 'get upstream content from db failed: '. $this->upstreamTbl->getDbError();
            return $return;
        }

        if(empty($ret)){
            $return['code'] = 1;
            $return['msg'] = "there is no such upstream file [$name:$gid]";
            return $return;
        }

        $oldUpstreamTmp = $ret['content'];
        $oldUpstream = explode("\n", $ret['content']);
        $newUpstream = $this->_fileDelNode($oldUpstream, $ips);
        if(empty($newUpstream)){
            $return['code'] = 1;
            $return['msg'] = 'upstream left no ip after delete, failed!';
            return $return;
        }
        $this->upstreamTbl->startTrans();
        $ret = $this->upstreamTbl
            ->where(['name' => $name, 'group_id' => $gid])
            ->save([
                'content' => implode("\n", $newUpstream),
                'update_time' => date("Y-m-d H:i:s"),
            ]);
        $success = true;
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = 'save new upstream to db failed: '. $this->upstreamTbl->getDbError();
            $success = false;
        }
        // ------- 对 consul变更的处理 -----------
        if($ret['is_consul']){
            $return['content']['is_consul'] = true;

            $consul = new Consul();
            $tmp = $consul->delNode($name, $gid, $ips);
            if($tmp['code'] != 0) $success = false;
        }

        if($success == false){
            $this->upstreamTbl->rollback();
            return $return;
        }

        $this->upstreamTbl->commit();
        /*
         * the code that write upstream content to file has been deprecated.
         * Because we don't need that anymore. Every time we take it from db for now.
         *          reposkeeper
         */

        return $return;

    }

    public function getUpstreamNamesByUnitId($id){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $mainConfName = "nginx.conf";

        $subQuery = "SELECT unit_id,name, max(version) AS mv FROM tbl_hubble_nginx_conf_main ";
        $subQuery .= "WHERE unit_Id = '$id' AND NAME = '$mainConfName' GROUP BY NAME ";

        $query  = "SELECT m.id,m.content,m.version FROM tbl_hubble_nginx_conf_main m ";
        $query .= "INNER JOIN ($subQuery) t ON m.name = t.name AND m.version = t.mv AND m.unit_id = t.unit_id";

        $model = M();
        $ret = $model->query($query);
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = "db error:".$model->getDbError();
            return $return;
        }
        if(empty($ret)){
            return $return;
        }

        $upstreams = [];
        foreach($ret as $line){
            if(preg_match('/^\sinclude upstream/(.*);.*/', $line, $matches) == 1){
                $upstreams[] = $matches[1];
            }
        }

        $return['content'] = $upstreams;
        return $return;
    }

    /*
     * 直接查询数据库unit 和 group 和 conf main
     */
    public function getUnitNamesByUpstreamId($id){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->getUpstreamDetail($id);
        if($ret['code'] != 0) return $ret;
        $groupId = $ret['content']['group_id'];
        $upstreamName = $ret['content']['name'];
        $mainConfName = "nginx.conf";

        $subQuery  = "SELECT m.*,max(version) AS mv FROM tbl_hubble_nginx_conf_main m ";
        $subQuery .= "INNER JOIN tbl_hubble_nginx_unit u ON m.unit_id = u.id ";
        $subQuery .= "WHERE u.group_id = '$groupId' AND m.name = '$mainConfName' AND m.deprecated = 0 GROUP BY NAME,unit_id ";

        $query  = "SELECT m.unit_id,m.id,m.NAME,tmp.mv FROM `tbl_hubble_nginx_conf_main` m ";
        $query .= "INNER JOIN ($subQuery) tmp ON m.version = tmp.mv AND m.name = tmp.name AND m.unit_id = tmp.unit_id";

        $model = M();
        $ret = $model->query($query);
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = "db error:".$model->getDbError();
            return $return;
        }
        if(empty($ret)){
            return $return;
        }

        $files = [];
        foreach($ret as $item){
            if(strpos($item['content'], $upstreamName) !== false)
                $files[] = $item['name'];
        }

        $return['content'] = $files;
        return $return;
    }

    /*
     *
     */
    public function callTunnel($script_id, $filename, $user, $is_group, $group_id, $ids = '',$cor_id=''){
        // 准备脚本内容

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];


        $shell = new Shell();
        $ret = $shell->getShellDetail($script_id);
        if($ret['code'] != 0) return $ret;
        $script = [['action' => [
            "module" => "longscript",
            "content" => $ret['content']['content']]]];

        // 准备脚本参数
        $host  = "http://".C('HUBBLE_HOST').':'.C('HUBBLE_PORT');
        $host .= "/v1/nginx/upstream/content/?action=upstream";
        $scriptArg = [
            'HUBBLE_FILE_COUNT' => 1,
            'HUBBLE_GROUP_ID'   => $group_id,
            'HUBBLE_FILE_NAMES' => $filename,
            'HUBBLE_HOST'       => $host
        ];

        // 获取 reload nginx 的ip 列表
        $node = new NodeModel();
        if($is_group)
            $ret = $node->getNodeIpsByGroupId($group_id);
        else
            $ret = $node->getNodeIpsByUnitIds($ids);

        if ($ret['code'] != 0) return $ret;

        $nginxIps = $ret['content'];

        // 启动一个任务
        $channel = new Channel();
        $task = $channel->ansible($nginxIps, 'root', $script,$scriptArg,1,$cor_id);
        if($task['code'] != 0) return $task;
        $return['content'] = $task['content'];

        return $return;
    }

    public function publishManuel($upstream_id, $unit_ids, $script_id, $user, $tunnel = 'ANSIBLE'){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        $upstream = $this->getUpstreamDetail($upstream_id);
        if($upstream['code']!= 0) return $upstream;

        $group_id = $upstream['content']['group_id'];
        $name = $upstream['content']['name'];

        $task =  $this->callTunnel($script_id, $name, $user, false, $group_id, $unit_ids);
        if($task['code'] != 0) return $task;
        // 记录变更表
        $task = $task['content'];
        $history = new AlterationHistory();
        $ret = $history->addRecord('async', $task['ansible_id'], $task['ansible_name'], 'ansible', $user);
        if($ret['code'] != 0) return $ret;
        hubble_log(HUBBLE_INFO, 'task already sent, record id: '.json_encode($ret['content']));

        $this->upstreamTbl->where(['id' => $upstream_id])->save(['release_id' => $ret['content']]);
        $return['content'] = [
            'type' => 'async',
            'task_id' => $ret['content'],
        ];

        return $return;
    }
}
