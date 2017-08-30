<?php
/**
 * Created by IntelliJ IDEA.
 * User: Gavin
 * Date: 2017/8/7
 * Time: 15:57
 */

namespace V1\Controller\Nginx;
use Common\Dao\Adaptor\AlterationHistory;
use Common\Dao\Nginx\Upstream;
use Think\Controller\RestController;

class TimingController extends RestController
{
    private $timinginfo,$alterationHistory,$altreationType,$upstream;
    public function __construct()
    {
        ignore_user_abort(true);//忽略用户的断开
        set_time_limit(0);//设置脚本最大执行时间
        parent::__construct();
        $this->timinginfo = M("TimingInfo");
        $this->alterationHistory = M("AlterationHistory");
        $this->altreationType = M("AlterationType");
        $this->upstream = new Upstream();
    }
    //定时任务
    public function beginTimingReload()
    {
        ignore_user_abort(true);//忽略用户的断开
        set_time_limit(0);//设置脚本最大执行时间
        $correlation = I('correlation');//每个ip对应的correlion_id,数据表中为global_id
        $sid = I('id');//服务发现类型id
        $user = I('user');//操作用户
        $data['global_id']=$correlation;
        $data['sid']=$sid;
        $data['opr_user'] = $user;
        $this->alterationHistory->add($data);
        //创建标志文件
        $entry_dir = C('HUBBLE_ROOT_DIR');
        $entry_file = C('HUBBLE_ROOT_DIR') . "timing";
        mkdir($entry_dir, 0755, true);
        //标志文件存在，说明循环计时任务已经开始
        if (!file_exists($entry_file)) {
            fclose(fopen($entry_file, "a"));
            //找出task_name为空的不重复的sid
            $map['task_name']  = array('eq','');
            $judge=1;
            while($judge>0){

                //1.查询history数据表是否存在未发布数据
                $resone = $this->alterationHistory->distinct(true)->where($map)->field('sid,opr_user')->select();
                $res=$resone;
                if($res==NULL){
                    //2.查询全部服务发现类型
                    $restwo = $this->altreationType->distinct(true)->field('id,content,opr_user')->select();
                    $res=$restwo;
                }
                for ($i=0;$i<count($res);$i++){
                    //统一化;避免二次查询
                    if(isset($res[$i]['sid'])){
                        $res[$i]['id']=$res[$i]['sid'];
                        //查询找出服务发现类型的content信息
                        $cont = $this->altreationType->field('content')
                            ->where(['id' => $res[$i]['id']])
                            ->find();
                        $content = json_decode($cont['content'], true);
                    }else{
                        $content = json_decode($res[$i]['content'], true);
                    }

                    //通过api获取当前服务发现类型的ip列表
                    $url = "http://orion:8080/pool/".$res[$i]['id']."/list_register";
                    $ch = curl_init();
                    curl_setopt($ch, CURLOPT_URL, $url);
                    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
                    curl_setopt($ch, CURLOPT_HEADER,0);
                    $json = curl_exec($ch);
                    curl_close($ch);
                    $ips=json_decode($json, true);
                    $content['ips']=$ips['data'];

                    //变更数据库中的upstream文件
                    $upstream = new Upstream();
                    //数据写入数据库
                    $ret = $upstream->addNode(
                        $content['name'], $content['group_id'], $content['ips'],
                        $content['port'], $content['weight']);
                    //print_r($ret);

                    if($ret['code'] != 0) return $ret;

                    // -------- 对 consul 的处理
                    if($ret['content']['is_consul']){
                        $return['content'] = [
                            'type' => 'sync',
                            'task_id' => 0,
                        ];
                        return $return;
                    }

                    //下发时，标志是服务发现类型的id
                    $task = $this->upstream->callTunnel($content['script_id'], $content['name'], $res[$i]['opr_user'], true, $content['group_id'], '', $res[$i]['id']);

                    //依据history表则记录变更表
                    if(isset($res[$i]['sid'])){
                        //记录变更表；根据服务发现类型的id变更所有与之相关的ip的task_id和task_name
                        $task = $task['content'];
                        $history = new AlterationHistory();
                        //将返回的task_id 和task_name保存至对应的sid列中的对应字段
                        $history->addTaskRecord('async', $task['ansible_id'], $task['ansible_name'], 'ansible', $res[$i]['id']);
                    }
                }
                sleep(50);//暂停50秒
                $judge=1;
            }
            unlink($entry_file);
        }
    }
}
