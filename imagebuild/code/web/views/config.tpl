<!DOCTYPE html>

<html>
<meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
<head>
  <title>Beego</title>
</head>

<body>
<h1>{{.project}}</h1>
{{if eq .showSubmit "show"}}
<form id="form_config" class="form-horizontal" action="/api/project/save?1=1" method="post">
{{end}}

{{str2html .view}}

{{if eq .showSubmit "show"}}
    <div class="form-group">
        <div class="col-sm-10">
            <input class="btn btn-primary" type="submit" id="btn_submit" style="margin-top:20px" value="{{i18n .Lang "submit"}}" />
        </div>
    </div>
    </form>
</body>
{{end}}
</html>
