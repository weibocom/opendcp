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
          <h4 class="modal-title" id="myModalLabel">发布新版本 - <?php echo $myParName;?></h4>
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
              <input type="text" class="form-control" id="name" name="name" onkeyup="check()" onchange="check()" placeholder="名称,eg:添加规则">
            </div>
          </div>
          <div class="form-group">
            <label for="type" class="col-sm-2 control-label">发布方式</label>
            <div class="col-sm-10">
              <select class="form-control" id="type" name="type" onchange="check()">
                <option value="ANSIBLE">Ansible</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="actions" class="col-sm-2 control-label">文件</label>
            <div class="col-sm-10 profile_details">
              <div class="well profile_view col-sm-12">
                <div class="col-sm-12">
                </div>
                <div class="col-sm-12">
                  <table class="table table-striped table-hover">
                    <thead>
                    <tr>
                      <th>#</th>
                      <th>文件</th>
                      <th>版本</th>
                      <th>#</th>
                    </tr>
                    </thead>
                    <tbody id="files_list">
                    </tbody>
                  </table>
                  <a class="btn btn-primary btn-xs" data-toggle="modal" data-target="#myChildModal" href="edit_nginx_ver_files.php">添加</a>
                </div>
              </div>
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="insert" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          cache.file_ver=[];
        </script>
