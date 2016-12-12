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
namespace Version\Controller;
use Think\Controller;
class IndexController extends Controller {
    public function index(){
        echo "Welcome to Hubble";
    }

    public function test(){
        $ret = file_get_contents('php://input');


        echo json_encode($ret);
        echo json_encode(I('server.'));
    }
}
