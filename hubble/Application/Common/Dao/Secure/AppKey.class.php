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
 * Date: 16/4/27
 * Time: 09:23
 */

namespace Common\Dao\Secure;

class AppKey{


    private $keyTable;
    private $interfaceTable;
    private $privilegeTable;
    private $root;
    private $keyTableName;
    private $interfaceTableName;
    private $privilegeTableName;


    function __construct(){
        $this->keyTable = M('SecureAppkey');
        $this->interfaceTable = M('SecureInterface');
        $this->privilegeTable = M('SecurePrivileges');

        $this->keyTableName = 'tbl_hubble_secure_appkey';
        $this->interfaceTableName = 'tbl_hubble_secure_interface';
        $this->privilegeTableName = 'tbl_hubble_secure_privileges';

        $this->root = C('HUBBLE_ROOT_APPKEY');
    }

    /*
     * 获取所有的权限key列表,可以使用 name 字段过滤
     *
     * @name  key的名称
     *
     * @return array
     *      code    状态  0 成功  1 无数据 100 数据库错误
     *      content  array 多行数据的集合
     */
    public function getAppkeyList($name = ''){

        $where = [];
        if(!empty($name)) $where = ['name' => $name];
        $ret = $this->keyTable->where($where)->select();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if($ret === NULL){

            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif($ret === false) {

            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->keyTable->getDbError();
            hubble_log(HUBBLE_ERROR, $this->keyTable->getLastSql() . " ERROR: ". $this->keyTable->getDbError());
        } else{

            $return['content'] = $ret;
        }

        return $return;

    }

    /*
     * 获取 Appkey的详细权限
     *
     * @param id 要获取详情的appkey的 id
     *
     */
    public function getPrivilegeDetails($id){

        $ret = M()->table($this->privilegeTableName . ' p')
                  ->where("p.appkey_id = '$id'")
                  ->join($this->interfaceTableName . ' i on p.interface_id = i.id')
                  ->select();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if($ret === NULL){

            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif($ret === false) {

            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->privilegeTable->getDbError();
            hubble_log(HUBBLE_ERROR, $this->privilegeTable->getLastSql() . " ERROR: ". $this->privilegeTable->getDbError());

        } else{
           $return['content'] = $ret;
        }

        return $return;
    }

    /*
     * 增加一个appkey
     */
    public function addAppkey($name, $desc)
    {


        if (!$this->existPrivilegeName($name)) return false;
        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];


        $other = [ 'extend' => 1, ];

        for ($i = 0; $i < 5; $i++) {  // 尝试 5次, 如果appkey仍有重复,则放弃
            $key = md5(time() . mt_rand(0, 10000));
            if (!$this->existPrivilegeAppkey($key))
                continue;
            else {
                $key = strtoupper($key);
                $ret = $this->keyTable->data(['name' => $name, 'appkey' => $key, 'describe' => $desc, 'other' => $other])->add();
                if ($ret === false) {
                    $return['code'] = HUBBLE_DB_ERR;
                    hubble_log(HUBBLE_ERROR, $this->keyTable->getLastSql() . " ERROR: ". $this->keyTable->getDbError());
                    $return['msg'] = 'db error: '.$this->keyTable->getDbError();
                    return $return;
                }
                $return['content'] = $key;
                return $return;
            }
        }
        $return['code'] = 1;
        $return['msg'] = 'try generate appkey for five times, all failed';
        hubble_log(HUBBLE_ERROR, "try generate appkey for five times, all failed");

        return $return;
    }

    /*
     * 删除appkey
     *
     * @param appkey 应用key
     *
     * @return true if success or false
     *
     */
    public function delAppkey($appkey)
    {

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->keyTable->where(['appkey' => $appkey])->delete();

        if ($ret === false) {

            $return['code'] = HUBBLE_DB_ERR;
            hubble_log(HUBBLE_ERROR, $this->keyTable->getLastSql() . " ERROR: ". $this->keyTable->getDbError());
            $return['msg'] = 'db error: '.$this->keyTable->getDbError();
            return $return;
        }

        return $return;

    }


    /*
     * 查看appkey是否存在
     */
    public function existPrivilegeAppkey($appkey)
    {
        return $this->_existPrivilege('appkey', $appkey);
    }

    /*
     * 查看appkey 的名字是否存在
     */
    public function existPrivilegeName($name)
    {
        return $this->_existPrivilege('name', $name);
    }

    /*
     * 查看一个某字段的值是否存在
     *
     * @param column  以哪一列查看
     * @param value    要查看的值
     *
     * @return true if exist or false
     */
    private function _existPrivilege($column, $value)
    {

        $ret = $this->keyTable->where([$column => $value])->find();

        if ($ret === false) {
            hubble_log(HUBBLE_ERROR,
                $this->keyTable->getDbError() . ':::::' . $this->keyTable->getLastSql());
            return false;
        }

        if ($ret === NULL) return true;

        return false;
    }

    /*
     * 检查权限
     * 使用了缓存, 5分钟
     *
     * @param appkey 应用钥匙
     * @param item  要检查的权限名
     * @param core  是否为核心权限
     *
     * @return true if has or false
     */

    public function checkPrivilege($appkey, $item)
    {
        if($appkey == $this->root) return true;

        $ret = M()->table($this->privilegeTableName . ' p')
                  ->join($this->keyTableName . " a on p.appkey_id = a.id")
                  ->join($this->interfaceTableName . " i on p.interface_id = i.id")
                  ->where(['appkey' => $appkey, 'addr' => $item])
                  ->find();

        if ($ret === NULL) return false;

        if ($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->privilegeTable->getLastSql());
            return false;
        }

        return true;

    }

