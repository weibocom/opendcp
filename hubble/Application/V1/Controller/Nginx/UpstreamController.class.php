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
 * Date: 16/9/1
 * Time: 下午4:32
 */

namespace V1\Controller\Nginx;

use Common\Dao\Adaptor\AlterationHistory;
use Common\Dao\Nginx\Shell;
use Common\Dao\Nginx\Upstream;
use Think\Controller\RestController;

class UpstreamController extends RestController{

    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }


    public function list_get(){

        $nameArg = I('name');
        $idArg = I('group_id');
        $likeArg = I('like', true);
        $bidArg = I('server.X-BIZ-ID',0);
        $page = I('page', 1);
        $limit = I('limit', 20);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('limit or page out of range'));

        // 设置过滤器
        $filter     = ['biz_id' => $bidArg];
        if(!empty($nameArg))
            $filter['name'] = $nameArg;

        $filter['group_id'] = $idArg;

        // init
        $content = [
            'page'      => (int)$page,
            'limit'     => (int)$limit,
            'content'   => array(),
            'count'     => 0,
            'total_page'=> 0,
        ];

        $upstream = new Upstream();

        $ret = $upstream->countUpstream($filter);
        if($ret === false) {
            $this->ajaxReturn(std_error('db error'));
        }

        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }

        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $upstream->getUpstreamList($filter, $page, $limit, $likeArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $content['content'] = $ret['content'];
            $this->ajaxReturn(std_return($content));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }


    public function detail_get(){
        $idArg = I('id');
        $bidArg = I('server.X-BIZ-ID',0);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));
        
        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        $upstream = new Upstream();

        $ret = $upstream->getUpstreamDetail(['id' => $idArg ,'biz_id' => $bidArg]);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function add_post(){
        $nameArg = I('name');
        $contentArg = I('content', '', 'unsafe_raw');
        $groupIdArg = I('group_id');
        $consulArg = I('is_consul');
        $userArg = I('user');
        $bidArg = I('server.X-BIZ-ID',0);

        if(empty($nameArg))
            $this->ajaxReturn(std_error('name is empty'));

        if(empty($contentArg))
            $this->ajaxReturn(std_error('content is empty'));

        if(empty($groupIdArg))
            $this->ajaxReturn(std_error('group is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $consulArg = $consulArg == 0 ? false:true;

        $upstream = new Upstream();

        if($upstream->countUpstream(['name' => $nameArg, 'group_id' => $groupIdArg, 'biz_id' => $bidArg], false) !== 0)
            $this->ajaxReturn(std_error("name [$nameArg] is exist."));


        $ret = $upstream->addUpstream($nameArg, $contentArg, $groupIdArg, $consulArg, $userArg, $bidArg);

        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Add upstream', I('server.HTTP_APPKEY'), $userArg, "name:$nameArg, group:$groupIdArg, consul:$consulArg");
            $this->ajaxReturn(std_return());
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function delete_delete(){

        $idArg = I('id');
        $userArg = I('user', '');
        $bidArg = I('server.X-BIZ-ID',0);

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        $upstream = new Upstream();

        $ret = $upstream->deleteUpstream($idArg, $bidArg);

        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Delete upstream', I('server.HTTP_APPKEY'), $userArg, "id:$idArg");
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function modify_put(){
        $idArg = I('id');
        $contentArg = I('content', '', 'unsafe_raw');
        $userArg = I('user', '');
        $bidArg = I('server.X-BIZ-ID',0);

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($contentArg))
            $this->ajaxReturn(std_error('content is empty'));

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        $upstream = new Upstream();

        $ret = $upstream->modifyUpstream($idArg, $contentArg, $bidArg);

        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Add upstream', I('server.HTTP_APPKEY'), $userArg, "id:$idArg, content: $contentArg");
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function content_get(){
        $nameArg = I('name');
        $groupIdArg = I('group_id');
        $bidArg = I('server.X-BIZ-ID',0);

        if(empty($nameArg) || empty($groupIdArg))
            $this->ajaxReturn(std_error('name or group_id is empty'));
        
        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        $upstream = new Upstream();

        $ret = $upstream->getUpstreamContent($groupIdArg, $nameArg, $bidArg);
        if($ret['code'] == 1){
            header('HTTP/1.0 404 Not Found');
        } else{
            header("Expires: ".gmdate("D, d M Y H:i:s", time() + 3600000)." GMT");
            header('Content-type:application/octet-stream', true);
            echo($ret['content']);
        }
    }

    public function unit_list_get(){
        $idArg = I('id');
        $bidArg = I('server.X-BIZ-ID',0);
        
        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }

        $upstream = new Upstream();

        $ret = $upstream->getUnitNamesByUpstreamId($idArg, $bidArg);
        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function upstream_list_get(){
        $unitIdArg = I('unit_id');
        $bidArg = I('server.X-BIZ-ID',0);

        if(empty($unitIdArg))
            $this->ajaxReturn(std_error('unit_id is empty'));

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        $upstream = new Upstream();

        $ret = $upstream->getUpstreamNamesByUnitId($unitIdArg, $bidArg);
        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function publish_post(){
        $upstreamIdArg = I('upstream_id');
        $unitIdArg = I('unit_ids');
        $tunnelArg = I('tunnel', 'ANSIBLE');
        $scriptIdArg = I('script_id');
        $userArg = I('user');
        $bidArg = I('server.X-BIZ-ID',0);

        if(empty($unitIdArg))
            $this->ajaxReturn(std_error('unit_ids is empty'));

        if(empty($upstreamIdArg))
            $this->ajaxReturn(std_error('upstream_id is empty'));

        $tunnelArg =strtoupper($tunnelArg);
        if(!in_array_case($tunnelArg, ['ANSIBLE']))
            $this->ajaxReturn(std_error("tunnel [$tunnelArg] dose not exist"));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($scriptIdArg))
            $this->ajaxReturn(std_error('script_id is empty'));

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        $shell = new Shell();
        $ret = $shell->getShellDetail(['id' => $scriptIdArg ,'biz_id' => $bidArg]);
        if($ret['code'] == HUBBLE_RET_NULL)
            $this->ajaxReturn(std_error('script id doese not exist'));

        $upstream = new Upstream();
        $ret = $upstream->publishManuel($upstreamIdArg, $unitIdArg, $scriptIdArg, $userArg, $bidArg, $tunnelArg);

        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }

    }
}
