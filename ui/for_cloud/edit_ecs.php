        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">创建机器</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="ClusterId" class="col-sm-2 control-label">机型模板</label>
            <div class="col-sm-10">
              <select class="form-control" id="ClusterId" name="ClusterId" onchange="check()">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="Number" class="col-sm-2 control-label">数量</label>
            <div class="col-sm-10">
              <input type="number" class="form-control" id="Number" name="Number" onkeyup="isNumberValid(this,1,1000,1)" onchange="isNumberValid(this,1,1000,1)" value="1" max="1000" min="1" placeholder="数量,eg:1">
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="insert" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          getCluster();
        </script>
