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
 * Date: 16/9/2
 * Time: 下午1:51
 */

namespace V1\Controller\Adaptor;

use Common\Dao\Adaptor\Adaptor;
use Common\Dao\Adaptor\AlterationHistory;
use Common\Dao\Adaptor\AlterationType;
use Common\Dao\Adaptor\Channel;
use Common\Dao\Nginx\NodeModel;
use Common\Dao\Nginx\UnitModel;
use Think\Controller\RestController;


class AutoAlterationController extends RestController {

    private $input;

    public function __construct()
    {
        parent::__construct();

        $this->input = hubble_parse_param();
        if(! IS_GET && empty($this->input)){
            $this->ajaxReturn(std_error('parameter is empty'));
        }

        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }

    public function addNode_post(){

        $params = ['sid', 'ips', 'user', ];
        foreach ($params as $p){
            if(!isset($this->input[$p]) || empty($this->input[$p]))
                $this->ajaxReturn(std_error("parameter [$p] is absent or empty, please check and try again."));
        }
        $ips = $this->input['ips'];
        $user = $this->input['user'];
        $sid['id'] = $this->input['sid'];

        if(isset($sid['id']) && !empty($sid['id'])){
            $content=M('AlterationType')
                ->where($sid)
                ->getField('content');
            $data=json_decode($content, true);
            $condition['group_id'] = $data["group_id"];
            $uid=M('NginxUnit')
                ->where($condition)
                ->getField('id');
            $unit_id=$uid;
        }


        if($unit_id <= 0 ){
            $this->ajaxReturn(std_error('unit_id error'));
        }


        //检查是否存在单元


        $unit = new UnitModel() ;

        $filter = [];
        $filter['id'] = $unit_id;
        $ret = $unit->existsUnit($filter);

        if($ret['code'] != 0){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        //添加
        $ip = array_unique(explode(",",$ips));
        $data = array();

        foreach($ip as $v ){
            $data[] = $v;
        }
        $filer = [];
        $arr = [];

        $filer['ip'] = ['in' , $data];

        $node = new NodeModel() ;
        //检查
        $check = $node->existsNode($filer);
        //错误
        if($check['code'] == 2){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        //存在
        if($check['code'] == 0){
            foreach($check['content'] as $v ){
                $arr[] = $v['ip'];
            }
        }

        //判断diff
        $diff = array_diff($data,$arr) ;
        $intersect = array_intersect($data,$arr);

        if(empty($diff)){
            $msg = "exits:".implode(",",$intersect);
            $this->ajaxReturn(std_error($msg));
        }
        $ret = $node->addNode($unit_id,$user,array_diff($data,$arr));

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }

        hubble_oprlog('Nginx', 'Add node', I('server.HTTP_APPKEY'), $user, "ips:".json_encode($data).'exits:'.json_encode($arr));
        if(empty($intersect)){
            $this->ajaxReturn(std_return($ret['msg']));
        }else{
            $msg = $ret['msg']." exits:".implode(",",$intersect);
            $this->ajaxReturn(std_return(array(),$msg));
        }



    }


    public function deleteNode_post(){

        $params = ['ips', 'user', ];
        foreach ($params as $p){
            if(!isset($this->input[$p]) || empty($this->input[$p]))
                $this->ajaxReturn(std_error("parameter [$p] is absent or empty, please check and try again."));
        }
        $ip['ip'] = $this->input['ips'];
        $user = $this->input['user'];

        if(isset($ip['ip']) && !empty($ip['ip'])) {
            $nodes = M('NginxNode')
                ->where($ip)
                ->getField('id');
            $uid = M('NginxNode')
                ->where($ip)
                ->getField('unit_id');
            if($nodes==Null || $uid==Null){
                $ret=array('code'=>0,'msg'=>"success");
                hubble_oprlog('Nginx', 'del node', I('server.HTTP_APPKEY'), $user, "id:unit_id  nodes:".json_encode($nodes));
                $this->ajaxReturn(std_return($ret['msg']));
            }
        }


        $filter = [];
        $where = [];
        $where['id'] = ['in',$nodes];

        $filter['unit_id'] = $uid;
        $node = new NodeModel() ;

        $ret = $node->deleteNode($filter,$where);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        }
        hubble_oprlog('Nginx', 'del node', I('server.HTTP_APPKEY'), $user, "id:unit_id  nodes:".json_encode($nodes));
        $this->ajaxReturn(std_return($ret['msg']));


    }

