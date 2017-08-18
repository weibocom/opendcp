cache={
  page:1,
  check_action: {}, //Action详情
  step_action: [], //Action列表
  params: {},
  action_default: {"action": {"args": "docker run -d --name={{name}} --net={{net}} {{tag}} ","module": "shell"}},
  actimpl: {},
  check_chinese: true, //是否校验中文字符
}

var hasChinese=function(str){
  if(/[\u4e00-\u9fa5]/g.test(str)){
    return str.replace(/[\u4e00-\u9fa5]/g,'');
  }else{
    return str;
  }
}

var isJson=function(str) {
  try {
    JSON.parse(str);
  } catch (e) {
    return false;
  }
  return true;
}

var list = function(page,tab) {
  $('.popovers').each(function(){$(this).popover('hide');});
  NProgress.start();
  var postData={};
  if(!tab){
    tab=$('#tab').val();
  }
  if(tab!='action'&&tab!='remote_step'){
    tab='action';
  }
  var fIdx=$('#fIdx').val();
  postData={"fIdx":fIdx};
  switch(tab){
    case 'action':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=add"> 创建 <i class="fa fa-plus"></i></a>');
      break;
    case 'remote_step':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=add"> 创建 <i class="fa fa-plus"></i></a>');
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
        switchery();
      }else{
        pageNotify('error','加载失败！','错误信息：'+listdata.msg);
      }
      NProgress.done();
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
            console.log(i+' '+p1);
            li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
          }else{
            li='<li><a href="javascript:;" onclick="list('+i+')">'+i+'</a></li>';
          }
        }else{
          if(i==p2){
            if(p2<data.pageCount-1){
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
        var btnEdit='',btnDel='',btnView='',btnRun='',t='';
        switch(tab){
          case 'action':
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'action\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            td = '<td>' + v.desc + '</td>';
            tr.append(td);
            $.each(v.params,function(key,val){
              t+=(t)?', '+key:key;
            });
            td = '<td>' + t + '</td>';
            tr.append(td);
            var btnEdit = '<a class="btn blue tooltips" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.id + '" style="padding-left: 0px;" title="修改"><i class="fa fa-edit"></i></a>';
            var btnDel = '<a class="btn red tooltips" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')" style="padding-left: 0px;" title="删除"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'remote_step':
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'remote_step\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            td = '<td>' + v.desc + '</td>';
            tr.append(td);
            var ii=0;
            $.each(v.actions,function(key,val){
              ii++;
              t+=(t)?' <i class="fa fa-angle-double-right text-danger"></i> '+'<a class="tooltips" title="第'+ii+'步">'+val+'</a>':'<a class="tooltips" title="第'+ii+'步">'+val+'</a>';
            });
            td = '<td>' + t + '</td>';
            tr.append(td);
            var btnEdit = '<a class="btn blue tooltips" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.id + '" style="padding-left: 0px;" title="修改"><i class="fa fa-edit"></i></a>';
            var btnDel = '<a class="btn red tooltips" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+v.name+'\')" style="padding-left: 0px;" title="删除"><i class="fa fa-trash-o"></i></a>';
            break;
        }
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

//增删改查
var change=function(){
  NProgress.start();
  var tab=$('#tab').val();
  var url='/api/for_layout/'+tab+'.php';
  var postData={};
  var form=$('#myModalBody').find("input,select,textarea");
  var vars='',templates='',tasks='';
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
              if(this.checked){
                  switch (this.name){
                      case 'var':
                          vars+=this.id+',';
                          break;
                      case "template":
                          templates+=this.id+',';
                          break;
                      case 'task':
                          tasks+=this.id+',';
                          break;
                  }
              }
          }
        break;
      default:
        if(this.name) postData[this.name]=this.value;
        break;
    }
  });

    if($("#cmd_child").val()=="role"){
        creatRole(vars,templates,tasks,postData['name'],postData['desc']);
    }
  var action=$("#page_action").val();
  delete postData['page_action'];
  //console.log("action="+action);
  //console.log(JSON.stringify(postData));
  var actionDesc='';
  switch(action){
    case 'insert':
      actionDesc='添加';
      switch(tab){
        case 'action':
          if($('#cmd_child').val()=='role'){
            postData['task_type']="ansible_role";
          }else{
            postData['task_type']="ansible_task";
          }
          postData['params']=cache.params;
          if(postData['config_senior']=='true'){
            postData['template']=JSON.parse(postData['cmd_content']);
            delete(postData['cmd_content']);
          }else{
            var template={};
            template[postData['cmd_parent']]={};
            template[postData['cmd_parent']]['module']=postData['cmd_child'];
            switch(postData['cmd_child']){
              case 'shell':
                template[postData['cmd_parent']]['args']=postData['cmd_content'];
                break;
              case 'longscript':
                template[postData['cmd_parent']]['content']=postData['cmd_content'];
                break;
            }
            postData['template']=template;
            delete(postData['cmd_parent']);
            delete(postData['cmd_child']);
            delete(postData['cmd_content']);
          }
          delete(postData['param_0']);
          delete(postData['type_0']);
          delete(postData['config_senior']);
          break;
        case 'remote_step':
          postData['actions']=[];
          $.each(cache.step_action,function(k,v){
            postData['actions'].push(v);
          });
          break;
      }
      break;
    case 'update':
      actionDesc='更新';
      delete postData['name'];
      switch(tab){
        case 'action':
          postData['params']=cache.params;
          if(postData['config_senior']=='true'){
            postData['template']=JSON.parse(postData['cmd_content']);
            delete(postData['cmd_content']);
          }else{
            var template={};
            template[postData['cmd_parent']]={};
            template[postData['cmd_parent']]['module']=postData['cmd_child'];
            switch(postData['cmd_child']){
              case 'shell':
                template[postData['cmd_parent']]['args']=postData['cmd_content'];
                break;
              case 'longscript':
                template[postData['cmd_parent']]['content']=postData['cmd_content'];
                break;
            }
            postData['template']=template;
            delete(postData['cmd_parent']);
            delete(postData['cmd_child']);
            delete(postData['cmd_content']);
          }
          delete(postData['param_0']);
          delete(postData['type_0']);
          delete(postData['config_senior']);
          break;
        case 'remote_step':
          postData['actions']=[];
          $.each(cache.step_action,function(k,v){
            postData['actions'].push(v);
          });
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

var get = function (idx) {
  var tab=$('#tab').val();
  var url='/api/for_layout/'+tab+'.php';
  var postData={"action":"info","fIdx":idx};
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
                        if(k=='params'){
                          $('#'+k).val(JSON.stringify(v));
                        }else{
                          $('#'+k).val(v);
                        }
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
              if(k=='actions'){
                cache.step_action=v;
                listStepAction();
              }
              if(k=='params'){
                cache.params={};
                if(typeof v == 'object' && !$.isArray(v)) cache.params=v;
                showArg();
              }
              if(k=='id'){
                if(tab=='action') getActimpl(v);
              }
            });
              setTimeout(function(){
                  if($("#cmd_child").val()=='role'){
                      $("#vars_file").parent().attr("hidden",true);
                      $("#tems_file").parent().attr("hidden",true);
                      $("#tasks_file").parent().attr("hidden",true);
                  }
              },100);
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
      case 'action':
        switch(action){
          case 'del':
            modalTitle='删除远程命令';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 命令ID : '+idx +'<br>命令名称 : '+desc;
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
      case 'remote_step':
        switch(action){
          case 'del':
            modalTitle='删除步骤';
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 步骤ID : '+idx +'<br>步骤名称 : '+desc;
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

//commit check
var check=function(tab){
  if(!tab) tab=$('#tab').val();
  switch(tab){
    case 'action':
      var disabled=false;
      var name=$('#name').val();
      if(name==''){
        disabled=true;
      }else{
        if(cache.check_chinese===true) $('#name').val(hasChinese(name));
      }
      if($('#desc').val()=='') disabled=true;
      if($('#type').val()=='') disabled=true;
      if(checkCmd(true)) disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'remote_step':
      var disabled=false;
      var name=$('#name').val();
      if(name==''){
        disabled=true;
      }else{
        if(cache.check_chinese===true) $('#name').val(hasChinese(name));
      }
      if($('#desc').val()=='') disabled=true;
      if(cache.step_action.length==0) disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'action_param':
      var disabled=false;
      if($('#action_name').val()=='') disabled=true;
      $("#btnCommitAction").attr('disabled',disabled);
      break;
  }
}

//view
var view=function(type,idx){
  NProgress.start();
  var url='/api/for_layout/'+type+'.php',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  postData={"action":"info","fIdx":idx};
  switch(type){
    case 'action':
      title='查看命令详情 - '+idx;
      break;
    case 'remote_step':
      title='查看步骤详情 - '+idx;
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
              case 'action':
                $.each(data.content,function(k,v){
                  var t = (k=='params') ? JSON.stringify(v) : v;
                  if(t==''||t=='[]') t='空';
                  if(typeof(locale[k])!='undefined') k=locale[k];
                  text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+t+'</span>'+"\n";
                });
                break;
              case 'remote_step':
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
    title='非法请求';
    text='<div class="note note-danger">错误信息：参数错误</div>';
    $('#myViewModalLabel').html(title);
    $('#myViewModalBody').html(text);
    NProgress.done();
  }
}


var getAction=function(idx){
  var postData={"action":"list","pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_layout/action.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.action = data.content;
          updateSelect('action_name');
          updateEle('action',idx);
        }else{
          pageNotify('warning','加载命令成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/action.php">创建命令</a>]！',false,6000);
        }
      }else{
        pageNotify('error','加载命令失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载命令失败！','错误信息：接口不可用');
    }
  });
}

