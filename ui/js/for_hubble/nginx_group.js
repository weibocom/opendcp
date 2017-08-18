cache = {
  page: 1,
}

var reset = function(){
  $('#fIdx').val('');
  list(1);
}

var list = function(page,tab) {
  $('.popovers').each(function(){$(this).popover('hide');});
  NProgress.start();
  var postData={};
  if(!tab) tab=$('#tab').val();
  if(tab!='group') tab='group';
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'group':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_3').attr('class','hidden');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_nginx_'+tab+'.php"> 创建 <i class="fa fa-plus"></i></a>');
      postData={"fIdx":fIdx};
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
        td = '<td>' + v.i + '</td>';
        tr.append(td);
        td = '<td><a class="tooltips" title="查看分组详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'group\',\''+v.id+'\')">' + v.name + '</a></td>';
        tr.append(td);
        td = '<td><a class="tooltips" title="查看单元列表" href="nginx_unit.php?idx='+ v.id +'"><i class="fa fa-bars"></i></a></td>';
        tr.append(td);
        td = '<td>' + v.opr_user + '</td>';
        tr.append(td);
        td = '<td>' + v.create_time + '</td>';
        tr.append(td);
        var btnAdd='',btnEdit='',btnDel='';
        btnEdit = '<a class="text-primary tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_nginx_group.php?action=edit&idx=' + v.id + '"><i class="fa fa-edit"></i></a>';
        btnDel = '<a class="text-danger tooltips" title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')"><i class="fa fa-trash-o"></i></a>';
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
    case 'group':
      url='/api/for_hubble/nginx_group.php';
      title='查看分组详情 - '+idx;
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
            switch(type){
              case 'group':
                var locale={};
                if(typeof(locale_messages.hubble.nginx)) locale = locale_messages.hubble.nginx;
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

//commit check
var check=function(tab){
  if(!tab) tab=$('#tab').val();
  switch(tab){
    case 'group':
      var disabled=false;
      if($('#name').val()=='') disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
  }
}

//获取详情
var get = function (idx) {
  var tab=$('#tab').val();
  var url='/api/for_hubble/nginx_'+tab+'.php',postData={};
  switch (tab){
    case 'group':
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
      case 'group':
        switch(action){
          case 'del':
            modalTitle='删除AppKey';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> ID : '+idx+'<br>' +
                '名称 : '+desc;
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


