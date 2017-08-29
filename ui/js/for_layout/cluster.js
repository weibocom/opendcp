cache = {
    page: 1,
}

var list = function (page) {
    NProgress.start();
    var postData = {};
    var url = '/api/for_layout/cluster.php';
    if (!page) {
        page = cache.page;
    } else {
        cache.page = page;
    }
    var fIdx = $('#fIdx').val();
    postData = {"action": "list", "fIdx": fIdx};
    url += '?page=' + page;
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (listdata) {
            if (listdata.code == 0) {
                var pageinfo = $("#table-pageinfo");//分页信息
                var paginate = $("#table-paginate");//分页代码
                var head = $("#table-head");//数据表头
                var body = $("#table-body");//数据列表
                //清除当前页面数据
                pageinfo.html("");
                paginate.html("");
                head.html("");
                body.html("");
                //生成页面
                //生成分页
                processPage(listdata, pageinfo, paginate);
                //生成列表
                processBody(listdata, head, body);
                $('.popovers').each(function () {
                    $(this).popover();
                });
                $('.tooltips').each(function () {
                    $(this).tooltip();
                });
                switchery();
            } else {
                pageNotify('error', '加载失败！', '错误信息：' + listdata.msg);
            }
            NProgress.done();
        },
        error: function () {
            pageNotify('error', '加载失败！', '错误信息：接口不可用');
            NProgress.done();
        }
    });
}

//生成分页
var processPage = function (data, pageinfo, paginate) {
    var begin = data.pageSize * ( data.page - 1 ) + 1;
    var end = ( data.count > begin + data.pageSize - 1 ) ? begin + data.pageSize - 1 : data.count;
    pageinfo.html('Showing ' + begin + ' to ' + end + ' of ' + data.count + ' records');
    var p1 = (data.page - 1 > 0) ? data.page - 1 : 1;
    var p2 = data.page + 1;
    prev = '<li><a href="javascript:;" onclick="list(' + p1 + ')"><i class="fa fa-angle-left"></i></a></li>';
    paginate.append(prev);
    for (var i = 1; i <= data.pageCount; i++) {
        var li = '';
        if (i == data.page) {
            li = '<li class="active"><a href="javascript:;" onclick="list(' + i + ')">' + i + '</a></li>';
        } else {
            if (i == 1 || i == data.pageCount) {
                li = '<li><a href="javascript:;" onclick="list(' + i + ')">' + i + '</a></li>';
            } else {
                if (i == p1) {
                    if (p1 > 2) {
                        console.log(i + ' ' + p1);
                        li = '<li class="disabled"><a href="#">...</a></li>' + "\n" + '<li><a href="javascript:;" onclick="list(' + i + ')">' + i + '</a></li>';
                    } else {
                        li = '<li><a href="javascript:;" onclick="list(' + i + ')">' + i + '</a></li>';
                    }
                } else {
                    if (i == p2) {
                        if (p2 < data.pageCount - 1) {
                            li = '<li><a href="javascript:;" onclick="list(' + i + ')">' + i + '</a></li>' + "\n" + '<li class="disabled"><a href="#">...</a></li>';
                        } else {
                            li = '<li><a href="javascript:;" onclick="list(' + i + ')">' + i + '</a></li>';
                        }
                    }
                }
            }
        }
        paginate.append(li);
    }
    if (p2 > data.pageCount) p2 = data.pageCount;
    next = '<li class="next"><a href="javascript:;" title="Next" onclick="list(' + p2 + ')"><i class="fa fa-angle-right"></i></a></li>';
    paginate.append(next);
}

//生成列表
var processBody = function (data, head, body) {
    var td = "";
    if (data.title) {
        var tr = $('<tr></tr>');
        for (var i = 0; i < data.title.length; i++) {
            var v = data.title[i];
            td = '<th>' + v + '</th>';
            tr.append(td);
        }
        head.html(tr);
    }
    if (data.content) {
        if (data.content.length > 0) {
            for (var i = 0; i < data.content.length; i++) {
                var v = data.content[i];
                var tr = $('<tr></tr>');
                td = '<td>' + v.i + '</td>';
                tr.append(td);
                td = '<td>' + v.id + '</td>';
                tr.append(td);
                td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'cluster\',\'' + v.id + '\')">' + v.name + '</a></td>';
                tr.append(td);
                td = '<td>' + v.desc + '</td>';
                tr.append(td);
                var btnEdit = '<a class="btn blue tooltips" data-toggle="modal" data-target="#myModal" href="edit_cluster.php?action=edit&idx=' + v.id + '" style="padding-left: 0px;" title="修改"><i class="fa fa-edit"></i></a>';
                var btnDel = '<a class="btn red tooltips" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\'' + v.id + '\',\'' + v.name + '\')" style="padding-left: 0px;" title="删除"><i class="fa fa-trash-o"></i></a>';
                td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnEdit + btnDel + '</div></td>';
                tr.append(td);

                body.append(tr);
            }
        } else {
            pageNotify('info', 'Warning', '数据为空！');
        }
    } else {
        pageNotify('warning', 'error', '接口异常！');
    }
}

//增删改查
var change = function () {
    NProgress.start();
    var url = '/api/for_layout/cluster.php';
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
    $.ajax({
        type: "POST",
        url: url,
        data: {"action": action, "data": JSON.stringify(postData)},
        dataType: "json",
        success: function (data) {
            //执行结果提示
            if (data.code == 0) {
                pageNotify('success', '操作成功！');
            } else {
                pageNotify('error', '操作失败！', '错误信息：' + data.msg);
            }
            //重载列表
            list();
            //处理模态框和表单
            $("#myModal :input").each(function () {
                $(this).val("");
            });
            $("#myModal").on("hidden.bs.modal", function () {
                $(this).removeData("bs.modal");
            });
            NProgress.done();
        },
        error: function () {
            pageNotify('error', '操作失败！', '错误信息：接口不可用');
            NProgress.done();
        }
    });
}

