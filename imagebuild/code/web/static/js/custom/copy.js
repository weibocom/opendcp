function addCopyElement() {
    var copy = "<div class='form-inline' style='top:10px'>" +
        "<div class='form-group col-sm-6' style='margin-top:5px'>" +
        "<label for='copy.key' class='col-sm-3 control-label'>源路径:</label><input class='form-control' name='copy.key' id='copy.key'></input>" +
        "</div>" +
        "<div class='form-group col-sm-6' style='margin-top:5px'>" +
        "<label for='copy.value' class='col-sm-3 control-label'>目标路径:</label><input class='form-control' name='copy.value' id='copy.value'></input>" +
        "<span class='glyphicon glyphicon-remove' style='left:5px' onclick='deleteCopy(event)'></span>" +
        "</div>" +
        "</div>";

    $("#div_copy").append(copy);
}

function deleteCopy(event) {
    event = event ? event : window.event;
    var obj = event.srcElement ? event.srcElement : event.target;
    var obj = $(obj);
    var panel = obj.parents(".form-inline")[0];
    $(panel).remove();
}

$("#btn_addcopy").click();