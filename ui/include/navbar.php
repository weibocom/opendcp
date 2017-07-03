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


class navbar{
  private $table='nav_bar';
  private $fields=array('nb_name','nb_href','nb_target','nb_desc','nb_icon','nb_new');
  
  function getCount($filter=''){
    global $db;
    $sqlWhere='';
    $arrField=$this->fields;
    if($filter){
      foreach($arrField as $field){
        $sqlWhere.=($sqlWhere)?" OR {$field} LIKE ?":" WHERE {$field} LIKE ?";
      }
    }
    $sql='SELECT COUNT(*) FROM ' . $this->table . $sqlWhere.' ORDER BY nb_fid,nb_sort,nb_id;';
    $stmt = $db->prepare($sql);
    $filter='%'.$filter.'%';
    $stmt->bind_param('ssssss', $filter, $filter, $filter, $filter, $filter, $filter);

    if($stmt->execute()){
      $result = $stmt->get_result();
      if($row=$result->fetch_row()) return $row[0];
    }
    return false;
  }
  
  function getNav($id=0,$filter=''){
    global $db;
    if($id!==''){
      $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . ' WHERE nb_id=? ORDER BY nb_fid,nb_sort,nb_id;';
    }else{
      $sqlWhere='';
      $arrField=$this->fields;
      if($filter){
        foreach($arrField as $field){
          $sqlWhere.=($sqlWhere)?" OR {$field} LIKE ?":" WHERE {$field} LIKE ?";
        }
      }
      $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . $sqlWhere.' ORDER BY nb_fid,nb_sort,nb_id DESC;';
    }
    $stmt = $db->prepare($sql);
    if($id!==''){
      $stmt->bind_param('d', $id);
    }else{
      $filter='%'.$filter.'%';
      $stmt->bind_param('ssssss', $filter, $filter, $filter, $filter, $filter, $filter);
    }
    if($stmt->execute()){
      $result = $stmt->get_result();
      $arrRe=array();
      if($id!==''){
        while($row=$result->fetch_array(MYSQLI_ASSOC)){
          $arrRe[$row['nb_id']]=$row;
        }
      }else{
        while($row=$result->fetch_array(MYSQLI_ASSOC)){
          $arrRe[$row['nb_fid']][$row['nb_id']]=$row;
        }
      }
      return $arrRe;
    }
    return false;
  }
  
  function getNav1($filter='',$page=1,$myPageSize=20){
    global $db;
    $page=($page>0)?$page:1;
    $pageBegin=$page*$myPageSize-$myPageSize;
    $sqlWhere='';
    $arrField=$this->fields;
    if($filter){
      foreach($arrField as $field){
        $sqlWhere.=($sqlWhere)?" OR {$field} LIKE ?":" WHERE {$field} LIKE ?";
      }
    }
    $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . $sqlWhere.' ORDER BY nb_fid,nb_sort,nb_id DESC LIMIT '.$pageBegin.','.$myPageSize.';';
    $stmt = $db->prepare($sql);
    if($filter!==''){
      $filter='%'.$filter.'%';
      $stmt->bind_param('ssssss', $filter, $filter, $filter, $filter, $filter, $filter);
    }
    if($stmt->execute()){
      $result = $stmt->get_result();
      $arrRe=array();
      while($row=$result->fetch_array(MYSQLI_ASSOC)){
        $arrRe[$row['nb_id']]=$row;
      }
      return $arrRe;
    }
    return false;
  }
  
  function getNavByHref($href=''){
    global $db;
    if($href){
      $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . " WHERE nb_href=? ORDER BY nb_fid,nb_sort,nb_id;";
      $stmt = $db->prepare($sql);
      $stmt->bind_param('s', $href);

      if($stmt->execute()){
        $result = $stmt->get_result();
        $arrRe=array();
        if($row=$result->fetch_array(MYSQLI_ASSOC)){
          $arrRe=$row;
        }
        return $arrRe;
      }
    }
    return false;
  }
  
  function add($arr){
    global $db;
    if($arr){
      $sqlKey='';$sqlValue='';
      foreach($arr as $k=>$v){
        $sqlKey.=($sqlKey)?',`'.$k.'`':'`'.$k.'`';
        $sqlValue.=($sqlValue)?", ?":"?";
      }
      $sql='INSERT INTO '.$this->table.' ('.$sqlKey.') VALUES ('.$sqlValue.');';
      $stmt = $db->prepare($sql);
      foreach($arr as $k=>$v){
        $stmt->mbind_param('s', $v);
      }
      if($stmt->execute()) return true;
    }
    return false;
  }
  
