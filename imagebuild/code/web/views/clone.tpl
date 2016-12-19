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
  <div class="form-group" style="margin-top:20px">
    <label for="newProjectName" class="col-sm-3 control-label">{{i18n .Lang "New Project Name"}}:</label>
    <div class="col-sm-6">
       <input id="projectName" name="projectName" class="form-control"/>
    </div>
  </div>

  <div class="form-group">
      <div class="col-sm-offset-3 col-sm-6">
          <input class="btn btn-primary" style="margin-top:20px" type="submit" value="{{i18n .Lang "submit"}}" />
      </div>
  </div>

</body>
</html>