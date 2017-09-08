cache = {
    page: 1,
    page_size: 20,
    cluster: [],
    ecs: [],
    ecs_filter: [],
    ecs_del: {},
    ecs_del_id: [],
    copy: {
        ip: [],
    },
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

var reset = function(){
    $('#fIdx').val('');
    list(1);
}

var getList=function(){
    var url='/api/for_cloud/ecs.php?action=list';
    var postData={'pagesize':1000};
    var actionDesc='ECS列表';
    NProgress.start();
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            NProgress.done();
            if(data.code==0){
                if(typeof data.content != 'undefined') cache.ecs = data.content;
            }else{
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
            }
            list();
        },
        error: function (){
            NProgress.done();
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
            list();
        }
    });
}

var list = function(page,tab) {
    if(!tab) tab='ecs';
    $('.popovers').each(function(){$(this).popover('hide');});
    NProgress.start();
    var data=cache.ecs;
    if (!page) {
        page = cache.page;
    }else{
        cache.page = page;
    }
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
    //检索
    var fIdx=$('#fIdx').val();
    if(fIdx){
        cache.ecs_filter=[];
        $.each(data,function(k,v){
            if(v.PrivateIpAddress.indexOf(fIdx)!=-1||v.Cluster.Name.indexOf(fIdx)!=-1||v.Cluster.Provider.indexOf(fIdx)!=-1)
                cache.ecs_filter.push(v);
        });
        data=cache.ecs_filter;
    }
    //生成分页
    processPage(data, page, pageinfo, paginate);
    //生成列表
    processBody(data, page, head, body);
    $('.popovers').each(function(){$(this).popover();});
    $('.tooltips').each(function(){$(this).tooltip();});
    NProgress.done();
}

//生成分页
var processPage = function(data,page,pageinfo,paginate,func){
    if(!func) func='list';
    var page_size = cache.page_size;
    var count = ($.isArray(data)) ? data.length : 0;
    var page_count = Math.ceil( count / page_size);
    var begin = (count > 0) ? page_size * ( page - 1 ) + 1 : 0;
    var end = ( count > begin + page_size - 1 ) ? begin + page_size - 1 : count;
    pageinfo.html('Showing '+begin+' to '+end+' of '+count+' records');
    var p1=(page-1>0)?page-1:1;
    var p2=page+1;
    var prev='<li><a href="javascript:;" onclick="'+func+'('+p1+')"><i class="fa fa-angle-left"></i></a></li>';
    paginate.append(prev);
    for (var i = 1; i <= page_count; i++) {
        var li='';
        if(i==page){
            li='<li class="active"><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
        }else{
            if(i==1||i==page_count){
                li='<li><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
            }else{
                if(i==p1){
                    if(p1>2){
                        li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
                    }else{
                        li='<li><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
                    }
                }else{
                    if(i==p2){
                        if(p2<page_count-1){
                            li='<li><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>'+"\n"+'<li class="disabled"><a href="#">...</a></li>';
                        }else{
                            li='<li><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
                        }
                    }
                }
            }
        }
        paginate.append(li);
    }
    if(p2>page_count) p2=page_count;
    next='<li class="next"><a href="javascript:;" title="Next" onclick="'+func+'('+p2+')"><i class="fa fa-angle-right"></i></a></li>';
    paginate.append(next);
}

