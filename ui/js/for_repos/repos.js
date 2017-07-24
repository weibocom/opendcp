cache = {
  page: 1,
  projects: [],
  repositories: [],
  copy: {
    ip: [],
  },
  tags:{},
  autocomplete:[],
}

var getDate = function(t){
  if(!t) t='';
  d=new Date(t);
  return d.getFullYear()+'.'+ (d.getMonth()+1)+'.'+ d.getDate()+' '+ d.getHours()+':'+ d.getMinutes()+':'+ d.getSeconds();
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
  if(tab!='projects'&&tab!='repositories'&&tab!='tags'){
    tab='projects';
  }
  var fProject=$('#fProject').val();
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'projects':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=add"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fIdx":fIdx};
      $('#fProject').parent().parent().attr('hidden',true);
      $('#fRepos').parent().parent().attr('hidden',true);
      $('#page_images').html('');
      $('#page_table').attr('hidden',false);
      break;
    case 'repositories':
      $('#tab_1').attr('class','hidden');
      $('#tab_2').attr('class','active');
      //$('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'edit\',\'\')"> 创建 <i class="fa fa-plus"></i></a>');
      if(idx){
        fProject=idx;
        $('#fProject').val(fProject);
      }
      postData={"fProject":fProject,"fIdx":fIdx};
      $('#fProject').parent().parent().attr('hidden',false);
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
        if(tab!='projects'){
          body.parent().attr('hidden',true);
          body.html('');
          var body = $("#page_images");
        }
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
        if(tab!='projects'){
          processBodyImages(listdata, body);
        }else{
          processBody(listdata, head, body);
        }
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
          case 'projects':
            td = '<td>' + v.name + '</td>';
            tr.append(td);
            var t=(v.owner_name)?v.owner_name:v.owner_id;
            td = '<td><a class="tooltips" title="owner_id: '+ v.owner_id +'">' + t + '</a></td>';
            tr.append(td);
            td = '<td><a class="tooltips" title="查看镜像列表" onclick="getList(\'repositories\',\'projects\',\''+ v.project_id+'\')">' + v.repo_count + '</a></td>';
            tr.append(td);
            td = '<td>' + getDate(v.creation_time) + '</td>';
            tr.append(td);
            //btnEdit = '<a class="text-success tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.Id + '"><i class="fa fa-edit"></i></a>';
            //btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.Id+'\',\''+v.Name+'\')"><i class="fa fa-trash-o"></i></a>';
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

