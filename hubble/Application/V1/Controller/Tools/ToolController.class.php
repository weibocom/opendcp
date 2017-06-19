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
 * User: yabo
 * Date: 16/9/1
 * Time: 下午1:38
 */

namespace V1\Controller\Tools;

use Common\Dao\Adaptor\AlterationType;
use Common\Dao\Nginx\GroupModel;
use Common\Dao\Nginx\Main;
use Common\Dao\Nginx\Shell;
use Common\Dao\Nginx\UnitModel;
use Common\Dao\Nginx\Upstream;
use Think\Controller\RestController;

class ToolController extends RestController{

    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function nginx_init_post(){

        $bidArg = I('server.HTTP_X_BIZ_ID',0);

        if($bidArg < 1 )
            $this->ajaxReturn(std_error('biz_id is empty'));

        //创建默认分组
        $group = new GroupModel();

        $ret = $group->getDetail(['biz_id' => $bidArg]);
        if($ret['code'] == HUBBLE_DB_ERR)
            $this->ajaxReturn(std_error($ret['msg']));

        if($ret['code'] == 0)
            $this->ajaxReturn(std_error("the group of $bidArg exists"));

        $ret = $group->addGroup('default_group','system',$bidArg);
        if($ret['code'] == 1)
            $this->ajaxReturn(std_error($ret['msg']));

        $gid = $ret['content']['gid'];

        //创建默认单元
        $unit = new UnitModel();

        $ret = $unit->getDetail(['biz_id' => $bidArg]);
        if($ret['code'] == HUBBLE_DB_ERR)
            $this->ajaxReturn(std_error($ret['msg']));

        if($ret['code'] == HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error("the unit of $bidArg exists"));

        $ret =$unit->addUnit('default_unit',$gid,'system',$bidArg);
        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error($ret['msg']));

        $uid = $ret['content']['uid'];


        //创建nginx 主配置文件
        $main = new Main();

        $ret = $main->getMainList(['biz_id' => $bidArg],1, 20);
        
        if($ret['code'] == HUBBLE_DB_ERR)
            $this->ajaxReturn(std_error('get main_conf error: db error'));

        if($ret['code'] == HUBBLE_RET_SUCCESS )
            $this->ajaxReturn(std_error("main_conf of $bidArg exist"));

        $ret = $main->getMainDetail(['id' => 1 ,'biz_id' => 0]);
        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' get init main_conf '.$ret['msg']));

        $ret = $main->addMain('nginx.conf', $ret['content']['content'], $uid,1, 'system', $bidArg);
        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' set init main_conf '.$ret['msg']));

        //创建upstream文件
        $upstream = new Upstream();

        $ret = $upstream->getUpstreamList(['biz_id' => $bidArg], 1, 20);

        if($ret['code'] == HUBBLE_DB_ERR)
            $this->ajaxReturn(std_error('get upstream_conf error: db error'));

        if($ret['code'] == HUBBLE_RET_SUCCESS )
            $this->ajaxReturn(std_error("upstream_conf of $bidArg exist"));

        $ret = $upstream->getUpstreamDetail(['id' => 1 ,'biz_id' => 0]);
        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' get init upstream_conf '.$ret['msg']));
        
        $ret = $upstream->addUpstream('default.upstream', $ret['content']['content'], $gid, 0, 'system', $bidArg  );
        
        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' set init upstream_conf '.$ret['msg']));

        //创建shell
        $shell = new Shell();

        $ret = $shell->getShellList(['biz_id' => $bidArg], 1 ,20);
        
        if($ret['code'] == HUBBLE_DB_ERR)
            $this->ajaxReturn(std_error('get shell_nginx error: db error'));

        if($ret['code'] == HUBBLE_RET_SUCCESS )
            $this->ajaxReturn(std_error("shell_nginx of $bidArg exist"));

        //主配文件脚本
        $ret = $shell->getShellDetail(['id' => 1 ,'biz_id' => 0]);
        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' get init shell_main_conf '.$ret['msg']));

        $ret = $shell->addShell('updateMainConf.sh', '更新主配置脚本', $ret['content']['content'], 'system', $bidArg);

        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' set init shell_main_conf '.$ret['msg']));

        //upstream配置脚本
        $ret = $shell->getShellDetail(['id' => 2 ,'biz_id' => 0]);
        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' get init shell_upstream_conf '.$ret['msg']));

        $ret = $shell->addShell('updateUpstreamConf.sh', '更新upstream配置脚本', $ret['content']['content'], 'system', $bidArg);

        if($ret['code'] != HUBBLE_RET_SUCCESS)
            $this->ajaxReturn(std_error(' set init shell_upstream_conf '.$ret['msg']));

        $sid = $ret['content'];
        
        //服务注册
        $AlterationType = new AlterationType();

        $ret = $AlterationType->exist(['biz_id' => $bidArg]);
        if($ret['code'] == 1)
            $this->ajaxReturn(std_error("get init Alteration_Type of $bidArg exist"));

        if($ret['code'] == 2)
            $this->ajaxReturn(std_error("get init Alteration_Type error: db error"));

        $ret = $AlterationType->exist(['id' => 1, 'biz_id' => 0]) ;
        if($ret['code'] != 1)
            $this->ajaxReturn(std_error(' get init Alteration_Type '.$ret['msg']));

        $content = json_decode($ret['content'],true);
        $content['group_id'] = $gid;
        $content['script_id'] = $sid;

        $data = $ret['content'];
        $ret = $AlterationType->add($data['type'],$data['name'],json_encode($content),'system', $bidArg);

        if($ret['code'] == 1)
            $this->ajaxReturn(std_error($ret['msg']));
        
        $this->ajaxReturn(std_return());
    }

}   