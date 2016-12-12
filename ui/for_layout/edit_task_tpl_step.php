<?php
/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
$myIdx=(isset($_GET['idx']))?trim($_GET['idx']):'';
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myChildModalLabel">选择步骤命令</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myChildModalBody">
          <div class="form-group">
            <label for="step_name" class="col-sm-2 control-label">名称</label>
            <div class="col-sm-10">
              <select class="form-control" id="step_name" onchange="listActionParams()">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="action_param" class="col-sm-2 control-label">参数</label>
            <div class="col-sm-10 profile_details">
              <div class="well profile_view col-sm-12">
                <div class="col-sm-12">
                  <table class="table table-striped table-hover">
                    <thead>
                    <tr>
                      <th>参数名</th>
                      <th>值</th>
                    </tr>
                    </thead>
                    <tbody id="step_param">
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
          <div class="form-group">
            <label for="retry_times" class="col-sm-2 control-label">重试次数</label>
            <div class="col-sm-10">
              <input class="form-control" id="retry_times" onkeyup="check('step_param')" onchange="check('step_param')" placeholder="重试次数,eg:1" value="0">
            </div>
          </div>
          <div class="form-group">
            <label for="ignore_error" class="col-sm-2 control-label">忽略错误</label>
            <div class="col-sm-10">
              <select class="form-control" id="ignore_error" onchange="check('step_param')">
                <option value="false">否</option>
                <option value="true">是</option>
              </select>
            </div>
          </div>
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommitAction" onclick="setAction('<?php echo $myIdx;?>')" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          getStep('<?php echo $myIdx;?>');
        </script>
