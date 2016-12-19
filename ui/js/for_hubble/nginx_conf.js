cache = {
  page: 1,
  unit: [],
  ip: [], //选中IP列表
  modal: {
    fIdx: '',
    fUnit: '',
    fUnitName: '',
  },
  file_ver: [],
  files: [],
  shell: [],
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
  if(tab!='main'&&tab!='ver'){
    tab='main';
  }
  var fUnit=$('#fUnit').val();
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'main':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php?action=add&par_id='+ $('#fUnit').val() +'&par_name='+ $('#fUnit').find("option:selected").text() +'"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fUnit":fUnit,"fIdx":fIdx};
      break;
    case 'ver':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php?action=add&par_id='+ $('#fUnit').val() +'&par_name='+ $('#fUnit').find("option:selected").text() +'"> 创建新版本 <i class="fa fa-plus"></i></a>' +
        '<a type="button" class="btn btn-primary" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'_auto.php?action=add&par_id='+ $('#fUnit').val() +'&par_name='+ $('#fUnit').find("option:selected").text() +'"> 生成新版本 <i class="fa fa-plus"></i></a>');
      postData={"fUnit":fUnit,"fIdx":fIdx};
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
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        td = '<td>' + v.i + '</td>';
        tr.append(td);
        var btnAdd='',btnEdit='',btnDel='';
        switch(tab){
          case 'main':
            td = '<td><a class="tooltips" title="查看文件" data-toggle="modal" data-target="#myViewModal" onclick="view(\'main\',\''+v.id+'\')">' + v.name + '</a>' +
              '<a class="pull-right tooltips" title="查看文件历史" data-toggle="modal" data-target="#myViewModal" href="list_ver.php?idx='+ v.name +'&par_id='+ v.unit_id +'&par_name='+getName('unit',v.unit_id)+'"><i class="fa fa-history"></i></a></td>';
            tr.append(td);
            td = '<td>' + v.version + '</td>';
            tr.append(td);
            td = '<td><a class="tooltips" title="查看单元详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'unit\',\''+v.unit_id+'\')">' + getName('unit',v.unit_id) + '</a></td>';
            tr.append(td);
            td = '<td>' + v.opr_user + '</td>';
            tr.append(td);
            td = '<td>' + v.create_time + '</td>';
            tr.append(td);
            btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php?action=edit&idx=' + v.id +'&par_id='+ v.unit_id +'&par_name='+getName('unit',v.unit_id) + '"><i class="fa fa-edit"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="废弃" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'deprecated\',\''+v.id+'\',\''+v.name+'\',\''+getName('unit',v.unit_id)+'\',\''+v.version+'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'ver':
            var t='';
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'ver\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            td = '<td>' + v.version + '</td>';
            tr.append(td);
            t = ( v.deprecated == 0 ) ? '<span class="badge bg-green">正常</span>' : '<span class="badge bg-red">已废弃</span>';
            td = '<td>' + t + '</td>';
            tr.append(td);
            t = ( v.is_release == 1 ) ? '<span class="badge bg-green">已发布</span>' : '<span class="badge">未发布</span>';
            td = '<td>' + t + '</td>';
            tr.append(td);
            td = '<td>' + getName('unit',v.unit_id) + '</td>';
            tr.append(td);
            td = '<td>' + v.opr_user + '</td>';
            tr.append(td);
            td = '<td>' + v.create_time + '</td>';
            tr.append(td);
            if( v.deprecated == 0 )
              btnEdit = '<a class="text-primary tooltips" title="发布" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'release\',\''+v.id+'\',\''+v.name+'\',\''+getName('unit',v.unit_id)+'\',\''+v.version+'\')"><i class="fa fa-upload"></i></a>';
            btnEdit += ' <a class="text-info tooltips" title="发布结果" data-toggle="modal" data-target="#myViewModal" onclick="viewResult(\''+v.release_id+'\',\''+v.name+'\')"><i class="fa fa-comment-o"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="废弃" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'deprecated\',\''+v.id+'\',\''+v.name+'\',\''+getName('unit',v.unit_id)+'\',\''+v.version+'\')"><i class="fa fa-trash-o"></i></a>';
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
  //console.log("action="+action);
  //console.log(JSON.stringify(postData));
  var actionDesc='';
  switch(action){
    case 'insert':
      actionDesc='添加';
      if(tab=='ver'){
        postData['files']=JSON.stringify(cache.file_ver);
        actionDesc='创建版本';
      }
      break;
    case 'update':
      actionDesc='更新';
      break;
    case 'delete':
      actionDesc='删除';
      break;
    case 'generate':
      actionDesc='生成版本';
      break;
    case 'release':
      actionDesc='发布版本指令下发';
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

//Change on ChildModal
var changeOnModal=function(){
  var tab=$('#tab').val();
  var page_other=$('#page_other').val();
  var url='/api/for_hubble/nginx_'+tab+'.php';
  var postData={};
  var form=$('#myChildModalBody').find("input,select,textarea");

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
      modalList();
      list();
      //处理模态框和表单
      $("#myChildModal :input").each(function () {
        $(this).val("");
      });
      $("#myChildModal").on("hidden.bs.modal", function() {
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
  postData={"action":"info","fIdx":idx};
  url='/api/for_hubble/nginx_'+type+'.php';
  switch(type){
    case 'group':
      title='查看分组详情 - '+idx;
      break;
    case 'unit':
      title='查看单元详情 - '+idx;
      break;
    case 'main':
      title='查看文件 - '+idx;
      break;
    case 'ver':
      title='查看版本详情 - '+idx;
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
              case 'main':
                var str='';
                $.each(data.content,function(k,v){
                  if(k=='content'){
                    str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                      '<h5 class="col-sm-12 text-primary"><strong>文件内容</strong></h5>' +
                      '<textarea class="form-control" rows="12">'+v+'</textarea>';
                  }else{
                    if(v=='') v='空';
                    if(typeof(locale[k])!='undefined') k=locale[k];
                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                  }
                });
                text+=str;
                break;
              case 'ver':
                var str='';
                $.each(data.content,function(k,v){
                  if(k=='files'){
                    str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                      '<h5 class="col-sm-12 text-primary"><strong>包含文件</strong></h5>';
                    str+='<div class="table-scrollable">' +
                      '<table class="table table-bordered table-striped table-hover">' +
                      '<thead><tr><th>文件</th><th>版本</th><th> </th><th>文件</th><th>版本</th></tr></thead>' +
                      '<tbody>';
                    var i=0;
                    $.each(JSON.parse(v),function(key,val){
                      if(i%2==0){
                        str+=(i>0)?'</tr><tr>':'<tr>';
                      }
                      str+='<td>'+val.name+'</td><td>'+val.version+'</td>';
                      if(i%2==0) str+='<td> </td>';
                      i++;
                    });
                    if(i%2>0) str+='<td> </td><td> </td>';
                    if(i>0) str+='</tr>'
                    str+='</tbody></table></div>';
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

//View on ChildModal
var viewOnModal=function(type,idx){
  NProgress.start();
  var url='',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  postData={"action":"info","fIdx":idx};
  url='/api/for_hubble/nginx_'+type+'.php';
  switch(type){
    case 'group':
      title='查看分组详情 - '+idx;
      break;
    case 'unit':
      title='查看单元详情 - '+idx;
      break;
    case 'main':
      title='查看文件 - '+idx;
      break;
    case 'ver':
      title='查看版本详情 - '+idx;
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
            switch(type){
              case 'main':
                var str='';
                $.each(data.content,function(k,v){
                  if(k=='content'){
                    str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                      '<h5 class="col-sm-12 text-primary"><strong>Content</strong></h5>' +
                      '<textarea class="form-control" rows="8">'+v+'</textarea>';
                  }else{
                    if(v=='') v='空';
                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                  }
                });
                text+=str;
                break;
              case 'ver':
                var str='';
                $.each(data.content,function(k,v){
                  if(k=='files'){
                    str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                      '<h5 class="col-sm-12 text-primary"><strong>包含文件</strong></h5>';
                    $.each(JSON.parse(v),function(key,val){
                      str+='<span class="title col-sm-5" style="font-weight: bold;">文件: '+val.name+'</span><span class="col-sm-4" style="'+tStyle+'">版本: '+val.version+'</span>';
                    });
                  }else{
                    if(v=='') v='空';
                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                  }
                });
                text+=str;
                break;
              default:
                $.each(data.content,function(k,v){
                  if(v=='') v='空';
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
            $('#myViewChildModalBody').css('height',height);
          }
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
    title='非法请求 - '+action;
    text='<div class="note note-danger">错误信息：参数错误</div>';
    $('#myViewChildModalLabel').html(title);
    $('#myViewChildModalBody').html(text);
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
      pageNotify('info','此版本暂未发布');
      text='<div class="note note-info">此版本暂未发布</div>';
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
                text+='<h5 class="col-sm-12">版本名称 : '+desc+'</h5>';
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
            $('.tooltips').each(function(){$(this).tooltip();});
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
    case 'main':
      var disabled=false;
      if($('#unit_id').val()=='') disabled=true;
      if($('#name').val()=='') disabled=true;
      if($('#content').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'ver':
      var disabled=false;
      if($('#unit_id').val()=='') disabled=true;
      if($('#name').val()=='') disabled=true;
      if($('#type').val()=='') disabled=true;
      if(cache.file_ver.length<1) disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'ver_auto':
      var disabled=false;
      if($('#unit_id').val()=='') disabled=true;
      if($('#name').val()=='') disabled=true;
      if($('#type').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'file_ver':
      var disabled=false;
      if($('#file').val()=='') disabled=true;
      if(!$('input:radio[name="version"]:checked').val()) disabled=true;
      $("#btnCommitFile").attr('disabled',disabled);
      break;
    case 'release':
      var disabled=false;
      if($('#shell_id').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
  }
}

var twiceCheck=function(action,idx,desc,pidx,ver){
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
      case 'main':
        switch(action){
          case 'deprecated':
            modalTitle='废弃文件版本';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认废弃? <span class="text text-danger">只废弃此版本!</span></h4> 文件名 : '+desc+'<br>' +
              '文件版本 : '+ver+'<br>隶属单元: '+pidx;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="id" name="id" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="deprecated">';
            break;
          default:
            modalTitle='非法请求';
            notice='<div class="note note-danger">错误信息：参数错误</div>';
            pageNotify('error','非法请求！','错误信息：参数错误');
            break;
        }
        break;
      case 'ver':
        switch(action){
          case 'deprecated':
            modalTitle='废弃此版本';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认废弃? <span class="text text-danger">只废弃此版本!</span></h4> 描述 : '+desc+'<br>' +
              '版本 : '+ver+'<br>隶属单元: '+pidx;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="id" name="id" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="deprecated">';
            break;
          case 'release':
            modalTitle='发布此版本';
            modalBody+='<div class="form-group">' +
              '<label for="shell_id" class="col-sm-2 control-label">选择发布脚本</label>' +
              '<div class="col-sm-10">' +
              '<select class="form-control" id="shell_id" name="shell_id" onchange="check(\'release\')"><option value="">请选择</option></select>' +
              '</div>' +
              '</div>';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认发布? <span class="text text-danger">发布此版本上线!</span></h4> 描述 : '+desc+'<br>' +
              '版本 : '+ver+'<br>隶属单元: '+pidx;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="id" name="id" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="release">';
            getShell();
            btnDisable=true;
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

//twiceCheck on Modal
var twiceCheckOnModal=function(action,idx,desc,pidx,ver){
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
      case 'main':
        switch(action){
          case 'deprecated':
            modalTitle='废弃文件版本';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认废弃? <span class="text text-danger">只废弃此版本!</span></h4> 文件名 : '+desc+'<br>' +
              '文件版本 : '+ver+'<br>隶属单元: '+pidx;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="id" name="id" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="deprecated">';
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
  $('#myChildModalLabel').html(modalTitle);
  $('#myChildModalBody').html(modalBody);
  if(notice!=''){
    $('#modalNotice').html(notice);
    $('#btnChildCommit').attr('disabled',true);
  }else{
    $('#btnChildCommit').attr('disabled',btnDisable);
  }
  NProgress.done();
}

var getList=function(idx){
  var actionDesc='单元';
  var fGroup = $('#fGroup').val();
  var url='/api/for_hubble/nginx_unit.php?action=list';
  if(!fGroup){
    pageNotify('warning','获取'+actionDesc+'失败！','分组ID为空！');
    return false;
  }
  var postData={'pagesize':1000,'fGroup':fGroup};
  NProgress.start();
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      NProgress.done();
      if(data.code==0){
        if(data.content.length==0){
          pageNotify('info','获取'+actionDesc+'成功！','数据为空！');
        }else{
          cache.unit = data.content;
          updateSelect('fUnit',idx);
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
    case 'fUnit':
      data=cache.unit;
      break;
    case 'file':
      data=cache.files;
      break;
    case 'shell_id':
      data=cache.shell;
      break;
  }
  tSelect.empty();
  if(name=='shell_id') tSelect.append('<option value="">请选择</option>');
  if(data){
    $.each(data,function(k,v){
      tSelect.append('<option value="' + v.id + '">' + v.name + '</option>');
    });
    if(idx){
      tSelect.val(idx).trigger('change');
    }else{
      if(name!='shell_id') tSelect.val($('#'+name+' option:nth-child(1)').val()).trigger('change');
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

var checkAll=function(o){
  $('[id=list]:checkbox').prop('checked', o.checked);
}

//modal list
var modalList = function(page) {
  cache.modal.fIdx=$('#modal_idx').val();
  cache.modal.fUnit=$('#modal_unit_id').val();
  cache.modal.fUnitName=$('#modal_unit_name').val();
  var tab=$('#tab').val();
  NProgress.start();
  var postData={};
  if(!tab||!cache.modal.fIdx||!cache.modal.fUnit){
    pageNotify('error','加载失败！','错误信息：非法请求');
    NProgress.done();
  }
  var url='/api/for_hubble/nginx_'+tab+'.php';
  postData={"action":"ver_list","fUnit":cache.modal.fUnit,"fIdx":cache.modal.fIdx};
  if (!page) page = 1;
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
        //生成页面
        //生成分页
        processModalPage(listdata, pageinfo, paginate, tab, cache.modal.fIdx);
        //生成列表
        processModalBody(listdata, tab, head, body);
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

//modal 生成分页
var processModalPage = function(data,pageinfo,paginate,tab,findex){
  var begin = data.pageSize * ( data.page - 1 ) + 1;
  var end = ( data.count > begin + data.pageSize - 1 ) ? begin + data.pageSize - 1 : data.count;
  pageinfo.html('Showing '+begin+' to '+end+' of '+data.count+' records');
  var p1=(data.page-1>0)?data.page-1:1;
  var p2=data.page+1;
  prev='<li><a href="javascript:;" onclick="modalList('+p1+',\''+tab+'\',\''+findex+'\')"><i class="fa fa-angle-left"></i></a></li>';
  paginate.append(prev);
  for (var i = 1; i <= data.pageCount; i++) {
    var li='';
    if(i==data.page){
      li='<li class="active"><a href="javascript:;" onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
    }else{
      if(i==1||i==data.pageCount){
        li='<li><a href="javascript:;" onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
      }else{
        if(i==p1){
          if(p1>2){
            console.log(i+' '+p1);
            li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a href="javascript:;" onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
          }else{
            li='<li><a href="javascript:;" onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
          }
        }else{
          if(i==p2){
            if(p2<data.pageCount-1){
              console.log(i+' '+p2);
              li='<li><a href="javascript:;" onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>'+"\n"+'<li class="disabled"><a href="#">...</a></li>';
            }else{
              li='<li><a href="javascript:;" onclick="modalList('+i+',\''+tab+'\',\''+findex+'\')">'+i+'</a></li>';
            }
          }
        }
      }
    }
    paginate.append(li);
  }
  if(p2>data.pageCount) p2=data.pageCount;
  next='<li class="next"><a href="javascript:;" title="Next" onclick="modalList('+p2+',\''+tab+'\',\''+findex+'\')"><i class="fa fa-angle-right"></i></a></li>';
  paginate.append(next);
}

//modal 生成列表
var processModalBody = function(data,tab,head,body){
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
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        td = '<td>' + (i+1) + '</td>';
        tr.append(td);
        var btnEdit='',btnDel='';
        var checked='';
        td = '<td><a class="tooltips" title="查看文件" data-toggle="modal" data-target="#myViewChildModal" onclick="viewOnModal(\''+tab+'\',\''+v.id+'\')">' + v.name + '</a></td>';
        tr.append(td);
        td = '<td>' + v.version + '</td>';
        tr.append(td);
        t = ( v.deprecated == 0 ) ? '<span class="badge bg-green">正常</span>' : '<span class="badge bg-red">已废弃</span>';
        td = '<td>' + t + '</td>';
        tr.append(td);
        td = '<td>' + getName('unit',v.unit_id) + '</td>';
        tr.append(td);
        td = '<td>' + v.opr_user + '</td>';
        tr.append(td);
        td = '<td>' + v.create_time + '</td>';
        tr.append(td);
        btnDel = '<a class="text-danger tooltips" title="废弃" data-toggle="modal" data-target="#myChildModal" onclick="twiceCheckOnModal(\'deprecated\',\''+v.id+'\',\''+v.name+'\',\''+getName('unit',v.unit_id)+'\',\''+v.version+'\')"><i class="fa fa-trash-o"></i></a>';
        if (v.deprecated != 0) btnDel='-';
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

//发布相关
var getFileList=function(){
  var fUnit=$('#fUnit').val();
  var postData={"action":"list","fUnit":fUnit,"pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/nginx_main.php?action=list',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.files = data.content;
          updateSelect('file');
        }else{
          pageNotify('info','加载文件列表成功！','数据为空！');
        }
      }else{
        pageNotify('error','加载文件列表失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载文件列表失败！','错误信息：接口不可用');
    }
  });
}

var listFileVer = function(){
  var fUnit=$('#fUnit').val();
  var fIdx=$('#file').find("option:selected").text();
  if(!fUnit||!fIdx) return false;
  var str='';
  var postData={"action":"ver_list","fUnit":fUnit,"fIdx":fIdx,"pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/nginx_main.php?action=ver_list',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          $.each(data.content,function(k,v){
            if( v.deprecated == 0 ){
              //var t=( v.deprecated == 0 )? '<span class="badge bg-green">正常</span>' : '<span class="badge bg-red">已废弃</span>';
              var t = '<span class="badge bg-green">正常</span>';
              str+='<tr><td><input type="radio" name="version" value="'+ v.version +'" onclick="check(\'file_ver\')"></td>' +
                '<td>'+ v.name +'</td>' +
                '<td>'+ v.version +'</td>' +
                '<td>'+ t +'</td>' +
                '<td>'+ v.opr_user +'</td>' +
                '<td>'+ v.create_time +'</td>' +
                '</tr>'+"\n";
            }
          });
        }else{
          pageNotify('info','加载文件版本列表成功！','数据为空！');
        }
        if(!str) str='<tr><td colspan="6">无可用版本</td></tr>';
        $('#file_ver').html(str);
      }else{
        pageNotify('error','加载文件版本列表失败！','错误信息：'+data.msg);
      }
      check('file_ver');
    },
    error: function (){
      pageNotify('error','加载文件版本列表失败！','错误信息：接口不可用');
      check('file_ver');
    }
  });
}

var setFile=function(){
  var file=$('#file').find("option:selected").text();
  var ver=$('input:radio[name="version"]:checked').val();
  var t = {'name':file,'version':ver,'is_changed':true};
  cache.file_ver.push(t);
  listFileVerChecked();
  check();
  //处理模态框和表单
  $("#myChildModal :input").each(function () {
    $(this).val("");
  });
  $("#myChildModal").on("hidden.bs.modal", function() {
    $(this).removeData("bs.modal");
  });
}

var listFileVerChecked=function(){
  var str='',i=0;
  $.each(cache.file_ver,function(k,v){
    i++;
    str+='<tr><td>'+i+'</td><td>'+ v.name +'</td><td>'+ v.version +'</td>' +
      '<td><a class="btn red btn-xs tooltips" title="删除" style="padding-left: 0px;" onclick="delFileVerChecked(\''+ k +'\')"><i class="fa fa-trash-o"></i></a></td>' +
      '</tr>';
  });
  $('#files_list').html(str);
  $('.tooltips').each(function(){$(this).tooltip();});
}

var delFileVerChecked=function(k){
  if(!$.isNumeric(k)) return false;
  if(k<cache.file_ver.length) cache.file_ver.splice(k,1);
  listFileVerChecked();
  check();
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
          updateSelect('shell_id');
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

var getDefault=function(){
  var postData={"action":"info","fIdx":1};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/nginx_main.php',
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
