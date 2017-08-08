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

    <title><?php echo $mySiteTitle;?></title>

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
    <link href="gentelella/vendors/pagewalkthrough/dist/css/jquery.pagewalkthrough.css" rel="stylesheet" />

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
                    <a href="./" class="site_title"><i class="fa fa-cloud"></i> <span><?php echo $mySiteAlias;?></span></a>
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
                                    <li class="current-page"><a href="/opendcp_data.php">平台数据</a></li>
                                </ul>

                            </li>
                            <?php echo $navLeft;?>
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
                            <a href="javascript:;" class="user-profile dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
                                <?php echo $myCn;?>
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
                            <div class="count">168</div>
                            <h3>Cluster Count</h3>
                        </div>
                    </div>
                    <div class="animated flipInY col-lg-3 col-md-3 col-sm-6 col-xs-12">
                        <div class="tile-stats">
                            <div class="icon"><i class="fa fa-laptop"></i></div>
                            <div class="count">2860</div>
                            <h3>Machine Count</h3>
                        </div>
                    </div>
                    <div class="animated flipInY col-lg-3 col-md-3 col-sm-6 col-xs-12">
                        <div class="tile-stats">
                            <div class="icon"><i class="fa fa-cubes"></i></div>
                            <div class="count">5188</div>
                            <h3>Container Count</h3>
                        </div>
                    </div>
                    <div class="animated flipInY col-lg-3 col-md-3 col-sm-6 col-xs-12">
                        <div class="tile-stats">
                            <div class="icon"><i class="fa fa-tasks"></i></div>
                            <div class="count">10068</div>
                            <h3>Schedual Task</h3>
                        </div>
                    </div>
                </div>

                <div class="clearfix"></div>

                <div class="row">
                    <div class="col-md-12">
                        <div class="x_panel">
                            <div class="x_title">
                                <h2>Summary <small>Activity shares</small></h2>
                                <ul class="nav navbar-right panel_toolbox">
                                    <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                    </li>
                                    <li class="dropdown">
                                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>
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
                                <div class="row" style="border-bottom: 1px solid #E0E0E0; padding-bottom: 5px; margin-bottom: 5px;">
                                    <div class="col-md-7" style="overflow:hidden;">
                    <span class="sparkline_one" style="height: 160px; padding: 10px 25px;">
                      <canvas width="200" height="60" style="display: inline-block; vertical-align: top; width: 94px; height: 30px;"></canvas>
                    </span>
                                        <h4 style="margin:18px">Weekly dynamic capacity expansion progress</h4>
                                    </div>
                                    <div class="col-md-5">
                                        <div class="row" style="text-align: center;">
                                            <div class="col-md-4">
                                                <canvas id="canvas1i" height="110" width="110" style="margin: 5px 10px 10px 0"></canvas>
                                                <h4 style="margin:0">IAAS</h4>
                                            </div>
                                            <div class="col-md-4">
                                                <canvas id="canvas1i2" height="110" width="110" style="margin: 5px 10px 10px 0"></canvas>
                                                <h4 style="margin:0">PAAS</h4>
                                            </div>
                                            <div class="col-md-4">
                                                <canvas id="canvas1i3" height="110" width="110" style="margin: 5px 10px 10px 0"></canvas>
                                                <h4 style="margin:0">SAAS</h4>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>


                <div class="row">
                    <div class="col-md-4 col-sm-6 col-xs-12">
                        <div class="x_panel fixed_height_320">
                            <div class="x_title">
                                <h2>Versions</h2>
                                <ul class="nav navbar-right panel_toolbox">
                                    <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                    </li>
                                    <li class="dropdown">
                                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>
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
                                <div class="widget_summary">
                                    <div class="w_left w_25">
                                        <span>1.0.0</span>
                                    </div>
                                    <div class="w_center w_55">
                                        <div class="progress">
                                            <div class="progress-bar bg-green" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 66%;">
                                                <span class="sr-only">60% Complete</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="w_right w_20">
                                        <span></span>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>

                                <div class="widget_summary">
                                    <div class="w_left w_25">
                                        <span>0.1.1</span>
                                    </div>
                                    <div class="w_center w_55">
                                        <div class="progress">
                                            <div class="progress-bar bg-green" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 45%;">
                                                <span class="sr-only">60% Complete</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="w_right w_20">
                                        <span></span>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>
                                <div class="widget_summary">
                                    <div class="w_left w_25">
                                        <span>0.1.0</span>
                                    </div>
                                    <div class="w_center w_55">
                                        <div class="progress">
                                            <div class="progress-bar bg-green" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 25%;">
                                                <span class="sr-only">60% Complete</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="w_right w_20">
                                        <span></span>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>
                                <div class="widget_summary">
                                    <div class="w_left w_25">
                                        <span>0.0.2</span>
                                    </div>
                                    <div class="w_center w_55">
                                        <div class="progress">
                                            <div class="progress-bar bg-green" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 5%;">
                                                <span class="sr-only">60% Complete</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="w_right w_20">
                                        <span></span>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>
                                <div class="widget_summary">
                                    <div class="w_left w_25">
                                        <span>0.0.1</span>
                                    </div>
                                    <div class="w_center w_55">
                                        <div class="progress">
                                            <div class="progress-bar bg-green" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 2%;">
                                                <span class="sr-only">60% Complete</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="w_right w_20">
                                        <span></span>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>

                            </div>
                        </div>
                    </div>

                    <div class="col-md-4 col-sm-6 col-xs-12">
                        <div class="x_panel fixed_height_320">
                            <div class="x_title">
                                <h2>Cloud</h2>
                                <ul class="nav navbar-right panel_toolbox">
                                    <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                    </li>
                                    <li class="dropdown">
                                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>
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
                                <table class="" style="width:100%">
                                    <tr>
                                        <th style="width:37%;">
                                            <p>Top 5</p>
                                        </th>
                                        <th>
                                            <div class="col-lg-7 col-md-7 col-sm-7 col-xs-7">
                                                <p class="">Market</p>
                                            </div>
                                            <div class="col-lg-5 col-md-5 col-sm-5 col-xs-5">
                                                <p class=""> &nbsp&nbsp&nbsp;Per</p>
                                            </div>
                                        </th>
                                    </tr>
                                    <tr>
                                        <td>
                                            <canvas id="canvas1" height="130" width="130" style="margin: 15px 10px 10px 0"></canvas>
                                        </td>
                                        <td>
                                            <table class="tile_info">
                                                <tr>
                                                    <td>
                                                        <p><i class="fa fa-square blue"></i>AliCloud </p>
                                                    </td>
                                                    <td>31%</td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <p><i class="fa fa-square green"></i>Telecom </p>
                                                    </td>
                                                    <td>13%</td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <p><i class="fa fa-square purple"></i>Unicom </p>
                                                    </td>
                                                    <td>7%</td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <p><i class="fa fa-square aero"></i>Aws </p>
                                                    </td>
                                                    <td>4.3%</td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <p><i class="fa fa-square red"></i>Others </p>
                                                    </td>
                                                    <td>45%</td>
                                                </tr>
                                            </table>
                                        </td>
                                    </tr>
                                </table>
                            </div>
                        </div>
                    </div>

                    <div class="col-md-4 col-sm-6 col-xs-12">
                        <div class="x_panel fixed_height_320">
                            <div class="x_title">
                                <h2>Profile Settings <small>Sessions</small></h2>
                                <ul class="nav navbar-right panel_toolbox">
                                    <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                    </li>
                                    <li class="dropdown">
                                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>
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
                                <div class="dashboard-widget-content">
                                    <ul class="quick-list">
                                        <li><i class="fa fa-line-chart"></i><a href="#">Achievements</a></li>
                                        <li><i class="fa fa-thumbs-up"></i><a href="#">Favorites</a></li>
                                        <li><i class="fa fa-calendar-o"></i><a href="#">Activities</a></li>
                                        <li><i class="fa fa-cog"></i><a href="#">Settings</a></li>
                                        <li><i class="fa fa-area-chart"></i><a href="#">Logout</a></li>
                                    </ul>

                                    <div class="sidebar-widget">
                                        <h4>Profile Completion</h4>
                                        <canvas width="150" height="80" id="foo" class="" style="width: 160px; height: 100px;"></canvas>
                                        <div class="goal-wrapper">
                                            <span id="gauge-text" class="gauge-value pull-left">0</span>
                                            <span class="gauge-value pull-left">%</span>
                                            <span id="goal-text" class="goal-value pull-right">100%</span>
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
        <?php echo $myFooter;?>
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

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;" onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-2">
        <h1>多云对接 / 首次流程</h1>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 机型模板 &gt;&gt; 创建新模板, 并设置配额</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 机器管理 &gt;&gt; 创建一台机器</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;" onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-3">
        <h1>镜像市场 / 首次流程</h1>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 打包系统 &gt;&gt; 创建打包配置, 并构建镜像</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 镜像仓库 &gt;&gt; 查看新构建的镜像</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;" onclick="closePagewalkthrough()">点击不再显示</a></p>
    </div>

    <div id="walkthrough-4">
        <h1>服务编排 / 首次流程</h1>
        <p style="text-align: left;"><span style="color: orange;">第一步:</span> 集群管理 &gt;&gt; 创建集群</p>
        <p style="text-align: left;"><span style="color: orange;">第二步:</span> 远程命令 &gt;&gt; 创建命令、命令组</p>
        <p style="text-align: left;"><span style="color: orange;">第三步:</span> 任务管理 &gt;&gt; 创建任务模板(步骤=命令组)</p>
        <p style="text-align: left;"><span style="color: orange;">第四步:</span> 服务管理 &gt;&gt; 依次创建服务、服务池</p>

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;" onclick="closePagewalkthrough()">点击不再显示</a></p>
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

        <p style="font-size: 16px;">首次使用建议完整观看此引导, <a style="color: yellow;font-size: 16px;" href="javascript:;" onclick="closePagewalkthrough()">点击不再显示</a></p>
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
<!-- Chart.js -->
<script src="gentelella/vendors/Chart.js/dist/Chart.min.js"></script>
<!-- jQuery Sparklines -->
<script src="gentelella/vendors/jquery-sparkline/dist/jquery.sparkline.min.js"></script>
<!-- morris.js -->
<script src="gentelella/vendors/raphael/raphael.min.js"></script>
<script src="gentelella/vendors/morris.js/morris.min.js"></script>
<!-- gauge.js -->
<script src="gentelella/vendors/bernii/gauge.js/dist/gauge.min.js"></script>
<!-- Flot -->
<script src="gentelella/vendors/Flot/jquery.flot.js"></script>
<script src="gentelella/vendors/Flot/jquery.flot.pie.js"></script>
<script src="gentelella/vendors/Flot/jquery.flot.time.js"></script>
<script src="gentelella/vendors/Flot/jquery.flot.stack.js"></script>
<script src="gentelella/vendors/Flot/jquery.flot.resize.js"></script>
<!-- Flot plugins -->
<script src="gentelella/production/js/flot/jquery.flot.orderBars.js"></script>
<script src="gentelella/production/js/flot/date.js"></script>
<script src="gentelella/production/js/flot/jquery.flot.spline.js"></script>
<script src="gentelella/production/js/flot/curvedLines.js"></script>
<!-- bootstrap-progressbar -->
<script src="gentelella/vendors/bootstrap-progressbar/bootstrap-progressbar.min.js"></script>
<!-- iCheck -->
<script src="gentelella/vendors/iCheck/icheck.min.js"></script>
<!-- PNotify -->
<script src="gentelella/vendors/pnotify/dist/pnotify.js"></script>
<script src="gentelella/vendors/pnotify/dist/pnotify.buttons.js"></script>
<script src="gentelella/vendors/pnotify/dist/pnotify.nonblock.js"></script>
<!-- cookie -->
<script type="text/javascript" src="gentelella/vendors/jquery.cookie/jquery.cookie.js"></script>
<!-- Page walkthrough -->
<script type="text/javascript" src="gentelella/vendors/pagewalkthrough/dist/jquery.pagewalkthrough.js"></script>

