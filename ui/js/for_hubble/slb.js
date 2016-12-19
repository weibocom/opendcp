cache = {
  page: 1,
  page_size: 20,
  region_id: '',
  region: [],
  zone: [],
  vpc: [],
  subnet: [],
  security_group: [],
  https: [],
  slb: [],
  slb_filter: [],
  slb_info: {},
  slb_desc: {},
  listener:[],
  listener_filter: [],
  listener_info: {},
  backend:[],
  backend_filter: [],
  backend_info: [],
  backend_list: [],
  copy: {
    ip: [],
  },
  ip: [], //选中IP列表
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

var list = function(page,tab) {
  $('.popovers').each(function(){$(this).popover('hide');});
  NProgress.start();
  var data=[];
  if(!tab) tab=$('#tab').val();
  if(tab!='slb'&&tab!='listener'&&tab!='backend'){
    tab='slb';
  }
  switch(tab){
    case 'slb':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','hidden');
      $('#tab_3').attr('class','hidden');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php"> 创建 <i class="fa fa-plus"></i></a>');
      $('#fRegion').parent().parent().attr('hidden',false);
      $('#fSlb').parent().parent().attr('hidden',true);
      data=cache.slb;
      break;
    case 'listener':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','active');
      $('#tab_3').attr('class','hidden');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myStepModal" href="edit_slb_'+tab+'.php"> 创建 <i class="fa fa-plus"></i></a>');
      $('#fRegion').parent().parent().attr('hidden',true);
      $('#fSlb').parent().parent().attr('hidden',false);
      data=cache.listener;
      break;
    case 'backend':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','hidden');
      $('#tab_3').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myTableModal" onclick="getBackend()"> 添加 <i class="fa fa-plus"></i></a>' +
        '<a type="button" class="btn btn-primary" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'update\')"> 批量调整 <i class="fa fa-exchange"></i></a>' +
        '<a type="button" class="btn btn-danger" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\')"> 批量删除 <i class="fa fa-minus"></i></a>');
      $('#fRegion').parent().parent().attr('hidden',true);
      $('#fSlb').parent().parent().attr('hidden',false);
      data=cache.backend;
      break;
  }
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
    switch(tab){
      case 'slb':
        cache.slb_filter=[];
        $.each(data,function(k,v){
          if(v.LoadBalancerName.indexOf(fIdx)!=-1||v.Address.indexOf(fIdx)!=-1)
            cache.slb_filter.push(v);
        });
        data=cache.slb_filter;
        break;
      case 'listener':
        cache.listener_filter=[];
        $.each(data,function(k,v){
          if(v.ListenerPort.toString().indexOf(fIdx)!=-1||v.ListenerProtocol.indexOf(fIdx)!=-1)
            cache.listener_filter.push(v);
        });
        data=cache.listener_filter;
        break;
      case 'backend':
        cache.backend_filter=[];
        $.each(data,function(k,v){
          if(v.ServerId.indexOf(fIdx)!=-1||v.Address.indexOf(fIdx)!=-1)
            cache.backend_filter.push(v);
        });
        data=cache.backend_filter;
        break;
    }
  }
  //生成分页
  processPage(data, page, pageinfo, paginate);
  //生成列表
  processBody(data, page, head, body);
  $('.popovers').each(function(){$(this).popover();});
  $('.tooltips').each(function(){$(this).tooltip();});
  switchery();
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
  var title=[];
  var tab=$('#tab').val();
  switch(tab){
    case 'slb':
      title=["#","名称","服务地址","服务端口","Listener","Backend","实例状态","类型","创建时间","#"];
      break;
    case 'listener':
      title=["#","SLB名称","协议","端口","后端端口","状态","#"];
      break;
    case 'backend':
      title=["#",'<input type="checkbox" name="SelectAll" id="SelectAll" onclick="checkAll(this)">',"SLB名称","MemberId","服务地址","权重","状态","#"];
      break;
  }
  if(title){
    var tr = $('<tr></tr>');
    for (var i = 0; i < title.length; i++) {
      var v = title[i];
      var t='';
      if(title[i]=='IP') t='<a class="pull-right tooltips" data-container="body" data-trigger="hover" data-original-title="复制整列" data-toggle="modal" data-target="#myViewModal" onclick="copy(\'ip\')"><i class="fa fa-copy"></i></a>';
      td = '<th>' + v + t + '</th>';
      tr.append(td);
    }
    head.html(tr);
  }
  if(data.length>0){
    var tab=$('#tab').val();
    cache.copy.ip=[];
    cache.listener_info={};
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
      switch(tab){
        case 'slb':
          td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'slb\',\''+v.LoadBalancerId+'\')">' + ((v.LoadBalancerName)?v.LoadBalancerName:'空') + '</a></td>';
          tr.append(td);
          td = '<td>' + v.Address + '</td>';
          tr.append(td);
          td = '<td id="wan_port_'+ j +'"></td>';
          tr.append(td);
          td = '<td><a class="tooltips" title="查看服务端口" onclick="getList(\'listener\',\''+ v.LoadBalancerId +'\')"><i class="fa fa-bars"></i></a></td>';
          tr.append(td);
          td = '<td><a class="tooltips" title="查看后端节点" onclick="getList(\'backend\',\''+ v.LoadBalancerId +'\')"><i class="fa fa-bars"></i></a></td>';
          tr.append(td);
          if(v.LoadBalancerStatus=='active'){
            td = '<td><div><label class="tooltips" title="点击停用" onclick="return false;"><input type="checkbox" class="js-switch" checked onchange="switchs(\'inactive\',\'' + v.LoadBalancerId + '\')"/> 启用</label></div></td>';
          }else{
            td = '<td><div><label class="tooltips" title="点击启用" onclick="return false;"><input type="checkbox" class="js-switch" onchange="switchs(\'active\',\'' + v.LoadBalancerId + '\')"/> 停用</label></div></td>';
          }
          tr.append(td);
          td = '<td>' + ((v.AddressType=='intranet')?'内网':'公网') + '</td>';
          tr.append(td);
          td = '<td>' + getDate(v.CreateTime) + '</td>';
          tr.append(td);
          getSlbPort(v.LoadBalancerId,j);
          btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.LoadBalancerId+'\',\''+v.LoadBalancerName+'\')"><i class="fa fa-trash-o"></i></a>';
          break;
        case 'listener':
          td = '<td><a class="tooltips" title="查看SLB详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'slb\',\''+cache.slb_info.LoadBalancerId+'\')">' + ((cache.slb_info.LoadBalancerName)?cache.slb_info.LoadBalancerName:cache.slb_info.LoadBalancerId) + '</a></td>';
          tr.append(td);
          td = '<td>' + v.ListenerProtocol + '</td>';
          tr.append(td);
          td = '<td><a class="tooltips" title="查看Listener详情" data-toggle="modal" data-target="#myViewModal" onclick="viewDesc('+j+')">' + v.ListenerPort + '</a></td>';
          tr.append(td);
          td = '<td id="listener_backend_port_'+j+'"></td>';
          tr.append(td);
          td = '<td id="listener_status_'+j+'"></td>';
          tr.append(td);
          getDesc('listener',j,cache.slb_info.LoadBalancerId, v.ListenerProtocol, v.ListenerPort);
          btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myStepModal" href="edit_slb_'+tab+'.php?action=edit&protocol=' + v.ListenerProtocol + '&port=' + v.ListenerPort + '"><i class="fa fa-edit"></i></a>';
          btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+cache.slb_info.LoadBalancerId+'\',\''+v.ListenerProtocol+'\',\''+v.ListenerPort+'\')"><i class="fa fa-trash-o"></i></a>';
          break;
        case 'backend':
          cache.backend_list.push(v.ServerId);
          td = '<td><input type="checkbox" id="list" name="list[]" value="'+ v.ServerId +'" /></td>';
          tr.append(td);
          td = '<td><a class="tooltips" title="查看SLB详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'slb\',\''+cache.slb_info.LoadBalancerId+'\')">' + ((cache.slb_info.LoadBalancerName)?cache.slb_info.LoadBalancerName:cache.slb_info.LoadBalancerId) + '</a></td>';
          tr.append(td);
          td = '<td>' + v.ServerId + '</td>';
          tr.append(td);
          td = '<td id="backend_ip_'+v.ServerId+'"></td>';
          tr.append(td);
          td = '<td><input id="weight_' + v.ServerId + '" type="number" value="'+ v.Weight +'" max=100 min=1 onkeyup="isNumberValid(this,1,100,100)" onchange="isNumberValid(this,1,100,100)"></td>';
          tr.append(td);
          td = '<td id="backend_status_'+v.ServerId+'"></td>';
          tr.append(td);
          btnEdit = '<a class="text-success tooltips" title="权重" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'update\',\'' + v.ServerId + '\')"><i class="fa fa-eject"></i></a>';
          btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\'' + v.ServerId + '\')"><i class="fa fa-trash-o"></i></a>';
          break;
      }
      td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnEdit + ' ' + btnDel + '</div></td>';
      tr.append(td);

      body.append(tr);
    }
    if(tab=='backend') getDesc('backend','',cache.slb_info.LoadBalancerId);
  }else{
    pageNotify('info','Warning','数据为空！');
  }
}

