//Login
var login=function(action){
  var post={};
  var url='/api/login.php';
  switch(action){
    case 'login':
      var type=$('#authtype').val();
      var user=$('#username').val();
      var pass=$('#password').val();
      post={'action':action,'data':JSON.stringify({'type':type,'user':user,'pass':pass})};
      break;
    case 'logout':
      post={'action':action};
      break;
    default:
      pageNotify('error','非法操作','错误信息: 非法请求');
      return false;
  }
  $.ajax({
    type: "POST",
    url: url,
    data: post,
    dataType: "json",
    success: function (data) {
      //执行结果提示
      if(data.code==0){
        pageNotify('success','您已成功登录系统！');
        setTimeout("window.parent.location.href='/'",1000);
      }else{
        pageNotify('error','操作失败！','错误信息：'+data.msg);
      }
    },
    error: function (){
      pageNotify('error','操作失败！','错误信息：接口不可用');
    }
  });
}
