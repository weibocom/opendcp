cache={
    page:1,
    resourceData: {},
}

var getDate = function(t){
    if(!t) t='';
    var d= new Date(t);
    var M= (d.getMonth()+1);
    var D= d.getDate();
    var h= d.getHours();
    var i= d.getMinutes();
    var s= d.getSeconds();
    return d.getFullYear()+'.'+ ((M<10)?'0'+M:M) +'.'+ ((D<10)?'0'+D:D) +' '+ ((h<10)?'0'+h:h) +':'+ ((i<10)?'0'+i:i);
};



var list = function(page,tab) {
    $('.popovers').each(function(){$(this).popover('hide');});
    NProgress.start();
    var postData={};
    if(!tab){
        tab=$('#tab').val();
    }
    if(tab!='roleresource'&&tab!='role'){
        tab='roleresource';
    }

    switch (tab) {
        case 'roleresource':
            $('#tab_1').attr('class', 'active');
            $('#tab_2').attr('class', '');
            $("#resourceType").parent().attr("hidden", false);
            $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_' + tab + '.php?action=add"> 创建 <i class="fa fa-plus"></i></a>');
            postData['type']=$('#resourceType').val();
            break;
        case 'role':
            $('#tab_1').attr('class', '');
            $('#tab_2').attr('class', 'active');
            $("#resourceType").parent().attr("hidden", true);
            $('#tab_toolbar').html('<div>;</div>');
            break;
    }

    $('#tab').val(tab);
    var url='/api/for_layout/'+tab+'.php';
    if (!page) {
        page = cache.page;
    }else{
        cache.page = page;
    }
    url+='?action=list&page=' + page+'&type='+postData['type'];
    var head = $("#table-head");//数据表头
    var body = $("#table-body");//数据列表
    head.html("<tr><td>Loading ...</td></tr>");
    body.html("");
    delete postData['type'];
    $.ajax({
        type: "POST",
        url: url,
        data: JSON.stringify(postData),
        dataType: "json",
        success: function (listdata) {
            if(listdata.code==0){
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
                $('#tab').val(tab);
                //生成分页
                processPage(listdata, pageinfo, paginate);
                //生成列表
                processBody(listdata, head, body);
                $('.popovers').each(function(){$(this).popover();});
                $('.tooltips').each(function(){$(this).tooltip();});
                switchery();
            }else{
                pageNotify('error','加载失败！','错误信息：'+listdata.msg);
            }
            NProgress.done();
        },
        error: function (){
            pageNotify('error','加载失败！','错误信息：接口不可用');
            NProgress.done();
        }
    });
}

//生成分页
var processPage = function(data,pageinfo,paginate){
    var begin = data.pageSize * ( data.page - 1 ) + 1;
    var end = ( data.count > begin + data.pageSize - 1 ) ? begin + data.pageSize - 1 : data.count;
    pageinfo.html('Showing '+begin+' to '+end+' of '+data.count+' records');
    var p1=(data.page-1>0)?data.page-1:1;
    var p2=data.page+1;
    prev='<li><a href="javascript:;" onclick="list('+p1+')"><i class="fa fa-angle-left"></i></a></li>';
    paginate.append(prev);
    for (var i = 1; i <= data.pageCount; i++) {
        var li='';
        if(i==data.page){
            li='<li class="active"><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
        }else{
            if(i==1||i==data.pageCount){
                li='<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
            }else{
                if(i==p1){
                    if(p1>2){
                        console.log(i+' '+p1);
                        li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
                    }else{
                        li='<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
                    }
                }else{
                    if(i==p2){
                        if(p2<data.pageCount-1){
                            li='<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>'+"\n"+'<li class="disabled"><a href="#">...</a></li>';
                        }else{
                            li='<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
                        }
                    }
                }
            }
        }
        paginate.append(li);
    }
    if(p2>data.pageCount) p2=data.pageCount;
    next='<li class="next"><a href="javascript:;" title="Next" onclick="list('+p2+')"><i class="fa fa-angle-right"></i></a></li>';
    paginate.append(next);
}

