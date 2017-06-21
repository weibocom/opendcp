/**
 * Created by Administrator on 2017/6/21.
 */

$(document).ready(function() {

    setInterval(getDate,1000);


});
//获取当前时间
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
    var currentTime = d.getFullYear()+'/'+month+'/'+day+' '+hour+':'+min+':'+sec;
    $("#currentTime").text(currentTime);
};


var getColudList = function(){
    var url = var url='/api/login.php';
    $.ajax({
        type: "POST",
        url: url,
        data: post,
        dataType: "json",
        success: function (data) {
            //执行结果提示
            if(data.code==0){
                switch(action) {
                    case 'login':
                        pageNotify('success','您已成功登录系统！');
                        break;
                    case 'logout':
                        pageNotify('success','您已成功登出系统！');
                        break;
                }
                setTimeout("window.parent.location.href='/'",1000);
            }else{
                pageNotify('error','操作失败！','错误信息：'+data.msg, false);
            }
        },
        error: function (){
            pageNotify('error','操作失败！','错误信息：接口不可用');
        }
    });


}


