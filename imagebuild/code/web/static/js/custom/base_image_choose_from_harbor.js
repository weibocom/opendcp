//　刷新列表
function refresh() {
    var harborAddress = $("#base_image_choose_from_harbor\\.harborAddress").val();
    var user = $("#base_image_choose_from_harbor\\.user").val();
    var password = $("#base_image_choose_from_harbor\\.password").val();
    $.ajax({
        url: extensionInterfaceUrl,
        method: "post",
        data: { plugin: "base_image_choose_from_harbor",
                method: "BaseImageList",
                harborAddress: harborAddress,
                user: user,
                password: password },
        success: function(result){
            var imageList = $("#base_image_choose_from_harbor\\.selectImage");
            // var images = JSON.parse(result);
            var code = result.code;
            if(code != 10000) return;
            var images = result.content;
            imageList.empty();
            for (var image in images) {
                imageList.append($("<option></option>").attr("value", images[image]).text(images[image]));
            }

            $("#base_image_choose_from_harbor\\.selectImage").change(function () {
                var val = $("#base_image_choose_from_harbor\\.selectImage").val();
                $("#base_image_choose_from_harbor\\.selectImage").parent().parent().next().find("input[name='base_image_choose_from_harbor.handleInput']").val(val);
            });
            var val = $("#base_image_choose_from_harbor\\.selectImage").val();
            $("#base_image_choose_from_harbor\\.selectImage").parent().parent().next().find("input[name='base_image_choose_from_harbor.handleInput']").val(val);
        },
        error: function(result){
            alert("call " + extensionInterfaceUrl +" error!");
        }
    });
}