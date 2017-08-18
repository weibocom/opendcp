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

?>

        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myRoleModalLabel">添加</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myRoleModalBody">
            <div class="form-group">
                <label for="desc" class="col-sm-2 control-label">Role名称</label>
                <div class="col-sm-10">
                    <input type="text" class="form-control" id="name" name="name"  placeholder="role名称">
                </div>
            </div>
            <div class="form-group" >
                <label for="args" class="col-sm-2 control-label">var文件</label>
                <div id="var_file">
                </div>
            </div>
            <div class="form-group">
                <label for="args" class="col-sm-2 control-label">template文件</label>
                <div  id="template_file">
                </div>
            </div>
            <div class="form-group">
                <label for="args" class="col-sm-2 control-label">task文件</label>
                <div id="task_file">
                </div>
            </div>
            <div class="form-group">
                <label for="args" class="col-sm-2 control-label">role描述</label>
                <div class="col-sm-10">
                    <input type="text" class="form-control" id="desc" name="desc" placeholder="描述,eg:测试">
                </div>
            </div>
        </div>

        <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-success" id="btnCommit" data-dismiss="modal" onclick="creatRole()" style="margin-bottom: 5px;">提交</button>
        </div>

