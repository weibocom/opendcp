cache = {
    page: 1,
    projects: [],
    tags:{},
    autocomplete:[],
    state: {},
    currentProjectLogId: 0,  //保存当前项目的ID
    currentProjectName: ""  //保存当前项目的名称
}
//自动刷新
var autoRefresh = null;

var isJson = function(str) {
  try {
    JSON.parse(str);
  } catch (e) {
    return false;
  }
  return true;
}

var getDate = function(t){
  if(!t) t='';
  var d=new Date(t);
  var month=d.getMonth()+1;
  var day=d.getDate();
  var hour=d.getHours();
  var min=d.getMinutes();
  var sec=d.getSeconds();
  if(month<10)  month='0'+month;
  if(day<10)    day='0'+day;
  if(hour<10)   hour='0'+hour;
  if(min<10)    min='0'+min;
  if(sec<10)    sec='0'+sec;
  return d.getFullYear()+'.'+month+'.'+day+' '+hour+':'+min+':'+sec;
};

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
  if(tab!='package'){
    tab='package';
  }
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'package':
      $('#tab_1').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'edit\',\'\')"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fIdx":fIdx};
      break;
  }
  var url='/api/for_repos/'+tab+'.php';
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
  var begin = ( data.count > 0 ) ? data.pageSize * ( data.page - 1 ) + 1 : 0;
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
      for (var i = 0; i < data.content.length; i++) {
        cache.state[i] = {};
        var v = data.content[i];
        var tr = $('<tr></tr>');
        td = '<td>' + v.i + '</td>';
        tr.append(td);
        var btnAdd='',btnEdit='',btnDel='';
        switch(tab){
          case 'package':
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'package\',\''+ v.name+'\')">' + v.name + '</a></td>';
            tr.append(td);
            td = '<td>' + v.creator + '</td>';
            tr.append(td);
            td = '<td>' + getDate(v.lastModifyTime) + '</td>';
            tr.append(td);
            td = '<td id="state_'+i+'">-</td>';
            tr.append(td);
            getState(i, v.name);
            btnAdd = '<a class="text-success tooltips" title="克隆" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'clone\',\''+v.name+'\')"><i class="fa fa-copy"></i></a>' +
                ' <a class="text-primary tooltips" title="构建镜像" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'build\',\''+v.name+'\',\''+v.Cluster+'\')"><i class="fa fa-building-o"></i></a>';
            td = '<td>' + btnAdd + '</td>';
            tr.append(td);
            btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'edit\',\''+v.name+'\')"><i class="fa fa-edit"></i></a>';
            btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.name+'\',\''+v.creator+'\')"><i class="fa fa-trash-o"></i></a>';
            break;
        }
        var tb=(!btnEdit && !btnDel) ? '-' : btnEdit + ' ' + btnDel;
        td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + tb + '</div></td>';
        if(tb!='-') tr.append(td);

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
  var url='/api/for_repos/'+tab+'.php';
  var postData={};
  var form=$('#myModalBody').find("input,select,textarea");

  var action=$("#page_action").val();
  if(action=='update'){
    postData='';
    //处理表单内容--不需要修改
    $.each(form,function(i){
      switch(this.type){
        case 'radio':
          if(this.name) postData+=(postData)?'&'+this.name+'='+encodeURIComponent($('input[name="'+this.name+'"]:checked').val()):this.name+'='+encodeURIComponent($('input[name="'+this.name+'"]:checked').val());
          break;
        case 'checkbox':
          break;
        default:
          if(this.name!='page_action'){
            if(this.name) postData+=(postData)?'&'+this.name+'='+encodeURIComponent(this.value):this.name+'='+encodeURIComponent(this.value);
          }
          break;
      }
    });
  }else{
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
  }
  delete postData['page_action'];
  var actionDesc='';
  switch(action){
    case 'insert':
      actionDesc='添加';
      break;
    case 'clone':
      actionDesc='克隆';
      break;
    case 'build':
      actionDesc='构建镜像';
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
      list();
    }
  });
}

