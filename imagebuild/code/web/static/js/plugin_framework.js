function addExtensionPlug(server, selectId, updateId, type) {
    selectPlugin = document.getElementById(selectId)
    plugin = selectPlugin.options[selectPlugin.selectedIndex].value
    if (!canAdd(plugin)) return;
    $.ajax({
            type: "GET",
            url: pluginViewUrl,
            data: {type: type, plugin: plugin},
            success: function(result){
                // id中带有空格，转义处理
                var updateIdEscape = updateId.replace(/\ /g,"\\ ")
                // 此处使用jquery使因为原生的innerHtml会出现丢失input值的情况
                $("#" + updateIdEscape).append(result);
                $("#pluginTitle").html("操作列表");
                $.getScript(server + "/static/js/custom/" + plugin + ".js");
                bindChange();
            },
            error: function(result){
                alert(result);
            }
    });
}

var singlePlugins = "base_image_choose_from_harbor,env_set,expose,download_dockerfile,copy,entrypoint,run,work_dir"
function canAdd(plugin) {
    var canAdd = true;
    if (singlePlugins.indexOf(plugin) > -1){
        var updateIdEscape = updateDivId.replace(/\ /g,"\\ ");
        $("#"+updateDivId).find("input[id='$$plugin']").each(function () {
                if($(this).val()==plugin) {
                    canAdd = false;
                    return false;
                }
        });
        $("#"+updateIdEscape).find("input[id='$$plugin']").each(function () {
            if($(this).val()==plugin) {
                canAdd = false;
                return false;
            }
        });
    }
    return canAdd;
}

function movePluginUp(event) {
    event = event ? event : window.event;
    var obj = event.srcElement ? event.srcElement : event.target;
    var obj = $(obj);
    var panel = obj.parents(".panel")[0];
    var prev = $(panel).prev();
    if (prev.length == 0) {
        return;
    }

    var panelBefore;
    panelBefore = prev[0];
    $(panel).insertBefore($(panelBefore));
}

function movePluginDown(event) {
    event = event ? event : window.event;
    var obj = event.srcElement ? event.srcElement : event.target;
    var obj = $(obj);
    var panel = obj.parents(".panel")[0];
    var next = $(panel).next();
    if (next.length == 0) {
        return;
    }

    var panelAfter;
    panelAfter = next[0];
    $(panel).insertAfter($(panelAfter));
}

function deletePlugin(event) {
    event = event ? event : window.event;
    var obj = event.srcElement ? event.srcElement : event.target;
    var obj = $(obj);
    var panel = obj.parents(".panel")[0];
    $(panel).remove();
}

function bindChange() {
    $("#myModalBody").find("input").keyup( function() {
        check('update');
    });
    $("#myModalBody").find("select").keyup( function() {
        check('update');
    });
    $("#myModalBody").find("textarea").keyup( function() {
        check('update');
    });
}