//生成列表
var processBody = function(data,head,body){
    var td="";
    if(data.title){
        var tr = $('<tr></tr>');
        for (var i = 0; i < data.title.length; i++) {
            var v = data.title[i];
            td = '<th>' + v + '</th>';
            tr.append(td);
        }
        head.html(tr);
    }
    if(data.content){
        if(data.content.length>0){
            var tab=$('#tab').val();
            for (var i = 0; i < data.content.length; i++) {
                var v = data.content[i];
                var tr = $('<tr></tr>');
                td = '<td>' + v.i + '</td>';
                tr.append(td);
                var btnEdit='',btnDel='',btnView='',btnRun='',t='';
                switch(tab){
                    case 'roleresource':
                        td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'roleresource\',\''+v.id+'\')">' + v.name + '</a></td>';
                        tr.append(td);
                        td = '<td>' + v.resource_type + '</td>';
                        tr.append(td);
                        td = '<td>' + v.user + '</td>';
                        tr.append(td);
                        td = '<td>' + getDate(v.update_time) + '</td>';
                        tr.append(td)
                        var btnEdit = '<a class="btn blue tooltips" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.id + '" style="padding-left: 0px;" title="修改"><i class="fa fa-edit"></i></a>';
                        var btnDel = '<a class="btn red tooltips" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')" style="padding-left: 0px;" title="删除"><i class="fa fa-trash-o"></i></a>';
                        break;
                    case 'role':
                        td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'role\',\''+v.id+'\')">' + v.name + '</a></td>';
                        tr.append(td);
                        td = '<td>' + v.role_file_path + '</td>';
                        tr.append(td);

                        td = '<td>' + v.tasks + '</td>';
                        tr.append(td);
                        td = '<td>' + v.templates + '</td>';
                        tr.append(td);
                        td = '<td>' + getDate(v.update_time) + '</td>';
                        tr.append(td);
                        var btnEdit = '<a class="btn blue tooltips" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.id + '" style="padding-left: 0px;" title="修改"><i class="fa fa-edit"></i></a>';
                        var btnDel = '<a class="btn red tooltips" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')" style="padding-left: 0px;" title="删除"><i class="fa fa-trash-o"></i></a>';
                        break;
                }
                td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnEdit + btnDel + '</div></td>';
                tr.append(td);

                body.append(tr);
            }
        }else{
            pageNotify('info','Warning','数据为空！');
        }
    }else{
        pageNotify('warning','error','接口异常！');
    }
}

//获得所有的资源配置数据
var getResourceData = function(idx){
    var url = '/api/for_layout/roleresource.php?action=list&pagesize=1000';
    var postData = {"action": "list", "pagesize": 1000};
    cache.resourceData={};
    $.ajax({
        type:"POST",
        url:url,
        data:postData,
        dataType:"json",
        success: function(data){
            if(data.code==0){
                get(idx,data.content);
            }
        },
        error: function(){
            pageNotify('error','加载详情失败！','错误信息：接口不可用');
        }

    });
}

