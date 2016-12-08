<?php
  $myAction=(isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'add';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
  switch($myAction){
    case 'add':
      $myTitle='添加服务发现类型';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改服务发现类型';
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
              <input type="text" class="form-control" id="name" name="name" onkeyup="check()" onchange="check()" placeholder="名称,eg:for_V4_Core" <?php if($myIdx) echo 'readonly';?>>
            </div>
          </div>
          <div class="form-group">
            <label for="type" class="col-sm-2 control-label">发现类型</label>
            <div class="col-sm-10">
              <select class="form-control" id="type" name="type" onchange="getType()">
                <option value="">请选择</option>
                <option value="NGINX">Nginx</option>
                <option value="SLB">AliYun</option>
              </select>
            </div>
          </div>
          <div class="row col-sm-12"><hr class="col-sm-12" style="margin-top: 5px;margin-bottom: 5px;"></div>
          <h5 class="col-sm-12 text-primary"><strong style="margin-left:20px;">关联选项</strong></h5>
          <div class="profile_details" id="params">
          </div>
          <input type="hidden" id="id" name="id" value="<?php echo $myIdx;?>" />
          <input type="hidden" id="page_action" name="page_action" value="<?php echo $pageAction;?>" />
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">取消</button>
            <button class="btn btn-success" data-dismiss="modal" id="btnCommit" onclick="change()" style="margin-bottom: 5px;" disabled>确认</button>
        </div>
        <script>
          cache.alteration={};cache.params_value={};
          <?php if($myAction=='edit'){echo 'get(\''.$myIdx.'\');'."\n";} ?>
        </script>