<!-- Custom Theme Scripts -->
<script src="gentelella/build/js/custom.min.js"></script>
<!-- page level -->
<script src="js/pnotify.js"></script>
<script src="js/login.js"></script>
<script src="js/pagewalkthrough.js?_t=<?php echo date('U');?>"></script>
<!-- Custom Notification -->

<!-- jQuery Sparklines -->
<script>
    $(document).ready(function() {
        $(".sparkline_one").sparkline([2, 4, 3, 4, 5, 4, 5, 4, 3, 4, 5, 6, 4, 5, 6, 3, 5, 4, 5, 4, 5, 4, 3, 4, 5, 6, 7, 5, 4, 3, 5, 6], {
            type: 'bar',
            height: '125',
            barWidth: 13,
            colorMap: {
                '7': '#a1a1a1'
            },
            barSpacing: 2,
            barColor: '#26B99A'
        });

        $(".sparkline11").sparkline([2, 4, 3, 4, 5, 4, 5, 4, 3, 4, 6, 2, 4, 3, 4, 5, 4, 5, 4, 3], {
            type: 'bar',
            height: '40',
            barWidth: 8,
            colorMap: {
                '7': '#a1a1a1'
            },
            barSpacing: 2,
            barColor: '#26B99A'
        });

        $(".sparkline22").sparkline([2, 4, 3, 4, 7, 5, 4, 3, 5, 6, 2, 4, 3, 4, 5, 4, 5, 4, 3, 4, 6], {
            type: 'line',
            height: '40',
            width: '200',
            lineColor: '#26B99A',
            fillColor: '#ffffff',
            lineWidth: 3,
            spotColor: '#34495E',
            minSpotColor: '#34495E'
        });
    });
