<div class="modal-header">
    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span>
    </button>
    <h4 class="modal-title" id="myModalLabel">创建物理机器</h4>
</div>

<div class="modal-body" style="overflow:auto;" id="myModalBody">
    <div class="form-group">
        <div class="col-md-8 col-md-offset-2">
            <label>输入列表</label>
            <textarea class="form-control" id="instancelist" name="InstanceList" onkeyup="checkPhyDev()" rows="3"
                      placeholder="每行一个，按照 'PublicIP,PrivateIP,Password,Port(可选)' 格式输入"></textarea>
        </div>

    </div>


    <!--
    <div class="form-group">
        <label for="DataDiskSize" class="col-sm-2 control-label">DataDiskSize</label>
        <div class="col-sm-4">
            <input type="text" class="form-control" id="DataDiskSize" name="DataDiskSize" onkeyup="checkPhyDev()" placeholder="DataDiskSize">
        </div>
        <label for="DataDiskNum" class="col-sm-2 control-label">DataDiskNum</label>
        <div class="col-sm-4">
            <input type="text" class="form-control" id="DataDiskNum" name="DataDiskNum" onkeyup="checkPhyDev()"  placeholder="DataDiskNum">
        </div>
    </div>
    <div class="form-group">
        <label for="DataDiskCategory" class="col-sm-2 control-label">DataDiskCategory</label>
        <div class="col-sm-4">
            <input type="text" class="form-control" id="DataDiskCategory" name="DataDiskCategory" onkeyup="checkPhyDev()" placeholder="DataDiskCategory">
        </div>
        <label for="SystemDiskCategory" class="col-sm-2 control-label">SystemDiskCategory</label>
        <div class="col-sm-4">
            <input type="text" class="form-control" id="SystemDiskCategory" name="SystemDiskCategory" onkeyup="checkPhyDev()" placeholder="SystemDiskCategory">
        </div>
    </div>
    <div class="form-group">
        <label for="PrivateIpAddress" class="col-sm-2 control-label">PrivateIpAddress</label>
        <div class="col-sm-4">
            <input type="text" class="form-control" id="PrivateIpAddress" name="PrivateIpAddress" onkeyup="checkPhyDev()" placeholder="PrivateIpAddress">
        </div>

        <label for="NatIpAddress" class="col-sm-2 control-label">NatIpAddress</label>
        <div class="col-sm-4">
            <input type="text" class="form-control" id="NatIpAddress" name="NatIpAddress" onkeyup="checkPhyDev()" placeholder="NatIpAddress">
        </div>
    </div>
    -->


    <input type="hidden" id="page_action" name="page_action" value="addPhyDev">
</div>

<div class="modal-footer">
    <button class="btn btn-default" data-dismiss="modal">取消</button>
    <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;"
            disabled>确认
    </button>
</div>
