        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">添加SLB</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="LoadBalancerName" class="col-sm-2 control-label">名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="LoadBalancerName" name="LoadBalancerName" onkeyup="check()" onchange="check()" placeholder="SLB名称,eg:for_test">
            </div>
          </div>
          <div class="form-group">
            <label for="NetworkType" class="col-sm-2 control-label">网络类型</label>
            <div class="col-sm-10">
              <select class="form-control" id="NetworkType" name="NetworkType" onchange="getNetworkType()">
                <option value="common">经典网络</option>
                <option value="custom">专有网络</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="AddressType" class="col-sm-2 control-label">地址类型</label>
            <div class="col-sm-10">
              <select class="form-control" id="AddressType" name="AddressType" onchange="check()">
                <option value="internet">外网</option>
                <option value="intranet" disabled>内网 - 只适用于专用网络</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="InternetChargeType" class="col-sm-2 control-label">付费方式</label>
            <div class="col-sm-10">
              <select class="form-control" id="InternetChargeType" name="InternetChargeType" onchange="check()">
                <option value="">请选择</option>
                <option value="paybytraffic">paybytraffic - 按流量付费</option>
                <option value="paybybandwidth">paybybandwidth - 按带宽付费</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="Bandwidth" class="col-sm-2 control-label">带宽峰值</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="Bandwidth" name="Bandwidth" onkeyup="check()" placeholder="取值：1-1000（单位为Mbps）">
            </div>
          </div>
          <div class="form-group">
            <label for="RegionId" class="col-sm-2 control-label">可用地域</label>
            <div class="col-sm-10">
              <select class="form-control" id="RegionId" name="RegionId" onchange="getZoneAndVpc()">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group hidden">
            <label for="ZoneId" class="col-sm-2 control-label">可用区</label>
            <div class="col-sm-10">
              <select class="form-control" id="ZoneId" name="ZoneId" onchange="getVSwitchAndSecurityGroup()" disabled>
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group hidden">
            <label for="VpcId" class="col-sm-2 control-label">VPC</label>
            <div class="col-sm-10">
              <select class="form-control" id="VpcId" name="VpcId" onchange="getVSwitchAndSecurityGroup()" disabled>
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group hidden">
            <label for="VSwitchId" class="col-sm-2 control-label">子网</label>
            <div class="col-sm-10">
              <select class="form-control" id="VSwitchId" name="VSwitchId" onchange="check()" disabled>
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group hidden">
            <label for="SecurityGroup" class="col-sm-2 control-label">安全组</label>
            <div class="col-sm-10">
              <select class="form-control" id="SecurityGroup" name="SecurityGroup" onchange="check()" disabled>
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <input type="hidden" id="page_action" name="page_action" value="insert" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          updateSelect('RegionId');
        </script>
