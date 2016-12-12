<?php
/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
$myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
$myParId=(isset($_GET['par_id'])&&!empty($_GET['par_id']))?trim($_GET['par_id']):'';
$myParName=(isset($_GET['par_name'])&&!empty($_GET['par_name']))?trim($_GET['par_name']):'';
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myViewModalLabel">查看文件历史 - <?php echo $myIdx;?></h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myViewModalBody">
          <div class="table-scrollable">
            <table class="table table-bordered table-striped table-hover" id="modal_table">
              <thead id="modal_table-head">
                <tr>
                  <td>Loading ...</td>
                </tr>
              </thead>
              <tbody id="modal_table-body">
              </tbody>
            </table>
          </div>
          <div class="row">
            <div class="col-md-5 col-sm-5">
              <div class="dataTables_info" id="modal_table-pageinfo" role="status" aria-live="polite">Showing 0 to 0 of 0 records</div>
            </div>
            <div class="col-md-7 col-sm-7">
              <div class="dataTables_paginate paging_bootstrap_full_number">
                <ul class="pagination" style="visibility: visible;margin-top: 0px;margin-bottom: 0px;" id="modal_table-paginate">
                </ul>
              </div>
            </div>
          </div>
        </div>
        <input type="hidden" id="modal_idx" value="<?php echo $myIdx;?>" />
        <input type="hidden" id="modal_unit_id" value="<?php echo $myParId;?>" />
        <input type="hidden" id="modal_unit_name" value="<?php echo $myParName;?>" />
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">关闭</button>
        </div>
        <script>
          modalList(1);
        </script>