//生成列表
var processBody = function(data,page,head,body){
    var td="";
    var title=[ '<input type="checkbox" name="SelectAll" id="SelectAll" onclick="checkAll(this)">', '#', 'Ecs_Id', 'IP', '机型模板', '云厂商', '创建时间', '状态', '#'];
    if(title){
        var tr = $('<tr></tr>');
        for (var i = 0; i < title.length; i++) {
            var t='';
            if(title[i]=='IP') t='<a class="pull-right tooltips" data-container="body" data-trigger="hover" data-original-title="复制整列" data-toggle="modal" data-target="#myViewModal" onclick="copy(\'ip\')"><i class="fa fa-copy"></i></a>';
            td = '<th>' + title[i] + t + '</th>';
            tr.append(td);
        }
        head.html(tr);
    }
    if(data.length>0){
        cache.copy.ip=[];
        var j = 0;
        for (var i = 0; i < data.length; i++) {
            if(i<(page-1)*cache.page_size) continue;
            if(i>=page*cache.page_size) break;
            j++;
            var v = data[i];
            if(v.PrivateIpAddress) cache.copy.ip.push(v.PrivateIpAddress);
            if(v.PublicIpAddress) cache.copy.ip.push(v.PublicIpAddress);
            var tr = $('<tr></tr>');
            td = '<td><input type="checkbox" id="list" name="list[]" value="'+ v.InstanceId + ':' + ((v.PrivateIpAddress)?v.PrivateIpAddress: v.PublicIpAddress) +':' + v.Status +'" /></td>';
            tr.append(td);
            td = '<td>' + j + '</td>';
            tr.append(td);
            var btnAdd='',btnEdit='',btnDel='';
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'ecs\',\''+v.InstanceId+'\')">' + v.InstanceId + '</a></td>';
            tr.append(td);

            var ipShow = (v.PrivateIpAddress ? (v.PrivateIpAddress + " [内网]<br/>") : "") + (v.PublicIpAddress ? (v.PublicIpAddress + " [公网]") : "");

            td = '<td>' + ipShow + '</td>';
            tr.append(td);
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'cluster\',\''+v.Cluster.Id+'\')">' + v.Cluster.Name + '</a></td>';
            tr.append(td);
            td = '<td>' + v.Cluster.Provider + '</td>';
            tr.append(td);
            td = '<td>' + getDate(v.CreateTime) + '</td>';
            tr.append(td);
            td = '<td>' + getStatusAlias(v.Status) + '</td>';
            tr.append(td);
            btnEdit = '<a class="tooltips" title="查看创建日志" data-toggle="modal" data-target="#myViewModal" onclick="viewLog(\''+ v.InstanceId +'\',\''+ ((v.PrivateIpAddress)?v.PrivateIpAddress: v.PublicIpAddress) +'\')"><i class="fa fa-history"></i></a>';
            btnEdit += (v.Status==0)?'':' <a class="text-primary tooltips" title="下载私钥" target="_blank" href="/api/for_cloud/ecs.php?action=down_sshkey&fIdx='+ ((v.PrivateIpAddress)?v.PrivateIpAddress: v.PublicIpAddress) +'"><i class="fa fa-download"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.InstanceId+'\',\''+((v.PrivateIpAddress)?v.PrivateIpAddress: v.PublicIpAddress)+'\',\''+ v.Status +'\')"><i class="fa fa-trash-o"></i></a>';
            td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnEdit + ' ' + btnDel + '</div></td>';
            tr.append(td);

            body.append(tr);
        }
    }else{
        pageNotify('info','Warning','数据为空！');
    }
}

//增删改查
var change=function(step){
    NProgress.start();
    var tab=$('#tab').val(),modal='myModal';
    var page_other=$('#page_other').val();
    var url='/api/for_cloud/'+tab+'.php';
    var postData={};
    if(step) modal='myStepModal';
    var form=$('#'+modal+'Body').find("input,select,textarea");

    //处理表单内容--不需要修改
    $.each(form,function(i){
        switch(this.type){
            case 'radio':
                if(typeof(postData[this.name])=='undefined'){
                    if(this.name) postData[this.name]=$('input[name="'+this.name+'"]:checked').val();
                }
                break;
            case 'checkbox':
                if(this.id){
                    if(typeof(postData[this.id])=='undefined'){
                        postData[this.id]={};
                    }
                    if(this.checked){
                        postData[this.id][i]=this.value;
                    }
                }
                break;
            default:
                if(this.name) postData[this.name]=this.value;
                break;
        }
    });

    var opts = $("#sql_labels").find("option");
    var labels='',sql_labels='';
    if(postData['user_label']!=''){
        labels=postData['user_label'];
    }
    $.each(opts,function(i){
        if(this.selected){
            sql_labels += (sql_labels=='')?this.value:','+this.value;
        }
    });
    labels+=(labels=='')?sql_labels:((sql_labels=='')?'':","+sql_labels);
    postData['label']=labels;
    var action=$("#page_action").val();
    delete postData['page_action'];
    delete postData['page_other'];
    var actionDesc='';
    switch(action){
        case 'insert':
            actionDesc='添加';
            break;
        case 'delete':
            postData=JSON.parse(postData.ecs);
            actionDesc='删除';
            break;
        case 'addPhyDev':
            actionDesc='添加';
            break;
        default:
            actionDesc=action;
            break;
    }
    $.ajax({
        type: "POST",
        url: url,
        data: {"action":action,"data":JSON.stringify(postData)},
        dataType: "json",
        success: function (data) {
            NProgress.done();
            //执行结果提示
            if(data.code==0){
                pageNotify('success','【'+actionDesc+'】操作成功！');
            }else{
                if(action=='insert'&&data.msg=='timeout'){
                    pageNotify('info','【'+actionDesc+'】操作已提交！','机器创建中...请稍后!');
                }else{
                    pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+((typeof data.msg == 'object')?JSON.stringify(data.msg):data.msg));
                }
            }
            //重载列表
            getList();
            //处理模态框和表单
            $('#'+modal+' :input').each(function () {
                $(this).val("");
            });
            $('#'+modal).on("hidden.bs.modal", function() {
                $(this).removeData("bs.modal");
            });
            $('#'+modal).modal('hide');
        },
        error: function (){
            NProgress.done();
            pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
            getList();
        }
    });
}

