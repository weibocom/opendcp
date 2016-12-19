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
    $myTitle='添加用户';
    $pageAction='insert';
    break;
  case 'edit':
    $myTitle='修改用户';
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
    <label for="type" class="col-sm-2 control-label">用户类型</label>
    <div class="col-sm-10">
      <select class="form-control" id="type" name="type" onchange="check()">
        <option value="local">本地用户</option>
        <option value="ldap">LDAP</option>
      </select>
    </div>
  </div>
  <div class="form-group">
    <label for="en" class="col-sm-2 control-label">账号</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="en" name="en" onkeyup="check()" placeholder="eg:admin">
    </div>
  </div>
  <div class="form-group">
    <label for="cn" class="col-sm-2 control-label">姓名</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="cn" name="cn" onkeyup="check()" placeholder="eg:管理员">
    </div>
  </div>
  <div class="form-group">
    <label for="mail" class="col-sm-2 control-label">邮箱</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="mail" name="mail" onkeyup="check()" placeholder="eg:admin@xxx.com">
    </div>
  </div>
  <div class="form-group">
    <label for="pw" class="col-sm-2 control-label">密码</label>
    <div class="col-sm-10">
      <input type="password" class="form-control" id="pw" name="pw" onkeyup="check()" placeholder="为空时不修改">
    </div>
  </div>
  <div class="form-group">
    <label for="status" class="col-sm-2 control-label">用户类型</label>
    <div class="col-sm-10">
      <select class="form-control" id="status" name="status" onchange="check()">
        <option value="0">启用</option>
        <option value="1">停用</option>
      </select>
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
