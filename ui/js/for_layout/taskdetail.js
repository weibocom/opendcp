cache = {
    task_id: 0,
    task: {},
    nodes: {},
    state: {},
    flag: true, //是否自动更新
    waiting: {  //上次加载是否完成
        info: true,
        statistics: true,
        state: true,
    },
    listcount: {
        ready: 20,
        success: 20,
        running: 20,
    },
    ip: {
        ready: [],
        running: [],
        success: [],
        failed: [],
        stoped: []
    },
    arg_name: [],
    tasklist: [],
    result: {},
    refreshInterval: null,

}

var getDate = function (t, type) {
    if (!t) t = '';
    var d = new Date(t);
    var M = (d.getMonth() + 1);
    var D = d.getDate();
    var h = d.getHours();
    var i = d.getMinutes();
    var s = d.getSeconds();
    var ret = '';
    switch (type) {
        case 'time':
            ret = ((h < 10) ? '0' + h : h) + ':' + ((i < 10) ? '0' + i : i) + ':' + ((s < 10) ? '0' + s : s);
            break;
        default:
            ret = d.getFullYear() + '.' + ((M < 10) ? '0' + M : M) + '.' + ((D < 10) ? '0' + D : D) + ' ' + ((h < 10) ? '0' + h : h) + ':' + ((i < 10) ? '0' + i : i) + ':' + ((s < 10) ? '0' + s : s);
            break;
    }
    return ret;
}

//changeTaskId -- 未启用
var setId = function (action) {
    return false;
    switch (action) {
        case 'prev':
            cache.task_id--;
            break;
        case 'next':
            cache.task_id++;
            break;
    }
    getTask('info');
    cache.flag = true;
}

//获取任务信息
var getTask = function (action) {
    if (cache.flag == false && action != 'state') return false;
    //NProgress.start();
    var actionDesc = '', postData = {};
    switch (action) {
        case 'info':
            actionDesc = '获取任务详情';
            postData = {"action": action, "fIdx": cache.task_id};
            if (cache.waiting.info == false) return false;
            cache.waiting.info = false;
            pageNotify('info', actionDesc, '获取中...');
            break;
        case 'state':
            actionDesc = '获取任务节点';
            var pagesize = 2000;
            postData = {"action": "nodes", "fIdx": cache.task_id, "pagesize": pagesize};
            break;
    }
    if (!cache.task_id) {
        pageNotify('error', actionDesc + '失败！', '错误信息：错误的任务ID');
        //NProgress.done();
        return false;
    }
    $.ajax({
        type: "POST",
        url: '/api/for_layout/task.php?action=' + action,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (data.code == 0) {
                if (!$.isEmptyObject(data.content)) {
                    switch (action) {
                        case 'info':
                            cache.task = data.content;
                            updateEle('task');
                            updateEle('num');
                            updateEle('overview');
                            getTask('state');
                            updateEle('state', cache.task.state);
                            cache.waiting.info = true;
                            break;
                        case 'state':
                            cache.nodes = data.content;
                            cache.ip.ready = [];
                            cache.ip.running = [];
                            cache.ip.success = [];
                            cache.ip.failed = [];
                            cache.ip.stoped = [];
                            $.each(data.content, function (k, v) {
                                switch (k) {
                                    case 0:
                                    case '0':
                                        cache.ip.ready = v;
                                        break;
                                    case 1:
                                    case '1':
                                        cache.ip.running = v;
                                        break;
                                    case 2:
                                    case '2':
                                        cache.ip.success = v;
                                        break;
                                    case 3:
                                    case '3':
                                        cache.ip.failed = v;
                                        break;
                                    case '4':
                                        cache.ip.stoped = v;
                                        break;
                                }
                            });
                            updateEle('ip', 'ready');
                            updateEle('ip', 'running');
                            updateEle('ip', 'success');
                            updateEle('ip', 'failed');
                            updateEle('ip', 'stoped');
                            if (cache.ip.ready.length > 0) {
                                $('#tab_1').attr('class', 'tab-pane fade active in');
                                $('#tab_2').attr('class', 'tab-pane fade');
                                $('#tab_3').attr('class', 'tab-pane fade');
                                $('#tab_4').attr('class', 'tab-pane fade');
                                $('#tab_home_1').attr('class', 'active');
                                $('#tab_home_2').attr('class', '');
                                $('#tab_home_3').attr('class', '');
                                $('#tab_home_4').attr('class', '');
                                $('#tab_home_5').attr('class', '');
                            } else {
                                if (cache.ip.running.length > 0) {
                                    $('#tab_1').attr('class', 'tab-pane fade');
                                    $('#tab_2').attr('class', 'tab-pane fade active in');
                                    $('#tab_3').attr('class', 'tab-pane fade');
                                    $('#tab_4').attr('class', 'tab-pane fade');
                                    $('#tab_5').attr('class', 'tab-pane fade');
                                    $('#tab_home_1').attr('class', '');
                                    $('#tab_home_2').attr('class', 'active');
                                    $('#tab_home_3').attr('class', '');
                                    $('#tab_home_4').attr('class', '');
                                    $('#tab_home_5').attr('class', '');
                                } else {
                                    if (cache.ip.stoped.length > 0) {
                                        $('#tab_1').attr('class', 'tab-pane fade');
                                        $('#tab_2').attr('class', 'tab-pane fade');
                                        $('#tab_5').attr('class', 'tab-pane fade active in');
                                        $('#tab_3').attr('class', 'tab-pane fade');
                                        $('#tab_4').attr('class', 'tab-pane fade');
                                        $('#tab_home_1').attr('class', '');
                                        $('#tab_home_2').attr('class', '');
                                        $('#tab_home_5').attr('class', 'active');
                                        $('#tab_home_3').attr('class', '');
                                        $('#tab_home_4').attr('class', '');

                                    } else {
                                        if (cache.ip.success.length > 0) {
                                            $('#tab_1').attr('class', 'tab-pane fade');
                                            $('#tab_2').attr('class', 'tab-pane fade');
                                            $('#tab_3').attr('class', 'tab-pane fade active in');
                                            $('#tab_4').attr('class', 'tab-pane fade');
                                            $('#tab_5').attr('class', 'tab-pane fade');
                                            $('#tab_home_1').attr('class', '');
                                            $('#tab_home_2').attr('class', '');
                                            $('#tab_home_3').attr('class', 'active');
                                            $('#tab_home_4').attr('class', '');
                                            $('#tab_home_5').attr('class', '');
                                        } else {
                                            $('#tab_1').attr('class', 'tab-pane fade');
                                            $('#tab_2').attr('class', 'tab-pane fade');
                                            $('#tab_3').attr('class', 'tab-pane fade');
                                            $('#tab_4').attr('class', 'tab-pane fade active in');
                                            $('#tab_5').attr('class', 'tab-pane fade');
                                            $('#tab_home_1').attr('class', '');
                                            $('#tab_home_2').attr('class', '');
                                            $('#tab_home_3').attr('class', '');
                                            $('#tab_home_4').attr('class', 'active');
                                            $('#tab_home_5').attr('class', '');
                                        }
                                    }
                                }
                            }
                            break;
                    }
                } else {
                    pageNotify('info', actionDesc + '成功！', '数据为空！');
                }
            } else {
                pageNotify('error', actionDesc + '失败！', '错误信息：' + data.msg);
            }
            $('.tooltips').each(function () {
                $(this).tooltip();
            });
            //NProgress.done();
        },
        error: function () {
            pageNotify('error', actionDesc + '失败！', '错误信息：接口不可用');
            //NProgress.done();
        }
    });
}