var listActionParams = function(){
  cache.check_action={};
  var id=$('#action_name').val();
  if(!id) return false;
  var data=cache.action;
  var str='';
  $.each(data, function (k,v) {
    if(v.id==id){
      if(v.params){
        $.each(v.params,function (key,val){
          str+='<tr><td>'+key+'</td><td>'+val+'</td></tr>'+"\n";
        });
      }else{
        str+='<tr><td colspan="3">无参数</td></tr>';
      }
      cache.check_action=v;
    }
  });
  $('#action_param').html(str);
  check('action_param');
}

var updateSelect=function(name){
  var tSelect=$('#'+name),data='';
  switch(name){
    case 'action_name':
      data=cache.action;
      break;
  }
  if(data){
    tSelect.empty();
    tSelect.append('<option value="">请选择</option>');
    $.each(data,function(k,v){
      tSelect.append('<option value="' + v.id + '">' + v.name + '</option>');
    });
  }
}

//处理Action参数值
var setAction=function(idx){
  cache.step_action.push(cache.check_action.name);
  listStepAction();
  check();
  //处理模态框和表单
  $("#myChildModal :input").each(function () {
    $(this).val("");
  });
  $("#myChildModal").on("hidden.bs.modal", function() {
    $(this).removeData("bs.modal");
  });
}

