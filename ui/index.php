<?php
/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */


require_once('include/config.inc.php');
require_once('include/function.php');
require_once('include/func_session.php');
require_once('include/navbar.php');
?>
<!DOCTYPE html>
<html lang="en">

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <!-- Meta, title, CSS, favicons, etc. -->
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title><?php echo $mySiteTitle; ?></title>

    <!-- Bootstrap -->
    <link href="gentelella/vendors/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
    <!-- Font Awesome -->
    <link href="gentelella/vendors/font-awesome/css/font-awesome.min.css" rel="stylesheet">
    <!-- iCheck -->
    <link href="gentelella/vendors/iCheck/skins/flat/green.css" rel="stylesheet">
    <!-- bootstrap-progressbar -->
    <link href="gentelella/vendors/bootstrap-progressbar/css/bootstrap-progressbar-3.3.4.min.css" rel="stylesheet">
    <!-- PNotify -->
    <link href="gentelella/vendors/pnotify/dist/pnotify.css" rel="stylesheet">
    <link href="gentelella/vendors/pnotify/dist/pnotify.buttons.css" rel="stylesheet">
    <link href="gentelella/vendors/pnotify/dist/pnotify.nonblock.css" rel="stylesheet">
    <!-- pagewalkthrough -->
    <link href="gentelella/vendors/pagewalkthrough/dist/css/jquery.pagewalkthrough.css" rel="stylesheet"/>

    <!-- Custom Theme Style -->
    <link href="gentelella/build/css/custom.min.css" rel="stylesheet">
    <link href="css/custom.css" rel="stylesheet">
</head>

