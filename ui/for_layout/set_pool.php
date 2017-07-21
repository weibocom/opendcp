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
$myTitle='任务设置';
$pageAction='setpool';
?>
<div class="modal-header">
    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
    <h4 class="modal-title" id="myModalLabel"><?php echo $myTitle;?></h4>
</div>
<div class="modal-body" style="overflow:auto;" id="myModalBody">
    <div class = "x_panel">
        <div class="x_title">
            <h2 class="text-primary">  <i class="fa fa-bank"> 任务选择 </i> </h2>
            <div class="clearfix"></div>
        </div>
        <div class="x_content">
            <div class="row">
                <div class="col-sm-7 form-group">
                    <div class="btn-group">
                        <div class="input-group col-sm-5" style="padding-left:0px;">
                            <span class="input-group-addon">任务类型</span>
                            <select class="form-control" id="task_type" name="service_id" onchange="listCronOrDepen(<?php echo $myIdx;?>)">
                                <option value="expandList">扩容</option>
                                <option value="uploadList">上线</option>
                            </select>
                        </div>
                    </div>
                </div>
                <div class="col-md-5">
                    <div class="btn-group pull-right">
                        <span id="addCron" class="btn btn-success" style="border-radius: 3px;margin-right:5px;" onclick="addTaskCron()"><i class="fa fa-plus"></i> 定时任务</span>
                        <span id="addDepen" class="btn btn-success" style="border-radius: 3px;margin-left:5px" onclick="addTaskDepen()"><i class="fa fa-plus"></i> 依赖任务</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class = "x_panel">
        <div class="x_title">
            <h2 class="text-primary"><i class="fa fa-retweet"> 定时任务</i> </h2>
            <div class="clearfix"></div>
        </div>
        <div class="x_content">
            <div class="table-scrollable">
                <table class="table table-bordered table-hover" id="cron_table">
                    <thead id="cron_head">
                    </thead>
                    <tbody id="cron_body">
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    <div class = "x_panel">
        <div class="x_title">
            <h2 class="text-primary">  <i class="fa fa-caret-square-o-up"></i> 依赖任务</h2>
            <div class="clearfix"></div>
        </div>
        <div class="x_content">
            <div class="table-scrollable">
                <table class="table table-bordered table-hover" id="depen_table">
                    <thead id="depen_head">
                    <tr>
                        <th>#</th>
                        <th>依赖服务</th>
                        <th>依赖步骤</th>
                        <th>比例</th>
                        <th>机器冗余数量</th>
                        <th>忽略</th>
                        <th>#</th>
                    </tr>
                    </thead>
                    <tbody id="depend_body">
                    </tbody>
                </table>
            </div>
        </div>
   </div>
</div>

<div class="modal-footer">
    <button class="btn btn-default" data-dismiss="modal">取消</button>
    <button class="btn btn-success" data-dismiss="modal" id="btnSaveTask" onclick="saveCronAndDependTask()" style="margin-bottom: 5px;" disabled>保存</button>
</div>
<script>
    listCronOrDepen(<?php echo $myIdx;?>);
</script>
