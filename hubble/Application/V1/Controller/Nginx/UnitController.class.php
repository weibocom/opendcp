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
 * Time: 下午4:57
 */
namespace V1\Controller\Nginx;

use Common\Dao\Nginx\NodeModel;
use Common\Dao\Nginx\UnitModel;
use Common\Dao\Nginx\GroupModel;
use Think\Controller\RestController;

class UnitController extends RestController
{

    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }


    /*
     * 添加单元
     * @param string  name   单元名
     * @param int  group_id   分组ID
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function add_post(){

        $name = I('name','');
        $group_id = I('group_id',0);
        $user = I('user','');
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if($group_id <= 0 ){
            $this->ajaxReturn(std_error('group_id error'));
        }

        if(empty($name)){
            $this->ajaxReturn(std_error('name is empty'));
        }

        if(empty($user)){
            $this->ajaxReturn(std_error('user is empty'));
        }

        if($bid < 0){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }


        $unit = new UnitModel() ;
        $ret = $unit->existsUnitName($name,$group_id,$bid);
        if($ret['code'] != 0 ){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $ret = $unit->addUnit($name,$group_id,$user,$bid);

        if($ret['code'] == HUBBLE_DB_ERR){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        //创建目录
        $conf_main = C('HUBBLE_ROOT_DIR').C('HUBBLE_NGINX_DIR')."/group_".$group_id."/unit_".$ret['content']['uid']."/main";

        $main_dir = mkdir($conf_main,0755,true);

        if(!$main_dir){
            $this->ajaxReturn(std_error("mkdir $conf_main failed"));
        }

        hubble_oprlog('Nginx', 'Add unit', I('server.HTTP_APPKEY'), $user, " name:$name, group:$group_id");
        $this->ajaxReturn(std_return($ret['msg']));

    }

    public function detail_get(){
        $id = I('id',0);
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if($id <= 0){
            $this->ajaxReturn(std_error('id error'));
        }
        
        if($bid < 0){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        $unit = new UnitModel();
        $ret = $unit->getDetail(['id' => $id, 'biz_id' => $bid]);

        if($ret['code'] != HUBBLE_RET_SUCCESS){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $this->ajaxReturn(std_return($ret['content']));

    }

    /*
     * 更改单元 可更改名称和分组
     * @param string  name   单元名
     * @param int  group_id   分组ID
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function modify_put(){

        $id = I('id',0);
        $name = I('name','');
        $group_id = I('group_id',0);
        $user = I('user','');
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if($group_id <=0 ){
            $this->ajaxReturn(std_error('group_id error'));
        }

        if($id <=0 ){
            $this->ajaxReturn(std_error('id error'));
        }

        if(empty($name)  ){
            $this->ajaxReturn(std_error('name is empty'));
        }
        if(empty($user)  ){
            $this->ajaxReturn(std_error('user is empty'));
        }

        if($bid < 0){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        //检查ID是否存在
        if($group_id > 0 ){
            $group = new GroupModel();

            //检查分组是否存在
            $ret = $group->existsGroup($group_id,$bid);
            if($ret['code'] == 1){
                $this->ajaxReturn(std_error($ret['msg']));
            }
        }


        //更新
        $unit = new UnitModel();

        $ret = $unit->existsUnitName($name,$group_id,$bid);
        if($ret['code'] == 1 ){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $ret = $unit->updateUnit($id, $user, $group_id, $name);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        hubble_oprlog('Nginx', 'Update unit', I('server.HTTP_APPKEY'), $user, "name:$name, group:$group_id");
        $this->ajaxReturn(std_return($ret['msg']));



    }
    /*
     * 删除单元
     * @param int  id   单元id
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function delete_delete(){
        $unit_id = I('id',0);
        $user = I('user','');
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if($unit_id <=0 ){
            $this->ajaxReturn(std_error('id error'));
        }

        if(empty($user)){
            $this->ajaxReturn(std_error('user is empty'));
        }

        if($bid < 0){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        //检查单元下是否有节点
        $node = new NodeModel();
        $filter = [ 'unit_id' => $unit_id, 'biz_id' => $bid];
        
        $ret = $node->existsNode($filter);
        if($ret['code'] == 0){
            $this->ajaxReturn(std_error('node exists'));
        }
        if($ret['code'] == 2 ){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        //删除
        $unit = new UnitModel();
        $gret = $unit->getGid($unit_id,$bid);

        if($gret['code'] !== 0 ){
            $this->ajaxReturn(std_error($gret['msg']));
        }
        $ret = $unit->deleteUnit($unit_id);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        $dir = C('HUBBLE_ROOT_DIR').C('HUBBLE_NGINX_DIR')."/group_".$gret['content']['gid']."/unit_".$unit_id;

        rmdir_recursive($dir);
        hubble_oprlog('Nginx', 'Delete unit', I('server.HTTP_APPKEY'), $user, "id:$unit_id content:".json_encode($gret));
        $this->ajaxReturn(std_return($ret['msg']));

    }
    /*
     * 分页
     * @param int  page  页数
     * @param int  limit   分页数量
     * @param int  group_id   分组ID
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function list_get(){

        $page = I('page',1);
        $limit = I('limit',20);
        $group_id = I('group_id',0);
        $uname = I('name','');
        $bid = I('server.HTTP_X_BIZ_ID',0);
        $like  = I('like',true);

        if($page <=0 ){
            $this->ajaxReturn(std_error('page error'));
        }

        if($limit <=0 ){
            $this->ajaxReturn(std_error('limit error'));
        }

        if($bid < 0){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        $filter  = ['biz_id' => $bid];
        if(!empty($uname)){
            $filter['name'] = $uname;
        }

        if($group_id > 0){
            $filter['group_id'] = $group_id;
        }


        $unit = new UnitModel();

        $ret = $unit->listUnit($filter,$page,$limit,$like);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $this->ajaxReturn(std_return($ret['content']));

    }

}
