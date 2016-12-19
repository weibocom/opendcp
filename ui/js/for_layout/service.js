cache = {
  page: 1,
  cluster_id: '',
  service_id: '',
  pool_id: '',
  cluster: [],
  service: [],
  pool: [],
  pool_vm: {},
  pool_quota: {},
  copy: {
    ip: [],
  },
  task_tpl: [],
  balance: [],
  vm_type: [],
  ip: [], //选中IP列表
}

var reset = function(){
  $('#fIdx').val('');
}

var list = function(page,tab,idx) {
  $('.popovers').each(function(){$(this).popover('hide');});
  NProgress.start();
  var postData={};
  if(!tab){
    tab=$('#tab').val();
  }
  if(tab!='service'&&tab!='pool'&&tab!='node'){
    tab='service';
  }
  var fClusterId=$('#fClusterId').val();
  var fService=$('#fService').val();
  var fPool=$('#fPool').val();
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'service':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_3').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=add"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fClusterId":fClusterId,"fIdx":fIdx};
      cache.cluster_id=fClusterId;
      $('#fClusterId').parent().parent().attr('hidden',false);
      $('#fService').parent().parent().attr('hidden',true);
      $('#fPool').parent().parent().attr('hidden',true);
      break;
    case 'pool':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','active');
      $('#tab_3').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=add&par_id='+ $('#fService').val() +'&par_name='+ $('#fService').find("option:selected").text() +'"> 创建 <i class="fa fa-plus"></i></a>');
      if(idx){
        fService=idx;
        $('#fService').val(fService);
      }
      postData={"fService":fService,"fIdx":fIdx};
      cache.service_id=fService;
      $('#fClusterId').parent().parent().attr('hidden',true);
      $('#fService').parent().parent().attr('hidden',false);
      $('#fPool').parent().parent().attr('hidden',true);
      break;
    case 'node':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','');
      $('#tab_3').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-danger" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\')"> 批量删除 <i class="fa fa-trash-o"></i></a>');
      if(idx){
        fPool=idx;
        $('#fPool').val(fPool);
      }
      postData={"fPool":fPool,"fIdx":fIdx};
      cache.pool_id=fPool;
      $('#fClusterId').parent().parent().attr('hidden',true);
      $('#fService').parent().parent().attr('hidden',true);
      $('#fPool').parent().parent().attr('hidden',false);
      break;
  }
  var url='/api/for_layout/'+tab+'.php';
  if (!page) {
    page = cache.page;
  }else{
    cache.page = page;
  }
  url+='?action=list&page=' + page;
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
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
        NProgress.done();
      }else{
        pageNotify('error','加载失败！','错误信息：'+listdata.msg);
        NProgress.done();
      }
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
  prev='<li><a onclick="list('+p1+')"><i class="fa fa-angle-left"></i></a></li>';
  paginate.append(prev);
  for (var i = 1; i <= data.pageCount; i++) {
    var li='';
    if(i==data.page){
      li='<li class="active"><a onclick="list('+i+')">'+i+'</a></li>';
    }else{
      if(i==1||i==data.pageCount){
        li='<li><a onclick="list('+i+')">'+i+'</a></li>';
      }else{
        if(i==p1){
          if(p1>2){
            li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a onclick="list('+i+')">'+i+'</a></li>';
          }else{
            li='<li><a onclick="list('+i+')">'+i+'</a></li>';
          }
        }else{
          if(i==p2){
            if(p2<data.pageCount-1){
              li='<li><a onclick="list('+i+')">'+i+'</a></li>'+"\n"+'<li class="disabled"><a href="#">...</a></li>';
            }else{
              li='<li><a onclick="list('+i+')">'+i+'</a></li>';
            }
          }
        }
      }
    }
    paginate.append(li);
  }
  if(p2>data.pageCount) p2=data.pageCount;
  next='<li class="next"><a title="Next" onclick="list('+p2+')"><i class="fa fa-angle-right"></i></a></li>';
  paginate.append(next);
}