  function update($arrNew){
    global $db;
    if(isset($arrNew['id'])&&$arrNew['id']!==''){
      $id=$arrNew['id'];
      $arrOld=$this->getNav($id);
      $sqlSet='';
      $ret=', nb_id:'.$id;
      foreach($arrNew as $k=>$v){
        if($k=='id') continue;
        if($arrOld[$id][$k]!=$v){
          $sqlSet.=($sqlSet)?", `{$k}`=?":"`{$k}`=?";
          $ret.=", {$k}:{$v}";
        }
      }
      if($sqlSet){
        $sql='UPDATE '.$this->table.' SET '.$sqlSet." WHERE nb_id=?;";
        $stmt = $db->prepare($sql);
        foreach($arrNew as $k=>$v){
          if($k=='id') continue;
          if($arrOld[$id][$k]!=$v) {
            $stmt->mbind_param('s', $v);
          }
        }
        $stmt->mbind_param('d', $id);
        if($stmt->execute()) return $ret;
      }
    }
    return false;
  }
  
  function delete($id){
    global $db;
    if($id){
      if($arrOld=$this->getNav($id)){
        $sql='DELETE FROM '.$this->table.' WHERE nb_id=?;';
        $stmt = $db->prepare($sql);
        $stmt->bind_param('d', $id);
        if($stmt->execute()){
          $ret='';
          foreach($arrOld[$id] as $k=>$v){
            $ret.=', '.$k.':'.$v;
          }
          return $ret;
        }
      }
    }
    return false;
  }

  function getChild($name,$arr=array()){
    global $mySite,$myExt,$arrUri;
    $ret='';$navLi='';$flag=false;
    if(is_array($arr)&&!empty($arr)){
      foreach($arr as $k1=>$v1){
        if($arrUri[0]===$mySite.'/extlink.php'&&strpos($v1['nb_href'],$myExt)===0){
          $myExt=$v1['nb_href'];
          $pageExt=$v1['nb_name'];
          $flag=true;
        }
        if($v1['nb_status']>1){
          if($arrUri[0]===$mySite.'/'.$v1['nb_href']) $flag=true;
        }
        $liIcon=($v1['nb_icon'])?'<i class="'.$v1['nb_icon'].'"></i> ':'';
        $liNew=($v1['nb_new'])?'<span class="label label-'.$v1['nb_new'].' pull-right">new</span>':'';
        $liTarget=($v1['nb_target'])?' target="'.$v1['nb_target'].'"':'';
        $liClass=($v1['nb_status'])?'disabled-link':'';
        $hrClass=($v1['nb_status'])?' class="disable-target"':'';
        $liHref=($v1['nb_status'])?'javascript:;':((strpos($v1['nb_href'],'http')===0)?(($v1['nb_target']!='_self')?$v1['nb_href']:$mySite.'/extlink.php?site='.$v1['nb_href']):$mySite.'/'.$v1['nb_href']);
        if($arrUri[0]===$mySite.'/'.$v1['nb_href']){
          $flag=true;
          $navLi.='            <li class="current-page '.$liClass.'">'."\n";
        }else{
          $flag=false;
          if($liClass) $liClass=' class="disabled-link"';
          $navLi.='            <li'.$liClass.'>'."\n";
        }
        $curr_color=($flag)?' style="color: orange;"':(($v1['nb_status']>0)?' style="color: gray;"':'');
        $navLi.='              <a href="'.$liHref.'"'.$liTarget.$hrClass.$curr_color.'>'."\n";
        $navLi.='              '.$liIcon.$liNew.$v1['nb_name'].'</a>'."\n";
        $navLi.='            </li>'."\n";
      }
      $liClass='disabled-link';
      $ret.='        <li class="active '.$liClass.'">'."\n";
      $liTitle='<span class="title">'.$name.'</span>';
      $ret.='          <a href="javascript:;">'."\n";
      $ret.='          '.$liTitle.' <span class="fa fa-chevron-down"></span></a>'."\n";
      if($flag){
        $ret.='<ul class="nav child_menu" style="display: block;">'."\n".$navLi."\n".'</ul>'."\n";
      }else{
        $ret.='<ul class="nav child_menu">'."\n".$navLi."\n".'</ul>'."\n";
      }
      $ret.="        </li>\n";
    }
    return $ret;
  }
  
}
$navbar=new navbar();

/*权限检查*/
$pageForSuper = true;//当前页面是否需要管理员权限
$hasLimit = ($pageForSuper)?isSuper($myUser):true;

$arrUri=explode('?',$_SERVER['REQUEST_URI']);
if($mySite){
  $pageHref=str_replace($mySite.'/','',$arrUri[0]);
}else{
  $pageHref=substr($arrUri[0],1);
}
$pageInfo=$navbar->getNavByHref($pageHref);
$pageName=(isset($pageInfo['nb_name']))?$pageInfo['nb_name']:basename($pageHref);
$pageDesc=(isset($pageInfo['nb_desc']))?$pageInfo['nb_desc']:'';
$pageExt='';//外链页面名称
if(!isset($myExt)) $myExt='';

