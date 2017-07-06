cache = {
    page: 1,
    page_size: 20,
    cluster: [],
    cluster_filter: [],
    provider: [],
    charge: [],
    region: [],
    zone: [],
    vpc: [],
    subnet: [],
    security_group: [],
    ecs_type: [],
    image: [],
    disk: [],
    quota:{},
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
    var url='/api/for_cloud/cluster.php?action=list';
    var postData={'pagesize':1000};
    var actionDesc='机型模板';
    NProgress.start();
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            NProgress.done();
            if(data.code==0){
                if(typeof data.content != 'undefined') cache.cluster = data.content;
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
    console.log(new Date());
    if(!tab) tab='cluster';
    $('.popovers').each(function(){$(this).popover('hide');});
    NProgress.start();
    var data=cache.cluster;
    $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myStepModal" href="edit_'+tab+'.php?action=add"> 创建 <i class="fa fa-plus"></i></a>');
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
        cache.cluster_filter=[];
        $.each(data,function(k,v){
            if(v.Name.indexOf(fIdx)!=-1||v.Provider.indexOf(fIdx)!=-1)
                cache.cluster_filter.push(v);
        });
        data=cache.cluster_filter;
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
    var td="",tab=$('#tab').val();
    var title=[ '#', '名称', '云厂商', '配额用量/余量', '创建时间', '#'];
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
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'cluster\',\''+v.Id+'\')">' + v.Name + '</a></td>';
            tr.append(td);
            td = '<td>' + v.Provider + '</td>';
            tr.append(td);
            td = '<td><span id="quota_'+ v.Id +'">0</span><a class="text-success pull-right tooltips" title="追加配额" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'quota\',\''+v.Id+'\',\''+v.Name+'\')"><i class="fa fa-edit"></i></a></td>';
            tr.append(td);
            td = '<td>' + getDate(v.CreateTime) + '</td>';
            tr.append(td);
            getQuota(v.Id);
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.Id+'\',\''+v.Name+'\')"><i class="fa fa-trash-o"></i></a>';
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
    var actionDesc='';
    switch(action){
        case 'insert':
            actionDesc='添加';
            postData.Network={};
            postData.Zone={};
            if($('#Provider').val()=='aliyun') {
                switch (postData.NetworkType) {
                    case 'common':
                        if (typeof postData.InternetChargeType != 'undefined') {
                            postData.Network.InternetChargeType = postData.InternetChargeType;
                            delete postData.InternetChargeType;
                        }
                        if (typeof postData.InternetMaxBandwidthOut != 'undefined') {
                            postData.Network.InternetMaxBandwidthOut = parseInt(postData.InternetMaxBandwidthOut);
                            delete postData.InternetMaxBandwidthOut;
                        }
                        delete postData.VpcId;
                        delete postData.SubnetId;
                        break;
                    case 'custom':
                        if (typeof postData.VpcId != 'undefined') {
                            postData.Network.VpcId = postData.VpcId;
                            delete postData.VpcId;
                        }
                        if (typeof postData.SubnetId != 'undefined') {
                            postData.Network.SubnetId = postData.SubnetId;
                            delete postData.SubnetId;
                        }
                        delete postData.InternetChargeType;
                        delete postData.InternetMaxBandwidthOut;
                        break;
                }
                if (typeof postData.SecurityGroupId != 'undefined') {
                    postData.Network.SecurityGroupId = postData.SecurityGroupId;
                    delete postData.SecurityGroupId;
                }
                if(typeof postData.ZoneName != 'undefined'){
                    postData.Zone.ZoneName=postData.ZoneName;
                    delete postData.ZoneName;
                }
                if(typeof postData.RegionName != 'undefined'){
                    postData.Zone.RegionName=postData.RegionName;
                    delete postData.RegionName;
                }
                if(typeof postData.DataDiskNum != 'undefined') postData.DataDiskNum=parseInt(postData.DataDiskNum);
                if(typeof postData.DataDiskSize != 'undefined') postData.DataDiskSize=parseInt(postData.DataDiskSize);

            }else if($('#Provider').val()=='openstack'){
                postData.Network.VpcId = postData.NetworkOP;
                postData.Zone.ZoneName=postData.AvalibilityZone;
                postData.InstanceType=postData.DiskType;

                delete postData.DataDiskCategory
                delete postData.SystemDiskCategory
                delete postData.SecurityGroupId
                delete postData.SubnetId
                delete postData.VpcId
                delete postData.DataDiskNum;
                delete postData.DataDiskSize;
                delete postData.InternetChargeType;
                delete postData.InternetMaxBandwidthOut;
                delete postData.ZoneName;
                delete postData.RegionName;
            }
            delete postData.NetworkOP;
            delete  postData.DiskType;
            delete postData.AvalibilityZone;
            delete postData['NetworkType'];

            break;
        case 'update':
            url='/api/for_cloud/quota.php';
            actionDesc='追加配额';
            break;
        case 'delete':
            actionDesc='删除';
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
            //执行结果提示
            if(data.code==0){
                pageNotify('success','【'+actionDesc+'】操作成功！');
            }else{
                pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+data.msg);
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
    title='查看详情';
    postData={"action":"info","fIdx":idx};
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
                        var locale={};
                        if(typeof(locale_messages.cloud)) locale = locale_messages.cloud;
                        var str='';
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
        title='非法请求 - '+action;
        text='<div class="note note-danger">错误信息：参数错误</div>';
        $('#myViewModalLabel').html(title);
        $('#myViewModalBody').html(text);
        NProgress.done();
    }
}

//commit check
var check=function(tab){
    if(!tab) tab=$('#tab').val();
    switch(tab){
        case 'cluster':
            var disabled=false;
            var network=$('#NetworkType').val();
            var InternetChargeType=$('#InternetChargeType').val();
            var InternetMaxBandwidthOut=$('#InternetMaxBandwidthOut').val();
            var DataDiskSize=$('#DataDiskSize').val();
            var DataDiskNum=$('#DataDiskNum').val();
            if($('#Name').val()=='') disabled=true;
            if($('#Provider').val()=='') disabled=true;
            if($('#Provider').val()=='aliyun') {
                if (network == '') disabled = true;
                switch (network) {
                    case 'common':
                        if (InternetChargeType == '') disabled = true;
                        if (InternetMaxBandwidthOut == '') disabled = true;
                        switch (InternetChargeType) {
                            case 'PayByBandwidth':
                                if (InternetMaxBandwidthOut < 1 || InternetMaxBandwidthOut > 100) $('#InternetMaxBandwidthOut').val(1);
                                break;
                            case 'PayByTraffic':
                                if (InternetMaxBandwidthOut < 0 || InternetMaxBandwidthOut > 100) $('#InternetMaxBandwidthOut').val(0);
                                break;
                        }
                        break;
                    case 'custom':
                        if ($('#VpcId').val() == '') disabled = true;
                        if ($('#SubnetId').val() == '') disabled = true;
                        break;
                }
                if ($('#RegionName').val() == '') disabled = true;
                if ($('#ZoneName').val() == '') disabled = true;
                if ($('#InstanceType').val() == '') disabled = true;
                if ($('#ImageId').val() == '') disabled = true;
                if ($('#SystemDiskCategory').val() == '') disabled = true;
                if (DataDiskSize == '' || DataDiskSize < 5 || DataDiskSize > 32768) disabled = true;
                if (DataDiskNum == '' || DataDiskNum < 1 || DataDiskNum > 4) disabled = true;
                if ($('#DataDiskCategory').val() == '') disabled = true;
            }else if($('#Provider').val()=='openstack'){
                if($('#NetworkOP').val() == '') disabled =true;
                if($('#AvabilityZone').val() == '') disabled=true
                if($('#ImageId').val()=='') disabled=true
                if($('#DiskType').val()=='') disabled=true
            }
            $("#btnStepCommit").attr('disabled',disabled);
            break;
        case 'quota':
            var disabled=false;
            if($('#hours').val()=='') disabled=true;
            $("#btnCommit").attr('disabled',disabled);
            break;
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
            case 'cluster':
                switch(action){
                    case 'quota':
                        modalTitle='追加配额';
                        var credit=(typeof cache.quota[idx].credit != 'undefined') ? cache.quota[idx].credit : 0;
                        modalBody+='<div class="form-group">' +
                            '<label for="name" class="col-sm-2 control-label">模板名称</label>' +
                            '<div class="col-sm-10">' +
                            '<input type="text" class="form-control" id="name" name="name" value="'+desc+'" readonly>' +
                            '</div>' +
                            '</div>' +
                            '<div class="form-group">' +
                            '<label for="hours" class="col-sm-2 control-label">追加配额(h)</label>' +
                            '<div class="col-sm-10">' +
                            '<input type="number" class="form-control" id="hours" name="hours" onkeyup="isNumberValid(this,0,99999999,1)" value="1" min="1" max="99999999" placeholder="小时数,eg:1">' +
                            '<span class="help">当前配额余量: <span id="hours_help" class="badge bg-green">'+credit+'</span>小时, 追加配额将累加.</span>' +
                            '</div>' +
                            '</div>';
                        modalBody+='<input type="hidden" id="id" name="id" value="'+idx+'">';
                        modalBody+='<input type="hidden" id="page_action" name="page_action" value="update">';
                        break;
                    case 'del':
                        modalTitle='删除类型';
                        modalBody+='<div class="form-group col-sm-12">';
                        modalBody+='<div class="note note-danger">';
                        modalBody+='<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 类型 : '+idx+'<br>名称 : '+desc;
                        modalBody+='</div>';
                        modalBody+='</div>';
                        modalBody+='<input type="hidden" id="id" name="id" value="'+idx+'">';
                        modalBody+='<input type="hidden" id="page_action" name="page_action" value="delete">';
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
//*********
var updateSelect=function(name,idx){
    var tSelect=$('#'+name),data='';
    switch(name){
        case 'Provider':
            data=cache.provider;
            break;
        case 'InternetChargeType':
            data=cache.charge;
            break;
        case 'RegionName':
            data=cache.region;
            break;
        case 'ZoneName':
            data=cache.zone;
            break;
        case 'VpcId':
            data=cache.vpc;
            break;
        case 'NetworkOP':
            data=cache.vpc;
            break;
        case 'SubnetId':
            data=cache.subnet;
            break;
        case 'SecurityGroupId':
            data=cache.security_group;
            break;
        case 'InstanceType':
            data=cache.ecs_type;
            break;
        case 'DiskType':
            data=cache.ecs_type;
            break;
        case 'ImageId':
            data=cache.image;
            break;
        case 'SystemDiskCategory':
        case 'DataDiskCategory':
            data=cache.disk;
            break;
    }
    tSelect.empty();
    if(data.length>0){
        tSelect.append('<option value="">请选择</option>');
        switch(name){
            case 'InternetChargeType':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.name + '">' + v.name + '</option>');
                });
                break;
            case 'RegionName':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.RegionName + '">' + v.RegionName + '</option>');
                });
                break;
            case 'ZoneName':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.ZoneName + '">' + ((v.ZoneName)?v.ZoneName:v.ZoneId) + '</option>');
                });
                break;
            case 'VpcId':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.VpcId + '">' + v.VpcId + '</option>');
                });
                break;
            case 'NetworkOP':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.VpcId + '">' + v.State + '</option>');
                });
                break;
            case 'SubnetId':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.SubnetId + '">' + v.CidrBlock + '</option>');
                });
                break;
            case 'SecurityGroupId':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.GroupId + '">' + ((v.GroupName)?v.GroupName:v.GroupId) + '</option>');
                });
                break;
            case 'InstanceType':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.name + '">' + v.name + '</option>');
                });
                break;
            case 'DiskType':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    //传来的字符串由#分隔，第一段是flavorID，第二段是flavorName
                    strs=v.name.split("#")
                    tSelect.append('<option value="' + strs[0] + '">' + strs[1] + '</option>');
                })
                break;
            case 'ImageId':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.ImageId + '">' + ((v.Description)?'【'+v.Description+'】'+v.ImageId:v.ImageId) + '</option>');
                });
                break;
            case 'SystemDiskCategory':
            case 'DataDiskCategory':
                tSelect.removeAttr("disabled");
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.name + '">' + v.name + '</option>');
                });
                break;
            default:
                $.each(data,function(k,v){
                    tSelect.append('<option value="' + v.id + '">' + ((v.name)?v.name:v.id) + '</option>');
                });
                break;
        }
        if(idx){
            tSelect.val(idx).trigger('change');
        }else{
            tSelect.val($('#'+name+' option:nth-child(1)').val()).trigger('change');
        }
    }else{
        tSelect.append('<option value="">请选择</option>');
    }
    $("select.form-control").select2({width:'100%'});
}

