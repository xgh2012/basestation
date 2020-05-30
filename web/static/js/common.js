var layer = null,$ = null;
layui.use(['jquery', 'layer', 'miniAdmin','miniTongji'], function () {
    $ = layui.jquery,
    layer = layui.layer
});

function post(url,data) {
    $.ajax({
        type: "POST",
        dataType: "json",
        url: url,
        data: data,
        success: function(data){
            console.log(data);
            if (data.result == "success"){
                success()
            }
        },
        beforeSend:addLoading,
        complete:closeLoading,
        error:closeLoading
    });
}

function success() {
    layer.msg('提交成功', {icon: 1});
}
function addLoading() {
    layer.load()
}

function closeLoading() {
    layer.closeAll('loading');
}