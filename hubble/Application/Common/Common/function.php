<?php
/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
/**
 * Created by PhpStorm.
 * User: reposkeeper
 * Date: 16/4/29
 * Time: 12:42
 */



// --------------- DEFINE  ---------------------

// ------------  日志相关   ---------------------
define("HUBBLE_LOG_FILE", 'hubble.log');
define("HUBBLE_LOG_REQUEST", 'hubble_request.log');
define("HUBBLE_ERROR", 3);
define("HUBBLE_WARN",  2);
define("HUBBLE_INFO",  1);
define("HUBBLE_DEBUG", 0);
define('HUBBLE_LOG_RECORD_LEVEL', 1);

// ------------ 核心接口规范 --------------------

define("HUBBLE_DB_ERR", 100);
define("HUBBLE_RET_SUCCESS", 0);
define("HUBBLE_RET_NULL", 200);

// ------------ 外部URL ------------------------

// ---------------- 公共函数 ---------------------

/*
 * http 封装了 cURL的操作, 用来请求一个http的链接
 * @param url      string  请求的http地址
 * @param params   array   请求的参数
 * @param method   string  请求的方法
 * @param header   array   需要添加的header
 * @param mutil    bool    是否传输文件
 */
function http($url, $params, $method = 'GET', $timeout = 5, $header = array()){
    $userAgent =
        "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; 360SE";

    $header[] = "Accept-Language: zh-cn,zh;q=0.5";
    $header[] = "Accept-Charset:GB2312,utf-8;q=0.7,*;q=0.7";

    $opts = array(
        CURLOPT_TIMEOUT        => $timeout,
        CURLOPT_RETURNTRANSFER => 1,
        CURLOPT_SSL_VERIFYPEER => false,
        CURLOPT_SSL_VERIFYHOST => false,
        CURLOPT_HTTPHEADER     => $header,
        CURLOPT_USERAGENT      => $userAgent,
        CURLOPT_RETURNTRANSFER => true,       // return transfer as a string
        CURLOPT_HEADER         => false,      // don't return headers
        CURLOPT_ENCODING       => "",         // handle all encodings
        CURLOPT_CONNECTTIMEOUT => 10,         // timeout on connect
    );
    /* 根据请求类型设置特定参数 */
    switch(strtoupper($method)){
        case 'GET':
            $opts[CURLOPT_URL] = $url . '?' . is_array($params)? http_build_query($params): $params;
            break;
        case 'POST':
            $params = is_array($params)? http_build_query($params): $params;
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_POST] = 1;
            $opts[CURLOPT_POSTFIELDS] = $params;
            break;
        case 'PUT':
        case 'DELETE':
            $opts[CURLOPT_URL] = $url . '?' . is_array($params)? http_build_query($params): $params;
            $opts[CURLOPT_CUSTOMREQUEST] = strtoupper($method);
            break;
        default:
            throw new Exception('不支持的请求方式！');
    }
    /* 初始化并执行curl请求 */
    $ch = curl_init();
    curl_setopt_array($ch, $opts);

    $transfer = curl_getinfo($ch);
    $data     = curl_exec($ch);
    $error    = curl_error($ch);
    $errno    = curl_errno($ch);

    $return =  [
        'code' => 0,
        'data' => $data,
        'error' => $error,
        'errno' => $errno,
        'http' => [
            'code' => $transfer['http_code'],
            'url' => $transfer['url'],
        ],
    ];

    curl_close($ch);

    if($error)
        $return['code'] = 1;

    return $return;
}

