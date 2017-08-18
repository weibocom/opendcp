cache = {
  page: 1,
  alteration: {},
  params_value: {},
  params: {
    NGINX: {
      group_id: { name: '分组', type: 'select', onchange: 'getList(\'nginx_upstream\')', onload: 'getList(\'nginx_group\')' },
      name: { name: 'Upstream', type: 'select' },
      port: { name: '端口', type: 'input', is_num: true, min: 1, max: 65535, value: 8080 },
      weight: { name: '权重', type: 'input', is_num: true, min: 0, max: 1000, value: 20 },
      script_id: { name: '发布脚本', type: 'select', onload: 'getList(\'shell\')' },
    },
    SLB: {
      region: { name: '地域', type: 'select', onchange: 'getList(\'aliyun_slb\')', onload: 'getList(\'aliyun_region\')' },
      slb_id: { name: '阿里云SLB', type: 'select' },
      weight: { name: '权重', type: 'input', is_num: true, min: 0, max: 1000, value: 100 },
    },
  },
  //for Nginx
  nginx: {
    group: [],
    upstream: [],
  },
  shell: [],
  //for aliyun
  aliyun: {
    region: [],
    slb: [],
  }
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
  if(tab!='balance') tab='balance';
  var fIdx=$('#fIdx').val();
  switch(tab){
    case 'balance':
      $('#tab_1').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php"> 创建 <i class="fa fa-plus"></i></a>');
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
        td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'balance\',\''+v.id+'\')">' + v.name + '</a></td>';
        tr.append(td);
        td = '<td>' + v.type + '</td>';
        tr.append(td);
        td = '<td>' + v.opr_user + '</td>';
        tr.append(td);
        td = '<td>' + v.create_time + '</td>';
        tr.append(td);
        var btnAdd='',btnEdit='',btnDel='';
        btnEdit = '<a class="text-primary tooltips" title="修改" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.id + '"><i class="fa fa-edit"></i></a>';
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
  var url='/api/for_hubble/'+tab+'.php';
  var postData={},content={};
  var form=$('#myModalBody').find("input,select,textarea");

  //处理表单内容--不需要修改
  $.each(form,function(i){
    if(this.name.indexOf('param_')!=0){
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
      switch(postData.type){
        case 'NGINX':
          $.each(cache.params.NGINX,function(k,v){
            content[k]=$('#param_'+k).val();
          });
          postData.content=JSON.stringify(content);
          break;
        case 'SLB':
          $.each(cache.params.SLB,function(k,v){
            content[k]=$('#param_'+k).val();
          });
          postData.content=JSON.stringify(content);
          break;
        default:
          postData.content=$('#content').val();
          break;
      }
      break;
    case 'update':
      actionDesc='修改';
      switch(postData.type){
        case 'NGINX':
          $.each(cache.params.NGINX,function(k,v){
            content[k]=$('#param_'+k).val();
          });
          postData.content=JSON.stringify(content);
          break;
        case 'SLB':
          $.each(cache.params.SLB,function(k,v){
            content[k]=$('#param_'+k).val();
          });
          postData.content=JSON.stringify(content);
          break;
        default:
          postData.content=$('#content').val();
          break;
      }
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
  url='/api/for_hubble/'+type+'.php';
  postData={"action":"info","fIdx":idx};
  switch(type){
    case 'balance':
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
            switch(type){
              case 'balance':
                var str='',locale={};
                if(typeof(locale_messages.hubble.balance)) locale = locale_messages.hubble.balance;
                $.each(data.content,function(k,v){
                  if(k=='content'){
                    str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                      '<h5 class="col-sm-12 text-primary"><strong>Content</strong></h5>' +
                      '<textarea class="form-control" rows="8">'+v+'</textarea>';
                  }else{
                    if(v=='') v='空';
                    if(typeof(locale[k])!='undefined') k=locale[k];
                    text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
                  }
                });
                text+=str;
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
    case 'balance':
      var disabled=false;
      var params;
      if($('#name').val()=='') disabled=true;
      if($('#type').val()=='') disabled=true;
      switch($('#type').val()){
        case 'NGINX':
          params=cache.params.NGINX;
          $.each(params,function(k,v){
            if($('#param_'+k).val()=='') disabled=true;
          });
          break;
        case 'SLB':
          params=cache.params.SLB;
          $.each(params,function(k,v){
            if($('#param_'+k).val()=='') disabled=true;
          });
          break;
        default:
          if($('#content').val()=='') disabled=true;
          break;
      }
      $("#btnCommit").attr('disabled',disabled);
      break;
  }
}

//获取详情
var get = function (idx) {
  var tab=$('#tab').val(),postData={};
  url='/api/for_hubble/'+tab+'.php';
  postData={"action":"info","fIdx":idx};
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
            cache.params_value=JSON.parse(data.content.content);
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
            listParams(data.content.type);
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
      case 'balance':
        switch(action){
          case 'del':
            modalTitle='删除类型';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> ID : '+idx+'<br>' +
              '名称 : '+desc+'<br>';
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

//获取服务发现类型参数
var getType=function(){
  var actionDesc='服务发现类型参数';
  var fIdx=$('#type').val();
  if(!fIdx) return false;
  var postData={action:"info",fIdx:fIdx,pagesize:1000};
  $.ajax({
    type: "POST",
    url: '/api/for_hubble/alteration.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      cache.alteration = {};
      if(data.code==0){
        if(data.content){
          cache.alteration = data.content;
        }else{
          pageNotify('info','加载'+actionDesc+'成功！','数据为空！');
        }
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：'+data.msg);
      }
      listParams(fIdx);
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
      cache.alteration = {};
      listParams(fIdx);
    }
  });
}

//列出参数
var listParams=function(fIdx){
  var str='',data=cache.alteration,params_value=cache.params_value,params;
  var name,type,is_num,min,max,value,onchange,onload=[];
  switch (fIdx){
    case 'NGINX':
      params=cache.params.NGINX;
      $.each(params,function(k,v){
        name = (typeof v.name != 'undefined') ? v.name : k;
        type = (typeof v.type != 'undefined') ? v.type : 'input';
        is_num = (typeof v.is_num != 'undefined') ? v.is_num : '';
        min = (typeof v.min != 'undefined') ? v.min : 'A';
        max = (typeof v.max != 'undefined') ? v.max : 'A';
        if(typeof params_value[k] != 'undefined'){
          value = params_value[k];
        }else{
          value = (typeof v.value != 'undefined') ? v.value : '';
        }
        onchange = (typeof v.onchange != 'undefined') ? 'onchange="' + v.onchange +'"' : 'onchange="check()"';
        if(typeof v.onload != 'undefined') onload.push(v.onload);
        switch (type){
          case 'select':
            str+='<div class="form-group">'+"\n"+
              '<label for="param_'+k+'" class="col-sm-2 control-label">'+name+'</label>'+"\n" +
              '<div class="col-sm-10">'+"\n" +
              '<select class="form-control" id="param_'+k+'" name="param_'+k+'" '+onchange+'>'+"\n" +
              '<option value="">请选择</option>'+"\n" +
              '</select>'+"\n" +
              '</div>'+"\n" +
              '</div>'+"\n";
            break;
          default:
            str+='<div class="form-group">'+"\n"+
              '<label for="param_'+k+'" class="col-sm-2 control-label">'+name+'</label>'+"\n" +
              '<div class="col-sm-10">'+"\n" +
              '<input type="'+((is_num)?'number':'text')+'"'+((is_num&&min!='A')?' min="'+min+'"':'')+' '+((is_num&&max!='A')?' max="'+max+'"':'')+' class="form-control" id="param_'+k+'" name="param_'+k+'" value="'+value+'" '+onchange+'>'+"\n" +
              '</div>'+"\n" +
              '</div>'+"\n";
            break;
        }
      });
      break;
    case 'SLB':
      params=cache.params.SLB;
      $.each(params,function(k,v){
        name = (typeof v.name != 'undefined') ? v.name : k;
        type = (typeof v.type != 'undefined') ? v.type : 'input';
        is_num = (typeof v.is_num != 'undefined') ? v.is_num : '';
        min = (typeof v.min != 'undefined') ? v.min : 'A';
        max = (typeof v.max != 'undefined') ? v.max : 'A';
        if(typeof params_value[k] != 'undefined'){
          value = params_value[k];
        }else{
          value = (typeof v.value != 'undefined') ? v.value : '';
        }
        onchange = (typeof v.onchange != 'undefined') ? 'onchange="' + v.onchange +'"' : 'onchange="check()"';
        if(typeof v.onload != 'undefined') onload.push(v.onload);
        switch (type){
          case 'select':
            str+='<div class="form-group">'+"\n"+
              '<label for="param_'+k+'" class="col-sm-2 control-label">'+name+'</label>'+"\n" +
              '<div class="col-sm-10">'+"\n" +
              '<select class="form-control" id="param_'+k+'" name="param_'+k+'" '+onchange+'>'+"\n" +
              '<option value="">请选择</option>'+"\n" +
              '</select>'+"\n" +
              '</div>'+"\n" +
              '</div>'+"\n";
            break;
          default:
            str+='<div class="form-group">'+"\n"+
              '<label for="param_'+k+'" class="col-sm-2 control-label">'+name+'</label>'+"\n" +
              '<div class="col-sm-10">'+"\n" +
              '<input type="'+((is_num)?'number':'text')+'"'+((is_num&&min!='A')?' min="'+min+'"':'')+' '+((is_num&&max!='A')?' max="'+max+'"':'')+' class="form-control" id="param_'+k+'" name="param_'+k+'" value="'+value+'" '+onchange+'>'+"\n" +
              '</div>'+"\n" +
              '</div>'+"\n";
            break;
        }
      });
      break;
    default:
      str+='<textarea rows="8" class="form-control" id="content" name="content" onkeyup="check()" placeholder="关联" style="background:transparent;border-color: #fff;">'+JSON.stringify(data, null, 4)+'</textarea>';
      break;
  }
  if(str){
    $('#params').empty();
  }else{
    str='<p> </p><p> </p>'
  }
  $('#params').append(str);
  var select=$('#params').find("select");
  $.each(select,function(i){
    $('#'+this.id).select2({width:'100%'});
  });
  $('.tooltips').each(function(){$(this).tooltip();});

  $.each(onload,function(k,v){
    eval(v);
  });
}

//获取列表
var getList=function(type,idx){
  if(!type) return false;
  var actionDesc='服务发现',url='/api/for_hubble/'+type+'.php',select=type;
  var postData={action:"list",pagesize:1000};
  switch (type){
    case 'nginx_group':
      actionDesc='分组';
      break;
    case 'nginx_upstream':
      actionDesc='Upstream';
      postData.fGroup=$('#param_group_id').val();
      break;
    case 'shell':
      actionDesc='发布脚本';
      break;
    case 'aliyun_region':
      actionDesc='阿里云地域';
      url='/api/for_cloud/region.php';
      break;
    case 'aliyun_slb':
      actionDesc='阿里云SLB';
      postData.fRegion=$('#param_region').val();
      url='/api/for_cloud/slb.php';
      break;
  }
  $.ajax({
    type: "POST",
    url: url,
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          switch (type){
            case 'nginx_group':
              cache.nginx.group = data.content;
              break;
            case 'nginx_upstream':
              cache.nginx.upstream = data.content;
              break;
            case 'shell':
              cache.shell = data.content;
              break;
            case 'aliyun_region':
              cache.aliyun.region = data.content;
              break;
            case 'aliyun_slb':
              cache.aliyun.slb = data.content;
              break;
          }
        }else{
          switch (type){
            case 'nginx_group':
              cache.nginx.group=[];
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/nginx_group.php">创建'+actionDesc+'</a>]！',false);
              break;
            case 'nginx_upstream':
              cache.nginx.upstream=[];
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/nginx_unit.php">创建'+actionDesc+'</a>]！',false);
              break;
            case 'shell':
              cache.shell=[];
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/shell.php">创建'+actionDesc+'</a>]！',false);
              break;
            case 'aliyun_region':
              cache.aliyun.region=[];
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="https://www.aliyun.com/">创建'+actionDesc+'</a>]！',false);
              break;
            case 'aliyun_slb':
              cache.aliyun.slb=[];
              pageNotify('warning','获取'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_hubble/slb.php">创建'+actionDesc+'</a>]！',false);
              break;
          }
        }
        updateSelect(select);
      }else{
        pageNotify('error','加载'+actionDesc+'失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载'+actionDesc+'失败！','错误信息：接口不可用');
    }
  });
}

var updateSelect=function(name){
  var tSelect=$('#'+name),data='',params_value='';
  switch(name){
    case 'nginx_group':
      data=cache.nginx.group;
      tSelect=$('#param_group_id');
      if(typeof cache.params_value.group_id != 'undefined') params_value=cache.params_value.group_id;
      break;
    case 'nginx_upstream':
      data=cache.nginx.upstream;
      tSelect=$('#param_name');
      if(typeof cache.params_value.name != 'undefined') params_value=cache.params_value.name;
      break;
    case 'shell':
      data=cache.shell;
      tSelect=$('#param_script_id');
      if(typeof cache.params_value.script_id != 'undefined') params_value=cache.params_value.script_id;
      break;
    case 'aliyun_region':
      data=cache.aliyun.region;
      tSelect=$('#param_region');
      if(typeof cache.params_value.region != 'undefined') params_value=cache.params_value.region;
      break;
    case 'aliyun_slb':
      data=cache.aliyun.slb;
      tSelect=$('#param_slb_id');
      if(typeof cache.params_value.slb_id != 'undefined') params_value=cache.params_value.slb_id;
      break;
  }
  tSelect.empty();
  tSelect.append('<option value="">请选择</option>');
  if(data){
    switch(name){
      case 'nginx_upstream':
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.name + '">' + v.name + '</option>');
        });
        break;
      case 'aliyun_region':
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.RegionName + '">' + v.RegionName + '</option>');
        });
        break;
      case 'aliyun_slb':
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.LoadBalancerId + '">' + v.Address + ' - ' + v.LoadBalancerName + '</option>');
        });
        break;
      default:
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.id + '">' + v.name + '</option>');
        });
        break;
    }
  }
  if(params_value){
    if(tSelect.find("option[value='"+params_value+"']").length==0) tSelect.append('<option value="' + params_value + '">' + params_value + ' - 此记录不存在</option>');;
  }
  tSelect.select2({width:'100%'});
  if(params_value) tSelect.val(params_value).trigger('change');
}
