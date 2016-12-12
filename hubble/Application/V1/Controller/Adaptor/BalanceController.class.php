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
 * User: yabo
 * Date: 16/9/9
 * Time: 下午1:37
 */
namespace V1\Controller\Adaptor;
use Think\Controller\RestController;
use \Common\Dao\Adaptor\AlterationType;

class BalanceController extends RestController{

    //private $input;

    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }


    public function add_post(){

        $name = I('name','');
        $type = I('type','');
        $content = I('content','', 'unsafe_raw');
        $opr_user = I('user','');

        if(empty($name)){
            $this->ajaxReturn(std_error('name is empty'));
        }

        if(empty($type)){
            $this->ajaxReturn(std_error('type is empty'));
        }

        if(empty($content)){
            $this->ajaxReturn(std_error('content is empty'));
        }

        if(empty($opr_user)){
            $this->ajaxReturn(std_error('opr_user is empty'));
        }

        $filter = [];
        $filter['name'] = $name;

        $AlterationType = new AlterationType();
        $ret = $AlterationType->exist($filter);
        if($ret['code'] != 0){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        $ret = $AlterationType->add($type,$name,$content,$opr_user);


        if($ret['code'] == 1) {
            $this->ajaxReturn(std_error($ret['msg']));
        }
        hubble_oprlog('Alteration Type', 'add', I('server.HTTP_APPKEY'), $opr_user, "type:$type content:".json_encode($content));
        $this->ajaxReturn(std_return($ret['msg']));

    }

    public function list_get(){

        $filter = [];
        $page = I('page',1);
        $limit = I('limit',20);
        $opr_user = I('user','');
        $name = I('name','');
        $type = I('type','');
        $like = I('like',true);

        if($page < 0 ){
            $this->ajaxReturn(std_error('page is empty'));
        }

        if($limit < 0){
            $this->ajaxReturn(std_error('limit error'));
        }

        if($like){
            if(!empty($name)){
                $filter['name'] = ['LIKE',"%$name%"];
            }
            if(!empty($user)){
                $filter['opr_user'] = ['LIKE',"%$user%"];
            }
            if(!empty($type)){
                $filter['type'] = ['LIKE',"%$type%"];
            }

        }else{
            if(!empty($name)){
                $filter['name'] =  $name;
            }
            if(!empty($opr_user)){
                $filter['opr_user'] = $opr_user;
            }
            if(!empty($type)){
                $filter['type'] = $type;
            }
        }

        $AlterationType = new AlterationType();

        $ret = $AlterationType->getList($filter,$page,$limit);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $this->ajaxReturn(std_return($ret['content']));



    }

    public function detail_get(){
        $id = I('id',0);
        if($id <= 0){
            $this->ajaxReturn(std_error("id is empty"));
        }

        $filter = [];
        $filter['id'] = $id;
        $AlterationType = new AlterationType();
        $ret = $AlterationType->exist($filter);
        if($ret['code'] != 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $this->ajaxReturn(std_return($ret['content']));


    }

    public function modify_put(){

        $id = I('id',0);
        $opr_user = I('user','');
        $name = I('name','');
        $type = I('type','');
        $content = I('content','','unsafe_raw');


        if(empty($id)){
            $this->ajaxReturn(std_error("id is empty"));
        }

        if(empty($opr_user)){
            $this->ajaxReturn(std_error("user is empty"));
        }

        if(empty($name) && empty($type) && empty($content)){
            $this->ajaxReturn(std_error("modify content is empty"));
        }

        $data = [];
        if(!empty($name)){
            $data['name']  = $name;
        }
        if(!empty($type)){
            $data['type']  = $type;
        }
        if(!empty($name)){
            $data['content']  = $content;
        }
        $data['opr_user']  = $opr_user;
        $data['update_time'] = date('Y-m-d H:i:s');

        $AlterationType = new AlterationType();
        $ret = $AlterationType->update($id,$data);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        hubble_oprlog('Alteration Type', 'update', I('server.HTTP_APPKEY'), $opr_user, json_encode($data));
        $this->ajaxReturn(std_return($ret['msg']));

    }

    public function delete_delete(){

        $id = I('id',0);
        $opr_user = I('user','');
        if(empty($opr_user)){
            $this->ajaxReturn(std_error('opr_user is empty'));
        }
        if($id <= 0 ){
            $this->ajaxReturn(std_error('id is empty'));
        }

        $AlterationType = new AlterationType();
        $result = $AlterationType->exist($id);

        if($result['code'] != 1){
            $this->ajaxReturn(std_error($result['msg']));
        }

        $ret = $AlterationType->remove($id);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        hubble_oprlog('Alteration Type', 'remove', I('server.HTTP_APPKEY'), $opr_user, $id);
        $this->ajaxReturn(std_return($ret['msg']));
    }
}
