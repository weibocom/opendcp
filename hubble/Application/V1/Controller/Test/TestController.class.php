<?php
/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 2016/11/4
 * Time: 下午3:44
 */
namespace V1\Controller\Test;


use Common\Dao\Slb\Slb;
use Think\Controller\RestController;

class TestController extends RestController
{
    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }


    public function test_post(){

        hubble_log(HUBBLE_INFO, 'THIS IS A TEST');
        $this->ajaxReturn(I('server.'));

    }

    public function _empty(){ $this->response('404','', 404); }
}