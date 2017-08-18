var switchery=function(){
  if($(".js-switch")[0]){
    $.each($(".js-switch"),function(){
      if($(this).attr('is_switchery')!="true"){
        $(this).attr('is_switchery','true');
        new Switchery(this,{color:"#26B99A"});
      }
    });
  }
}