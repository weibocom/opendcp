function addRunElement() {
    var Run = "  <div class='form-group' style='top:10px'> "+
    " <label for='run.value' class='col-sm-2 control-label'>执行命令:</label>"+
    " <div class='col-sm-10'>"+
    " <div class='col-sm-9'>"+
    "   <input class='form-control' name='run.value' id='run.value'/>"+
        "</div>"+
    "   <span class='glyphicon glyphicon-remove' style='left:5px;top:10px;' onclick='deleteRun(event)'></span>"+
        " </div>"+
    " </div>";

    $("#div_run").append(Run);
}

function deleteRun(event) {
    event = event ? event : window.event;
    var obj = event.srcElement ? event.srcElement : event.target;
    var obj = $(obj);
    var panel = obj.parents(".form-group")[0];
    $(panel).remove();
}