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
 * Date: 16/5/4
 * Time: 17:45
 */

namespace Common\Dao\Secure;

class Oprlog {

    private $oprlog;

    function __construct(){
        $this->oprlog = M('SecureOprlog');
    }

    /*
     * 获取任务类型列表
     *
     * @param where 条件
     * @param limit 一次获取多少条记录
     * @param page 分页的话,第几页
     *
     * @return array
     *      code    状态  HUBBLE_RET_SUCCESS 成功  HUBBLE_RET_NULL 无数据 HUBBLE_DB_ERR 数据库错误
     *      content  array 多行数据的集合
     */
    public function getOprlogList($where, $limit, $page = 1){


        if(isset($where['operation'])){
            $where['operation'] = "%{$where['operation']}%";
        }

        $ret = $this->oprlog->where($where)->page($page, $limit)->order('opr_time desc')->select();

        $return = ['code' => HUBBLE_RET_SUCCESS, 'CONTENT' => ''];

        if($ret === NULL){

            $return['code'] = HUBBLE_RET_NULL;
            $return['content'] = 'no such content';
        } elseif($ret === false) {

            $return['code'] = HUBBLE_DB_ERR;
            $return['content'] = $this->oprlog->getDbError();
            hubble_log(HUBBLE_ERROR, '获取脚本执行记录' . $this->oprlog->getLastSql());
        } else{

            $return['content'] = $ret;
        }

        return $return;
    }

    /*
     * 获取时间线的总数
     */
    public function countOprlog($where){
        if(isset($where['operation'])){
            $where['operation'] = "%{$where['operation']}%";
        }

        $ret = $this->oprlog->where($where)->count();


        if($ret === false) {
            hubble_log(HUBBLE_ERROR, '获取脚本执行记录' . $this->oprlog->getLastSql());
            $ret = 0;
        }

        return $ret;
    }

    /*
     * 添加一个操作时间点
     *
     * @param module 模块名
     * @param operation 操作名
     * @param appkey 操作appkey
     * @param opr_time 操作时间
     * @param user  用户名
     * @param args  参数
     *
     * @return mixed 成功返回数据ID 失败返回 false
     */
    public function addItem($module, $operation, $appkey, $user, $args = ''){

        $time = date("Y-m-d H:i:s");

        $data = [
            'module'     => $module,
            'operation'  => $operation,
            'appkey'     => $appkey,
            'opr_time'   => $time,
            'user'       => $user,
            'args'       => $args
        ];

        $ret = $this->oprlog->add($data);

        if($ret === false){
            hubble_log(HUBBLE_ERROR,$this->oprlog->getLastSql());
            return false;
        }
        return $ret;
    }

    /*
     * 删除 xxxx-xx-xx日期之前的记录, MC会超时后自动删除
     * @param date xxxx-xx-xx yy:yy:yy 格式的日期
     *
     * @return mixed 成功返回 删除的数量 失败返回 false
     */
    public function deleteItems($date){

        $ret = $this->oprlog->where(['opr_time' => ['LT', $date]])->delete();
        if($ret === false){
            hubble_log(HUBBLE_ERROR,$this->oprlog->getLastSql());
            return false;
        }
        return $ret;
    }
}
