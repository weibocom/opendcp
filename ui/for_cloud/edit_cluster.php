<div class="modal-header">
    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
    <h4 class="modal-title" id="myStepModalLabel">创建机型模板</h4>
</div>
<div class="modal-body" style="overflow:auto;" id="myStepModalBody">
    <div id="wizard" class="form_wizard wizard_horizontal">
        <ul class="wizard_steps anchor">
            <li>
                <a href="#step-1" class="done" isdone="1" rel="1">
                    <span class="step_no">1</span>
                    <span class="step_descr">云厂商</span>
                </a>
            </li>
            <li>
                <a href="#step-2" class="done" isdone="1" rel="2">
                    <span class="step_no">2</span>
                    <span class="step_descr">网络配置</span>
                </a>
            </li>
            <li>
                <a href="#step-3" class="done" isdone="1" rel="3">
                    <span class="step_no">3</span>
                    <span class="step_descr">机器规格</span>
                </a>
            </li>
            <li>
                <a href="#step-4" class="selected" isdone="1" rel="4">
                    <span class="step_no">4</span>
                    <span class="step_descr">磁盘选项</span>
                </a>
            </li>
        </ul>
        <div class="stepContainer" style="overflow: auto;">
            <div id="step-1" class="content" style="display: block;">
                <div class="form-group">
                    <label for="Name" class="col-sm-2 control-label">名称</label>
                    <div class="col-sm-10">
                        <input type="text" class="form-control" id="Name" name="Name" onkeypress="check()" onchange="check()" placeholder="名称,eg:测试">
                    </div>
                </div>
                <div class="form-group">
                    <label for="Provider" class="col-sm-2 control-label">云厂商</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="Provider" name="Provider" onchange="selectProvider()">
                            <option value="default">请选择</option>
                        </select>

                    </div>
                </div>
            </div>

            <div id="step-2" class="content" style="display: none;">
                <div class="form-group">
                    <label for="NetworkType" class="col-sm-2 control-label">网络类型</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="NetworkType" name="NetworkType" onchange="getNetworkType()">
                            <option value="common">经典网络</option>
                            <option value="custom">专有网络</option>
                        </select>
                    </div>
                </div>
                <div class="form-group hidden">
                    <label for="NetworkOP" class="col-sm-2 control-label">网络</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="NetworkOP" name="NetworkOP" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group hidden">
                    <label for="AvabilityZone" class="col-sm-2 control-label">可用域</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="AvalibilityZone" name="AvalibilityZone" onchange="check()" >
                            <option value="">请选择</option>
                            <option value="nova">nova</option>
                        </select>
                    </div>
                </div>
                <div class="form-group hidden">
                    <label for="InternetChargeType" class="col-sm-2 control-label">经典网络-计费类型</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="InternetChargeType" name="InternetChargeType" onchange="check()">
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group hidden">
                    <label for="InternetMaxBandwidthOut" class="col-sm-2 control-label">经典网络-公网出带宽峰值</label>
                    <div class="col-sm-10">
                        <input type="number" min="0" max="100" value="1" class="form-control" id="InternetMaxBandwidthOut" name="InternetMaxBandwidthOut" onkeyup="check()" placeholder="单位: Mbps">
                        <span class="help-block"> 单位Mpbs, s按带宽付费取值1-100, 按流量付费取值0-100</span>
                    </div>
                </div>
                <div class="form-group">
                    <label for="RegionName" class="col-sm-2 control-label">可用地域</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="RegionName" name="RegionName" onchange="getZoneAndVpcAndETypeAndImage()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label for="ZoneName" class="col-sm-2 control-label">可用区</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="ZoneName" name="ZoneName" onchange="getVSwitchAndSecurityGroup()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label for="VpcId" class="col-sm-2 control-label">专有网络-VPC</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="VpcId" name="VpcId" onchange="getVSwitchAndSecurityGroup()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label for="SubnetId" class="col-sm-2 control-label">专有网络-子网</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="SubnetId" name="SubnetId" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label for="SecurityGroupId" class="col-sm-2 control-label">安全组</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="SecurityGroupId" name="SecurityGroupId" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
            </div>

            <div id="step-3" class="content" style="display: none;">
                <div class="form-group">
                    <label for="InstanceType" class="col-sm-2 control-label">机器规格</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="InstanceType" name="InstanceType" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label for="ImageId" class="col-sm-2 control-label">镜像</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="ImageId" name="ImageId" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
            </div>

            <div id="step-4" class="content" style="display: none;">
                <div class="form-group">
                    <label for="SystemDiskCategory" class="col-sm-2 control-label">系统盘类型</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="SystemDiskCategory" name="SystemDiskCategory" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group hidden">
                    <label for="DiskType" class="col-sm-2 control-label">磁盘规格</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="DiskType" name="DiskType" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>

                <div class="form-group">
                    <label for="DataDiskCategory" class="col-sm-2 control-label">数据盘类型</label>
                    <div class="col-sm-10">
                        <select class="form-control" id="DataDiskCategory" name="DataDiskCategory" onchange="check()" disabled>
                            <option value="">请选择</option>
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label for="DataDiskSize" class="col-sm-2 control-label">数据盘大小</label>
                    <div class="col-sm-10">
                        <input type="number" class="form-control" id="DataDiskSize" name="DataDiskSize" onkeyup="isNumberValid(this,1,32768,100)" value="100" max="32768" min="1" placeholder="数据盘大小,单位G,eg:100">
                        <span class="help-block">单位G, 范围5-32768, 请参阅官方手册</span>
                    </div>
                </div>

                <div class="form-group">
                    <label for="DataDiskNum" class="col-sm-2 control-label">数据盘数量</label>
                    <div class="col-sm-10">
                        <input type="number" class="form-control" id="DataDiskNum" name="DataDiskNum" onkeyup="isNumberValid(this,1,4,0)" value="1" max="4" min="0" placeholder="数据盘数量,eg:1">
                        <span class="help-block">单位G, 范围1-4, 请参阅官方手册</span>
                    </div>
                </div>
            </div>
        </div>
        <input type="hidden" id="page_action" name="page_action" value="insert" />
    </div>
</div>
<script>
    $(document).ready(function() {
        $('#wizard').smartWizard();
        $('#wizard_verticle').smartWizard({
            transitionEffect: 'slide'
        });

        $('.buttonNext').addClass('btn btn-success');
        $('.buttonPrevious').addClass('btn btn-primary');
        $('.buttonFinish').addClass('btn btn-info');

        $('#LoadBalancerId').val($('#fSlb').val());

         getProvider();

    });
</script>