//增删改查
var change=function(step){
  var tab=$('#tab').val(),modal='myModal';
  var page_other=$('#page_other').val();
  var url='/api/for_cloud/slb_'+tab+'.php';
  if(tab=='slb') url='/api/for_cloud/'+tab+'.php';
  var fIdx='',postData={};
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
  switch(tab){
    case 'slb':
      if(postData.NetworkType=='common'){
        delete postData.ZoneId;
        delete postData.VpcId;
        delete postData.VSwitchId;
        delete postData.SecurityGroup;
      }
      if(postData.AddressType=='intranet'){
        if(typeof postData.InternetChargeType != 'undefined') delete postData.InternetChargeType;
        if(typeof postData.Bandwidth != 'undefined') delete postData.Bandwidth;
      }else{
        if(typeof postData.Bandwidth != 'undefined') postData.Bandwidth=parseInt(postData.Bandwidth);
      }
      if(typeof postData.RegionId != 'undefined') cache.region_id=postData.RegionId;
      delete postData.NetworkType;
      break;
    case 'listener':
      $.each(postData,function(k,v){
        if(!v) delete postData[k];
      });
      var intField=[
        "Bandwidth","ListenerPort","BackendServerPort","PersistenceTimeout","HealthCheckConnectPort","HealthCheckTimeout",
        "HealthCheckInterval","UnhealthyThreshold","HealthyThreshold","CookieTimeout"
      ];
      $.each(intField,function(k,v){
        if(typeof postData[v] != 'undefined') postData[v]=parseInt(postData[v]);
      });
      break;
    case 'backend':
      fIdx=postData.LoadBalancerId;
      postData=JSON.parse(postData.backend);
      break;
  }
  delete postData['page_action'];
  delete postData['page_other'];
  //console.log("action="+action);
  //console.log(JSON.stringify(postData));
  var actionDesc='';
  switch(action){
    case 'insert':
      actionDesc='创建';
      break;
    case 'update':
      actionDesc='修改';
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
    data: {"action":action,"fIdx":fIdx,"data":JSON.stringify(postData)},
    dataType: "json",
    success: function (data) {
      //执行结果提示
      if(data.code==0){
        pageNotify('success','【'+actionDesc+'】操作成功！');
      }else{
        pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+data.msg);
      }
      //重载列表
      switch(tab){
        case 'slb':
          getList('slb');
          break;
        default:
          getInfo();
          break;
      }
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
    }
  });
}