</script>
<!-- /jQuery Sparklines -->
<script>
    $(document).ready(function() {

        var canvasDoughnut,
            options = {
                legend: false,
                responsive: false
            };

        new Chart(document.getElementById("canvas1i"), {
            type: 'doughnut',
            tooltipFillColor: "rgba(51, 51, 51, 0.55)",
            data: {
                labels: [
                    "21Vianet",
                    "Unicom",
                    "Others",
                    "Telecom",
                    "Aliyun"
                ],
                datasets: [{
                    data: [4, 7, 45, 13, 31],
                    backgroundColor: [
                        "#BDC3C7",
                        "#9B59B6",
                        "#E74C3C",
                        "#26B99A",
                        "#3498DB"
                    ],
                    hoverBackgroundColor: [
                        "#CFD4D8",
                        "#B370CF",
                        "#E95E4F",
                        "#36CAAB",
                        "#49A9EA"
                    ]

                }]
            },
            options: options
        });

        new Chart(document.getElementById("canvas1i2"), {
            type: 'doughnut',
            tooltipFillColor: "rgba(51, 51, 51, 0.55)",
            data: {
                labels: [
                    "Salesforce",
                    "Amazon",
                    "Others",
                    "Microsoft",
                    "IBM"
                ],
                datasets: [{
                    data: [24, 16.8, 45.8, 10, 3.4],
                    backgroundColor: [
                        "#BDC3C7",
                        "#9B59B6",
                        "#E74C3C",
                        "#26B99A",
                        "#3498DB"
                    ],
                    hoverBackgroundColor: [
                        "#CFD4D8",
                        "#B370CF",
                        "#E95E4F",
                        "#36CAAB",
                        "#49A9EA"
                    ]

                }]
            },
            options: options
        });

        new Chart(document.getElementById("canvas1i3"), {
            type: 'doughnut',
            tooltipFillColor: "rgba(51, 51, 51, 0.55)",
            data: {
                labels: [
                    "Salesforce",
                    "Microsoft",
                    "Others",
                    "Adobe",
                    "SAP"
                ],
                datasets: [{
                    data: [11, 8, 70, 6, 5],
                    backgroundColor: [
                        "#BDC3C7",
                        "#9B59B6",
                        "#E74C3C",
                        "#26B99A",
                        "#3498DB"
                    ],
                    hoverBackgroundColor: [
                        "#CFD4D8",
                        "#B370CF",
                        "#E95E4F",
                        "#36CAAB",
                        "#49A9EA"
                    ]

                }]
            },
            options: options
        });

    });