    public function add_post(){

        $params = ['type_id', 'ips', 'user', ];

        foreach ($params as $p){
            if(!isset($this->input[$p]) || empty($this->input[$p]))
                $this->ajaxReturn(std_error("parameter [$p] is absent or empty, please check and try again."));
        }
        $ipStr = $this->input['ips'];
        $this->input['ips'] = explode(',', $this->input['ips']);

        $adaptor = new Adaptor();
        $ret = $adaptor->doAddNode($this->input['type_id'], $this->input, $this->input['user']);
        if($ret['code'] == 0){
            hubble_log(HUBBLE_INFO, 'auto alteration add success'.json_encode($ret['content']));
            hubble_oprlog('Adaptor', 'auto alteration add success',
                I('server.HTTP_APPKEY'), $this->input['user'], "type_id:{$this->input['type_id']}, ips:$ipStr");
            $this->ajaxReturn(std_return($ret['content']));
        }
        else{
            hubble_log(HUBBLE_WARN, $ret['msg']);
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function remove_post(){

        $params = ['type_id', 'ips', 'user', ];

        foreach ($params as $p){
            if(!isset($this->input[$p]) || empty($this->input[$p]))
                $this->ajaxReturn(std_error("parameter [$p] is absent or empty, please check and try again."));
        }
        $ipStr = $this->input['ips'];
        $this->input['ips'] = explode(',', $this->input['ips']);

        $adaptor = new Adaptor();
        $ret = $adaptor->doDelNode($this->input['type_id'], $this->input, $this->input['user']);
        if($ret['code'] == 0){
            hubble_log(HUBBLE_INFO, 'auto alteration del success'.json_encode($ret['content']));
            hubble_oprlog('Adaptor', 'auto alteration del',
                I('server.HTTP_APPKEY'), $this->input['user'], "type_id:{$this->input['type_id']}, ips:$ipStr");
            $this->ajaxReturn(std_return($ret['content']));
        }
        else{
            hubble_log(HUBBLE_WARN, $ret['msg']);
            $this->ajaxReturn(std_error($ret['msg']));
        }

    }

    public function check_state_get(){

        $gid = I('server.HTTP_X_CORRELATION_ID');
        $rid = I('release_id');

        if(empty($gid) && empty($rid))
            $this->ajaxReturn(std_error('correlation-id and release_id are both empty'));

        $record =  new AlterationHistory();

        if(!empty($rid)){// if there is release_id, use it in first
            $ret = $record->exist($rid);
            $gid = $ret['content']['global_id'];
        }
        else
            $ret = $record->existGid($gid);


        if($ret['code'] == 1){
            $this->ajaxReturn(std_error('check task state: '.$ret['msg']));
        }

        $ret = $ret['content'];
        if($ret['type'] == 'sync')
            $this->ajaxReturn(std_return(['task_id'=>$ret['task_id']]));

        switch(strtoupper($ret['channel'])){
            case 'ANSIBLE':

                $channel = new Channel();
                $result = $channel->ansibleCheck($ret['task_name']);

                if($result['code'] != 0){
                    $this->ajaxReturn(std_error("http: ".$result['error']));
                }

                $data = json_decode($result['data'],true);
                if(empty($data))
                    $this->ajaxReturn(std_error('wrong json format'));

                if($data['code'] != 0)
                    $this->ajaxReturn(std_error("ansible: ".$data['message']));

                $content = [];
                $content['state'] = $data['content']['task']['status'];

                if(!empty($data['content']['nodes'])){
                    foreach($data['content']['nodes'] as $v){
                        $content['detail'][] = [
                            'ip'=> $v['ip'],
                            'state'=> $v['status']];
                    }
                }
                $content['X-CORRELATION-ID'] = $gid;
                $this->ajaxReturn(std_return($content));
                break;

            default:
                $this->ajaxReturn(std_error('no such channel to deal with.'));
        }
    }


    public function type_param_get(){
        $typeArg = I('type');

        if(empty($typeArg))
            $this->ajaxReturn(std_error('type is empty'));

        $alteration = new AlterationType();

        $ret = $alteration->getTypeColumns($typeArg);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        } else{
            $this->ajaxReturn(std_return($ret['content']));
        }
    }

}