//view
var view=function(type,idx){
  NProgress.start();
  var url='',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  url='/api/for_cloud/'+type+'.php';
  postData={"action":"info","fIdx":idx};
  switch(type){
    case 'slb':
      title='查看详情 - '+idx;
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
            if(typeof(locale_messages.cloud)) locale = locale_messages.cloud;
            switch(type){
              default:
                $.each(data.content,function(k,v){
                  if(locale[k]!=false){
                    if(v=='') v='空';
                    if(typeof(locale[k])!='undefined') k=locale[k];
                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                  }
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

//commit check
var check=function(tab){
  if(!tab) tab=$('#tab').val();
  switch(tab){
    case 'slb':
      var disabled=false;
      var LoadBalancerName=$('#LoadBalancerName').val();
      var RegionId=$('#RegionId').val();
      var ZoneId=$('#ZoneId').val();
      var NetworkType=$('#NetworkType').val();
      var AddressType=$('#AddressType').val();
      var InternetChargeType=$('#InternetChargeType').val();
      var Bandwidth=$('#Bandwidth').val();
      var VpcId=$('#VpcId').val();
      var VSwitchId=$('#VSwitchId').val();
      var SecurityGroup=$('#SecurityGroup').val();

      if(LoadBalancerName=='') disabled=true;
      if(RegionId=='') disabled=true;
      if(AddressType=='') disabled=true;
      switch(NetworkType){
        case 'common':
          if(AddressType=='intranet') disabled=true;
          if($('#Bandwidth').parent().parent().attr('class').indexOf('hidden')!=-1) $('#Bandwidth').parent().parent().removeClass('hidden');
          break;
        case 'custom':
          if(ZoneId=='') disabled=true;
          if(AddressType=='internet'){
            if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#InternetChargeType').parent().parent().removeClass('hidden');
            if(InternetChargeType=='paybybandwidth'){
              if($('#Bandwidth').parent().parent().attr('class').indexOf('hidden')!=-1) $('#Bandwidth').parent().parent().removeClass('hidden');
            }else{
              if($('#Bandwidth').parent().parent().attr('class').indexOf('hidden')==-1) $('#Bandwidth').parent().parent().addClass('hidden');
            }
          }else{
            if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')==-1) $('#InternetChargeType').parent().parent().addClass('hidden');
            if($('#Bandwidth').parent().parent().attr('class').indexOf('hidden')==-1) $('#Bandwidth').parent().parent().addClass('hidden');
          }
          if(AddressType=='internet'&&InternetChargeType=='') disabled=true;
          if(AddressType=='internet'&&Bandwidth=='') disabled=true;
          if(VpcId=='') disabled=true;
          if(VSwitchId=='') disabled=true;
          if(SecurityGroup=='') disabled=true;
          break;
        default:
          disabled=true;
          break;
      }
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'listener':
      var disabled=false;
      var LoadBalancerId=$('#fSlb').val();
      $('#LoadBalancerId').val(LoadBalancerId);
      var Protocol=$('#Protocol').val();
      var Scheduler=$('#Scheduler').val();
      var ListenerPort=$('#ListenerPort').val();
      var BackendServerPort=$('#BackendServerPort').val();
      var PersistenceTimeout=$('#PersistenceTimeout').val();
      var ServerCertificateId=$('#ServerCertificateId').val();
      var Bandwidth=$('#Bandwidth').val();
      var XForwarderFor=$('#XForwarderFor').val();
      var HealthCheck=$('#HealthCheck').val();
      var HealthCheckConnectPort=$('#HealthCheckConnectPort').val();
      var HealthCheckTimeout=$('#HealthCheckTimeout').val();
      var HealthCheckInterval=$('#HealthCheckInterval').val();
      var UnhealthyThreshold=$('#UnhealthThreshold').val();
      var HealthyThreshold=$('#HealthThreshold').val();
      var HealthCheckDomain=$('#HealthCheckDomain').val();
      var HealthCheckURI=$('#HealthCheckURI').val();
      var HealthCheckHttpCode=$('#HealthCheckHttpCode').val();
      var StickySession=$('#StickySession').val();
      var StickySessionType=$('#StickySessionType').val();
      var CookieTimeout=$('#CookieTimeout').val();
      var Cookie=$('#Cookie').val();

      if(LoadBalancerId=='') disabled=true;
      if(Protocol=='') disabled=true;
      if(Scheduler=='') disabled=true;
      if(ListenerPort=='') disabled=true;
      if(BackendServerPort=='') disabled=true;
      if(Protocol=='tcp'||Protocol=='udp'){
        if(PersistenceTimeout=='') disabled=true;
      }
      if(Protocol=='https'&&ServerCertificateId=='') disabled=true;
      if(Bandwidth=='') disabled=true;
      if(XForwarderFor=='') disabled=true;
      if(HealthCheck=='') disabled=true;
      if(HealthCheck=='on'){
        if(HealthCheckConnectPort=='') disabled=true;
        if(HealthCheckTimeout=='') disabled=true;
        if(HealthCheckInterval=='') disabled=true;
        if(UnhealthyThreshold=='') disabled=true;
        if(HealthyThreshold=='') disabled=true;
        if(Protocol=='http'||Protocol=='https'){
          //if(HealthCheckDomain=='') disabled=true;
          if(HealthCheckURI=='') disabled=true;
          if(HealthCheckHttpCode=='') disabled=true;
        }
      }
      if(StickySession=='on'){
        switch(StickySessionType){
          case 'insert':
            if(CookieTimeout=='') disabled=true;
            break;
          case 'server':
            if(Cookie=='') disabled=true;
            break;
          default:
            disabled=true;
            break;
        }
      }
      $("#btnStepCommit").attr('disabled',disabled);
      break;
    case 'backend':
      var disabled=false;
      if($('#unit_id').val()=='') disabled=true;
      if($('#ip').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
  }
  $("select.form-control").select2({width:'100%'});
}

var checkProtocol=function(){
  if($('#Protocol').val()=='https') {
    if($('#ServerCertificateId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#ServerCertificateId').parent().parent().removeClass('hidden');
    if($('#PersistenceTimeout').parent().parent().attr('class').indexOf('hidden')==-1) $('#PersistenceTimeout').parent().parent().addClass('hidden');
    $('#StickySession').attr('disabled',false);
    getCertificate();
  }else{
    if($('#Protocol').val()=='http') {
      $('#StickySession').attr('disabled',false);
      if($('#PersistenceTimeout').parent().parent().attr('class').indexOf('hidden')==-1) $('#PersistenceTimeout').parent().parent().addClass('hidden');
    }else{
      $('#StickySession').attr('disabled',true);
      if($('#PersistenceTimeout').parent().parent().attr('class').indexOf('hidden')!=-1) $('#PersistenceTimeout').parent().parent().removeClass('hidden');
    }
    if($('#ServerCertificateId').parent().parent().attr('class').indexOf('hidden')==-1) $('#ServerCertificateId').parent().parent().addClass('hidden');
  }
  checkHealthCheck();
  checkStickySession();
}
var checkAccessControlStatus=function(){
  if($('#AccessControlStatus').val()=='close'){
    if($('#SourceItems').parent().parent().attr('class').indexOf('hidden')==-1) $('#SourceItems').parent().parent().addClass('hidden');
  }else{
    if($('#SourceItems').parent().parent().attr('class').indexOf('hidden')!=-1) $('#SourceItems').parent().parent().removeClass('hidden');
  }
}
var checkHealthCheck=function(){
  if($('#HealthCheck').val()=='on'){
    if($('#HealthCheckConnectPort').parent().parent().attr('class').indexOf('hidden')!=-1) $('#HealthCheckConnectPort').parent().parent().removeClass('hidden');
    if($('#HealthCheckTimeout').parent().parent().attr('class').indexOf('hidden')!=-1) $('#HealthCheckTimeout').parent().parent().removeClass('hidden');
    if($('#HealthCheckInterval').parent().parent().attr('class').indexOf('hidden')!=-1) $('#HealthCheckInterval').parent().parent().removeClass('hidden');
    if($('#UnhealthyThreshold').parent().parent().attr('class').indexOf('hidden')!=-1) $('#UnhealthyThreshold').parent().parent().removeClass('hidden');
    if($('#HealthyThreshold').parent().parent().attr('class').indexOf('hidden')!=-1) $('#HealthyThreshold').parent().parent().removeClass('hidden');
    var protocol=$('#Protocol').val();
    if(protocol=='http'||protocol=='https'){
      if($('#HealthCheckDomain').parent().parent().attr('class').indexOf('hidden')!=-1) $('#HealthCheckDomain').parent().parent().removeClass('hidden');
      if($('#HealthCheckURI').parent().parent().attr('class').indexOf('hidden')!=-1) $('#HealthCheckURI').parent().parent().removeClass('hidden');
      if($('#HealthCheckHttpCode').parent().parent().attr('class').indexOf('hidden')!=-1) $('#HealthCheckHttpCode').parent().parent().removeClass('hidden');
    }else{
      if($('#HealthCheckDomain').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckDomain').parent().parent().addClass('hidden');
      if($('#HealthCheckURI').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckURI').parent().parent().addClass('hidden');
      if($('#HealthCheckHttpCode').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckHttpCode').parent().parent().addClass('hidden');
    }
  }else{
    if($('#HealthCheckConnectPort').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckConnectPort').parent().parent().addClass('hidden');
    if($('#HealthCheckTimeout').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckTimeout').parent().parent().addClass('hidden');
    if($('#HealthCheckInterval').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckInterval').parent().parent().addClass('hidden');
    if($('#UnhealthyThreshold').parent().parent().attr('class').indexOf('hidden')==-1) $('#UnhealthyThreshold').parent().parent().addClass('hidden');
    if($('#HealthyThreshold').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthyThreshold').parent().parent().addClass('hidden');
    if($('#HealthCheckDomain').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckDomain').parent().parent().addClass('hidden');
    if($('#HealthCheckURI').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckURI').parent().parent().addClass('hidden');
    if($('#HealthCheckHttpCode').parent().parent().attr('class').indexOf('hidden')==-1) $('#HealthCheckHttpCode').parent().parent().addClass('hidden');
  }
  check();
}
var checkStickySession=function(){
  if($('#StickySession').attr('disabled')=='disabled'){
    if($('#StickySessionType').parent().parent().attr('class').indexOf('hidden')==-1) $('#StickySessionType').parent().parent().addClass('hidden');
  }else{
    if($('#StickySession').val()=='on'){
      if($('#StickySessionType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#StickySessionType').parent().parent().removeClass('hidden');
    }else{
      if($('#StickySessionType').parent().parent().attr('class').indexOf('hidden')==-1) $('#StickySessionType').parent().parent().addClass('hidden');
    }
  }
  checkStickySessionType();
}
var checkStickySessionType=function(){
  if($('#StickySessionType').parent().parent().attr('class').indexOf('hidden')!=-1){
    if($('#CookieTimeout').parent().parent().attr('class').indexOf('hidden')==-1) $('#CookieTimeout').parent().parent().addClass('hidden');
    if($('#Cookie').parent().parent().attr('class').indexOf('hidden')==-1) $('#Cookie').parent().parent().addClass('hidden');
    check();
    return;
  }
  if($('#StickySessionType').val()=='insert'){
    if($('#CookieTimeout').parent().parent().attr('class').indexOf('hidden')!=-1) $('#CookieTimeout').parent().parent().removeClass('hidden');
    if($('#Cookie').parent().parent().attr('class').indexOf('hidden')==-1) $('#Cookie').parent().parent().addClass('hidden');
  }else{
    if($('#CookieTimeout').parent().parent().attr('class').indexOf('hidden')==-1) $('#CookieTimeout').parent().parent().addClass('hidden');
    if($('#Cookie').parent().parent().attr('class').indexOf('hidden')!=-1) $('#Cookie').parent().parent().removeClass('hidden');
  }
  check();
}

var twiceCheck=function(action,idx,desc,ipidx){
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
      case 'slb':
        switch(action){
          case 'del':
            modalTitle='删除SLB';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认删除? <span class="text text-danger">警告! 操作不可回退!</span></h4>' +
              'Id : '+idx+'<br/> SLB名称 : '+desc;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="LoadBalancerId" name="LoadBalancerId" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="delete">';
            break;
          default:
            modalTitle='非法请求';
            notice='<div class="note note-danger">错误信息：参数错误</div>';
            pageNotify('error','非法请求！','错误信息：参数错误');
            break;
        }
        break;
      case 'listener':
        switch(action){
          case 'del':
            modalTitle='删除单元';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认删除? <span class="text text-danger">警告! 操作不可回退!</span></h4> Slb Id : '+idx;
            if(cache.slb_info.LoadBalancerId==idx) modalBody+='<br/> Slb名称 : '+cache.slb_info.LoadBalancerName;
            modalBody+='<br/> 协议 : '+desc+'<br/> 端口 : '+ipidx;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="LoadBalancerId" name="LoadBalancerId" value="'+idx+'">';
            modalBody+='<input type="hidden" id="Protocol" name="Protocol" value="'+desc+'">';
            modalBody+='<input type="hidden" id="ListenerPort" name="ListenerPort" value="'+ipidx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="delete">';
            break;
          default:
            modalTitle='非法请求';
            notice='<div class="note note-danger">错误信息：参数错误</div>';
            pageNotify('error','非法请求！','错误信息：参数错误');
            break;
        }
        break;
      case 'backend':
        var count = 0, postUpdate = [], postDel = [], list='', ServerId ='', ServerIp='', Weight='';
        if(idx){
          count++;
          $('input:checkbox[value="'+idx+'"]').attr('checked','true');
          Weight=$('#weight_'+idx).val();
          ServerIp=$('#backend_ip_'+idx).val();
          if(!ServerIp) ServerIp=idx;
          list+='<span class="col-sm-3">'+ServerIp+' : '+Weight+'</span>';
          postUpdate.push({ServerId: idx, Weight: parseInt(Weight)});
          postDel.push(idx);
        }else{
          $('input:checkbox[id=list]:checked').each(function(i){
            count++;
            ServerId=$(this).val();
            Weight=$('#weight_'+ServerId).val();
            ServerIp=$('#backend_ip_'+ServerId).val();
            if(!ServerIp) ServerIp=ServerId;
            list+='<span class="col-sm-3">'+ServerIp+' : '+Weight+'</span>';
            postUpdate.push({ServerId: ServerId, Weight: parseInt(Weight)});
            postDel.push(ServerId);
          });
        }
        if(action=='update'||action=='del'){
          modalBody+='<div class="form-group col-sm-12">';
          modalBody+='<h5><strong>当前操作, 将影响如下列表:</strong></h5>';
          modalBody+='<p><strong class="text-primary">涉及总数</strong>: 共 <span class="badge badge-danger">'+count+'</span> 个 </p>';
          modalBody+='<p><strong class="text-primary">涉及列表</strong>:';
          modalBody+='<div class="col-sm-12">'+list+'</div>';
          modalBody+='</div>';
          modalBody+='<input type="hidden" id="LoadBalancerId" name="LoadBalancerId" value="'+ cache.slb_info.LoadBalancerId +'">';
          switch(action){
            case 'update':
              modalTitle='调整Backend权重';
              modalBody+='<textarea class="hidden" id="backend" name="backend">'+JSON.stringify(postUpdate)+'</textarea>'
              modalBody+='<input type="hidden" id="page_action" name="page_action" value="update">';
              break;
            case 'del':
              modalTitle='移除Backend';
              modalBody+='<textarea class="hidden" id="backend" name="backend">'+JSON.stringify(postDel)+'</textarea>'
              modalBody+='<input type="hidden" id="page_action" name="page_action" value="delete">';
              break;
          }
          if(count==0) btnDisable=true;
        }else{
          modalTitle='非法请求';
          notice='<div class="note note-danger">错误信息：参数错误</div>';
          pageNotify('error','非法请求！','错误信息：参数错误');
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

var getList=function(type,idx,tab){
  if(!type) type='region';
  if(!tab) tab='slb';
  if(type=='listener'||type=='backend'){
    updateSelect('fSlb',idx,type);
    return false;
  }
  var url='';
  switch(type){
    case 'region':
      url='/api/for_cloud/region.php?action=list';
      break;
    case 'slb':
      url='/api/for_cloud/slb.php?action=list';
      break;
  }
  var postData={'pagesize':1000};
  $('#tab').val(tab);
  var actionDesc='';
  switch (type){
    case 'region':
      actionDesc='地域';
      if(cache.region_id) idx=cache.region_id;
      break;
    case 'slb':
      postData.fRegion=$('#fRegion').val();
      actionDesc='AliSLB';
      cache.region_id=postData.fRegion;
      break;
  }
  NProgress.start();
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      NProgress.done();
      if(data.code==0){
        switch (type){
          case 'region':
            cache.region = data.content;
            updateSelect('fRegion',idx);
            if(data.content.length==0){
              switch (type){
                case 'region':
                  $('#fRegion').parent().parent().attr('hidden',false);
                  $('#fSlb').parent().parent().attr('hidden',true);
                  pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先在云服务商控制台[创建'+actionDesc+']！',false);
                  $('#table-head').html('<tr><td>'+actionDesc+'数据为空！请先在云服务商控制台[创建'+actionDesc+']！</td></tr>');
                  break;
              }
            }
            break;
          case 'slb':
            cache.slb = data.content;
            break;
        }
      }else{
        pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
      }
      if(type=='slb') list(1);
    },
    error: function (){
      NProgress.done();
      pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
      if(type=='slb') list(1);
    }
  });
}

var getInfo=function(idx){
  cache.backend_list = [];
  var tab=$('#tab').val();
  var url='',postData={};
  url='/api/for_cloud/slb.php';
  if(!idx) idx=$('#fSlb').val();
  if(!idx) return false;
  postData={"action":"info","fIdx":idx};
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      //执行结果提示
      if(data.code==0){
        if(typeof(data.content)!='undefined'){
          cache.slb_info=data.content;
          switch(tab){
            case 'listener':
              if(typeof cache.slb_info.ListenerPortsAndProtocol != 'undefined')
                cache.listener=cache.slb_info.ListenerPortsAndProtocol.ListenerPortAndProtocol;
              break;
            case 'backend':
              if(typeof cache.slb_info.BackendServers != 'undefined')
                cache.backend=cache.slb_info.BackendServers.BackendServer;
              break;
          }
        }
      }else{
        pageNotify('error','加载失败！','错误信息：'+data.msg);
      }
      list(1);
    },
    error: function (){
      pageNotify('error','加载失败！','错误信息：接口不可用');
      list(1);
    }
  });
}

var getSlbPort=function(idx,j){
  var tab=$('#tab').val();
  var url='',postData={};
  url='/api/for_cloud/slb.php?'+j+'_'+idx;
  if(!idx) return false;
  postData={"action":"info","fIdx":idx};
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      //执行结果提示
      var wan='';
      if(data.code==0){
        if(typeof(data.content)!='undefined'){
          if(typeof data.content.ListenerPortsAndProtocol != 'undefined'){
            $.each(data.content.ListenerPortsAndProtocol.ListenerPortAndProtocol,function(k,v){
              if(wan) wan+='<br/>';
              wan += v.ListenerProtocol + ': '+ v.ListenerPort;
            });
            if(!wan) wan='未配置';
          }
        }
      }else{
        pageNotify('error','加载失败！','错误信息：'+data.msg);
        wan='获取失败';
      }
      $('#wan_port_'+j).html(wan);
    },
    error: function (){
      $('#wan_port_'+j).html('获取失败');
    }
  });
}

//获取属性
var getDesc=function(type,id,idx,protocol,port){
  var url='',postData={};
  url='/api/for_cloud/slb_'+type+'.php';
  switch(type){
    case 'listener':
      postData={"action":"info","fIdx":idx,"fProtocol":protocol,"fPort":port};
      break;
    case 'backend':
      postData={"action":"status","fIdx":idx};
      break;
    default:
      return false;
      break;
  }
  if(!idx) return false;
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      //执行结果提示
      if(data.code==0){
        if(typeof(data.content)!='undefined'){
          switch(type){
            case 'listener':
              cache.listener_info[id]=data.content;
              var v=data.content;
              $('#listener_backend_port_'+id).html(v.BackendServerPort);
              switch(v.Status){
                case 'running':
                  $('#listener_status_'+id).html('<div><label class="tooltips" title="点击停止" onclick="return false;"><input type="checkbox" class="js-switch" checked onchange="switchs(\'stop\',\'' + idx + '\',\'' + port + '\')"/> 运行中</label></div>');
                  switchery();
                  break;
                case 'stopped':
                  $('#listener_status_'+id).html('<div><label class="tooltips" title="点击启动" onclick="return false;"><input type="checkbox" class="js-switch" onchange="switchs(\'start\',\'' + idx + '\',\'' + port + '\')"/> 已停止</label></div>');
                  switchery();
                  break;
                case 'starting':
                  $('#listener_status_'+id).html('<span class="badge bg-blue-sky">启动中</span>');
                  break;
                case 'configuring':
                  $('#listener_status_'+id).html('<span class="badge bg-blue-sky">配置中</span>');
                  break;
                case 'stopping':
                  $('#listener_status_'+id).html('<span class="badge bg-blue-sky">停止中</span>');
                  break;
                default:
                  $('#listener_status_'+id).html('<span class="badge">'+ v.Status +'</span>');
                  break;
              }
              break;
            case 'backend':
              $.each(data.content, function (k,v) {
                cache.backend_info[v.ServerId]=v;
                $('#backend_ip_'+ v.ServerId).html(v.Address);
                switch(v.ServerHealthStatus){
                  case 'normal':
                    $('#backend_status_'+ v.ServerId).html('<span class="badge bg-green">'+v.ServerHealthStatus+'</span>');
                    break;
                  case 'abnormal':
                    $('#backend_status_'+ v.ServerId).html('<span class="badge bg-red">'+v.ServerHealthStatus+'</span>');
                    break;
                  default:
                    $('#backend_status_'+ v.ServerId).html('<span class="badge">'+v.ServerHealthStatus+'</span>');
                    break;
                }
              });
              break;
          }
        }
      }else{
        pageNotify('error','加载失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载失败！','错误信息：接口不可用');
      NProgress.done();
    }
  });
}

//展示Listener详情
var viewDesc=function(id){
  NProgress.start();
  var title='查看Listener详情',text='',height='';
  var tStyle='word-break:break-all;word-warp:break-word;';
  if(id){
    var locale={};
    if(typeof(locale_messages.cloud)) locale = locale_messages.cloud;
    if(typeof cache.listener_info[id] != 'undefined'){
      $.each(cache.listener_info[id],function(k,v){
        if(locale[k]!=false){
          if(v=='') v='空';
          if(typeof(locale[k])!='undefined') k=locale[k];
          text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
        }
      });
    }else{
      pageNotify('warning','数据为空！');
      text='<div class="note note-warning">数据为空！</div>';
    }
  }else{
    pageNotify('warning','错误操作！','错误信息：参数错误');
    title='非法请求';
    text='<div class="note note-danger">错误信息：参数错误</div>';
  }
  $('#myViewModalLabel').html(title);
  $('#myViewModalBody').html(text);
  NProgress.done();
}

var getNetworkType=function(){
  switch ($('#NetworkType').val()){
    case 'common':
      if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')!=-1) $('#InternetChargeType').parent().parent().removeClass('hidden');
      if($('#ZoneId').parent().parent().attr('class').indexOf('hidden')==-1) $('#ZoneId').parent().parent().addClass('hidden');
      if($('#VpcId').parent().parent().attr('class').indexOf('hidden')==-1) $('#VpcId').parent().parent().addClass('hidden');
      if($('#VSwitchId').parent().parent().attr('class').indexOf('hidden')==-1) $('#VSwitchId').parent().parent().addClass('hidden');
      if($('#SecurityGroup').parent().parent().attr('class').indexOf('hidden')==-1) $('#SecurityGroup').parent().parent().addClass('hidden');
      $('#AddressType').val('internet').trigger('change');
      $('#AddressType').find("option[value=intranet]").attr('disabled',true).trigger('change');
      break;
    case 'custom':
      if($('#InternetChargeType').parent().parent().attr('class').indexOf('hidden')==-1) $('#InternetChargeType').parent().parent().addClass('hidden');
      if($('#ZoneId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#ZoneId').parent().parent().removeClass('hidden');
      if($('#VpcId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#VpcId').parent().parent().removeClass('hidden');
      if($('#VSwitchId').parent().parent().attr('class').indexOf('hidden')!=-1) $('#VSwitchId').parent().parent().removeClass('hidden');
      if($('#SecurityGroup').parent().parent().attr('class').indexOf('hidden')!=-1) $('#SecurityGroup').parent().parent().removeClass('hidden');
      $('#AddressType').find("option[value=intranet]").attr('disabled',false).trigger('change');
      getVpcId();
      break;
  }
  check();
}

var getZoneAndVpc=function(){
  check();
  if($('#NetworkType').val()=='custom') {
    getZoneId();
    getVpcId();
  }
}

var getZoneId=function(){
  var actionDesc="可用区",tSelect='ZoneId';
  var url='/api/for_cloud/zone.php?action=list';
  var idx=$('#RegionId').val();
  if(!idx){
    $('#'+tSelect).empty();
    $('#'+tSelect).append('<option value="">请选择</option>');
    $('#'+tSelect).val($('#'+tSelect+' option:nth-child(1)').val()).trigger('change');
    return false;
  }
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

var getVpcId=function(){
  var actionDesc="可用区",tSelect='VpcId';
  var url='/api/for_cloud/vpc.php?action=list';
  var idx=$('#RegionId').val();
  if(!idx){
    $('#'+tSelect).empty();
    $('#'+tSelect).append('<option value="">请选择</option>');
    $('#'+tSelect).val($('#'+tSelect+' option:nth-child(1)').val()).trigger('change');
    return false;
  }
  var postData={"pagesize":1000,"fIdx":idx};
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
  if($('#NetworkType').val()=='custom'){
    getVSwitchId();
    getSecurityGroup();
  }
}

var getVSwitchId=function(){
  var actionDesc="子网",tSelect='VSwitchId';
  var url='/api/for_cloud/subnet.php?action=list';
  var zone=$('#ZoneId').val();
  var idx=$('#VpcId').val();
  if(!zone||!idx){
    $('#'+tSelect).empty();
    $('#'+tSelect).append('<option value="">请选择</option>');
    $('#'+tSelect).val($('#'+tSelect+' option:nth-child(1)').val()).trigger('change');
    return false;
  }
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
  var actionDesc="安全组",tSelect='SecurityGroup';
  var url='/api/for_cloud/security_group.php?action=list';
  var region=$('#RegionId').val();
  var idx=$('#VpcId').val();
  if(!region||!idx){
    $('#'+tSelect).empty();
    $('#'+tSelect).append('<option value="">请选择</option>');
    $('#'+tSelect).val($('#'+tSelect+' option:nth-child(1)').val()).trigger('change');
    return false;
  }
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

var getCertificate=function(){
  var actionDesc="安全证书",tSelect='ServerCertificateId';
  var url='/api/for_cloud/https.php?action=list';
  var idx=$('#fRegion').val();
  if(!idx) return false;
  var postData={"pagesize":1000,"fIdx":idx};
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      cache.https = (typeof data.content != 'undefined') ? data.content : [];
      updateSelect(tSelect);
      if(data.code!=0){
        pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
      }else{
        if(data.content.length==0) pageNotify('warning','获取'+actionDesc+'成功！','数据为空!');
      }
    },
    error: function (){
      cache.https = [];
      updateSelect(tSelect);
      pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}

var updateSelect=function(name,idx,tab){
  if(tab) $('#tab').val(tab);
  var tSelect=$('#'+name),data='';
  switch(name){
    case 'fRegion':
    case 'RegionId':
      data=cache.region;
      break;
    case 'fSlb':
      data=cache.slb;
      break;
    case 'ZoneId':
      data=cache.zone;
      break;
    case 'VpcId':
      data=cache.vpc;
      break;
    case 'VSwitchId':
      data=cache.subnet;
      break;
    case 'SecurityGroup':
      data=cache.security_group;
      break;
    case 'ServerCertificateId':
      data=cache.https;
      break;
  }
  tSelect.empty();
  if(data.length>0){
    switch(name){
      case 'fRegion':
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.RegionName + '">' + v.RegionName + '</option>');
        });
        break;
      case 'RegionId':
        tSelect.append('<option value="">请选择</option>');
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.RegionName + '">' + v.RegionName + '</option>');
        });
        break;
      case 'ZoneId':
        tSelect.removeAttr("disabled");
        tSelect.append('<option value="">请选择</option>');
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.ZoneName + '">' + ((v.ZoneName)?v.ZoneName:v.ZoneId) + '</option>');
        });
        break;
      case 'VpcId':
        tSelect.removeAttr("disabled");
        tSelect.append('<option value="">请选择</option>');
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.VpcId + '">' + v.VpcId + '</option>');
        });
        break;
      case 'VSwitchId':
        tSelect.removeAttr("disabled");
        tSelect.append('<option value="">请选择</option>');
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.SubnetId + '">' + v.CidrBlock + '</option>');
        });
        break;
      case 'SecurityGroup':
        tSelect.removeAttr("disabled");
        tSelect.append('<option value="">请选择</option>');
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.GroupId + '">' + ((v.GroupName)?v.GroupName:v.GroupId) + '</option>');
        });
        break;
      case 'ServerCertificateId':
        tSelect.removeAttr("disabled");
        tSelect.append('<option value="">请选择</option>');
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.ServerCertificateId + '">' + ((v.ServerCertificateName)?v.ServerCertificateName:v.ServerCertificateId) + '</option>');
        });
        break;
      default:
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.LoadBalancerId + '">' + ((v.LoadBalancerName)?v.LoadBalancerName:v.LoadBalancerId) + '</option>');
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

