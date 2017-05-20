layui.define(['layer', 'form'], function(exports) {
    var layer = layui.layer
    var form = layui.form()
    var $ = layui.jquery
    // var element = layui.element();


    // element.on('user', function(data) {
    //     element.tabChange('user', data.index);
    // });

    form.verify({
        password: [/(.+){6,12}$/, '密码必须6到12位'],
    });

    form.on('submit(forgetpwd)', function(data) {
         if (data.field.password != data.field.repassword) {
            layer.msg("两次密码输入不一致!")
            return false;
        }
        $.ajax({
            async: false,
            url: "/user/setnewpwd",
            data: {
              
                "password": data.field.password,
                "repassword":data.field.password,
               "uuid":data.field.uuid,
            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {
                    location.href = '/'
                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }

        });

        return false;

    });





    exports('forgetpwd', {});
});
