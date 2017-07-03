cache = {
    page: 1,
    page_size: 20,
    account: [],
    account_filter: [],
}

var reset = function(){
    $('#fIdx').val('');
    list(1);
}

var getList=function(){
    var url='/api/for_cloud/account.php?action=list';
    var postData={'pagesize':1000};
    var actionDesc='account列表';
    NProgress.start();
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            NProgress.done();
            if(data.code==0){
                if(typeof data.content != 'undefined') cache.account = data.content;
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
    if(!tab) tab='account';
    $('.popovers').each(function(){$(this).popover('hide');});
    NProgress.start();
    var data=cache.account;
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
        cache.account_filter=[];
        $.each(data,function(k,v){
            if(v.Provider.indexOf(fIdx)!=-1||v.KeyId.indexOf(fIdx)!=-1)
                cache.account_filter.push(v);
        });
        data=cache.account_filter;
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
    var title=['#', '云厂商', '云账号id', '云账号密码', '体验机使用额度', '体验机总额度', '账号操作'];
    if(title){
        var tr = $('<tr></tr>');
        for (var i = 0; i < title.length; i++) {
            td = '<th>' + title[i] + '</th>';
            tr.append(td);
        }
        head.html(tr);
    }
    if(data.length>0){
        var j = 0;
        for (var i = 0; i < data.length; i++) {
            if(i<(page-1)*cache.page_size) continue;
            if(i>=page*cache.page_size) break;
            j++;
            var v = data[i];
            var tr = $('<tr></tr>');
            td = '<td>' + j + '</td>';
            tr.append(td);
            var btnAdd='',btnEdit='',btnDel='';
            td = '<td>' + v.Provider + '</td>';
            tr.append(td);
            td = '<td>' + v.KeyId + '</td>';
            tr.append(td);
            var hideSecret = '';
            for(var h = 0; h < v.KeySecret.length; h++)hideSecret +='*';
            td = '<td>' + hideSecret +'</td>';
            tr.append(td);
            var theCredit = Math.round(parseFloat(v.Credit)*10, 1)/10.0;
            var theSpent = Math.round(parseFloat(v.Spent)*10, 1)/10.0;
            if(theSpent > theCredit) theSpent = theCredit;
            td = '<td>' + theSpent + '</td>';
            tr.append(td);
            td = '<td>' + theCredit + '</td>';
            tr.append(td);
            btnEdit = '<a class="tooltips" title="修改账号" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'alt\','+ v.Id +',\''+ v.KeyId +'\',\''+ v.KeySecret +'\',\''+ v.Provider +'\')"><i class="fa fa-edit"></i></a>';
            td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnEdit + '</div></td>';
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
    var action=$("#page_action").val();
    delete postData['page_action'];
    delete postData['page_other'];
    postData['id'] = parseInt(postData['id']);
    var actionDesc='';
    switch(action){
        case 'update':
            actionDesc='修改';
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
                pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+((typeof data.msg == 'object')?JSON.stringify(data.msg):data.msg));
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

var twiceCheck=function(action,idx,desc, secret,Provider){
    NProgress.start();
    if(!idx) idx=0;
    if(!desc) desc='';
    if(!secret) secret='';
    if(!Provider) Provider='';
    var modalTitle='',modalBody='',list='',notice='',btnDisable=false;
    if(!action){
        modalTitle='非法请求';
        notice='<div class="note note-danger">错误信息：参数错误</div>';
        pageNotify('error','非法请求！','错误信息：参数错误');
    }else{
        switch(action){
            case 'alt':
                modalTitle='修改云账号';
                modalBody+='<div class="col-sm-11">';
                modalBody+='<div class="form-group">';
                modalBody+='<label for="hours" class="col-sm-4 control-label">云厂商账号</label>';
                modalBody+='<div class="col-sm-7">';
                modalBody+='<input type="text" class="form-control" id="KeyId" name="KeyId" onkeyup="inputchange()" value= \''+ desc + '\' placeholder="云厂商账号">';
                modalBody+='</div>';
                modalBody+='</div>';
                modalBody+='<div class="form-group">';
                modalBody+='<label for="hours" class="col-sm-4 control-label">云账号密码</label>';
                modalBody+='<div class="col-sm-7">';
                modalBody+='<input type="password" class="form-control" id="KeySecret" name="KeySecret" onkeyup="inputchange()" value=\''+ secret + '\'  placeholder="请输入云账号密码">';
                modalBody+='</div>';
                modalBody+='</div>';
                modalBody+='</div>';
                modalBody+='<input type="hidden" id="provider" name="provider" value="'+Provider+'">';
                modalBody+='<input type="hidden" id="id" name="id" value='+idx+'>';
                modalBody+='<input type="hidden" id="page_action" name="page_action" value="update">';
                btnDisable=true;
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

var inputchange=function(input){
    $('#btnCommit').attr('disabled',false);
}