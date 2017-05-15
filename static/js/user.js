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
                    setTimeout(function() { location.href = '/user/set' }, 1000);
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
    });
    form.on('submit(modify-pwd)',function(data){
        if(data.field.newpassword!=data.field.repassword){
            layer.msg("两次密码输入不一致!")
            return false;
        }
        $.ajax({
            async:false,
            url:"/user/updatepwd",
            data:{"oldpassword":data.field.oldpassword,
                  "newpassword":data.field.newpassword},
            type:'POST',
            success:function(text){
                if(text.msg=='success'){
                    
                    layer.msg('修改成功,请重新登陆!')
                    setTimeout(function (){location.href='/user/login'}, 3000);
                    
                }else if(text.code!=0){
                    layer.msg(text.msg)

                }
            }


        });
        
        return false;

    });

    form.on('submit(forgetgo)',function(data){
        $.ajax({
            async:false,
            url:"/user/forgetpwd",
            data:{"email":data.field.email,
            "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id
        },
            type:'POST',
            success:function(text){
                if(text.msg=='success'){
                    
                    layer.msg('已发送找回密码邮件,前往重置')
                }else if(text.code!=0){
                    layer.msg(text.msg)

                }
            }


        });
        
        return false;

    });




    exports('login', {});
});