//展示任务详情
var updateEle = function (o, idx) {
    $('.popovers').each(function () {
        $(this).popover('hide');
    });
    var t = '', data = '';
    switch (o) {
        case 'task':
            data = cache.task;
            if (data.task_name) $('#task_name').html(data.task_name);
            if (data.pool_name) $('#pool_name').html(data.pool_name);
            if (data.tag) $('#tag').html(data.tag);
            if (data.step_len) $('#step_len').html(data.step_len);
            if (data.created) $('#created').html(getDate(data.created));
            if (data.template_id) $('#template_id').html(data.template_id);
            if (data.updated) $('#updated').html(getDate(data.updated));
            if (data.start) $('#start').html(data.start);
            if (data.auto) $('#auto').html(data.auto);
            if (data.timeout) $('#t_timeout').val(data.timeout);
            if (data.num) $('#t_num').val(data.num);
            if (data.rate) $('#t_rate').val(data.rate);
            cache.task.arg = data.params;
            break;
        case 'state':
            switch (idx) {
                case 0:
                    $('#state').html('<span class="badge bg-default">未开始</span>');
                    break;
                case 1:
                    $('#state').html('<span class="badge bg-blue">执行中</span>');
                    break;
                case 2:
                    $('#state').html('<span class="badge bg-green">已完成</span>');
                    if (cache.refreshInterval != null) {
                        clearInterval(cache.refreshInterval);
                        cache.refreshInterval = null;
                    }
                    break;
                case 3:
                    $('#state').html('<span class="badge bg-red">失败</span>');
                    if (cache.refreshInterval != null) {
                        clearInterval(cache.refreshInterval);
                        cache.refreshInterval = null;
                    }
                    break;
                case 4:
                    $('#state').html('<span class="badge bg-orange">已暂停</span>');
                    if (cache.refreshInterval != null) {
                        clearInterval(cache.refreshInterval);
                        cache.refreshInterval = null;
                    }
                    break;
                default:
                    $('#state').html('<span class="badge bg-red" title="' + idx + '">未知</span>');
                    break;
            }
            break;
        case 'num':
            var data = cache.task.Stat;
            var ready = 0, success = 0, running = 0, failed = 0, stoped = 0;
            if (!$.isEmptyObject(data)) {
                $('#num_ready').html('');
                $('#num_success').html('');
                $('#num_running').html('');
                $('#num_failed').html('');
                $('#num_stoped').html('');
                ready = data[0];
                running = data[1];
                success = data[2];
                failed = data[3];
                stoped = data[4];
                if (ready > 0) {
                    $('#num_ready').html(ready);
                } else {
                    $('#task_ready').html('');
                }
                if (success > 0) {
                    $('#num_success').html(success);
                } else {
                    $('#task_success').html('');
                }
                if (running > 0) {
                    $('#num_running').html(running);
                } else {
                    $('#task_running').html('');
                }
                if (failed > 0) {
                    $('#num_failed').html(failed);
                } else {
                    $('#task_failed').html('');
                }
                if (stoped > 0) {
                    $('#num_stoped').html(stoped);
                } else {
                    $('#num_stoped').html('');
                }
            }
            break;
        case 'overview':
            var data = cache.task;
            if (!$.isEmptyObject(data)) {
                var body = $('#task_process');
                body.html('');
                var tr = $('<tr></tr>');
                td = '<td>' + cache.task.pool_name + '</td>';
                tr.append(td);
                var ready = (typeof cache.task.Stat != 'undefined') ? cache.task.Stat[0] : 0;
                var running = (typeof cache.task.Stat != 'undefined') ? cache.task.Stat[1] : 0;
                var success = (typeof cache.task.Stat != 'undefined') ? cache.task.Stat[2] : 0;
                var failed = (typeof cache.task.Stat != 'undefined') ? cache.task.Stat[3] : 0;
                var stoped = (typeof cache.task.Stat != 'undefined') ? cache.task.Stat[4] : 0;
                var count = ready + running + success + failed + stoped;
                td = '<td>' + count + '</td>';
                tr.append(td);
                td = '<td><span class="label label-default">' + ready + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-info">' + running + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-warning">' + stoped + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-success">' + success + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-danger">' + failed + '</span></td>';
                tr.append(td);
                var rate = (success + failed) * 100 / count;
                var tmp = rate.toString().substr(0, 5);
                if (failed > 0) {
                    td = '<td><div class="progress progress-striped active" style="margin-bottom: 0px;"><div class="progress-bar progress-bar-danger" role="progressbar" aria-valuemin="0" aria-valuemax="100" style="width: ' + tmp + '%"><span>' + tmp + '%</span></div></div></td>';
                } else {
                    if (tmp > 10) {
                        td = '<td><div class="progress progress-striped active" style="margin-bottom: 0px;"><div class="progress-bar progress-bar-success" role="progressbar" aria-valuemin="0" aria-valuemax="100" style="width: ' + tmp + '%"><span>' + tmp + '%</span></div></div></td>';
                    } else {
                        td = '<td><div class="progress progress-striped active" style="margin-bottom: 0px;"><div class="progress-bar progress-bar-success" role="progressbar" aria-valuemin="0" aria-valuemax="100" style="width: ' + tmp + '%"><span class="text-primary">' + tmp + '%</span></div></div></td>';
                    }
                }
                tr.append(td);
                body.append(tr);
            }
            break;
        case 'ip':
            $("[id=SelectAll]:checkbox").each(function(i){
                 $(this).attr("checked",false);
             });
            var body = $('#task_' + idx);
            switch (idx) {
                case 'ready':
                    data = cache.ip.ready;
                    break;
                case 'running':
                    data = cache.ip.running;
                    break;
                case 'success':
                    data = cache.ip.success;
                    break;
                case 'failed':
                    data = cache.ip.failed;
                    break;
                case 'stoped':
                    data = cache.ip.stoped;
                    break;
            }
            var n = 1;
            body.html('');
            if (!$.isEmptyObject(data)) {
                for (var i = 0; i < data.length; i++) {
                    var v = data[i];
                    if (n > cache.listcount) break;
                    var tr = $('<tr></tr>');
                    td = '<td><input type="checkbox" id="list_' + idx + '" name="list_' + idx + '[]" value="' + v.id + '_' + v.ip +'_'+v.state+ '"></td>';
                    tr.append(td);
                    td = '<td>' + n + '</td>';
                    tr.append(td);
                    var err = (v.error) ? insertStr(v.error, ' ', 40) : v.error;
                    var warn = (v.error) ? '<a class="text-danger popovers pull-right" data-container="body" data-trigger="hover" data-original-title="错误信息" data-content="' + err + '"><i class="fa fa-warning"></i></a>' : '';
                    td = '<td><a class="tooltips" data-container="body" data-trigger="hover" data-original-title="查看结果" data-toggle="modal" data-target="#myViewModal" onclick="view(\'result\',\'' + v.id + '\',\'' + v.ip + '\')">' + v.ip + '</a>' + warn + '</td>';
                    tr.append(td);
                    if (typeof cache.task.options != 'undefined') {
                        var step = '', curr = -1, style = '';
                        $.each(cache.task.options, function (key, val) {
                            style = 'label label-default';
                            if (step) step += ' <i class="fa fa-angle-double-right text-danger"></i> ';
                            if (val.name == v.steps) {
                                curr = key;
                                switch (idx) {
                                    case 'running':
                                        style = 'badge bg-blue';
                                        break;
                                    case 'success':
                                        style = 'badge bg-green';
                                        break;
                                    case 'failed':
                                        style = 'badge bg-red';
                                        break;
                                    case 'stoped':
                                        style = 'badge bg-orange';
                                        break;
                                }
                            } else {
                                if (idx != 'ready') if (key < curr || curr == -1) style = 'label label-success';
                            }
                            step += '<a class="tooltips ' + style + '" title="第' + (key + 1) + '步">' + val.name + '</a>'
                        });
                        td = '<td>' + step + '</td>';
                    } else {
                        td = '<td>' + v.steps + '</td>';
                    }
                    tr.append(td);
                    if (idx != 'ready') {
                        var runTime = (typeof(v.runTime) != 'undefined') ? v.runTime : 0;
                        // var beginSec = (typeof(v.Created)!='undefined') ? Date.parse(v.Created) : 0;
                        // var endSec = (typeof(v.Updated)!='undefined') ? ((v.Updated!='0000-00-00 00:00:00') ? Date.parse(v.Updated) : new Date().getTime()) : new Date().getTime();
                        var timeLen = (runTime > 0) ? Math.ceil(runTime) + '秒' : '-';
                        td = '<td>' + timeLen + '</td>';
                        tr.append(td);
                    }
                    var btn = '<a class="tooltips" data-container="body" data-trigger="hover" data-original-title="查看结果" data-toggle="modal" data-target="#myViewModal" onclick="view(\'result\',\'' + v.id + '\',\'' + v.ip + '\')"><i class="fa fa-comment"></i></a>';
                    td = '<td>' + btn + '</td>';
                    tr.append(td);
                    body.append(tr);
                    n++;
                }
            }
            break;
    }
    $('.popovers').each(function () {
        $(this).popover();
    });
    $('.tooltips').each(function () {
        $(this).tooltip();
    });
}