function http_for_slb($url, $params, $method = 'GET', $timeout = 5, $header = array()){
    $userAgent =
        "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; 360SE";

    $header[] = "Accept-Language: zh-cn,zh;q=0.5";
    $header[] = "Accept-Charset:GB2312,utf-8;q=0.7,*;q=0.7";

    $opts = array(
        CURLOPT_TIMEOUT        => $timeout,
        CURLOPT_RETURNTRANSFER => 1,
        CURLOPT_SSL_VERIFYPEER => false,
        CURLOPT_SSL_VERIFYHOST => false,
        CURLOPT_HTTPHEADER     => $header,
        CURLOPT_USERAGENT      => $userAgent,
        CURLOPT_RETURNTRANSFER => true,       // return transfer as a string
        CURLOPT_HEADER         => false,      // don't return headers
        CURLOPT_ENCODING       => "",         // handle all encodings
        CURLOPT_CONNECTTIMEOUT => 10,         // timeout on connect
    );
    /* 根据请求类型设置特定参数 */
    switch(strtoupper($method)){
        case 'GET':
            $opts[CURLOPT_URL] = $url . '?' . is_array($params)? http_build_query($params): $params;
            break;
        case 'POST':
            $params = is_array($params)? http_build_query($params): $params;
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_POST] = 1;
            $opts[CURLOPT_POSTFIELDS] = $params;
            break;
        case 'PUT':
            $opts[CURLOPT_URL] = $url . '?' . is_array($params)? http_build_query($params): $params;
            $opts[CURLOPT_CUSTOMREQUEST] = strtoupper($method);
            break;
        case 'DELETE':
            $params = is_array($params)? http_build_query($params): $params;
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_POST] = 1;
            $opts[CURLOPT_POSTFIELDS] = $params;
            $opts[CURLOPT_CUSTOMREQUEST] = strtoupper($method);
            break;
        default:
            throw new Exception('不支持的请求方式！');
    }
    /* 初始化并执行curl请求 */
    $ch = curl_init();
    curl_setopt_array($ch, $opts);

    $data     = curl_exec($ch);
    $transfer = curl_getinfo($ch);
    $error    = curl_error($ch);
    $errno    = curl_errno($ch);

    $return =  [
        'code' => 0,
        'data' => $data,
        'error' => $error,
        'errno' => $errno,
        'http' => [
            'code' => $transfer['http_code'],
            'url' => $transfer['url'],
        ],
    ];

    curl_close($ch);

    if($error)
        $return['code'] = 1;

    return $return;
}

// ----------------------- 标准返回格式 -----------------------
/*
 * 这是一个标准的返回格式
 */
function std_return($data = array(), $msg = 'success'){
    if(empty($data)) $data = (object)array();
    $ret =  [ 'code' => 0, 'msg'  => $msg, 'data' => $data ];
    if(REQUEST_METHOD != 'GET')
        base_log(HUBBLE_LOG_REQUEST, HUBBLE_INFO,
            "请求: ".html_entity_decode(I('server.REQUEST_URI'))." 返回: ". json_encode($ret),
            REQUEST_METHOD);

    return $ret;
}

/*
 * 这是一个标准的错误返回格式
 */
function std_error($msg = 'default error message', $code = 1){
    $ret = [ 'code' => $code, 'msg'  => $msg, 'data' => (object)array() ];
    if(REQUEST_METHOD != 'GET')
        base_log(HUBBLE_LOG_REQUEST, HUBBLE_INFO,
            "请求: ".html_entity_decode(I('server.REQUEST_URI'))." 返回: ". json_encode($ret),
            REQUEST_METHOD);
    return $ret;
}

// ----------------------- 检查认证 ---------------------------

function check_auth($appkey, $item){

    if(REQUEST_METHOD == 'GET' && strpos($item, 'appkey') === false) return true;
    if(empty($appkey)) return false;

    $app = new \Common\Dao\Secure\AppKey();
    return $app->checkPrivilege($appkey, str_replace("/index.php", "", $item));
}

function hubble_parse_param(){
    if(REQUEST_METHOD == 'GET')
        $input = I('get.');
    else
        $input = file_get_contents('php://input');
    hubble_request_log($input);

    if(REQUEST_METHOD == 'GET')
        return $input;
    else
        return json_decode($input, true);

}

function check_gid(){

    if(IS_GET) return true;

    if(empty(I('server.HTTP_X_CORRELATION_ID')))
        return false;

    return true;
}

// ----------------------- 日志相关 ---------------------------


