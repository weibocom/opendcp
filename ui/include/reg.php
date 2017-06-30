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


class reg{
  private $table='reg';
  private $fields=array('en', 'cn', 'mobile', 'mail', 'biz');
  
  function getCount($filter=''){
    global $db;
    $sqlWhere='';
    $arrField=$this->fields;
    if($filter){
      foreach($arrField as $field){
        $sqlWhere.=($sqlWhere)?" OR {$field} LIKE ?":" WHERE {$field} LIKE ?";
      }
    }
    $sql='SELECT COUNT(*) FROM ' . $this->table . $sqlWhere.' ORDER BY id;';
    $stmt = $db->prepare($sql);
    $filter='%'.$filter.'%';
    $stmt->bind_param('sssss', $filter, $filter, $filter, $filter, $filter);

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
      $sql='SELECT * FROM ' . $this->table . $sqlWhere.' ORDER BY id DESC LIMIT '.$pageBegin.','.$myPageSize.';';
    }
    $stmt = $db->prepare($sql);
    if($id){
      $stmt->bind_param('d', $id);
    }else{
      $filter='%'.$filter.'%';
      $stmt->bind_param('sssss', $filter, $filter, $filter, $filter, $filter);
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

  function add($arr,$flag = false){
    global $db;
    $ret = array('code' => 3, 'msg' => 'param error', 'content' => '');
    if($arr){
      //判断业务方名称是否已存在
      if($this->isBIZNameExite($arr['biz'])){
          $ret['code'] = 1;
          return $ret;
      }
      //判断数据中是否有该用户，以及该用户的状态
      $usrsql='SELECT * FROM ' . $this->table . " WHERE en = ? ORDER BY en;";
      $usrstmt = $db->prepare($usrsql);
      $usrstmt->bind_param('s', $arr['en']);
      $usrexit = ture;
      $usrarrRe=array();
      if($usrstmt->execute()){
          $result = $usrstmt->get_result();
          while($row=$result->fetch_array(MYSQLI_ASSOC)){
              $usrarrRe[1]=$row;
          }
          if(empty($usrarrRe)){
              $usrexit = false;
          }
      }else{
          $ret['code'] = 3;
          $ret['msg'] = $usrstmt->error;
          return $ret;
      }
      //当用户是拒绝状态时允许默认用户此次为更改操作
      if($usrexit && $usrarrRe[1]['status'] == 1){
          $arr['id'] = $usrarrRe[1]['id'];
          $ret = $this->update($arr);
          if($ret['code'] !== 0) $ret['code'] = 3;
          return $ret;
      }
      if($usrexit && $usrarrRe[1]['status'] !== 1){
          $ret['code'] = 2;
          return $ret;
      }
      $sqlKey='';$sqlValue='';
      foreach($arr as $k=>$v){
        if($k=='pw'&&$v==='') continue;
        $sqlKey.=($sqlKey)?',`'.$k.'`':'`'.$k.'`';
        $sqlValue.=($sqlValue)?", ?":"?";
      }
      $sql='INSERT INTO '.$this->table.' ('.$sqlKey.') VALUES ('.$sqlValue.');';
      $stmt = $db->prepare($sql);
      foreach($arr as $k=>$v){
        if($k=='pw'&&$v==='') continue;
        if($flag === false){
          if($k=='pw') $v=md5($v);
        }
        $stmt->mbind_param('s', $v);
      }
      if($stmt->execute()){
        $ret['code'] = 0;
        $ret['content'] = $stmt->insert_id;
      }else{
          $ret['code'] = 3;
          $ret['msg'] = $stmt->error;
      }
    }
    return $ret;
  }
  function isBIZNameExite($bizname){
      global $db;
      $sql='SELECT * FROM ' . $this->table . " WHERE biz = ? and status != 1;";
      $stmt = $db->prepare($sql);
      $stmt->bind_param('s', $bizname);
      if($stmt->execute()){
          $result = $stmt->get_result();
          $usrarrRe = array();
          while($row=$result->fetch_array(MYSQLI_ASSOC)){
              $usrarrRe=$row;
          }
          if(empty($usrarrRe)){
              return false;
          }
      }
      return true;
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
        if($k=='pw' && $arrOld[$id][$k]!=$v) $v=md5($v);
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
          if($k=='pw' && $arrOld[$id][$k]!=$v) $v=md5($v);
          if($arrOld[$id][$k]!=$v) {
            $stmt->mbind_param('s', $v);
          }
        }
        $stmt->mbind_param('d', $id);
        if($stmt->execute()){
          $ret['code'] = 0;
        }
        $ret['msg'] = $db->error;
      }else{
          $ret['code'] = 0;
          $ret['msg'] ='';
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
  
}

$reg=new reg();
?>
