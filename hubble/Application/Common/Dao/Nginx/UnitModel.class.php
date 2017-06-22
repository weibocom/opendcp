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
 * Date: 16/9/1
 * Time: 下午2:39
 */
namespace Common\Dao\Nginx;

class UnitModel{

    private $table ;

    public function __construct()
    {
        $this->table= M('NginxUnit');
    }

    //添加单元
    public function addUnit($name,$gid,$user,$bid){

        $data = [
            'name' => $name,
            'group_id' =>$gid,
            'opr_user' => $user,
            'biz_id'  => $bid,
            'create_time' => date("Y-m-d H:i:s"),
        ];
        $ret = $this->table->where(['group_id'=>$gid])->add($data);
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>HUBBLE_RET_SUCCESS ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        //创建

        return array('code'=>HUBBLE_RET_SUCCESS,'msg'=>"success",'content'=>array('uid'=>$ret));


    }
    //修改单元
    public function updateUnit($id, $user, $gid='', $name=''){

        $ret = $this->table->where(['id'=>$id])->save(
            ['name' => $name, 'group_id' => $gid, 'user' => $user, 'create_time' => date("Y-m-d H:i:s")]);
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        return array('code'=>0,'msg'=>"success");

    }
    //获取GID
    public function getGid($uid,$bid){

        $ret = $this->table->where(['id'=>$uid,'biz_id'=>$bid])->select();
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>2 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        if(empty($ret)){
            return array('code'=>1 ,'msg'=>"no such content") ;
        }
        return array('code'=>0,'msg'=>"success",'content'=>array('gid'=>$ret[0]['group_id']));
    }
    //删除分组
    public function deleteUnit($id){

        $ret = $this->table->where(['id' => $id])->delete();
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        return array('code'=>0,'msg'=>"success");
    }

    public function getDetail($where){

        $ret = $this->table
            ->where($where)
            ->find();

        $return = ['code' => HUBBLE_RET_SUCCESS, 'msg' => 'success', 'content' => ''];
        if($ret === NULL){
            $return['code'] = HUBBLE_RET_NULL;
            $return['msg'] = 'no such content';
        } elseif($ret === false) {
            $return['code'] = HUBBLE_DB_ERR;
            $return['msg'] = 'db error: '.$this->table->getDbError();
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
        } else{
            $return['content'] = $ret;
        }
        return $return;
    }

    //查询单元
    public function existsUnit($filter){

        $ret = $this->table->where($filter)->select();
        // 0 成功 1失败 2 数据库连接错误
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>2 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        if(empty($ret)){
            return array('code'=>1 ,'msg'=>"no such comment") ;
        }

        return array('code'=>0,'msg'=>"success");

    }

    public function  existsUnitName($name,$gid,$bid){
        $ret = $this->table->where(['group_id'=>$gid,'name'=>$name,'biz_id'=>$bid])->select();

        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        if(empty($ret)){
            return array('code'=>0 ,'msg'=>"no such content") ;
        }
        return array('code'=>1,'msg'=>"$name exists");

    }
    //分页
    public function listUnit($filter,$page,$limit,$like=true){

        $return = array('code'=>0,'msg'=>'success');

        foreach($filter as $k => $v ){
            if($k == 'name' && $like){
                $filter[$k] = ['LIKE',"%$v%"];
            }
        }

        ///数量

        $count = $this->table->where($filter)->count();

        //页数
        $total_page = ceil($count/$limit);
        //数据

        $ret  = $this->table->where($filter)->page($page, $limit)->select();

        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }


        $return['content']['count'] = $count;
        $return['content']['page'] = $page;
        $return['content']['limit'] = $limit;
        $return['content']['total_page'] = $total_page;
        $return['content']['content'] = $ret;

        return  $return;


    }

    public function getGroupId($id){
        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];

        $ret = $this->table->field('group_id')
            ->where("id = '$id'")
            ->find();

        if($ret === false){
            $return['code'] = 1;
            $return ['msg'] = 'get group id failed: '.$this->table->getDbError();
            return $return;
        }

        if(is_null($ret)){
            $return['code'] = 1;
            $return ['msg'] = 'no unit id record matched';
            return $return;
        }

        $return['content']['group_id'] = $ret['group_id'];
        return $return;
    }


    public function isExist($unitId, $bid){

        $ret = $this->table->field('id')
            ->where(['id' => $unitId, 'biz_id' => $bid])
            ->find();

        if($ret === false) return false;

        if(empty($ret)) return null;

        return true;
    }

    public  function getNamesByIds($ids){

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
        $ret = $this->table->field('name')->where(['id', ['IN', $ids]])->select();

        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            $ret['code'] = 1;
            $ret['msg'] = "db error: ".$this->table->getDbError();
        }elseif(empty($ret)){
            $return['content'] = [];
        }else{
            $return['content'] = array_map(function($i){return $i['name'];}, $ret);
        }

        return $return;
    }


}
