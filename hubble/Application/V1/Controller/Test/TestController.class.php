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
