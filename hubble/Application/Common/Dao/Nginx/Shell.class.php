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
 * Date: 16/9/7
 * Time: 下午2:59
 */

namespace Common\Dao\Nginx;

class Shell {

    private $shellTbl;

    function __construct()
    {

        $this->shellTbl = M('NginxShell');
    }

    /*
     * 计算给定条件下的main conf 的个数
     *
     * @param $where Array 表示需要过滤的字段名和值
     * @param $like bool 如果是 true 则会对name字段使用模糊匹配,反之则不
     *
     * @return mixed 成功返回 数值, 数据库错误返回 false
     */
    public function countShell($where, $like = true)
    {

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->shellTbl
            ->where($where)
            ->count();

        if ($ret === NULL) {
            return 0;
        } elseif ($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->shellTbl->getLastSql() . ' ERROR: ' . $this->shellTbl->getDbError());
            return false;
        } else {
            return (int)$ret;
        }
    }

    /*
     * 获取Shell 的列表 支持字段过滤和 name的模糊匹配
     */
    public function getShellList($where, $page, $limit, $like = true)
    {

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->shellTbl
            ->where($where)
            ->page($page, $limit)
            ->select();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->shellTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->shellTbl->getLastSql() . ' ERROR: ' . $this->shellTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    /*
     * 获取一个main conf 的具体信息
     */
    public function getShellDetail($id)
    {

        $ret = $this->shellTbl
            ->where(['id' => $id])
            ->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if (empty($ret)) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such shell';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->shellTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->shellTbl->getLastSql() . ' ERROR: ' . $this->shellTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }


    /*
     * 添加一个main conf ,或者添加一个新版本的main conf
     */
    public function addShell($name, $desc, $content, $user){

        $data = [
            'name' => $name,
            'desc' => $desc,
            'content' => $content,
            'create_time' => date("Y-m-d H:i:s"),
            'update_time' => date("Y-m-d H:i:s"),
            'opr_user' => $user,
        ];

        $ret = $this->shellTbl->add($data);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->shellTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->shellTbl->getLastSql() . ' ERROR: ' . $this->shellTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    public function deleteShell($id){
        $ret = $this->shellTbl
            ->where(['id' => $id])
            ->delete();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->shellTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->shellTbl->getLastSql() . ' ERROR: ' . $this->shellTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    public function modifyShell($id, $name, $desc, $content, $user){

        $data = [
            'desc' => $desc,
            'content' => $content,
            'opr_user' => $user,
        ];

        $ret = $this->shellTbl
            ->where(['id' => $id])
            ->save($data);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->shellTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->shellTbl->getLastSql() . ' ERROR: ' . $this->shellTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }
}
