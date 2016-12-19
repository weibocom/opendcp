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


$myIdx=(isset($_GET['idx'])&&!empty($_GET['idx']))?trim($_GET['idx']):'';
?>
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">AppKey授权 - <?php echo $myIdx;?></h4>
        </div>
        <div class="modal-body" style="overflow:auto;" id="myModalBody">
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
              <div class="dataTables_paginate paging_bootstrap_full_number" id="sample_1_paginate">
                <ul class="pagination" style="visibility: visible;" id="modal_table-paginate">
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal">关闭</button>
        </div>
        <script>
          modalList(1,'appkey','<?php echo $myIdx;?>');
        </script>
