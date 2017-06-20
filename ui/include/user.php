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
  private $fields=array('en','cn','mail','type');
  private $log_fields=array('t_user','t_module','t_action','t_desc');
  
  function getCount($filter=''){
    global $db;
    $sqlWhere='';
    $arrField=$this->fields;
    if($filter){
      foreach($arrField as $field){
        $sqlWhere.=($sqlWhere)?" OR {$field} LIKE ?":" WHERE {$field} LIKE ?";
      }
    }
    $sql='SELECT COUNT(*) FROM ' . $this->table . $sqlWhere;
    $stmt = $db->prepare($sql);
    $filter='%'.$filter.'%';
    $stmt->bind_param('ssss', $filter, $filter, $filter, $filter);

    if($stmt->execute()){
      $result = $stmt->get_result();
      if($row=$result->fetch_row()) return $row[0];
    }
    return false;
  }
  
  function get($id=0,$filter='',$page=1,$myPageSize=20){
    global $db;
    $page=($page>0)?$page:1;
    $pageBegin=$page*$myPageSize-$myPageSize;
    if($id){
      $sql='SELECT * FROM ' . $this->table . " WHERE id = ? ORDER BY id;";
    }else{
      $sqlWhere='';
      $arrField=$this->fields;
      if($filter){
        foreach($arrField as $field){
          $sqlWhere.=($sqlWhere)?" OR {$field} LIKE ?":" WHERE {$field} LIKE ?";
        }
      }
      $sql='SELECT * FROM ' . $this->table . $sqlWhere.' ORDER BY id LIMIT '.$pageBegin.','.$myPageSize.';';
    }
    $stmt = $db->prepare($sql);
    if($id){
      $stmt->bind_param('d', $id);
    }else{
      $filter='%'.$filter.'%';
      $stmt->bind_param('ssss', $filter, $filter, $filter, $filter);
    }
    if($stmt->execute()){
      $result = $stmt->get_result();
      $arrRe=array();
      while($row=$result->fetch_array(MYSQLI_ASSOC)){
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
    $sqlFilter1.=($user)?"t_user = ?":'';
    $arrField=$this->log_fields;
    if($filter){
      foreach($arrField as $field){
        $sqlFilter2.=($sqlFilter2)?" OR {$field} LIKE ?":" {$field} LIKE ?";
      }
    }
    $sqlWhere=($sqlFilter1)?$sqlFilter1:'';
    if($sqlFilter2) $sqlWhere.=($sqlWhere)?' AND ('.$sqlFilter2.')':$sqlFilter2;
    if($sqlWhere) $sqlWhere='WHERE '.$sqlWhere;
    $sql='SELECT COUNT(*) FROM ' . $this->logTable . " {$sqlWhere} ORDER BY t_time DESC,id DESC;";
    $stmt = $db->prepare($sql);
    if($user){
      if($filter){
        $stmt->bind_param('sssss', $user, $filter, $filter, $filter, $filter);
      }else{
        $stmt->bind_param('s', $user);
      }
    }else{
      if($filter){
        $stmt->bind_param('ssss', $filter, $filter, $filter, $filter);
      }
    }
    if($stmt->execute()){
      $result = $stmt->get_result();
      if($row=$result->fetch_row()) return $row[0];
    }
    return false;
  }
  
  function getLog($user='',$filter='',$page=1,$myPageSize=20){
    global $db;
    $page=($page>0)?$page:1;
    $pageBegin=$page*$myPageSize-$myPageSize;
    $sqlFilter1='';$sqlFilter2='';
    $sqlFilter1.=($user)?"t_user = ?":'';
    $arrField=$this->log_fields;
    if($filter){
      foreach($arrField as $field){
        $sqlFilter2.=($sqlFilter2)?" OR {$field} LIKE ?":" {$field} LIKE ?";
      }
    }
    $sqlWhere=($sqlFilter1)?$sqlFilter1:'';
    if($sqlFilter2) $sqlWhere.=($sqlWhere)?' AND ('.$sqlFilter2.')':$sqlFilter2;
    if($sqlWhere) $sqlWhere='WHERE '.$sqlWhere;
    $sql='SELECT * FROM ' . $this->logTable . " {$sqlWhere} ORDER BY t_time DESC,id DESC LIMIT ".$pageBegin.','.$myPageSize.';';
    $stmt = $db->prepare($sql);
    if($user){
      if($filter){
        $stmt->bind_param('sssss', $user, $filter, $filter, $filter, $filter);
      }else{
        $stmt->bind_param('s', $user);
      }
    }else{
      if($filter){
        $stmt->bind_param('ssss', $filter, $filter, $filter, $filter);
      }
    }
    if($stmt->execute()){
      $result = $stmt->get_result();
      $arrRe=array();
      while($row=$result->fetch_array(MYSQLI_ASSOC)){
        $arrRe[$row['id']]=$row;
      }
      return $arrRe;
    }
    return false;
  }
  
  function getLogById($id=''){
    global $db;
    if($id){
      $sql='SELECT * FROM ' . $this->logTable . " WHERE id = ? ORDER BY t_time DESC,id DESC;";
      $stmt = $db->prepare($sql);
      $stmt->bind_param('s', $id);
      if($stmt->execute()){
        $result = $stmt->get_result();
        $arrRe=array();
        while($row=$result->fetch_array(MYSQLI_ASSOC)){
          $arrRe[$row['id']]=$row;
        }
        return $arrRe;
      }
    }
    return false;
  }

  function add($arr, $flag = true){
    global $db;
    $ret = array('code' => 1, 'msg' => 'param error', 'content' => '');
    if($arr){
      $sqlKey='';$sqlValue='';
      foreach($arr as $k=>$v){
        if($k=='pw'&&$v==='') continue;
        $sqlKey.=($sqlKey)?',`'.$k.'`':'`'.$k.'`';
        $sqlValue.=($sqlValue)?", ?":"?";
      }
      $sql='INSERT INTO '.$this->table.' ('.$sqlKey.') VALUES ('.$sqlValue.')';
      $stmt = $db->prepare($sql);
      foreach($arr as $k=>$v){
        if($k=='pw'&&$v==='') continue;
        if($k=='pw'){
          if($flag) $v=md5($v);
        }
        $stmt->mbind_param('s', $v);
      }
      if($stmt->execute()){
        $ret['code'] = 0;
        $ret['content'] = $stmt->insert_id;
      }
      $ret['msg'] = $stmt->error;
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
          $sqlSet.=($sqlSet)?", `{$k}`=?":"`{$k}`=?";
        }
      }
      if($sqlSet){
        $ret['older'] = $arrOld[$id];
        $sql='UPDATE '.$this->table.' SET '.$sqlSet." WHERE id=?";
        $stmt = $db->prepare($sql);
        foreach($arrNew as $k=>$v){
          if($k=='id') continue;
          if($k=='pw'&&$v==='') continue;
          if($k=='pw') $v=md5($v);
          if($arrOld[$id][$k]!=$v) {
            $stmt->mbind_param('s', $v);
          }
        }
        $stmt->mbind_param('d', $id);
        if($stmt->execute()){
          $ret['code'] = 0;
        }
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
        $sql='DELETE FROM '.$this->table.' WHERE id = ?';
        $ret['content'] = $arrOld[$arr['id']];
        $stmt = $db->prepare($sql);
        $stmt->bind_param('d', $arr['id']);
        if($stmt->execute()){
          $ret['code'] = 0;
        }
        $ret['msg'] = $stmt->error;
      }
    }
    return $ret;
  }
  
  function deleteLog($id){
    global $db;
    if($id){
      if($arrOld=$this->getLogById($id)){
        if(strlen($arrOld[$id]['desc'])<1000) return false;
        $sql='DELETE FROM '.$this->logTable.' WHERE id = ?';
        $stmt = $db->prepare($sql);
        $stmt->bind_param('d', $id);
        if($stmt->execute()) return true;
      }
    }
    return false;
  }
  
}

$user=new user();
?>