<body class="nav-md">
<div class="container body">
    <div class="main_container">
        <div class="col-md-3 left_col">
            <div class="left_col scroll-view">
                <div class="navbar nav_title" style="border: 0;background-color: #FB5557;">
                    <a href="./" class="site_title"><i class="fa fa-cloud"></i> <span><?php echo $mySiteAlias; ?></span></a>
                </div>

                <div class="clearfix"></div>

                <!-- sidebar menu -->
                <div id="sidebar-menu" class="main_menu_side hidden-print main_menu">
                    <div class="menu_section" style="margin-bottom: 0px;">
                        <ul class="nav side-menu">
                            <li class="active">
                                <a><i class="fa fa-home"></i> Home </a>
                                <ul class="nav child_menu" style="display: block;">
                                    <!--<li class="current-page"><a href="/">Dashboard</a></li>-->
                                    <li class="current-page"><a href="/">Dashboard</a></li>
                                </ul>

                            </li>
                            <?php echo $navLeft; ?>
                        </ul>
                    </div>

                </div>
                <!-- /sidebar menu -->

                <!-- /menu footer buttons -->
                <div class="sidebar-footer hidden-small">
                </div>
                <!-- /menu footer buttons -->
            </div>
        </div>

        <!-- top navigation -->
        <div class="top_nav">
            <div class="nav_menu">
                <nav class="" role="navigation">
                    <div class="nav toggle">
                        <a id="menu_toggle"><i class="fa fa-bars"></i></a>
                    </div>

                    <ul class="nav navbar-nav navbar-right">
                        <li class="">
                            <a href="javascript:;" class="user-profile dropdown-toggle" data-toggle="dropdown"
                               aria-expanded="false">
                                <?php echo $myCn; ?>
                                <span class=" fa fa-angle-down"></span>
                            </a>
                            <ul class="dropdown-menu dropdown-usermenu pull-right">
                                <li><a href="user/log.php"><i class="fa fa-history"></i> 我的日志</a></li>
                                <li><a onclick="login('logout')"><i class="fa fa-sign-out"></i> 退出</a></li>
                            </ul>
                        </li>
                    </ul>
                </nav>
            </div>
        </div>
        <!-- /top navigation -->

        <!-- page content -->
        <div class="right_col" role="main">
            <div class="">

                <div class="row top_tiles">
                    <div class="animated flipInY col-lg-3 col-md-3 col-sm-6 col-xs-12">
                        <div class="tile-stats">
                            <div class="icon"><i class="fa fa-database"></i></div>
                            <div id="clusterCount" class="count">0</div>
                            <h3>Cluster Count</h3>
                        </div>
                    </div>
                    <div class="animated flipInY col-lg-3 col-md-3 col-sm-6 col-xs-12">
                        <div class="tile-stats">
                            <div class="icon"><i class="fa fa-laptop"></i></div>
                            <div id="machineCount" class="count">0</div>
                            <h3>Machine Count</h3>
                        </div>
                    </div>
                    <div class="animated flipInY col-lg-3 col-md-3 col-sm-6 col-xs-12">
                        <div class="tile-stats">
                            <div class="icon"><i class="fa fa-cubes"></i></div>
                            <div id="poolCount" class="count">0</div>
                            <h3>Pool Count</h3>
                        </div>
                    </div>
                    <div class="animated flipInY col-lg-3 col-md-3 col-sm-6 col-xs-12">
                        <div class="tile-stats">
                            <div class="icon"><i class="fa fa-tasks"></i></div>
                            <div id="taskCount" class="count">0</div>
                            <h3>Schedual Task</h3>
                        </div>
                    </div>
                </div>

                <div class="clearfix"></div>
                <div class="row">
                    <div class="col-md-6 col-sm-6 col-xs-12">
                        <div class="x_panel">
                            <div class="x_title">
                                <h2>弹性扩容
                                    <small>机器数量</small>
                                </h2>
                                <ul class="nav navbar-right panel_toolbox">
                                    <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                    </li>
                                    <li class="dropdown">
                                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button"
                                           aria-expanded="false"><i class="fa fa-wrench"></i></a>
                                        <ul class="dropdown-menu" role="menu">
                                            <li><a href="#">Settings</a></li>
                                        </ul>
                                    </li>
                                    <li><a class="close-link"><i class="fa fa-close"></i></a>
                                    </li>
                                </ul>
                                <div class="clearfix"></div>
                            </div>
                            <div class="x_content">
                                <div class="col-sm-4 navbar-right" style="padding-right:0px;">
                                    <div class="input-group">
                                        <input type="name" id="time_nume" class="form-control" placeholder="输入时间"
                                               value="1">
                                        <span class="input-group-btn dropdown">
                                            <button id="time_unit" data-toggle="dropdown"
                                                    class="btn btn-default dropdown-toggle" type="button"
                                                    aria-expanded="false">
                                                <a onclick="changeTime(0)">小时 </a><span class="caret"></span>
                                            </button>
                                            <ul role="menu" class="dropdown-menu col-sm-3 navbar-right">
                                              <li><a onclick="changeTime(0)">小时</a></li>
                                              <li><a onclick="changeTime(1)">天</a></li>
                                              <li><a onclick="changeTime(2)">周</a></li>
                                              <li><a onclick="changeTime(3)">月</a></li>
                                              <li><a onclick="changeTime(4)">年</a> </li>
                                            </ul>
                                         </span>
                                    </div>
                                </div>
                            </div>
                            <div class="x_content" style="margin-top: -30px;">
                                <div id="container"
                                     style="height:310px; margin-right: -50px; margin-bottom: -35px;"></div>
                            </div>
                        </div>
                    </div>

                    <div class="col-md-6 col-sm-6 col-xs-12">
                        <div class="x_panel">
                            <div class="x_title">
                                <h2>虚拟化
                                    <small>机器详情</small>
                                </h2>
                                <ul class="nav navbar-right panel_toolbox">
                                    <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                    </li>
                                    <li class="dropdown">
                                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button"
                                           aria-expanded="false"><i class="fa fa-wrench"></i></a>
                                        <ul class="dropdown-menu" role="menu">
                                            <li><a href="#">Settings 1</a>
                                            </li>
                                            <li><a href="#">Settings 2</a>
                                            </li>
                                        </ul>
                                    </li>
                                    <li><a class="close-link"><i class="fa fa-close"></i></a>
                                    </li>
                                </ul>
                                <div class="clearfix"></div>
                            </div>
                            <div class="x_content">
                                <div class="col-sm-4 navbar-right" style="padding-right:0px;">
                                    <div class="input-group">
                                        <input type="name" id="time_open_nume" class="form-control" placeholder="输入时间"
                                               value="1">
                                        <span class="input-group-btn dropdown">
                                            <button id="time_open_unit" data-toggle="dropdown"
                                                    class="btn btn-default dropdown-toggle" type="button"
                                                    aria-expanded="false">
                                                <a onclick="changeOpenTime(0)">小时 </a><span class="caret"></span>
                                            </button>
                                            <ul role="menu" class="dropdown-menu col-sm-3 navbar-right">
                                              <li><a onclick="changeOpenTime(0)">小时</a></li>
                                              <li><a onclick="changeOpenTime(1)">天</a></li>
                                              <li><a onclick="changeOpenTime(2)">周</a></li>
                                              <li><a onclick="changeOpenTime(3)">月</a></li>
                                              <li><a onclick="changeOpenTime(4)">年</a> </li>
                                            </ul>
                                         </span>
                                    </div>
                                </div>
                            </div>
                            <div class="x_content" style="margin-top: -30px">
                                <div id="container_stack"
                                     style="height:310px; margin-right: -50px;margin-bottom: -35px;"
                                ">
                            </div>
                        </div>
                    </div>
                </div>

            </div>
            <div class="clearfix"></div>
            <div class="row">
                <div class="col-md-12 col-sm-12 col-xs-12">
                    <div class="x_panel">
                        <div class="x_title">
                            <h2>任务信息</h2>
                            <ul class="nav navbar-right panel_toolbox">
                                <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                </li>
                                <li class="dropdown">
                                    <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button"
                                       aria-expanded="false"><i class="fa fa-wrench"></i></a>
                                    <ul class="dropdown-menu" role="menu">
                                        <li><a href="#">Settings 1</a>
                                        </li>
                                        <li><a href="#">Settings 2</a>
                                        </li>
                                    </ul>
                                </li>
                                <li><a class="close-link"><i class="fa fa-close"></i></a>
                                </li>
                            </ul>
                            <div class="clearfix"></div>
                        </div>
                        <div class="x_content">
                            <table id="datatable" class="table table-striped table-bordered">
                                <thead id="table-head">
                                <tr></tr>
                                </thead>
                                <tbody id="table-body">
                                </tbody>
                            </table>
                            <div class="row">
                                <div class="col-md-5 col-sm-5">
                                    <div class="dataTables_info" id="table-pageinfo" role="status" aria-live="polite">
                                        Showing 1 to 0 of 0 entries
                                    </div>
                                </div>
                                <div class="col-md-7 col-sm-7">
                                    <div class="dataTables_paginate paging_bootstrap_full_number"
                                         id="sample_1_paginate">
                                        <ul class="pagination"
                                            style="visibility: visible;margin-top: 0px;margin-bottom: 0px;"
                                            id="table-paginate">
                                            <li><a href="javascript:;" onclick="getTask(1)"><i
                                                            class="fa fa-angle-left"></i></a></li>
                                            <li class="active">
                                                <a href="javascript:;" onclick="getTask(1)">1</a>
                                            </li>
                                            <li class="next">
                                                <a href="javascript:;" title="Next" onclick="getTask(1)"><i
                                                            class="fa fa-angle-right"></i></a>
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="clearfix"></div>
    </div>
    <!-- /page content -->

    <!-- footer content -->
    <?php echo $myFooter; ?>
    <!-- /footer content -->
