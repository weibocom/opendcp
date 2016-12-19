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
 * User: yabo
 * Date: 16/9/9
 * Time: 下午1:44
 */
namespace Common\Dao\Adaptor;

class AlterationType {
    private $typeTbl;

    function __construct(){
        $this->typeTbl = M('AlterationType');
    }

    private $alterationTypeArgs = [
        'NGINX' => [
            'name' => 'upstream name that you want alteration',
            'group_id' => 'Group ID',
            'port'=> 'Port',
            'weight' => 'weight of new ip',
            'script_id' => 'script id of nginx-reload script'
        ],
        'SLB' => [
            'weight' => 'weight of new ip',
            'slb_id' => 'slb\'s id',
            'region' => 'ecs region'
        ],
    ];


    public function add($type,$name,$content,$opr_user){
        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $data = [
            'type' => $type,
            'name' => $name,
            'content' => $content,
            'opr_user' => $opr_user,
            'create_time' => date('Y-m-d H:i:s'),
            'update_time' =>  date('Y-m-d H:i:s'),
        ];

        $ret = $this->typeTbl->add($data);
        if($ret === false){
            $return['code'] = 1;
            $return['msg'] = "add alteration type failed, error: " .$this->typeTbl->getDbError();
            return $return;
        }
        hubble_log(HUBBLE_ERROR, $this->typeTbl->getLastSql().' ERROR: '. $this->typeTbl->getDbError());
        return $return;
    }

    public function remove($id){

        $ret = $this->typeTbl->where(['id' => $id])->delete();
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->typeTbl->getLastSql().' ERROR: '. $this->typeTbl->getDbError());
            return ['code'=>1 ,'msg'=>"db error: {$this->typeTbl->getDbError()}"] ;
        }

        if(is_null($ret)) {
            return ['code'=>1 ,'msg'=>"db error: {$this->typeTbl->getDbError()}"] ;
        }

        return ['code'=>0,'msg'=>"success"];

    }

    public function update($id,$data){

        $ret = $this->typeTbl->where(['id' => $id])->save($data);
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->typeTbl->getLastSql().' ERROR: '. $this->typeTbl->getDbError());
            return ['code'=>1 ,'msg'=>"db error: {$this->typeTbl->getDbError()}"] ;
        }

        if(is_null($ret)) {
            return ['code'=>1 ,'msg'=>"db error: {$this->typeTbl->getDbError()}"] ;
        }

        return ['code'=>0,'msg'=>"success"];

    }

    public function exist($filter){

        $return = ['code' => 1, 'msg' => 'name exist', 'content' => ''];
        $ret = $this->typeTbl->where($filter)->find() ;

        if($ret === false){
            $return['code'] = 2;
            $return['msg'] = "check alteration type failed, error: " .$this->typeTbl->getDbError();
            return $return;
        }

        if(empty($ret)){
            $return['code'] = 0;
            $return['msg'] = 'no such content';
            return $return;
        }
        $return['content'] = $ret;
        return $return;
    }

    public function getList($filter,$page,$limit){

        $return = array('code'=>0,'msg'=>'success');

        //数量
        $count = $this->typeTbl->where($filter)->count();
        //页数
        $total_page = ceil($count/$limit);
        //数据

        $ret  = $this->typeTbl->where($filter)->page($page, $limit)->select();

        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->typeTbl->getLastSql().' ERROR: '. $this->typeTbl->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->typeTbl->getDbError()}") ;
        }

        $return['content']['count'] = $count;
        $return['content']['page'] = $page;
        $return['content']['limit'] = $limit;
        $return['content']['total_page'] = $total_page;
        $return['content']['content'] = $ret;

        return  $return;
    }

    public function getTypeColumns($type){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        if(!isset($this->alterationTypeArgs[$type])){
            $return['code'] = 1;
            $return['msg'] = "we do not support type [$type] yet.";
            return $return;
        }
        $return['content'] = $this->alterationTypeArgs[$type];
        return $return;
    }

}
