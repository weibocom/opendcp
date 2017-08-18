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
      $myTitle='新增资源配置';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改资源配置';
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
            <label for="name" class="col-sm-2 control-label">资源类型</label>
            <div class="col-sm-10">
                <select class="form-control" id="resource_type" name="resource_type" onchange="hiddenFile()">
                    <option value="task">task</option>
                    <option value="template">template</option>
                    <option value="var">var</option>
                    <option value="file">file</option>
                    <option value="handle">handle</option>
                    <option value="meta">meta</option>
                </select>
            </div>
          </div>
          <div class="form-group">
            <label for="desc" class="col-sm-2 control-label">资源名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name"  placeholder="资源名称">
            </div>
          </div>
          <div class="form-group">
            <label for="desc" class="col-sm-2 control-label">描述</label>
            <div class="col-sm-10">
                <input type="text" class="form-control" id="desc" name="desc"  placeholder="资源描述">
            </div>
           </div>
          <div id="template_show" hidden>
          <div class="form-group">
            <label for="args" class="col-sm-2 control-label">文件路径</label>
             <div class="col-sm-10">
                  <input type="text" class="form-control" id="template_file_path" name="template_file_path" placeholder="配置文件拷贝路径">
             </div>
          </div>
          <div class="form-group">
              <label for="args" class="col-sm-2 control-label">文件权限</label>
              <div class="col-sm-10">
                  <input type="text" class="form-control" id="template_file_perm" name="template_file_perm"  placeholder="描述,eg:测试">
              </div>
          </div>
           <div class="form-group">
                <label for="args" class="col-sm-2 control-label">文件Owner</label>
                <div class="col-sm-10">
                    <input type="text" class="form-control" id="template_file_owner" name="template_file_owner"  placeholder="描述,eg:测试">
                </div>
           </div>
          </div>
          <div class="form-group">
            <label for="template" class="col-sm-2 control-label">资源内容</label>
            <div class="col-sm-10">
                  <textarea rows="12" class="form-control" id="resource_content" name="resource_content" "></textarea>
            </div>
          </div>
          <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;">确认</button>
        </div>
        <script>
          <?php if($myAction=='edit'){echo 'get(\''.$myIdx.'\');'."\n";} ?>
        </script>
