<?php

include_once('include/config.inc.php');
include_once('include/function.php');
include_once('include/node_init.php');
include_once('include/cloud.php');
include_once('include/layout.php');

$mycloud = new cloud();
$mylayout = new layout();

$arr_type = array(
	1=>array(
		'pool_id'=>4,
		'template_id'=>8,
	),
	2=>array(
		'pool_id'=>3,
		'template_id'=>7,
	),
);

while(true) {
$arr = node_init::getNodeInitList();
foreach($arr['data'] as $oneinit){
	if($oneinit['status']!=0){
		continue;
	}

	node_init::modifyOneNodeInit($oneinit['id'], array('status'=>1));
	
	doaddlog($oneinit['id'], '初始化开始', 0, '');
	
	$ip = $oneinit['ip'];
	$password = $oneinit['password'];
	
	$arrJson = array();
        $arrJson['InstanceList'][] = array(
		'publicip'=>$ip,
		'privateip'=>$ip,
		'password'=>$password,
	);
	dolog('add machine: '.print_r($arrJson,true));
	$ret = $mycloud->get('root', 'instance/phydev', 'POST', $arrJson);
	dolog('add machine ret: '.$ret);
	doaddlog($oneinit['id'], '机器录入成功', 1, '');

	doaddlog($oneinit['id'], '等待机器初始化', 0, '');
	$initok = false;
	while(true){
		sleep(5);
		$arrJson = array(
			'page' => 1,
			'pagesize' => 1000,
		);
		$machinelist = $mycloud->get('root', 'instance/list', 'GET', $arrJson);
		$arr_ml = @json_decode($machinelist, true);
		foreach($arr_ml['content'] as $onemachine){
			if($onemachine['PrivateIpAddress']==$ip){
				dolog('initializing '.print_r($onemachine, true));
				if($onemachine['Status']==1){
					$initok = true;
					break;
				}
			}
		}
		if($initok) {
			dolog('init ok');
			break;
		}
	}
	doaddlog($oneinit['id'], '机器初始化成功', 1, '');

	dolog('add machine to pool');
	$pool_id = $arr_type[$oneinit['type']]['pool_id'];
        $retArr=$mylayout->get(
		'root', 
		'pool/'.$pool_id,
		'add_nodes', 
		array(
			'nodes'=>array($ip),
		)
	);
	dolog('add nodes ret: '.$retArr);
	doaddlog($oneinit['id'], '将机器加入服务池', 1, '');
	
	dolog('sleep 5');
	sleep(5);
	
	dolog('go do task');
        $retArr=$mylayout->get(
		'root',
		'task',
		'create', 
		array(
			'template_id'=>$arr_type[$oneinit['type']]['template_id'],
			'task_name'=>'init_compute_'.$ip,
			'max_num'=>1,
			'max_ratio'=>30,
			'nodes'=>array(
				array('ip'=>$ip,),
			),
			'opr_user'=>'root',
		)
	);
	dolog('go do task ret: '.$retArr);
	doaddlog($oneinit['id'], '创建初始化任务完成', 1, '');
	break;
}
sleep(5);
}

function dolog($msg){
        global $logid;
        printf('[logid:'.$logid.']['.date('Y-m-d H:i:s').'] '.$msg."\n");
}

function doaddlog($task_id, $title, $status, $text){
	exec('curl -v "http://10.39.59.73:8888/api/for_openstack/machine.php?action=addlog&task_id='.$task_id.'&title='.$title.'&status='.$status.'&text='.$text.'"', $ret);
	dolog('curl ret: '.print_r($ret,true));
}


?>