</div>
</div>

<div id="custom_notifications" class="custom-notifications dsp_none">
    <ul class="list-unstyled notifications clearfix" data-tabbed_notifications="notif-group">
    </ul>
    <div class="clearfix"></div>
    <div id="notif-group" class="tabbed_notifications"></div>
</div>

<div id="walkthrough-content">
    <div id="walkthrough-1">
        <h1>欢迎使用<br>DCP混合云开源平台</h1>
        <p>首次使用流程</p>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 多云对接 &gt;&gt; 创建机型模板, 并设置配额</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 镜像市场 &gt;&gt; 创建打包配置, 并构建镜像</p>
        <p style="text-align: left;"><span style="color: orange;">第三步:</span> 服务发现 &gt;&gt; 配置服务注册类型</p>
        <p style="text-align: left;"><span style="color: orange;">第四步:</span> 服务编排 &gt;&gt; 创建集群、服务、服务池</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;"
                                                      onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-2">
        <h1>多云对接 / 首次流程</h1>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 机型模板 &gt;&gt; 创建新模板, 并设置配额</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 机器管理 &gt;&gt; 创建一台机器</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;"
                                                      onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-3">
        <h1>镜像市场 / 首次流程</h1>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 打包系统 &gt;&gt; 创建打包配置, 并构建镜像</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 镜像仓库 &gt;&gt; 查看新构建的镜像</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;"
                                                      onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-4">
        <h1>服务编排 / 首次流程</h1>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 集群管理 &gt;&gt; 创建集群</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 远程命令 &gt;&gt; 创建命令、命令组</p>
        <p style="text-align: left;"><span style="color: orange;">第三步:</span> 任务管理 &gt;&gt; 创建任务模板(步骤=命令组)</p>
        <p style="text-align: left;"><span style="color: orange;">第四步:</span> 服务管理 &gt;&gt; 依次创建服务、服务池</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;"
                                                      onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-5">
        <h1>服务发现 / 首次流程</h1>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 确认服务注册类型 &gt;&gt; Nginx、SLB</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 若使用SLB &gt;&gt; 跳过以下三、四、五、六步</p>
        <p style="text-align: left;"><span style="color: orange;">第三步:</span> 若使用Nginx &gt;&gt; 创建Nginx分组、单元</p>
        <p style="text-align: left;"><span style="color: orange;">第四步:</span> 若使用Nginx &gt;&gt; 导入Nginx节点(IP)</p>
        <p style="text-align: left;"><span style="color: orange;">第五步:</span> 若使用Nginx &gt;&gt; 创建或导入Upstream配置</p>
        <p style="text-align: left;"><span style="color: orange;">第六步:</span> 若使用Nginx &gt;&gt; 创建或导入Nginx主配置</p>
        <p style="text-align: left;"><span style="color: orange;">第七步:</span> 服务注册 &gt;&gt; 创建服务注册</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;"
                                                      onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-6">
        <h1>系统管理</h1>
        <p style="color: orange;">仅管理员有权限</p>
        <p><a style="color: yellow;font-size: 16px;" href="javascript:;" onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>
