cache = {
  page: 1,
  check_step: {}, //Action详情
  task_step: [], //步骤
  task_tpl: {}, //模板列表
  cluster: [], //集群列表
  service: [], //服务列表
  pool: [], //服务池列表
  ip: [], //选中IP列表
}

var getDate = function(){
  var d=new Date();
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
  return '_'+d.getFullYear()+month+day+'_'+ hour+min+sec;
};

var formatDate = function(t){
  if(!t) t='';
  var d=new Date(t);
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
  if(tab!='task_tpl'&&tab!='task'){
    tab='task';
  }
  if(idx) $('#fIdx').val(idx);
  var fIdx=$('#fIdx').val();
  postData={"action":"list","fIdx":fIdx};
  switch(tab){
    case 'task_tpl':
      $('#tab_1').attr('class','active');
      $('#tab_2').attr('class','');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=add"> 创建 <i class="fa fa-plus"></i></a>');
      break;
    case 'task':
      $('#tab_1').attr('class','');
      $('#tab_2').attr('class','active');
      $('#tab_toolbar').html('<a type="button" class="btn btn-success" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=add"> 发起新任务 <i class="fa fa-plus"></i></a>');
      break;
  }
  $('#tab').val(tab);
  var url='/api/for_layout/'+tab+'.php';
  if (!page) {
    page = cache.page;
  }else{
    cache.page = page;
  }
  url+='?page=' + page;
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
          case 'task_tpl':
            td = '<td><a class="tooltips" title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'task_tpl\',\''+v.id+'\')">' + v.name + '</a></td>';
            tr.append(td);
            td = '<td>' + v.desc + '</td>';
            tr.append(td);
            var ii=0;
            $.each(v.steps,function(key,val){
              ii++;
              t+=(t)?' <i class="fa fa-angle-double-right text-danger"></i> '+'<a class="tooltips" title="第'+ii+'步">'+val.name+'</a>':'<a class="tooltips" title="第'+ii+'步">'+val.name+'</a>';
            });
            td = '<td>' + t + '</td>';
            tr.append(td);
            //btnRun = '<a class="text-primary tooltips" data-container="body" data-trigger="hover" data-original-title="使用" data-toggle="modal" data-target="#myListModal" href="edit_task.php?idx=' + v.id + '"><i class="fa fa-play"></i></a>';
            btnEdit = '<a class="text-success tooltips" data-container="body" data-trigger="hover" data-original-title="修改" data-toggle="modal" data-target="#myModal" href="edit_'+tab+'.php?action=edit&idx=' + v.id + '"><i class="fa fa-edit"></i></a>';
            btnDel = '<a class="text-danger tooltips" data-container="body" data-trigger="hover" data-original-title="删除" data-toggle="modal" data-target="#myModal" onclick="twiceCheck(\'del\',\''+v.id+'\',\''+ v.name +'\')"><i class="fa fa-trash-o"></i></a>';
            break;
          case 'task':
            td = '<td title="'+ v.id +'" style="width:200px;word-wrap:break-word;word-break:break-all;"><a class="tooltips" data-container="body" data-trigger="hover" data-original-title="查看详情" data-toggle="modal" data-target="#myViewModal" onclick="view(\'task\',\''+v.id+'\')">' + v.task_name + '</a></td>';
            tr.append(td);
            var tState='';
            switch(v.state){
              case 0: tState='<span class="badge bg-default">未开始</span>'; break;
              case 1: tState='<span class="badge bg-blue">执行中</span>'; break;
              case 2: tState='<span class="badge bg-green">成功</span>'; break;
              case 3: tState='<span class="badge bg-red">失败</span>'; break;
              case 4: tState='<span class="badge bg-orange">已暂停</span>'; break;
              default: tState='<span class="badge bg-red" title="'+ v.state +'">未知</span>'; break;
            }
            td = '<td>' + tState + '</td>';
            tr.append(td);
            td = '<td>' + v.opr_user + '</td>';
            tr.append(td);
            td = '<td>' + formatDate(v.created) + '</td>';
            tr.append(td);
            btnView = '<a class="text-success tooltips" title="任务详情" href="task_detail.php?idx='+ v.id +'"><i class="fa fa-info-circle"></i></a>';
            break;
        }
        td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnRun + ' ' + btnEdit + ' ' + btnDel + ' ' + btnView + '</div></td>';
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
  //console.log("action="+action);
  //console.log(JSON.stringify(postData));
  var actionDesc='';
  switch(action){
    case 'insert':
      actionDesc='添加';
      if(tab=='task_tpl'){
        postData['steps']=cache.task_step;
      }
      break;
    case 'update':
      actionDesc='更新';
      delete postData['name'];
      if(tab=='task_tpl'){
        postData['steps']=cache.task_step;
      }
      break;
    case 'delete':
      actionDesc='删除';
      break;
    default:
      actionDesc=action;
      break;
  }
  //console.log({"action":action,"data":JSON.stringify(postData)});
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
  url='/api/for_layout/'+type+'.php';
  postData={"action":"info","fIdx":idx};
  switch(type){
    case 'task_tpl':
      title='查看任务模板详情 - '+idx;
      break;
    case 'task':
      title='查看任务详情 - '+idx;
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
            var json='';
            switch(type){
              case 'task_tpl':
                $.each(data.content,function(k,v){
                  if(locale[k]!=false) {
                    if (k == 'steps') {
                      json += '<textarea class="form-control" rows="20">' + JSON.stringify(v, null, "\t") + '</textarea>';
                    } else {
                      if (v == '') v = '空';
                      if (typeof(locale[k]) != 'undefined') k = locale[k];
                      text += '<span class="title col-sm-2" style="font-weight: bold;"><strong>' + k + '</strong></span> ' +
                        '<span class="col-sm-4" style="' + tStyle + '">' + v + '</span>' + "\n";
                    }
                  }
                });
                text+=(json)?'<div class="row col-sm-12"><hr class="col-sm-12" style="margin-bottom: 0px;margin-top: 10px;" /></div><h4 class="col-sm-12 text-primary"><strong>步骤</strong></h4>' +
                json:'<span class="col-sm-2"><strong>步骤</strong></span> ' +
                '<span class="col-sm-4" style="'+tStyle+'">步骤</span>'+"\n";
                break;
              default:
                $.each(data.content,function(k,v){
                  if(locale[k]!=false){
                    if(k=='options'){
                      json+='<textarea class="form-control" rows="16">'+JSON.stringify(v,null,"\t")+'</textarea>';
                    }else {
                      if (v == '') v = '空';
                      if (typeof(locale[k]) != 'undefined') k = locale[k];
                      text += '<span class="title col-sm-2" style="font-weight: bold;"><strong>' + k + '</strong></span> <span class="col-sm-4" style="' + tStyle + '">' + v + '</span>' + "\n";
                    }
                  }
                });
                text+=(json)?'<div class="row col-sm-12"><hr class="col-sm-12" style="margin-bottom: 0px;margin-top: 10px;" /></div><h4 class="col-sm-12 text-primary"><strong>任务参数</strong></h4>' +
                json:'<span class="col-sm-2"><strong>参数</strong></span> ' +
                '<span class="col-sm-4" style="'+tStyle+'">任务参数</span>'+"\n";
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
        },1000);
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
    case 'task_tpl':
      var disabled=false;
      if($('#name').val()=='') disabled=true;
      if($('#desc').val()=='') disabled=true;
      if(cache.task_step.length==0) disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
    case 'step_param':
      var disabled=false;
      if($('#step_name').val()=='') disabled=true;
      if(cache.check_step.params){
        $.each(cache.check_step.params,function (key,val){
          if($('#param_'+key).val()=='') disabled=true;
        });
      }
      if($('#retry_times').val()=='') disabled=true;
      if($('#ignore_error').val()=='') disabled=true;
      $("#btnCommitAction").attr('disabled',disabled);
      break;
    case 'task':
      var disabled=false;
      if($('#template_id').val()=='') disabled=true;
      if($('#task_name').val()=='') disabled=true;
      if($('#auto').val()=='') disabled=true;
      if($('#max_num').val()=='') disabled=true;
      if($('#max_ratio').val()=='') disabled=true;
      if(parseInt($('#run_num').html())<1) disabled=true;
      $("#btnCommit").attr('disabled',disabled);
      break;
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
      case 'task_tpl':
        switch(action){
          case 'del':
            modalTitle='删除任务模板 - '+idx;
            modalBody=modalBody+'<div class="form-group col-sm-12">';
            modalBody=modalBody+'<div class="note note-danger">';
            modalBody=modalBody+'<h4>确认删除? <span class="text text-primary">警告! 操作不可回退!</span></h4> 模板ID : '+idx+'<br/>模板名称 : '+desc;
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

var getStep=function(idx){
  var postData={"action":"list","pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_layout/task_step.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.step = data.content;
          updateSelect('step_name');
          updateEle('step',idx);
        }else{
          pageNotify('info','加载服务成功！','数据为空！');
        }
      }else{
        pageNotify('error','加载服务失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','加载服务失败！','错误信息：接口不可用');
    }
  });
}

