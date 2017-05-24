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

    form.on('submit(logingo)', function(data) {

        $.ajax({
            async: false,
            url: "/user/login",
            data: {
                "email": data.field.email,
                "password": data.field.password,
                "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id
            },
            type: 'POST',
            success: function(text) {
                if (text.code == 0) {
                    if(text.mfa==true){
                        location.href = '/user/mfaverify'
                    }else{
                        location.href = '/'
                    }
                    
                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }

        });

        return false;

    });





    exports('login', {});
});