//二次操作确认
var twiceCheck = function (action, idx, ip) {
    NProgress.start();
    if (!idx) idx = cache.task_id;
    var modalTitle = '', modalBody = '', notice = '',list = '', btnDisable = false;
    if (!action || !idx) {
        modalTitle = '非法请求';
        notice = '<div class="alert alert-danger">错误信息：参数错误</div>';
        pageNotify('error', '非法请求！', '错误信息：参数错误');
    } else {
        var count = 0, diableCount = 0,postNodeIds=[];
        switch (action) {
            case 'start':
                modalTitle = '启动任务';
                //除任务处在执行中不能重新启动任务外，其他状况均可重新启动任务
                // if (cache..state == '1') {
                //     notice = '<div class="alert alert-danger">错误信息：任务执行中</div>';
                // }
                // if(cache.task.state=='1'||cache.task.state=='2'||cache.task.state=='3'){
                //   notice='<div class="alert alert-danger">错误信息：任务执行中或已完成</div>';
                // }

                if(cache.ip.ready.length > 0){
                        $('input:checkbox[id=list_ready]:checked').each(function(i){
                            count++;
                            var node = $(this).val().split("_");
                            var node_id = node[0];
                            var node_ip = node[1];
                            postNodeIds.push(node_id);
                            list+='<span class="col-sm-3" id="check_'+node_ip+'">'+node_ip+'</span>';

                        });
                    }
                    if(cache.ip.running.length > 0){
                        $('input:checkbox[id=list_running]:checked').each(function(i){
                            var node = $(this).val().split("_");
                            var node_ip = node[1];
                            diableCount++;
                            list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'(<span style="margin-left:3px;" class="badge bg-purple">任务执行中</span>)</span>';
                            notice = '<div class="alert alert-danger">错误信息：任务执行中或者已成功</div>';

                        });
                    }
                    if(cache.ip.success.length > 0){
                        $('input:checkbox[id=list_success]:checked').each(function(i){
                            var node = $(this).val().split("_");
                            var node_ip = node[1];
                            diableCount++;
                            list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(2)+')</span>';
                            notice = '<div class="alert alert-danger">错误信息：任务执行中或者已成功</div>';
                        });
                    }
                    if(cache.ip.failed.length > 0){
                        $('input:checkbox[id=list_failed]:checked').each(function(i){
                            count++;
                            var node = $(this).val().split("_");
                            var node_id = node[0];
                            var node_ip = node[1];
                            postNodeIds.push(node_id);
                            list+='<span class="col-sm-3" id="check_'+node_ip+'">'+node_ip+'</span>';

                        });
                }
                if(cache.ip.stoped.length > 0){
                    $('input:checkbox[id=list_stoped]:checked').each(function(i){
                        count++;
                        var node = $(this).val().split("_");
                        var node_id = node[0];
                        var node_ip = node[1];
                        postNodeIds.push(node_id);
                        list+='<span class="col-sm-3" id="check_'+node_ip+'">'+node_ip+'</span>';

                    });
                }

                break;
            case 'pause':
                modalTitle = '暂停任务';
                if(cache.ip.ready.length > 0){
                    $('input:checkbox[id=list_ready]:checked').each(function(i){
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        diableCount++;
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(0)+')</span>';
                        notice = '<div class="alert alert-danger">错误信息：只有执行中的任务才允许暂停</div>';

                    });
                }
                if(cache.ip.running.length > 0){
                    $('input:checkbox[id=list_running]:checked').each(function(i){
                        count++;
                        var node = $(this).val().split("_");
                        var node_id = node[0];
                        var node_ip = node[1];
                        postNodeIds.push(node_id);
                        list+='<span class="col-sm-3" id="check_'+node_ip+'">'+node_ip+'</span>';
                    });
                }
                if(cache.ip.success.length > 0){
                    $('input:checkbox[id=list_success]:checked').each(function(i){
                        diableCount++;
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(2)+')</span>';
                        notice = '<div class="alert alert-danger">错误信息：只有执行中的任务才允许暂停</div>';

                    });
                }
                if(cache.ip.failed.length > 0){
                    $('input:checkbox[id=list_failed]:checked').each(function(i){
                        diableCount++;
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(3)+')</span>';
                        notice = '<div class="alert alert-danger">错误信息：只有执行中的任务才允许暂停</div>';

                    });
                }
                if(cache.ip.stoped.length > 0){
                    $('input:checkbox[id=list_stoped]:checked').each(function(i){
                        diableCount;
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(4)+')</span>';
                        notice = '<div class="alert alert-danger">错误信息：只有执行中的任务才允许暂停</div>';

                    });
                }
                // if (cache.task.state != '1') {
                //     notice = '<div class="alert alert-danger">错误信息：只有执行中的任务才允许暂停</div>';
                // }
                break;
            case 'finish':
                modalTitle = '完成任务';
                //当状态是未启动或者是已经完成或时，完成任务不可操作，其他状态均可完成
                // if (cache.task.state == '0' || cache.task.state == '2') {
                //     notice = '<div class="alert alert-danger">错误信息：任务未启动或已完成</div>';
                // }
                // if(cache.task.state=='0'||cache.task.state=='2'||cache.task.state=='3'){
                //   notice='<div class="alert alert-danger">错误信息：任务未启动或已完成</div>';
                // }


                if(cache.ip.ready.length > 0){
                    $('input:checkbox[id=list_ready]:checked').each(function(i){
                        diableCount++;
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(0)+')</span>';
                        notice='<div class="alert alert-danger">错误信息：任务未启动、暂停或已完成</div>';

                    });
                }
                if(cache.ip.running.length > 0){
                    $('input:checkbox[id=list_running]:checked').each(function(i){
                        count++;
                        var node = $(this).val().split("_");
                        var node_id = node[0];
                        var node_ip = node[1];
                        postNodeIds.push(node_id);
                        list+='<span class="col-sm-3" id="check_'+node_ip+'">'+node_ip+'</span>';
                    });
                }
                if(cache.ip.success.length > 0){
                    $('input:checkbox[id=list_success]:checked').each(function(i){
                        diableCount++;
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(2)+')</span>';
                        notice='<div class="alert alert-danger">错误信息：任务未启动、暂停或已完成</div>';

                    });
                }

                if(cache.ip.failed.length > 0){
                    $('input:checkbox[id=list_failed]:checked').each(function(i){
                        diableCount++;
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(3)+')</span>';
                        notice='<div class="alert alert-danger">错误信息：任务未启动、暂停或已完成</div>';

                    });
                }

                if(cache.ip.stoped.length > 0){
                    $('input:checkbox[id=list_stoped]:checked').each(function(i){
                        diableCount++;
                        var node = $(this).val().split("_");
                        var node_ip = node[1];
                        list+='<span class="col-sm-3 text-success" id="check_'+node_ip+'">'+node_ip+'('+getStatusAlias(4)+')</span>';
                        notice='<div class="alert alert-danger">错误信息：任务未启动、暂停或已完成</div>';

                    });
                }
                break;
            default:
                modalTitle = '';
                break;
        }
        var total_count = count+diableCount;
        modalBody = modalBody + '<div class="form-group col-sm-12">';
        modalBody = modalBody + '<h4>确认' + modalTitle + '? <span class="text text-primary">警告! 请谨慎操作!</span></h4>';
        modalBody+='<p><strong class="text-primary">选中总数</strong>: 共 <span class="badge badge-danger">'+total_count+'</span> 个 </p>';
        modalBody+='<p><strong class="text-primary">可用总数</strong>: 共 <span class="badge badge-danger">'+count+'</span> 个 </p>';
        modalBody = modalBody + '<strong class="text-primary">任务ID: ' + idx + '</strong><br/>';
        modalBody+='<div style="margin-top:5px;" class="col-sm-12">'+list+'</div>';
        modalBody = modalBody + '</div>';
        modalBody = modalBody + '<div class="col-sm-12" id="modalNotice"></div>';
        modalBody+='<textarea class="hidden" id="node_idss" name="node_ids">'+JSON.stringify(postNodeIds)+'</textarea>'
        modalBody = modalBody + '<input type="hidden" id="page_action" name="page_action" value="' + action + '">';
        modalBody = modalBody + '<input type="hidden" id="id" name="id" value="' + idx + '">';
    }
    if (!modalTitle) {
        modalTitle = '非法请求';
        notice = '<div class="note note-danger">错误信息：参数错误</div>';
        pageNotify('error', '非法请求！', '错误信息：参数错误');
    }
    modalTitle += ' - ' + idx;
    if (ip) modalTitle += ' / ' + idx;
    $('#myModalLabel').html(modalTitle);
    $('#myModalBody').html(modalBody);
    if(count == 0){
        $('#btnCommit').attr('disabled', true);
        if (notice != '') {
            $('#modalNotice').html(notice);
        }else{
            $('#modalNotice').html('<div class="alert alert-danger">错误信息：没有选中任何节点</div>');
        }
    }else{
        $('#btnCommit').attr('disabled', btnDisable);
        if (notice != '') {
            $('#modalNotice').html(notice);
        }
    }

    NProgress.done();
}