//生成列表Image
var processBodyImages = function(data,body){
  var str='',p=$('#fProject').find('option:selected').text();
  if(data.content){
    if(data.content.length>0){
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<div class="panel"></div>');
        str='<a class="panel-heading collapsed" role="tab" id="heading_'+i+'" data-toggle="collapse" data-parent="#accordion" href="#collapse'+i+'" aria-expanded="false" aria-controls="collapse'+i+'">' +
          '<h4 class="panel-title"><i class="fa fa-book"></i> '+ v.name.replace(p+'/','') +' <i class="badge bg-green" id="tagsNum_'+i+'">0</i></h4></a>' +
          /*
          '<div class="panel-heading" role="tab">' +
          '<h4 class="panel-title">' +
          '<a class="collapsed" role="button" id="heading_'+i+'" data-toggle="collapse" data-parent="#accordion" href="#collapse'+i+'" aria-expanded="false" aria-controls="collapse'+i+'">' +
          '<i class="fa fa-book"></i> '+ v.name.replace(p+'/','') +' <i class="badge bg-green" id="tagsNum_'+i+'">0</i>' +
          '</a>' +
          '<div class="btn-group btn-group-xs btn-group-solid pull-right">' +
          '<a class="text-success tooltips" title="克隆" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'clone\',\''+v.name+'\')"><i class="fa fa-copy"></i></a> ' +
          '<a class="text-primary tooltips" style="margin-left:10px;" title="构建镜像" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'build\',\''+v.name+'\')"><i class="fa fa-building-o"></i></a> ' +
          '<a class="text-success tooltips" style="margin-left:10px;" title="修改" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'edit\',\''+v.name+'\')"><i class="fa fa-edit"></i></a> ' +
          '<a class="text-danger tooltips" style="margin-left:10px;" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.name+'\',\''+v.creator+'\')"><i class="fa fa-trash-o"></i></a>' +
          '</div>' +
          '</h4>' +
          '</div>' +
          */
          '<div id="collapse'+i+'" class="panel-collapse collapse" role="tabpanel" aria-labelledby="heading'+i+'" aria-expanded="false">' +
          '<div class="panel-body" style="padding-top: 2px;padding-bottom: 2px;">' +
          '<table class="table table-hover" style="margin-bottom: 0px;"><thead><tr><td>序号</td><td>标签</td><td>操作</td></tr></thead><tbody id="tags_'+i+'"></tbody></table>' +
          '<div class="row">' +
          '<div class="col-md-5 col-sm-5"><div class="dataTables_info" id="tags_'+i+'-pageinfo" role="status" aria-live="polite">Showing 1 to 0 of 0 entries</div></div>' +
          '<div class="col-md-7 col-sm-7">' +
          '<div class="dataTables_paginate paging_bootstrap_full_number"><ul class="pagination" style="visibility: visible;margin-top: 0px;margin-bottom: 0px;" id="tags_'+i+'-paginate"></ul></div>' +
          '</div></div>' +
          '</div></div>';
        tr.append(str);

        body.append(tr);
        listTags(1, i, v.name);
      }
    }else{
      pageNotify('info','Warning','数据为空！');
    }
  }else{
    pageNotify('warning','error','接口异常！');
  }
}

var listTags = function(page,id,fRepos) {
  $('.popovers').each(function(){$(this).popover('hide');});
  var postData={"fRepos":fRepos};
  var url='/api/for_repos/tags.php';
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
        var pageinfo = $('#tags_'+id+'-pageinfo');//分页信息
        var paginate = $('#tags_'+id+'-paginate');//分页代码
        var body = $("#tags_"+id);//数据列表
        //清除当前页面数据
        pageinfo.html("");
        paginate.html("");
        body.html("");
        //生成页面
        //生成分页
        processPage(listdata, pageinfo, paginate);
        //生成列表
        processBodyTags(listdata, body, id, fRepos);
        $('.popovers').each(function(){$(this).popover();});
        $('.tooltips').each(function(){$(this).tooltip();});
      }else{
        pageNotify('error','加载失败！','错误信息：'+listdata.msg);
      }
    },
    error: function (){
      pageNotify('error','加载失败！','错误信息：接口不可用');
    }
  });
}

//生成列表Tags
var processBodyTags = function(data,body,id, fRepos){
  var td="";
  if(data.content){
    cache.tags[fRepos]=data.content;
    if(data.content.length>0){
      $('#tagsNum_'+id).html(data.content.length);
      var tab=$('#tab').val();
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        if(tab!='node'){
          td = '<td>' + v.i + '</td>';
          tr.append(td);
        }
        var btnDel='';
        td = '<td><a class="tooltips" title="查看镜像详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'tags\',\''+fRepos+'\',\''+v.name+'\')">' + v.name + '</a></td>';
        tr.append(td);
        btnDel = '<a class="text-success tooltips" title="" data-original-title="查看镜像" data-toggle="modal" data-target="#myViewModal" onclick="showDetailImage(\''+data.imageAddress+'\',\''+fRepos+'\',\''+v.name+'\')"><i class="fa fa-history"></i></a>&nbsp;&nbsp;';
        btnDel += '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'delTag\',\''+fRepos+'\',\''+v.name+'\')"><i class="fa fa-trash-o"></i></a>';
        td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnDel + '</div></td>';
        tr.append(td);
        body.append(tr);
      }
    //}else{
      //pageNotify('info','Warning','数据为空！');
    }
  }else{
    pageNotify('warning','error','接口异常！');
  }
}