//列出已有步骤
var listStepAction=function(){
  var str='',i=0,up='',down='';
  $.each(cache.step_action,function(k,v){
    up='<a class="btn green btn-xs tooltips" title="上移" style="padding-left: 0px;" onclick="sortStepAction('+ i +','+(i-1)+')"><i class="fa fa-arrow-up"></i></a>';
    down='<a class="btn purple btn-xs tooltips" title="下移" style="padding-left: 0px;" onclick="sortStepAction('+ i +','+(i+1)+')"><i class="fa fa-arrow-down"></i></a>';
    if(i==0) up='<a class="btn dark btn-xs" style="padding-left: 0px;" onclick="return false;" disabled><i class="fa fa-arrow-up"></i></a>';
    i++;
    if(i==cache.step_action.length) down='<a class="btn dark btn-xs" style="padding-left: 0px;" onclick="return false;" disabled><i class="fa fa-arrow-down"></i></a>';
    str+='<tr><td>'+i+'</td><td>'+ v +'</td>' +
      '<td>' + up + down +
      '<a class="btn red btn-xs tooltips" title="删除" style="padding-left: 0px;" onclick="delStepAction('+ (i-1) +')"><i class="fa fa-trash-o"></i></a></td>' +
      '</tr>';
  });
  $('#step_action').html(str);
  $('.tooltips').each(function(){$(this).tooltip();});
}