//commit check
var check=function(tab){
  if(!tab) tab=$('#tab').val();
  switch(tab){
    case 'new':
      var disabled=false;
      if($('#projectName').val()=='') disabled=true;
      if($('#staticDockerfile').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'clone':
      var disabled=false;
      if($('#srcProjectName').val()=='') disabled=true;
      if($('#dstProjectName').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'build':
      var disabled=false;
      if($('#projectName').val()=='') disabled=true;
      if($('#tag').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'update':
      var disabled=false;
      $('#myModalBody').find("input,select,textarea").each(function(i){
        if(!disabled){
          if($(this).val()=='') disabled=true;
        }
      });
      $("#btnCommit").attr('disabled',disabled);
      break;
  }
}

//view
var view=function(type,idx,idx2){
  NProgress.start();
  var url='',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  url='/api/for_repos/'+type+'.php';
  switch(type){
    case 'package':
      title='查看项目详情 - '+idx;
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
            //pageNotify('success','加载成功！');
            switch(type){
              case 'package':
                var tStr='',locale={};
                if(typeof(locale_messages.package)) locale = locale_messages.package;
                $.each(data.content,function(k,v){
                  if(v=='') v='空';
                  if(typeof(locale[k])!='undefined') k=locale[k];
                  text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                });
                text+=tStr;
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
      case 'package':
        switch(action){
          case 'clone':
            modalTitle='克隆项目';
            modalBody+='<div class="form-group">' +
                '<label for="srcProjectName" class="col-sm-2 control-label">源项目名称</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="srcProjectName" name="srcProjectName" onkeyup="check(\'clone\')" value="'+idx+'" readonly>' +
                '</div>' +
                '</div>';
            modalBody+='<div class="form-group">' +
                '<label for="dstProjectName" class="col-sm-2 control-label">新项目名称</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="dstProjectName" name="dstProjectName" onkeyup="check(\'clone\')" placeholder="新项目名称,eg:测试">' +
                '</div>' +
                '</div>';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="clone">';
            btnDisable=true;
            break;
          case 'build':
            modalTitle='构建镜像';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认为项目 <span class="text text-danger">'+idx+'</span> 构建镜像? </h4>';
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<div class="form-group">' +
                '<label for="tag" class="col-sm-2 control-label">标签</label>' +
                '<div class="col-sm-10">' +
                '<input type="text" class="form-control" id="tag" name="tag" onkeyup="check(\'build\')" placeholder="标签(不支持中文和特殊符号) ,eg:test">' +
                '<div id="tag-container" style="position: relative; float: left; width: 400px; margin: 10px;" z-index="99999"></div>' +
                '</div>' +
                '</div>';
            modalBody+='<input type="hidden" id="projectName" name="projectName" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="build">';
            break;
          case 'edit':
            modalTitle=(idx)?'项目配置 - '+idx:'创建项目';
            modalBody=getConfigDep(idx);
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="update">';
            btnDisable=true;
            break;
          case 'del':
            modalTitle='删除项目';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认删除? <span class="text text-danger">警告! 操作不可回退!</span></h4> 项目名称 : '+idx+'<br>创建人 : '+desc;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="projectName" name="projectName" value="'+idx+'">';
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
  var modal='myModal',btn='btnCommit';
  if(notice!=''){
    modalBody=notice;
    $('#'+btn).attr('disabled',true);
  }else{
    $('#'+btn).attr('disabled',btnDisable);
  }
  $('#'+modal+'Label').html(modalTitle);
  $('#'+modal+'Body').html(modalBody);
  if(action=='build'){
    getTag(desc,idx);
    $('#'+btn).attr('disabled',true);
    $('#'+modal+'Body').css('overflow','');
    var p = $("#projectName").val(),count=0;
    if (typeof cache.tags[p] != "undefined"){
      $.each(cache.tags[p],function(k,v){
        count++;
        cache.autocomplete.push({data: v.name,value: v.name});
      });
    }
    $("#tag").autocomplete({
      lookup: cache.autocomplete,
      appendTo: '#tag-container',
      showNoSuggestionNotice: true,
      autoSelectFirst: true
    });
  }else{
    $('#'+modal+'Body').css('overflow','auto');
  }
  NProgress.done();
}

//获取配置
var getConfigDep=function(idx){
  var actionDesc='项目配置';
  var postData={'action':'dep','fIdx':idx};
  var ret=false;
  $.ajax({
    async: false,
    type: "POST",
    url: '/api/for_repos/package.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content){
          ret=data.content;
        }else{
          pageNotify('info','加载'+actionDesc+'成功！','数据为空！');
          ret='接口返回数据为空';
        }
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：依赖接口不可用');
        ret=data.content;
      }
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
      ret='错误信息：接口不可用';
    }
  });
  return ret;
}

//获取TAG
var getTag=function(cluster,o){
  if(!o) return false;
  var postData={"fRepos":cluster+'/'+o};
  var url='/api/for_repos/tags.php?action=list&page=1000';
  $.ajax({
    async: false,
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content){
          cache.tags[o] = data.content;
        }
      }
    }
  });
}

