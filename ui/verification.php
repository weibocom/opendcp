<?php
session_start();
$img = imagecreatetruecolor(100, 35);
$black = imagecolorallocate($img, 0x00, 0x00, 0x00);
$green = imagecolorallocate($img, 0x00, 0xFF, 0x00);
$white = imagecolorallocate($img, 0xFF, 0xFF, 0xFF);
imagefill($img,0,0,$white);
//生成随机的验证码  
$code = '';
for($i = 0; $i < 6; $i++) {
  $code .= rand(0, 9);
}
$_SESSION['verification_code'] = $code;  //存储验证码
imagestring($img, 30, 28, 10, $code, $black);
//加入噪点干扰  
for($i=0;$i<200;$i++) {
  imagesetpixel($img, rand(0, 100) , rand(0, 100) , $black);
  imagesetpixel($img, rand(0, 100) , rand(0, 100) , $green);
}
//输出验证码  
header("content-type: image/png");
imagepng($img);
imagedestroy($img);
?>