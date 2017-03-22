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

    form.on('submit(reggo)', function(data) {
        if (data.field.password != data.field.repassword) {
            layer.msg("两次密码输入不一致!")
            return false;
        }
        $.ajax({
            async: false,
            url: "/user/register",
            data: {
                "email": data.field.email,
                "nickname": data.field.nickname,
                "password": data.field.password,
                "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id
            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {

                    layer.msg('欢迎注册')
                    setTimeout(function() { location.href = '/user/login' }, 3000);

                } else if (text.code != 0) {
                    layer.msg(text.msg)

                }
            }


        });

        return false;

    });





    exports('register', {});
});
