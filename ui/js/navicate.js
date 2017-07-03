/**
 * Created by Administrator on 2017/6/21.
 */
//额度使用比例
var usedRate = 0.9;
var scrollTips = null;

$(document).ready(function() {
    //动态刷新时间getCurrentDate
    getCurrentDate();
    setInterval(getCurrentDate,1000);
    getColudList();
    //每15分钟更新一下列表
    setInterval(getColudList,1000*60*15);
});
var ScrollTextLeft = function(){
    var speed=50;
    if(scrollTips != null){
        clearInterval(scrollTips);
    }
    var scroll_begin = document.getElementById("scroll_begin");
    var scroll_end = document.getElementById("scroll_end");
    var scroll_div = document.getElementById("scroll_div");

    scroll_end.innerHTML=scroll_begin.innerHTML;
    function Marquee(){
        if(scroll_end.offsetWidth-scroll_div.scrollLeft<=0)
            scroll_div.scrollLeft-=scroll_begin.offsetWidth;
        else
            scroll_div.scrollLeft++;
    }
    scrollTips=setInterval(Marquee,speed);
    scroll_div.onmouseover = function(){
        clearInterval(scrollTips);
        scrollTips = null;
    }
    scroll_div.onmouseout = function(){
        scrollTips = setInterval(Marquee,speed);
    }
};
//获取当前时间
var getCurrentDate = function(){
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
}


var getColudList = function(){
    var url='/api/for_cloud/quota.php?action=list';
    var postData={};
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
                pageNotify('error','操作失败！','错误信息：'+data.msg);
            }
        },
        error: function (){
            pageNotify('error','操作失败！','错误信息：接口不可用');
        }
    });
}

var showQuota = function(datalist){
    var result = "";
    var tipresult = "";
    for(var i = 0; i < datalist.length; i++){
        var spent = Math.round(parseFloat(datalist[i].Spent)*10, 1)/10.0;
        var credit = Math.round(parseFloat(datalist[i].Credit)*10, 1)/10.0;
        var provider = datalist[i].Provider;
        if(spent > credit) spent = credit;
        result += '<li><a href="javascript:;"><i class="fa fa-cubes"></i> 总额度:'+ credit +
            ' <i class="fa fa-cube"></i> 使用额度:' + spent +
            ' <i class="fa fa-cloud"></i> 云厂商:'+ provider +'</a> </li>';
        if(spent > credit * usedRate){
            if(spent > credit) spent = credit;
            tipresult += '<span style = "padding-right:2em;">' +
                '<span style = "color: red"> <i class="fa fa-exclamation-circle"></i> 体验机额度使用提醒：</span>' +
                ' <i class="fa fa-cubes"></i> 总额度:' + credit +
                ' <i class="fa fa-cube"></i> 使用额度已达:' + spent +
                ' <i class="fa fa-cloud"></i> 云厂商:'+ provider +
                '<span style = "color: red"> <i class="fa fa-exclamation-circle"></i> 体验的机器即将删除！</span>'+
                '</span>';
        }
    }
    var thescrollContent = '<div id="scroll_begin" style ="display: inline;">' +
        tipresult+
        '</div><div id="scroll_end" style ="display: inline;"></div>';

    var totalWidth =  $("#scroll_div").width();
    $("#scroll_div").html(thescrollContent);
    // $("#scroll_begin").width(totalWidth);
    // $("#scroll_end").width(totalWidth);
    var scrollWidth = $("#scroll_begin").width();
    if(scrollWidth <= totalWidth){
        $("#scroll_div").width(scrollWidth);
    }
    if(tipresult != ""){
        setTimeout(ScrollTextLeft, 1000);
    }
    $("#auotaCloud").html(result);
}