var getStatusAlias = function(status){
    var str=status;
    switch(status){
        case 0:
            str='<span style="margin-left:3px;" class="badge">准备中</span>';break;
        case 1:
            str='<span style="margin-left:3px;" class="badge bg-blue-sky">执行中</span>';break;
        case 2:
            str='<span style="margin-left:3px;" class="badge bg-green">已完成</span>';break;
        case 3:
            str='<span style="margin-left:3px;" class="badge bg-red">失败</span>';break;
        case 4:
            str='<span style="margin-left:3px;" class="badge bg-orange">暂停</span>';break;
        default:
            str='<span style="margin-left:3px;" class="badge">未知状态</span>';break;
    }
    return str;
}
//增删改查
var change = function () {
    var url = '/api/for_layout/task.php';
    var postData = {};
    var form = $('#myModalBody').find("input,select,textarea");

    //处理表单内容--不需要修改
    $.each(form, function (i) {
        switch (this.type) {
            case 'radio':
                if (typeof(postData[this.name]) == 'undefined') {
                    if (this.name) postData[this.name] = $('input[name="' + this.name + '"]:checked').val();
                }
                break;
            case 'checkbox':
                if (this.id) {
                    if (typeof(postData[this.id]) == 'undefined') {
                        postData[this.id] = {};
                    }
                    if (this.checked) {
                        postData[this.id][i] = this.value;
                    }
                }
                break;
            default:
                if (this.name) postData[this.name] = this.value;
                break;
        }
    });
    var action = $("#page_action").val();
    delete postData['page_action'];
    //console.log("action="+action);
    //console.log(JSON.stringify(postData));
    var actionDesc = '';
    switch (action) {
        case 'insert':
            actionDesc = '添加';
            break;
        case 'update':
            actionDesc = '更新';
            break;
        case 'delete':
            actionDesc = '删除';
            break;
        case 'start':
            if (cache.refreshInterval == null) {
                cache.refreshInterval = setInterval('getTask(\'info\');', 10000);
            }
            actionDesc += '启动任务';
            postData["node_ids"]=JSON.parse(postData.node_ids);
            break;
        case 'pause':
            actionDesc += '暂停任务';
            postData["node_ids"]=JSON.parse(postData.node_ids);
            break;
        case 'finish':
            actionDesc += '完成任务';
            postData["node_ids"]=JSON.parse(postData.node_ids);
            break;
        default:
            actionDesc = action;
            break;
    }
    $.ajax({
        type: "POST",
        url: url,
        data: {"action": action, "data": JSON.stringify(postData)},
        dataType: "json",
        success: function (data) {
            //执行结果提示
            if (data.code == 0) {
                pageNotify('success', '【' + actionDesc + '】操作成功！');
            } else {
                if (cache.refreshInterval != null) {
                    clearInterval(cache.refreshInterval);
                    cache.refreshInterval = null;
                }
                pageNotify('error', '【' + actionDesc + '】操作失败！', '错误信息：' + data.msg);
            }
            //重载列表
            getTask('info');
            //处理模态框和表单
            $("#myModal :input").each(function () {
                $(this).val("");
            });
            $("#myModal").on("hidden.bs.modal", function () {
                $(this).removeData("bs.modal");
            });
        },
        error: function () {
            if (cache.refreshInterval != null) {
                clearInterval(cache.refreshInterval);
                cache.refreshInterval = null;
            }
            pageNotify('error', '【' + actionDesc + '】操作失败！', '错误信息：接口不可用');
        }
    });
}

