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
        $bidArg = I('server.HTTP_X_BIZ_ID',0);
        $page = I('page', 1);
        $limit = I('limit', 20);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('limit or page out of range'));
        
        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        // 设置过滤器
        $filter     = ['biz_id' => $bidArg];
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
        $bidArg = I('server.HTTP_X_BIZ_ID',0);
        $page = I('page', 1);
        $limit = I('limit', 20);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('limit or page out of range'));

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        // 设置过滤器
        $filter     = ['biz_id' => $bidArg];
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
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if($bidArg < 1)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $main = new Main();

        $ret = $main->getMainDetail(['id' => $idArg, 'biz_id' => $bidArg]);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function add_post(){
        $nameArg = I('name');
        $bidArg = I('server.HTTP_X_BIZ_ID',0);
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

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        $main = new Main();

        $nextId = $main->getNextVersion($nameArg, $unitIdArg, $bidArg);
        if($nextId === false){
            $this->ajaxReturn(std_error('get new version from DB failed.'));
        }

        $ret = $main->addMain($nameArg, $contentArg, $unitIdArg, $nextId, $userArg, $bidArg);

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
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if($bidArg < 1){
            $this->ajaxReturn(std_error('biz_id is empty'));
        }
        
        $main = new Main();

        $ret = $main->setDeprecated($idArg, $bidArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'deprecated main conf', I('server.HTTP_APPKEY'), $userArg, "id:$idArg");
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

}
