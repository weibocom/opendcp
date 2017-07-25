<?php

include_once('include/config.inc.php');
include_once('include/function.php');
include_once('include/cloud.php');
include_once('include/layout.php');
include_once('include/keydata.php');

include_once('include/computepower.php');
include_once('include/openstack.php');

while(true){

        openstack::needOpenstackLogin();
        $arr_hyper = openstack::getHypervisorList(array(), 0, 1000);
        $arr_ret = array(
                'vcpus'=>0,
                'vcpus_used'=>0,
                'memory_mb'=>0,
                'memory_mb_used'=>0,
                'memory_gb'=>0,
                'memory_gb_used'=>0,
                'local_gb'=>0,
                'local_gb_used'=>0,
                'machine_count'=>0,
        );
        if(!empty($arr_hyper['hypervisors'])){
                foreach($arr_hyper['hypervisors'] as $onehyper){
                        if($onehyper['state']=='up' && $onehyper['status']=='enabled'){
                                $arr_ret['machine_count'] += 1;
                                $arr_ret['vcpus'] += $onehyper['vcpus'];
                                $arr_ret['vcpus_used'] += $onehyper['vcpus_used'];
                                $arr_ret['memory_mb'] += $onehyper['memory_mb'];
                                $arr_ret['memory_mb_used'] += $onehyper['memory_mb_used'];
                                $arr_ret['memory_gb'] += sprintf("%.2f", $onehyper['memory_mb']/1024);
                                $arr_ret['memory_gb_used'] += sprintf("%.2f", $onehyper['memory_mb_used']/1024);
                                $arr_ret['local_gb'] += $onehyper['local_gb'];
                                $arr_ret['local_gb_used'] += $onehyper['local_gb_used'];
                        }
                }
        }
	dolog('insert power: '.print_r($arr_ret, true));
	computepower::insertPower($arr_ret);
	sleep(60);
}

function dolog($msg){
        global $logid;
        printf('[logid:'.$logid.']['.date('Y-m-d H:i:s').'] '.$msg."\n");
}

?>