var get = function (idx,resourceData) {
    var tab=$('#tab').val();
    var url='/api/for_layout/'+tab+'.php',postData={};
    switch (tab){
        case 'roleresource':
            postData={"action":"info","fIdx":idx};
            break;
        case 'role':
            postData={"action":"info","fIdx":idx};
            break;
    }
    if(idx!='') {
        $.ajax({
            type: "POST",
            url: url,
            data: postData,
            dataType: "json",
            success: function (data) {
                //执行结果提示
                if (data.code == 0) {
                    if (typeof(data.content) != 'undefined') {
                        $('#var_file').html('');
                        $('#template_file').html('');
                        $('#task_file').html('');
                        if(tab=='role'){
                            if (data.content.vars != 'undefined'){
                                // var str = data.content.vars;
                                // var var_array = str.split(",");
                                var var_array = data.content.vars;
                            }

                            if (data.content.tasks != 'undefined') {
                                var task_array = data.content.tasks.split(",");
                            }
                            if (data.content.templates != 'undefined') {
                                var tem_array = data.content.templates.split(",");
                            }
                            if(resourceData.length>0){
                                for(var i=0;i<resourceData.length;i++){
                                    resource_id = resourceData[i].id.toString();
                                    if(resourceData[i].resource_type=='var'){
                                        // if(var_array.indexOf(resource_id)!=-1){
                                        //     var var_checkboxes = '<span class="col-sm-2"><input type="checkbox" checked id="' + resource_id + '" name="var">' + resourceData[i].name + '</span>';
                                        //     $('#var_file').append(var_checkboxes);
                                        // }else{
                                        //     var var_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="' + resource_id + '" name="var">' + resourceData[i].name + '</span>';
                                        //     $('#var_file').append(var_checkboxes);
                                        //
                                        // }

                                        if(var_array == resource_id){
                                            var var_checkboxes = '<span class="col-sm-2"><input type="radio" checked id="' + resource_id + '" name="var">' + resourceData[i].name + '</span>';
                                            $('#var_file').append(var_checkboxes);
                                        }else{
                                            var var_checkboxes = '<span class="col-sm-2"><input type="radio" id="' + resource_id + '" name="var">' + resourceData[i].name + '</span>';
                                            $('#var_file').append(var_checkboxes);

                                        }
                                    }

                                    if(resourceData[i].resource_type=='task'){
                                        // if(task_array.indexOf(resource_id)!=-1){
                                        //     var task_checkboxes = '<span class="col-sm-2"><input type="checkbox" checked id="' + resource_id + '" name="task">' + resourceData[i].name + '</span>';
                                        //     $('#task_file').append(task_checkboxes);
                                        // }else{
                                        //     var task_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="' + resource_id + '" name="task">' + resourceData[i].name + '</span>';
                                        //     $('#task_file').append(task_checkboxes);
                                        //
                                        // }

                                        if(task_array == resource_id){
                                            var task_checkboxes = '<span class="col-sm-2"><input type="radio" checked id="' + resource_id + '" name="task">' + resourceData[i].name + '</span>';
                                            $('#task_file').append(task_checkboxes);
                                        }else{
                                            var task_checkboxes = '<span class="col-sm-2"><input type="radio" id="' + resource_id + '" name="task">' + resourceData[i].name + '</span>';
                                            $('#task_file').append(task_checkboxes);

                                        }


                                    }

                                    if(resourceData[i].resource_type=='template'){
                                        if(tem_array.indexOf(resource_id)!=-1){
                                            var tem_checkboxes = '<span class="col-sm-2"><input type="checkbox" checked id="' + resource_id + '" name="template">' + resourceData[i].name + '</span>';
                                            $('#tem_file').append(tem_checkboxes);
                                        }else{
                                            var tem_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="' + resource_id + '" name="template">' + resourceData[i].name + '</span>';
                                            $('#tem_file').append(tem_checkboxes);

                                        }
                                    }


                                }
                            }
                        }


                        if(data.content.resource_type=='template'){
                            $("#template_show").attr("hidden",false);
                        }else{
                            $("#template_show").attr("hidden",true);
                        }
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
    }else{
        pageNotify('warning', '加载详情失败！', '错误信息：参数错误');
        title = '非法请求 - ' + action;
    }

}

var twiceCheck=function(action,idx,desc){
    NProgress.start();
    if(!idx) idx='';
    if(!desc) desc='';
    var modalTitle='',modalBody='',list='',notice='',btnDisable=false;
    var tab=$('#tab').val();
    if(!action){
        modalTitle='非法请求';
        notice='<div class="note note-danger">错误信息：参数错误</div>';
        pageNotify('error','非法请求！','错误信息：参数错误');
    }else{
        switch(tab){
            case 'roleresource':
                switch(action){
                    case 'del':
                        modalTitle='删除资源配置';
                        modalBody=modalBody+'<div class="form-group col-sm-12">';
                        modalBody=modalBody+'<div class="note note-danger">';
                        modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 资源ID : '+idx +'<br>资源配置名称 : '+desc;
                        modalBody=modalBody+'</div>';
                        modalBody=modalBody+'</div>';
                        modalBody=modalBody+'<input type="hidden" id="id" name="id" value="'+idx+'">';
                        modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="delete">';
                        break;
                    default:
                        modalTitle='非法请求';
                        notice='<div class="note note-danger">错误信息：参数错误</div>';
                        pageNotify('error','非法请求！','错误信息：参数错误');
                        break;
                }
                break;
            case 'role':
                switch(action){
                    case 'del':
                        modalTitle='删除Role配置';
                        modalBody=modalBody+'<div class="form-group col-sm-12">';
                        modalBody=modalBody+'<div class="note note-danger">';
                        modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> RoleID : '+idx +'<br>Role名称 : '+desc;
                        modalBody=modalBody+'</div>';
                        modalBody=modalBody+'</div>';
                        modalBody=modalBody+'<input type="hidden" id="id" name="id" value="'+idx+'">';
                        modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="delete">';
                        break;
                    default:
                        modalTitle='非法请求';
                        notice='<div class="note note-danger">错误信息：参数错误</div>';
                        pageNotify('error','非法请求！','错误信息：参数错误');
                        break;
                }
                break;
            default:
                modalTitle='非法请求';
                notice='<div class="note note-danger">错误信息：参数错误</div>';
                pageNotify('error','非法请求！','错误信息：参数错误');
                break;
        }
    }
    $('#myModalLabel').html(modalTitle);
    $('#myModalBody').html(modalBody);
    if(notice!=''){
        $('#modalNotice').html(notice);
        $('#btnCommit').attr('disabled',true);
    }else{
        $('#btnCommit').attr('disabled',btnDisable);
    }
    NProgress.done();
}

//view
var view=function(type,idx){
    NProgress.start();
    var url='/api/for_layout/'+type+'.php',title='',text='',illegal=false,height='',postData={};
    var tStyle='word-break:break-all;word-warp:break-word;';
    postData={"action":"info","fIdx":idx};
    switch(type){
        case 'roleresource':
            title='查看资源详情 - '+idx;
            break;
        case 'role':
            title='查看Role详情 - '+idx;
            break;
        default:
            illegal=true;
            break;
    }
    if(!illegal&&idx!=''){
        $.ajax({
            type: "POST",
            url: url,
            data: postData,
            dataType: "json",
            success: function (data) {
                //执行结果提示
                if(data.code==0){
                    if(typeof(data.content)!='undefined'){
                        //pageNotify('success','加载成功！');
                        var locale={};
                        switch(type){
                            case 'roleresource':
                                if(typeof(locale_messages.layout.resource)) locale = locale_messages.layout.resource;
                                $.each(data.content,function(k,v){
                                    if(v=='') v='空';
                                    if(typeof(locale[k])!='undefined'){
                                        k=locale[k];
                                    }
                                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                                });
                                break;
                            case 'role':
                                if(typeof(locale_messages.layout.role)) locale = locale_messages.layout.role;
                                $.each(data.content,function(k,v){
                                    if(v=='') v='空';
                                    if(typeof(locale[k])!='undefined') k=locale[k];
                                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                                });
                                break;
                        }
                        if(!text){
                            pageNotify('warning','数据为空！');
                            text='<div class="note note-warning">数据为空！</div>';
                        }
                    }else{
                        pageNotify('warning','数据为空！');
                        text='<div class="note note-warning">数据为空！</div>';
                    }
                }else{
                    pageNotify('error','加载失败！','错误信息：'+data.msg);
                    text='<div class="note note-danger">错误信息：'+data.msg+'</div>';
                }
                setTimeout(function(){
                    if(height!=''){
                        $('#myViewModalBody').css('height',height);
                    }
                    $('#myViewModalLabel').html(title);
                    $('#myViewModalBody').html(text);
                    NProgress.done();
                },200);
            },
            error: function (){
                pageNotify('error','加载失败！','错误信息：接口不可用');
                text='<div class="note note-danger">错误信息：接口不可用</div>';
                $('#myViewModalLabel').html(title);
                $('#myViewModalBody').html(text);
                NProgress.done();
            }
        });
    }else{
        pageNotify('warning','错误操作！','错误信息：参数错误');
        title='非法请求';
        text='<div class="note note-danger">错误信息：参数错误</div>';
        $('#myViewModalLabel').html(title);
        $('#myViewModalBody').html(text);
        NProgress.done();
    }
}
//增删改查
var change=function(){
    NProgress.start();
    var tab=$('#tab').val();
    var url='/api/for_layout/'+tab+'.php';
    var postData={};
    var form=$('#myModalBody').find("input,select,textarea");
    var vars='',templates='',tasks='';
    //处理表单内容--不需要修改
    $.each(form,function(i){
        switch(this.type){
            case 'radio':
                if(this.id){
                    // if(this.name) postData[this.name]=$('input[name="'+this.name+'"]:checked').val();
                    if(this.checked){
                        switch (this.name){
                            case 'var':
                                vars = this.id;
                                break;
                            case 'task':
                                tasks = this.id;
                                break;

                        }
                    }
                }
                break;
            case 'checkbox':
                if(this.id) {
                    if (this.checked) {
                        switch (this.name){
                            // case 'var':
                            //     vars+=this.id+',';
                            //     break;
                            case "template":
                                templates+=this.id+',';
                                break;
                        }
                    }
                }
                    break;
            default:
                if(this.name) postData[this.name]=this.value;
                break;
            }
        });
    var action=$("#page_action").val();
    // vars=vars.substring(0,vars.length-1);
    // tasks=tasks.substring(0,tasks.length-1);
    templates=templates.substring(0,templates.length-1);
    // postData['vars']=vars;
    postData['tasks']=tasks;
    postData['templates']=templates;
    postData['vars']=vars;
    if(postData['resource_type']!='template'){
        postData['template_file_owner']='';
        postData['template_file_path']='';
        postData['template_file_perm']='';
    }

    delete postData['page_action'];
    var actionDesc='';
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
        default:
            actionDesc = action;
            break;
    }
    $.ajax({
        type: "POST",
        url: url,
        data: {"action":action,"data":JSON.stringify(postData)},
        dataType: "json",
        success: function (data) {
            //执行结果提示
            if(data.code==0){
                pageNotify('success','【'+actionDesc+'】操作成功！');
            }else{
                pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+data.msg);
            }
            //重载列表
            list();
            //处理模态框和表单
            $("#myModal :input").each(function () {
                $(this).val("");
            });
            $("#myModal").on("hidden.bs.modal", function() {
                $(this).removeData("bs.modal");
            });
        },
        error: function (){
            pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
        }
    });
}

var hiddenFile=  function(){

    if($("#resource_type").val()=='template'){
        $("#template_show").attr("hidden",false);
    }else{
        $("#template_show").attr("hidden",true);
    }
}