$arrNavBar=$navbar->getNav('');
$navLeft='';
if(!empty($arrNavBar[0])){
  foreach($arrNavBar[0] as $k=>$v){
    if($k<=0||$v['nb_status']>1){continue;}
    if($k === 90001 && !$hasLimit){continue;}//如果是系统管理并且不是超级管理员则跳过这个导航栏
    $flag=false;
    $navLi='';
    if(!empty($arrNavBar[$k])){
      foreach($arrNavBar[$k] as $k1=>$v1){
        if(!empty($arrNavBar[$k1])){
          $navLi.=$navbar->getChild($v1['nb_name'],$arrNavBar[$k1]);
          continue;
        }
        if($arrUri[0]===$mySite.'/extlink.php'&&strpos($v1['nb_href'],$myExt)===0){
          $myExt=$v1['nb_href'];
          $pageExt=$v1['nb_name'];
          $flag=true;
        }
        if($v1['nb_status']>1){
          if($arrUri[0]===$mySite.'/'.$v1['nb_href']) $flag=true;
          continue;
        }
        $liIcon=($v1['nb_icon'])?'<i class="'.$v1['nb_icon'].'"></i> ':'';
        $liNew=($v1['nb_new'])?'<span class="label label-'.$v1['nb_new'].' pull-right">new</span>':'';
        $liTarget=($v1['nb_target'])?' target="'.$v1['nb_target'].'"':'';
        if($v1['nb_status']){
          $liClass='disabled-link';
          $hrClass=' class="disable-target "';
          $liDisabled='<span class="label label-default pull-right">已停用</span>';
          $liHref='javascript:;';
        }else{
          $liClass='';
          $hrClass='';
          $liDisabled='';
          $liHref=(strpos($v1['nb_href'],'http')===0)?(($v1['nb_target']!='_self')?$v1['nb_href']:$mySite.'/extlink.php?site='.$v1['nb_href']):$mySite.'/'.$v1['nb_href'];
        }
        if($arrUri[0]===$mySite.'/'.$v1['nb_href']){
          $flag=true;
          $navLi.='            <li class="current-page '.$liClass.'">'."\n";
        }else{
          if($liClass) $liClass=' class="disabled-link"';
          $navLi.='            <li'.$liClass.'>'."\n";
        }
        $navLi.='              <a href="'.$liHref.'"'.$liTarget.$hrClass.'>'."\n";
        $navLi.='              '.$liIcon.$liNew.$liDisabled.$v1['nb_name'].'</a>'."\n";
        $navLi.='            </li>'."\n";
      }
    }
    if(strpos($navLi,'current-page')!==false) $flag=true;
    $liClass=($v['nb_status'])?'disabled-link':'';
    $hrClass=($v['nb_status'])?' class="disable-target"':'';
    $liHref=($v['nb_status'])?'javascript:;':((strpos($v['nb_href'],'http')===0)?(($v['nb_target']!='_self')?$v['nb_href']:$mySite.'/extlink.php?site='.$v['nb_href']):$mySite.'/'.$v['nb_href']);
    if($flag||$arrUri[0]===$mySite.'/'.$v['nb_href']){
      $navLeft.='        <li class="active '.$liClass.'" id="walkthrough_'.$k.'">'."\n";
    }else{
      if($liClass) $liClass=' class="disabled-link"';
      $navLeft.='        <li'.$liClass.' id="walkthrough_'.$k.'">'."\n";
    }
    $liIcon=($v['nb_icon'])?'<i class="'.$v['nb_icon'].'"></i> ':'';
    $liNew=($v['nb_new'])?'<span class="label label-'.$v['nb_new'].' pull-right">new</span>':'';
    $liTarget=($v['nb_target'])?' target="'.$v['nb_target'].'"':'';
    $liTitle='<span class="title">'.$v['nb_name'].'</span>';
    if(empty($navLi)){
      $navLeft.='          <a href="'.$liHref.'"'.$liTarget.$hrClass.'>'."\n";
      $navLeft.='          '.$liIcon.$liNew.$liTitle.' <span class="fa fa-chevron-down"></span></a>'."\n";
    }else{
      $navLeft.='          <a href="javascript:;">'."\n";
      $navLeft.='          '.$liIcon.$liNew.$liTitle.' <span class="fa fa-chevron-down"></span></a>'."\n";
    }
    if($flag){
      $navLeft.='<ul class="nav child_menu" style="display: block;">'."\n".$navLi."\n".'</ul>'."\n";
    }else{
      $navLeft.='<ul class="nav child_menu">'."\n".$navLi."\n".'</ul>'."\n";
    }
    $navLeft.="        </li>\n";
  }
}
if($pageExt) $pageName=$pageExt;

?>
