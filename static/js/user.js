layui.define(['layer', 'form', 'element', 'upload'], function(exports) {
    var layer = layui.layer
    var form = layui.form()
    var elements = layui.element()
    var $ = layui.jquery
    

    elements.on('user', function(data) {
        elements.tabChange('user', data.index);
    });

    //头像设置上传
    layui.upload({
            url: "/user/imgupload",
            success: function(text) {
                if (text.msg == 'success') {
                    // setTimeout(function() { location.href = '/user/set' }, 1000);
                    layer.msg("cc")
                } else if (text.code != 0) {
                    layer.msg(text.msg)

                }
            }

        }

    )



    form.on('submit(setinfo)', function(data) {
        $.ajax({
            async: false,
            url: "/user/set",
            data: {
                // "email": data.field.Email,
                "nickname": data.field.Nickname,
                "sex": data.field.sex,
                "city": data.field.city,
                "sign": data.field.sign
            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {
                    setTimeout(function() { location.href = '/user/set' }, 1000);
                } else if (text.code != 0) {
                    layer.msg(text.msg)

                }
            }

        })
        return false;
    })




    exports('login', {});
});