//commit check
var check=function(tab){
  if(!tab) tab=$('#tab').val();
  switch(tab){
    case 'projects':
      var disabled=false;
      if($('#project_name').val()=='') disabled=true;
      if($('#public').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
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
  var url='/api/for_repos/'+tab+'.php';
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
var view=function(type,idx,idx2){
  NProgress.start();
  var url='',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  url='/api/for_repos/'+type+'.php';
  switch(type){
    case 'tags':
      title='查看镜像详情 - '+idx+' / '+idx2;
      postData={"action":"info","fRepos":idx,"fIdx":idx2};
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
              case 'tags':
                var tStr='',info={};
                /*
                $.each(data.content,function(k,v){
                  if(v=='') v='空';
                  text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                });
                */
                info=data.content;info.config=JSON.parse(info.config);
                tStr=JSON.stringify(info,null,"\t");
                text+='<textarea class="form-control" rows="30">'+tStr+'</textarea>';
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
      case 'repositories':
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
              '<input type="text" class="form-control" id="tag" name="tag" onkeyup="check(\'build\')" placeholder="标签,eg:测试">' +
              '<div id="tag-container" style="position: relative; float: left; width: 400px; margin: 10px;" z-index="99999"></div>' +
              '</div>' +
              '</div>';
            modalBody+='<input type="hidden" id="projectName" name="projectName" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="build">';
            break;
          case 'edit':
            modalTitle=(idx)?'项目配置 - '+idx:'创建项目';
            modalBody=getConfigDep(idx);
            //modalBody+='<input type="hidden" id="projectName" name="projectName" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="update">';
            break;
          case 'del':
            modalTitle='删除项目';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认删除? <span class="text text-danger">警告! 操作不可回退!</span></h4> 项目名称 : '+idx+'<br>';
            if(desc!='undefined') modalBody+='创建人 : '+desc;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody+='<input type="hidden" id="projectName" name="projectName" value="'+idx+'">';
            modalBody+='<input type="hidden" id="page_action" name="page_action" value="delete">';
            break;
          case 'delTag':
            modalTitle='删除镜像TAG';
            modalBody+='<div class="form-group col-sm-12">';
            modalBody+='<div class="note note-danger">';
            modalBody+='<h4>确认删除? <span class="text text-danger">警告! 操作不可回退!</span></h4> 镜像 : '+idx+'<br>';
            if(desc!='undefined') modalBody+='TAG : '+desc;
            modalBody+='</div>';
            modalBody+='</div>';
            modalBody=modalBody+'<input type="hidden" id="repo_name" name="repo_name" value="'+idx+'">';
            modalBody=modalBody+'<input type="hidden" id="tag" name="tag" value="'+desc+'">';
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
  if(action=='build'){
    $('#btnCommit').attr('disabled',true);
    $('#myModalBody').css('overflow','');
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
    $('#myModalBody').css('overflow','auto');
  }
  NProgress.done();
}

var getList=function(tab,type,idx){
  if(!type) type='projects';
  var url='/api/for_repos/'+type+'.php?action=list';
  var postData={'pagesize':1000};
  if(!tab) tab='repositories';
  $('#tab').val(tab);
  var actionDesc='';
  switch (type){
    case 'projects':
      actionDesc='项目';
      break;
    default:
      actionDesc='非法请求';
      return false;
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
          case 'projects':
            cache.projects = data.content;
            updateSelect('fProject',idx);
            break;
        }
        if(data.content.length==0){
          pageNotify('info','获取'+actionDesc+'成功！','数据为空！');
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
    case 'fProject':
      data=cache.projects;
      break;
  }
  tSelect.empty();
  if(data){
    switch(name){
      case 'fProject':
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.project_id + '">' + v.name + '</option>');
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
//查看镜像全称
var showDetailImage = function(address, idx,desc){
    var myaddresslist = address.split("/");
    var myaddress =  address.split("/")[myaddresslist.length - 1];
    var myidx = idx;
    text = "<div class='note note-danger'>" + myaddress + "/"+ myidx +":"+desc+"</div>";
    $('#myViewModalLabel').html("镜像全称");
    $('#myViewModalBody').html(text);

}