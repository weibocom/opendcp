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


$myAction=(isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'add';
$myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
switch($myAction){
  case 'add':
    $myTitle='添加业务方';
    $pageAction='insert';
    break;
  case 'edit':
    $myTitle='修改业务方';
    $pageAction='update';
    break;
  default:
    $myTitle='错误请求';
    $pageAction='Illegal';
    break;
}
?>
<div class="modal-header">
  <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
  <h4 class="modal-title" id="myModalLabel"><?php echo $myTitle;?></h4>
</div>
<div class="modal-body" style="overflow:auto;" id="myModalBody">
  <div class="form-group">
    <label for="name" class="col-sm-2 control-label">业务方名称</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="name" name="name" onkeyup="check()" placeholder="eg:新浪微博">
    </div>
  </div>
  <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
  <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
</div>
<div class="modal-footer">
  <button class="btn btn-default" data-dismiss="modal">取消</button>
  <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
</div>
<script>
  <?php if($myAction=='edit'){echo 'get(\''.$myIdx.'\');'."\n";} ?>
</script>
