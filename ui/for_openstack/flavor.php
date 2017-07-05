<?php

require_once('../include/config.inc.php');
require_once('../include/function.php');
require_once('../include/func_session.php');
require_once('../include/navbar.php');
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
                  <li><a href="/">Dashboard</a></li>
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
                <a data-toggle="tab" role="tab" aria-expanded="true" onclick="location.href='/for_openstack/machine.php';">物理节点管理</a>
              </li>
              <li id="tab_2" role="presentation">
                <a data-toggle="tab" role="tab" aria-expanded="false" onclick="location.href='/for_openstack/initlist.php';">初始化任务列表</a>
              </li>
              <li id="tab_3" role="presentation">
                <a data-toggle="tab" role="tab" aria-expanded="false" onclick="location.href='/for_openstack/network.php';">网络管理</a>
              </li>
              <li id="tab_4" role="presentation" class="active">
                <a data-toggle="tab" role="tab" aria-expanded="false" onclick="location.href='/for_openstack/flavor.php';">机型管理</a>
              </li>
            </ul>
            <div id="myTabContent" class="tab-content">
              <div role="tabpanel" class="tab-pane fade active in" id="tab_content1" aria-labelledby="home-tab">
                <div class="x_panel" style="border: 0px;">

<?php

require_once('../include/openstack.php');
openstack::needOpenstackLogin();
$arr_flavor = openstack::getFlavorList();

?>

		  <div class="col-md-9">
                  </div>
		  <div class="col-md-2">
                      <div class="btn-group pull-right" id="tab_toolbar"><a type="button" class="btn btn-success" data-toggle="modal" data-target="#myAddFlavorModal" href="add_flavor.php?action=add"> 创建机型 <i class="fa fa-plus"></i></a></div>
                  </div>


                  <table class="table table-bordered table-hover">
                    <thead class="flip-content">
                    <tr>
                      <th>flavor id</th>
                      <th>名称</th>
                      <th width="20%">CPU（核）</th>
                      <th width="20%">内存（M）</th>
                      <th width="20%">硬盘（G）</th>
                    </tr>
                    </thead>
                    <tbody id="task_process">
			<?php foreach($arr_flavor['flavors'] as $oneflavor) { ?>
			<tr>
				<td><?=$oneflavor['id']?></td>
				<td><?=$oneflavor['name']?></td>
				<td><?=$oneflavor['vcpus']?></td>
				<td><?=$oneflavor['ram']?></td>
				<td><?=$oneflavor['disk']?></td>
			</tr>
			<?php } ?>
		    </tbody>
                  </table>


                  <form method="post" class="form-horizontal">
                    <div class="modal fade bs-modal-lg" id="myAddFlavorModal" role="dialog" aria-hidden="true">
                      <div class="modal-dialog modal-lg">
                        <div class="modal-content">
                          <div class="modal-header">
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                            <h4 class="modal-title" id="myAddFlavorModalLabel">Loading ...</h4>
                          </div>
                          <div class="modal-body" style="overflow:auto;" id="myAddFlavorModalBody">
                            <p> 
				
			    </p>
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
                    <div class="modal fade bs-modal-lg" id="myViewNetworkModal" role="dialog" aria-hidden="true">
                      <div class="modal-dialog modal-lg">
                        <div class="modal-content">
                          <div class="modal-header">
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                            <h4 class="modal-title" id="myViewNetworkModalLabel">Loading ...</h4>
                          </div>
                          <div class="modal-body" style="overflow:auto;" id="myViewNetworkModalBody">
                            <p> 
				
			    </p>
                          </div>
                          <div class="modal-footer">
                            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                            <button type="button" class="btn btn-success" id="btnCommit" data-dismiss="modal" onclick="change()" style="margin-bottom: 5px;" disabled>提交</button>
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
<script src="../js/for_layout/service.js?_t=<?php echo date('U');?>"></script>


<!-- Custom Notification -->
<script>
  $(document).ready(function() {
    $("select.form-control").select2({width:'100%'});
    $('#fIdx').bind('keypress',function(event){
      if(event.keyCode == "13"){
        //list(1);
      }
    });
    window.setTimeout('getList()',200);
  });

</script>
<!-- /Custom Notification -->
</body>
</html>
