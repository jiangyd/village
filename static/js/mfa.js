layui.define(['layer', 'form', 'element', 'upload'], function(exports) {
    var layer = layui.layer
    var form = layui.form()
    var elements = layui.element()
    var $ = layui.jquery


    elements.on('mfa', function(data) {
        elements.tabChange('mfa', data.index);
    });



    form.on('submit(mfa_m)', function(data) {
        $.ajax({
            async: false,
            url: "/user/setmfa",
            data: {
                "code1": data.field.mcode1,
                "code2": data.field.mcode2

            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {

                    layer.msg('MFA开启成功')
                    location.href="/"
                } else if (text.code != 0) {
                    layer.msg(text.msg)

                }
            }


        });

        return false;

    });
    form.on('submit(mfa_a)', function(data) {
        $.ajax({
            async: false,
            url: "/user/setmfa",
            data: {
                "code1": data.field.acode1,
                "code2": data.field.acode2

            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {

                    layer.msg('MFA开启成功')
                    location.href="/"
                } else if (text.code != 0) {
                    layer.msg(text.msg)

                }
            }


        });

        return false;

    });
    form.on('submit(mfa_close)', function(data) {
        $.ajax({
            async: false,
            url: "/user/closemfa",
            data: {
                "code1": data.field.acode1,
                "code2": data.field.acode2

            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {

                    layer.msg('MFA开启成功')
                    location.href="/"
                } else if (text.code != 0) {
                    layer.msg(text.msg)

                }
            }


        });

        return false;

    });

        form.on('submit(mfasend)', function(data) {
        $.ajax({
            async: false,
            url: "/user/mfaverify",
            data: {
                "code": data.field.code
            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {
                    location.href = '/'
                } else if (text.code != 0) {
                    layer.msg(text.msg+"aa")

                }
            }


        });

        return false;

    });




    exports('mfa', {});
});