//删除步骤
var delStepAction=function(idx){
  if(!$.isNumeric(idx)) return false;
  if(idx<cache.step_action.length) cache.step_action.splice(idx,1);
  listStepAction();
  check();
}

//排序步骤
var sortStepAction=function(i,j){
  if(!$.isNumeric(i) || !$.isNumeric(j)) return false;
  if(i<0||j>=cache.step_action.length) return false;
  var v=cache.step_action[i];
  cache.step_action[i]=cache.step_action[j];
  cache.step_action[j]=v;
  listStepAction();
  check();
}

//修改步骤参数
var updateEle=function(o,idx){
  var t='',data='';
  switch(o){
    case 'action':
      //修改步骤时显示
      if(!$.isNumeric(idx)) return false;
      data=cache.step_action[idx];
      $('#action_name').val(data.id).trigger('change');
      if(data.params){
        $.each(data.params,function(k,v){
          $('#param_'+k).val(v);
        });
      }
      break;
    case 'name':
      t=$('#task_type').val();
      if(!t) return false;
      $('#task_name').val(t);
      get('arg',t);
      break;
  }
}

//列出参数
var showArg=function(){
  var data=cache.params;
  var body=$('#args'),td='';
  body.html('');
  var i=1;
  $.each(data,function(k,v){
    var tr=$('<tr></tr>');
    td='<td><input type="text" class="form-control input-sm" id="param_'+i+'" value="'+k+'" placeholder="参数名称" readonly></td>';
    tr.append(td);
    switch (v){
      case 'string':
        td='<td><select class="input-sm" id="type_'+i+'" style="width: 100%;"><option value="string" selected>string</option><option value="integer">integer</option></select></td>';
        break;
      case 'integer':
        td='<td><select class="input-sm" id="type_'+i+'" style="width: 100%;"><option value="string">string</option><option value="integer" selected>integer</option></select></td>';
        break;
    }
    tr.append(td);
    td='<td><a class="tooltips text-danger" title="删除" onclick="delArg(\''+k+'\')"><span style="padding-top: 8px;">删除</span></a></td>';
    tr.append(td);
    body.append(tr);
  });
  for(var i=0;i<data.length;i++){
    var j=i+1;
  }
  var tr=$('<tr></tr>');
  td='<td><input type="text" class="form-control input-sm" id="param_0" name="param_0" onkeyup="checkParam(this)" placeholder="参数名称"></td>';
  tr.append(td);
  td='<td><select class="input-sm" id="type_0" name="type_0" style="width: 100%;"><option value="string">string</option><option value="integer">integer</option></select></td>';
  tr.append(td);
  td='<td><a class="tooltips text-success" title="确认添加参数" onclick="addArg()"><span style="padding-top: 8px;">添加</span></a></td>';
  tr.append(td);
  body.append(tr);
  $('.tooltips').each(function(){$(this).tooltip();});
  check();
}

//添加参数
var addArg=function(){
  var key=$('#param_0').val();
  var type=$('#type_0').val();
  if(!key || !type){
    $('#param_0')[0].focus();
    return false;
  }
  cache.params[key]=type;
  showArg();
}

//移除参数
var delArg=function(idx){
  $.each(cache.params,function(k,v){
    if(k==idx) delete(cache.params[idx]);
  });
  showArg();
}

