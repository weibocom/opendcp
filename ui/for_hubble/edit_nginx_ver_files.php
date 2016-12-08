        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myChildModalLabel">选择文件版本</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myChildModalBody">
          <div class="form-group">
            <label for="file" class="col-sm-2 control-label">文件</label>
            <div class="col-sm-10">
              <select class="form-control" id="file" onchange="listFileVer()">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="file_ver" class="col-sm-2 control-label">版本</label>
            <div class="col-sm-10 profile_details">
              <div class="well profile_view col-sm-12">
                <div class="col-sm-12">
                  <table class="table table-striped table-hover">
                    <thead>
                    <tr>
                      <th>#</th>
                      <th>文件</th>
                      <th>版本</th>
                      <th>状态</th>
                      <th>用户</th>
                      <th>时间</th>
                    </tr>
                    </thead>
                    <tbody id="file_ver">
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommitFile" onclick="setFile()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          getFileList();
        </script>