</div>

<!-- jQuery -->
<script src="gentelella/vendors/jquery/dist/jquery.min.js"></script>
<!-- Bootstrap -->
<script src="gentelella/vendors/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- FastClick -->
<script src="gentelella/vendors/fastclick/lib/fastclick.js"></script>
<!-- NProgress -->
<script src="gentelella/vendors/nprogress/nprogress.js"></script>
<!-- bootstrap-progressbar -->
<script src="gentelella/vendors/bootstrap-progressbar/bootstrap-progressbar.min.js"></script>
<!-- iCheck -->
<script src="gentelella/vendors/iCheck/icheck.min.js"></script>
<!-- bootstrap-daterangepicker -->
<script src="gentelella/production/js/moment/moment.min.js"></script>
<script src="gentelella/production/js/datepicker/daterangepicker.js"></script>
<!-- bootstrap-wysiwyg -->
<script src="gentelella/vendors/bootstrap-wysiwyg/js/bootstrap-wysiwyg.min.js"></script>
<script src="gentelella/vendors/jquery.hotkeys/jquery.hotkeys.js"></script>
<script src="gentelella/vendors/google-code-prettify/src/prettify.js"></script>
<!-- jQuery Tags Input -->
<script src="gentelella/vendors/jquery.tagsinput/src/jquery.tagsinput.js"></script>
<!-- Switchery -->
<script src="gentelella/vendors/switchery/dist/switchery.min.js"></script>
<!-- Select2 -->
<script src="gentelella/vendors/select2/dist/js/select2.full.min.js"></script>
<!-- Parsley -->
<script src="gentelella/vendors/parsleyjs/dist/parsley.min.js"></script>
<!-- Autosize -->
<script src="gentelella/vendors/autosize/dist/autosize.min.js"></script>
<!-- jQuery autocomplete -->
<script src="gentelella/vendors/devbridge-autocomplete/dist/jquery.autocomplete.min.js"></script>
<!-- starrr -->
<script src="gentelella/vendors/starrr/dist/starrr.js"></script>
<!-- PNotify -->
<script src="gentelella/vendors/pnotify/dist/pnotify.js"></script>
<script src="gentelella/vendors/pnotify/dist/pnotify.buttons.js"></script>
<script src="gentelella/vendors/pnotify/dist/pnotify.nonblock.js"></script>
<!-- reveal -->
<script src="gentelella/vendors/reveal330/lib/js/head.min.js"></script>
<script src="gentelella/vendors/reveal330/js/reveal.js"></script>

<!-- ECharts -->
<script src="gentelella/vendors/echarts/dist/echarts.min.js"></script>
<script src="gentelella/vendors/echarts/map/js/world.js"></script>

<!-- cookie -->
<script type="text/javascript" src="gentelella/vendors/jquery.cookie/jquery.cookie.js"></script>
<!-- Page walkthrough -->
<script type="text/javascript" src="gentelella/vendors/pagewalkthrough/dist/jquery.pagewalkthrough.js"></script>

<!-- Custom Theme Scripts -->
<script src="gentelella/build/js/custom.min.js"></script>
<!-- page level -->
<script src="js/pnotify.js"></script>
<script src="js/login.js"></script>
<script src="js/pagewalkthrough.js?_t=<?php echo date('U'); ?>"></script>
<script src="js/opendcpData.js?_t=<?php echo date('U'); ?>"></script>
<!-- Custom Notification -->

<!-- /Custom Notification -->
</body>
</html>
