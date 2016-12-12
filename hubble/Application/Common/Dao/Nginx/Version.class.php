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
 * Date: 16/9/5
 * Time: 上午8:37
 */

namespace Common\Dao\Nginx;

use Common\Dao\Adaptor\AlterationHistory;
use Common\Dao\Adaptor\Channel;

class Version {

    private $versionTbl;

    function __construct(){

        $this->versionTbl = M('NginxVersion');
    }

    public function isExist($id){

        $ret = $this->versionTbl->field('id')
            ->where("id = '$id'")
            ->find();

        if($ret === false) return false;

        if(empty($ret)) return null;

        return true;
    }

    public function countVersion($where, $like = true)
    {

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->versionTbl
            ->where($where)
            ->count();

        if ($ret === NULL) {
            return 0;
        } elseif ($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->versionTbl->getLastSql() . ' ERROR: ' . $this->versionTbl->getDbError());
            return false;
        } else {
            return (int)$ret;
        }
    }


    public function getVersionList($where, $page, $limit, $like = true){

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->versionTbl
            ->where($where)
            ->order('id desc')
            ->page($page, $limit)
            ->select();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->versionTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->versionTbl->getLastSql() . ' ERROR: ' . $this->versionTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    public function getVersionDetail($id)
    {

        $ret = $this->versionTbl
            ->where(['id' => $id])
            ->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->versionTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->versionTbl->getLastSql() . ' ERROR: ' . $this->versionTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    public function createVersion($uid, $name, $files, $type, $user){
        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $lastVersion = $this->getLastVersionNumber($uid);
        if($lastVersion === false) {
            $return['code'] = 1;
            $return['msg'] = 'get last version faild';
            return $return;
        }

        $data = [
            'name' => $name,
            'unit_id' => $uid,
            'version' => $lastVersion+1,
            'files' => $files,
            'deprecated' => 0,
            'create_time' => date("Y-m-d H:i:s"),
            'opr_user'    => $user,
            'is_release'  => false,
            'type'        => $type,
            'release_id'  => 0,
        ];

        $ret = $this->versionTbl->add($data);
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = 'generate version, db failed: '.$this->versionTbl->getDbError();
            return $return;
        }
        return $return;
    }

    public function generateVersion($uid, $name, $type, $user){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        $ret = $this->_generateVersionContent($uid);

        if($ret['code'] != 0) return $ret;

        $data = [
            'name' => $name,
            'unit_id' => $uid,
            'version' => $ret['content']['version'],
            'files' => json_encode($ret['content']['files']),
            'deprecated' => 0,
            'create_time' => date("Y-m-d H:i:s"),
            'opr_user'    => $user,
            'type'        => $type,
            'is_release'  => false,
            'release_id'  => 0,
        ];

        $ret = $this->versionTbl->add($data);
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = 'generate version, db failed: '.$this->versionTbl->getDbError();
            return $return;
        }
        return $return;
    }


    /*
     * 使用所有的最新文件产生一个版本的描述
     */
    private function _generateVersionContent($uid){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $lastVersion = $this->getLastVersionNumber($uid);
        if($lastVersion === false) {
            $return['code'] = 1;
            $return['msg'] = 'get last version faild';
            return $return;
        }

        $main = M('NginxConfMain');

        // 获取主配置的文件名及配置
        $mainConf = $main->field('name, max(version) as version')
            ->where(['unit_id' => $uid, 'deprecated' => 0])
            ->group('name')
            ->select();

        if($mainConf === false){
            $return['code'] = 1;
            $return['msg'] = 'get main conf version failed: '.$main->getDbError();
            return $return;
        }

        $content = [];
        foreach ($mainConf as $item){
            $item['is_changed'] = true;
            $content[] = $item;
        }

        $return['content'] = [
            'version' => $lastVersion+1,
            'files' => $content,
        ];

        return $return;
    }

    /*
     * 获取某个单元下的最新版本
     */
    public function getLastVersionNumber($uid){

        $ret = $this->versionTbl->field('max(version) as version')
            ->where(['unit_id' => $uid, 'deprecated' => 0])
            ->find();

        if(is_null($ret)) return 0;
        if($ret === false) return false;
        return $ret['version'];
    }

    public function setDeprecated($id)
    {

        $ret = $this->versionTbl
            ->where(['id' => $id])
            ->setField('deprecated', 1);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->versionTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->versionTbl->getLastSql() . ' ERROR: ' . $this->versionTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    public function prepareReleaseFiles($id){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $version = $this->versionTbl
            ->where("id = '$id'")
            ->find();

        if ($version === false) {
            $return['code'] = 1;
            $return['msg'] = 'db error: ' . $this->versionTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->versionTbl->getLastSql() . ' ERROR: ' . $this->versionTbl->getDbError());
            return $return;
        }
        if (empty($version)) {
            $return['code'] = 1;
            $return['msg'] = 'no such version';
            return $return;
        }
        hubble_log(HUBBLE_INFO, "release version : ".json_encode($version));

        $unit_id = $version['unit_id'];

        $unit = new UnitModel();

        $ret = $unit->getGroupId($unit_id);
        if($ret['code'] != 0)
            return $ret;

        $groupId = $ret['content']['group_id'];
        $rootDir = C('HUBBLE_ROOT_DIR').C('HUBBLE_NGINX_DIR')."/group_$groupId/unit_$unit_id";


        $currentDir = $rootDir.'/current';
        $rootDir .= '/newer';

        if(file_exists($rootDir)){
            hubble_log(HUBBLE_INFO, "dir [$rootDir] exist. reday to remove");
            rmdir_recursive($rootDir);
            // 检查是否删除成功
            if(file_exists($rootDir)){
                $return['code'] = 1;
                $return['msg'] = "$rootDir remove failed.";
                return $return;
            }
        }

        // 建立主配置文件夹
        if(!mkdir($rootDir.'/main', 0777, true)){
            $return['code'] = 1;
            $return['msg'] = "create dir $rootDir/main failed.";
            return $return;
        }


        $mainConf = new Main();
        $files = json_decode($version['files'], true);
        if($files === null){
            $return['code'] = 1;
            $return['msg'] = "files is not a valid json";
            return $return;
        }
        foreach($files as $item){
            $ret = $mainConf->getContentByNVU($item['name'], $item['version'], $unit_id);
            // 如果取文件错误,直接就中断发布的过程
            if($ret['code'] != 0)
                return $ret;
            $tmpName = "$rootDir/main/{$item['name']}";
            hubble_log(HUBBLE_INFO, "file [$tmpName] ready to write");
            $fileRet = file_put_contents($tmpName, $ret['content']);
            if(!$fileRet){
                $return['code'] = 1;
                $return['msg'] = "file_put_content: $tmpName write failed.";
                hubble_log(HUBBLE_INFO, "file [$tmpName] ready to write");

                return $return;
            }
        }
        hubble_log(HUBBLE_INFO, "version release write file success");

        rmdir_recursive($currentDir);
        // 检查是否删除成功
        if(file_exists($currentDir)){
            $return['code'] = 1;
            $return['msg'] = "remove current version dir $currentDir failed.";
            return $return;
        }

        if(!rename($rootDir, $currentDir)) {
            $return['code'] = 1;
            $return['msg'] = "rename current version dir $currentDir failed.";
            return $return;
        }

        $return['content'] = [
            'unit_id' => $unit_id,
            'group_id' => $groupId,
        ];
        return $return;

    }

    public function releaseVersion($id, $shellId, $user){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->versionTbl->where("id = '$id'")->find();
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = $this->versionTbl->getDbError();
            return $return;
        }
        if(empty($ret)){
            $return['code'] = 1;
            $return['msg'] = "no such version id:$id";
            return $return;
        }
        $ver = $ret;
        $alterType = $ver['type'];

        $ret = $this->prepareReleaseFiles($id);
        if($ret['code'] != 0) return $ret;

        hubble_log(HUBBLE_INFO, "file update success, ready to start alteration");

        $unitId = $ret['content']['unit_id'];
        $groupId = $ret['content']['group_id'];

        switch(strtoupper($alterType)){
            case 'ANSIBLE':
                // 准备脚本内容
                $shell = new Shell();
                $ret = $shell->getShellDetail($shellId);
                if($ret['code'] != 0) return $ret;
                $script = [['action' => [
                    "module" => "longscript",
                    "content" => $ret['content']['content']]]];

                // 准备脚本参数
                $scriptArg = [
                    'HUBBLE_GROUP_ID'   => $groupId,
                    'HUBBLE_UNIT_ID'    => $unitId,
                    'HUBBLE_RSYNC_HOST' => C('HUBBLE_HOST')."::hubble/".C('HUBBLE_NGINX_DIR'),
                ];

                // 获取 reload nginx 的ip 列表
                $node = new NodeModel();
                $ret = $node->getNodeIpsByGroupId($groupId);
                if($ret['code'] != 0) {
                    if($ret['code'] == 2)
                        $ret['msg'] = 'there is no ip under unit';

                    return $ret;
                }
                $nginxIps = $ret['content'];

                // 启动一个任务
                $channel = new Channel();
                $task = $channel->ansible($nginxIps, 'root', $script,$scriptArg,1);
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
                break;
            default:
                $return['code'] = 1;
                $return['msg'] = "no such type [$alterType]";
                return $return;
        }

        $ret = $this->versionTbl->where("id = '$id'")
                                ->save(['is_release'=>1, 'release_id' => $return['content']['task_id']]);
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = "release success, version state update failed";
            return $return;
        }

        return $return;

    }

}
