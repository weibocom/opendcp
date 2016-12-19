<?php
/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */


class user{
  private $table='member';
  private $logTable='user_log';
  private $fields=array('t_en','t_cn','t_mail','t_type');
  private $log_fields=array('t_user','t_module','t_action','t_desc');
  
  function getCount($filter=''){
    global $db;
    $sqlWhere='';
    $arrField=$this->fields;
    if($filter){
      foreach($arrField as $field){
        $sqlWhere.=($sqlWhere)?" OR {$field} LIKE '%{$filter}%'":" WHERE {$field} LIKE '%{$filter}%'";
      }
    }
    $sql='SELECT COUNT(*) FROM ' . $this->table . $sqlWhere.' ORDER BY id;';
    if($query=$db->query($sql)){
      if($row=$query->fetch_row()) return $row[0];
    }
    return false;
  }
  
  function get($id=0,$filter='',$page=1,$myPageSize=20){
    global $db;
    $page=($page>0)?$page:1;
    $pageBegin=$page*$myPageSize-$myPageSize;
    if($id){
      $sql='SELECT * FROM ' . $this->table . " WHERE id = {$id} ORDER BY id;";
    }else{
      $sqlWhere='';
      $arrField=$this->fields;
      if($filter){
        foreach($arrField as $field){
          $sqlWhere.=($sqlWhere)?" OR {$field} LIKE '%{$filter}%'":" WHERE {$field} LIKE '%{$filter}%'";
        }
      }
      $sql='SELECT * FROM ' . $this->table . $sqlWhere.' ORDER BY id LIMIT '.$pageBegin.','.$myPageSize.';';
    }
    if($query=$db->query($sql)){
      $arrRe=array();
      while($row=$query->fetch_array(MYSQL_ASSOC)){
        unset($row['pw']);
        $arrRe[$row['id']]=$row;
      }
      return $arrRe;
    }
    return false;
  }
  
  function getLogCount($user='',$filter=''){
    global $db;
    $sqlFilter1='';$sqlFilter2='';
    $sqlFilter1.=($user)?"t_user='{$user}'":'';
    $arrField=$this->log_fields;
    if($filter){
      foreach($arrField as $field){
        $sqlFilter2.=($sqlFilter2)?" OR {$field} LIKE '%{$filter}%'":" {$field} LIKE '%{$filter}%'";
      }
    }
    $sqlWhere=($sqlFilter1)?$sqlFilter1:'';
    if($sqlFilter2) $sqlWhere.=($sqlWhere)?' AND ('.$sqlFilter2.')':$sqlFilter2;
    if($sqlWhere) $sqlWhere='WHERE '.$sqlWhere;
    $sql='SELECT COUNT(*) FROM ' . $this->logTable . " {$sqlWhere} ORDER BY t_time DESC,id DESC;";
    if($query=$db->query($sql)){
      if($row=$query->fetch_row()) return $row[0];
    }
    return false;
  }
  
  function getLog($user='',$filter='',$page=1,$myPageSize=20){
    global $db;
    $page=($page>0)?$page:1;
    $pageBegin=$page*$myPageSize-$myPageSize;
    $sqlFilter1='';$sqlFilter2='';
    $sqlFilter1.=($user)?"t_user='{$user}'":'';
    $arrField=$this->log_fields;
    if($filter){
      foreach($arrField as $field){
        $sqlFilter2.=($sqlFilter2)?" OR {$field} LIKE '%{$filter}%'":" {$field} LIKE '%{$filter}%'";
      }
    }
    $sqlWhere=($sqlFilter1)?$sqlFilter1:'';
    if($sqlFilter2) $sqlWhere.=($sqlWhere)?' AND ('.$sqlFilter2.')':$sqlFilter2;
    if($sqlWhere) $sqlWhere='WHERE '.$sqlWhere;
    $sql='SELECT * FROM ' . $this->logTable . " {$sqlWhere} ORDER BY t_time DESC,id DESC LIMIT ".$pageBegin.','.$myPageSize.';';
    if($query=$db->query($sql)){
      $arrRe=array();
      while($row=$query->fetch_array(MYSQL_ASSOC)){
        $arrRe[$row['id']]=$row;
      }
      return $arrRe;
    }
    return false;
  }
  
  function getLogById($id=''){
    global $db;
    if($id){
      $sql='SELECT * FROM ' . $this->logTable . " WHERE id={$id} ORDER BY t_time DESC,id DESC;";
      if($query=$db->query($sql)){
        $arrRe=array();
        while($row=$query->fetch_array(MYSQL_ASSOC)){
          $arrRe[$row['id']]=$row;
        }
        return $arrRe;
      }
    }
    return false;
  }

  function add($arr){
    global $db;
    $ret = array('code' => 1, 'msg' => 'param error', 'content' => '');
    if($arr){
      $sqlKey='';$sqlValue='';
      foreach($arr as $k=>$v){
        if($k=='pw'&&$v==='') continue;
        $sqlKey.=($sqlKey)?','.$k:$k;
        if($k=='pw') $v=md5($v);
        $sqlValue.=($sqlValue)?",'".$v."'":"'".$v."'";
      }
      $sql='INSERT INTO '.$this->table.' ('.$sqlKey.') VALUES ('.$sqlValue.');';
      if($query=$db->query($sql)){
        $ret = array('code' => 0, 'msg' => $db->error, 'content' => $db->insert_id);
      }else{
        $ret['msg'] = $db->error;
      }
    }
    return $ret;
  }
  
  function update($arrNew){
    global $db;
    $ret = array('code' => 1, 'msg' => 'param error', 'content' => $arrNew, 'older'=>'');
    if(isset($arrNew['id'])&&!empty($arrNew['id'])){
      $id=$arrNew['id'];
      $arrOld=$this->get($id);
      $sqlSet='';
      foreach($arrNew as $k=>$v){
        if($k=='id') continue;
        if($k=='pw'&&$v==='') continue;
        if($k=='pw') $v=md5($v);
        if($arrOld[$id][$k]!=$v){
          $sqlSet.=($sqlSet)?",{$k}='{$v}'":"{$k}='{$v}'";
        }
      }
      if($sqlSet){
        $sql='UPDATE '.$this->table.' SET '.$sqlSet." WHERE id={$id};";
        $ret['older'] = $arrOld[$id];
        if($query=$db->query($sql)) $ret['code'] = 0;
        $ret['msg'] = $db->error;
      }
    }
    return $ret;
  }

  function delete($arr=array()){
    global $db;
    $ret = array('code' => 1, 'msg' => 'id null', 'content' => '');
    if(isset($arr['id'])&&!empty($arr['id'])){
      if($arrOld=$this->get($arr['id'])){
        $sql='DELETE FROM '.$this->table.' WHERE id='.$arr['id'].';';
        $ret['content'] = $arrOld[$arr['id']];
        if($query=$db->query($sql)){
          $ret['code'] = 0;
        }
        $ret['msg'] = $db->error;
      }
    }
    return $ret;
  }
  
  function deleteLog($id){
    global $db;
    if($id){
      if($arrOld=$this->getLogById($id)){
        if(strlen($arrOld[$id]['desc'])<1000);return false;
        $sql='DELETE FROM '.$this->logTable.' WHERE id='.$id.';';
        if($query=$db->query($sql)) return true;
      }
    }
    return false;
  }
  
}

$user=new user();
?>