//任务操作
var controlTask = function (action, idx) {
    if (!idx) idx = cache.task_id;
    var postData = '', actionDesc = '';
    var url = '/api/for_layout/task.php';
    postData = {"action": action, "data": JSON.stringify({"task_id": idx})};
    switch (action) {
        case 'start':
            actionDesc = '启动任务';
            break;
        case 'stop':
            actionDesc = '暂停任务';
            break;
        case 'finish':
            actionDesc = '完成任务';
            break;
        case 'modify':
            actionDesc = '修改任务';
            break;
        default:
            postData = '';
            break;
    }
    if (postData) {
        $.ajax({
            type: "POST",
            url: url,
            data: postData,
            dataType: "json",
            success: function (data) {
                //执行结果提示
                if (data.code == 0) {
                    pageNotify('success', '【' + actionDesc + '】操作成功！');
                    if (cache.flag == false) {
                        cache.flag = true;
                        getTask('info');
                    }
                } else {
                    pageNotify('error', '【' + actionDesc + '】操作失败！', '错误信息：' + data.msg);
                }
                //处理模态框和表单
                $("#myModal :input").each(function () {
                    $(this).val("");
                });
                $("#myModal").on("hidden.bs.modal", function () {
                    $(this).removeData("bs.modal");
                });
                $("#myModal").draggable({
                    handle: ".modal-header"
                });
            },
            error: function () {
                pageNotify('error', '【' + actionDesc + '】操作失败！', '错误信息：接口不可用');
            }
        });
    } else {
        pageNotify('error', '非法请求！', '错误信息：参数错误');
    }
}

