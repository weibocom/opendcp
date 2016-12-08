<?php
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

        $page = I('page');
        $limit = I('limit');

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

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        $version = new Version();

        $ret = $version->getVersionDetail($idArg);

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

        if(empty($unitIdArg))
            $this->ajaxReturn(std_error('unit_id is empty'));

        if(empty($nameArg))
            $this->ajaxReturn(std_error('name is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        $channel = new Channel();
        if(!$channel->isIllegal($typeArg))
            $this->ajaxReturn(std_error('type is not illegal'));


        $unit = new UnitModel();

        $ret = $unit->isExist($unitIdArg);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm unit_id exist: failed, db error'));

        if(is_null($ret))
            $this->ajaxReturn(std_error('unit_id dose not exist!'));


        $version = new Version();
        $ret = $version->generateVersion($unitIdArg, $nameArg, $typeArg, $userArg);

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

        $filesJson = json_decode($filesArg, true);
        if(is_null($filesJson))
            $this->ajaxReturn(std_error('files json foramt error'));

        $channel = new Channel();
        if(!$channel->isIllegal($typeArg))
            $this->ajaxReturn(std_error('type is not illegal'));


        $unit = new UnitModel();

        $ret = $unit->isExist($unitIdArg);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm unit_id exist: failed, db error'));

        if(is_null($ret))
            $this->ajaxReturn(std_error('unit_id dose not exist!'));


        $version = new Version();
        $ret = $version->createVersion($unitIdArg, $nameArg, $filesArg, $typeArg, $userArg);

        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Nginx', 'generate version', I('server.HTTP_APPKEY'), $userArg, "name:$nameArg, unit_id: $unitIdArg, result: ".json_encode($ret));
        $this->ajaxReturn(std_return($ret['content']));

    }

    public function deprecated_post(){
        $idArg = I('id');
        $userArg = I('user');

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));


        $version = new Version();
        $ret = $version->isExist($idArg);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm version id exist: db error'));
        if(is_null($ret))
            $this->ajaxReturn(std_error("no such version id [$idArg]"));

        $ret = $version->setDeprecated($idArg);
        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Nginx', 'set version deprecated', I('server.HTTP_APPKEY'), $userArg, "id: $idArg");
        $this->ajaxReturn(std_return($ret['content']));
    }

    public function release_post(){
        $idArg = I('id');
        $userArg = I('user');
        $shellIdArg = I('shell_id');

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($shellIdArg))
            $this->ajaxReturn(std_error('shell_id is empty'));

        $version = new Version();
        $ret = $version->isExist($idArg);
        if($ret === false)
            $this->ajaxReturn(std_error('confirm version id exist: db error'));
        if(is_null($ret))
            $this->ajaxReturn(std_error("no such version id [$idArg]"));

        $ret = $version->releaseVersion($idArg, $shellIdArg, $userArg);
        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Nginx', 'release version', I('server.HTTP_APPKEY'), $userArg, "id: $idArg, shell_id: $shellIdArg");
        $this->ajaxReturn(std_return($ret['content']));
    }
}