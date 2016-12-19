/**
 * Created by junxian on 16/10/18.
 */
var staticDockerfile = "static_dockerfile";
var downloadExecutableSource = "download_executable_source";
var downloadDockerfile = "download_dockerfile";

function defineFileInit() {
    initDivSourceCode();
    initExtensionPlugins();
    initProjectName();
    bindChangeDefineType();
    bindAllChange();
}

function initProjectName() {
    if($.trim($("#project").val()) != defaultProjectName) {
        showPluginWithDefineTypeChange();
        $("#project").attr("readonly", "readonly");
        $("#div_StaticDockerfile_btn").remove();
        $("#div_DownloadDockerfile_btn").remove();
        $("#div_DownloadExecutableSource_btn").remove();
    } else {
        $("#project").val("");
        handleDefineTypeChange();
        $("#addOrUpdate").val("add");
    }
}
function initDivSourceCode() {
    if ($.trim($("#div_SourceCode").html()) == "") {
        $("#div_SourceCode").empty();
        addStaticPlug(downloadExecutableSource, "div_SourceCode");
    }
    $(document).find("input[name='$$plugin'][value='download_executable_source']:gt(0)").parent().remove();
}
function initExtensionPlugins() {
    $("#dockerfile_extension_plugins").find("option").each(function () {
        if ($(this).val()==staticDockerfile
            || $(this).val()==downloadDockerfile
            || $(this).val()==downloadExecutableSource) {
            $(this).remove();
        }
    })
}

function bindChangeDefineType() {
    $("#DefineDockerFileType").change(function () {
        setTitleWithDefineTypeChange();
        handleDefineTypeChange();
    });
    setTitleWithDefineTypeChange();
}

function setTitleWithDefineTypeChange() {
    var dType = $("#DefineDockerFileType").val();
    if (dType == "tool") {
        setTitle("");
    } else if (dType == "edit") {
        setTitle(staticDockerfile);
    } else if (dType == "download") {
        setTitle(downloadDockerfile);
    }
}

function handleDefineTypeChange() {
    var dType = $("#DefineDockerFileType").val();
    if (dType == "tool") {
        handleDefineTypeIsTool();
    } else if (dType == "edit") {
        handleDefineTypeIsEdit();
    } else if (dType == "download") {
        handleDefineTypeIsDownload();
    }
}

function showPluginWithDefineTypeChange() {
    var dType = $("#DefineDockerFileType").val();
    if (dType == "tool") {
        showTools();
    } else {
        hideTools();
    }
}

function handleDefineTypeIsTool() {
    showTools();
    cleanUpdateDivId();
    $("#dockerfile_extension_plugins").val('base_image_choose_from_harbor');
    addExtensionPlug(server, 'dockerfile_extension_plugins', updateDivId, 1);
}

function handleDefineTypeIsDownload() {
    hideTools();
    cleanUpdateDivId();
    addStaticPlug(downloadDockerfile, "");
}

function handleDefineTypeIsEdit() {
    hideTools();
    cleanUpdateDivId();
    addStaticPlug(staticDockerfile, "");
}

function showTools() {
    $("#div_dockerfile_extension_plugins").show();
    $("#div_dockerfile_extension_plugins").addClass("form-group");
}

function hideTools() {
    $("#div_dockerfile_extension_plugins").hide();
}

function cleanUpdateDivId() {
    var updateIdEscape = updateDivId.replace(/\ /g,"\\ ");
    $("#" + updateDivId).empty();
    $("#" + updateIdEscape).empty();
}

function appendHtml(result, parentDivID, plugin) {
    setTitle(plugin);
    if (parentDivID == "div_SourceCode") {
        $("#" + parentDivID).append(result);
    } else {
        var updateIdEscape = updateDivId.replace(/\ /g,"\\ ");
        $("#" + updateDivId).append(result);
        $("#" + updateIdEscape).append(result);
    }
}

function setTitle(plugin) {
    if (plugin == staticDockerfile) {
        $("#pluginTitle").html("直接编辑Dockerfile");
    } else if (plugin == downloadDockerfile) {
        $("#pluginTitle").html("下载Dockerfile");
    } else {
        $("#pluginTitle").html("操作列表");
    }
}

function addStaticPlug(plugin, parentDivID) {
    $.ajax({
        type: "GET",
        url: pluginViewUrl,
        async:false,
        data: {type: 1, plugin: plugin},
        success: function(result){
            appendHtml(result, parentDivID, plugin);
            $.getScript(server + "/static/js/custom/" + plugin + ".js");
            bindAllChange();
        },
        error: function(result){
            alert(result);
        }
    });
}

function bindAllChange() {
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



defineFileInit();