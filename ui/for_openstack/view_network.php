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


require_once('../include/config.inc.php');
require_once('../include/function.php');


  $myAction=(isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'add';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
  $myTitle = '网络详情';
  $pageAction = 'addcompute';
  $id = empty($_REQUEST['id']) ? '' : $_REQUEST['id'];
  include_once('../include/openstack.php');
  openstack::needOpenstackLogin();
  $onenetwork = openstack::getOneNetwork($id);
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel"><?php echo $myTitle;?></h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">网络名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" value="<?=$onenetwork['name']?>">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">状态</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" value="<?=$onenetwork['status']?>">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">子网</label>
            <div class="col-sm-10">
            </div>
          </div>
   	  <?php foreach($onenetwork['subnets'] as $onesubnet) { ?>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">子网信息</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" value="<?=$onesubnet['name']?> - <?=$onesubnet['gateway_ip']?> - <?=$onesubnet['cidr']?>">
            </div>
          </div>
	  <?php } ?>

          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">已分配端口</label>
            <div class="col-sm-10">
            </div>
          </div>
   	  <?php foreach($onenetwork['ports'] as $oneport) { ?>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">端口信息</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" value="<?=$oneport['device_owner']?> - <?=$oneport['fixed_ips'][0]['ip_address']?>">
            </div>
          </div>
	  <?php } ?>
          <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="type" name="type" value="<?=$type?>" />
          <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="addcompute()" style="margin-bottom: 5px;">确认</button>
        </div>
        <script>

function addcompute(){	
  url = '/api/for_openstack/machine.php';
  ip = $('#compute_ip').val();
  type = $('#type').val();
  password = $('#compute_password').val();
  postData = {"action": "addip", "ip": ip, "password": password, "type":type};
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (listdata) {
      if(listdata.code==0){
        NProgress.done();
      }else{
        pageNotify('error','加载失败！','错误信息：'+listdata.msg);
        NProgress.done();
      }
    },
    error: function (){
      pageNotify('error','加载失败！','错误信息：接口不可用');
      NProgress.done();
    }
  });
}


        </script>
