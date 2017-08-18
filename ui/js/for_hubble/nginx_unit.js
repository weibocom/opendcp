cache = {
  page: 1,
  group_id: 0,
  group: [],
  unit: [],
  copy: {
    ip: [],
  },
  ip: [], //选中IP列表
}

var reset = function(){
  $('#fIdx').val('');
  list(1);
}

var list = function(page,tab,idx) {
  $('.popovers').each(function(){$(this).popover('hide');});
  NProgress.start();
  var postData={};
  if(!tab){
    tab=$('#tab').val();
  }
  if(tab!='upstream'&&tab!='unit'&&tab!='node'){
    tab='unit';
  }
  var fGroup=$('#fGroup').val();
  var fUnit=$('#fUnit').val();
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'upstream':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_3').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php?action=add&par_id='+ $('#fGroup').val() +'&par_name='+ $('#fGroup').find("option:selected").text() +'"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fGroup":fGroup,"fIdx":fIdx};
      cache.group_id=fGroup;
      $('#fGroup').parent().parent().attr('hidden',false);
      $('#fUnit').parent().parent().attr('hidden',true);
      break;
    case 'unit':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','active');
      $('#tab_3').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php?action=add&par_id='+ $('#fGroup').val() +'&par_name='+ $('#fGroup').find("option:selected").text() +'"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fGroup":fGroup,"fIdx":fIdx};
      cache.group_id=fGroup;
      $('#fGroup').parent().parent().attr('hidden',false);
      $('#fUnit').parent().parent().attr('hidden',true);
      break;
    case 'node':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','');
      $('#tab_3').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-danger" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\')"> 批量删除 <i class="fa fa-trash-o"></i></a>');
      if(idx){
        fUnit=idx;
        $('#fUnit').val(fUnit);
      }
      postData={"fUnit":fUnit,"fIdx":fIdx};
      $('#fGroup').parent().parent().attr('hidden',true);
      $('#fUnit').parent().parent().attr('hidden',false);
      break;
  }
  var url='/api/for_hubble/nginx_'+tab+'.php';
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
            li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
          }else{
            li='<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
          }
        }else{
          if(i==p2){
            if(p2<data.pageCount-1){
              console.log(i+' '+p2);
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
      if(tab=='unit') cache.unit=data.content;
      cache.copy.ip=[];
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        if(tab!='node'){
          td = '<td>' + v.i + '</td>';
          tr.append(td);
        }
        var btnAdd='',btnEdit='',btnDel='',btnPublish='';
        switch(tab){
          case 'upstream':
            td = '<td><a class="tooltips" title="查看文件" data-toggle="modal" data-target="#myViewModal" onclick="view(\'upstream\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            td = '<td><a class="tooltips" title="查看分组详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'group\',\''+v.group_id+'\')">' + getName('group',v.group_id) + '</a></td>';
            tr.append(td);
            td = '<td>' + v.opr_user + '</td>';
            tr.append(td);
            td = '<td>' + v.update_time + '</td>';
            tr.append(td);
            btnPublish='<a class="text-primary tooltips" title="发布" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'publish\',\''+v.id+'\',\''+v.name+'\')"><i class="fa fa-upload"></i></a>' +
              ' <a class="text-info tooltips" title="发布结果" data-toggle="modal" data-target="#myViewModal" onclick="viewResult(\''+v.release_id+'\',\''+v.name+'\')"><i class="fa fa-comment-o"></i></a>';
            td = '<td>' + btnPublish + '</td>';
            tr.append(td);
            btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php?action=edit&idx=' + v.id + '"><i class="fa fa-edit"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'unit':
            td = '<td><a class="tooltips" title="查看单元详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'unit\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            td = '<td><a class="tooltips" title="查看分组详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'group\',\''+v.group_id+'\')">' + getName('group',v.group_id) + '</a></td>';
            tr.append(td);
            var tA='<a class="tooltips" title="查看节点列表" onclick="getList(\'unit\',\''+ v.id+'\')" id="ip_'+ v.id +'"><i class="fa fa-bars"></i></a>' +
              ' <a class="tooltips" title="添加节点" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'addnode\',\''+v.id+'\',\''+v.name+'\')"><i class="fa fa-plus"></i></a>';
            td = '<td>'+tA+'</td>';
            tr.append(td);
            td = '<td><a class="tooltips" title="查看主配置" href="nginx_conf.php?group='+ v.group_id +'&idx='+ v.id +'"><i class="fa fa-file-text"></i></a></td>';
            tr.append(td);
            td = '<td>' + v.opr_user + '</td>';
            tr.append(td);
            td = '<td>' + v.create_time + '</td>';
            tr.append(td);
            btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php?action=edit&idx=' + v.id + '&par_id='+ v.group_id +'&par_name='+ getName('group',v.group_id) +'"><i class="fa fa-edit"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\',\''+getName('group',v.group_id)+'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'node':
            cache.copy.ip.push(v.ip);
            td = '<td><input type="checkbox" id="list" name="list[]" value="'+ v.unit_id + ',' + getName('unit',v.unit_id) + ',' + v.ip+',' + v.id+'" /></td>';
            tr.append(td);
            td = '<td>' + v.i + '</td>';
            tr.append(td);
            td = '<td>' + getName('unit',v.unit_id) + '</td>';
            tr.append(td);
            td = '<td>' + v.ip + '</td>';
            tr.append(td);
            td = '<td>' + v.create_time + '</td>';
            tr.append(td);
            td = '<td>' + v.opr_user + '</td>';
            tr.append(td);
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\'' + v.unit_id + ',' + getName('unit',v.unit_id) + '\',\'' + v.ip + '\',\'' + v.id + '\')"><i class="fa fa-trash-o"></i></a>';
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
    case 'addnode': tab='node'; break;
  }
  var url='/api/for_hubble/nginx_'+tab+'.php';
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
    case 'publish':
      actionDesc='发布指令下发';
      postData.unit_ids='';
      if(typeof postData['ulist']){
        $.each(postData['ulist'],function(k,v){
          postData.unit_ids+=(postData.unit_ids)?','+v:v;
        });
        delete postData['ulist'];
      }
      delete postData['selectAll'];
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
  url='/api/for_hubble/nginx_'+type+'.php';
  postData={"action":"info","fIdx":idx};
  switch(type){
    case 'group':
      title='查看分组详情 - '+idx;
      break;
    case 'unit':
      title='查看单元详情 - '+idx;
      break;
    case 'upstream':
      title='查看文件 - '+idx;
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
            if(typeof(locale_messages.hubble.nginx)) locale = locale_messages.hubble.nginx;
            switch(type){
              case 'upstream':
                var str='';
                $.each(data.content,function(k,v){
                  if(k=='content'){
                    str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                      '<h5 class="col-sm-12 text-primary"><strong>Content</strong></h5>' +
                      '<textarea class="form-control" rows="12">'+v+'</textarea>';
                  }else{
                    if(v=='') v='空';
                    if(typeof(locale[k])!='undefined') k=locale[k];
                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                  }
                });
                text+=str;
                break;
              default:
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
    title='非法请求 - '+action;
    text='<div class="note note-danger">错误信息：参数错误</div>';
    $('#myViewModalLabel').html(title);
    $('#myViewModalBody').html(text);
    NProgress.done();
  }
}

var viewResult=function(idx,desc){
  NProgress.start();
  var url='',title='查看发布结果 - '+desc,text='',height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  postData={"action":"state","fIdx":idx};
  url='/api/for_hubble/alteration.php';
  if(idx!=''){
    if(idx==0){
      pageNotify('info','此文件暂未发布');
      text='<div class="note note-info">此文件暂未发布</div>';
      $('#myViewModalLabel').html(title);
      $('#myViewModalBody').html(text);
      NProgress.done();
    }else{
      $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
          //执行结果提示
          if(data.code==0){
            if(typeof(data.content)!='undefined'){
              if(typeof data.content.state != 'undefined'){
                text+='<h5 class="col-sm-12">文件名称 : '+desc+'</h5>';
                switch(data.content.state){
                  case 1: text+='<h5 class="col-sm-12">发布结果 : <span class="badge bg-blue">执行中</span></h5>';break;
                  case 2: text+='<h5 class="col-sm-12">发布结果 : <span class="badge bg-green">成功</span></h5>';break;
                  case 3: text+='<h5 class="col-sm-12">发布结果 : <span class="badge bg-red">失败</span></h5>';break;
                  default: text+='<h5 class="col-sm-12">发布结果 : <span class="badge bg-orange">未知状态</span></h5>';break;
                }
                var text1='',text2='',text3='',text0='';
                $.each(data.content.detail,function(k,v){
                  switch(v.state){
                    case 1: text1+='<span class="col-sm-2"><a class="tooltips text-primary" title="查看执行日志" data-toggle="modal" data-target="#myViewChildModal" onclick="viewIpResult(\''+ data.content['X-CORRELATION-ID'] +'\',\''+ v.ip +'\')" ">'+ v.ip +'</a></span>';break;
                    case 2: text2+='<span class="col-sm-2"><a class="tooltips text-success" title="查看执行日志" data-toggle="modal" data-target="#myViewChildModal" onclick="viewIpResult(\''+ data.content['X-CORRELATION-ID'] +'\',\''+ v.ip +'\')" ">'+ v.ip +'</a></span>';break;
                    case 3: text3+='<span class="col-sm-2"><a class="tooltips text-danger" title="查看执行日志" data-toggle="modal" data-target="#myViewChildModal" onclick="viewIpResult(\''+ data.content['X-CORRELATION-ID'] +'\',\''+ v.ip +'\')" ">'+ v.ip +'</a></span>';break;
                    default: text0+='<span class="col-sm-2"><a class="tooltips" title="查看执行日志" data-toggle="modal" data-target="#myViewChildModal" onclick="viewIpResult(\''+ data.content['X-CORRELATION-ID'] +'\',\''+ v.ip +'\')" ">'+ v.ip +'</a></span>';break;
                  }
                });
                if(text1) text1='<div class="col-sm-12" style="padding-left: 30px;">'+text1+'</div>';
                if(text2) text2='<div class="col-sm-12" style="padding-left: 30px;">'+text2+'</div>';
                if(text3) text3='<div class="col-sm-12" style="padding-left: 30px;">'+text3+'</div>';
                if(text0) text0='<div class="col-sm-12" style="padding-left: 30px;">'+text0+'</div>';
                text+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>';
                text+='<h5 class="col-sm-12 text-primary"><strong>执行中节点 : '+((text1)?'':'无')+'</strong></h5>'+text1;
                text+='<h5 class="col-sm-12 text-danger"><strong>失败节点 : '+((text3)?'':'无')+'</strong></h5>'+text3;
                text+='<h5 class="col-sm-12 text-success"><strong>成功节点 : '+((text2)?'':'无')+'</strong></h5>'+text2;
                if(text0) text+='<h5 class="col-sm-12"><strong>未知状态节点 : </strong></h5>'+text0;
              }
              if(!text){
                pageNotify('warning','数据为空！');
                text='<div class="note note-warning">'+JSON.stringify(data)+'</div>';
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
    }
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
    case 'upstream':
      var disabled=false;
      if($('#group_id').val()=='') disabled=true;
      if($('#name').val()=='') disabled=true;
      if($('#content').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'unit':
      var disabled=false;
      if($('#group_id').val()=='') disabled=true;
      if($('#name').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'addnode':
      var disabled=false;
      if($('#unit_id').val()=='') disabled=true;
      var ip = $('#ips').val();
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
    case 'publish':
      var disabled=false;
      if($('#myModalBody input:checkbox[id=ulist]:checked').length==0) disabled=true;
      if($('#tunnel').val()=='') disabled=true;
      if($('#script_id').val()=='') disabled=true;
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
      case 'upstream':
        switch(action){
          case 'publish':
            modalTitle='发布Upstream文件';
            modalBody+='<div class="form-group">' +
              '<label for="upstream_id" class="col-sm-2 control-label">文件名称</label>' +
              '<div class="col-sm-10">' +
              '<select class="form-control" id="upstream_id" name="upstream_id" disabled><option value="'+idx+'">'+desc+'</option></select>' +
              '</div>' +
              '</div>';
            modalBody+='<div class="form-group">' +
              '<label for="upstream_id" class="col-sm-2 control-label">目标单元</label>' +
              '<div class="col-sm-10 profile_details">' +
              '<div class="well profile_view col-sm-12">' +
              '<div class="col-sm-12" style="min-height:26px;" id="publish_unit"></div>' +
              '</div>' +
              '</div>' +
              '</div>';
            modalBody+='<div class="form-group">' +
              '<label for="tunnel" class="col-sm-2 control-label">发布方式</label>' +
              '<div class="col-sm-10">' +
              '<select class="form-control" id="tunnel" name="tunnel" onchange="check(\'publish\')"><option value="ANSIBLE">Ansible</option></select>' +
              '</div>' +
              '</div>';
            modalBody+='<div class="form-group">' +
              '<label for="script_id" class="col-sm-2 control-label">发布脚本</label>' +
              '<div class="col-sm-10">' +
              '<select class="form-control" id="script_id" name="script_id" onchange="check(\'publish\')"><option value="">请选择</option></select>' +
              '</div>' +
              '</div>';
            modalBody+='<input type="hidden" id="upstream_id" name="upstream_id" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="publish">';
            btnDisable=true;
            getShell();
            getUnit();
            break;
          case 'del':
            modalTitle='删除Upstream文件';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 文件名 : '+desc;
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
      case 'unit':
        switch(action){
          case 'del':
            modalTitle='删除单元';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 单元 : '+desc+' (隶属分组: '+ipidx+')';
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="id" name="id" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="delete">';
            break;
          case 'addnode':
            modalTitle='添加节点';
            modalBody+='<div class="form-group">' +
                '<label for="unit_id" class="col-sm-2 control-label">单元</label>' +
                '<div class="col-sm-10">' +
                '<select class="form-control" id="unit_id" name="unit_id" disabled><option value="'+idx+'">'+desc+'</option></select>' +
                '</div>' +
                '</div>';
            modalBody+='<div class="form-group">' +
                '<label for="ips" class="col-sm-2 control-label">IP</label>' +
                '<div class="col-sm-10">' +
                '<textarea rows="6" class="form-control" id="ips" name="ips" placeholder="支持换行,逗号,分号,空格分割" onkeyup="check(\'addnode\')"></textarea>' +
                '</div>' +
                '</div>';
            modalBody+='<input type="hidden" id="page_other" name="page_other" value="addnode">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="insert">';
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
            modalBody+='<div class="form-group">' +
                '<label for="unit_id" class="col-sm-2 control-label">单元</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="unit_id" name="unit_id" onkeyup="check(\'addnode\')" value="'+desc+'" readonly>' +
                '</div>' +
                '</div>';
            modalBody+='<div class="form-group">' +
                '<label for="ips" class="col-sm-2 control-label">IP</label>' +
                '<div class="col-sm-10">' +
                '<textarea rows="6" class="form-control" id="ips" name="ips" style="font-family: \'Lucida Console\';" placeholder="IP,eg:支持 换行,逗号,分号,空格 分割" onkeyup="check(\'addnode\')"></textarea>' +
                '</div>' +
                '</div>';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="insert">';
            btnDisable=true;
            break;
          case 'del':
            modalTitle='删除节点';
            var unit=[],ips=[],ipid=[],info=[],unitName=[];
            if(idx&&desc){
              var aIdx=idx.split(',');
              unit.push({"id":aIdx[0],"name":aIdx[1]});
              ips.push(desc);
              ipid.push(ipidx);
            }else{
              $('input:checkbox[id=list]:checked').each(function(){
                info=$(this).val().split(',');
                if($.inArray(info[1],unitName) == -1){
                  unitName.push(info[1]);
                  unit.push({"id":info[0],"name":info[1]});
                }
                if($.inArray(info[2],ips) == -1) ips.push(info[2]);
                if($.inArray(info[3],ipid) == -1) ipid.push(info[3]);
              });
            }
            console.log(ipid);
            if(unit.length>1||ipid.length==0) btnDisable=true;
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<h5><strong>当前操作, 将影响如下操作:</strong></h5>';
            modalBody+='<p><strong class="text-primary">涉及单元</strong>: 共 <span class="badge badge-danger">'+unit.length+'</span> 个 <strong class="text-danger">(每次只支持操作一个单元)</strong></p>';
            modalBody+='<div class="col-sm-12" style="margin-bottom: 5px;">';
            $.each(unit,function(k,v){
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
            if(unit.length>0){
              modalBody+='<input type="hidden" id="unit_id" name="unit_id" value="'+unit[0]['id']+'">';
            }else{
              modalBody+='<input type="hidden" id="unit_id" name="unit_id" value="0">';
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

var getList=function(type,idx,tab){
  if(!type) type='group';
  if(!tab) tab='unit';
  var url='/api/for_hubble/nginx_'+type+'.php?action=list';
  var postData={'pagesize':1000};
  $('#tab').val(tab);
  var actionDesc='';
  switch (type){
    case 'group':
      actionDesc='分组';
      if(cache.group_id) idx=cache.group_id;
      break;
    case 'unit':
      postData.fGroup=$('#fGroup').val();
      $('#tab').val('node');
      actionDesc='单元';
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
            case 'group':
              cache.group = data.content;
              updateSelect('fGroup',idx);
              break;
            case 'unit':
              cache.unit = data.content;
              updateSelect('fUnit',idx);
              break;
          }
        }else{
          switch (type){
            case 'group':
              $('#fGroup').parent().parent().attr('hidden',false);
              $('#fUnit').parent().parent().attr('hidden',true);
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/nginx_group.php">创建'+actionDesc+'</a>]！',false);
              $('#table-head').html('<tr><td>'+actionDesc+'数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/nginx_group.php">创建'+actionDesc+'</a>]！</td></tr>');
              break;
            case 'unit':
              $('#fGroup').parent().parent().attr('hidden',true);
              $('#fUnit').parent().parent().attr('hidden',false);
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/nginx_unit.php">创建'+actionDesc+'</a>]！',false);
              $('#table-head').html('<tr><td>'+actionDesc+'数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/nginx_unit.php">创建'+actionDesc+'</a>]！</td></tr>');
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
    case 'fGroup':
      data=cache.group;
      break;
    case 'group_id':
      data=cache.group;
      break;
    case 'fUnit':
      data=cache.unit;
      break;
    case 'script_id':
      data=cache.shell;
      break;
  }
  tSelect.empty();
  if(name=='script_id') tSelect.append('<option value="">请选择</option>');
  if(data){
    $.each(data,function(k,v){
      tSelect.append('<option value="' + v.id + '">' + v.name + '</option>');
    });
    if(idx){
      tSelect.val(idx).trigger('change');
    }else{
      tSelect.val($('#'+name+' option:nth-child(1)').val()).trigger('change');
    }
  }else{
    tSelect.append('<option value="">请选择</option>');
  }
}

var get = function (idx,tab) {
  if(!tab) tab=$('#tab').val();
  var url='/api/for_hubble/nginx_'+tab+'.php',postData={};
  switch (tab){
    default:
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
                    //$('#'+k).find("option[value='"+v+"']").attr("selected",true);
                    $('#'+k).val(v).trigger('change');
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
    case 'group':
      data=cache.group;
      $.each(data,function(k,v){
        if(v.id==idx){
          name= v.name;
        }
      });
      break;
    case 'unit':
      data=cache.unit;
      $.each(data,function(k,v){
        if(v.id==idx){
          name= v.name;
        }
      });
      break;
  }
  return name;
}

var checkAll=function(o,t){
  if(!t) t='list';
  $('[id='+t+']:checkbox').prop('checked', o.checked);
  if(t=='ulist'){
    check('publish');
  }
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

var getDefault=function(){
  var postData={"action":"info","fIdx":1};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/nginx_upstream.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(typeof data.content == 'object'){
          if(data.content.content.length>0){
            $('#content').val(data.content.content);
          }
        }
      }
    }
  });
}

var getShell=function(){
  var actionDesc='发布脚本列表';
  var postData={"action":"list","pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/shell.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.shell = data.content;
          updateSelect('script_id');
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

var getUnit=function(){
  var actionDesc='Nginx单元列表';
  var postData={"action":"list","pagesize":1000,"fGroup":$('#fGroup').val()};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/nginx_unit.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      var str='';
      if(data.code==0){
        if(data.content.length>0){
          cache.unit = data.content;
          str+='<label class="col-sm-12"><input type="checkbox" id="selectAll" onclick="checkAll(this,\'ulist\')"> 全选</label>';
          $.each(cache.unit,function(k,v){
            str+='<label class="col-sm-6"><input type="checkbox" id="ulist" name="ulist[]" value="'+ v.id +'" onchange="check(\'publish\')"> '+ v.name +'</label>';
          });
        }else{
          str+='<span class="text-danger" style="margin-bottom: 2px;">单元列表为空,请先创建单元</span><br/>';
        }
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：'+data.msg);
        str+='<span class="text-danger" style="margin-bottom: 2px;">data.msg</span><br/>';
      }
      $('#publish_unit').html(str);
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
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

var viewIpResult=function(idx,ip){
  NProgress.start();
  var url='',title='查看执行结果 - '+ip,text='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  postData={"action":"result","fIdx":idx,"fIp":ip};
  url='/api/for_hubble/oprlog.php';
  if(idx!=''&&ip!=''){
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
          $('#myViewChildModalLabel').html(title);
          $('#myViewChildModalBody').html(text);
          NProgress.done();
        },200);
      },
      error: function (){
        pageNotify('error','加载失败！','错误信息：接口不可用');
        text='<div class="note note-danger">错误信息：接口不可用</div>';
        $('#myViewChildModalLabel').html(title);
        $('#myViewChildModalBody').html(text);
        NProgress.done();
      }
    });
  }else{
    pageNotify('warning','错误操作！','错误信息：参数错误');
    title='非法请求';
    text='<div class="note note-danger">错误信息：参数错误</div>';
    $('#myViewChildModalLabel').html(title);
    $('#myViewChildModalBody').html(text);
    NProgress.done();
  }
}