var listActionParams = function(){
  cache.check_step={};
  var id=$('#step_name').val();
  if(!id) return false;
  var data=cache.step;
  var str='';
  $.each(data, function (k,v) {
    if(v.name==id){
      if(v.params){
        $.each(v.params,function (key,val){
          str+='<tr><td>'+key+'</td><td><input type="text" class="form-control input-sm" id="param_'+key+'" name="param_'+key+'" placeholder="'+val+'" onkeyup="check(\'step_param\')"></td></tr>'+"\n";
        });
      }else{
        str+='<tr><td colspan="3">无参数</td></tr>';
      }
      cache.check_step=v;
    }
  });
  $('#step_param').html(str);
  check('step_param');
}

var updateSelect=function(name){
  var tSelect=$('#'+name),data='';
  switch(name){
    case 'step_name':
      data=cache.step;
      break;
    case 'template_id':
      data=cache.task_tpl;
      break;
    case 'cluster':
      data=cache.cluster;
      break;
    case 'service':
      data=cache.service;
      break;
    case 'pool':
      data=cache.pool;
      break;
  }
  if(data){
    tSelect.empty();
    tSelect.append('<option value="">请选择</option>');
    switch(name){
      case 'step_name':
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.name + '">' + v.name + '</option>');
        });
        break;
      default:
        $.each(data,function(k,v){
          tSelect.append('<option value="' + v.id + '">' + v.name + '</option>');
        });
        break;
    }
  }
  tSelect.select2();
}

