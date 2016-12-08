<?php
/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 16/9/1
 * Time: 下午6:11
 */

namespace V1\Controller\Nginx;

use Common\Dao\Nginx\Shell;
use Think\Controller\RestController;

class ShellController extends RestController{

    function __construct(){

        parent::__construct();
        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }


    public function list_get(){

        $nameArg = I('name');
        $likeArg = I('like', true);

        $page = I('page', 1);
        $limit = I('limit', 20);

        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('limit or page out of range'));

        // 设置过滤器
        $filter     = [];
        if(!empty($nameArg))
            $filter['name'] = $nameArg;

        // init
        $content = [
            'page'      => (int)$page,
            'limit'     => (int)$limit,
            'content'   => array(),
            'count'     => 0,
            'total_page'=> 0,
        ];

        $shell = new Shell();

        $ret = $shell->countShell($filter);
        if($ret === false) {
            $this->ajaxReturn(std_error('db error'));
        }

        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }

        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $shell->getShellList($filter, $page, $limit, $likeArg);

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

        $shell = new Shell();

        $ret = $shell->getShellDetail($idArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function add_post(){
        $nameArg = I('name');
        $contentArg = I('content', '', 'unsafe_raw');
        $descArg = I('desc');
        $userArg = I('user');

        if(empty($nameArg))
            $this->ajaxReturn(std_error('name is empty'));

        if(empty($contentArg))
            $this->ajaxReturn(std_error('content is empty'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));


        $shell = new Shell();

        $ret = $shell->addShell($nameArg, $descArg, $contentArg, $userArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Add shell', I('server.HTTP_APPKEY'), $userArg, "name:$nameArg, desc: $descArg");
            $this->ajaxReturn(std_return());
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function modify_put(){
        $idArg = I('id');
        $nameArg = I('name');
        $descArg = I('desc');
        $contentArg = I('content', '', 'unsafe_raw');
        $userArg = I('user', '');

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));

        if(empty($contentArg))
            $this->ajaxReturn(std_error('content is empty'));


        $shell = new Shell();


        $ret = $shell->modifyShell($idArg, $nameArg, $descArg, $contentArg, $userArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Add upstream', I('server.HTTP_APPKEY'), $userArg, "id:$idArg, content: $contentArg");
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function delete_delete(){

        $idArg = I('id');
        $userArg = I('user', '');

        if(empty($userArg))
            $this->ajaxReturn(std_error('user is empty'));

        if(empty($idArg))
            $this->ajaxReturn(std_error('id is empty'));


        $shell = new Shell();

        $ret = $shell->deleteShell($idArg);

        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            hubble_oprlog('Nginx', 'Del Shell', I('server.HTTP_APPKEY'), $userArg, "id:$idArg");
            $this->ajaxReturn(std_return($ret['content']));
        } else{
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }
}