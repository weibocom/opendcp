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
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

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

        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $filter = ['biz_id' => $bidArg];
        $filter['name'] = $name;

        $AlterationType = new AlterationType();
        $ret = $AlterationType->exist($filter);
        if($ret['code'] != 0){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        $ret = $AlterationType->add($type,$name,$content,$opr_user,$bidArg);


        if($ret['code'] == 1) {
            $this->ajaxReturn(std_error($ret['msg']));
        }
        hubble_oprlog('Alteration Type', 'add', I('server.HTTP_APPKEY'), $opr_user, "type:$type content:".json_encode($content));
        $this->ajaxReturn(std_return($ret['msg']));

    }

    public function list_get(){

        $page = I('page',1);
        $limit = I('limit',20);
        $opr_user = I('user','');
        $name = I('name','');
        $type = I('type','');
        $bidArg = I('server.HTTP_X_BIZ_ID',0);
        $like = I('like',true);

        if($page < 0 ){
            $this->ajaxReturn(std_error('page is empty'));
        }

        if($limit < 0){
            $this->ajaxReturn(std_error('limit error'));
        }

        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

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

        $filter = ['biz_id' => $bidArg];
        
        $AlterationType = new AlterationType();

        $ret = $AlterationType->getList($filter,$page,$limit);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        $this->ajaxReturn(std_return($ret['content']));



    }

    public function detail_get(){
        $id = I('id',0);
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if($id <= 0){
            $this->ajaxReturn(std_error("id is empty"));
        }
        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $filter = ['biz_id' => $bidArg, 'id' => $id];

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
        $bidArg = I('server.HTTP_X_BIZ_ID',0);
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

        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

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
        $ret = $AlterationType->update($id,$bidArg,$data);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        hubble_oprlog('Alteration Type', 'update', I('server.HTTP_APPKEY'), $opr_user, json_encode($data));
        $this->ajaxReturn(std_return($ret['msg']));

    }

    public function delete_delete(){

        $id = I('id',0);
        $opr_user = I('user','');
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if(empty($opr_user)){
            $this->ajaxReturn(std_error('opr_user is empty'));
        }
        if($id <= 0 ){
            $this->ajaxReturn(std_error('id is empty'));
        }
        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $AlterationType = new AlterationType();
        $result = $AlterationType->exist(['id' => $id, 'biz_id' => $bidArg]);

        if($result['code'] != 1){
            $this->ajaxReturn(std_error($result['msg']));
        }

        $ret = $AlterationType->remove($id,$bidArg);
        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        hubble_oprlog('Alteration Type', 'remove', I('server.HTTP_APPKEY'), $opr_user, $id);
        $this->ajaxReturn(std_return($ret['msg']));
    }
}
