<?php
  $myAction=(isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'add';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
  switch($myAction){
    case 'add':
      $myTitle='添加任务模板';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改任务模板';
      $pageAction='update';
      break;
    default:
      $myTitle='错误请求';
      $pageAction='Illegal';
      break;
  }
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel"><?php echo $myTitle;?></h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" onkeyup="check()" onchange="check()" placeholder="名称,eg:扩容" <?php if($myAction=='edit'){echo 'readonly';} ?>>
            </div>
          </div>
          <div class="form-group">
            <label for="desc" class="col-sm-2 control-label">描述</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="desc" name="desc" onkeyup="check()" onchange="check()" placeholder="描述,eg:扩容">
            </div>
          </div>
          <div class="form-group">
            <label for="steps" class="col-sm-2 control-label">步骤</label>
            <div class="col-sm-10 profile_details">
              <div class="well profile_view col-sm-12">
                <div class="col-sm-12">
                </div>
                <div class="col-sm-12">
                  <table class="table table-striped table-hover">
                    <thead>
                    <tr>
                      <th>顺序</th>
                      <th>步骤名称</th>
                      <th>重试次数</th>
                      <th>忽略错误</th>
                      <th>#</th>
                    </tr>
                    </thead>
                    <tbody id="task_step">
                    </tbody>
                  </table>
                  <a class="btn btn-primary btn-xs" data-toggle="modal" data-target="#myChildModal" href="edit_task_tpl_step.php">添加</a>
                </div>
              </div>
            </div>
          </div>
          <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          <?php if($myAction=='edit'){echo 'get(\''.$myIdx.'\');'."\n";}else{echo 'cache.task_step=[];'."\n";} ?>
          listTaskStep();
        </script>
