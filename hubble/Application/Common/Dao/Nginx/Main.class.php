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
 * Date: 16/9/1
 * Time: 下午5:44
 */

namespace Common\Dao\Nginx;


class Main
{

    private $mainTbl;
    private $mainTblName = 'tbl_hubble_nginx_conf_main';

    function __construct()
    {

        $this->mainTbl = M('NginxConfMain');
    }

    /*
     * 计算给定条件下的main conf 的个数
     *
     * @param $where Array 表示需要过滤的字段名和值
     * @param $like bool 如果是 true 则会对name字段使用模糊匹配,反之则不
     *
     * @return mixed 成功返回 数值, 数据库错误返回 false
     */
    public function countMain($where, $like = true)
    {

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->mainTbl->where($where)->count();


        if ($ret === NULL) {
            return 0;
        } elseif ($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
            return false;
        } else {
            return (int)$ret;
        }
    }

    /*
     * 获取Main conf 的列表 支持字段过滤和 name的模糊匹配
     */
    public function getMainList($where, $page, $limit, $like = true)
    {

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->mainTbl
            ->where($where)->page($page, $limit)->select();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->mainTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    public function countMainSingleVersion($where, $like = true)
    {

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }
        $table = $this->mainTblName;
        $subQuery = $this->mainTbl->field('name,unit_id,max(version) as mv')->where($where)->group('name')->fetchSql(true)->select();

        $query  = "SELECT count(*) as count FROM $table AS m RIGHT JOIN ($subQuery) AS t ";
        $query .= "ON t.name = m.name and t.unit_id = m.unit_id AND t.mv = m.version";

        $ret = $this->mainTbl->query($query);

        if ($ret === NULL) {
            return 0;
        } elseif ($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
            return false;
        } else {
            return (int)$ret[0]['count'];
        }
    }

    /*
     * 获取Main conf 的列表 支持字段过滤和 name的模糊匹配
     */
    public function getMainListSingleVersion($where, $page, $limit, $like = true)
    {

        foreach ($where as $k => $v) {

            if ($k == 'name' && $like) {
                $where[$k] = ['LIKE', "%$v%"];
            }
        }


        $table = $this->mainTblName;
        $subQuery = $this->mainTbl->field('name,unit_id,max(version) as mv')->where($where)->group('name')->fetchSql(true)->select();

        $query  = "SELECT m.* FROM $table AS m RIGHT JOIN ($subQuery) AS t ";
        $query .= "ON t.name = m.name and t.unit_id = m.unit_id AND t.mv = m.version ";
        $query .= "order by create_time desc ";
        $query .= "limit ". ($page-1)*$limit ."," .$page*$limit;

        $ret = $this->mainTbl->query($query);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->mainTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    public function isExist($name, $unit_id, $version){
        $ret = $this->mainTbl->field('id')
            ->where(['name' => $name, 'unit_id' => $unit_id, 'version' => $version])
            ->find();

        if($ret === false) return false;
        if(empty($ret)) return null;
        return true;
    }

    /*
     * 获取一个main conf 的具体信息
     */
    public function getMainDetail($id)
    {

        $ret = $this->mainTbl
            ->where(['id' => $id])
            ->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->mainTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    /*
     * 添加一个main conf ,或者添加一个新版本的main conf
     */
    public function addMain($name, $content, $unitId, $version, $user)
    {

        $data = [
            'name' => $name,
            'content' => $content,
            'unit_id' => $unitId,
            'version' => $version,
            'deprecated' => 0,
            'create_time' => date("Y-m-d H:i:s"),
            'update_time' => date("Y-m-d H:i:s"),
            'opr_user' => $user,
        ];

        $ret = $this->mainTbl->add($data);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === NULL) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->mainTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;


    }

    /*
     * 获取当前name 和 group的 conf 的下一个 version 数值
     */
    public function getNextVersion($name, $unit_id){

        $ret = $this->mainTbl
            ->field('version')
            ->where(['name' => $name, 'unit_id' => $unit_id])
            ->order('version desc')
            ->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];


        if ($ret === NULL) {
            return 1;
        } elseif ($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
            return false;
        } else {
            $return['content'] = $ret;
        }
        return $ret['version']+1;
    }

    /*
     * 设置conf 为 废弃状态
     */
    public function setDeprecated($id)
    {

        $ret = $this->mainTbl
            ->where(['id' => $id])
            ->setField('deprecated', 1);

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->mainTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    /*
     * 根据 组ID 删除 conf
     */
    public function deleteMain($group_ids){
        $ret = $this->mainTbl
            ->where(['group_id' => ['IN', implode(',', $group_ids)]])
            ->delete();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->mainTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
        } else {
            $return['content'] = $ret;
        }
        return $return;
    }

    /*
     * 获取文件的内容 通过 文件名、版本、单元ID
     */
    public function getContentByNVU($name, $version, $unitId){

        $ret = $this->mainTbl->field('content')
            ->where([
                'name' => $name,
                'version' => $version,
                'unit_id' => $unitId,
            ])->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if ($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: ' . $this->mainTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->mainTbl->getLastSql() . ' ERROR: ' . $this->mainTbl->getDbError());
        } elseif (empty($ret)) {
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = "no such file. name:$name, version:$version, unit_id:$unitId";
        } else {
            $return['content'] = $ret['content'];
        }

        return $return;
    }
}
