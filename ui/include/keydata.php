<?php

class keydata {
	static $table='keydata';

	static function getContentByKey($key){
		if(empty($key)) return false;
		global $db;
		$sql = 'select * from '.self::$table.' where datakey=\''.@mysql_escape_string($key).'\'';
		$query = $db->query($sql);
		$row = $query->fetch_array(MYSQL_ASSOC);
		$row['datacontent'] = @json_decode($row['datacontent'], true);
		return $row['datacontent'];
	}
	static function insert($key, $content) {

		$now = time();
		$sql = 'insert into '.self::$table.' (datakey, datacontent) values (\''.@mysql_escape_string($key).'\',\''.@mysql_escape_string(json_encode($content)).'\')';
		global $db;
		$db->query($sql);
		return true;
	}
	static function update($key, $content){
		global $db;
		$now = time();
		$sql = 'update '.self::$table.' set datacontent=\''.@mysql_escape_string(json_encode($content)).'\' where datakey=\''.@mysql_escape_string($key).'\'';
	    	$db->query($sql);
		return true;
	}
}

?>
