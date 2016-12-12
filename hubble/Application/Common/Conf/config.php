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
return array(
    'DEFAULT_MODULE'        =>  'Version',  // 默认模块
    'URL_MODEL'             =>  2,
    # -------------------------------------------------
    'URL_PATHINFO_DEPR'     =>  '/',           // PATHINFO模式下，各参数之间的分割符号
    # -------------------------------------------------
    'ACTION_SUFFIX'         =>  '',      // 操作方法后缀
    # -------------------------------------------------
    'LOG_RECORD'            =>  true,
    'LOG_EXCEPTION_RECORD'  =>  true,    // 是否记录异常信息日志
    'LOG_RECORD_LEVEL'      =>  'EMERG,ALERT,CRIT,ERR,WARN,SQL',
    # -------------------------------------------------
    # ----------  MySQL 相关  --------------------------
    'DB_TYPE'               =>  'mysql',       // 数据库类型
    'DB_HOST'               =>  '127.0.0.1',   // 服务器地址
    'DB_NAME'               =>  'hubble',       // 数据库名
    'DB_USER'               =>  'hubble',        // 用户名
    'DB_PWD'                =>  'hubble',    // 密码
    'DB_PORT'               =>  '3306',        // 端口
    'DB_PREFIX'             =>  'tbl_hubble_',// 数据库表前缀
    'DB_PARAMS'             =>  array(),       // 数据库连接参数
    'DB_FIELDS_CACHE'       =>  true,          // 启用字段缓存
    'DB_CHARSET'            =>  'utf8',        // 数据库编码默认采用utf8
    'DB_DEBUG'  			=>  true,
    # ------------ Cache 相关  ----------------------------
    'DATA_CACHE_TIME'       =>  0,             // 数据缓存有效期 0表示永久缓存
    'DATA_CACHE_COMPRESS'   =>  false,         // 数据缓存是否压缩缓存
    'DATA_CACHE_CHECK'      =>  false,         // 数据缓存是否校验缓存
    'DATA_CACHE_PREFIX'     =>  'HUBBLE_',     // 缓存前缀
    'DATA_CACHE_TYPE'       =>  'Redis',       // 数据缓存类型,
    # ------------ Redis 服务设置 -------------------------
    'REDIS_HOST'            => '127.0.0.1',
    'REDIS_PORT'            => 6379,
    # ----------------------------------------------------
    'URL_HTML_SUFFIX'       =>  '',  // URL伪静态后缀设

    # ------------- 自定义配置 --------------------------
    # ------ 目录配置 --------
    'HUBBLE_ROOT_DIR'       => RUNTIME_PATH.'Hubble/',
    'HUBBLE_NGINX_DIR'      => 'Nginx',
    'HUBBLE_TENGINE_DIR'    => 'Tengine',
    'HUBBLE_VARNISH_DIR'    => 'Varnish',
    'HUBBLE_HAPROXY_DIR'    => 'Haproxy',
    # ----- Cache Lock ------
    'HUBBLE_CACHE_LOCK_WAIT_TIME' => 5, // 等待锁的时间(秒),超时则返回失败
    'HUBBLE_CACHE_LOCK_TIME' => 2, //分钟
    # ----- Appkey ----------
    'HUBBLE_ROOT_APPKEY'    => '6741bc42-9e21-4763-977c-ac3a1fc0bdd8', # ROOT appkey
    # ----- Host ------------
    'HUBBLE_HOST'           => '127.0.0.1',   #本机地址
    'HUBBLE_PORT'           => 5454,            #本机端口
    # ----- address --------
    'HUBBLE_ANSIBLE_HTTP'   => 'http://127.0.0.1:8000', #下发通道ANSIBLE的端口
    'HUBBLE_SLB_HTTP'       => 'http://127.0.0.1:8888'
);
