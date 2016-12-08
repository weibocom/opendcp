function addEnvElement() {
    var env = "<div class='form-inline' style='top:10px'>" +
                        "<div class='form-group col-sm-6' style='margin-top:5px'>" +
                            "<label for='env_set.key' class='col-sm-3 control-label'>键:</label><input class='form-control' name='env_set.key' id='env_set.key'></input>" +
                        "</div>" +
                        "<div class='form-group col-sm-6' style='margin-top:5px'>" +
                             "<label for='env_set.value' class='col-sm-3 control-label'>值:</label><input class='form-control' name='env_set.value' id='env_set.value'></input>" +
                             "<span class='glyphicon glyphicon-remove' style='left:5px' onclick='deleteEnv(event)'></span>" +
                        "</div>" +
                  "</div>";

    $("#div_env").append(env);
}

function deleteEnv(event) {
    event = event ? event : window.event;
    var obj = event.srcElement ? event.srcElement : event.target;
    var obj = $(obj);
    var panel = obj.parents(".form-inline")[0];
    $(panel).remove();
}