//获取命令实现
var getActimpl=function(idx){
  var actionDesc='命令实现';
  var postData={"action":"info","fIdx":idx};
  $.ajax({
    type: "POST",
    url: '/api/for_layout/actimpl.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content){
          $('#type').val(data.content.type).trigger('change');
          cache.actimpl=data.content;
          showCmd();
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

//配置模式
var updateSenior= function () {
  if($('#config_senior').val()=='false'){
    $('#cmd_parent').parent().parent().removeClass('hidden');
    $('#cmd_child').parent().parent().removeClass('hidden');
    if($('#page_action').val()=='insert'){
      $('#cmd_content').val('');
    }else{
      showCmd();
    }
  }else{
    $('#cmd_parent').parent().parent().addClass('hidden');
    $('#cmd_child').parent().parent().addClass('hidden');
    if($('#page_action').val()=='insert'){
      $('#cmd_content').val(JSON.stringify(cache.action_default,null,"\t"));
    }else{
      $('#cmd_content').val(JSON.stringify(cache.actimpl.template,null,"\t"));
    }
  }
}

var updateCmd=function(){
  checkCmd();
}
getModal= function(){
    alert("getmodal");
}
var checkCmd=function(fromCheck){
  var disabled=false;
  if($('#config_senior').val()=='false'){
    if($('#cmd_parent').val()=='') disabled=true;
    if($('#cmd_child').val()=='') disabled=true;
    if($('#cmd_content').val()=='') disabled=true;
  }else{
    if($('#cmd_content').val()=='') disabled=true;
  }

  if(fromCheck) return disabled;
  $("#btnCommit").attr('disabled',disabled);
  if(!disabled) check();
}


var showRole=function(){
    if($('#cmd_child').val()=='role'){
        $("#cmd_content").parent().parent().addClass('hidden');
        $("#cmd_content").val("{}");
        $("#vars_file").parent().attr("hidden",false);
        $("#tems_file").parent().attr("hidden",false);
        $("#tasks_file").parent().attr("hidden",false);
        var postData = {" action": "list", "pagesize": 1000};
        var url = '/api/for_layout/roleresource.php?action=list&pagesize=1000';
        var actionDesc="添加Role";
        $.ajax({
            type: "POST",
            url: url,
            data: {"action":'list',"data":JSON.stringify(postData)},
            dataType: "json",
            success: function (data) {
                //执行结果提示
                if(data.code==0){
                    $('#vars_file').html('');
                    $('#tems_file').html('');
                    $('#tasks_file').html('');
                    for(var i=0;i<data.content.length;i++){
                        var v = data.content[i];
                        switch (v.resource_type){
                            case 'var':
                                var var_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="'+v.id+'" name="var">'+v.name+'</span>';
                                $('#vars_file').append(var_checkboxes);
                                break;
                            case "template":
                                var tem_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="'+v.id+'" name="template">'+v.name+'</span>';
                                $('#tems_file').append(tem_checkboxes);
                                break;
                            case 'task':
                                var task_checkboxes = '<span class="col-sm-2"><input type="checkbox" id="'+v.id+'" name="task">'+v.name+'</span>';
                                $('#tasks_file').append(task_checkboxes);
                                break;
                        }

                    }
                }else{
                    pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+data.msg);
                }

            },
            error: function (){
                pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
            }
        });
    }else{
        $("#cmd_content").val('');
        $("#cmd_content").parent().parent().removeClass('hidden');
        $("#vars_file").parent().attr("hidden",true);
        $("#tems_file").parent().attr("hidden",true);
        $("#tasks_file").parent().attr("hidden",true);

    }
    checkCmd();
}
var showCmd= function () {
  var o = cache.actimpl;
  if(o){
    var cmd_parent='',cmd_child='';
    $.each(o.template,function(k,v){
      if(!cmd_parent) cmd_parent=k;
      if(cmd_parent=='action'){
        if(typeof v.module != 'undefined') cmd_child = v.module;
      }
    });
    if(o.type=='ansible'&&cmd_parent=='action'){
        if($('#cmd_parent').find("option[value='action']").length==0) $('#cmd_parent').append('<option value="'+cmd_parent+'">'+cmd_parent+'</option>');
        $('#cmd_parent').val(cmd_parent).trigger('change');
        if(cmd_child){
          if(cmd_child=='shell'||cmd_child=='longscript'){
            $('#cmd_child').val(cmd_child).trigger('change');
            switch (cmd_child){
              case 'shell':
                $('#cmd_content').val(o.template.action.args);
                break;
              case 'longscript':
                $('#cmd_content').val(o.template.action.content);
                break;
            }
          }else{
            if($('#cmd_child').find("option[value='"+cmd_child+"']").length==0) $('#cmd_child').append('<option value="'+cmd_child+'">'+cmd_child+'</option>');
            $('#cmd_child').val(cmd_child).trigger('change');
            $('#cmd_content').val(JSON.stringify(o.template,null,"\t"));
          }
        }else{
          $('#cmd_child').append('<option value="'+cmd_child+'">空</option>');
          $('#cmd_child').val(cmd_child).trigger('change');
          $('#cmd_content').val(JSON.stringify(o.template,null,"\t"));
        }
    }else{
      $('#config_senior').val('true').trigger('change');
      $('#cmd_content').val(JSON.stringify(o.template,null,"\t"));
    }
  }
}

var checkParam = function(o){
  if(!cache.check_chinese) return false;
  var name=$(o).val();
  if(name!=''){
    $(o).val(hasChinese(name));
  }
}

var creatRole=function(vars,templates,tasks,name,desc){
    NProgress.start();
    var tab='role'
    var url='/api/for_layout/role.php';
    var postData={};
    var actionDesc="添加"

    vars=vars.substring(0,vars.length-1);
    tasks=tasks.substring(0,tasks.length-1);
    templates=templates.substring(0,templates.length-1);
    postData['vars']=vars;
    postData['templates']=templates;
    postData['tasks']=tasks;
    postData['name']=name;
    postData['desc']=desc;
    var action='insert';
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

        },
        error: function (){
            pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
        }
    });
}

