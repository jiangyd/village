layui.define(['layer', 'form', 'element', 'tree'], function(exports) {
    var layer = layui.layer;
    var elements = layui.element();
    var $ = layui.jquery;
    var form = layui.form();


    // //添加根目录弹出
    // $("#addnode").on('click', function() {
    //     layer.open({
    //         type: 1,
    //         title: "添加根目录",
    //         area: ['500px', '260px'],
    //         content: $("#div_addnode")
    //     });
    //     form.render('select');
    // });
    //添加根目录
    // form.on('submit(addnode)', function(data) {
    //     $.ajax({
    //         async: false,
    //         url: "/docnode/add",
    //         data: {

    //             "pid": "",
    //             "node": data.field.node,
    //             "content": data.field.content
    //         },
    //         type: 'POST',
    //         success: function(text) {
    //             if (text.code == 0) {
    //                 location.href = '/document'

    //             } else if (text.code != 0) {
    //                 layer.msg(text.msg)
    //             }
    //         }
    //     });
    //     return false;
    // });

    //编辑目录
    $("#editnode").on('click', function() {
        var tnode = $("#tree").tree('getSelectedNode');
        location.href = "/editdocument/" + tnode.id
    });

    //添加子目录弹出
    $("#addsubnode").on('click', function() {
        var tnode = $("#tree").tree('getSelectedNode');
        location.href = "/adddocument/" + tnode.id
    });

    //删除目录
    $("#delnode").on('click', function() {
        var tnode = $("#tree").tree('getSelectedNode');
        $.ajax({
            url: "/docnode/del",
            data: {
                "id": tnode.id
            },
            type: "POST",
            success: function(text) {
                if (text.code == 0) {
                    location.href = '/document'

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }
        });

    });

    


    exports('document', {});
});
