<?php

class node_init_log {
	static $table='node_init_log';

	static function getLogListByTaskid($task_id, $page = 1, $pagesize = 20){

		$page--;
		$sql = 'select * from '.self::$table.' where task_id='.$task_id.' order by id desc limit '.$page*$pagesize.','.$pagesize;
		global $db;
		$query = $db->query($sql);
		$arrRet=array();
		while($row=$query->fetch_array(MYSQL_ASSOC)){
			$row['data'] = @json_decode($row['data'], true);
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
	static function insertLog($data) {

		$now = time();
		$status = empty($data['status']) ? 0 : $data['status'];
		$sql = 'insert into '.self::$table.' (task_id, title, status, data, create_time) values (\''.@mysql_escape_string($data['task_id']).'\',\''.@mysql_escape_string($data['title']).'\', \''.@mysql_escape_string($status).'\',\''.@mysql_escape_string(json_encode($data['data'])).'\', '.$now.')';
		global $db;
		$db->query($sql);
		$id = $db->insert_id;
		return $id;
	}
}

?>