// var creatRole=function(){
//     NProgress.start();
//     var tab='role'
//     var url='/api/for_layout/role.php';
//     var postData={};
//     var form=$('#myRoleModal').find("input,select,textarea");
//     //处理表单内容--不需要修改
//     var vars='',templates='',tasks='';
//     var actionDesc="添加"
//     $.each(form,function(i){
//         switch(this.type){
//             case 'radio':
//                 if(typeof(postData[this.name])=='undefined'){
//                     if(this.name) postData[this.name]=$('input[name="'+this.name+'"]:checked').val();
//                 }
//                 break;
//             case 'checkbox':
//                 if(this.id){
//                   if(this.checked){
//                       switch (this.name){
//                           case 'var':
//                              vars+=this.id+',';
//                              break;
//                           case "template":
//                               templates+=this.id+',';
//                               break;
//                           case 'task':
//                               tasks+=this.id+',';
//                               break;
//                       }
//                   }
//                 }
//                 break;
//             default:
//                 if(this.name) postData[this.name]=this.value;
//                 break;
//         }
//     });
//     vars=vars.substring(0,vars.length-1);
//     tasks=tasks.substring(0,tasks.length-1);
//     templates=templates.substring(0,templates.length-1);
//     postData['vars']=vars;
//     postData['templates']=templates;
//     postData['tasks']=tasks;
//     var action='insert';
//     $.ajax({
//         type: "POST",
//         url: url,
//         data: {"action":action,"data":JSON.stringify(postData)},
//         dataType: "json",
//         success: function (data) {
//             //执行结果提示
//             if(data.code==0){
//                 pageNotify('success','【'+actionDesc+'】操作成功！');
//             }else{
//                 pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：'+data.msg);
//             }
//             //处理模态框和表单
//             $("#myRoleModal :input").each(function () {
//                 $(this).val("");
//             });
//             $("#myRoleModal").on("hidden.bs.modal", function() {
//                 $(this).removeData("bs.modal");
//             });
//
//
//             $("#myRoleModal").on("hidden.bs.modal", function() {
//                 setTimeout(function(){
//                     $('body').addClass('modal-open')
//                 },800)
//
//             });
//
//
//         },
//         error: function (){
//             pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
//         }
//     });
// }