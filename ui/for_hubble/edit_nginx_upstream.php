<?php
  $myAction=(isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'add';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
  $myParId=(isset($_GET['par_id'])&&!empty($_GET['par_id']))?trim($_GET['par_id']):'';
  $myParName=(isset($_GET['par_name'])&&!empty($_GET['par_name']))?trim($_GET['par_name']):'';
  switch($myAction){
    case 'add':
      $myTitle='添加Upstream文件';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改Upstream文件';
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
            <label for="group_id" class="col-sm-2 control-label">隶属分组</label>
            <div class="col-sm-10">
              <select class="form-control" id="group_id" name="group_id" disabled>
                <option value="<?php echo $myParId;?>"><?php echo $myParName;?></option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">文件名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" onkeyup="check()" onchange="check()" placeholder="名称,eg:default_idc.upstream" value="default.upstream" <?php if($myIdx) echo 'readonly';?>>
            </div>
          </div>
          <div class="form-group">
            <label for="content" class="col-sm-2 control-label">文件内容</label>
            <div class="col-sm-10">
              <textarea rows="10" class="form-control" id="content" name="content" onkeyup="check()" placeholder="文件内容">upstream default_upstream{
		keepalive 1;
		server 127.0.0.1:8080 max_fails=0 fail_timeout=30s weight=50;
		check interval=1000 rise=3 fall=2 timeout=3000 type=http default_down=false;
		check_http_send "GET / HTTP/1.0\r\n\r\n";
		check_http_expect_alive http_2xx;
}</textarea>
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
          <?php
          switch($myAction){
            case 'edit':
              echo '$(\'#group_id\').select2({disabled:true});'."\n";
              echo 'get(\''.$myIdx.'\');'."\n";
              break;
          }
          ?>
        </script>
