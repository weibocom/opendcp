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
 * Date: 16/9/21
 * Time: 下午2:37
 */

namespace V1\Controller\Adaptor;

use Common\Dao\Adaptor\AlterationHistory;
use Think\Controller\RestController;

class HistoryController extends RestController {

    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }


    public function list_get(){

        $nameArg = I('task_name');
        $idArg = I('id');
        $taskIdArg = I('task_id');
        $likeArg = I('like', true);

        $page = I('page');
        $limit = I('limit');



        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('limit or page out of range'));

        // 设置过滤器
        $filter     = [];
        if(!empty($nameArg))
            $filter['name'] = $nameArg;

        if(!empty($idArg))
            $filter['id'] = $idArg;

        if(!empty($nameArg))
            $filter['task_id'] = $taskIdArg;


        // init
        $content = [
            'page'      => (int)$page,
            'limit'     => (int)$limit,
            'content'   => array(),
            'count'     => 0,
            'total_page'=> 0,
        ];

        $history = new AlterationHistory();

        $ret = $history->countHistory($filter);
        if($ret === false) {
            $this->ajaxReturn(std_error('db error'));
        }

        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }

        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $history->getHistoryList($filter, $page, $limit, $likeArg);

        if($ret['state'] == HUBBLE_RET_SUCCESS) {
            $content['content'] = $ret['content'];
            $this->ajaxReturn(std_return($content));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }
}