var clearSelect=function(name){
  var tSelect=$('#'+name);
  tSelect.empty();
  tSelect.append('<option value="">请选择</option>');
  tSelect.select2();
}

//处理Action参数值
var setAction=function(idx){
  var t={},tt={},retry_times=0,ignore_error='false';
  $.each(cache.check_step,function(k,v){
    switch(k){
      case 'params':
        if(v){
          $.each(v,function(key,val){
            if(val.toLowerCase()=='integer'){
              tt[key]=parseInt($('#param_'+key).val());
            }else{
              tt[key]=$('#param_'+key).val();
            }
          });
          t['param_values']=tt;
        }else{
          t['param_values']='';
        }
        break;
      case 'name':
        t[k]= v;
        break;
    }
  });
  retry_times=parseInt($('#retry_times').val());
  ignore_error=($('#ignore_error').val()=='false')?false:true;
  t['retry']={
    retry_times: retry_times,
    ignore_error: ignore_error
  };
  if(idx){
    cache.task_step[idx]=t;
  }else{
    cache.task_step.push(t);
  }
  listTaskStep();
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
var listTaskStep=function(){
  var str='',i=0,up='',down='';
  $.each(cache.task_step,function(k,v){
    up='<a class="btn green btn-xs tooltips" title="上移" style="padding-left: 0px;" onclick="sortTaskStep('+ i +','+(i-1)+')"><i class="fa fa-arrow-up"></i></a>';
    down='<a class="btn purple btn-xs tooltips" title="下移" style="padding-left: 0px;" onclick="sortTaskStep('+ i +','+(i+1)+')"><i class="fa fa-arrow-down"></i></a>';
    if(i==0) up='<a class="btn dark btn-xs" style="padding-left: 0px;" onclick="return false;" disabled><i class="fa fa-arrow-up"></i></a>';
    i++;
    if(i==cache.task_step.length) down='<a class="btn dark btn-xs" style="padding-left: 0px;" onclick="return false;" disabled><i class="fa fa-arrow-down"></i></a>';
    str+='<tr><td>'+i+'</td><td>'+ v.name +'</td><td>'+ ((v.retry)?v.retry.retry_times:'-') +'</td><td>'+ ((v.retry)?((v.retry.ignore_error)?'忽略':'否'):'-') +'</td>' +
      '<td><a class="btn blue btn-xs tooltips" title="修改" data-toggle="modal" data-target="#myChildModal" href="edit_task_tpl_step.php?idx='+ (i-1) +'"><i class="fa fa-edit"></i></a>' +
      up + down +
      '<a class="btn red btn-xs tooltips" title="删除" style="padding-left: 0px;" onclick="delTaskStep('+ (i-1) +')"><i class="fa fa-trash-o"></i></a></td>' +
      '</tr>';
  });
  $('#task_step').html(str);
  $('.tooltips').each(function(){$(this).tooltip();});
}

//删除步骤
var delTaskStep=function(idx){
  if(!$.isNumeric(idx)) return false;
  if(idx<cache.task_step.length) cache.task_step.splice(idx,1);
  listTaskStep();
  check();
}

//排序步骤
var sortTaskStep=function(i,j){
  if(!$.isNumeric(i) || !$.isNumeric(j)) return false;
  if(i<0||j>=cache.task_step.length) return false;
  var v=cache.task_step[i];
  cache.task_step[i]=cache.task_step[j];
  cache.task_step[j]=v;
  listTaskStep();
  check();
}

//修改步骤参数
var updateEle=function(o,idx){
  var t='',data='',retry_times=0,ignore_error='false';
  switch(o){
    case 'step':
      //修改步骤时显示
      if(!$.isNumeric(idx)) return false;
      data=cache.task_step[idx];
      $('#step_name').val(data.name).trigger('change');
      if(data.param_values){
        $.each(data.param_values,function(k,v){
          $('#param_'+k).val(v);
        });
      }
      if(data.retry){
        retry_times=data.retry.retry_times;
        ignore_error=data.retry.ignore_error.toString();
      }
      $('#retry_times').val(retry_times).trigger('change');
      $('#ignore_error').val(ignore_error).trigger('change');
      break;
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
          updateSelect('template_id');
        }else{
          pageNotify('warning','加载'+actionDesc+'成功！','数据为空！请先[<a class="tooltips text-danger" title="点击跳转" href="../../for_layout/task_tpl.php">创建任务模板</a>]！',true,6000);
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

//设置任务名称
var setTaskName=function(){
  var id=$('#template_id').val();
  if(!id) return false;
  for(var i=0;i<cache.task_tpl.length;i++){
    if(cache.task_tpl[i].id==id){
      $('#task_name').val(cache.task_tpl[i].name+getDate());
    }
  }
  check();
}

var get = function (idx) {
  tab=$('#tab').val();
  var postData={"action":"info","fIdx":idx};
  var url='/api/for_layout/'+tab+'.php';
  if(idx){
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
              if(k=='app') k='service';
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
            switch (tab){
              case 'task_tpl':
                cache.task_step=data.content.steps;
                listTaskStep();
                break;
              default:
                break;
            }
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
    title='非法请求 - '+tab;
  }
}

//获取集群列表
var getCluster=function(){
  clearSelect('service');
  clearSelect('pool');
  var actionDesc='集群列表';
  var postData={"action":"list","pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_layout/cluster.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.cluster = data.content;
          updateSelect('cluster');
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

//获取服务列表
var getService=function(){
  var fIdx=$('#cluster').val();
  clearSelect('service');
  clearSelect('pool');
  if(!fIdx) return false;
  var actionDesc='服务列表';
  var postData={"action":"list","fClusterId":fIdx,"pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_layout/service.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.service = data.content;
          updateSelect('service');
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

//获取服务池列表
var getPool=function(){
  var fIdx=$('#service').val();
  clearSelect('pool');
  if(!fIdx) return false;
  var actionDesc='服务池列表';
  var postData={"action":"list","fService":fIdx,"pagesize":1000};
  $.ajax({
    type: "POST",
    url: '/api/for_layout/pool.php',
    data: postData,
    dataType: "json",
    success: function (data) {
      if(data.code==0){
        if(data.content.length>0){
          cache.pool = data.content;
          updateSelect('pool');
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
    check();
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

var checkTask=function(){
  var disabled=false;
  if($('#task_type').val()=='') disabled=true;
  if($('#task_name').val()=='') disabled=true;
  if($('#service').val()=='') disabled=true;
  if($('#tag').val()=='') disabled=true;
  if($('#auto').val()=='') disabled=true;
  if($('#rate').val()=='') disabled=true;
  if($('#num').val()=='') disabled=true;
  if($('#timeout').val()=='') disabled=true;
  if($('#service_path').val()=='') disabled=true;
  //检查参数
  if(cache.arg_name.length>0){
    for(var i=0;i<cache.arg_name.length;i++){
      if(cache.arg_name[i].optional!='1'){
        if($('#value_'+cache.arg_name[i].arg_name).val()==''){
          disabled=true;
        }
      }
    }
  }
  //检查目标
  var num=0;
  if(!disabled){
    $.each(cache.pools,function(p,l){
      if(l.length>0){
        num+= l.length;
      }
    });
  }
  if(num==0) disabled=true;
  $('#form_wizard_1').find('.button-submit').attr('disabled',disabled);
}

var controlTask=function(action,order,idx,ip){
  var postData='',actionDesc='';
  var url='/api/for_layout/task.php';
  switch(action){
    case 'task_command':
      postData={"action":action,"data":{"order":order,"task_id":idx}};
      switch(order){
        case 'start':
          actionDesc='启动任务';
          break;
        case 'pause':
          actionDesc='暂停任务';
          break;
        case 'finish':
          actionDesc='完成任务';
          break;
        case 'modify':
          actionDesc='修改任务';
          break;
        default:
          postData='';
          break;
      }
      break;
    case 'sub_command':
      postData={"action":action,"data":{"order":order,"task_id":idx,"ip":ip}};
      switch(order){
        case 'exec':
          actionDesc='执行子任务';
          break;
        case 'redo':
          actionDesc='重做子任务';
          break;
        case 'skip':
          actionDesc='跳过子任务';
          break;
        default:
          postData='';
          break;
      }
      break;
    case 'agent_command':
      switch(order){
        case 'cancel':
          postData={"action":action,"data":{"order":order,"task_id":idx}};
          actionDesc='操作Agent-取消任务';
          break;
        default:
          postData='';
          break;
      }
      break;
  }
  if(postData){
    $.ajax({
      type: "POST",
      url: url,
      data: postData,
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
        $("#myModal").draggable({
          handle: ".modal-header"
        });
      },
      error: function (){
        pageNotify('error','【'+actionDesc+'】操作失败！','错误信息：接口不可用');
      }
    });
  }else{
    pageNotify('error','非法请求！','错误信息：参数错误');
  }
}

var twiceCheckTask=function(action,order,idx,ip){
  NProgress.start();
  if(!idx) idx='';
  var modalTitle='',modalBody='',notice='',btnDisable=false;
  if(!action||!order||!idx){
    modalTitle='非法请求';
    notice='<div class="note note-danger">错误信息：参数错误</div>';
    pageNotify('error','非法请求！','错误信息：参数错误');
  }else{
    switch(action){
      case 'task_command':
        switch(order){
          case 'start':
            modalTitle='启动任务';
            break;
          case 'pause':
            modalTitle='暂停任务';
            break;
          case 'finish':
            modalTitle='完成任务';
            break;
          case 'modify':
            modalTitle='修改任务';
            break;
          default:
            modalTitle='';
            break;
        }
        break;
      case 'sub_command':
        switch(order){
          case 'exec':
            modalTitle='执行子任务';
            break;
          case 'redo':
            modalTitle='重做子任务';
            break;
          case 'skip':
            modalTitle='跳过子任务';
            break;
          default:
            modalTitle='';
            break;
        }
        break;
      case 'agent_command':
        switch(order){
          case 'cancel':
            modalTitle='取消任务';
            break;
          default:
            modalTitle='';
            break;
        }
        break;
    }
    modalBody=modalBody+'<div class="form-group col-sm-12">';
    modalBody=modalBody+'<div class="note note-danger">';
    modalBody=modalBody+'<h4>确认'+modalTitle+'? <span class="text text-primary">警告! 请谨慎操作!</span></h4>';
    modalBody=modalBody+'任务ID: '+idx+'<br/>';
    if(ip) modalBody=modalBody+'目标IP: '+ip+'<br/>';
    modalBody=modalBody+'</div>';
    modalBody=modalBody+'</div>';
    modalBody=modalBody+'<input type="hidden" id="page_action" name="page_action" value="'+action+'">';
    modalBody=modalBody+'<input type="hidden" id="order" name="order" value="'+order+'">';
    modalBody=modalBody+'<input type="hidden" id="task_id" name="task_id" value="'+idx+'">';
    if(ip) modalBody=modalBody+'<input type="hidden" id="ip" name="ip" value="'+ip+'">';
  }
  if(!modalTitle){
    modalTitle='非法请求';
    notice='<div class="note note-danger">错误信息：参数错误</div>';
    pageNotify('error','非法请求！','错误信息：参数错误');
  }
  modalTitle+=' - '+idx;
  if(ip) modalTitle+=' / '+idx;
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