//view
var view=function(type,idx){
    NProgress.start();
    var url='',title='',text='',illegal=false,height='',postData={};
    var tStyle='word-break:break-all;word-warp:break-word;';
    url='/api/for_cloud/'+type+'.php';
    switch(type){
        case 'cluster':
            title='查看机型模板详情';
            postData={"action":"info","fIdx":idx};
            break;
        case 'ecs':
            title='查看ECS详情';
            postData={"action":"info","fIdx":idx};
            break;
        default:
            NProgress.done();
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
                        var locale={},str='';
                        switch(type){
                            case 'cluster':
                                if(typeof(locale_messages.cloud)) locale = locale_messages.cloud;
                                $.each(data.content,function(k,v){
                                    switch (k){
                                        case 'Network':
                                            str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                                                '<h5 class="col-sm-12 text-primary"><strong>'+((typeof locale[k])?locale[k]:'网络选项')+'</strong></h5>';
                                            $.each(v,function(key,val){
                                                if(locale[key]!=false){
                                                    if(val=='') val='空';
                                                    if(typeof(locale[key])!='undefined') key=locale[key];
                                                    str+='<span class="title col-sm-2" style="font-weight: bold;">'+key+'</span> <span class="col-sm-4" style="'+tStyle+'">'+val+'</span>'+"\n";
                                                }
                                            });
                                            break;
                                        case 'Zone':
                                            str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                                                '<h5 class="col-sm-12 text-primary"><strong>'+((typeof locale[k])?locale[k]:'区域选项')+'</strong></h5>';
                                            $.each(v,function(key,val){
                                                if(locale[key]!=false){
                                                    if(val=='') val='空';
                                                    if(typeof(locale[key])!='undefined') key=locale[key];
                                                    str+='<span class="title col-sm-2" style="font-weight: bold;">'+key+'</span> <span class="col-sm-4" style="'+tStyle+'">'+val+'</span>'+"\n";
                                                }
                                            });
                                            break;
                                        default:
                                            if(locale[k]!=false){
                                                if(v=='') v='空';
                                                if(typeof(locale[k])!='undefined') k=locale[k];
                                                text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                                            }
                                            break;
                                    }
                                });
                                text+=str;
                                break;
                            case 'ecs':
                                if(typeof(locale_messages.cloud)) locale = locale_messages.cloud;
                                $.each(data.content,function(k,v){
                                    if(locale[k]!=false){
                                        if(v=='') v='空';
                                        if(typeof(locale[k])!='undefined') k=locale[k];
                                        text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                                    }
                                });
                                text+=str;
                                break;
                        }
                        if(!text){
                            pageNotify('warning','无可展示数据！');
                            text='<div class="note note-warning">无可展示数据！</div>';
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
        title='非法请求 - '+action;
        text='<div class="note note-danger">错误信息：参数错误</div>';
        $('#myViewModalLabel').html(title);
        $('#myViewModalBody').html(text);
        NProgress.done();
    }
}

var check=function(){
    var disabled=false;
    if($('#ClusterId').val()=='') disabled=true;
    if($('#number').val()=='') disabled=true;
    $('#btnCommit').attr('disabled',false)
}

var twiceCheck=function(action,idx,desc,status){
    NProgress.start();
    if(!idx) idx='';
    if(!desc) desc='';
    var modalTitle='',modalBody='',list='',notice='',btnDisable=false;
    if(!action){
        modalTitle='非法请求';
        notice='<div class="note note-danger">错误信息：参数错误</div>';
        pageNotify('error','非法请求！','错误信息：参数错误');
    }else{
        switch(action){
            case 'del':
                cache.ecs_del={};
                modalTitle='删除机器';
                var count = 0, postDel = [], ecsId='', ecsIp ='', ecsStatus='', ecsIps='';
                if(idx){
                    status=parseInt(status);
                    count++;
                    $('input:checkbox[value="'+idx+':'+desc+':'+status+'"]').attr('checked','true');
                    if(!desc) desc=idx;
                    if(desc) ecsIps=(ecsIps)?','+desc:desc;
                    cache.ecs_del[desc]=idx;
                    if((status>0&&status<5)||status==7 || status == 8){
                        postDel.push(idx);
                        list+='<span class="col-sm-3">'+desc+'</span>';
                    }else{
                        notice='只能删除"未初始化","初始化中","初始化超时","初始化完成","初始化失败"的机器!';
                        list+='<span class="col-sm-3">'+desc+'('+getStatusAlias(status)+')</span>';
                        count--;
                    }
                }else{
                    $('input:checkbox[id=list]:checked').each(function(i){
                        count++;
                        var ecs=$(this).val().split(':');
                        ecsId=ecs[0];
                        ecsIp=ecs[1];
                        ecsStatus=parseInt(ecs[2]);
                        if(ecsIp) {
                            ecsIps+=(ecsIps)?','+ecsIp:ecsIp;
                            cache.ecs_del[ecsIp]=ecsId;
                        }else{
                            ecsIp=ecsId;
                        }
                        if(ecsStatus>0&&ecsStatus<5){
                            postDel.push(ecsId);
                            list+='<span class="col-sm-3" id="check_'+ecsIp+'">'+ecsIp+'</span>';
                        }else{
                            list+='<span class="col-sm-3 text-success" id="check_'+ecsIp+'">'+ecsIp+'('+getStatusAlias(ecsStatus)+')</span>';
                            notice='只能删除"未初始化","初始化中","初始化超时","初始化完成"的机器!';
                            count--;
                        }
                    });
                }
                modalBody+='<div class="form-group col-sm-12">';
                modalBody+='<h5><strong>当前操作, 将影响如下列表:</strong></h5>';
                modalBody+='<p><strong class="text-primary">涉及总数</strong>: 共 <span class="badge badge-danger">'+count+'</span> 个 </p>';
                modalBody+='<p><strong class="text-primary">涉及列表</strong>:';
                modalBody+='<div class="col-sm-12">'+list+'</div>';
                modalBody+='</div>';
                modalBody+='<div class="col-sm-12"><div class="alert alert-danger" id="del_notice"></div></div>';
                cache.ecs_del_id=postDel;
                modalBody+='<textarea class="hidden" id="ecs" name="ecs">'+JSON.stringify(postDel)+'</textarea>'
                modalBody+='<input type="hidden" id="page_action" name="page_action" value="delete">';
                btnDisable=true;
                if(count>0){
                    if(ecsIps){
                        btnDisable=false;
                        checkEcsState(ecsIps,notice);
                    }else{
                        btnDisable=false;
                    }
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
        $('#del_notice').html(notice);
        $('#btnCommit').attr('disabled',true);
    }else{
        $('#btnCommit').attr('disabled',btnDisable);
    }
    NProgress.done();
}

var updateSelect=function(name,idx){
    var tSelect=$('#'+name),data='';
    switch(name){
        case 'ClusterId':
            data=cache.cluster;
            break;
    }
    tSelect.empty();
    tSelect.append('<option value="">请选择</option>');
    if(data.length>0){
        $.each(data,function(k,v){
            tSelect.append('<option value="' + v.Id + '">' + v.Name + '</option>');
        });
    }
}

var get = function (idx) {
    var tab=$('#tab').val(),postData={};
    url='/api/for_cloud/'+tab+'.php';
    switch (tab){
        case 'service':
            postData={"action":"info","fIdx":idx};
            break;
        case 'pool':
            postData={"action":"info","fIdx":idx};
            break;
    }
    if(idx!=''){
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
                        $.each(data.content,function(k,v){
                            if($('#'+k).length>0){
                                switch ($('#'+k).get(0).tagName){
                                    case 'INPUT':
                                        switch ($('#'+k).attr('type')){
                                            case 'radio':
                                                $("input[name='"+k+"'][value='"+v+"']").attr("checked",true);
                                                break;
                                            case 'checkbox':
                                                $.each(v,function(k1,v1){
                                                    $("input[id='"+k+"']:checkbox[value='"+v1+"']").attr('checked','true');
                                                });
                                                break;
                                            default:
                                                $('#'+k).val(v);
                                                break;
                                        }
                                        break;
                                    case 'SELECT':
                                        if($('#'+k).find("option[value='"+v+"']").length==0){
                                            $('#'+k).append('<option value="' + v + '">' + v + '</option>');
                                        }
                                        $('#'+k).find("option[value='"+v+"']").attr("selected",true);
                                        break;
                                    default:
                                        $('#'+k).val(v);
                                        break;
                                }
                            }
                        });
                    }else{
                        pageNotify('warning','数据为空！');
                    }
                }else{
                    pageNotify('error','加载失败！','错误信息：'+data.msg);
                }
            },
            error: function (){
                pageNotify('error','加载详情失败！','错误信息：接口不可用');
            }
        });
    }else{
        pageNotify('warning','加载详情失败！','错误信息：参数错误');
        title='非法请求 - '+action;
    }
}