function base_log($file, $level, $msg, $module = 'base'){

    $time = date("Y-m-d H:i:s");
    $gid = I('server.HTTP_X_CORRELATION_ID');
    switch($level){
        case 0:
            $level_str = 'DEBUG'; break;
        case 1:
            $level_str = 'INFO'; break;
        case 2:
            $level_str = 'WARN'; break;
        case 3:
            $level_str = 'ERROR'; break;
        default:
            return;
    }
    if($level >= HUBBLE_LOG_RECORD_LEVEL) {
        if ($level == 0 || IS_GET || empty($gid)){
            file_put_contents(LOG_PATH . $file, "$time $level_str ] $module $gid $msg \n", FILE_APPEND);
        } else {
            file_put_contents(LOG_PATH . $file, "$time $level_str ] $module $gid $msg \n", FILE_APPEND);
            $log_db = new \Common\Dao\Common\LogDb();
            $log_db->insert($gid, $module, "$time $level_str ] $module $gid $msg", $level_str);
        }
    }
}

/*
 * 应用的日志记录,不与框架的混在一块
 * @param level 日志等级
 * @param msg   日志内容
 */
function hubble_log($level, $msg, $middle = 'core'){
    
    base_log(HUBBLE_LOG_FILE, $level, $msg, $middle);
}

function hubble_request_log($param){

    // 记录日志
    switch(REQUEST_METHOD) {

        case 'GET':
        case 'DELETE':
        case 'PUT':
        case 'POST':
            base_log(HUBBLE_LOG_REQUEST, HUBBLE_INFO,
                '地址: [ '.I('server.REQUEST_URI') . ' ] 参数: ' . $param, REQUEST_METHOD);
            break;
        default:
            base_log(HUBBLE_LOG_REQUEST, HUBBLE_WARN, 'unknow request', REQUEST_METHOD);
    }
}

/*
 * 请求的记录日志
 */
function hubble_middle_layer(){

    // 记录日志
    switch(REQUEST_METHOD) {

        case 'GET':
        case 'DELETE':
            base_log(HUBBLE_LOG_REQUEST, HUBBLE_INFO, html_entity_decode(I('server.REQUEST_URI')), REQUEST_METHOD);
            break;
        case 'PUT':
        case 'POST':
            base_log(HUBBLE_LOG_REQUEST, HUBBLE_INFO,
                '地址: [ '.I('server.REQUEST_URI') . ' ] 参数: ' .
                html_entity_decode(urldecode(http_build_query(I('param.')))),
                REQUEST_METHOD);
            break;
    }

    if(REQUEST_METHOD != 'GET' && empty(I('server.HTTP_APPKEY')))
        return [false, 'no appkey'];

    if(!check_auth(I('server.HTTP_APPKEY'), I('server.DOCUMENT_URI')))
        return [false, 'this appkey do not have permission.'];

    if(!check_gid())
        return [false, 'check correlation_id failed'];

    return [true, 'success'];
}

// ---------------------- 性能Log的全局函数 -------------------------
function performance_str($msg, $time){

    return "操作:[ $msg ], 时间:[ $time ] \n";

}

// ----------------------- 生成 UUID -------------------------------
/*
 * 生成一个UUID
 * @param prefix 前缀
 * @return str
 */
function UUID($prefix = ""){

    $str   = md5(uniqid(mt_rand(), true));
    $uuid  = substr($str,0,8) . '-';
    $uuid .= substr($str,8,4) . '-';
    $uuid .= substr($str,12,4) . '-';
    $uuid .= substr($str,16,4) . '-';
    $uuid .= substr($str,20,12);
    return $prefix . $uuid;
}


function hubble_oprlog($module, $operation, $appkey, $user, $args = ''){

    $oprlog = new \Common\Dao\Secure\Oprlog();
    $ret = $oprlog->addItem($module, $operation, $appkey, $user, $args);
    if($ret === false)
        hubble_log(HUBBLE_WARN,
            '变更记录失败: ' . "[$module - $operation by $user with appkey $appkey, args=$args]"
        );
}

// 递归删除文件及文件夹
function rmdir_recursive($dir) {
    if(!is_dir($dir)) return;
    foreach(scandir($dir) as $file) {
        if ('.' === $file || '..' === $file) continue;
        if (is_dir("$dir/$file")) rmdir_recursive("$dir/$file");
        else unlink("$dir/$file");
    }
    rmdir($dir);
}
