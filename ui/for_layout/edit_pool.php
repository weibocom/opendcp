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
  $myParId=(isset($_GET['par_id'])&&!empty($_GET['par_id']))?trim($_GET['par_id']):'';
  $myParName=(isset($_GET['par_name'])&&!empty($_GET['par_name']))?trim($_GET['par_name']):'';
  switch($myAction){
    case 'add':
      $myTitle='添加服务池';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改服务池';
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
            <label for="service_id" class="col-sm-2 control-label">隶属服务</label>
            <div class="col-sm-10">
              <select class="form-control" id="service_id" name="service_id" readonly>
                <option value="<?php echo $myParId;?>"><?php echo $myParName;?></option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" onkeyup="check('pool')" onchange="check('pool')" placeholder="名称,eg:test_pool" <?php if($myAction=='edit'){echo 'readonly';} ?>>
            </div>
          </div>
          <div class="form-group">
            <label for="desc" class="col-sm-2 control-label">描述</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="desc" name="desc" onkeyup="check('pool')" onchange="check('pool')" placeholder="描述,eg:测试服务池">
            </div>
          </div>
          <div class="form-group">
            <label for="vm_type" class="col-sm-2 control-label">机型模板</label>
            <div class="col-sm-10">
              <select class="form-control" id="vm_type" name="vm_type" onchange="check('pool')">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="sd_id" class="col-sm-2 control-label">服务发现类型</label>
            <div class="col-sm-10">
              <select class="form-control" id="sd_id" name="sd_id" onchange="check('pool')">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="tpl_expand" class="col-sm-2 control-label">扩容任务模板</label>
            <div class="col-sm-10">
              <select class="form-control" id="tpl_expand" name="tpl_expand" onchange="check('pool')">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="tpl_shrink" class="col-sm-2 control-label">缩容任务模板</label>
            <div class="col-sm-10">
              <select class="form-control" id="tpl_shrink" name="tpl_shrink" onchange="check('pool')">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="tpl_deploy" class="col-sm-2 control-label">上线任务模板</label>
            <div class="col-sm-10">
              <select class="form-control" id="tpl_deploy" name="tpl_deploy" onchange="check('pool')">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
          <input type="hidden" id="page_other" name="page_other" value="addpool">
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          <?php if($myAction=='edit'){echo '$(\'#service_id\').select2({disabled:true});'."\n";echo 'get(\''.$myIdx.'\');'."\n";}else{?>
          updateSelect('tpl_expand');
          updateSelect('tpl_shrink');
          updateSelect('tpl_deploy');
          updateSelect('sd_id');
          updateSelect('vm_type');
          <?php } ?>
        </script>
