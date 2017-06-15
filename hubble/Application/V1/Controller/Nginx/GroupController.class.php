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
 * Time: 下午1:38
 */
namespace V1\Controller\Nginx;

use \Common\Dao\Nginx\GroupModel;
use Common\Dao\Nginx\UnitModel;
use Common\Dao\Nginx\Upstream;
use Think\Controller\RestController;

class GroupController extends RestController{

    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }

    /*
     * 添加分组
     * @param string name 分组名
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function add_post(){

        $name = I('name','');
        $user = I('user','');
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if(empty($name)){
            $this->ajaxReturn(std_error('name is empty'));
        }

        if(empty($user)){
            $this->ajaxReturn(std_error('user is empty'));
        }

        if($bid < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        $group = new GroupModel() ;

        $ret = $group->existsGroupName($name,$bid);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $ret = $group->addGroup($name,$user,$bid);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        $dir  = C('HUBBLE_ROOT_DIR').C('HUBBLE_NGINX_DIR')."/group_".$ret['content']['gid'];

        if(!mkdir($dir,0755,true)){
            $this->ajaxReturn(std_error("mkdir group dir failed"));

        }
        hubble_oprlog('Nginx', 'Add group', I('server.HTTP_APPKEY'), $user, "name:$name, id:".$ret['content']['gid']);
        $this->ajaxReturn(std_return($ret['msg']));

    }

    public function detail_get(){
        $id = I('id',0);
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if($id <= 0){
            $this->ajaxReturn(std_error('id is error'));
        }

        if($bid < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $group = new GroupModel() ;
        $ret = $group->getDetail(['id' => $id, 'biz_id' => $bid]);

        if($ret['code'] != HUBBLE_RET_SUCCESS){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $this->ajaxReturn(std_return($ret['content']));

    }

    /*
     * 更新分组
     * @param string name 分组名
     * @param int  id   分组id
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function modify_put(){

        $gid = I('id',0);
        $name = I('name','');
        $user = I('user','');
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if($gid <= 0 ){
            $this->ajaxReturn(std_error('id is error'));
        }

        if(empty($name)){
            $this->ajaxReturn(std_error('name is empty'));
        }

        if(empty($user)){
            $this->ajaxReturn(std_error('user is empty'));
        }

        if($bid < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        //查找ID是否存在
        $group = new GroupModel();
        $ret = $group->existsGroup($gid,$bid);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        $ret = $group->existsGroupName($name, $bid);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        //更新
        $ret = $group->updateGroup($gid, $user, $name);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        hubble_oprlog('Nginx', 'Update group', I('server.HTTP_APPKEY'), $user, "name:$name, id:".$ret['content']['gid']);
        $this->ajaxReturn(std_return($ret['msg']));
        
    }

    /*
     * 删除分组
     * @param int  id   分组id
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function delete_delete(){
        $gid = I('id',0);
        $user = I('user','');
        $bid = I('server.HTTP_X_BIZ_ID',0);

        if($gid <= 0){
            $this->ajaxReturn(std_error('id is error'));
        }

        if(empty($user)){
            $this->ajaxReturn(std_error('user is empty'));
        }

        if($bid < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        $group = new GroupModel();

        //查找分组下单元
        $unit = new UnitModel();
        $filter = [];
        $filter['group_id'] = $gid;
        $filter['biz_id'] = $bid;

        $ret = $unit->existsUnit($filter);
        if($ret['code'] == 0 ){
            $this->ajaxReturn(std_error("there still have unit in group, delete that first"));
        }

        if($ret['code'] == 2 ){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        $upstream = new Upstream();
        if($upstream->countUpstream(['group_id' => $gid]) > 0)
            $this->ajaxReturn(std_error("there still have upstream in group, delete that first"));

        //删除
        $arr = $group->existsGroup($gid,$bid);
        if($arr['code'] != 0 )
            $this->ajaxReturn(std_error($ret['msg']));

        $ret = $group->deleteGroup($gid,$bid);
        if($ret['code'] == 1 ){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        $dir  = C('HUBBLE_ROOT_DIR').C('HUBBLE_NGINX_DIR')."/group_".$gid;

        rmdir_recursive($dir);


        hubble_oprlog('Nginx', 'Delete group', I('server.HTTP_APPKEY'), $user, " id:$gid content:".json_encode($arr));

        $this->ajaxReturn(std_return($ret['msg']));

    }

    /*
     * 获取分组 带分页
     * @param int  page   页数
     * @param int  limit   分页数量
     * @return mixed
     * 成功  int 0
     * 失败  int 1
     */
    public function list_get(){

        $page = I('page',1);
        $limit = I('limit',20);
        $gname = I('name','');
        $bid = I('server.HTTP_X_BIZ_ID',0);
        $like= I('like', true);

        if(empty($page) || $page <= 0 ){
            $this->ajaxReturn(std_error('page is error'));
        }

        if(empty($limit) || $limit <= 0){
            $this->ajaxReturn(std_error('limit is error'));
        }

        if($bid < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        $group = new GroupModel();
        $filter = ['biz_id' => $bid];
        if(!empty($gname)){
            $filter['name'] = $gname;
        }

        $ret = $group->listGroup($filter,$page,$limit,$like);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $this->ajaxReturn(std_return($ret['content']));

    }




}
