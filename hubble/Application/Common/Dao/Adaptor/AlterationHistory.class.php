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
 * Date: 16/9/8
 * Time: 上午10:03
 */

namespace Common\Dao\Adaptor;


class AlterationHistory {

    private $historyTbl;

    function __construct(){
        $this->historyTbl = M('AlterationHistory');
    }

    public function addRecord($type, $task_id, $task_name, $channel, $user,$cor_id=''){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $data =[
            'type'        => $type,
            'task_id'     => $task_id,
            'task_name'   => $task_name,
            'channel'     => $channel,
            'global_id'   => I('server.HTTP_X_CORRELATION_ID')?I('server.HTTP_X_CORRELATION_ID'):$cor_id,
            'create_time' => date("Y-m-d H:i:s"),
            'opr_user'    => $user,
        ];

        $ret = $this->historyTbl->add($data);
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = "add alteration record failed, ERROR: " .$this->historyTbl->getDbError();
            return $return;
        }

        $return['content'] = $ret;
        return $return;
    }

    public function exist($id){
        $ret = $this->historyTbl->where(['id'=>$id])->find();

        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->historyTbl->getLastSql().' ERROR: '. $this->historyTbl->getDbError());
            return array('code'=>1,'msg'=>"db error: {$this->historyTbl->getDbError()}") ;
        }

        if(empty($ret)){
            return array('code'=>1,'msg'=>'no such content');
        }

        return array('code'=>0,'msg'=>"success",'content'=>$ret);

    }

    public function existGid($gid){
        $in['global_id']=$gid;
        $in['task_id']=array('neq',0);
        $ret = $this->historyTbl->where($in)->order('id desc')->limit(1)->select();
        $ret=$ret[0];

        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->historyTbl->getLastSql().' ERROR: '. $this->historyTbl->getDbError());
            return array('code'=>1,'msg'=>"db error: {$this->historyTbl->getDbError()}") ;
        }

        if(empty($ret)){
            return array('code'=>1,'msg'=>'no such content');
        }

        return array('code'=>0,'msg'=>"success",'content'=>$ret);

    }

    public function countHistory($where, $like = true){

        foreach($where as $k => $v){

            if($k == 'task_name' && $like){
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->historyTbl
            ->where($where)
            ->count();

        if($ret === NULL){
            return 0;
        } elseif($ret === false) {
            hubble_log(HUBBLE_ERROR, $this->historyTbl->getLastSql().' ERROR: '. $this->historyTbl->getDbError());
            return false;
        } else{
            return (int)$ret;
        }
    }


    public function getHistoryList($where, $page, $limit, $like = true){

        foreach($where as $k => $v){

            if($k == 'task_name' && $like){
                $where[$k] = ['LIKE', "%$v%"];
            }
        }

        $ret = $this->historyTbl
            ->where($where)
            ->page($page, $limit)
            ->select();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        if(empty($ret)){
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->historyTbl->getDbError();
            hubble_log(HUBBLE_ERROR, $this->historyTbl->getLastSql().' ERROR: '. $this->historyTbl->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;
    }
}
