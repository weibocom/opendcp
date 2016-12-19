function addPort() {
    var env = "<div class='form-inline' style='top:10px'>" +
                   "<div class='form-group col-sm-6' style='margin-top:5px'>" +
                        "<label for='expose.port' class='col-sm-3 control-label'>port:</label><input class='form-control' name='expose.port' id='expose.port'></input>" +
                        "<span class='glyphicon glyphicon-remove' style='left:5px' onclick='deletePort(event)'></span>" +
                    "</div>" +
              "</div>";
     $("#div_port").append(env);
}

function deletePort(event) {
    event = event ? event : window.event;
    var obj = event.srcElement ? event.srcElement : event.target;
    var obj = $(obj);
    var panel = obj.parents(".form-inline")[0];
    $(panel).remove();
}