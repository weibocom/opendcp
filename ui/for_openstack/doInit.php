<?php

include_once('include/config.inc.php');
include_once('include/function.php');
include_once('include/node_init.php');
include_once('include/cloud.php');
include_once('include/layout.php');
include_once('include/keydata.php');

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
    3=>array(
        'pool_id'=>5,
        'template_id'=>10,
    ),
);


$arr = node_init::getNodeInitList();
foreach($arr['data'] as $oneinit){
	if($oneinit['status']!=0){
		continue;
	}

	node_init::modifyOneNodeInit($oneinit['id'], array('status'=>1));
	
	doaddlog($oneinit['id'], '初始化开始', 0, '');
	
	$ip = $oneinit['ip'];
	$password = $oneinit['password'];

	//del
	delip($ip, $mylayout, $mycloud);
	if($oneinit['type']==2){
		//del original controllerip
		$cip = keydata::getContentByKey('controller_ip');
		delip($cip, $mylayout, $mycloud);
	}

	sleep(3);

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
				if($onemachine['Status']==7 || $onemachine['Status']==4){
					dolog('init machine failed');
					doaddlogfinal($oneinit['id'], '机器初始化失败', 2, '', 11);
					exit;
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


function dolog($msg){
        global $logid;
        printf('[logid:'.$logid.']['.date('Y-m-d H:i:s').'] '.$msg."\n");
}

//打印日志到后台
function doaddlog($task_id, $title, $status, $text){
	exec('curl -v "http://host_ip:8888/api/for_openstack/machine.php?action=addlog&task_id='.$task_id.'&title='.$title.'&status='.$status.'&text='.$text.'"', $ret);
	dolog('curl ret: '.print_r($ret,true));
}

//打印日志到前端
function doaddlogfinal($task_id, $title, $status, $text, $final = 10){
	exec('curl -v "http://host_ip:8888/api/for_openstack/machine.php?action=addlog&task_id='.$task_id.'&title='.$title.'&status='.$status.'&text='.$text.'&final='.$final.'"', $ret);
	dolog('curl ret: '.print_r($ret,true));
}

function delip($ip, $mylayout, $mycloud){
	dolog('del ip: '.$ip);
	$ip = trim($ip);
	if(empty($ip)) return false;
        $getipret = $mylayout->get('root', 'pool/search_by_ip', $ip);
        $getipret = @json_decode($getipret, true);
        dolog('search ip ret: '.print_r($getipret,true));
        if(!empty($getipret['data'][$ip]) && $getipret['data'][$ip]!=-1){
                $arrJson = array(
                        'page' => 1,
                        'page_size' => 100,
                        'pool_id' => $getipret['data'][$ip],
                );
                $poollist = $mylayout->get('root', 'pool/'.$getipret['data'][$ip], 'list_nodes', $arrJson);
                $poollist = @json_decode($poollist, true);
                dolog('getpoolnode ret:'.print_r($poollist, true));
                foreach($poollist['data'] as $onenode){
                        if($onenode['ip']==$ip){
                                $arrJson = array(
                                        'nodes'=>array(
                                                (int)$onenode['id'],
                                        ),
                                );
                                dolog('remove nodes :'.print_r($arrJson, true));
                                $removeret = $mylayout->get('root', 'pool/'.$getipret['data'][$ip], 'remove_nodes', $arrJson);
                                dolog('remove nodes ret:'.print_r($removeret, true));
                        }
                }
        }



        $mlist = $mycloud->get('root', 'instance/list', 'GET');
        $mlist = @json_decode($mlist, true);
        dolog('machine list : '.print_r($mlist ,true));
        foreach($mlist['content'] as $onemachine){
                if($onemachine['PrivateIpAddress']==$ip){
                        dolog('delete machine: '.print_r($onemachine, true));
                        $delret = $mycloud->get('root', 'instance/'.$onemachine['InstanceId'], 'DELETE');
                        dolog('delete machine ret: '.print_r($delret, true));
                }
        }
        return true;
}


?>
