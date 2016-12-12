<?php
/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
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
        $sqlWhere.=($sqlWhere)?" OR {$field} LIKE '%{$filter}%'":" WHERE {$field} LIKE '%{$filter}%'";
      }
    }
    $sql='SELECT COUNT(*) FROM ' . $this->table . $sqlWhere.' ORDER BY nb_fid,nb_sort,nb_id;';
    if($query=$db->query($sql)){
      if($row=$query->fetch_row()) return $row[0];
    }
    return false;
  }
  
  function getNav($id=0,$filter=''){
    global $db;
    if($id!==''){
      $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . ' WHERE nb_id='.$id.' ORDER BY nb_fid,nb_sort,nb_id;';
    }else{
      $sqlWhere='';
      $arrField=$this->fields;
      if($filter){
        foreach($arrField as $field){
          $sqlWhere.=($sqlWhere)?" OR {$field} LIKE '%{$filter}%'":" WHERE {$field} LIKE '%{$filter}%'";
        }
      }
      $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . $sqlWhere.' ORDER BY nb_fid,nb_sort,nb_id DESC;';
    }
    if($query=$db->query($sql)){
      $arrRe=array();
      if($id!==''){
        while($row=$query->fetch_array(MYSQL_ASSOC)){
          $arrRe[$row['nb_id']]=array(
            'nb_id'=>"{$row['nb_id']}",
            'nb_fid'=>"{$row['nb_fid']}",
            'nb_name'=>"{$row['nb_name']}",
            'nb_href'=>"{$row['nb_href']}",
            'nb_target'=>"{$row['nb_target']}",
            'nb_desc'=>"{$row['nb_desc']}",
            'nb_sort'=>"{$row['nb_sort']}",
            'nb_icon'=>"{$row['nb_icon']}",
            'nb_new'=>"{$row['nb_new']}",
            'nb_status'=>"{$row['nb_status']}",
          );
        }
      }else{
        while($row=$query->fetch_array(MYSQL_ASSOC)){
          $arrRe[$row['nb_fid']][$row['nb_id']]=array(
            'nb_id'=>"{$row['nb_id']}",
            'nb_fid'=>"{$row['nb_fid']}",
            'nb_name'=>"{$row['nb_name']}",
            'nb_href'=>"{$row['nb_href']}",
            'nb_target'=>"{$row['nb_target']}",
            'nb_desc'=>"{$row['nb_desc']}",
            'nb_sort'=>"{$row['nb_sort']}",
            'nb_icon'=>"{$row['nb_icon']}",
            'nb_new'=>"{$row['nb_new']}",
            'nb_status'=>"{$row['nb_status']}",
          );
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
        $sqlWhere.=($sqlWhere)?" OR {$field} LIKE '%{$filter}%'":" WHERE {$field} LIKE '%{$filter}%'";
      }
    }
    $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . $sqlWhere.' ORDER BY nb_fid,nb_sort,nb_id DESC LIMIT '.$pageBegin.','.$myPageSize.';';
    if($query=$db->query($sql)){
      $arrRe=array();
      while($row=$query->fetch_array(MYSQL_ASSOC)){
        $arrRe[$row['nb_id']]=array(
          'nb_id'=>"{$row['nb_id']}",
          'nb_fid'=>"{$row['nb_fid']}",
          'nb_name'=>"{$row['nb_name']}",
          'nb_href'=>"{$row['nb_href']}",
          'nb_target'=>"{$row['nb_target']}",
          'nb_desc'=>"{$row['nb_desc']}",
          'nb_sort'=>"{$row['nb_sort']}",
          'nb_icon'=>"{$row['nb_icon']}",
          'nb_new'=>"{$row['nb_new']}",
          'nb_status'=>"{$row['nb_status']}",
        );
      }
      return $arrRe;
    }
    return false;
  }
  
  function getNavByHref($href=''){
    global $db;
    if($href){
      $sql='SELECT nb_id,nb_fid,nb_name,nb_href,nb_target,nb_desc,nb_sort,nb_icon,nb_new,nb_status FROM ' . $this->table . " WHERE nb_href='".$href."' ORDER BY nb_fid,nb_sort,nb_id;";
      if($query=$db->query($sql)){
        $arrRe=array();
        if($row=$query->fetch_array(MYSQL_ASSOC)){
          $arrRe=array(
            'nb_id'=>"{$row['nb_id']}",
            'nb_fid'=>"{$row['nb_fid']}",
            'nb_name'=>"{$row['nb_name']}",
            'nb_href'=>"{$row['nb_href']}",
            'nb_target'=>"{$row['nb_target']}",
            'nb_desc'=>"{$row['nb_desc']}",
            'nb_sort'=>"{$row['nb_sort']}",
            'nb_icon'=>"{$row['nb_icon']}",
            'nb_new'=>"{$row['nb_new']}",
            'nb_status'=>"{$row['nb_status']}",
          );
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
        $sqlKey.=($sqlKey)?','.$k:$k;
        $sqlValue.=($sqlValue)?",'".$v."'":"'".$v."'";
      }
      $sql='INSERT INTO '.$this->table.' ('.$sqlKey.') VALUES ('.$sqlValue.');';
      if($query=$db->query($sql)) return true;
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
          $sqlSet.=($sqlSet)?",{$k}='{$v}'":"{$k}='{$v}'";
          $ret.=", {$k}:{$v}";
        }
      }
      if($sqlSet){
        $sql='UPDATE '.$this->table.' SET '.$sqlSet." WHERE nb_id={$id};";
        if($query=$db->query($sql)) return $ret;
      }
    }
    return false;
  }
  
  function delete($id){
    global $db;
    if($id){
      if($arrOld=$this->getNav($id)){
        $sql='DELETE FROM '.$this->table.' WHERE nb_id='.$id.';';
        if($query=$db->query($sql)){
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