var get = function (idx) {
    var tab = $('#tab').val();
    url = '/api/for_layout/' + tab + '.php';
    switch (tab) {
        case 'cluster':
            postData = {"action": "info", "fIdx": idx};
            break;
    }
    if (idx != '') {
        $.ajax({
            type: "POST",
            url: url,
            data: postData,
            dataType: "json",
            success: function (data) {
                //执行结果提示
                if (data.code == 0) {
                    if (typeof(data.content) != 'undefined') {
                        //pageNotify('success','加载成功！');
                        $.each(data.content, function (k, v) {
                            if ($('#' + k).length > 0) {
                                switch ($('#' + k).get(0).tagName) {
                                    case 'INPUT':
                                        switch ($('#' + k).attr('type')) {
                                            case 'radio':
                                                $("input[name='" + k + "'][value='" + v + "']").attr("checked", true);
                                                break;
                                            case 'checkbox':
                                                $.each(v, function (k1, v1) {
                                                    $("input[id='" + k + "']:checkbox[value='" + v1 + "']").attr('checked', 'true');
                                                });
                                                break;
                                            default:
                                                $('#' + k).val(v);
                                                break;
                                        }
                                        break;
                                    case 'SELECT':
                                        if ($('#' + k).find("option[value='" + v + "']").length == 0) {
                                            $('#' + k).append('<option value="' + v + '">' + v + '</option>');
                                        }
                                        $('#' + k).find("option[value='" + v + "']").attr("selected", true);
                                        break;
                                    default:
                                        $('#' + k).val(v);
                                        break;
                                }
                            }
                        });
                        getCloudOrg();
                    } else {
                        pageNotify('warning', '数据为空！');
                    }
                } else {
                    pageNotify('error', '加载失败！', '错误信息：' + data.msg);
                }
            },
            error: function () {
                pageNotify('error', '加载详情失败！', '错误信息：接口不可用');
            }
        });
    } else {
        pageNotify('warning', '加载详情失败！', '错误信息：参数错误');
    }
}

var twiceCheck = function (action, idx, desc) {
    NProgress.start();
    if (!idx) idx = '';
    if (!desc) desc = '';
    var modalTitle = '', modalBody = '', list = '', notice = '', btnDisable = false;
    var tab = $('#tab').val();
    if (!action) {
        modalTitle = '非法请求';
        notice = '<div class="note note-danger">错误信息：参数错误</div>';
        pageNotify('error', '非法请求！', '错误信息：参数错误');
    } else {
        switch (tab) {
            case 'cluster':
                switch (action) {
                    case 'del':
                        modalTitle = '删除集群';
                        modalBody = modalBody + '<div class="form-group col-sm-12">';
                        modalBody = modalBody + '<div class="note note-danger">';
                        modalBody = modalBody + '<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 集群ID : ' + idx + '<br>集群名称 : ' + desc;
                        modalBody = modalBody + '</div>';
                        modalBody = modalBody + '</div>';
                        modalBody = modalBody + '<input type="hidden" id="id" name="id" value="' + idx + '">';
                        modalBody = modalBody + '<input type="hidden" id="page_action" name="page_action" value="delete">';
                        break;
                    default:
                        modalTitle = '非法请求';
                        notice = '<div class="note note-danger">错误信息：参数错误</div>';
                        pageNotify('error', '非法请求！', '错误信息：参数错误');
                        break;
                }
                break;
            default:
                modalTitle = '非法请求';
                notice = '<div class="note note-danger">错误信息：参数错误</div>';
                pageNotify('error', '非法请求！', '错误信息：参数错误');
                break;
        }
    }
    $('#myModalLabel').html(modalTitle);
    $('#myModalBody').html(modalBody);
    if (notice != '') {
        $('#modalNotice').html(notice);
        $('#btnCommit').attr('disabled', true);
    } else {
        $('#btnCommit').attr('disabled', btnDisable);
    }
    NProgress.done();
}

//commit check
var check = function (tab) {
    if (!tab) tab = $('#tab').val();
    switch (tab) {
        case 'cluster':
            var disabled = false;
            if ($('#name').val() == '') disabled = true;
            if ($('#desc').val() == '') disabled = true;
            if ($('#biz').val() == '') disabled = true;
            $("#btnCommit").attr('disabled', disabled);
            break;
    }
}

//view
var view = function (type, idx) {
    NProgress.start();
    var url = '', title = '', text = '', illegal = false, height = '', postData = {};
    var tStyle = 'word-break:break-all;word-warp:break-word;';
    switch (type) {
        case 'cluster':
            url = '/api/for_layout/cluster.php';
            title = '查看集群详情 - ' + idx;
            postData = {"action": "info", "fIdx": idx};
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
                        //pageNotify('success','加载成功！');
                        var locale = {};
                        if (typeof(locale_messages.layout)) locale = locale_messages.layout;
                        switch (type) {
                            case 'cluster':
                                $.each(data.content, function (k, v) {
                                    if (locale[k] != false) {
                                        if (v == '') v = '空';
                                        if (typeof(locale[k]) != 'undefined') k = locale[k];
                                        text += '<span class="title col-sm-2" style="font-weight: bold;">' + k + '</span> <span class="col-sm-4" style="' + tStyle + '">' + v + '</span>' + "\n";
                                    }
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
        title = '非法请求';
        text = '<div class="note note-danger">错误信息：参数错误</div>';
        $('#myViewModalLabel').html(title);
        $('#myViewModalBody').html(text);
        NProgress.done();
    }
}
