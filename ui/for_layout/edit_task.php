        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">创建任务</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="template_id" class="col-sm-2 control-label">任务模板</label>
            <div class="col-sm-4">
              <select class="form-control" id="template_id" name="template_id" onchange="setTaskName()">
                <option value="">请选择</option>
              </select>
            </div>
            <label for="task_name" class="col-sm-2 control-label">任务名称</label>
            <div class="col-sm-4">
              <input type="text" class="form-control" id="task_name" name="task_name" onkeyup="check()" onchange="check()" placeholder="任务名称">
            </div>
          </div>
          <div class="form-group">
            <label for="max_num" class="col-sm-2 control-label">最大并发数</label>
            <div class="col-sm-4">
              <input type="text" class="form-control" id="max_num" name="max_num" onkeyup="check()" onchange="check()" value="1" placeholder="最大同时执行数">
            </div>
            <label for="max_ratio" class="col-sm-2 control-label">最大并发比例</label>
            <div class="col-sm-4">
              <input type="text" class="form-control" id="max_ratio" name="max_ratio" onkeyup="check()" onchange="check()" value="30" placeholder="最大同时执行比例">
            </div>
          </div>
          <div class="form-group">
            <label for="pool" class="col-sm-2 control-label">服务</label>
            <div class="col-sm-10" style="padding-bottom: 0px;">
              <div class="panel panel-default" style="padding: 15px 15px 0px;margin-bottom: 0px;">
                <div class="pannel-body">
                  <div class="col-sm-12" style="padding: 5px 0px 5px 0px;">
                    <div class="col-sm-4" style="padding: 0px 0px;">
                      <div class="input-group">
                        <span class="input-group-addon">集群</span>
                        <select class="form-control" name="cluster" id="cluster" onchange="getService()">
                          <option value="">选择集群</option>
                        </select>
                      </div>
                    </div>
                    <div class="col-sm-4" style="padding: 0px 0px;">
                      <div class="input-group">
                        <span class="input-group-addon">服务</span>
                        <select class="form-control" name="service" id="service" onchange="getPool()">
                          <option value="">请选择</option>
                        </select>
                      </div>
                    </div>
                    <div class="col-sm-4"style="padding: 0px 0px;">
                      <div class="input-group">
                        <span class="input-group-addon">服务池</span>
                        <select class="form-control" name="pool" id="pool" onchange="getNodes()">
                          <option value="">请选择</option>
                        </select>
                      </div>
                    </div>
                  </div>
                  <div class="col-sm-12" style="padding: 5px 0px 5px 0px;">
                    <div class="input-group" style="margin-bottom: 0px;">
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
            <label for="ip" class="col-sm-2 control-label">将执行IP<br>数量：<span id="run_num" class="badge bg-red">0</span></label>
            <div class="col-sm-10">
              <textarea rows="3" class="form-control" id="ip" name="ip" placeholder="支持逗号,分号,空格,冒号,换行" onkeyup="manualIp()"></textarea>
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="insert" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          getTaskTpl();
          getCluster();
        </script>
