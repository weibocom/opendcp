<?php
namespace Version\Controller;
use Think\Controller;
class IndexController extends Controller {
    public function index(){
        echo "Welcome to Hubble";
    }

    public function test(){
        $ret = file_get_contents('php://input');


        echo json_encode($ret);
        echo json_encode(I('server.'));
    }
}