//view
var view = function (type, idx, ip, offset) {
    NProgress.start();
    if (!offset) offset = 0;
    var url = '', title = '', text = '', illegal = false, height = '', postData = {};
    var tStyle = 'word-break:break-all;word-warp:break-word;';
    url = '/api/for_layout/task.php';
    switch (type) {
        case 'result':
            title = '查看任务结果 - ' + ip;
            postData = {"action": "result", "fIdx": idx};
            break;
        case 'tasklog':
            title = '查看任务总日志';
            postData = {"action": "tasklog", "fIdx": idx};
            break;
        default:
            illegal = true;
            break;
    }
    if (!illegal && idx != '') {
        $.ajax({
            type: "POST",
            url: url,
            data: postData,
            dataType: "json",
            success: function (data) {
                //执行结果提示
                if (data.code == 0) {
                    if (typeof(data.content) != 'undefined') {
                        switch (type) {
                            case 'result':
                                cache.result = {};
                                text = '<div class="col-sm-12" id="result_' + idx + '">';
                                $.each(data.content, function (k, v) {
                                    $.each(v, function (key, val) {
                                        cache.result[key] = val;
                                    });
                                });
                                $.each(cache.task.options, function (k, v) {
                                    text += '<h5 class="col-sm-12 text-primary">步骤 ' + (k + 1) + ' : ' + v.name + '</h5>';
                                    var log = (typeof cache.result[v.name] != 'undefined') ? cache.result[v.name].replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/\n/g, '<br>') : '此步骤暂无日志';
                                    if (!log) log = '日志数据为空';
                                    text += '<span class="col-sm-12" style="background-color:#000;color:#ccc;line-height: 150%">' + log + '</span>';
                                });
                                text += '</div>';
                                if (cache.task.state != 4) window.setInterval('getResult("' + idx + '","' + ip + '")', 5000);
                                break;
                            case 'tasklog':
                                cache.result = {};
                                var log = '';
                                text = '<div class="col-sm-12" id="tasklog_' + idx + '">';
                                $.each(data.content, function (k, v) {
                                    log += '[ ' + timeStampToSting(v['ctime']) + '] ' + v['message'] + "<br>"
                                });

                                text += '<span class="col-sm-12" style="background-color:#000;color:#ccc;line-height: 150%">' + log + '</span>';
                                text += '</div>';
                                console.log(cache.task);
                                if (cache.task.state != 4) window.setInterval('getTaskLog("' + idx + '")', 5000);
                                break;
                            default:
                                var locale = {};
                                if (typeof(locale_messages.layout)) locale = locale_messages.layout;
                                $.each(data.content, function (k, v) {
                                    if (v == '') v = '空';
                                    if (typeof(locale[k]) != 'undefined') k = locale[k];
                                    text += '<span class="col-sm-2"><strong>' + k + '</strong></span> <span class="col-sm-4" style="' + tStyle + '">' + v + '</span>' + "\n";
                                });
                                break;
                        }
                        if (!text) {
                            pageNotify('warning', '数据为空！');
                            text = '<div class="note note-warning">数据为空！</div>';
                        }
                    } else {
                        pageNotify('warning', '数据为空！');
                        text = '<div class="note note-warning">数据为空！</div>';
                    }
                } else {
                    pageNotify('error', '加载失败！', '错误信息：' + data.msg);
                    text = '<div class="note note-danger">错误信息：' + data.msg + '</div>';
                }
                setTimeout(function () {
                    if (height != '') {
                        $('#myViewModalBody').css('height', height);
                    }
                    $('#myViewModalLabel').html(title);
                    $('#myViewModalBody').html(text);
                    NProgress.done();
                }, 200);
            },
            error: function () {
                pageNotify('error', '加载失败！', '错误信息：接口不可用');
                text = '<div class="note note-danger">错误信息：接口不可用</div>';
                $('#myViewModalLabel').html(title);
                $('#myViewModalBody').html(text);
                NProgress.done();
            }
        });
    } else {
        pageNotify('warning', '错误操作！', '错误信息：参数错误');
        title = '非法请求 - ' + type;
        text = '<div class="note note-danger">错误信息：参数错误</div>';
        $('#myViewModalLabel').html(title);
        $('#myViewModalBody').html(text);
        NProgress.done();
    }
}

