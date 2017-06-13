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
 * Time: 下午12:59
 */
namespace Common\Dao\Nginx;
class GroupModel{

    private $table ;

    public function __construct()
    {
        $this->table= M('NginxGroup');
    }

    //添加分组
    public function addGroup($name,$user,$bid){

        $data = [
            'name' => $name,
            'opr_user' => $user,
            'biz_id' => $bid,
            'create_time' => date("Y-m-d H:i:s"),
        ];
        $ret = $this->table->add($data);
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        return array('code'=>0,'msg'=>"success",'content'=>array('gid'=>$ret));

    }
    //修改分组
    public function updateGroup($id, $user, $name){

        $ret = $this->table->where(['id'=>$id])
            ->save(['name' => $name, 'user' => $user, 'create_time' => date("Y-m-d H:i:s")]);
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        return array('code'=>0,'msg'=>"success");

    }
    //删除分组
    public function deleteGroup($id,$bid){

        $ret = $this->table->where(['id'=>$id, 'biz_id' => $bid])->select();
        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }
        if(empty($ret)){
            return array('code'=>1 ,'msg'=>"no such group") ;
        }

        $ret  = $this->table->where(['id'=>$id])->delete();

        if($ret === false){
            hubble_log(HUBBLE_ERROR, $this->table->getLastSql().' ERROR: '. $this->table->getDbError());
            return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
        }

        return array('code'=>0,'msg'=>"success");
    }


   public function existsGroup($id,$bid){

       $ret = $this->table->where(['id'=>$id, 'biz_id' => $bid])->select();
       if($ret === false){
           return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
       }
       if(empty($ret)){
           return array('code'=>1 ,'msg'=>"no such content") ;
       }
       return array('code'=>0,'msg'=>"success");
   }

   public function existsGroupName($name,$bid){
       $ret = $this->table->where(['name'=>$name, 'biz_id' => $bid])->select();
       if($ret === false){
           return array('code'=>1 ,'msg'=>"db error: {$this->table->getDbError()}") ;
       }
       if(empty($ret)){
           return array('code'=>0 ,'msg'=>"no such content") ;
       }
       return array('code'=>1,'msg'=>"$name exists");
   }

    public function getDetail($where){

        $ret = $this->table
            ->where($where)
            ->find();

        $return = ['code' => 0, 'msg' => 'success', 'content' => ''];
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

   //查询
   public function listGroup($filter,$page,$limit,$like=true){

       $return = array('code'=>0,'msg'=>'success');

       foreach($filter as $k => $v ){
           if($k == 'name' && $like){
               $filter[$k] = ['LIKE', "%$v%"];
           }

       }

       //数量
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

}
