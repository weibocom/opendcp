<?php

class node_init {
	static $table='node_init';
	static $arr_status = array(
		0 => '准备执行',
		1 => '正在执行',
		10 => '执行成功',
		11 => '执行失败',
	);
	static $arr_type = array(
		1 => '计算节点初始化',
		2 => '控制节点初始化',
		3 => '存储节点初始化',
	);

	static function getOneNodeInit($id){
		if(empty($id)) return false;
		global $db;
		$sql = 'select * from '.self::$table.' where id='.$id;
		$query = $db->query($sql);
		return $query->fetch_array(MYSQL_ASSOC);
	}
	static function getOneTaskByIp($ip){
		$sql = 'select * from '.self::$table.' where ip=\''.@mysql_escape_string($ip).'\' and status=1 order by id desc limit 1';
		global $db;
		$query = $db->query($sql);
		$arrRet=array();
		$row=$query->fetch_array(MYSQL_ASSOC);
		return empty($row) ? array() : $row;
	}

	static function getOneDiskNameByIp($ip){

        return self::getOneTaskByIp($ip);
	}
	static function getNodeInitList($page = 1, $pagesize = 20){

		$page--;
		$sql = 'select * from '.self::$table.' order by id desc limit '.$page*$pagesize.','.$pagesize;
		global $db;
		$query = $db->query($sql);
		$arrRet=array();
		while($row=$query->fetch_array(MYSQL_ASSOC)){
			$arrRet[]=$row;
		}

		$sql = 'select count(*) as num from '.self::$table;
		$query = $db->query($sql);
		$rowc=$query->fetch_array(MYSQL_ASSOC);
		return array(
			'count'=>$rowc['num'],
			'data'=>$arrRet,
		);
	}
	static function insertNodeInit($data) {

		$now = time();
		$type = empty($data['type']) ? 1 : $data['type'];
		$sql = 'insert into '.self::$table.' (ip, password, type, create_time, disk_name) values (\''.@mysql_escape_string($data['ip']).'\',\''.@mysql_escape_string($data['password']).'\',\''.$type.'\',\' '.$now.'\',\''.@mysql_escape_string($data['disk_name']).'\')';
		global $db;
		$ret = $db->query($sql);
		$id = $db->insert_id;
		if(!empty($id) && $type==2){
			require_once('keydata.php');
			keydata::update('controller_ip', $data['ip']);
			require_once('cloud.php');
			$mycloud = new cloud();
			$ret = $mycloud->get('root', 'instance/openstack', 'POST', array(
				'OpIp'=>$data['ip'],
				'OpPort'=>'5000',
				'OpUserName'=>'admin',
				'OpPassWord'=>'root',
			));

		}
		return $id;
	}
	//在数据库中更新节点状态
	static function modifyOneNodeInit($id, $data){
		global $db;
		$now = time();
		$sql = 'update '.self::$table.' set ';
		foreach($data as $k=>$v){
			$sql .= $k.'=\''.@mysql_escape_string($v).'\'';
		}
		$sql .= ' where id='.$id;
	    	$query = $db->query($sql);
		return $query;
	}
}

?>