//查看任务参数
var viewArg = function () {
    NProgress.start();
    var data = cache.task.options, title = '', text = '', tStyle = 'word-break:break-all;word-warp:break-word;';
    title = '查看任务参数';
    var locale = {};
    if (typeof(locale_messages.layout)) locale = locale_messages.layout;
    if (!$.isEmptyObject(data)) {
        var i = 0;
        $.each(data, function (k, v) {
            i++;
            text += '<h5 class="text-primary col-sm-12">步骤' + i + ': ' + v.name + '</h5>' + "\n";

            if (v == '') {
                text += '<span class="col-sm-12">此步骤无参数</span>' + "\n";
            } else {
                var j = 0;
                $.each(v, function (key, val) {
                    if (key != 'name') {
                        j++;
                        if (j % 4 == 0) text += '<div class="clearfix"></div>' + "\n";
                        if (val == '') val = '参数值为空';
                        if (typeof(locale[key]) != 'undefined') key = locale[key];
                        text += '<div class="col-sm-12"><span class="title col-sm-1" style="font-weight: bold;">' + key + '</span> <span class="col-sm-11" style="' + tStyle + '">' + ((typeof val == 'object') ? JSON.stringify(val) : val) + '</span></div>' + "\n";
                    }
                });
            }
        });
    } else {
        pageNotify('warning', '数据为空！');
        text = '<div class="note note-warning">数据为空！</div>';
    }
    $('#myViewModalLabel').html(title);
    $('#myViewModalBody').html(text);
    NProgress.done();
}

//获取最近任务列表
var getList = function () {
    var url = '/api/for_layout/task.php?action=list';
    var postData = {"action": "list", "page": 1, "pagesize": 20};
    if (cache.task_id == 1) {
        $('#task_prev').attr('disabled', true).attr('href', 'javascripts:;').attr('onclick', 'return false;');
    }
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (listdata) {
            if (listdata.code == 0) {
                cache.tasklist = listdata.content;
                var body = $('#list');
                if (cache.tasklist.length > 0) {
                    body.html('');
                    $.each(cache.tasklist, function (k, v) {
                        var li = $('<li></li>');
                        var t = (v.task_name) ? v.task_name : '任务ID: ' + v.id;
                        if (v.id == cache.task_id) {
                            li.append('<a href="task_detail.php?idx=' + v.id + '" style="padding:3px 8px;"><span class="text-primary">' + t + '</span><span class="badge bg-green pull-right">当前</span></a>');
                            if (k == 0) {
                                $('#task_next').attr('href', 'javascripts:;').attr('disabled', true).attr('onclick', 'return false;');
                            }
                        } else {
                            li.append('<a href="task_detail.php?idx=' + v.id + '" style="padding:3px 8px;">' + t + '</a>');
                        }
                        body.append(li);
                    });
                    if (cache.tasklist.length > 10) body.css('height', '200px');
                } else {
                    $('#task_next').attr('href', 'javascripts:;').attr('disabled', true).attr('onclick', 'return false;');
                }
            } else {
                pageNotify('error', '加载最近任务列表失败！', '错误信息：' + listdata.msg);
                $('#task_next').attr('href', 'javascripts:;').attr('disabled', true).attr('onclick', 'return false;');
            }
        },
        error: function () {
            pageNotify('error', '加载最近任务列表失败！', '错误信息：接口不可用');
            $('#task_next').attr('href', 'javascripts:;').attr('disabled', true).attr('onclick', 'return false;');
        }
    });
}