//生成列表
var processBody = function(data,head,body){
  var td="";
  if(data.title){
    var tr = $('<tr></tr>');
    for (var i = 0; i < data.title.length; i++) {
      var v = data.title[i];
      var t='';
      if(data.title[i]=='IP') t='<a class="pull-right tooltips" data-container="body" data-trigger="hover" data-original-title="复制整列" data-toggle="modal" data-target="#myViewModal" onclick="copy(\'ip\')"><i class="fa fa-copy"></i></a>';
      td = '<th>' + v + t + '</th>';
      tr.append(td);
    }
    head.html(tr);
  }
  if(data.content){
    if(data.content.length>0){
      var tab=$('#tab').val();
      cache.copy.ip=[];
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        if(tab!='node'){
          td = '<td>' + v.i + '</td>';
          tr.append(td);
        }
        var btnAdd='',btnEdit='',btnDel='';
        switch(tab){
          case 'service':
            if(i==0) cache.service_id = v.id;
            cache.service.push(v);
            td = '<td title="服务ID: '+ v.id +'"><a class="tooltips" title="查看服务详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'service\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            var tA='<a class="tooltips" title="查看服务池列表" onclick="getList(\'service\',\''+v.id+'\')"><i class="fa fa-bars"></i></a> ' +
                '<a class="pull-right tooltips" title="添加服务池" data-toggle="modal" data-target="#myModal" href="edit_pool.php?action=add&par_id='+ v.id +'&par_name='+ v.name +'"><i class="fa fa-plus"></i></a>';
            td = '<td>'+tA+'</td>';
            tr.append(td);
            td = '<td title="集群ID: '+ v.cluster_id +'">' + getName('cluster',v.cluster_id) + '</td>';
            tr.append(td);
            td = '<td>' + v.desc + '</td>';
            tr.append(td);
            td = '<td>' + v.service_type + '</td>';
            tr.append(td);
            //td = '<td style="width:200px;word-wrap:break-word;word-break:break-all;">' + v.docker_image + '</td>';
            td = '<td>' + v.docker_image + '</td>';
            tr.append(td);
            btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.id + '"><i class="fa fa-edit"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'pool':
            if(i==0) cache.pool_id = v.id;
            td = '<td title="服务池ID: '+ v.id +'"><a class="tooltips" title="查看服务池详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'pool\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            var tA='<a class="tooltips" title="查看服务详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'service\',\''+v.service_id+'\')">' + getName('service',v.service_id) + '</a>';
            td = '<td title="服务ID: '+ v.service_id +'">'+tA+'</td>';
            tr.append(td);
            td = '<td>' + v.desc + '</td>';
            tr.append(td);
            var tA='<a class="tooltips" title="查看节点列表" onclick="getList(\'pool\',\''+ v.id+'\')" id="ip_'+ v.id +'">'+ v.node_count +'</a>' +
              '<a class="pull-right tooltips" title="添加节点" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'addnode\',\''+v.id+'\',\''+v.name+'\')"><i class="fa fa-plus"></i></a>';
            td = '<td>'+tA+'</td>';
            tr.append(td);
            td = '<td>' + getName('vm_type',v.vm_type) + '</td>';
            tr.append(td);
            cache.pool_vm[v.id] = v.vm_type;
            td = '<td>' + getName('balance',v.sd_id) + '</td>';
            tr.append(td);
            var tExpand=(typeof(v.tasks.expand)!='undefined')?v.tasks.expand:'0';
            var tShrink=(typeof(v.tasks.shrink)!='undefined')?v.tasks.shrink:'0';
            var tDeploy=(typeof(v.tasks.deploy)!='undefined')?v.tasks.deploy:'0';
            btnAdd = '<a class="text-success tooltips" title="扩容" data-toggle="modal" data-target="#myModal" href="expand_'+tab+'.php?id='+ v.id +'&name='+ v.name +'&idx=' + tExpand + '"><i class="fa fa-plus"></i></a>' +
              ' <a class="text-danger tooltips" title="缩容" data-toggle="modal" data-target="#myModal" href="shrink_'+tab+'.php?id='+ v.id +'&name='+ v.name +'&idx=' + tShrink + '"><i class="fa fa-minus"></i></a>' +
              ' <a class="text-primary tooltips" title="上线" data-toggle="modal" data-target="#myModal" href="deploy_'+tab+'.php?id='+ v.id +'&name='+ v.name +'&idx=' + tDeploy + '"><i class="fa fa-refresh"></i></a>';
            td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnAdd + '</div></td>';
            tr.append(td);
            btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&par_id='+ $('#fService').val() +'&par_name='+ $('#fService').find("option:selected").text() +'&idx=' + v.id + '"><i class="fa fa-edit"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'node':
            cache.copy.ip.push(v.ip);
            td = '<td><input type="checkbox" id="list" name="list[]" value="'+ v.pool_id + ',' + getName('pool',v.pool_id) + ',' + v.ip+',' + v.id+'" /></td>';
            tr.append(td);
            td = '<td>' + v.i + '</td>';
            tr.append(td);
            td = '<td>' + getName('pool',v.pool_id) + '</td>';
            tr.append(td);
            td = '<td>' + v.ip + '</td>';
            tr.append(td);
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\'' + v.pool_id + ',' + getName('pool',v.pool_id) + '\',\'' + v.ip + '\',\'' + v.id + '\')"><i class="fa fa-trash-o"></i></a>';
            break;
        }
        td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnEdit + ' ' + btnDel + '</div></td>';
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

//增删改查
var change=function(){
  var tab=$('#tab').val();
  var page_other=$('#page_other').val();
  switch (page_other){
    case 'addpool': tab='pool'; break;
    case 'addnode': tab='node'; break;
  }
  var url='/api/for_layout/'+tab+'.php';
  var postData={};
  var form=$('#myModalBody').find("input,select,textarea");

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
  //console.log("action="+action);
  //console.log(JSON.stringify(postData));
  var actionDesc='';
  switch(action){
    case 'insert':
      actionDesc='添加';
      break;
    case 'update':
      actionDesc='更新';
      break;
    case 'delete':
      actionDesc='删除';
      break;
    case 'expand':
      actionDesc='扩容';
      break;
    case 'shrink':
      actionDesc='缩容';
      break;
    case 'deploy':
      actionDesc='上线';
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

//view
var view=function(type,idx){
  NProgress.start();
  var url='',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  switch(type){
    case 'service':
      url='/api/for_layout/service.php';
      title='查看服务详情 - '+idx;
      postData={"action":"info","fIdx":idx};
      break;
    case 'pool':
      url='/api/for_layout/pool.php';
      title='查看服务池详情 - '+idx;
      postData={"action":"info","fIdx":idx};
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
            if(typeof(locale_messages.layout)) locale = locale_messages.layout;
            switch(type){
              case 'service':
                $.each(data.content,function(k,v){
                  if(v=='') v='空';
                  if(typeof(locale[k])!='undefined') k=locale[k];
                  text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                });
                break;
              case 'pool':
                $.each(data.content,function(k,v){
                  if(v=='') v='空';
                  if(k=='tasks') v=JSON.stringify(v);
                  if(k=='vm_type') v=getName('vm_type',v);
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
    case 'service':
      var disabled=false;
      if($('#cluster_id').val()=='') disabled=true;
      if($('#name').val()=='') disabled=true;
      if($('#desc').val()=='') disabled=true;
      if($('#service_type').val()=='') disabled=true;
      if($('#docker_image').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'pool':
      var disabled=false;
      if($('#service_id').val()=='') disabled=true;
      if($('#name').val()=='') disabled=true;
      if($('#desc').val()=='') disabled=true;
      //if($('#sd_id').val()=='') disabled=true;
      if($('#vm_type').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'addnode':
      var disabled=false;
      if($('#service_pool').val()=='') disabled=true;
      var ip = $('#nodes').val();
      var arrIp = ip.split(/[\s\n\r\,\;\:\#\_]+/);
      for(var i=0;i<arrIp.length;i++){
        if(arrIp[i]!=''){
          if(!checkIp(arrIp[i])){
            disabled=true;
            break;
          }
        }
      }
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'expand':
      var disabled=false;
      var pool=$('#pool').val();
      var num=$('#num').val();
      if(pool=='') disabled=true;
      if(num=='') disabled=true;
      if(!disabled){
        if(num<1){
          disabled=true;
          pageNotify('warning','输入错误','扩容数量必须大于0!')
        }else{
          if(num>cache.pool_quota[pool]){
            disabled=true;
            pageNotify('warning','配额余量不足','请追加配额或者减少扩容数量!')
          }
        }
      }
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'shrink':
      var disabled=false;
      if($('#pool').val()=='') disabled=true;
      if($('#ip').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'deploy':
      var disabled=false;
      if($('#pool').val()=='') disabled=true;
      if($('#tag').val()=='') disabled=true;
      if($('#task_name').val()=='') disabled=true;
      if($('#max_num').val()=='') disabled=true;
      if($('#max_ratio').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
  }
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
      case 'service':
        switch(action){
          case 'del':
            modalTitle='删除服务';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 服务 : '+idx+'<br>服务名称 : '+desc;
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
      case 'pool':
        switch(action){
          case 'del':
            modalTitle='删除服务池';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 服务池 : '+idx+' (隶属服务: '+desc+')';
            modalBody=modalBody+'</div>';
            modalBody=modalBody+'</div>';
            modalBody=modalBody+'<input type="hidden" id="id" name="id" value="'+idx+'">';
            modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="delete">';
            break;
          case 'addnode':
            modalTitle='添加节点';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="service" class="col-sm-2 control-label">服务池</label>' +
                '<div class="col-sm-10">' +
                '<select class="form-control" id="id" name="id" disabled><option value="'+idx+'">'+desc+'</option></select>' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="service" class="col-sm-2 control-label">IP</label>' +
                '<div class="col-sm-10">' +
                '<textarea rows="6" class="form-control" id="nodes" name="nodes" style="font-family: \'Lucida Console\';" placeholder="支持换行,逗号,分号,空格分割" onkeyup="check(\'addnode\')"></textarea>' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<input type="hidden" id="page_other" name="page_other" value="addnode">';
            modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="insert">';
            btnDisable=true;
            break;
          default:
            modalTitle='非法请求';
            notice='<div class="note note-danger">错误信息：参数错误</div>';
            pageNotify('error','非法请求！','错误信息：参数错误');
            break;
        }
        break;
      case 'node':
        switch(action){
          case 'add':
            modalTitle='添加节点';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="service_pool" class="col-sm-2 control-label">服务池</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="service_pool" name="service_pool" onkeyup="check(\'addnode\')" value="'+desc+'" readonly>' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="nodes" class="col-sm-2 control-label">IP</label>' +
                '<div class="col-sm-10">' +
                '<textarea rows="6" class="form-control" id="nodes" name="nodes" style="font-family: \'Lucida Console\';" placeholder="IP,eg:支持 换行,逗号,分号,空格 分割" onkeyup="check(\'addnode\')"></textarea>' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="insert">';
            btnDisable=true;
            break;
          case 'del':
            modalTitle='删除节点';
            var pool=[],ips=[],ipid=[],info=[],poolName=[];
            if(idx&&desc){
              var aIdx=idx.split(',');
              pool.push({"id":aIdx[0],"name":aIdx[1]});
              ips.push(desc);
              ipid.push(ipidx);
            }else{
              $('input:checkbox[id=list]:checked').each(function(){
                info=$(this).val().split(',');
                if($.inArray(info[1],poolName) == -1){
                  poolName.push(info[1]);
                  pool.push({"id":info[0],"name":info[1]});
                }
                if($.inArray(info[2],ips) == -1) ips.push(info[2]);
                if($.inArray(info[3],ipid) == -1) ipid.push(info[3]);
              });
            }
            if(pool.length>1||ipid.length==0) btnDisable=true;
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<h4 class="text-danger"><strong>【警告】</strong>此处删除节点不会同步删除ECS, 如需同步删除ECS, 请使用缩容!</h4>';
            modalBody+='<h5><strong>当前操作, 将影响如下操作:</strong></h5>';
            modalBody+='<p><strong class="text-primary">涉及服务池</strong>: 共 <span class="badge badge-danger">'+pool.length+'</span> 个 <strong class="text-danger">(每次只支持操作一个服务池)</strong></p>';
            modalBody+='<div class="col-sm-12" style="margin-bottom: 5px;">';
            $.each(pool,function(k,v){
              modalBody+='<span class="col-sm-3">'+ v.name +'</span>';
            });
            modalBody+='</div>';
            modalBody+='<p><strong class="text-primary">涉及节点</strong>: 共 <span class="badge badge-danger">'+ipid.length+'</span> 个</p>';
            modalBody+='<div class="col-sm-12">';
            $.each(ips,function(k,v){
              modalBody+='<span class="col-sm-2">'+v+'</span>';
            });
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="nodes" name="nodes" value="'+ipid.toString()+'">';
            if(pool.length>0){
              modalBody+='<input type="hidden" id="id" name="id" value="'+pool[0]['id']+'">';
            }else{
              modalBody+='<input type="hidden" id="id" name="id" value="0">';
            }
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

var getList=function(type,idx){
  if(!type) type='cluster';
  var url='/api/for_layout/'+type+'.php?action=list';
  var postData={'pagesize':1000};
  $('#tab').val('service');
  var actionDesc='';
  switch (type){
    case 'cluster':
      actionDesc='集群';
      getTaskTpl();
      getHubbleBalance();
      getCloudCluster();
      break;
    case 'service':
      var fClusterId=$('#fClusterId').val();
      if(!fClusterId){
        pageNotify('warning','集群数据为空！','请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/cluster.php">创建集群</a>]！');
        $('#table-head').html('<tr><td>集群数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/cluster.php">创建集群</a>]！</td></tr>');
        $('#table-body').html('');
        $('#fClusterId').parent().parent().attr('hidden',true);
        $('#fService').parent().parent().attr('hidden',false);
        $('#fService').empty().append('<option value="">全部服务</option>').select2({width:'100%'});
        cache.service_id='';
        cache.pool_id='';
        $('#fPool').parent().parent().attr('hidden',true);
        return;
      }
      postData.fClusterId=fClusterId;
      $('#tab').val('pool');
      actionDesc='服务';
      getTaskTpl();
      getHubbleBalance();
      getCloudCluster();
      break;
    case 'pool':
      var fService=$('#fService').val();
      if(!fService){
        if(cache.service.length>0){
          if(cache.cluster_id==cache.service[0].cluster_id) fService=cache.service[0].id;
        }
      }
      if(!fService){
        pageNotify('warning','服务数据为空！','请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/service.php">创建服务</a>]！');
        $('#table-head').html('<tr><td>服务数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/service.php">创建服务</a>]！</td></tr>');
        $('#table-body').html('');
        $('#fClusterId').parent().parent().attr('hidden',true);
        $('#fService').parent().parent().attr('hidden',true);
        $('#fPool').parent().parent().attr('hidden',false);
        $('#fPool').empty().append('<option value="">全部服务池</option>').select2({width:'100%'});
        cache.pool_id='';
        return;
      }
      postData.fService=fService;
      $('#tab').val('node');
      actionDesc='服务池';
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
        if(data.content.length>0){
          switch (type){
            case 'cluster':
              cache.cluster = data.content;
              if(!idx) idx=cache.cluster_id;
              updateSelect('fClusterId',idx);
              break;
            case 'service':
              cache.service = data.content;
              if(!idx) idx=cache.service_id;
              updateSelect('fService',idx);
              break;
            case 'pool':
              cache.pool = data.content;
              if(!idx) idx=cache.pool_id;
              updateSelect('fPool',idx);
              break;
          }
        }else{
          switch (type){
            case 'cluster':
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/cluster.php">创建'+actionDesc+'</a>]！',false);
              $('#table-head').html('<tr><td>'+actionDesc+'数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/cluster.php">创建'+actionDesc+'</a>]！</td></tr>');
              $('#table-body').html('');
              $('#fClusterId').parent().parent().attr('hidden',false);
              $('#fService').parent().parent().attr('hidden',true);
              $('#fPool').parent().parent().attr('hidden',true);
              break;
            case 'service':
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/service.php">创建'+actionDesc+'</a>]！',false);
              $('#table-head').html('<tr><td>'+actionDesc+'数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/service.php">创建'+actionDesc+'</a>]！</td></tr>');
              $('#fClusterId').parent().parent().attr('hidden',true);
              $('#fService').parent().parent().attr('hidden',false);
              $('#fService').empty().append('<option value="">全部服务</option>').select2({width:'100%'});
              cache.service_id='';
              cache.pool_id='';
              $('#fPool').parent().parent().attr('hidden',true);
              break;
            case 'pool':
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/pool.php">创建'+actionDesc+'</a>]！',false);
              $('#table-head').html('<tr><td>'+actionDesc+'数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/pool.php ">创建'+actionDesc+'</a>]！</td></tr>');
              $('#table-body').html('');
              $('#fClusterId').parent().parent().attr('hidden',true);
              $('#fService').parent().parent().attr('hidden',true);
              $('#fPool').parent().parent().attr('hidden',false);
              $('#fPool').empty().append('<option value="">全部服务池</option>').select2({width:'100%'});
              cache.pool_id='';
              break;
          }
          $('.tooltips').each(function(){$(this).tooltip();});
          $('#table-body').html('');
        }
      }else{
        pageNotify('error','获取'+actionDesc+'失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      NProgress.done();
      pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}

var updateSelect=function(name,idx){
  var tSelect=$('#'+name),data='';
  switch(name){
    case 'fClusterId':
      data=cache.cluster;
      break;
    case 'cluster_id':
      data=cache.cluster;
      break;
    case 'fService':
      data=cache.service;
      break;
    case 'service_id':
      data=cache.cluster;
      break;
    case 'fPool':
      data=cache.pool;
      break;
    case 'tpl_expand':
      data=cache.task_tpl;
      break;
    case 'tpl_shrink':
      data=cache.task_tpl;
      break;
    case 'tpl_deploy':
      data=cache.task_tpl;
      break;
    case 'sd_id':
      data=cache.balance;
      break;
    case 'vm_type':
      data=cache.vm_type;
      break;
  }
  tSelect.empty();
  if(data.length>0){
    switch (name){
      case 'sd_id':
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.id + '">' + v.type + ' - ' + v.name + '</option>');
        });
        break;
      case 'vm_type':
        $.each(data,function(k,v){
          var t=(v.Name) ? v.Name : v.Id;
          tSelect.append('<option value="' + v.Id + '">' + v.Provider + ' - ' + t + '</option>');
        });
        break;
      default:
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.id + '">' + v.name + '</option>');
        });
        break;
    }
    if(idx){
      tSelect.val(idx).trigger('change');
      if(!tSelect.val()){
        tSelect.append('<option value="'+idx+'">'+idx+' - 此选项已不存在</option>');
        tSelect.val(idx).trigger('change');
      }
    }else{
      tSelect.val($('#'+name+' option:nth-child(1)').val()).trigger('change');
    }
  }else{
    if(idx){
      tSelect.append('<option value="'+idx+'">'+idx+'</option>');
    }else{
      tSelect.append('<option value="">请选择</option>');
    }
  }
}

var get = function (idx,tab) {
  if(!tab) tab=$('#tab').val();
  var url='/api/for_layout/'+tab+'.php',postData={};
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
              if(k=='tasks'){
                updateSelect('tpl_expand',v.expand);
                updateSelect('tpl_shrink',v.shrink);
                updateSelect('tpl_deploy',v.deploy);
                $('#tpl_expand').val(v.expand).trigger('change');
                $('#tpl_shrink').val(v.shrink).trigger('change');
                $('#tpl_deploy').val(v.deploy).trigger('change');
              }else{
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
              }
            });
            if(tab=='pool'){
              updateSelect('sd_id',$('#sd_id').val());
              updateSelect('vm_type',$('#vm_type').val());
            }
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
    case 'cluster':
      data=cache.cluster;
      for(var i=0;i<data.length;i++){
        if(data[i].id==idx){
          name= data[i].name;
          break;
        }
      }
      break;
    case 'service':
      data=cache.service;
      for(var i=0;i<data.length;i++){
        if(data[i].id==idx){
          name= data[i].name;
          break;
        }
      }
      break;
    case 'pool':
      data=cache.pool;
      for(var i=0;i<data.length;i++){
        if(data[i].id==idx){
          name= data[i].name;
          break;
        }
      }
      break;
    case 'vm_type':
      data=cache.vm_type;
      for(var i=0;i<data.length;i++){
        if(data[i].Id==idx){
          name=(data[i].Name)? data[i].Provider+' - '+data[i].Name : data[i].Provider+' - '+data[i].Id;
          break;
        }
      }
      if(name==idx) name+=' - 此模板已不存在';
      break;
    case 'balance':
      data=cache.balance;
      for(var i=0;i<data.length;i++){
        if(data[i].id==idx){
          name=(data[i].name)? data[i].type+' - '+data[i].name : data[i].type+' - '+data[i].id;
          break;
        }
      }
      if(name==idx) name+=' - 服务注册类型已不存在';
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


//获取模板列表
var getTaskTpl=function(){
  var actionDesc='模板列表';
  var postData={"action":"list","pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_layout/task_tpl.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.task_tpl = data.content;
        }else{
          pageNotify('info','加载'+actionDesc+'成功！','数据为空！');
        }
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}

//获取服务发现
var getHubbleBalance=function(){
  var actionDesc='服务发现列表';
  var postData={"action":"list","pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/balance.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.balance = data.content;
        }else{
          pageNotify('info','加载'+actionDesc+'成功！','数据为空！');
        }
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}

//获取机型模板
var getCloudCluster=function(){
  var actionDesc='机型模板列表';
  var postData={"action":"list","pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_cloud/cluster.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.vm_type = data.content;
        }else{
          pageNotify('info','加载'+actionDesc+'成功！','数据为空！');
        }
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}

//获取节点
var getNodes=function(){
  var actionDesc='节点列表';
  var idx=$('#pool').val();
  var list=$('#iplist');
  $('#ipnum').html('全选');
  $('#check_all').attr('disabled',true);
  $('#check_all').attr('checked',false);
  list.html('');
  if(!idx){
    return false;
  }
  $.ajax({
    type: "POST",
    url: '/api/for_layout/node.php',
    data: {"action":"list","fPool":idx,"pagesize":10000},
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          $.each(data.content, function (k,v) {
            list.append('<label style="width:130px;"><input type="checkbox" id="list" name="list[]" value="'+ v.ip +'" onchange="selectIp()" \>'+ v.ip+'</label>');
          });
          $('#ipnum').html('全选('+data.content.length+')');
          if(data.content.length>20) $('#iplist').css('height','200px');
          $('#check_all').attr('disabled',false);
        }else{
          pageNotify('info','加载'+actionDesc+'成功！','数据为空！');
        }
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}

var checkAll=function(o){
  $('[id=list]:checkbox').prop('checked', o.checked);
  $('#check_num').html($('[id=list]:checked').length);
  updateIp();
}

//勾选IP
var selectIp=function () {
  $('#check_num').html($('[id=list]:checked').length);
  updateIp();
}

//检查IP格式
var checkIp=function(ip) {
  if(ip){
    var re=/^(\d+)\.(\d+)\.(\d+)\.(\d+)$/;//正则表达式
    if(re.test(ip)){
      if( RegExp.$1<256 && RegExp.$2<256 && RegExp.$3<256 && RegExp.$4<256)
        return true;
    }
  }
  return false;
}

//待执行ip处理
var updateIp=function(){
  $('[id=list]:checkbox').each(function(){
    var idx = $.inArray(this.value,cache.ip);
    if(this.checked){
      if(idx == -1) cache.ip.push(this.value);
    }else{
      if(idx != -1) cache.ip.splice(idx,1);
    }
  });
  var str='';
  if(cache.ip.length>0){
    $.each(cache.ip,function(k,v){
      str+=(str)?','+v:v;
    });
  }
  $('#ip').val(str);
  $('#run_num').html(cache.ip.length);
  check();
}

//手动输入IP时
var manualIp = function () {
  var ip = $('#ip').val();
  var arrIp = ip.split(/[\s\n\r\,\;\:\#\_]+/);
  var disabled=false;
  var arr=[];
  for(var i=0;i<arrIp.length;i++){
    if(arrIp[i]!=''){
      if(!checkIp(arrIp[i])){
        disabled=true;
      }else{
        if($.inArray(arrIp[i],arr) == -1) arr.push(arrIp[i]);
        if($.inArray(arrIp[i],cache.ip) == -1) cache.ip.push(arrIp[i]);
      }
    }
  }
  if(arr.length==0) cache.ip=[];
  for(var i=0;i<cache.ip.length;i++){
    var idx=$.inArray(cache.ip[i],arr);
    if(idx == -1 ) cache.ip.splice(idx,1);
  }
  $('#run_num').html(cache.ip.length);
  if(disabled){
    $('#btnCommit').attr('disabled',disabled);
  }else{
    check('run');
  }
}

//过滤指定IP
var autoCheck=function(checked){
  var  ip_keyword = $('#check_input').val();
  ip_keyword = $.trim(ip_keyword);
  if(ip_keyword == '') return false;
  $('[id=list]:checkbox').each(function(){
    if(this.value.indexOf(ip_keyword) < 0) return true;
    //$(this).attr('checked',checked).trigger('change');
    if(checked){
      $(this).iCheck('check');
    }else{
      $(this).iCheck('uncheck');
    }
  });
  selectIp();
}

var getBiz=function(){
  var biz='',cluster_id=$('#fClusterId').val();
  for(var i=0;i<cache.cluster.length;i++){
    if(cache.cluster[i].id == cluster_id){
      biz=cache.cluster[i].biz;
      break;
    }
  }
  return biz;
}

var checkIp=function(ip) {
  if(ip){
    var re=/^(\d+)\.(\d+)\.(\d+)\.(\d+)$/;//正则表达式
    if(re.test(ip)){
      if( RegExp.$1<256 && RegExp.$2<256 && RegExp.$3<256 && RegExp.$4<256)
        return true;
    }
  }
  return false;
}

var getQuota=function(idx){
  var url='/api/for_cloud/quota.php?action=info';
  var pool=$('#pool').val();
  if(!pool) return false;
  var idx=(typeof cache.pool_vm[pool] != 'undefined')?cache.pool_vm[pool]:'';
  if(!idx) return false;
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
        $('#quota').html(data['content']['credit']);
        cache.pool_quota[pool]=data['content']['credit'];
      }else{
        $('#quota').html('获取配额余量失败');
      }
      $('.tooltips').each(function(){$(this).tooltip();});
    },
    error: function (){
      NProgress.done();
      pageNotify('error','获取'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}