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
    private $timinginfo,$upstream;
    public function __construct()
    {
        ignore_user_abort(true);//忽略用户的断开
        set_time_limit(0);//设置脚本最大执行时间
        parent::__construct();
        $this->timinginfo = M("TimingInfo");
        $this->upstream = new Upstream();
    }
    //测试函数
    public function test(){
        echo "cli - php test";
        echo I('aa');
    }
    //定时任务
    public function beginTimingReload()
    {
        ignore_user_abort(true);//忽略用户的断开
        set_time_limit(0);//设置脚本最大执行时间
        //接受值，并且写入数据库
        $id = I('id');
        $sid = I('sid');
        $name = I('name');
        $gid = I('gid');
        $user = I('user');
        $correlation = I('correlation');
        $data['sdid'] = intval($id);
        $data['script_id'] = intval($sid);
        $data['name'] = $name;
        $data['user'] = $user;
        $data['group_id'] = intval($gid);
        $data['correlation_id'] = $correlation;
        if (!empty($data['script_id']) && !empty($data['name']) && !empty($data['user']) && !empty($data['group_id']) && !empty($data['correlation_id'])){

            $this->timinginfo->add($data);
            $entry_dir = C('HUBBLE_ROOT_DIR');
            $entry_file = C('HUBBLE_ROOT_DIR') . "test";
            mkdir($entry_dir, 0755, true);
            //标志文件存在，说明循环计时任务已经开始
            if (!file_exists($entry_file)) {
                fclose(fopen($entry_file, "a"));
                $res = $this->timinginfo->find();
                //数据库数据不为空
                while (is_array($res) && count($res) > 0) {
                    $result = $this->timinginfo->find();
                    $dele['correlation_id']=$result['correlation_id'];
                    //删除correlation_id相同数据
                    $this->timinginfo->where($dele)->delete();
                    //下发.
                    $task = $this->upstream->callTunnel($result['script_id'], $result['name'], $result['user'], true, $result['group_id'], '', $result['correlation_id']);
                    //记录变更表
                    $task = $task['content'];
                    $history = new AlterationHistory();
                    $history->addRecord('async', $task['ansible_id'], $task['ansible_name'], 'ansible', $result['user'], $result['correlation_id']);
                    sleep(3);//倒计时
                    $res = $this->timinginfo->find();
                }
                unlink($entry_file);
            }
        }
    }
}
