<?php
  $myAction=(isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'add';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
  switch($myAction){
    case 'add':
      $myTitle='添加服务';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改服务';
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
            <label for="cluster_id" class="col-sm-2 control-label">隶属集群</label>
            <div class="col-sm-10">
              <select class="form-control" id="cluster_id" name="cluster_id">
                <option value="">请选择</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" onkeyup="check()" onchange="check()" placeholder="名称,eg:test_cluster" <?php if($myAction=='edit'){echo 'readonly';} ?>>
            </div>
          </div>
          <div class="form-group">
            <label for="desc" class="col-sm-2 control-label">描述</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="desc" name="desc" onkeyup="check()" onchange="check()" placeholder="描述,eg:测试集群">
            </div>
          </div>
          <div class="form-group">
            <label for="service_type" class="col-sm-2 control-label">服务类型</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="service_type" name="service_type" onkeyup="check()" placeholder="服务类型,eg:JAVA,PHP">
            </div>
          </div>
          <div class="form-group">
            <label for="docker_image" class="col-sm-2 control-label">镜像名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="docker_image" name="docker_image" onkeyup="check()" placeholder="镜像名称,eg:registry.xxx.com/base/nginx:default">
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
          updateSelect('cluster_id');
          $('#cluster_id').select2({disabled:true});
          $('#cluster_id').val($('#fClusterId').val()).trigger('change');
          <?php if($myAction=='edit'){echo 'get(\''.$myIdx.'\');'."\n";} ?>
        </script>
