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
$myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):0;
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

        <div class="" style="background-color:#fff;">
          <div class="" role="tabpanel" data-example-id="togglable-tabs">
            <ul id="myTab" class="nav nav-tabs bar_tabs" role="tablist" style="margin-bottom: 10px;">
              <li id="tab_1" role="presentation">
                <a data-toggle="tab" role="tab" aria-expanded="true" onclick="getList('group','','upstream')">Upstream</a>
              </li>
              <li id="tab_2" role="presentation" class="active">
                <a data-toggle="tab" role="tab" aria-expanded="true" onclick="getList('group')">单元</a>
              </li>
              <li id="tab_3" role="presentation">
                <a data-toggle="tab" role="tab" aria-expanded="false" onclick="getList('unit')">节点</a>
              </li>
            </ul>
            <div id="myTabContent" class="tab-content">
              <div role="tabpanel" class="tab-pane fade active in" id="tab_content1" aria-labelledby="home-tab">
                <div class="x_panel" style="border: 0px;">
                  <div class="row">
                    <div class="col-md-9 form-group">
                      <div class="btn-group">
                        <div class="hidden">
                          <input type="hidden" id="tab" name="tab" value="unit">
                        </div>
                        <div class="col-sm-4" style="padding-left:0px;">
                          <div class="input-group">
                            <span class="input-group-addon">分组</span>
                            <select class="form-control" id="fGroup" onchange="list(1)">
                              <option value="">全部分组</option>
                            </select>
                          </div>
                        </div>
                        <div class="col-sm-4" style="padding-left:0px;" hidden="hidden">
                          <div class="input-group">
                            <span class="input-group-addon">单元</span>
                            <select class="form-control" id="fUnit" onchange="list(1)">
                              <option value="">全部单元</option>
                            </select>
                          </div>
                        </div>
                        <div class="col-sm-3" style="padding-left:0px;">
                          <div class="input-group">
                            <input type="name" id="fIdx" class="form-control" placeholder="关键字">
                                <span class="input-group-btn">
                                  <button class="btn btn-default" type="button" onclick="list(1)">Go!</button>
                                </span>
                          </div>
                        </div>
                        <div style="padding-left:0px;">
                          <button class="btn btn-default" type="button" onclick="reset()">重置</button>
                        </div>
                      </div>
                    </div>
                    <div class="col-md-3">
                      <div class="btn-group pull-right" id="tab_toolbar">
                      </div>
                    </div>
                  </div>
                  <table class="table table-bordered table-hover" id="page_table">
                    <thead id="table-head">
                    <tr>
                      <td>Loading ...</td>
                    </tr>
                    </thead>
                    <tbody id="table-body">
                    </tbody>
                  </table>
                  <div class="row">
                    <div class="col-md-5 col-sm-5">
                      <div class="dataTables_info" id="table-pageinfo" role="status" aria-live="polite">Showing 1 to 0 of 0 entries</div>
                    </div>
                    <div class="col-md-7 col-sm-7">
                      <div class="dataTables_paginate paging_bootstrap_full_number" id="sample_1_paginate">
                        <ul class="pagination" style="visibility: visible;margin-top: 0px;margin-bottom: 0px;" id="table-paginate">
                          <li><a href="javascript:;" onclick="list(1)"><i class="fa fa-angle-left"></i></a></li>
                          <li class="active">
                            <a href="javascript:;" onclick="list(1)">1</a>
                          </li>
                          <li class="next">
                            <a href="javascript:;" title="Next" onclick="list(1)"><i class="fa fa-angle-right"></i></a>
                          </li>
                        </ul>
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
                    <div class="modal fade bs-modal-lg" id="myViewChildModal" role="dialog" aria-hidden="true">
                      <div class="modal-dialog modal-lg">
                        <div class="modal-content">
                          <div class="modal-header">
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                            <h4 class="modal-title" id="myViewChildModalLabel">Loading ...</h4>
                          </div>
                          <div class="modal-body" style="overflow:auto;line-height:200%" id="myViewChildModalBody">
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
<script src="../js/for_hubble/nginx_unit.js?_t=<?php echo date('U');?>"></script>


<!-- Custom Notification -->
<script>
  $(document).ready(function() {
    $("select.form-control").select2({width:'100%'});
    $('#fIdx').bind('keypress',function(event){
      if(event.keyCode == "13"){
        list(1);
      }
    });
    cache.group_id=<?php echo $myIdx;?>;
    window.setTimeout("getList()",200);
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
