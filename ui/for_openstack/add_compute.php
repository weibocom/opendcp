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


$myAction = (isset($_GET['action']) && !empty($_GET['action'])) ? trim($_GET['action']) : 'add';
$myIdx = (isset($_GET['idx']) && !empty($_GET['idx'])) ? trim($_GET['idx']) : '';
$pageAction = 'addcompute';
$ip = empty($_REQUEST['ip']) ? '' : $_REQUEST['ip'];
$disk_name = empty($_REQUEST['disk_name']) ? '' : $_REQUEST['disk_name'];
$type = empty($_REQUEST['type']) ? 1 : $_REQUEST['type'];
//type为1是计算节点，2是控制节点，3是存储节点
$myTitle = ($type == 1 ? '初始化计算节点' : ($type == 3 ? '初始化存储节点' : '初始化控制节点'));
?>
<div class="modal-header">
    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span>
    </button>
    <h4 class="modal-title" id="myModalLabel"><?php echo $myTitle; ?></h4>
</div>
<div class="modal-body" style="overflow:auto;" id="myModalBody">
    <div class="form-group">
        <label for="name" class="col-sm-2 control-label">ip</label>
        <div class="col-sm-10">
            <input type="text" class="form-control" id="compute_ip" name="compute_ip" placeholder="节点的ip"
                   value="<?= $ip ?>">
        </div>
    </div>
    <?php if (empty($ip)) { ?>
        <div class="form-group">
            <label for="name" class="col-sm-2 control-label">登录root密码</label>
            <div class="col-sm-10">
                <input type="text" class="form-control" id="compute_password" name="compute_password"
                       placeholder="请提供root密码">
            </div>
        </div>
    <?php } ?>
    <?php if ($type == 3) { ?>
        <div class="form-group">
            <label for="name" class="col-sm-2 control-label">输入挂载磁盘</label>
            <div class="col-sm-10">
                <input type="text" class="form-control" id="disk_name" name="disk_name"
                       placeholder="请输入挂载磁盘，例如：sdb">
            </div>
        </div>
    <?php } ?>
    <input type="hidden" id="id" name="id" value="<?php echo $myIdx; ?>"/>
    <input type="hidden" id="type" name="type" value="<?= $type ?>"/>
    <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction; ?>"/>
</div>
<div class="modal-footer">
    <button class="btn btn-default" data-dismiss="modal">取消</button>
    <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="addcompute()"
            style="margin-bottom: 5px;">确认
    </button>
</div>
<script>

    function addcompute() {
        url = '/api/for_openstack/machine.php';
        ip = $('#compute_ip').val();
        type = $('#type').val();
        disk_name = ($('#disk_name').val() == null) ? '' : $('#disk_name').val();
        password = $('#compute_password').val();
        postData = {"action": "addip", "ip": ip, "password": password, "type": type, "disk_name" : disk_name,};
        $.ajax({
            type: "POST",
            url: url,
            data: postData,
            dataType: "json",
            success: function (listdata) {
                if (listdata.code == 0) {
                    location.href = '/for_openstack/initlist.php';
                    return;
                } else {
                    pageNotify('error', '加载失败！', '错误信息：' + listdata.msg);
                    NProgress.done();
                }
            },
            error: function () {
                pageNotify('error', '加载失败！', '错误信息：接口不可用');
                NProgress.done();
            }
        });
    }


</script>
