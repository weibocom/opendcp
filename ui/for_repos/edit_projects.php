        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">新建项目</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="project_name" class="col-sm-2 control-label">项目名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="project_name" name="project_name" onkeyup="check()" onchange="check()" placeholder="名称,eg:测试">
            </div>
          </div>
          <div class="form-group">
            <label for="public" class="col-sm-2 control-label">是否公共项目</label>
            <div class="col-sm-10">
              <select class="form-control" id="public" name="public" onchange="check()">
                <option value="false">False</option>
                <option value="true">True</option>
              </select>
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="insert" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
