$(document).ready(function() {
  var cookie='cookie_pagewalkthrough';
  var expire=3650;
  //$.cookie(cookie,0,{expires: expire});
  var count=$.cookie(cookie);
  var showCount=2;
  if(count>showCount) return false;

  // Set up tour
  $('body').pagewalkthrough({
    name: 'introduction',
    steps: [{
      popup: {
        content: '#walkthrough-1',
        type: 'modal',
      }
    }, {
      wrapper: '#walkthrough_1',
      popup: {
        content: '#walkthrough-2',
        type: 'tooltip',
        position: 'right'
      },
      onEnter: function(){
        $('#walkthrough_1').addClass('active');
        $('#walkthrough_1 > ul').css('display','block');
      },
      onLeave: function(){
        $('#walkthrough_1').removeClass('active');
        $('#walkthrough_1 > ul').css('display','none');
      },
    }, {
      wrapper: '#walkthrough_2',
      popup: {
        content: '#walkthrough-3',
        type: 'tooltip',
        position: 'right'
      },
      onEnter: function(){
        $('#walkthrough_2').addClass('active');
        $('#walkthrough_2 > ul').css('display','block');
      },
      onLeave: function(){
        $('#walkthrough_2').removeClass('active');
        $('#walkthrough_2 > ul').css('display','none');
      },
    }, {
      wrapper: '#walkthrough_4',
      popup: {
        content: '#walkthrough-5',
        type: 'tooltip',
        position: 'right'
      },
      onEnter: function(){
        $('#walkthrough_4').addClass('active');
        $('#walkthrough_4 > ul').css('display','block');
      },
      onLeave: function(){
        $('#walkthrough_4').removeClass('active');
        $('#walkthrough_4 > ul').css('display','none');
      },
    }, {
      wrapper: '#walkthrough_3',
      popup: {
        content: '#walkthrough-4',
        type: 'tooltip',
        position: 'right'
      },
      onEnter: function(){
        $('#walkthrough_3').addClass('active');
        $('#walkthrough_3 > ul').css('display','block');
      },
      onLeave: function(){
        $('#walkthrough_3').removeClass('active');
        $('#walkthrough_3 > ul').css('display','none');
      },
    }, {
      wrapper: '#walkthrough_90001',
      popup: {
        content: '#walkthrough-6',
        type: 'tooltip',
        position: 'right'
      },
      onEnter: function(){
        $('#walkthrough_90001').addClass('active');
        $('#walkthrough_90001 > ul').css('display','block');
      },
      onLeave: function(){
        $('#walkthrough_90001').removeClass('active');
        $('#walkthrough_90001 > ul').css('display','none');
      },
    }]
  });

  // Show the tour
  $('body').pagewalkthrough('show');

  if(count<=showCount){
    count++;
    $.cookie(cookie,count,{expires: expire});
  }else{
    $.cookie(cookie,1,{expires: expire});
  }
});

var closePagewalkthrough= function () {
  $.cookie('cookie_pagewalkthrough',11,{expires: 3650});
  pageNotify('success','已关闭引导流程');
  $.pagewalkthrough('close');
}

var openPagewalkthrouph=function(){
  $.cookie('cookie_pagewalkthrough',0,{expires: 3650});
  $('body').pagewalkthrough('show');
}