var get = function (idx) {
    var tab=$('#tab').val(),postData;
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

var getQuota=function(idx){
    cache.quota[idx]={};
    var url='/api/for_cloud/quota.php?action=info';
    var postData={fIdx:idx};
    var actionDesc='配额';
    NProgress.start();
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            NProgress.done();
            if(data.code==0){
                $('#quota_'+idx).html('<a class="tooltips" title="用量">'+data['content']['costs'] + '</a>/' +
                    '<a class="tooltips" title="余量">'+data['content']['credit']+'</a>');
                cache.quota[idx]=data['content'];
            }else{
                $('#quota_'+idx).html(data.msg);
            }
            $('.tooltips').each(function(){$(this).tooltip();});
        },
        error: function (){
            NProgress.done();
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
        }
    });
}

var getProvider=function(){
    var actionDesc="云厂商",tSelect='Provider';
    var url='/api/for_cloud/provider.php?action=list';
    var postData={"pagesize":1000};
    cache.provider = [];
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (typeof data.content != 'undefined') cache.provider = data.content;
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

var getDiskAndNetwork=function(){
    check();
    getDisk();
    getRegion();
    getNetworkType();
}

var getNetworkType=function(){
    switch ($('#NetworkType').val()){
        case 'common':
            if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#InternetChargeType').parent().parent().removeClass('hidden');
            if($('#InternetMaxBandwidthOut').parent().parent().attr('class').indexOf('hidden')!=-1) $('#InternetMaxBandwidthOut').parent().parent().removeClass('hidden');
            if($('#VpcId').parent().parent().attr('class').indexOf('hidden')==-1) $('#VpcId').parent().parent().addClass('hidden');
            if($('#SubnetId').parent().parent().attr('class').indexOf('hidden')==-1) $('#SubnetId').parent().parent().addClass('hidden');
            getChargeType();
            break;
        case 'custom':
            if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')==-1) $('#InternetChargeType').parent().parent().addClass('hidden');
            if($('#InternetMaxBandwidthOut').parent().parent().attr('class').indexOf('hidden')==-1) $('#InternetMaxBandwidthOut').parent().parent().addClass('hidden');
            if($('select[name="VpcId"]').parent().parent().attr('class').indexOf('hidden')!=-1) $('select[name="VpcId"]').parent().parent().removeClass('hidden');
            if($('#SubnetId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#SubnetId').parent().parent().removeClass('hidden');
            getVpcId();
            break;
    }
    check();
}

var getChargeType=function(){
    var actionDesc="网络计费类型",tSelect='InternetChargeType';
    var url='/api/for_cloud/charge.php?action=list';
    var idx=$('#Provider').val();
    if(!idx) return false;
    var postData={"pagesize":1000,"fIdx":idx};
    cache.charge = [];
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (typeof data.content != 'undefined') cache.charge = data.content;
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

var getRegion=function(){
    var actionDesc="可用地域",tSelect='RegionName';
    var url='/api/for_cloud/region.php?action=list';
    var idx=$('#Provider').val();
    if(!idx) return false;
    var postData={"pagesize":1000,"fIdx":idx};
    cache.region = [];
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (typeof data.content != 'undefined') cache.region = data.content;
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

var getDisk=function(){
    var actionDesc="磁盘类型";
    var url='/api/for_cloud/disk_category.php?action=list';
    var idx=$('#Provider').val();
    if(!idx) return false;
    var postData={"pagesize":1000,"fIdx":idx};
    cache.disk = [];
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (typeof data.content != 'undefined') cache.disk = data.content;
            updateSelect('SystemDiskCategory');
            updateSelect('DataDiskCategory');
            if(data.code!=0){
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
            }else{
                if(data.content.length==0) pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
            }
        },
        error: function (){
            updateSelect('SystemDiskCategory');
            updateSelect('DataDiskCategory');
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
        }
    });
}

var getZoneAndVpcAndETypeAndImage=function(){
    check();
    getInstanceType();
    getImage();
    getZoneId();
    if($('#NetworkType').val()=='custom') getVpcId();
}

var getInstanceType=function(){
    var actionDesc="机器规格";
    var idx=$('#RegionName').val();
    var provider=$('#Provider').val();
    var tSelect = 'InstanceType';
    if(provider == 'aliyun'){
        if(!provider||!idx) return false;
    }else if(provider == 'openstack'){
        tSelect = 'DiskType';
    }
    var url='/api/for_cloud/ecs_type.php?action=list';
    var postData={"pagesize":1000,"fProvider":provider,"fIdx":idx};
    cache.ecs_type = [];
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (typeof data.content != 'undefined') cache.ecs_type = data.content;
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

var getImage=function(){
    var actionDesc="镜像",tSelect='ImageId';
    var url='/api/for_cloud/image.php?action=list';
    var provider=$('#Provider').val();
    var idx=$('#RegionName').val();
    if($('#Provider').val()=='aliyun') {
        if (!provider || !idx) return false;
    }else if($('#Provider').val()=='openstack'){
        //openstack未对idx作要求，此处是象征性地传递
        idx=1;
    }
    var postData={"pagesize":1000,"fProvider":provider,"fIdx":idx};
    cache.image = [];
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            if (typeof data.content != 'undefined') cache.image = data.content;
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

var getZoneId=function(){
    var actionDesc="可用区",tSelect='ZoneName';
    var url='/api/for_cloud/zone.php?action=list';
    var idx=$('#RegionName').val();
    if(!idx) return false;
    var postData={"pagesize":1000,"fIdx":idx};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            cache.zone = (typeof data.content != 'undefined') ? data.content : [];
            updateSelect(tSelect);
            if(data.code!=0){
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
            }else{
                if(data.content.length==0) pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
            }
        },
        error: function (){
            cache.zone = [];
            updateSelect(tSelect);
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
        }
    });
}
//******对openstack的情况有所修改
var getVpcId=function(){
    var actionDesc="可用区";
    var provider=$('#Provider').val();
    var url='/api/for_cloud/vpc.php?action=list';
    if($('#Provider').val()=="aliyun") {
        tSelect='VpcId';
        var idx = $('#RegionName').val();
        if (!idx) return false;
    }else if($('#Provider').val()=='openstack'){
        tSelect='NetworkOP'
        var idx=1;
    }
    var postData={"pagesize":1000,"fProvider":provider,"fIdx":idx};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            cache.vpc = (typeof data.content != 'undefined') ? data.content : [];
            updateSelect(tSelect);
            if(data.code!=0){
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
            }else{
                if(data.content.length==0) pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
            }
        },
        error: function (){
            cache.vpc = [];
            updateSelect(tSelect);
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
        }
    });
}