    /*
     * 添加一个权限
     *
     * @param keyId  appkey的id
     * @param interfaceId  接口的id
     */
    public function addPrivilege($keyId, $interfaceId){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->privilegeTable->add(
            ['appkey_id' => $keyId, 'interface_id' => $interfaceId]
        );

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            hubble_log(HUBBLE_ERROR, $this->privilegeTable->getLastSql() . " ERROR: ". $this->privilegeTable->getDbError());
            $return['msg'] = 'db error: '.$this->privilegeTable->getDbError();
            return $return;
        }

        return $return;
    }

    /*
     * 删除一个权限
     *
     * @param keyId  appkey的id
     * @param interfaceId  接口的id
     */
    public function delPrivilege($keyId, $interfaceId){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->privilegeTable->where(['appkey_id' => $keyId, 'interface_id' => $interfaceId])->delete();

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            hubble_log(HUBBLE_ERROR, $this->privilegeTable->getLastSql() . " ERROR: ". $this->privilegeTable->getDbError());
            $return['msg'] = 'db error: '.$this->privilegeTable->getDbError();
            return $return;
        }

        return $return;
    }

    /*
     * 获取 接口列表
     *
     * @param addr 过滤 addr
     * @param page 第几页
     * @param limit 每页限制
     * @param like 是否启用模糊匹配
     *
     * @return 列表
     */
    public function getInterfaceList($filter, $page, $limit = 20, $like = true){

        foreach($filter as $k => $v){
            if($like && isset($filter['addr']))
                $filter['addr'] = ['LIKE', "%{$filter['addr']}%"];

        }

        if($filter)
            $ret = $this->interfaceTable->where($filter)->page($page, $limit)->select();
        else
            $ret = $this->interfaceTable->page($page, $limit)->select();


        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];


        if($ret === NULL){

            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif($ret === false) {

            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->interfaceTable->getDbError();
            hubble_log(HUBBLE_ERROR, $this->interfaceTable->getLastSql() . " ERROR: ". $this->interfaceTable->getDbError());

        } else{

            $return['content'] = $ret;
        }

        return $return;
    }

    /*
     * 获取 接口的数量
     */
    public function countInterface($filter, $like = true){

        foreach($filter as $k => $v) {
            if ($like && isset($filter['addr']))
                $filter['addr'] = ['LIKE', "%{$filter['addr']}%"];
        }
        if($filter)
            $ret = $this->interfaceTable->where($filter)->count();
        else
            $ret = $this->interfaceTable->count();

        if($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->interfaceTable->getLastSql() . " ERROR: ". $this->interfaceTable->getDbError());
            $ret = 0;
        }

        return (int)$ret;
    }

    /*
     * 删除一个接口
     */
    public function delInterface($id){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->interfaceTable->where(['id' => $id])->delete();

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            hubble_log(HUBBLE_ERROR, $this->interfaceTable->getLastSql() . " ERROR: ". $this->interfaceTable->getDbError());
            $return['msg'] = 'db error: '.$this->keyTable->getDbError();
            return $return;
        }

        return $return;
    }

    /*
     * 添加一个接口
     */
    public function addInterface($addr, $desc, $methodArg){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->interfaceTable->add(['addr' => $addr, 'desc' => $desc, 'method' => $methodArg]);

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            hubble_log(HUBBLE_ERROR, $this->interfaceTable->getLastSql() . " ERROR: ". $this->interfaceTable->getDbError());
            $return['msg'] = 'db error: '.$this->interfaceTable->getDbError();
            return $return;
        }

        return $return;
    }
}
