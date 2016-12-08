<?php
/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 16/4/27
 * Time: 08:53
 */


namespace V1\Controller\Secure;

use Common\Dao\Secure\AppKey;
use Think\Controller\RestController;

class AppkeyController extends RestController{


    /*
     * 获取所有的appkey的列表,不支持分页,也不支持搜索
     *
     * 这是因为这个最多也就 10来条数据,没有必要做那些.
     * 如果有一天,appkey的数据多到必须要分页了,那这个系统肯定不再需要了.
     *
     */
    public function __construct()
    {
        parent::__construct();
        $ret = hubble_middle_layer();
        if (!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }



    public function list_get(){

        $nameArg   = I('name', '');

        $appkey = new AppKey();

        $ret = $appkey->getAppkeyList($nameArg);

        if($ret['code'] == HUBBLE_RET_NULL) 
            $this->ajaxReturn(std_return());
        
        if($ret['code'] == HUBBLE_DB_ERR)
            $this->ajaxReturn(std_error('db error'));

        $content['content'] = $ret['content'];
        $this->ajaxReturn(std_return($content));

    }

    /*
     * 获取appkey 的权限信息
     */
    public function detail_get(){

        $keyArg = I('key_id', '');

        if(empty($keyArg))
            $this->ajaxReturn(std_error('要查看的key为空'));


        $appkey = new Appkey();

        $ret = $appkey->getPrivilegeDetails($keyArg);
        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $content['content'] = $ret['content'];
            $this->ajaxReturn(std_return($content));
        } else{
            $this->ajaxReturn(std_error('db error'));
        }
    }