var getVSwitchAndSecurityGroup=function(){
    check();
    if($('#NetworkType').val()=='custom') getVSwitchId();
    getSecurityGroup();
}

var getVSwitchId=function(){
    var actionDesc="子网",tSelect='SubnetId';
    var url='/api/for_cloud/subnet.php?action=list';
    var zone=$('#ZoneName').val();
    var idx=$('#VpcId').val();
    if(!zone||!idx) return false;
    var postData={"pagesize":1000,fZone:zone,"fIdx":idx};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            cache.subnet = (typeof data.content != 'undefined') ? data.content : [];
            updateSelect(tSelect);
            if(data.code!=0){
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
            }else{
                if(data.content.length==0) pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
            }
        },
        error: function (){
            cache.subnet = [];
            updateSelect(tSelect);
            pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
        }
    });
}

var getSecurityGroup=function(){
    var actionDesc="安全组",tSelect='SecurityGroupId';
    var url='/api/for_cloud/security_group.php?action=list';
    var region=$('#RegionName').val(),network=$('#NetworkType').val();
    var idx=$('#VpcId').val();
    if(network!='custom'){
        idx='';
    }else{
        if(!idx) return false;
    }
    if(!region) return false;
    var postData={"pagesize":1000,fRegion:region,"fIdx":idx};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            cache.security_group = (typeof data.content != 'undefined') ? data.content : [];
            updateSelect(tSelect);
            if(data.code!=0){
                pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
            }else{
                if(data.content.length==0) pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
            }
        },
        error: function (){
            cache.security_group = [];
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

var selectProvider=function () {
    switch ($('#Provider').val()){
        case 'aliyun':
            switchToAliyun();
            getDiskAndNetwork();
            break;
        case 'openstack':
            switchToOpenStack();
            getOpenStackAttr();
            break;
    }
}

var getOpenStackAttr=function () {
    getImage();
    getInstanceType();
    getVpcId()
}

var switchToAliyun=function () {
    //显示
    if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#InternetChargeType').parent().parent().removeClass('hidden');
    if($('#InternetMaxBandwidthOut').parent().parent().attr('class').indexOf('hidden')!=-1) $('#InternetMaxBandwidthOut').parent().parent().removeClass('hidden');
    if($('#VpcId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#VpcId').parent().parent().removeClass('hidden');
    if($('#SubnetId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#SubnetId').parent().parent().removeClass('hidden');
    if($('#NetworkType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#NetworkType').parent().parent().removeClass('hidden');
    if($('#RegionName').parent().parent().attr('class').indexOf('hidden')!=-1) $('#RegionName').parent().parent().removeClass('hidden');
    if($('#ZoneName').parent().parent().attr('class').indexOf('hidden')!=-1) $('#ZoneName').parent().parent().removeClass('hidden');
    if($('#SecurityGroupId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#SecurityGroupId').parent().parent().removeClass('hidden');
    if($('#InstanceType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#InstanceType').parent().parent().removeClass('hidden');
    if($('#SystemDiskCategory').parent().parent().attr('class').indexOf('hidden')!=-1) $('#SystemDiskCategory').parent().parent().removeClass('hidden');
    if($('#DataDiskCategory').parent().parent().attr('class').indexOf('hidden')!=-1) $('#DataDiskCategory').parent().parent().removeClass('hidden');
    if($('#DataDiskSize').parent().parent().attr('class').indexOf('hidden')!=-1) $('#DataDiskSize').parent().parent().removeClass('hidden');
    if($('#DataDiskNum').parent().parent().attr('class').indexOf('hidden')!=-1) $('#DataDiskNum').parent().parent().removeClass('hidden');
    //隐藏
    if($('#NetworkOP').parent().parent().attr('class').indexOf('hidden')==-1) $('#NetworkOP').parent().parent().addClass('hidden');
    if($('#AvalibilityZone').parent().parent().attr('class').indexOf('hidden')==-1) $('#AvalibilityZone').parent().parent().addClass('hidden');
    if($('#DiskType').parent().parent().attr('class').indexOf('hidden')==-1) $('#DiskType').parent().parent().addClass('hidden');
}

var switchToOpenStack=function () {
    //隐藏
    if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')==-1) $('#InternetChargeType').parent().parent().addClass('hidden');
    if($('#InternetMaxBandwidthOut').parent().parent().attr('class').indexOf('hidden')==-1) $('#InternetMaxBandwidthOut').parent().parent().addClass('hidden');
    if($('#VpcId').parent().parent().attr('class').indexOf('hidden')==-1) $('#VpcId').parent().parent().addClass('hidden');
    if($('#SubnetId').parent().parent().attr('class').indexOf('hidden')==-1) $('#SubnetId').parent().parent().addClass('hidden');
    if($('#NetworkType').parent().parent().attr('class').indexOf('hidden')==-1) $('#NetworkType').parent().parent().addClass('hidden');
    if($('#RegionName').parent().parent().attr('class').indexOf('hidden')==-1) $('#RegionName').parent().parent().addClass('hidden');
    if($('#ZoneName').parent().parent().attr('class').indexOf('hidden')==-1) $('#ZoneName').parent().parent().addClass('hidden');
    if($('#SecurityGroupId').parent().parent().attr('class').indexOf('hidden')==-1) $('#SecurityGroupId').parent().parent().addClass('hidden');
    if($('#InstanceType').parent().parent().attr('class').indexOf('hidden')==-1) $('#InstanceType').parent().parent().addClass('hidden');
    if($('#SystemDiskCategory').parent().parent().attr('class').indexOf('hidden')==-1) $('#SystemDiskCategory').parent().parent().addClass('hidden');
    if($('#DataDiskCategory').parent().parent().attr('class').indexOf('hidden')==-1) $('#DataDiskCategory').parent().parent().addClass('hidden');
    if($('#DataDiskSize').parent().parent().attr('class').indexOf('hidden')==-1) $('#DataDiskSize').parent().parent().addClass('hidden');
    if($('#DataDiskNum').parent().parent().attr('class').indexOf('hidden')==-1) $('#DataDiskNum').parent().parent().addClass('hidden');
    //显示

    if($('#NetworkOP').parent().parent().attr('class').indexOf('hidden')!=-1) $('#NetworkOP').parent().parent().removeClass('hidden');
    if($('#AvalibilityZone').parent().parent().attr('class').indexOf('hidden')!=-1) $('#AvalibilityZone').parent().parent().removeClass('hidden');
    if($('#DiskType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#DiskType').parent().parent().removeClass('hidden');

}