var getName=function(type,idx){
    var name=idx;
    switch (type){
        case 'organization':
            data=cache.organization;
            $.each(data,function(k,v){
                if(v.Id==idx){
                    name= v.Name;
                }
            });
            break;
        case 'cluster':
            data=cache.cluster;
            $.each(data,function(k,v){
                if(v.Id==idx){
                    name= v.Name;
                }
            });
            break;
    }
    return name;
}

var checkAll=function(o){
    $('[id=list]:checkbox').prop('checked', o.checked);
}

var copy=function(col){
    NProgress.start();
    var url='',title='查看整列 - '+col,text='';
    var str='';
    switch (col){
        case 'ip':
            $.each(cache.copy.ip, function (k,v) {
                str+=(str)?"\n"+v:v;
            });
            break;
    }
    text+='<textarea class="form-control" rows="10">'+str+'</textarea>';
    $('#myViewModalLabel').html(title);
    $('#myViewModalBody').html(text);
    NProgress.done();
    if(!str){
        pageNotify('info','没有可复制的数据');
    }
}

var getStatusAlias = function(status){
    var str=status;
    switch(status){
        case 0:
            str='<span class="badge">创建中</span>';break;
        case 1:
            str='<span class="badge bg-green">初始化完成</span>';break;
        case 2:
            str='<span class="badge bg-blue">未初始化</span>';break;
        case 3:
            str='<span class="badge bg-blue-sky">初始化中</span>';break;
        case 4:
            str='<span class="badge bg-orange">初始化超时</span>';break;
        case 5:
            str='<span class="badge bg-blue">机器已删除</span>';break;
        case 6:
            str='<span class="badge bg-purple">机器删除中</span>';break;
        case 7:
            str='<span class="badge bg-red">初始化失败</span>';break;
        case 8:
            str='<span class="badge bg-green">录入服务池成功</span>';break;
        default:
            str='<span class="badge">未知状态</span>';break;
    }
    return str;
}

