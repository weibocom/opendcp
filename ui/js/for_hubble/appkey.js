cache = {
  page: 1,
  detail: {
    'id':'0',
    'appkey':'无',
    'name':'ERROR',
    'describe':'非法操作'
  },
  interface: {},
  privilege: [],
  checked: [],
  appkey: [],
}

var reset = function(){
  $('#fIdx').val('');
  list(1);
}

var list = function(page,tab) {
  $('.popovers').each(function(){$(this).popover('hide');});
  NProgress.start();
  var postData={};
  if(!tab){
    tab=$('#tab').val();
  }
  if(tab!='appkey'&&tab!='interface'){
    tab='appkey';
  }
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'appkey':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_3').attr('class','hidden');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'add\')"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fIdx":fIdx};
      break;
    case 'interface':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','active');
      $('#tab_3').attr('class','hidden');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'add\')"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fIdx":fIdx};
      break;
  }
  var url='/api/for_hubble/'+tab+'.php';
  if (!page) {
    page = cache.page;
  }else{
    cache.page = page;
  }
  var url=url+'?action=list&page=' + page;
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (listdata) {
      if(listdata.code==0){
        if(tab=='appkey') cache.appkey=listdata.content;
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
            console.log(i+' '+p1);
            li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a onclick="list('+i+')">'+i+'</a></li>';
          }else{
            li='<li><a onclick="list('+i+')">'+i+'</a></li>';
          }
        }else{
          if(i==p2){
            if(p2<data.pageCount-1){
              console.log(i+' '+p2);
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
        if(tab!='node'){
          td = '<td>' + v.i + '</td>';
          tr.append(td);
        }
        var btnAdd='',btnEdit='',btnDel='';
        switch(tab){
          case 'appkey':
            td = '<td>' + v.appkey + '</td>';
            tr.append(td);
            td = '<td>' + v.name + '</td>';
            tr.append(td);
            td = '<td>' + v.describe + '</td>';
            tr.append(td);
            btnEdit = '<a class="text-primary tooltips" title="授权列表" data-toggle="modal" data-target="#myChildModal" href="edit_privilege.php?idx=' + v.id + '"><i class="fa fa-delicious"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.appkey+'\',\''+v.name+'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'interface':
            td = '<td title="'+ v.id +'">' + v.addr + '</td>';
            tr.append(td);
            td = '<td>' + v.desc + '</td>';
            tr.append(td);
            td = '<td>' + v.method + '</td>';
            tr.append(td);
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+ v.addr +'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'privilege':
            td = '<td>' + v.addr_id + '</td>';
            tr.append(td);
            td = '<td>' + v.key_id + '</td>';
            tr.append(td);
            td = '<td>' + v.describe + '</td>';
            tr.append(td);
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+ v.desc +'\')"><i class="fa fa-trash-o"></i></a>';
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
  var url='/api/for_hubble/'+tab+'.php';
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

//commit check
var check=function(tab){
  if(!tab) tab=$('#tab').val();
  switch(tab){
    case 'appkey':
      var disabled=false;
      if($('#name').val()=='') disabled=true;
      if($('#desc').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'addinterface':
      var disabled=false;
      if($('#addr').val()=='') disabled=true;
      if($('#method').val()=='') disabled=true;
      if($('#desc').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'privilege':
      var disabled=false;
      if($('#addr_id').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
  }
}

//二次确认
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
      case 'appkey':
        switch(action){
          case 'add':
            modalTitle='生成AppKey';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="name" class="col-sm-2 control-label">名称</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="name" name="name" onkeyup="check()" placeholder="名称,eg:XX">' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="desc" class="col-sm-2 control-label">描述</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="desc" name="desc" onkeyup="check()" placeholder="描述,eg:XX系统用">' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="insert">';
            btnDisable=true;
            break;
          case 'del':
            modalTitle='删除AppKey';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> AppKey : '+idx+'<br>' +
                '名称 : '+desc;
            modalBody=modalBody+'</div>';
            modalBody=modalBody+'</div>';
            modalBody=modalBody+'<input type="hidden" id="key" name="key" value="'+idx+'">';
            modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="delete">';
            break;
          default:
            modalTitle='非法请求';
            notice='<div class="note note-danger">错误信息：参数错误</div>';
            pageNotify('error','非法请求！','错误信息：参数错误');
            break;
        }
        break;
      case 'interface':
        switch(action){
          case 'add':
            modalTitle='添加接口';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="addr" class="col-sm-2 control-label">接口地址</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="addr" name="addr" onkeyup="check(\'addinterface\')" placeholder="接口,eg:/api/auth.php">' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="method" class="col-sm-2 control-label">方法</label>' +
                '<div class="col-sm-10">' +
                '<select class="form-control" id="method" name="method" onchange="check(\'addinterface\')">' +
                '<option value="GET">GET</option>' +
                '<option value="POST">POST</option>' +
                '<option value="PUT">PUT</option>' +
                '<option value="DELETE">DELETE</option>' +
                '</select>' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<div class="form-group">' +
                '<label for="desc" class="col-sm-2 control-label">描述</label>' +
                '<div class="col-sm-10">' +
                '<textarea rows="3" class="form-control" id="desc" name="desc" style="font-family: \'Lucida Console\';" placeholder="描述,eg:调用参数,返回结果" onkeyup="check(\'addinterface\')"></textarea>' +
                '</div>' +
                '</div>';
            modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="insert">';
            btnDisable=true;
            break;
          case 'del':
            modalTitle='删除接口';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 接口ID : '+idx+'<br/>' +
                '接口地址 : '+desc;
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

//modal list
var modalList = function(page,tab,fidx) {
  NProgress.start();
  var postData={};
  if(!tab||!fidx){
    pageNotify('error','加载失败！','错误信息：非法请求');
    NProgress.done();
  }
  switch(tab){
    case 'appkey':
      var url='/api/for_hubble/appkey.php';
      break;
    default:
      pageNotify('error','加载失败！','错误信息：非法请求');
      NProgress.done();
      break;
  }
  postData={"action":"info","fIdx":fidx};
  if (!page) {
    page = 1;
  }
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (listdata) {
      if(listdata.code==0){
        var pageinfo = $("#modal_table-pageinfo");//分页信息
        var paginate = $("#modal_table-paginate");//分页代码
        var head = $("#modal_table-head");//数据表头
        var body = $("#modal_table-body");//数据列表
        //清除当前页面数据
        pageinfo.html("");
        paginate.html("");
        head.html("");
        body.html("");
        //定位appkey
        for(var i=0;i<cache.appkey.length;i++){
          if(cache.appkey[i].id==fidx){
            cache.detail=cache.appkey[i];
            $('#myModalLabel').html('AppKey授权 - '+fidx+' - '+cache.appkey[i].name);
            break;
          }
        }
        cache.privilege=listdata.content.content;
        getChecked();
        getInterface(fidx, page, pageinfo, paginate, head, body);
        //生成页面
        //生成分页
        //processModalPage(listdata, pageinfo, paginate, tab, fidx);
        //生成列表
        //processModalBody(listdata, fidx, tab, head, body);
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

//modal 生成分页
var processModalPage = function(data,pageinfo,paginate,findex){
  var tab='appkey';
  var begin = data.pageSize * ( data.page - 1 ) + 1;
  var end = ( data.count > begin + data.pageSize - 1 ) ? begin + data.pageSize - 1 : data.count;
  pageinfo.html('Showing '+begin+' to '+end+' of '+data.count+' records');
  var p1=(data.page-1>0)?data.page-1:1;
  var p2=data.page+1;
  prev='<li><a onclick="modalList('+p1+',\''+tab+'\',\''+findex+'\')"><i class="fa fa-angle-left"></i></a></li>';
  paginate.append(prev);
  for (var i = 1; i <= data.pageCount; i++) {
    var li='';
    if(i==data.page){
      li='<li class="active"><a onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
    }else{
      if(i==1||i==data.pageCount){
        li='<li><a onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
      }else{
        if(i==p1){
          if(p1>2){
            console.log(i+' '+p1);
            li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
          }else{
            li='<li><a onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
          }
        }else{
          if(i==p2){
            if(p2<data.pageCount-1){
              console.log(i+' '+p2);
              li='<li><a onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>'+"\n"+'<li class="disabled"><a href="#">...</a></li>';
            }else{
              li='<li><a onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
            }
          }
        }
      }
    }
    paginate.append(li);
  }
  if(p2>data.pageCount) p2=data.pageCount;
  next='<li class="next"><a title="Next" onclick="modalList('+p2+',\''+tab+'\',\''+findex+'\')"><i class="fa fa-angle-right"></i></a></li>';
  paginate.append(next);
}

//modal 生成列表
var processModalBody = function(data,head,body){
  var td="";
  var title=['授权','#','AppKey','AppKey名称','接口','接口描述'];
  if(title){
    var tr = $('<tr></tr>');
    for (var i = 0; i < title.length; i++) {
      var v = title[i];
      td = '<th>' + v + '</th>';
      tr.append(td);
    }
    head.html(tr);
  }
  if(cache.interface){
    if(cache.interface.content.length>0){
      for (var i = 0; i < cache.interface.content.length; i++) {
        var v = cache.interface.content[i];
        var tr = $('<tr></tr>');
        var checked='';
        if($.inArray(v.id,cache.checked) != -1) checked='checked="checked"';
        td = '<td><input type="checkbox" id="pri_'+v.id+'" onchange="setPrivilege(\''+ v.id +'\',\''+ cache.detail.id +'\',\''+ v.id +'\')" '+checked+'\></td>';
        tr.append(td);
        td = '<td>' + (i+1) + '</td>';
        tr.append(td);
        td = '<td title="'+cache.detail.id+'">' + cache.detail.appkey + '</td>';
        tr.append(td);
        td = '<td>' + cache.detail.name + '</td>';
        tr.append(td);
        td = '<td title="'+v.id+'">' + v.addr + '</td>';
        tr.append(td);
        td = '<td>' + v.desc + '</td>';
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

//已授权接口ID
var getChecked = function (){
  cache.checked=[];
  if(cache.privilege){
    for(var i=0;i<cache.privilege.length;i++){
      cache.checked.push(cache.privilege[i].interface_id);
    }
  }
}

//获取接口列表
var getInterface = function (fidx, page, pageinfo, paginate, head, body) {
  var postData={};
  var url='/api/for_hubble/interface.php';
  postData={"action":"list","pagesize":10,"page":page};
  if(head && body){
    $.ajax({
      type: "POST",
      url: url,
      data: postData,
      dataType: "json",
      success: function (data) {
        //执行结果提示
        if(data.code==0){
          if(typeof(data.content)!='undefined'){
            cache.interface=data;
            processModalPage(data, pageinfo, paginate, fidx);
            processModalBody(data, head, body);
          }else{
            pageNotify('warning','数据为空！');
          }
        }else{
          pageNotify('error','加载失败！','错误信息：'+data.msg);
        }
      },
      error: function (){
        pageNotify('error','加载失败！','错误信息：接口不可用');
      }
    });
  }else{
    pageNotify('warning','加载失败！','错误信息：参数错误');
  }
}

//修改授权
var setPrivilege = function (o,appkey,addr){
  var url='/api/for_hubble/privilege.php';
  var action='',postData={'key_id':appkey,'addr_id':addr},ret='';
  if($('#pri_'+o).prop('checked')){
    action='insert';
    actionDesc='添加接口授权';
    ret=false;
  }else{
    action='delete';
    actionDesc='移除接口授权';
    ret=true;
  }
  if(appkey&&addr){
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
          $('#pri_'+o).attr('checked',ret);
        }
      },
      error: function (){
        pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
        $('#pri_'+o).attr('checked',ret);
      }
    });
  }
}

