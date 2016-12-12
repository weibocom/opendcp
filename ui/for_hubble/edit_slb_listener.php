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
$myProtocol=(isset($_GET['protocol'])&&!empty($_GET['protocol']))?trim($_GET['protocol']):'';
$myPort=(isset($_GET['port'])&&!empty($_GET['port']))?trim($_GET['port']):'';
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myStepModalLabel">添加Listener</h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myStepModalBody">
          <div id="wizard" class="form_wizard wizard_horizontal">
            <ul class="wizard_steps anchor">
              <li>
                <a href="#step-1" class="done" isdone="1" rel="1">
                  <span class="step_no">1</span>
                  <span class="step_descr">基础配置</span>
                </a>
              </li>
              <li>
                <a href="#step-2" class="done" isdone="1" rel="2">
                  <span class="step_no">2</span>
                  <span class="step_descr">来访追踪配置</span>
                </a>
              </li>
              <li>
                <a href="#step-3" class="done" isdone="1" rel="3">
                  <span class="step_no">3</span>
                  <span class="step_descr">健康检查配置</span>
                </a>
              </li>
              <li>
                <a href="#step-4" class="selected" isdone="1" rel="4">
                  <span class="step_no">4</span>
                  <span class="step_descr">回话保持设置</span>
                </a>
              </li>
            </ul>
            <div class="stepContainer" style="overflow: auto;">
              <div id="step-1" class="content" style="display: block;">
                <div class="form-group">
                  <label for="Protocol" class="col-sm-2 control-label">协议类型</label>
                  <div class="col-sm-10">
                    <select class="form-control" id="Protocol" name="Protocol" onchange="checkProtocol()">
                      <option value="http">HTTP</option>
                      <option value="https">HTTPS</option>
                      <option value="tcp">TCP</option>
                      <option value="udp">UDP</option>
                    </select>
                  </div>
                </div>
                <div class="form-group hidden">
                  <label for="ServerCertificateId" class="col-sm-2 control-label">安全证书</label>
                  <div class="col-sm-10">
                    <select class="form-control" id="ServerCertificateId" name="ServerCertificateId" onchange="check()">
                      <option value="">请选择</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label for="Scheduler" class="col-sm-2 control-label">调度算法</label>
                  <div class="col-sm-10">
                    <select class="form-control" id="Scheduler" name="Scheduler" onchange="check()">
                      <option value="wrr">轮询模式</option>
                      <option value="wlc">最小连接数</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label for="ListenerPort" class="col-sm-2 control-label">服务端口</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="ListenerPort" name="ListenerPort" onkeyup="check()" placeholder="端口取值: 1-65535">
                  </div>
                </div>
                <div class="form-group">
                  <label for="BackendServerPort" class="col-sm-2 control-label">后端端口</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="BackendServerPort" name="BackendServerPort" onkeyup="check()" placeholder="端口取值: 1-65535">
                    <span class="help-block"> 创建以后不可修改</span>
                  </div>
                </div>
                <div class="form-group hidden">
                  <label for="PersistenceTimeout" class="col-sm-2 control-label">超时时间</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="PersistenceTimeout" name="PersistenceTimeout" onkeyup="check()" placeholder="连接持久化超时时间(0关闭,1-1000)">
                  </div>
                </div>
                <div class="form-group">
                  <label for="Bandwidth" class="col-sm-2 control-label">带宽峰值</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="Bandwidth" name="Bandwidth" onkeyup="check()" placeholder="取值:1-1000 | -1">
                  </div>
                </div>
              </div>

              <div id="step-2" class="content" style="display: none;">
                <div class="form-group">
                  <label for="XForwarderFor" class="col-sm-2 control-label">获取来访者IP *</label>
                  <div class="col-sm-10">
                    <select class="form-control" id="XForwarderFor" name="XForwarderFor" onChange="check()">
                      <option value="on">开启</option>
                      <option value="off">关闭</option>
                    </select>
                    <span class="help-block"> 是否开启获取来访者真实IP</span>
                  </div>
                </div>
              </div>

              <div id="step-3" class="content" style="display: none;">
                <div class="form-group">
                  <label for="HealthCheck" class="col-sm-2 control-label">健康检查 *</label>
                  <div class="col-sm-10">
                    <select class="form-control" id="HealthCheck" name="HealthCheck" onChange="checkHealthCheck()">
                      <option value="on">开启</option>
                      <option value="off">关闭</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label for="HealthCheckConnectPort" class="col-sm-2 control-label">检查端口 *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="HealthCheckConnectPort" name="HealthCheckConnectPort" onkeyup="check()" placeholder="取值：1-65535,或者-520(使用后端服务端口)">
                  </div>
                </div>
                <div class="form-group">
                  <label for="HealthCheckTimeout" class="col-sm-2 control-label">响应超时时间 *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="HealthCheckTimeout" name="HealthCheckTimeout" value="5" onkeyup="check()" placeholder="取值：1-50">
                  </div>
                </div>
                <div class="form-group">
                  <label for="HealthCheckInterval" class="col-sm-2 control-label">健康检查间隔 *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="HealthCheckInterval" name="HealthCheckInterval" value="2" onkeyup="check()" placeholder="取值：1-5">
                  </div>
                </div>
                <div class="form-group">
                  <label for="UnhealthyThreshold" class="col-sm-2 control-label">不健康阈值 *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="UnhealthyThreshold" name="UnhealthyThreshold" value="3" onkeyup="check()" placeholder="判定健康检查结果为fail的阈值。取值：1-10">
                    <span class="help-block">判定健康检查结果为fail的阈值.取值:1-10</span>
                  </div>
                </div>
                <div class="form-group">
                  <label for="HealthyThreshold" class="col-sm-2 control-label">健康阈值 *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="HealthyThreshold" name="HealthyThreshold" value="3" onkeyup="check()" placeholder="取值：1-10">
                    <span class="help-block">判定健康检查结果为success的阈值.取值:1-10</span>
                  </div>
                </div>
                <div class="form-group">
                  <label for="HealthCheckDomain" class="col-sm-2 control-label">检查域名</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="HealthCheckDomain" name="HealthCheckDomain" onkeyup="check()">
                    <span class="help-block">只能使用字母、数字、‘-’、‘.’，默认使用各后端服务器的内网IP为域名</span>
                  </div>
                </div>
                <div class="form-group">
                  <label for="HealthCheckURI" class="col-sm-2 control-label">检查路径</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="HealthCheckURI" name="HealthCheckURI" onkeyup="check()" value="/">
                    <span class="help-block">用于健康检查页面文件的URI</span>
                  </div>
                </div>
                <div class="form-group">
                  <label for="HealthCheckHttpCode" class="col-sm-2 control-label">http状态码 *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="HealthCheckHttpCode" name="HealthCheckHttpCode" onkeyup="check()" value="http_2xx" placeholder="健康检查正常的http状态码">
                    <span class="help-block"> 健康检查正常的http状态码</span>
                  </div>
                </div>
              </div>

              <div id="step-4" class="content" style="display: none;">
                <div class="form-group">
                  <label for="StickySession" class="col-sm-2 control-label">会话保持 *</label>
                  <div class="col-sm-10">
                    <select class="form-control" id="StickySession" name="StickySession" onChange="checkStickySession()">
                      <option value="off" selected>关闭</option>
                      <option value="on">开启</option>
                    </select>
                    <span class="help-block"> 仅适用于HTTP/HTTPS协议</span>
                  </div>
                </div>
                <div class="form-group hidden">
                  <label for="StickySessionType" class="col-sm-2 control-label">会话保持类型 *</label>
                  <div class="col-sm-10">
                    <select class="form-control" id="StickySessionType" name="StickySessionType" onChange="checkStickySessionType()">
                      <option value="insert">insert (由负载均衡插入)</option>
                      <option value="server">server (从后端学习)</option>
                    </select>
                  </div>
                </div>
                <div class="form-group hidden">
                  <label for="CookieTimeout" class="col-sm-2 control-label">cookie超时时间 *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="CookieTimeout" name="CookieTimeout" onkeyup="check()" value="10" placeholder="取值：1-86400">
                  </div>
                </div>
                <div class="form-group hidden">
                  <label for="Cookie" class="col-sm-2 control-label">cookie *</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" id="Cookie" name="Cookie" onkeyup="check()" placeholder="服务器上配置的cookie">
                  </div>
                </div>
              </div>
            </div>
            <input type="hidden" id="LoadBalancerId" name="LoadBalancerId" value="">
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

            <?php if($myProtocol&&$myPort) echo 'get("'.$myProtocol.'","'.$myPort.'");'; ?>

          });
        </script>