//开关
var switchs=function(action,idx,port){
  NProgress.start();
  var actionDesc="";
  var url='/api/for_cloud/slb.php';
  var postData={LoadBalancerId:idx};
  var tab=$('#tab').val();
  switch(tab){
    case 'listener':
      url='/api/for_cloud/slb_listener.php';
      postData.ListenerPort=parseInt(port);
      actionDesc=(action=='start')?'启动监听':'停止监听';
      break;
    default:
      actionDesc=(action=='active')?'启用SLB':'停用SLB';
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
        pageNotify('success',actionDesc+'操作成功！');
      }else{
        pageNotify('error',actionDesc+'操作失败！','错误信息：'+data.msg);
      }
      //重载列表
      switch(tab){
        case 'slb':
          getList('slb');
          break;
        case 'listener':
          getInfo();
          break;
      }
      NProgress.done();
    },
    error: function (){
      pageNotify('error',actionDesc+'操作失败！','错误信息：接口不可用');
      NProgress.done();
    }
  });
}

var get = function (protocol,port) {
  var tab=$('#tab').val();
  var url='/api/for_cloud/slb_'+tab+'.php',postData={};
  var idx=$('#LoadBalancerId').val();
  switch (tab){
    default:
      postData={"action":"info","fIdx":idx,"fProtocol":protocol,"fPort":port};
      break;
  }
  if(idx&&protocol&&port){
    $('#page_action').val('update');
    $.ajax({
      type: "POST",
      url: url,
      data: postData,
      dataType: "json",
      success: function (data) {
        //执行结果提示
        if(data.code==0){
          if(typeof(data.content)!='undefined'){
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
                    //$('#'+k).find("option[value='"+v+"']").attr("selected",true);
                    $('#'+k).val(v).trigger('change');
                    break;
                  default:
                    $('#'+k).val(v);
                    break;
                }
              }
            });
            switch($('#Protocol').val()){
              case 'tcp':
                if($('#PersistenceTimeout').parent().parent().attr('class').indexOf('hidden')!=-1) $('#PersistenceTimeout').parent().parent().removeClass('hidden');
                break;
              case 'udp':
                if($('#PersistenceTimeout').parent().parent().attr('class').indexOf('hidden')!=-1) $('#PersistenceTimeout').parent().parent().removeClass('hidden');
                break;
              default:
                if($('#PersistenceTimeout').parent().parent().attr('class').indexOf('hidden')==-1) $('#PersistenceTimeout').parent().parent().addClass('hidden');
                break;
            }
            $('#ListenerPort').attr('disabled',true);
            $('#BackendServerPort').attr('disabled',true);
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
  }
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

//Backend part
var getBackend=function(page){
  if(!page) page=1;
  var url='/api/for_cloud/ecs.php?action=list&page='+page;
  NProgress.start();
  $.ajax({
    type: "POST",
    url: url,
    data: '',
    dataType: "json",
    success: function (listdata) {
      if(listdata.code==0){
        var modalTitle=$('#myTableModalLabel');
        modalTitle.html('添加 Backend');
        var pageinfo = $("#modal_table-pageinfo");//分页信息
        var paginate = $("#modal_table-paginate");//分页代码
        var head = $("#modal_table-head");//数据表头
        var body = $("#modal_table-body");//数据列表
        //清除当前页面数据
        pageinfo.html("");
        paginate.html("");
        head.html('<tr><td>Loading ...</td></tr>');
        body.html('');
        //生成分页
        processModalPage(listdata, pageinfo, paginate,'getBackend');
        //生成列表
        processModalBody(listdata, head, body);
        body.append('<input type="hidden" id="page_action" name="page_action" value="insert">');
        NProgress.done();
      }else{
        pageNotify('error','加载ECS失败！','错误信息：'+listdata.msg);
        NProgress.done();
      }
    },
    error: function (){
      pageNotify('error','加载ECS失败！','错误信息：接口不可用');
      NProgress.done();
    }
  });
}

var processModalPage = function(data,pageinfo,paginate,func){
  if(!func) func='list';
  var begin = data.pageSize * ( data.page - 1 ) + 1;
  var end = ( data.count > begin + data.pageSize - 1 ) ? begin + data.pageSize - 1 : data.count;
  pageinfo.html('Showing '+begin+' to '+end+' of '+data.count+' records');
  var p1=(data.page-1>0)?data.page-1:1;
  var p2=data.page+1;
  prev='<li><a href="javascript:;" onclick="'+func+'('+p1+')"><i class="fa fa-angle-left"></i></a></li>';
  paginate.append(prev);
  for (var i = 1; i <= data.pageCount; i++) {
    var li='';
    if(i==data.page){
      li='<li class="active"><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
    }else{
      if(i==1||i==data.pageCount){
        li='<li><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
      }else{
        if(i==p1){
          if(p1>2){
            li='<li class="disabled"><a href="javascript:;" href="#">...</a></li>'+"\n"+'<li><a onclick="'+func+'('+i+')">'+i+'</a></li>';
          }else{
            li='<li><a href="javascript:;" onclick="'+func+'('+i+')">'+i+'</a></li>';
          }
        }else{
          if(i==p2){
            if(p2<data.pageCount-1){
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
  if(p2>data.pageCount) p2=data.pageCount;
  next='<li class="next"><a href="javascript:;" title="Next" onclick="'+func+'('+p2+')"><i class="fa fa-angle-right"></i></a></li>';
  paginate.append(next);
}

var processModalBody = function(data,head,body){
  var td="";
  var title=['<input type="checkbox" name="modalSelectAll" id="modalSelectAll" onclick="modalCheckAll(this)">','ECS_ID','ECS地址','机型模板','云厂商','权重'];
  if(title){
    var tr = $('<tr></tr>');
    for (var i = 0; i < title.length; i++) {
      td = '<th>' + title[i] + '</th>';
      tr.append(td);
    }
    head.html(tr);
  }
  if(data.content){
    if(data.content.length>0) {
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        if ($.inArray(v.InstanceId, cache.backend_list) != -1) {
          td = '<td><input type="checkbox" id="modal_list" name="modal_list[]" value="'+ v.InstanceId+'" onclick="checkBackend()" checked disabled></td>';
        }else{
          td = '<td><input type="checkbox" id="modal_list" name="modal_list[]" value="'+ v.InstanceId+'" onclick="checkBackend()"></td>';
        }
        tr.append(td);
        td = '<td>' + v.InstanceId + '</td>';
        tr.append(td);
        td = '<td>' + v.PrivateIpAddress + '</td>';
        tr.append(td);
        td = '<td>' + v.Cluster.Name + '</td>';
        tr.append(td);
        td = '<td>' + v.Cluster.Provider + '</td>';
        tr.append(td);
        td = '<td><input id="' + v.InstanceId + '_weight" type="number" value="100" max=100 min=0 onkeyup="isNumberValid(this,1,100,100)"></td>';
        tr.append(td);
        body.append(tr);
      }
    }else{
      pageNotify('info','Warning','ECS数据为空！');
    }
  }
}

var modalCheckAll=function(o){
  $('[id=modal_list]:checkbox').prop('checked', o.checked);
  checkBackend();
}

var isNumberValid=function(o,a,b,d){
  var id = o.id.replace('weight_','');
  $('input:checkbox[value="'+id+'"]').attr('checked','true');
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
}

var checkBackend=function(){
  var o=document.getElementsByName("modal_list[]");
  var disabled=true;
  for(var i=0;i<o.length;i++){
    if(o[i].checked){
      disabled=false;
      break;
    }
  }
  $('#btnCommitModal').attr('disabled',disabled);
}

var addBackend=function(){
  var tab=$('#tab').val();
  var url='/api/for_cloud/slb_'+tab+'.php?action=insert';
  var postData=[];
  var form=$('#myModalBody').find("input,select,textarea");
  $('input:checkbox[id=modal_list]:checked').each(function (i) {
    var t = {};
    t['ServerId'] = $(this).val();
    t['weight'] = parseInt($('#' + $(this).val() + '_weight').val());
    postData.push(t);
  });
  var action=$("#page_action").val();
  var actionDesc='添加Backend';
  $.ajax({
    type: "POST",
    url: url,
    data: {"action":action,"fIdx": $('#fSlb').val(),"data":JSON.stringify(postData)},
    dataType: "json",
    success: function (data) {
      //执行结果提示
      if(data.code==0){
        pageNotify('success','【'+actionDesc+'】操作成功！');
      }else{
        pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+data.msg);
      }
      //重载列表
      getInfo();
      //处理模态框和表单
      $("#myTableModal :input").each(function () {
        $(this).val("");
      });
      $("#myTableModal").on("hidden.bs.modal", function() {
        $(this).removeData("bs.modal");
      });
      $("#myTableModal").modal('hide');
    },
    error: function (){
      pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
    }
  });
}