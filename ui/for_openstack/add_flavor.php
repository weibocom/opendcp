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
  $myTitle = '添加机型';
  $pageAction = 'addflavor';
  $vcpus = empty($_REQUEST['vcpus']) ? '' : $_REQUEST['vcpus'];
  $ram = empty($_REQUEST['ram']) ? 1 : $_REQUEST['ram'];
  $disk = empty($_REQUEST['disk']) ? 1 : $_REQUEST['disk'];
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel"><?php echo $myTitle;?></h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">机型名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="flavor_name" name="flavor_name" placeholder="机型名称" value="<?=$ip?>">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">vcpu核数</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="flavor_vcpus" name="flavor_vcpus" placeholder="vcpu核数">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">内存大小</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="flavor_ram" name="flavor_ram" placeholder="内存大小，M为单位">
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">硬盘大小</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="flavor_disk" name="flavor_disk" placeholder="硬盘大小，G为单位">
            </div>
          </div>
          <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="addflavor()" style="margin-bottom: 5px;">确认</button>
        </div>
        <script>

function addflavor(){	
  url = '/api/for_openstack/flavor.php';
  flavor_vcpus = $('#flavor_vcpus').val();
  flavor_ram = $('#flavor_ram').val();
  flavor_disk = $('#flavor_disk').val();
  flavor_name = $('#flavor_name').val();

  postData = {"action": "addflavor", "flavor_vcpus":  flavor_vcpus, "flavor_name": flavor_name, "flavor_ram": flavor_ram, "flavor_disk": flavor_disk};
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
