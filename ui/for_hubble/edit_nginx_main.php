<?php
  $myAction=(isset($_GET['action'])&&!empty($_GET['action']))?trim($_GET['action']):'add';
  $myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
  $myParId=(isset($_GET['par_id'])&&!empty($_GET['par_id']))?trim($_GET['par_id']):'';
  $myParName=(isset($_GET['par_name'])&&!empty($_GET['par_name']))?trim($_GET['par_name']):'';
  switch($myAction){
    case 'add':
      $myTitle='添加主配置文件';
      $pageAction='insert';
      break;
    case 'edit':
      $myTitle='修改主配置文件';
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
            <label for="unit_id" class="col-sm-2 control-label">隶属单元</label>
            <div class="col-sm-10">
              <select class="form-control" id="unit_id" name="unit_id" disabled>
                <option value="<?php echo $myParId;?>"><?php echo $myParName;?></option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="name" class="col-sm-2 control-label">文件名称</label>
            <div class="col-sm-10">
              <input type="text" class="form-control" id="name" name="name" onkeyup="check()" onchange="check()" placeholder="名称,eg:nginx.conf" value="nginx.conf" <?php if($myIdx) echo 'readonly';?>>
            </div>
          </div>
          <div class="form-group">
            <label for="content" class="col-sm-2 control-label">文件内容</label>
            <div class="col-sm-10">
              <textarea rows="10" class="form-control" id="content" name="content" onkeyup="check()" placeholder="文件内容">worker_processes 12;
worker_cpu_affinity 000000000001 000000000010 000000000100 000000001000 000000010000 000000100000 000001000000 000010000000 000100000000 001000000000 010000000000 100000000000;
worker_rlimit_nofile 65535;

error_log /data0/www/logs/error.log warn;
pid /data0/www/logs/nginx.pid;

events {
	use epoll;
	worker_connections 65535;
}

http {
	default_type  text/plain;

	log_format main $server_addr  ' $remote_addr - $remote_user [$time_local] "$request" '
		'$status $body_bytes_sent "$http_referer" '
		'"$http_user_agent" "$http_x_forwarded_for" "$proxy_host" "$upstream_addr" "$upstream_status" "$upstream_response_time" "$request_time"';

	access_log  /data0/www/logs/access.log  main buffer=1024 flush=1m;

	keepalive_timeout  5;
	keepalive_requests 10000;
	client_header_timeout 5s;
	client_body_timeout 10s;
	send_timeout 5s;
        underscores_in_headers  on;
	server_names_hash_bucket_size 64;

        client_body_buffer_size 256k;

	server_tokens off;

	merge_slashes on;
	reset_timedout_connection on;
        check_shm_size 10M;

	include upstream/openapi_webv2-yf-core-inner.upstream;
        include upstream/openapi_webv2-yf-statuses-inner.upstream;
        include upstream/openapi_webv2-yfali-feed-inner.upstream;
        include upstream/openapi_webv2-yfali-statuses-inner.upstream;
        include upstream/openapi_friendship-tcali-core.upstream;
        include upstream/openapi_friendship-tc-core.upstream;

	include vhost/openapi.vhost.conf;

	server {
		listen 80;
		server_name _;
		location / {
			deny all;
		}
		error_page 500 502 503 504 /50x.html;
		location = /50x.html {
			root html;
		}
	}
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
              case 'add':
                echo 'getDefault();'."\n";
                break;
              case 'edit':
                echo '$(\'#unit_id\').select2({disabled:true});'."\n";
                echo 'get(\''.$myIdx.'\');'."\n";
                break;
            }
          ?>
        </script>
