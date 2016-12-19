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


require_once("../include/config.inc.php");
require_once('../include/function.php');
require_once('../include/func_session.php');
require_once('../include/navbar.php');
$thisClass=$navbar;

/*权限检查*/
$pageForSuper=true;//当前页面是否需要管理员权限
$hasLimit=($pageForSuper)?isSuper($myUser):false;
$pageDeny='
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">I Am Sorry!</h4>
        </div>
        <div class="modal-body">
          <div class="note note-danger">
            <h4 class="block">Permission Denied!</h4>
            <p>请联系管理员申请权限！</p>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        </div>';

$myAction=(isset($_POST['action'])&&!empty($_POST['action']))?trim($_POST['action']):((isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):"add");
$myIndex=(isset($_GET['index'])&&!empty($_GET['index']))?intval($_GET['index']):0;

switch($myAction){
  case "add":
    if(!$hasLimit) {echo $pageDeny;return;}
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">添加导航</h4>
        </div>
        <div class="modal-body" id="myModalBody" style="height:400px;overflow:auto;">
          <div class="form-group">
            <label for="nb_fid" class="col-sm-2 control-label">父页面ID</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_fid" name="nb_fid" placeholder="父页面ID" value="0">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_id" class="col-sm-2 control-label">页面ID</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_id" name="nb_id" placeholder="页面ID" value="" onkeyup="sort_id()">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_name" class="col-sm-2 control-label">页面名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_name" name="nb_name" placeholder="页面名称" onkeyup="desc()" value="">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_href" class="col-sm-2 control-label">页面链接</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_href" name="nb_href" placeholder="页面链接,eg:admin/user.php" value="">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_target" class="col-sm-2 control-label">跳转方式</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_target" name="nb_target" placeholder="跳转方式,eg:_self,_blank" value="_self">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_desc" class="col-sm-2 control-label">页面描述</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_desc" name="nb_desc" placeholder="页面描述" value="">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_sort" class="col-sm-2 control-label">排序</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_sort" name="nb_sort" placeholder="排序,eg:1" value="1">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_icon" class="col-sm-2 control-label">页面图标</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_icon" name="nb_icon" placeholder="页面图标" value="">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_new" class="col-sm-2 control-label">新页标识</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_new" name="nb_new" placeholder="新页标识,eg:danger,warning" value="">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_status" class="col-sm-2 control-label">状态</label>
            <div class="col-sm-10">
              <div style="padding-top: 8px;">
                <input type="radio" class="flat" name="nb_status" id="nb_status1" value="0" checked="" /> 激活
                <input type="radio" class="flat" name="nb_status" id="nb_status2" value="1" /> 关闭
                <input type="radio" class="flat" name="nb_status" id="nb_status3" value="2" /> 隐藏
              </div>
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="insert"></input>
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" onclick="change()" style="margin-bottom: 5px;">确认添加</button>
        </div>
        <script>
          function sort_id(){
            $("#nb_sort").val($("#nb_id").val());
          }
          function desc(){
            $("#nb_desc").val($("#nb_name").val());
          }
        </script>
<?php
    break;
  case "edit":
    if(!$hasLimit) {echo $pageDeny;return;}
    $arrFilter=$thisClass->getNav($myIndex);
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">修改导航</h4>
        </div>
        <div class="modal-body" id="myModalBody" style="height:400px;overflow:auto;">
          <div class="form-group">
            <label for="nb_fid" class="col-sm-2 control-label">父页面ID</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_fid" name="nb_fid" placeholder="父页面ID" value="<?php echo $arrFilter[$myIndex]['nb_fid'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_id" class="col-sm-2 control-label">页面ID</label>
            <div class="col-sm-10">
              <input type="hidden" class="form-control" id="id" name="id" value="<?php echo $arrFilter[$myIndex]['nb_id'];?>">
              <input type="text" class="form-control" id="nb_id" name="nb_id" placeholder="页面ID" value="<?php echo $arrFilter[$myIndex]['nb_id'];?>" onkeyup="sort_id()">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_name" class="col-sm-2 control-label">页面名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_name" name="nb_name" placeholder="页面名称" onkeyup="desc()" value="<?php echo $arrFilter[$myIndex]['nb_name'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_href" class="col-sm-2 control-label">页面链接</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_href" name="nb_href" placeholder="页面链接,eg:admin/user.php" value="<?php echo $arrFilter[$myIndex]['nb_href'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_target" class="col-sm-2 control-label">跳转方式</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_target" name="nb_target" placeholder="跳转方式,eg:_self,_blank" value="<?php echo $arrFilter[$myIndex]['nb_target'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_desc" class="col-sm-2 control-label">页面描述</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_desc" name="nb_desc" placeholder="页面描述" value="<?php echo $arrFilter[$myIndex]['nb_desc'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_sort" class="col-sm-2 control-label">排序</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_sort" name="nb_sort" placeholder="排序,eg:1" value="<?php echo $arrFilter[$myIndex]['nb_sort'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_icon" class="col-sm-2 control-label">页面图标</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_icon" name="nb_icon" placeholder="页面图标" value="<?php echo $arrFilter[$myIndex]['nb_icon'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="nb_new" class="col-sm-2 control-label">新页标识</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="nb_new" name="nb_new" placeholder="新页标识,eg:danger,warning" value="<?php echo $arrFilter[$myIndex]['nb_new'];?>">
            </div>
          </div>
          <div class="form-group">
            <label for="t_switch" class="col-sm-2 control-label">状态</label>
            <div class="col-sm-10">
              <div style="padding-top: 8px;">
<?php
  switch($arrFilter[$myIndex]['nb_status']){
    case '0':
?>
                  <input type="radio" class="flat" name="nb_status" id="nb_status1" value="0" checked="" /> 激活
                  <input type="radio" class="flat" name="nb_status" id="nb_status2" value="1" /> 关闭
                  <input type="radio" class="flat" name="nb_status" id="nb_status3" value="2" /> 隐藏
<?php
    break;
    case '1':
?>
                  <input type="radio" class="flat" name="nb_status" id="nb_status1" value="0" /> 激活
                  <input type="radio" class="flat" name="nb_status" id="nb_status2" value="1" checked="" /> 关闭
                  <input type="radio" class="flat" name="nb_status" id="nb_status3" value="2" /> 隐藏
<?php
    break;
    case '2':
?>
                  <input type="radio" class="flat" name="nb_status" id="nb_status1" value="0" /> 激活
                  <input type="radio" class="flat" name="nb_status" id="nb_status2" value="1" /> 关闭
                  <input type="radio" class="flat" name="nb_status" id="nb_status3" value="2" checked="" /> 隐藏
<?php
    break;
  }
?>
              </div>
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="update"></input>
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" onclick="change()" style="margin-bottom: 5px;">确认修改</button>
        </div>
        <script>
          function sort_id(){
            $("#nb_sort").val($("#nb_id").val());
          }
          function desc(){
            $("#nb_desc").val($("#nb_name").val());
          }
          $("input.flat").iCheck({checkboxClass:"icheckbox_flat-green",radioClass:"iradio_flat-green"});
        </script>
<?php
    break;
    case "del":
      if(!$hasLimit) {echo $pageDeny;return;}
      $arrFilter=$thisClass->getNav($myIndex);
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">确认删除？</h4>
        </div>
        <div class="modal-body" id="myModalBody">
<?php
  echo "          <input type=\"hidden\" id=\"nb_id\" name=\"nb_id\" value=\"{$arrFilter[$myIndex]['nb_id']}\" readonly>\n";
  echo "          <input type=\"hidden\" id=\"page_action\" name=\"page_action\" value=\"delete\" readonly>\n";
  echo "          <p>序号：{$arrFilter[$myIndex]['nb_id']}</p>\n";
  echo "          <p>父页面：{$arrFilter[$myIndex]['nb_fid']}</p>\n";
  echo "          <p>名称：{$arrFilter[$myIndex]['nb_name']}</p>\n";
  echo "          <p>链接：{$arrFilter[$myIndex]['nb_href']}</p>\n";
?>
          <div class="note note-danger"><p>警告，删除操作谨慎使用！</p></div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button class="btn btn-danger" data-dismiss="modal" onclick="change()" style="margin-bottom: 5px;">确认删除</button>
        </div>
<?php
    break;
    default:
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">错误请求</h4>
        </div>
        <div class="modal-body">
          <div class="note note-danger"><p>非法操作</p></div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        </div>
<?php
    break;
  }
?>
