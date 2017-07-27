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


require_once('../include/config.inc.php');
require_once('../include/function.php');
require_once('../include/func_session.php');
require_once('../include/navbar.php');
$myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?intval($_GET['idx']):0;
$arrIdx=array(
  'prev' => ($myIdx>1)?$myIdx-1:$myIdx,
  'next' => $myIdx+1,
);
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
  <link href="../gentelella/vendors/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
  <!-- Font Awesome -->
  <link href="../gentelella/vendors/font-awesome/css/font-awesome.min.css" rel="stylesheet">
  <!-- iCheck -->
  <link href="../gentelella/vendors/iCheck/skins/flat/green.css" rel="stylesheet">
  <!-- bootstrap-wysiwyg -->
  <link href="../gentelella/vendors/google-code-prettify/bin/prettify.min.css" rel="stylesheet">
  <!-- bootstrap-progressbar -->
  <link href="../gentelella/vendors/bootstrap-progressbar/css/bootstrap-progressbar-3.3.4.min.css" rel="stylesheet">
  <!-- Select2 -->
  <link href="../gentelella/vendors/select2/dist/css/select2.min.css" rel="stylesheet">
  <!-- Switchery -->
  <link href="../gentelella/vendors/switchery/dist/switchery.min.css" rel="stylesheet">
  <!-- starrr -->
  <link href="../gentelella/vendors/starrr/dist/starrr.css" rel="stylesheet">
  <!-- PNotify -->
  <link href="../gentelella/vendors/pnotify/dist/pnotify.css" rel="stylesheet">
  <link href="../gentelella/vendors/pnotify/dist/pnotify.buttons.css" rel="stylesheet">
  <link href="../gentelella/vendors/pnotify/dist/pnotify.nonblock.css" rel="stylesheet">
  <!-- reveal -->
  <link href="../gentelella/vendors/reveal330/css/reveal.css" rel="stylesheet">
  <link href="../gentelella/vendors/reveal330/css/theme/solarized.css" rel="stylesheet" id="theme">
  <link href="../gentelella/vendors/reveal330/lib/css/zenburn.css" rel="stylesheet">

  <!-- Custom Theme Style -->
  <link href="../gentelella/build/css/custom.min.css" rel="stylesheet">
  <link href="../css/custom.css" rel="stylesheet">
</head>

