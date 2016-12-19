<!DOCTYPE html>

<html>
<meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
<head>
  <title>Beego</title>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/submit_form.js"></script>
</head>

<body>
<div class="form-horizontal">
  <div class="form-group">
    <label for="projectName" class="col-sm-3 control-label">{{i18n .Lang "Project Name"}}:</label>
    <div class="col-sm-6">
       <input id="projectName" name="projectName" class="form-control" value="{{.name}}" disabled/>
    </div>
  </div>
  <div class="form-group">
    <label for="createTime" class="col-sm-3 control-label">{{i18n .Lang "Create Time"}}:</label>
    <div class="col-sm-6">
       <input id="createTime" name="createTime" class="form-control" value="{{.createTime}}" disabled/>
    </div>
  </div>
  <div class="form-group">
    <label for="creator" class="col-sm-3 control-label">{{i18n .Lang "Creator"}}:</label>
    <div class="col-sm-6">
       <input id="creator" name="creator" class="form-control" value="{{.creator}}" disabled/>
    </div>
  </div>
  <div class="form-group">
    <label for="lastModifyTime" class="col-sm-3 control-label">{{i18n .Lang "Last Modify Time"}}:</label>
    <div class="col-sm-6">
       <input id="lastModifyTime" name="lastModifyTime" class="form-control" value="{{.lastModifyTime}}" disabled/>
    </div>
  </div>
    <div class="form-group">
      <label for="lastModifyOperator" class="col-sm-3 control-label">{{i18n .Lang "Last Modify Operator"}}:</label>
      <div class="col-sm-6">
         <input id="lastModifyOperator" name="lastModifyOperator" class="form-control" value="{{.lastModifyOperator}}" disabled/>
      </div>
    </div>
</body>
</html>