var getCluster=function(){
    var actionDesc="机型模板",tSelect='ClusterId';
    var url='/api/for_cloud/cluster.php?action=list';
    var postData={"pagesize":1000};
    cache.cluster = [];
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (typeof data.content != 'undefined') cache.cluster = data.content;
            updateSelect(tSelect);
            if(data.code!=0){
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
            }else{
                if(data.content.length==0) pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
            }
        },
        error: function (){
            updateSelect(tSelect);
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
        }
    });
}

var isNumberValid=function(o,a,b,d){
    var v= o.value;
    var flag=false;
    var ex = /^\d+$/;
    if (ex.test(v)) {
        if(a&&parseInt(v)<a) flag=true;
        if(b&&parseInt(v)>b) flag=true;
    }else{
        flag=true;
    }
    if(flag){
        if(d){
            o.value=d;
        }else{
            o.value='';
        }
    }
    check();
}

var checkEcsState=function(ips,flag){
    var actionDesc='校验机器使用状态',notice='',count=0;
    var url='/api/for_layout/pool.php?action=state';
    var postData={"action":"state","fIdx":ips};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if(data.code!=0){
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
                notice=data.msg;
            }else{
                if(data.content){
                    $.each(data.content,function(k,v){
                        if( v >=0 ){
                            count++;
                            if($('#btnCommit').attr('disabled')!='true') $('#btnCommit').attr('disabled',true);
                            $("[id='check_"+k+"']").addClass('text-danger').html(k+'(使用中)');
                            //移除选中机器
                            var id=cache.ecs_del[k];
                            cache.ecs_del_id.splice($.inArray(id,cache.ecs_del_id),1);
                            $('#ecs').val(JSON.stringify(cache.ecs_del_id));
                        }
                    });
                    if(count>0){
                        notice='有<span class="badge bg-green">'+count+'</span>台正在使用中(如上标红), 禁止直接删除, 请使用服务编排的缩容!';
                    }else{
                        if(!flag){
                            notice='机器使用状态校验结果显示没有正在使用中的机器, 确认删除?';
                            $('#del_notice').attr('class','alert alert-success');
                        }
                    }
                }else{
                    pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
                    notice='机器使用状态校验返回结果数据为空, 请谨慎操作!';
                }
            }
            if(flag) notice=(notice)?flag+'<br>'+notice:flag;
            $('#del_notice').html(notice);
        },
        error: function (){
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
        }
    });
}

