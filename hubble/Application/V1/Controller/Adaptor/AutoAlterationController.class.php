<?php
/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 16/9/2
 * Time: 下午1:51
 */

namespace V1\Controller\Adaptor;

use Common\Dao\Adaptor\Adaptor;
use Common\Dao\Adaptor\AlterationHistory;
use Common\Dao\Adaptor\AlterationType;
use Common\Dao\Adaptor\Channel;
use Think\Controller\RestController;

class AutoAlterationController extends RestController {

    private $input;

    public function __construct()
    {
        parent::__construct();

        $this->input = hubble_parse_param();
        if(! IS_GET && empty($this->input)){
            $this->ajaxReturn(std_error('parameter is empty'));
        }

        $ret = hubble_middle_layer();
        if(!$ret[0])
            $this->ajaxReturn(std_error($ret[1]));
    }

    public function _empty(){ $this->response('404','', 404); }


    public function add_post(){

        $params = ['type_id', 'ips', 'user', ];

        foreach ($params as $p){
            if(!isset($this->input[$p]) || empty($this->input[$p]))
                $this->ajaxReturn(std_error("parameter [$p] is absent or empty, please check and try again."));
        }
        $ipStr = $this->input['ips'];
        $this->input['ips'] = explode(',', $this->input['ips']);

        $adaptor = new Adaptor();
        $ret = $adaptor->doAddNode($this->input['type_id'], $this->input, $this->input['user']);
        if($ret['code'] == 0){
            hubble_log(HUBBLE_INFO, 'auto alteration add success'.json_encode($ret['content']));
            hubble_oprlog('Adaptor', 'auto alteration add success',
                I('server.HTTP_APPKEY'), $this->input['user'], "type_id:{$this->input['type_id']}, ips:$ipStr");
            $this->ajaxReturn(std_return($ret['content']));
        }
        else{
            hubble_log(HUBBLE_WARN, $ret['msg']);
            $this->ajaxReturn(std_error($ret['msg']));
        }
    }

    public function remove_post(){

        $params = ['type_id', 'ips', 'user', ];

        foreach ($params as $p){
            if(!isset($this->input[$p]) || empty($this->input[$p]))
                $this->ajaxReturn(std_error("parameter [$p] is absent or empty, please check and try again."));
        }
        $ipStr = $this->input['ips'];
        $this->input['ips'] = explode(',', $this->input['ips']);

        $adaptor = new Adaptor();
        $ret = $adaptor->doDelNode($this->input['type_id'], $this->input, $this->input['user']);
        if($ret['code'] == 0){
            hubble_log(HUBBLE_INFO, 'auto alteration del success'.json_encode($ret['content']));
            hubble_oprlog('Adaptor', 'auto alteration del',
                I('server.HTTP_APPKEY'), $this->input['user'], "type_id:{$this->input['type_id']}, ips:$ipStr");
            $this->ajaxReturn(std_return($ret['content']));
        }
        else{
            hubble_log(HUBBLE_WARN, $ret['msg']);
            $this->ajaxReturn(std_error($ret['msg']));
        }

    }

    public function check_state_get(){

        $gid = I('server.HTTP_X_CORRELATION_ID');
        $rid = I('release_id');

        if(empty($gid) && empty($rid))
            $this->ajaxReturn(std_error('correlation-id and release_id are both empty'));

        $record =  new AlterationHistory();

        if(!empty($rid)){// if there is release_id, use it in first
            $ret = $record->exist($rid);
            $gid = $ret['content']['global_id'];
        }
        else
            $ret = $record->existGid($gid);


        if($ret['code'] == 1){
            $this->ajaxReturn(std_error('check task state: '.$ret['msg']));
        }

        $ret = $ret['content'];
        if($ret['type'] == 'sync')
            $this->ajaxReturn(std_return(['task_id'=>$ret['task_id']]));

        switch(strtoupper($ret['channel'])){
            case 'ANSIBLE':

                $channel = new Channel();
                $result = $channel->ansibleCheck($ret['task_name']);

                if($result['code'] != 0){
                    $this->ajaxReturn(std_error("http: ".$result['error']));
                }

                $data = json_decode($result['data'],true);
                if(empty($data))
                    $this->ajaxReturn(std_error('wrong json format'));

                if($data['code'] != 0)
                    $this->ajaxReturn(std_error("ansible: ".$data['message']));

                $content = [];
                $content['state'] = $data['content']['task']['status'];

                if(!empty($data['content']['nodes'])){
                    foreach($data['content']['nodes'] as $v){
                        $content['detail'][] = [
                            'ip'=> $v['ip'],
                            'state'=> $v['status']];
                    }
                }
                $content['X-CORRELATION-ID'] = $gid;
                $this->ajaxReturn(std_return($content));
                break;

            default:
                $this->ajaxReturn(std_error('no such channel to deal with.'));
        }
    }


    public function type_param_get(){
        $typeArg = I('type');

        if(empty($typeArg))
            $this->ajaxReturn(std_error('type is empty'));

        $alteration = new AlterationType();

        $ret = $alteration->getTypeColumns($typeArg);

        if($ret['code'] == 1){
            $this->ajaxReturn(std_error($ret['msg']));
        } else{
            $this->ajaxReturn(std_return($ret['content']));
        }
    }

}