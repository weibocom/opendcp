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
 * Date: 16/4/19
 * Time: 10:15
 */

namespace V1\Controller\Secure;

use Common\Dao\Common\LogDb;
use Common\Dao\Secure\Oprlog;
use Think\Controller\RestController;


class OprlogController extends RestController{

    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));


    }

    public function _empty(){ $this->response('404','', 404); }

    public function list_get(){

        // 参数获取
        $fileArg    = I('operation', '');
        $page       = I('page', 1);
        $limit      = I('limit', 20);

        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('wrong page or limit'));


        // 设置过滤器
        $filter     = [];
        if(!empty($fileArg))
            $filter['operation'] = $fileArg;

        // init
        $content = [
            'page'      => (int)$page,
            'limit'     => (int)$limit,
            'content'   => array(),
            'count'     => 0,
            'total_page'=> 0,
        ];

        $oprlog = new Oprlog();

        $ret = $oprlog->countOprlog($filter);
        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }
        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $oprlog->getOprlogList($filter, $limit, $page);
        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            $content['content'] = $ret['content'];
            $this->ajaxReturn(std_return($content));
        } else{
            $this->ajaxReturn(std_error('db error'));
        }
    }

    public function log_get(){

        $gidArg = I('correlation_id', '');

        if(empty($gidArg))
            $this->ajaxReturn(std_error('correlation_id is empty'));


        $logDb = new LogDb();
        $ret =  $logDb->getAllLog($gidArg);

        if($ret['code'] != 0)
            $this->ajaxReturn(std_error('get log from db:'.$ret['msg']));

        $this->ajaxReturn(std_return($ret['content']));
    }

    public function iplog_get(){
        $gidArg = I('correlation_id', '');
        $ipArg = I('ip');

        if(empty($gidArg))
            $this->ajaxReturn(std_error('correlation_id is empty'));

        if(empty($ipArg))
            $this->ajaxReturn(std_error('ip is empty'));

        $logDb = new LogDb();
        $ret =  $logDb->getOctanLog($gidArg, $ipArg);

        $this->ajaxReturn(std_return(implode('\n', $ret)));
    }
}

