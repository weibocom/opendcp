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

        $page = I('page', 1);
        $limit = I('limit', 20);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));



        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('limit or page out of range'));

        // 设置过滤器
        $filter     = [];
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

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));


        $upstream = new Upstream();

        $ret = $upstream->getUpstreamDetail($idArg);

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

        if(empty($nameArg))
            $this->ajaxReturn(std_error('name is empty'));

        if(empty($contentArg))
            $this->ajaxReturn(std_error('content is empty'));

        if(empty($groupIdArg))
            $this->ajaxReturn(std_error('group is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        $consulArg = $consulArg == 0 ? false:true;

        $upstream = new Upstream();

        if($upstream->countUpstream(['name' => $nameArg, 'group_id' => $groupIdArg], false) !== 0)
            $this->ajaxReturn(std_error("name [$nameArg] is exist."));


        $ret = $upstream->addUpstream($nameArg, $contentArg, $groupIdArg, $consulArg, $userArg);

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

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));


        $upstream = new Upstream();

        $ret = $upstream->deleteUpstream($idArg);

        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Add upstream', I('server.HTTP_APPKEY'), $userArg, "id:$idArg");
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function modify_put(){
        $idArg = I('id');
        $contentArg = I('content', '', 'unsafe_raw');
        $userArg = I('user', '');

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($contentArg))
            $this->ajaxReturn(std_error('content is empty'));


        $upstream = new Upstream();

        $ret = $upstream->modifyUpstream($idArg, $contentArg);

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

        if(empty($nameArg) || empty($groupIdArg))
            $this->ajaxReturn(std_error('name or group_id is empty'));

        $upstream = new Upstream();

        $ret = $upstream->getUpstreamContent($groupIdArg, $nameArg);
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

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        $upstream = new Upstream();

        $ret = $upstream->getUnitNamesByUpstreamId($idArg);
        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function upstream_list_get(){
        $unitIdArg = I('unit_id');

        if(empty($unitIdArg))
            $this->ajaxReturn(std_error('unit_id is empty'));

        $upstream = new Upstream();

        $ret = $upstream->getUpstreamNamesByUnitId($unitIdArg);
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

        $shell = new Shell();
        $ret = $shell->getShellDetail($scriptIdArg);
        if($ret['code'] == HUBBLE_RET_NULL)
            $this->ajaxReturn(std_error('script id doese not exist'));

        $upstream = new Upstream();
        $ret = $upstream->publishManuel($upstreamIdArg, $unitIdArg, $scriptIdArg, $userArg, $tunnelArg);

        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }

    }
}
