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
 * Time: 下午6:11
 */

namespace V1\Controller\Nginx;

use Common\Dao\Nginx\Main;
use Think\Controller\RestController;

class MainController extends RestController{

    function __construct(){

        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }


    public function list_get(){

        $nameArg = I('name');
        $idArg = I('unit_id');
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

        $filter['unit_id'] = $idArg;
        // init
        $content = [
            'page'      => (int)$page,
            'limit'     => (int)$limit,
            'content'   => array(),
            'count'     => 0,
            'total_page'=> 0,
        ];

        $main = new Main();

        $ret = $main->countMain($filter, $likeArg);
        if($ret === false) {
            $this->ajaxReturn(std_error('db error'));
        }

        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }

        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $main->getMainList($filter, $page, $limit, $likeArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $content['content'] = $ret['content'];
            $this->ajaxReturn(std_return($content));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function list_ver_get(){

        $nameArg = I('name');
        $idArg = I('unit_id');
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

        $filter['unit_id'] = $idArg;
        $filter['deprecated'] = 0;
        // init
        $content = [
            'page'      => (int)$page,
            'limit'     => (int)$limit,
            'content'   => array(),
            'count'     => 0,
            'total_page'=> 0,
        ];

        $main = new Main();

        $ret = $main->countMainSingleVersion($filter, $likeArg);
        if($ret === false) {
            $this->ajaxReturn(std_error('db error'));
        }

        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }

        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $main->getMainListSingleVersion($filter, $page, $limit, $likeArg);

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

        $main = new Main();

        $ret = $main->getMainDetail($idArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function add_post(){
        $nameArg = I('name');
        $contentArg = I('content', '', 'unsafe_raw');
        $unitIdArg = I('unit_id');
        $userArg = I('user');

        if(empty($nameArg))
            $this->ajaxReturn(std_error('name is empty'));

        if(empty($contentArg))
            $this->ajaxReturn(std_error('content is empty'));

        if(empty($unitIdArg))
            $this->ajaxReturn(std_error('unit_id is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));


        $main = new Main();

        $nextId = $main->getNextVersion($nameArg, $unitIdArg);
        if($nextId === false){
            $this->ajaxReturn(std_error('get new version from DB failed.'));
        }

        $ret = $main->addMain($nameArg, $contentArg, $unitIdArg, $nextId, $userArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Add main conf', I('server.HTTP_APPKEY'), $userArg, "name:$nameArg, unit_id:$unitIdArg");
            $this->ajaxReturn(std_return());
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function deprecated_post(){
        $idArg = I('id');
        $userArg = I('user');

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));


        $main = new Main();

        $ret = $main->setDeprecated($idArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'deprecated main conf', I('server.HTTP_APPKEY'), $userArg, "id:$idArg");
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

}