<body class="nav-md">
<div class="container body">
  <div class="main_container">
    <div class="col-md-3 left_col">
      <div class="left_col scroll-view">
        <div class="navbar nav_title" style="border: 0;background-color: #FB5557;">
          <a href="../" class="site_title"><i class="fa fa-cloud"></i> <span><?php echo $mySiteAlias;?></span></a>
        </div>

        <div class="clearfix"></div>

        <!-- sidebar menu -->
        <div id="sidebar-menu" class="main_menu_side hidden-print main_menu">
          <div class="menu_section" style="margin-bottom: 0px;">
            <ul class="nav side-menu">
              <li>
                <a><i class="fa fa-home"></i> Home <span class="fa fa-chevron-down"></span></a>
                <ul class="nav child_menu">
                    <!--<li class="current-page"><a href="/">Dashboard</a></li>-->
                    <li class="current-page"><a href="/">Dashboard</a></li>
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
                <li><a href="../user/log.php"><i class="fa fa-history"></i> 我的日志</a></li>
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
        <div class="page-title">
          <div class="title_left">
            <h3><?php echo $pageName;?> <small><?php echo $pageDesc;?></small></h3>
          </div>
          <div class="pull-right">
            <h3><a class="text-primary tooltips" title="查看帮助" data-toggle="modal" data-target="#myRevealModal" onclick="showHelp()"><i class="fa fa-question-circle"></i></a></h3>
          </div>
        </div>

        <div class="clearfix"></div>

        <div>
          <div class="x_panel">
            <div class="x_title">
              <h2 class="text-primary"><i class="fa fa-bank"></i> 任务信息 <small>任务参数和属性</small></h2>
              <ul class="nav navbar-right panel_toolbox">
                <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                </li>
                <li><a class="close-link"><i class="fa fa-close"></i></a>
                </li>
              </ul>
              <div class="clearfix"></div>
            </div>
            <div class="x_content" id="task_info">
              <div class="col-sm-12">
                <a class="btn btn-primary btn-xs" data-toggle="modal" data-target="#myModal" onclick="twiceCheck('start')"><i class="fa fa-play"></i> 启动</a>
                <a class="btn btn-warning btn-xs" data-toggle="modal" data-target="#myModal" onclick="twiceCheck('pause')"><i class="fa fa-pause"></i> 暂停</a>
                <a class="btn btn-success btn-xs" data-toggle="modal" data-target="#myModal" onclick="twiceCheck('finish')"><i class="fa fa-stop"></i> 完成</a>
                <div class="btn-group btn-group-circle btn-group-xs pull-right  ">
                  <a class="btn btn-default tooltips" data-container="body" data-original-title="前一个任务" id="task_prev" href="task_detail.php?idx=<?php echo $arrIdx['prev'];?>"><i class="fa fa-angle-left"></i> Prev</a>
                  <div class="btn-group btn-group-xs">
                    <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">
                      近期任务 <i class="fa fa-angle-down"></i>
                    </button>
                    <ul class="dropdown-menu pull-right" role="menu" style="overflow:auto;" id="list">
                      <li>
                        <a href="javascript:;"> 获取中 </a>
                      </li>
                    </ul>
                  </div>
                  <a class="btn btn-default tooltips" data-container="body" data-original-title="后一个任务" id="task_next" href="task_detail.php?idx=<?php echo $arrIdx['next'];?>">Next <i class="fa fa-angle-right"></i></a>
                </div>
              </div>
              <div class="col-sm-12">
                <hr style="margin-top:5px;margin-bottom: 5px;" />
              </div>
              <span class="col-sm-2">任务状态</span>
              <span class="col-sm-10" id="state"><span class="badge badge-default">获取中</span></span>
              <span class="col-sm-2">任务名称</span>
              <span class="col-sm-10" id="task_name">&nbsp;</span>
              <span class="col-sm-2">服务池</span>
              <span class="col-sm-4" id="pool_name">&nbsp;</span>
              <span class="col-sm-2">上线TAG</span>
              <span class="col-sm-4" id="tag">&nbsp;</span>
              <span class="col-sm-2">任务步长</span>
              <span class="col-sm-4" id="step_len">&nbsp;</span>
              <span class="col-sm-2">创建时间</span>
              <span class="col-sm-4" id="created">&nbsp;</span>
              <span class="col-sm-2">任务模板</span>
              <span class="col-sm-4" id="template_id">&nbsp;</span>
              <span class="col-sm-2">更新时间</span>
              <span class="col-sm-4" id="updated">&nbsp;</span>
              <span class="col-sm-2">任务参数</span>
              <span class="col-sm-4" id="arg"><a class="tooltips" title="查看任务参数" data-toggle="modal" data-target="#myViewModal" onclick="viewArg()"><i class="fa fa-bars"></i></a></span>
              <div class="col-sm-12">
                <hr style="margin-top:5px;margin-bottom: 5px;" />
              </div>
            </div>
          </div>

          <div class="x_panel">
            <div class="x_title">
              <h2 class="text-primary"><i class="fa fa-tasks"></i> 任务概览 <small>概要统计</small></h2>
              <ul class="nav navbar-right panel_toolbox">
                <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                </li>
                <li><a class="close-link"><i class="fa fa-close"></i></a>
                </li>
              </ul>
              <div class="clearfix"></div>
            </div>
            <div class="x_content">
              <div class="table-scrollable">
                <table class="table table-bordered table-hover">
                  <thead class="flip-content">
                  <tr>
                    <th>应用池子名称</th>
                    <th width="6%">总台数</th>
                    <th width="6%">准备中</th>
                    <th width="6%">执行中</th>
                    <th width="6%">暂停</th>
                    <th width="6%">成功</th>
                    <th width="6%">失败</th>
                    <th width="20%">进度</th>
                  </tr>
                  </thead>
                  <tbody id="task_process"></tbody>
                </table>
              </div>
            </div>
          </div>

          <div class="x_panel">
            <div class="x_title">
              <h2 class="text-primary"><i class="fa fa-list-ul"></i> 任务细节
                <small>任务分发状态详情
                  <a class="tooltips" id="viewTaskLog" data-container="body" data-trigger="hover" data-original-title="查看日志" data-toggle="modal" data-target="#myViewModal"><i class="fa fa-comment"></i></a>
                </small>
              </h2>
              <ul class="nav navbar-right panel_toolbox">
                <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                </li>
                <li><a class="close-link"><i class="fa fa-close"></i></a>
                </li>
              </ul>
              <div class="clearfix"></div>
            </div>
            <div class="x_content">
              <div class="" role="tabpanel" data-example-id="togglable-tabs">
                <ul id="myTab" class="nav nav-tabs bar_tabs" role="tablist">
                  <li role="presentation" class="active" id="tab_home_1">
                    <a href="#tab_1" id="home-tab" role="tab" data-toggle="tab" aria-expanded="true">准备中<span id="num_ready" class="badge bg-info" style="position: absolute;top: -3px;right: -10px;"></span></a>
                  </li>
                  <li role="presentation" class="" id="tab_home_2">
                    <a href="#tab_2" role="tab" id="profile-tab" data-toggle="tab" aria-expanded="false">执行中<span id="num_running" class="badge bg-blue" style="position: absolute;top: -3px;right: -10px;"></span></a>
                  </li>
                  <li role="presentation" class="" id="tab_home_5">
                     <a href="#tab_5" role="tab" id="profile-tab2" data-toggle="tab" aria-expanded="false">暂停<span id="num_stoped" class="badge bg-orange" style="position: absolute;top: -3px;right: -10px;"></span></a>
                  </li>
                  <li role="presentation" class="" id="tab_home_3">
                    <a href="#tab_3" role="tab" id="profile-tab2" data-toggle="tab" aria-expanded="false">成功<span id="num_success" class="badge bg-green" style="position: absolute;top: -3px;right: -10px;"></span></a>
                  </li>
                  <li role="presentation" class="" id="tab_home_4">
                    <a href="#tab_4" role="tab" id="profile-tab2" data-toggle="tab" aria-expanded="false">失败<span id="num_failed" class="badge bg-red" style="position: absolute;top: -3px;right: -10px;"></span></a>
                  </li>

                </ul>
                <div id="myTabContent" class="tab-content">
                  <div role="tabpanel" class="tab-pane fade active in" id="tab_1">
                    <div class="table-scrollable">
                      <table class="table table-bordered table-hover">
                        <thead class="flip-content">
                        <tr>
                          <th width="4%"><div class="checker" id="uniform-SelectAll"><span><input type="checkbox" name="SelectAll" id="SelectAll" onclick="checkAll(this,'ready')"></span></div></th>
                          <th width="6%">#</th>
                          <th width="15%">IP</th>
                          <th>步骤</th>
                          <th width="10%">操作</th>
                        </tr>
                        </thead>
                        <tbody id="task_ready"></tbody>
                      </table>
                    </div>
                  </div>
                  <div role="tabpanel" class="tab-pane fade" id="tab_2">
                    <div class="table-scrollable">
                      <table class="table table-bordered table-hover">
                        <thead class="flip-content">
                        <tr>
                          <th width="4%"><div class="checker" id="uniform-SelectAll"><span><input type="checkbox" name="SelectAll" id="SelectAll" onclick="checkAll(this,'running')"></span></div></th>
                          <th width="6%">#</th>
                          <th width="15%">IP</th>
                          <th>步骤 <span class="badge bg-blue">当前</span></th>
                          <th width="8%">时间</th>
                          <th width="10%">操作</th>
                        </tr>
                        </thead>
                        <tbody id="task_running"></tbody>
                      </table>
                    </div>
                  </div>
                  <div role="tabpanel" class="tab-pane fade" id="tab_3">
                    <div class="table-scrollable">
                      <table class="table table-bordered table-hover">
                        <thead class="flip-content">
                        <tr>
                          <th width="4%"><div class="checker" id="uniform-SelectAll"><span><input type="checkbox" name="SelectAll" id="SelectAll" onclick="checkAll(this,'success')"></span></div></th>
                          <th width="6%">#</th>
                          <th width="15%">IP</th>
                          <th>步骤 <span class="badge bg-green">当前</span></th>
                          <th width="8%">时间</th>
                          <th width="10%">操作</th>
                        </tr>
                        </thead>
                        <tbody id="task_success"></tbody>
                      </table>
                    </div>
                  </div>
                  <div role="tabpanel" class="tab-pane fade" id="tab_4">
                    <div class="table-scrollable">
                      <table class="table table-bordered table-hover">
                        <thead class="flip-content">
                        <tr>
                          <th width="4%"><div class="checker" id="uniform-SelectAll"><span><input type="checkbox" name="SelectAll" id="SelectAll" onclick="checkAll(this,'failed')"></span></div></th>
                          <th width="6%">#</th>
                          <th width="15%">IP</th>
                          <th>步骤 <span class="badge bg-red">当前</span></th>
                          <th width="8%">时间</th>
                          <th width="10%">操作</th>
                        </tr>
                        </thead>
                        <tbody id="task_failed"></tbody>
                      </table>
                    </div>
                  </div>
                  <div role="tabpanel" class="tab-pane fade" id="tab_5">
                        <div class="table-scrollable">
                            <table class="table table-bordered table-hover">
                                <thead class="flip-content">
                                <tr>
                                    <th width="4%"><div class="checker" id="uniform-SelectAll"><span><input type="checkbox" name="SelectAll" id="SelectAll" onclick="checkAll(this,'stoped')"></span></div></th>
                                    <th width="6%">#</th>
                                    <th width="15%">IP</th>
                                    <th>步骤 <span class="badge bg-orange">当前</span></th>
                                    <th width="8%">时间</th>
                                    <th width="10%">操作</th>
                                </tr>
                                </thead>
                                <tbody id="task_stoped"></tbody>
                            </table>
                        </div>
                    </div>
                </div>
              </div>
            </div>
          </div>

          <form method="post" id="my_form_1" class="form-horizontal">
            <div class="modal fade bs-modal-lg" id="myModal" role="dialog" aria-hidden="true">
              <div class="modal-dialog modal-lg">
                <div class="modal-content">
                  <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myModalLabel">Loading ...</h4>
                  </div>
                  <div class="modal-body" style="overflow:auto;" id="myModalBody">
                    <p> </p>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                    <button type="button" class="btn btn-success" id="btnCommit" data-dismiss="modal" onclick="change()" style="margin-bottom: 5px;" disabled>提交</button>
                  </div>
                </div>
              </div>
            </div>
          </form>
          <form method="post" class="form-horizontal">
            <div class="modal fade bs-modal-lg" id="myChildModal" role="dialog" aria-hidden="true">
              <div class="modal-dialog modal-lg">
                <div class="modal-content">
                  <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myChildModalLabel">Loading ...</h4>
                  </div>
                  <div class="modal-body" style="overflow:auto;" id="myChildModalBody">
                    <p> </p>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                    <button type="button" class="btn btn-success" id="btnCommit" data-dismiss="modal" onclick="change()" style="margin-bottom: 5px;" disabled>提交</button>
                  </div>
                </div>
              </div>
            </div>
          </form>
          <form method="post" class="form-horizontal">
            <div class="modal fade bs-modal-lg" id="myViewModal" role="dialog" aria-hidden="true">
              <div class="modal-dialog modal-lg">
                <div class="modal-content">
                  <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myViewModalLabel">Loading ...</h4>
                  </div>
                  <div class="modal-body" style="overflow:auto;line-height:200%" id="myViewModalBody">
                    <p> </p>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                  </div>
                </div>
              </div>
            </div>
          </form>
          <form method="post" class="form-horizontal">
            <div class="modal fade bs-modal-lg" id="myRevealModal" role="dialog" aria-hidden="true">
              <div class="modal-dialog modal-lg">
                <div class="modal-content">
                  <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myRevealModalLabel">帮助</h4>
                  </div>
                  <div class="modal-body" style="height:500px;" id="myRevealModalBody">
                    <p> </p>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                  </div>
                </div>
              </div>
            </div>
          </form>
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