//字符串指定长度插入字符
var insertStr = function (str, flg, idx) {
    if (idx) idx = 40;
    var arr = str.split(' '), ret = '';
    for (var j = 0; j < arr.length; j++) {
        var len = arr[j].length;
        if (len < idx) {
            ret += arr[j] + flg;
            continue;
        }
        var num = len / idx;
        var start, end;
        for (i = 0; i < num; i++) {
            var tmp = '';
            if (len % idx != 0) {//不能完整分段
                start = i * idx - 1;
                end = i * idx + (idx - 1);
            } else {
                start = i * idx;
                end = (i + 1) * idx;
            }
            start = start < 0 ? 0 : start;
            if (end <= len) {
                tmp = arr[j].substring(start, end);
            }
            ret += (end >= len) ? tmp : tmp + flg;
        }
    }
    return ret;
}

//获取结果
var getResult = function (idx, ip) {
    if (!idx || !ip) return false;
    if ($('#result_' + idx).length == 0) return false;
    var postData = {"action": "result", "fIdx": idx};
    url = '/api/for_layout/task.php?action=result';
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            //执行结果提示
            if (data.code == 0) {
                if (typeof(data.content) != 'undefined') {
                    if ($('#result_' + idx).length > 0) {
                        cache.result = {};
                        var text = '';
                        $.each(data.content, function (k, v) {
                            $.each(v, function (key, val) {
                                cache.result[key] = val;
                            });
                        });
                        $.each(cache.task.options, function (k, v) {
                            text += '<h5 class="col-sm-12 text-primary">步骤 ' + (k + 1) + ' : ' + v.name + '</h5>';
                            var log = (typeof cache.result[v.name] != 'undefined') ? cache.result[v.name].replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/\n/g, '<br>') : '此步骤暂无日志';
                            if (!log) log = '日志数据为空';
                            text += '<span class="col-sm-12" style="background-color:#000;color:#ccc;line-height: 150%">' + log + '</span>';
                        });
                        $('#result_' + idx).html(text);
                        $('#result_' + idx).append('<span class="pull-right text-danger">Updated @' + getDate(new Date(), 'time') + '</span>');
                    }
                } else {
                    pageNotify('warning', '数据为空！');
                }
            } else {
                pageNotify('error', '加载结果失败！', '错误信息：' + data.msg);
            }
        },
        error: function () {
            pageNotify('error', '加载结果失败！', '错误信息：接口不可用');
        }
    });
}


//获取任务主日志
var getTaskLog = function (taskId) {
    if (!taskId) return false;
    if ($('#tasklog_' + taskId).length == 0) return false;

    var postData = {"action": "tasklog", "fIdx": taskId};
    url = '/api/for_layout/task.php';
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            //执行结果提示
            if (data.code == 0) {
                if (typeof(data.content) != 'undefined') {
                    if ($('#tasklog_' + taskId).length > 0) {

                        var text = '<span class="col-sm-12" style="background-color:#000;color:#ccc;line-height: 150%">';
                        $.each(data.content, function (k, v) {
                            text += '[ ' + timeStampToSting(v['ctime']) + '] ' + v['message'] + "<br>"
                        });

                        text += '</span>';
                        text += '<span class="pull-right text-danger">Updated @' + getDate(new Date(), 'time') + '</span>'

                        $('#tasklog_' + taskId).html(text);
                    }
                } else {
                    pageNotify('warning', '数据为空！');
                }
            } else {
                pageNotify('error', '加载结果失败！', '错误信息：' + data.msg);
            }
        },
        error: function () {
            pageNotify('error', '加载结果失败！', '错误信息：接口不可用');
        }
    });
}

/**
 * 将UNIX时间戳 转化为 2015-06-11 11:10:37 形式的字符串
 *
 *
 * @access public
 * @param int data  UNIX时间戳 格式如：1430982663 [Must]
 * @return string $str
 */
var timeStampToSting = function (data) {
    var time = new Date(data * 1000);
    var y = time.getFullYear();
    var m = time.getMonth() + 1;
    var d = time.getDate();
    var h = time.getHours();
    var mm = time.getMinutes();
    var s = time.getSeconds();

    m = (m < 10) ? '0' + m : m;
    d = (d < 10) ? '0' + d : d;
    h = (h < 10) ? '0' + h : h;
    mm = (mm < 10) ? '0' + mm : mm;
    s = (s < 10) ? '0' + s : s;

    return y + '-' + m + '-' + d + ' ' + h + ':' + mm + ':' + s;
}

var checkAll = function(o,idx){
    var body = 'list_' + idx;
    $('[id='+body+']:checkbox').prop('checked', o.checked);

}