    /*
     * 添加一个appkey
     */
    public function add_post(){


        $nameArg   = I('name', '');
        $descArg   = I('desc', '');
        $userArg   = I('user', '');

        if(empty($nameArg))
            $this->ajaxReturn(std_error('名称为空'));

        if(empty($descArg))
            $this->ajaxReturn(std_error('描述为空'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user为空'));



        $appkey = new Appkey();

        if (!$appkey->existPrivilegeName($nameArg))
            $this->ajaxReturn(std_error("$nameArg 已经存在"));



        $ret = $appkey->addAppkey($nameArg, $descArg);
        if($ret == false)
            $this->ajaxReturn(std_error('db error'));

        hubble_oprlog('Secure', 'Add Appkey', I('server.HTTP_APPKEY'), $userArg, "name:$nameArg, desc:$descArg, ret:".json_encode($ret));
        $this->ajaxReturn(std_return(['key' => $ret]));

    }

    /*
     *  删除一个appkey
     */
    public function delete_delete(){

        $appkeyArg = I('server.HTTP_APPKEY');
        $keyArg = I('key', '');
        $userArg = I('user', '');

        if(empty($userArg))
            $this->ajaxReturn(std_error('user 为空'));

        if(empty($keyArg))
            $this->ajaxReturn(std_error('要删除key为空'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user 为空'));



        $appkey = new Appkey();

        $ret = $appkey->existPrivilegeAppkey($keyArg);
        if($ret)
            $this->ajaxReturn(std_error('no such appkey'));

        $ret = $appkey->delAppkey($keyArg);
        if($ret != true)
            $this->ajaxReturn(std_error('db error'));

        hubble_oprlog('Secure', 'Del Appkey', $appkeyArg, $userArg, $keyArg);
        $this->ajaxReturn(std_return());


    }


    // ------------  url 管理  -------------------------

    /*
     * 添加一个url
     */
    public function add_interface_post(){


        $addrArg   = I('addr', '');
        $descArg   = I('desc', '');
        $methodArg   = I('method', '');
        $appkeyArg = I('server.HTTP_APPKEY');
        $userArg   = I('user', '');

        if(empty($addrArg))
            $this->ajaxReturn(std_error('接口为空'));

        if(empty($descArg))
            $this->ajaxReturn(std_error('描述为空'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user为空'));



        $appkey = new Appkey();

        $ret = $appkey->addInterface($addrArg, $descArg, $methodArg);
        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Secure', 'Add Interface', $appkeyArg, $userArg, "addr:$addrArg : desc: $descArg, ret:".json_encode($ret));
        $this->ajaxReturn(std_return(['key' => $ret]));
    }

    /*
     * 删除一个url
     */
    public function delete_interface_delete(){

        $idArg   = I('id', '');
        $appkeyArg = I('server.HTTP_APPKEY');
        $userArg   = I('user', '');

        if(empty($idArg))
            $this->ajaxReturn(std_error('接口为空'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user为空'));


        $appkey = new Appkey();

        $ret = $appkey->delInterface($idArg);
        if($ret['code'] != 0)
            $this->ajaxReturn(std_error($ret['msg']));

        hubble_oprlog('Secure', 'Del Interface', $appkeyArg, $userArg, "$idArg : $ret");
        $this->ajaxReturn(std_return(['key' => $ret]));
    }

    /*
     * interface 的列表
     */
    public function list_interface_get(){

        $addrArg   = I('addr', false);
        $page       = I('page', 1);
        $limit      = I('limit', 20);

        // 参数检查
        if($page <= 0 || $limit <= 0)
            $this->ajaxReturn(std_error('page或limit错误'));


        // 设置过滤器
        $filter     = [];
        if(!empty($addrArg))
            $filter['addr'] = $addrArg;

        // init
        $content = [
            'page'      => (int)$page,
            'limit'     => (int)$limit,
            'content'   => array(),
            'count'     => 0,
            'total_page'=> 0,
        ];

        $appkey = new Appkey();

        $ret = $appkey->countInterface($filter);
        if($ret == 0) {
            $this->ajaxReturn(std_return($content));
        }
        $content['count'] = $ret;
        $content['total_page'] = ceil($ret / $limit);

        $ret = $appkey->getInterfaceList($filter, $page, $limit);
        if($ret['code'] == HUBBLE_RET_SUCCESS) {
            $content['content'] = $ret['content'];
            $this->ajaxReturn(std_return($content));

        } elseif($ret['code'] == HUBBLE_RET_NULL) {

            $this->ajaxReturn(std_return((object)[], "no such interface"));
        }else{

            $this->ajaxReturn(std_error('db error'));
        }
    }

    // ------------  权限管理  --------------------------

    /*
     * 给指定appkey 添加某个权限
     */
    public function add_privilege_post(){

        $addrIdArg   = I('addr_id', '');
        $keyIdArg    = I('key_id', '');
        $appkeyArg   = I('server.HTTP_APPKEY');
        $userArg     = I('user', '');

        if(empty($addrIdArg))
            $this->ajaxReturn(std_error('接口ID为空'));

        if(empty($keyIdArg))
            $this->ajaxReturn(std_error('key ID 为空'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user为空'));


        $appkey = new Appkey();

        $ret = $appkey->addPrivilege($keyIdArg, $addrIdArg);
        if($ret == false)
            $this->ajaxReturn(std_error('db error'));

        hubble_oprlog('Secure', 'Add privilege', $appkeyArg, $userArg, "addr_id:$addrIdArg, key_id:$keyIdArg, ret:".json_encode($ret));
        $this->ajaxReturn(std_return(['key' => $ret]));
    }

    /*
     * 删除
     */
    public function delete_privilege_delete(){

        $addrIdArg   = I('addr_id', '');
        $keyIdArg   = I('key_id', '');
        $appkeyArg = I('server.HTTP_APPKEY');
        $userArg   = I('user', '');

        if(empty($addrIdArg))
            $this->ajaxReturn(std_error('接口ID为空'));

        if(empty($keyIdArg))
            $this->ajaxReturn(std_error('key ID 为空'));

        if(empty($userArg))
            $this->ajaxReturn(std_error('user为空'));


        $appkey = new Appkey();

        $ret = $appkey->delPrivilege($keyIdArg, $addrIdArg);
        if($ret == false)
            $this->ajaxReturn(std_error('db error'));

        hubble_oprlog('Secure', 'Del privilege', $appkeyArg, $userArg, "addr_id:$addrIdArg, key_id:$keyIdArg, ret:".json_encode($ret));
        $this->ajaxReturn(std_return(['key' => $ret]));
    }

}