var viewLog=function(idx,desc){
    NProgress.start();
    var url='',title='查看创建 - ',text='',postData={};
    title+=(desc)?desc:idx;
    var tStyle='word-break:break-all;word-warp:break-word;';
    postData={"action":"log","fIdx":idx};
    url='/api/for_cloud/ecs.php';
    if(idx!=''){
        $.ajax({
            type: "POST",
            url: url,
            data: postData,
            dataType: "json",
            success: function (data) {
                //执行结果提示
                if(data.code==0){
                    var result = (typeof data.content != 'undefined') ? data.content.replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/\n/g,'<br>') : '';
                    if(!result) result='日志数据为空';
                    text+='<span class="col-sm-12" style="background-color:#000;color:#ccc;line-height: 150%">'+ result +'</span>';
                }else{
                    pageNotify('error','加载失败！','错误信息：'+data.msg);
                    text='<div class="note note-danger">错误信息：'+data.msg+'</div>';
                }
                setTimeout(function(){
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




var checkPhyDev=function(){
    var disabled=false;
    if($('#InstanceList').val()=='') disabled=true;
    $('#btnCommit').attr('disabled',disabled)
}

var getPoolAndLabel = function(){
    var actionDesc = '服务池列表';
    var postData = {"action": "poolList", "pagesize": 1000};
    $.ajax({
        type:"POST",
        url: '/api/for_layout/pool.php',
        data:postData,
        dataType:"json",
        success: function(data){
            if(data.code==0){
                $('#pools').empty();
                if(data.content.length>0){
                    $.each(data.content,function(k,v){
                        $('#pools').append('<option value="'+v.id+'">'+v.name+'</option>').select2({width: '100%'});
                    });
                }
            }else{
                pageNotify('error','加载'+actionDesc+"失败",'错误信息：'+data.msg);
            }
        },
        error: function(){
            pageNotify('error', '加载' + actionDesc + '失败！', '错误信息：接口不可用');
        }
    });

    var actionDesc_l = '标签列表';
    var postDataLabel = {"page": 1,'pagesize':1000};
    $.ajax({
        type:"POST",
        url: '/api/for_layout/node.php?action=labelList',
        data:{"data": JSON.stringify(postDataLabel)},
        dataType:"json",
        success: function(data){
            if(data.code==0){
                $('#sql_labels').empty();
                var label_array=[];
                $.each(data.content,function(k,v){
                    if(/,/.test(v)){
                        var tmp_labels = v.split(",");
                        for(var i=0;i<tmp_labels.length;i++){
                            if(label_array.length == 0){
                                label_array.push(tmp_labels[i]);
                            }else{
                                if(label_array.indexOf(tmp_labels[i])==-1){
                                    label_array.push(tmp_labels[i]);
                                }
                            }
                        }
                    }else{
                        if(label_array.length == 0){
                            label_array.push(v);
                        }else{
                            if(label_array.indexOf(v)==-1){
                                label_array.push(v);
                            }
                        }

                    }
                });
                if(label_array.length>0){
                    $.each(label_array,function(k,v){
                        if(v!=""){
                            $('#sql_labels').append('<option value="'+v+'">'+v+'</option>').select2({width: '100%'});
                        }
                    });
                }
            }else{
                pageNotify('error','加载'+actionDesc_l+"失败",'错误信息：'+data.msg);
            }
        },
        error: function(){
            pageNotify('error', '加载' + actionDesc_l + '失败！', '错误信息：接口不可用');
        }
    });


}

var getLabel=function(){
    $(".js-example-tags").select2({
        tags: true
    })
}


var addLabels=function(){
    $("#user_label").parent().attr("hidden",false);
}