layui.define(['layer', 'form', 'element','tree'], function(exports) {
    var layer = layui.layer;
    var elements = layui.element();
    var $ = layui.jquery;
    var form = layui.form();
    exports('document', {});
        //添加分类单弹出
    $("#addnode").on('click', function() {
        layer.open({
            type: 1,
            title: "添加分类",
            area: ['500px', '260px'],
            content: $("#div_addnode")
        });
        form.render('select');
    });
    //添加子菜单form
    form.on('submit(addnode)', function(data) {
        $.ajax({
            async: false,
            url: "/docnode/add",
            data: {
                
                "pid": "",
                "node": data.field.node,
                "content": data.field.content
            },
            type: 'POST',
            success: function(text) {
                if (text.code == 0) {
                    location.href = '/admin/submenumanagelist'

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }
        });
        return false;
    });
});
