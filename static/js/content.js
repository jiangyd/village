layui.define(['layer', 'layedit', 'form'], function(exports) {
    var layer = layui.layer
    var layedit = layui.layedit
    var $ = layui.jquery
    var form = layui.form()
    layedit.set({
        uploadImage: {
            url: '' //接口url
                ,
            type: '' //默认post
        }
    });

    //发表主题的富文本编辑器
    var topic_content = layedit.build('topic_content', {
        tool: ['strong', 'face', 'image', 'link', 'unlink', 'code']
    });

    //发表回复的富文本编辑器
    var reply_content = layedit.build('reply_content', {
        tool: ['strong', 'face', 'image', 'link', 'unlink', 'code'],
        height: 180,
    });

    //编辑主题富文本编辑器
    var edittopic_content = layedit.build('edittopic_content', {
        tool: ['strong', 'face', 'image', 'link', 'unlink', 'code'],
       
    });


    form.on('submit(topicgo)', function(data) {
        $.ajax({
            async: false,
            url: "/topic/create",
            data: {
                "title": data.field.title,
                "content": layedit.getContent(topic_content),
                "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id
            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {
                    location.href = '/topic/' + text.tid

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }

        });

        return false;

    });

    form.on('submit(replygo)', function(data) {
        $.ajax({
            async: false,
            url: "/topic/reply",
            data: {
                "topic_id": data.field.topic_id,
                "title": data.field.title,
                "content": layedit.getContent(reply_content),
                "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id
            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {
                    location.href = '/topic/' + text.tid

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }

        });

        return false;

    });

    form.on('submit(edittopicgo)', function(data) {
        $.ajax({
            async: false,
            url: "/topic/edit",
            data: {
                "topic_id": data.field.topic_id,
                "title":data.field.title,
                "content": layedit.getContent(edittopic_content),
                "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id
            },
            type: 'POST',
            success: function(text) {
                if (text.msg == 'success') {
                    location.href = '/topic/' + text.tid
                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }

        });

        return false;

    });




    exports('content', {});
});