</script>

<!-- Doughnut Chart -->
<script>
    $(document).ready(function() {
        var options = {
            legend: false,
            responsive: false
        };

        new Chart(document.getElementById("canvas1"), {
            type: 'doughnut',
            tooltipFillColor: "rgba(51, 51, 51, 0.55)",
            data: {
                labels: [
                    "21Vianet",
                    "Unicom",
                    "Others",
                    "Telecom",
                    "Aliyun"
                ],
                datasets: [{
                    data: [4, 7, 45, 13, 31],
                    backgroundColor: [
                        "#BDC3C7",
                        "#9B59B6",
                        "#E74C3C",
                        "#26B99A",
                        "#3498DB"
                    ],
                    hoverBackgroundColor: [
                        "#CFD4D8",
                        "#B370CF",
                        "#E95E4F",
                        "#36CAAB",
                        "#49A9EA"
                    ]
                }]
            },
            options: options
        });
    });
</script>
<!-- /Doughnut Chart -->

<!-- gauge.js -->
<script>
    var opts = {
        lines: 12,
        angle: 0,
        lineWidth: 0.4,
        pointer: {
            length: 0.75,
            strokeWidth: 0.042,
            color: '#1D212A'
        },
        limitMax: 'false',
        colorStart: '#1ABC9C',
        colorStop: '#1ABC9C',
        strokeColor: '#F0F3F3',
        generateGradient: true
    };
    var target = document.getElementById('foo'),
        gauge = new Gauge(target).setOptions(opts);

    gauge.maxValue = 100;
    gauge.animationSpeed = 32;
    gauge.set(80);
    gauge.setTextField(document.getElementById("gauge-text"));

</script>
<!-- /gauge.js -->
<!-- /Custom Notification -->
</body>
</html>
