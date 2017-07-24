<?php

class computepower {
	static $table='compute_power';

	static function getListByTime($time){
		$sql = 'select * from '.self::$table.' where create_time>=\''.$time.'\' order by id desc';
		global $db;
		$query = $db->query($sql);
		$arrRet=array();
		while($row=$query->fetch_array(MYSQL_ASSOC)){
			$arrRet[]=$row;
		}
		return $arrRet;
	}
	static function insertPower($data) {

		$now = time();
		$sql = 'insert into '.self::$table.' (data, create_time) values (\''.@mysql_escape_string(json_encode($data)).'\','.$now.')';
		global $db;
		$db->query($sql);
		$id = $db->insert_id;
		return $id;
	}
}

?>
