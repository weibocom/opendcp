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


  $myParId=(isset($_GET['par_id'])&&!empty($_GET['par_id']))?trim($_GET['par_id']):'';
  $myParName=(isset($_GET['par_name'])&&!empty($_GET['par_name']))?trim($_GET['par_name']):'';
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">生成新版本 - <?php echo $myParName;?></h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="unit_id" class="col-sm-2 control-label">隶属单元</label>
            <div class="col-sm-10">
              <select class="form-control" id="unit_id" name="unit_id" disabled>
                <option value="<?php echo $myParId;?>"><?php echo $myParName;?></option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">版本描述</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" onkeyup="check('ver_auto')" onchange="check('ver_auto')" placeholder="名称,eg:添加规则">
            </div>
          </div>
          <div class="form-group">
            <label for="type" class="col-sm-2 control-label">发布方式</label>
            <div class="col-sm-10">
              <select class="form-control" id="type" name="type" onchange="check('ver_auto')">
                <option value="ANSIBLE">Ansible</option>
              </select>
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="generate" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          cache.files_list=[];
        </script>
