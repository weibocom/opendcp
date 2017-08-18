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
      $myTitle='添加远程命令';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改远程命令';
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
            <label for="name" class="col-sm-2 control-label">名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" onkeyup="check()" onchange="check()" placeholder="名称,eg:test" <?php if($myAction=='edit'){echo 'readonly';} ?>>
            </div>
          </div>
          <div class="form-group">
            <label for="desc" class="col-sm-2 control-label">描述</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="desc" name="desc" onkeyup="check()" onchange="check()" placeholder="描述,eg:测试">
            </div>
          </div>
          <div class="form-group">
            <label for="args" class="col-sm-2 control-label">参数定义</label>
            <div class="col-sm-10">
              <div class="panel panel-default" style="margin-bottom: 0px;">
                <div class="pannel-body" style="padding: 0px 15px 0px;">
                  <table class="table table-hover">
                    <thead>
                    <tr>
                      <td>参数名称</td>
                      <td>值类型</td>
                      <td>#</td>
                    </tr>
                    </thead>
                    <tbody id="args">
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
          <div class="form-group">
            <label for="type" class="col-sm-2 control-label">命令实现</label>
            <div class="col-sm-10">
              <select class="form-control" id="type" name="type" onchange="check()">
                <option value="ansible">Ansible</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="template" class="col-sm-2 control-label">参数</label>
            <div class="col-sm-10">
              <div class="panel panel-default" style="margin-bottom: 0px;">
                <div class="pannel-body" style="padding: 10px 15px 0px;">
                  <div class="form-group">
                    <label for="config_senior" class="col-sm-2 control-label">配置模式</label>
                    <div class="col-sm-10">
                      <select class="form-control" id="config_senior" onchange="updateSenior()">
                        <option value="false">默认</option>
                        <option value="true">专家模式</option>
                      </select>
                    </div>
                  </div>
                  <div class="form-group">
                    <label for="cmd_parent" class="col-sm-2 control-label">命令实现类型</label>
                    <div class="col-sm-10">
                      <select class="form-control" id="cmd_parent" name="cmd_parent" onchange="updateCmd()">
                        <option value="action">Action</option>
                      </select>
                    </div>
                  </div>
                  <div class="form-group">
                    <label for="cmd_child" class="col-sm-2 control-label">命令实现模板</label>
                    <div class="col-sm-10">
                      <select class="form-control" id="cmd_child" name="cmd_child"  onchange="showRole()">
                          <option value="shell">shell - 命令</option>
                          <option value="longscript">longscript - 脚本</option>
                          <option value="role">Ansible Role</option>
                      </select>
                    </div>
                  </div>
                  <div class="form-group">
                    <label for="cmd_content" class="col-sm-2 control-label">执行的命令/脚本</label>
                    <div class="col-sm-10">
                      <textarea rows="12" class="form-control" id="cmd_content" name="cmd_content" onkeyup="checkCmd()"></textarea>
                    </div>
                  </div>
                    <div class="form-group" hidden>
                        <label for="desc" class="col-sm-2 control-label">var文件</label>
                        <div id="vars_file" >

                        </div>
                    </div>
                    <div class="form-group" hidden>
                        <label for="desc" class="col-sm-2 control-label">task文件</label>
                        <div id="tasks_file" >

                        </div>
                    </div>
                    <div class="form-group" hidden>
                        <label for="desc" class="col-sm-2 control-label">template文件</label>
                        <div id="tems_file">

                        </div>
                    </div>

                </div>

              </div>
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
          <?php if($myAction=='edit'){echo 'get(\''.$myIdx.'\');'."\n";}else{echo 'cache.params={};'."\n".'showArg();'."\n";} ?>
//              $("#role").click(function(){
//                 alert("click role");
//                 var postData = {"action": "list", "pagesize": 1000};
//                 var url = '/api/for_layout/roleresource.php?action=list&pagesize=1000';
//                 var actionDesc="添加Role";
//                 $.ajax({
//                       type: "POST",
//                       url: url,
//                       data: {"action":'list',"data":JSON.stringify(postData)},
//                       dataType: "json",
//                       success: function (data) {
//                           //执行结果提示
//                           if(data.code==0){
//                               $('#var_file').html('');
//                               $('#template_file').html('');
//                               $('#task_file').html('');
//                               for(var i=0;i<data.content.length;i++){
//                                 var v = data.content[i];
//                                 switch (v.resource_type){
//                                     case 'var':
//                                       var var_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="'+v.id+'" name="var">'+v.name+'</span>';
//                                       $('#var_file').append(var_checkboxes);
//                                       break;
//                                     case "template":
//                                         var tem_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="'+v.id+'" name="template">'+v.name+'</span>';
//                                         $('#template_file').append(tem_checkboxes);
//                                       break;
//                                     case 'task':v
//                                         var task_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="'+v.id+'" name="task">'+v.name+'</span>';
//                                         $('#task_file').append(task_checkboxes);
//                                       break;
//                                 }
//                                 $('#myRoleModal').modal('show');
//                               }
//                           }else{
//                               pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+data.msg);
//                           }
//                           //处理模态框和表单
//                           $("#myRoleModal :input").each(function () {
//                               $(this).val("");
//                           });
//                           $("#myRoleModal").on("hidden.bs.modal", function() {
//                               $(this).removeData("bs.modal");
//                           });
//                       },
//                       error: function (){
//                           pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
//                       }
//                   });
//
//
//
//
//              });

        </script>