//获取构建状态
var getState=function(i,o){
  if(!o) return false;
  var postData={"fIdx":o},str='';
  var url='/api/for_repos/package.php?action=state';
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      cache.state[i] = data;
      if(data.code==0){
        if(typeof data.content != 'undefined'){
          switch(data.content.state){
            case 0:
              str = '<span class="badge bg-blue">构建中</span>';
              break;
            case 1:
              str = '<span class="badge bg-green">成功</span>';
              break;
            case 2:
              str = '<span class="badge bg-red">失败</span>';
              break;
            default:
              str = '<span class="badge">未知状态</span>';
              break;
          }
          str += '<a class="pull-right" data-toggle="modal" data-target="#myViewModal" title="查看日志" onclick="showLog(' + i + ',\''+ o +'\')"><i class="fa fa-history"></i></a>';
          $('#state_'+i).html(str);
        }
      }
    }
  });
}

//关闭自动刷新
var closeRefresh = function(){
    if(autoRefresh){
        clearInterval(autoRefresh);
    }
}
//显示日志
var showLog = function (i, o){
    cache.currentProjectLogId = i;
    cache.currentProjectName = o;
    refreshLog();
    autoRefresh = setInterval(refreshLog,3000);//3秒刷新一次
}
//刷新日志
var refreshLog = function (){
    var index = cache.currentProjectLogId;
    var name =  cache.currentProjectName;
    NProgress.start();
    var title='查看构建日志',text='';
    if(typeof index != 'undefined'){
        if(typeof cache.state[index] != 'undefined'){
            if(typeof cache.state[index].content != 'undefined'){
                var result = cache.state[index].content.logs.replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/\n/g,'<br>');
                text ='<span class="col-sm-12" style="background-color:#000;color:#ccc;line-height: 150%">'+ result +'</span>';
            }else{
                text='<div class="note note-danger">'+JSON.stringify(cache.state[index])+'</div>';
            }
        } else {
            text='<div class="note note-danger">加载失败：未找到对应日志</div>';
        }
    }else{
        text='<div class="note note-danger">加载失败：参数错误</div>';
    }
    text += '<span class="pull-right text-danger">Updated:'+getCurrentDate(new Date(),'time')+'</span>';
    $('#myViewModalLabel').html(title);
    $('#myViewModalBody').html(text);
    getState(index,name);
    NProgress.done();
}
//获取当前时间
var getCurrentDate = function(t,type){
    if(!t) t='';
    var d= new Date(t);
    var M= (d.getMonth()+1);
    var D= d.getDate();
    var h= d.getHours();
    var i= d.getMinutes();
    var s= d.getSeconds();
    var ret='';
    switch (type){
        case 'time':
            ret=((h<10)?'0'+h:h) +':'+ ((i<10)?'0'+i:i) +':'+ ((s<10)?'0'+s:s);
            break;
        default:
            ret=d.getFullYear()+'.'+ ((M<10)?'0'+M:M) +'.'+ ((D<10)?'0'+D:D) +' '+ ((h<10)?'0'+h:h) +':'+ ((i<10)?'0'+i:i) +':'+ ((s<10)?'0'+s:s);
            break;
    }
    return ret;
}