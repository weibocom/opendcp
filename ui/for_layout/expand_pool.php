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
  $myId=(isset($_GET['id'])&&!empty($_GET['id']))?trim($_GET['id']):'';
  $myName=(isset($_GET['name'])&&!empty($_GET['name']))?trim($_GET['name']):'';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">服务池扩容</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="pool" class="col-sm-2 control-label">服务池</label>
            <div class="col-sm-10">
              <select class="form-control" id="pool" name="pool" onchange="check('expand')" readonly>
                <option value="<?php echo $myId;?>"><?php echo $myName;?></option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="num" class="col-sm-2 control-label">扩容数量</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="num" name="num" onkeyup="check('expand')" placeholder="扩容数量">
              <span class="help-block">当前配额余量: <span class="badge bg-red" id="quota">0</span></span>
            </div>
          </div>
          <input type="hidden" id="template_id" name="template_id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="page_action" name="page_action" value="expand" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          getQuota();
        </script>