<!-- jQuery -->
<script src="../gentelella/vendors/jquery/dist/jquery.min.js"></script>
<!-- Bootstrap -->
<script src="../gentelella/vendors/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- FastClick -->
<script src="../gentelella/vendors/fastclick/lib/fastclick.js"></script>
<!-- NProgress -->
<script src="../gentelella/vendors/nprogress/nprogress.js"></script>
<!-- bootstrap-progressbar -->
<script src="../gentelella/vendors/bootstrap-progressbar/bootstrap-progressbar.min.js"></script>
<!-- iCheck -->
<script src="../gentelella/vendors/iCheck/icheck.min.js"></script>
<!-- bootstrap-daterangepicker -->
<script src="../gentelella/production/js/moment/moment.min.js"></script>
<script src="../gentelella/production/js/datepicker/daterangepicker.js"></script>
<!-- bootstrap-wysiwyg -->
<script src="../gentelella/vendors/bootstrap-wysiwyg/js/bootstrap-wysiwyg.min.js"></script>
<script src="../gentelella/vendors/jquery.hotkeys/jquery.hotkeys.js"></script>
<script src="../gentelella/vendors/google-code-prettify/src/prettify.js"></script>
<!-- jQuery Tags Input -->
<script src="../gentelella/vendors/jquery.tagsinput/src/jquery.tagsinput.js"></script>
<!-- Switchery -->
<script src="../gentelella/vendors/switchery/dist/switchery.min.js"></script>
<!-- Select2 -->
<script src="../gentelella/vendors/select2/dist/js/select2.full.min.js"></script>
<!-- Parsley -->
<script src="../gentelella/vendors/parsleyjs/dist/parsley.min.js"></script>
<!-- Autosize -->
<script src="../gentelella/vendors/autosize/dist/autosize.min.js"></script>
<!-- jQuery autocomplete -->
<script src="../gentelella/vendors/devbridge-autocomplete/dist/jquery.autocomplete.min.js"></script>
<!-- starrr -->
<script src="../gentelella/vendors/starrr/dist/starrr.js"></script>
<!-- PNotify -->
<script src="../gentelella/vendors/pnotify/dist/pnotify.js"></script>
<script src="../gentelella/vendors/pnotify/dist/pnotify.buttons.js"></script>
<script src="../gentelella/vendors/pnotify/dist/pnotify.nonblock.js"></script>
<!-- reveal -->
<script src="../gentelella/vendors/reveal330/lib/js/head.min.js"></script>
<script src="../gentelella/vendors/reveal330/js/reveal.js"></script>

