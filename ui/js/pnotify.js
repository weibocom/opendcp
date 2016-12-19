var pageNotify = function(type, title, msg, hide, delay){
  if(!msg) msg = false;
  if(hide!==false) hide = true;
  if(!delay) delay=3000;
  new PNotify({
    title: title, //标题
    text: msg,  //正文
    type: type, //success,info,error,warning
    hide: hide, //是否自动关闭
    delay: delay, //延迟关闭时间
    styling: 'bootstrap3'
  });
}
