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
  $myTitle = '添加网络';
  $pageAction = 'addcompute';
  $ip = empty($_REQUEST['ip']) ? '' : $_REQUEST['ip'];
  $type = empty($_REQUEST['type']) ? 1 : $_REQUEST['type'];
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel"><?php echo $myTitle;?></h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">openstack网络名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="network_name" name="network_name" placeholder="网络名称" value="<?=$ip?>">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">物理网络名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="network_providername" name="network_providername" placeholder="provider" value="provider">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">子网名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="network_subnet_name" name="network_subnet_name" placeholder="">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">子网网段</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="network_subnet_range" name="network_subnet_range" placeholder="例如,192.168.11.0/24">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">子网网关</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="network_subnet_gateway" name="network_subnet_gateway" placeholder="例如,192.168.11.1">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">子网ip段</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="network_subnet_ip" name="network_subnet_ip" placeholder="例如,192.168.11.10,192.168.11.12">
            </div>
          </div>
          <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="type" name="type" value="<?=$type?>" />
          <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="addnetwork()" style="margin-bottom: 5px;">确认</button>
        </div>
        <script>

function addnetwork(){	
  url = '/api/for_openstack/network.php';
  network_name = $('#network_name').val();
  network_providername = $('#network_providername').val();
  network_subnet_name = $('#network_subnet_name').val();
  network_subnet_range = $('#network_subnet_range').val();
  network_subnet_gateway = $('#network_subnet_gateway').val();
  network_subnet_ip = $('#network_subnet_ip').val();
  postData = {"action": "addnetwork", "network_name": network_name, "network_providername": network_providername, "network_subnet_name": network_subnet_name, "network_subnet_range":network_subnet_range, "network_subnet_gateway": network_subnet_gateway, "network_subnet_ip": network_subnet_ip};
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (listdata) {
      if(listdata.code==0){
	location.reload();
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