<!-- Custom Theme Scripts -->
<script src="../gentelella/build/js/custom.min.js"></script>
<!-- page level -->
<script src="../js/pnotify.js"></script>
<script src="../js/switchery.js"></script>
<script src="../js/login.js"></script>
<script src="../js/locale_messages.js"></script>
<script src="../js/reveal.js?_t=<?php echo date('U');?>"></script>
<script src="../js/for_layout/taskdetail.js?_t=<?php echo date('U');?>"></script>


<!-- Custom Notification -->
<script>
  $(document).ready(function() {
    $("select.form-control").select2({width:'100%'});
    cache.task_id=<?php echo $myIdx;?>;
    window.setTimeout('getTask(\'info\');',200);
    cache.refreshInterval = setInterval('getTask(\'info\');',10000);
    getList();
    $('#fIdx').bind('keypress',function(event){
      if(event.keyCode == "13"){
        list(1);
      }
    });

    //查看任务主日志
    $('#viewTaskLog').click(function(){
        view('tasklog',cache.task_id,0,0);
    })



  });
  $("#myModal").on("shown.bs.modal", function(){
    $("select.form-control").select2();
  });
  $("#myModal").on("hidden.bs.modal", function() {
    $(this).removeData("bs.modal");
    $('#myModalBody').css('height','');
    $('#myModalLabel').html('Loading ...');
    $("#myModalBody").html('<p> </p>');
  });

  $("#myChildModal").on("shown.bs.modal", function(){
    $("select.form-control").select2();
  });
  $("#myChildModal").on("hidden.bs.modal", function() {
    $(this).removeData("bs.modal");
    $('#myChildModalBody').css('height','');
    $('#myChildModalLabel').html('Loading ...');
    $("#myChildModalBody").html('<p> </p>');
  });

  $("#myViewModal").on("shown.bs.modal", function(){
    $("select.form-control").select2();
  });
  $("#myViewModal").on("hidden.bs.modal", function() {
    $(this).removeData("bs.modal");
    $('#myViewModalBody').css('height','');
    $('#myViewModalLabel').html('Loading ...');
    $("#myViewModalBody").html('<p> </p>');
  });


</script>
<!-- /Custom Notification -->
</body>
</html>
