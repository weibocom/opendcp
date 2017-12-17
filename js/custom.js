//首页加载过渡动画
$(window).load(function() {
  $('.preloader').fadeOut(1000);
});
//展示回到顶部
$(window).scroll(function() {
  if ($(".navbar").offset().top > 50) {
    $(".navbar-fixed-top").addClass("top-nav-collapse");
  } else {
    $(".navbar-fixed-top").removeClass("top-nav-collapse");
  }
});
$(document).ready(function() {
  //单击链接后隐藏移动菜单
  $('.navbar-collapse a').click(function() {
    $(".navbar-collapse").collapse('hide');
  });
  //Owl Carousel
  $(document).ready(function() {
    $("#owl-speakers").owlCarousel({
      autoPlay: 6000,
      items: 4,
      itemsDesktop: [1199, 2],
      itemsDesktopSmall: [979, 1],
      itemsTablet: [768, 1],
      itemsTabletSmall: [985, 2],
      itemsMobile: [479, 1]
    });
  });
  //回到顶部
  $(window).scroll(function() {
    //console.log("滚动"+$('#video').offset().top);
    if ($(this).scrollTop() > 200) {
      $('.go-top').fadeIn(200);
    } else {
      $('.go-top').fadeOut(200);
    }
  });
  //动画回到顶部
  $('.go-top').click(function(event) {
    event.preventDefault();
    $('html, body').animate({
        scrollTop: 0
      },
      1000);
  });
  //锚点动画
  $('.anchor').click(function(event) {
    event.preventDefault();
    $('html, body').animate({
      scrollTop: $('#'+$(this).attr("name")).offset().top
    }, 1000);
  });
  //wow
  new WOW({
    mobile: true
  }).init();
  
});
