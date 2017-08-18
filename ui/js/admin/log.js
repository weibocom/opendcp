cache={
  page:1
}
var list = function(page) {
  NProgress.start()
  var url='/api/admin/log.php';
  if (!page) {
    page = cache.page;
  }else{
    cache.page = page;
  }
  var filter=$('#filter').val();
  if(!filter){
    filter='';
  }
  var url=url+'?action=list&page=' + page+'&filter='+filter;
  $.ajax({
    type: "POST",
    url: url,
    data: '',
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
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        td = '<td>' + v.i + '</td>';
        tr.append(td);
        td = '<td>' + v.time + '</td>';
        tr.append(td);
        td = '<td>' + v.user + '</td>';
        tr.append(td);
        td = '<td>' + v.module + '</td>';
        tr.append(td);
        td = '<td>' + v.action + '</td>';
        tr.append(td);
        td = '<td>' + v.desc + '</td>';
        tr.append(td);
        var btnView = '<a class="text-primary tooltips" data-original-title="查看详情" data-placement="right" data-toggle="modal" data-target="#myViewModal" onclick="info(' + v.id + ')"><i class="fa fa-history"></i></a>';
        td = '<td><div class="btn-group btn-group-xs btn-group-solid">' + btnView + '</div></td>';
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

//信息
var info=function(index){
  NProgress.start();
  if(index!=''){
    $.ajax({
      type: "POST",
      url: '/api/admin/log.php',
      data: {"action":"info","index":index},
      dataType: "json",
      success: function (data) {
        //执行结果提示
        if(data.code==0){
          if(data.content.length>0){
            //pageNotify('success','加载成功！');
            var id = data.content[0].id;
            var module = data.content[0].module;
            var code = data.content[0].code;
            if(code==''){
              code='nothing';
            }
            var text = '<textarea rows="16" class="form-control" name="t_code" style="color:#FFFFFF;background-color:#000000;font-family:\'Lucida Console\';">'+code+'</textarea>';
            $('#myViewModalLabel').html('查看CODE - '+module+' - '+id);
            $('#myViewModalBody').html(text);
          }else{
            pageNotify('warning','数据为空！');
          }
        }else{
          pageNotify('error','加载失败！','错误信息：'+data.msg);
        }
        NProgress.done();
      },
      error: function (){
        pageNotify('error','加载失败！','错误信息：接口不可用');
        NProgress.done();
      }
    });
  }else{
    pageNotify('warning','错误操作！','错误信息：参数错误');
    NProgress.done();
  }
}