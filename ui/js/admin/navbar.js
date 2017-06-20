cache={
  page:1
}
var list = function(page) {
  NProgress.start();
  var url='/api/admin/navbar.php';
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
      for (var i = 0; i < data.content.length; i++) {
        var v = data.content[i];
        var tr = $('<tr></tr>');
        td = '<td>' + v.i + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_id + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_fid + '</td>';
        tr.append(td);
        switch(v.nb_status){
          case 0:
            td = '<td><div><label class="tooltips" title="点击停用" onclick="return false;"><input type="checkbox" class="js-switch" checked onchange="switchs(\'off\',\'' + v.nb_id + '\')"/> 启用</label></div></td>';
            //td = '<td><a class="btn btn-round btn-warning btn-xs tooltips" onclick="switchs(\'on\',\'' + v.nb_id + '\')" title="点击停用">启用</a></td>';
            tr.append(td);
            break;
          case 1:
            td = '<td><div><label class="tooltips" title="点击启用" onclick="return false;"><input type="checkbox" class="js-switch" onchange="switchs(\'on\',\'' + v.nb_id + '\')"/> 停用</label></div></td>';
            //td = '<td><a class="btn btn-round btn-warning btn-xs tooltips" onclick="switchs(\'on\',\'' + v.nb_id + '\')" title="点击启用">停用</a></td>';
            tr.append(td);
            break;
          case 2:
            td = '<td><span class="btn btn-round btn-danger btn-sm btn-xs">Hidden</span></td>';
            tr.append(td);
            break;
          default:
            td = '<td>' + v.nb_status + '</td>';
            tr.append(td);
            break;
        }
        td = '<td>' + v.nb_name + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_href + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_target + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_desc + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_sort + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_icon + '</td>';
        tr.append(td);
        td = '<td>' + v.nb_new + '</td>';
        tr.append(td);
        var btnEdit = '<a class="btn blue tooltips" data-toggle="modal" data-target="#myModal" href="edit_nav.php?action=edit&index=' + v.nb_id + '" style="padding-left: 0px;" title="修改"><i class="fa fa-edit"></i></a>';
        var btnDel = '<a class="btn red tooltips" data-toggle="modal" data-target="#myModal" href="edit_nav.php?action=del&index=' + v.nb_id + '" style="padding-left: 0px;padding-right: 0px;" title="删除"><i class="fa fa-trash-o"></i></a>';
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
  var url='/api/admin/navbar.php';
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
  $.ajax({
    type: "POST",
    url: url,
    data: {"action":action,"data":JSON.stringify(postData)},
    dataType: "json",
    success: function (data) {
      //执行结果提示
      if(data.code==0){
        pageNotify('success','操作成功！');
      }else{
        pageNotify('error','操作失败！','错误信息：'+data.msg);
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
      NProgress.done();
    },
    error: function (){
      pageNotify('error','操作失败！','错误信息：接口不可用');
      NProgress.done();
    }
  });
}

//开关
var switchs=function(action,index){
  NProgress.start();
  var url='/api/admin/navbar.php';
  var postData={nb_id: index};
  $.ajax({
    type: "POST",
    url: url,
    data: {"action":action,"data":JSON.stringify(postData)},
    dataType: "json",
    success: function (data) {
      //执行结果提示
      if(data.code==0){
        pageNotify('success','操作成功！');
      }else{
        pageNotify('error','操作失败！','错误信息：'+data.msg);
      }
      //重载列表
      list();
      NProgress.done();
    },
    error: function (){
      pageNotify('error','操作失败！','错误信息：接口不可用');
      NProgress.done();
    }
  });
}