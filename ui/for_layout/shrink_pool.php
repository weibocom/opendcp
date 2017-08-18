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


  $myId=(isset($_GET['id'])&&!empty($_GET['id']))?trim($_GET['id']):'';
  $myName=(isset($_GET['name'])&&!empty($_GET['name']))?trim($_GET['name']):'';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">服务池缩容</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="pool" class="col-sm-2 control-label">服务池</label>
            <div class="col-sm-10">
              <select class="form-control" id="pool" name="pool" onchange="check('shrink')" readonly>
                <option value="<?php echo $myId;?>"><?php echo $myName;?></option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="pool" class="col-sm-2 control-label">服务池</label>
            <div class="col-sm-10" style="padding-bottom: 0px;">
              <div class="panel panel-default" style="padding: 15px 15px 0px;margin-bottom: 0px;">
                <div class="pannel-body">
                  <div style="padding-top: 5px;padding-bottom: 5px;">
                    <div class="input-group">
                      <span class="input-group-btn"><button type="button" class="btn btn-default" disabled>IP段</button></span>
                      <input type="text" id="check_input" placeholder="请输入需过滤的IP段" class="form-control col-sm-4" style="width:30%">
                      <input type="button" value="选定指定IP" class="btn btn-default" onclick="autoCheck(true)" style="margin-left: 20px;">
                      <input type="button" value="取消指定IP" class="btn btn-default" onclick="autoCheck(false)">
                      <span class="btn btn-default">当前选择数量：<span id="check_num">0</span></span>
                    </div>
                  </div>
                  <label><input type="checkbox" name="check_all" id="check_all" onchange="checkAll(this)"><span id="ipnum">全选</span></label>
                  <div id="iplist" name="iplist" style="overflow:auto;"></div>
                </div>
              </div>
            </div>
          </div>
          <div class="form-group">
            <label for="ip" class="col-sm-2 control-label">将缩容IP<br>数量：<span id="run_num" class="badge bg-red">0</span></label>
            <div class="col-sm-10">
              <textarea rows="3" class="form-control" id="ip" name="ip" placeholder="支持逗号,分号,空格,冒号,换行" onkeyup="manualIp()"></textarea>
            </div>
          </div>
          <input type="hidden" id="template_id" name="template_id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="page_action" name="page_action" value="shrink" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          cache.ip=[];
          getNodes();
        </script>
