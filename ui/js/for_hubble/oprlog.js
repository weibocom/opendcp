cache = {
  page: 1,
  content: {},
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
  if(tab!='oprlog'&&tab!='history') tab='oprlog';
  var fIdx=$('#fIdx').val();
  postData={"fIdx":fIdx};
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
      cache.content = listdata;
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
  var td="",tab=$('#tab').val();
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
        td = '<td>' + v.i + '</td>';
        tr.append(td);
        var btnView='';
        switch (tab) {
          case 'oprlog':
            td = '<td>' + v.module + '</td>';
            tr.append(td);
            td = '<td>' + v.operation + '</td>';
            tr.append(td);
            td = '<td>' + v.user + '</td>';
            tr.append(td);
            td = '<td>' + v.opr_time + '</td>';
            tr.append(td);
            btnView = '<a class="text-primary tooltips" data-original-title="查看详情" data-placement="right" data-toggle="modal" data-target="#myViewModal" onclick="view(\'oprlog\',' + v.id + ')"><i class="fa fa-history"></i></a>';
            td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnView + '</div></td>';
            tr.append(td);
            break;
          case 'history':
            td = '<td>' + v.task_name + '</td>';
            tr.append(td);
            td = '<td>' + v.type + '</td>';
            tr.append(td);
            td = '<td>' + v.channel + '</td>';
            tr.append(td);
            td = '<td>' + v.opr_user + '</td>';
            tr.append(td);
            td = '<td>' + v.create_time + '</td>';
            tr.append(td);
            btnView = '<a class="text-primary tooltips" data-original-title="查看详情" data-placement="right" data-toggle="modal" data-target="#myViewModal" onclick="viewResult(\'state\',' + v.id + ')"><i class="fa fa-history"></i></a>';
            td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnView + '</div></td>';
            tr.append(td);
            break;
        }
        body.append(tr);
      }
    }else{
      pageNotify('info','Warning','数据为空！');
    }
  }else{
    pageNotify('warning','error','接口异常！');
  }
}

var view=function(type,idx){
  NProgress.start();
  var url='',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  postData={"action":"info","fIdx":idx};
  title='查看详情 - '+idx;
  if(idx!=''){
    if(!$.isEmptyObject(cache.content)){
      var str='';
      for(var i=0;i<cache.content.content.length;i++){
        var val=cache.content.content[i];
        if(val.id != idx) continue;
        var locale={};
        if(typeof(locale_messages.hubble.opr_log)) locale = locale_messages.hubble.opr_log;
        $.each(val,function(k,v){
          if(k=='args'){
            str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
              '<h5 class="col-sm-12 text-primary"><strong>请求参数</strong></h5>' +
              '<textarea class="form-control" rows="8">'+v+'</textarea>';
          }else{
            if(k!='i'){
              if(v=='') v='空';
              if(typeof(locale[k])!='undefined') k=locale[k];
              text+='<span class="title col-sm-2" style="font-weight: bold;">'+k+'</span> <span class="col-sm-4" style="'+tStyle+'">'+v+'</span>'+"\n";
            }
          }
        });
      }
      text+=str;
      if(!text){
        pageNotify('warning','未找到对应数据！');
        text='<div class="note note-warning">未找到对应数据！</div>';
      }
    }else{
      pageNotify('warning','数据为空！');
      text='<div class="note note-warning">数据为空！</div>';
    }
  }else{
    pageNotify('warning','错误操作！','错误信息：参数错误');
    title='非法请求 - '+action;
    text='<div class="note note-danger">错误信息：参数错误</div>';
  }
  $('#myViewModalLabel').html(title);
  $('#myViewModalBody').html(text);
  NProgress.done();
}

var viewResult=function(type,idx){
  NProgress.start();
  var url='',title='',text='',illegal=false,height='',postData={};
  var tStyle='word-break:break-all;word-warp:break-word;';
  url='/api/for_hubble/'+type+'.php';
  postData={"action":"info","fIdx":idx};
  switch(type){
    case 'state':
      title='查看详情 - '+idx;
      url='/api/for_hubble/alteration.php';
      postData={"action":"state","fIdx":idx};
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
            var str='',strSucc='',strFail='',strRun='',strUn='';
            var num={succ:0,fail:0,run:0,un:0};
            var locale={};
            if(typeof(locale_messages.hubble.oprlog)) locale = locale_messages.hubble.oprlog;
            switch(type){
              case 'state':
                var state='';
                switch (data.content.state){
                  case 1: state='<span class="badge bg-blue">执行中</span>';break;
                  case 2: state='<span class="badge bg-green">成功</span>';break;
                  case 3: state='<span class="badge bg-red">失败</span>';break;
                  default: state='<span class="badge bg-red">未知</span>';break;
                }
                str+='<h5 class="col-sm-12 text-primary"><strong>任务状态: </strong>'+state+'</h5>';
                $.each(data.content.detail,function(k,v){
                  switch(v.state){
                    case 1: strRun='<label class="col-sm-2">'+ v.ip+'</label>';num.run++;break;
                    case 2: strSucc='<label class="col-sm-2">'+ v.ip+'</label>';num.succ++;break;
                    case 3: strFail='<label class="col-sm-2">'+ v.ip+'</label>';num.fail++;break;
                    default: strUn='<label class="col-sm-2">'+ v.ip+'</label>';num.un++;break;
                  }
                  str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                    '<h5 class="col-sm-12 text-primary"><strong class="text-danger">失败 : '+num.fail+'</strong></h5>'+'<div style="margin-left:30px;">'+strFail+'</div>';
                  str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                    '<h5 class="col-sm-12 text-primary"><strong class="text-warning">未知 : '+num.un+'</strong></h5>'+'<div style="margin-left:30px;">'+strUn+'</div>';
                  str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                    '<h5 class="col-sm-12 text-primary"><strong class="text-info">执行中 : '+num.run+'</strong></h5>'+'<div style="margin-left:30px;">'+strRun+'</div>';
                  str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                    '<h5 class="col-sm-12 text-primary"><strong class="text-success">成功 : '+num.succ+'</strong></h5>'+'<div style="margin-left:30px;">'+strSucc+'</div>';
                });
                text+=str;
                break;
              default:
                $.each(data.content,function(k,v){
                  if(k=='content'){
                    str+='<div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>' +
                      '<h5 class="col-sm-12 text-primary"><strong>内容</strong></h5>' +
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
