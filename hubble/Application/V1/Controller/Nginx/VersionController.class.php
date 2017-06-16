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
 * Date: 16/9/7
 * Time: 上午11:09
 */

namespace V1\Controller\Nginx;

use Common\Dao\Adaptor\Channel;
use Common\Dao\Nginx\UnitModel;
use Common\Dao\Nginx\Version;
use Think\Controller\RestController;

class VersionController extends RestController {

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
        $page = I('page');
        $limit = I('limit');

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('limit or page out of range'));

        if($bidArg < 0)
            $this->ajaxReturn(std_error('biz_id is empty'));
        
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

        $version = new Version();

        $ret = $version->countVersion($filter);
        if($ret === false) {
            $this->ajaxReturn(std_error('db error'));
        }

        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }

        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $version->getVersionList($filter, $page, $limit, $likeArg);

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

        if($bidArg < 0)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $version = new Version();

        $ret = $version->getVersionDetail(['id' => $idArg, 'biz_id' => $bidArg]);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function generate_post(){
        $unitIdArg = I('unit_id');
        $nameArg = I('name');
        $userArg = I('user');
        $typeArg = I('type');
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if(empty($unitIdArg))
            $this->ajaxReturn(std_error('unit_id is empty'));

        if(empty($nameArg))
            $this->ajaxReturn(std_error('name is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if($bidArg < 0)
            $this->ajaxReturn(std_error('biz_id is empty'));
        
        $channel = new Channel();
        if(!$channel->isIllegal($typeArg))
            $this->ajaxReturn(std_error('type is not illegal'));


        $unit = new UnitModel();

        $ret = $unit->isExist($unitIdArg, $bidArg);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm unit_id exist: failed, db error'));

        if(is_null($ret))
            $this->ajaxReturn(std_error('unit_id dose not exist!'));


        $version = new Version();
        $ret = $version->generateVersion($unitIdArg, $nameArg, $typeArg, $userArg, $bidArg);

        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Nginx', 'generate version', I('server.HTTP_APPKEY'), $userArg, "name:$nameArg, unit_id: $unitIdArg, result: ".json_encode($ret));
        $this->ajaxReturn(std_return($ret['content']));

    }

    public function create_post(){
        $unitIdArg = I('unit_id');
        $nameArg = I('name');
        $userArg = I('user');
        $typeArg = I('type');
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        $filesArg = I('files', '','unsafe_raw');

        if(empty($unitIdArg))
            $this->ajaxReturn(std_error('unit_id is empty'));

        if(empty($nameArg))
            $this->ajaxReturn(std_error('name is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($filesArg))
            $this->ajaxReturn(std_error('files is empty'));

        if(empty($typeArg))
            $this->ajaxReturn(std_error('type is empty'));
        
        if($bidArg < 0)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $filesJson = json_decode($filesArg, true);
        if(is_null($filesJson))
            $this->ajaxReturn(std_error('files json foramt error'));

        $channel = new Channel();
        if(!$channel->isIllegal($typeArg))
            $this->ajaxReturn(std_error('type is not illegal'));


        $unit = new UnitModel();

        $ret = $unit->isExist($unitIdArg, $bidArg);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm unit_id exist: failed, db error'));

        if(is_null($ret))
            $this->ajaxReturn(std_error('unit_id dose not exist!'));


        $version = new Version();
        $ret = $version->createVersion($unitIdArg, $nameArg, $filesArg, $typeArg, $userArg, $bidArg);

        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Nginx', 'generate version', I('server.HTTP_APPKEY'), $userArg, "name:$nameArg, unit_id: $unitIdArg, result: ".json_encode($ret));
        $this->ajaxReturn(std_return($ret['content']));

    }

    public function deprecated_post(){
        $idArg = I('id');
        $userArg = I('user');
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if($bidArg < 0)
            $this->ajaxReturn(std_error('biz_id is empty'));

        $version = new Version();
        $ret = $version->isExist(['id' => $idArg, 'biz_id' => $bidArg]);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm version id exist: db error'));
        if(is_null($ret))
            $this->ajaxReturn(std_error("no such version id [$idArg]"));

        $ret = $version->setDeprecated($idArg, $bidArg);
        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Nginx', 'set version deprecated', I('server.HTTP_APPKEY'), $userArg, "id: $idArg");
        $this->ajaxReturn(std_return($ret['content']));
    }

    public function release_post(){
        $idArg = I('id');
        $userArg = I('user');
        $shellIdArg = I('shell_id');
        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($shellIdArg))
            $this->ajaxReturn(std_error('shell_id is empty'));

        if($bidArg < 0)
            $this->ajaxReturn(std_error('biz_id is empty'));
        
        $version = new Version();
        $ret = $version->isExist(['id' => $idArg, 'biz_id' => $bidArg]);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm version id exist: db error'));
        if(is_null($ret))
            $this->ajaxReturn(std_error("no such version id [$idArg]"));

        $ret = $version->releaseVersion($idArg, $shellIdArg, $userArg, $bidArg);
        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Nginx', 'release version', I('server.HTTP_APPKEY'), $userArg, "id: $idArg, shell_id: $shellIdArg");
        $this->ajaxReturn(std_return($ret['content']));
    }
}
