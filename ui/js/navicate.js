/**
 * Created by Administrator on 2017/6/21.
 */

$(document).ready(function() {
    //动态刷新时间
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
    var url='/api/for_cloud/quota.php?action=list';
    var postData={};
    NProgress.start();
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            //执行结果提示
            if(data.code==0){
                showQuota(data.content);
            }else{
                pageNotify('error','操作失败！','错误信息：'+data.msg, false);
            }
            NProgress.done();
        },
        error: function (){
            pageNotify('error','操作失败！','错误信息：接口不可用');
            NProgress.done();
        }
    });
}

var showQuota = function(datalist){
    var result = ""
    for(var i = 0; i < datalist.length; i++){
        var spent = datalist[i].Spent;
        var credit = datalist[i].Credit;
        var provider = datalist[i].Provider;
        result += '<a href="javascript:;"><i class="fa fa-cubes"></i> 总额度:'+ credit +
            ' <i class="fa fa-cube"></i> 使用额度:' + spent +
            ' <i class="fa fa-cloud"></i> 云厂商:'+ provider +'</a>';
    }
    $("#auotaCloud").html(result);
}
