layui.define(['layer', 'layedit', 'form'], function(exports) {
    var layer = layui.layer
    var layedit = layui.layedit
    var $ = layui.jquery
    var form = layui.form()
    layedit.set({
        uploadImage: {
            url: '/topicupload' //接口url
                ,
            type: 'POST' //默认post
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
        layer.msg(data.field.category)
        $.ajax({
            async: false,
            url: "/topic/create",
            data: {
                "title": data.field.title,
                "content": layedit.getContent(topic_content),
                "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id,
                "category": data.field.category
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
                "title": data.field.title,
                "content": layedit.getContent(edittopic_content),
                "vercode": data.field.vercode,
                "captcha_id": data.field.captcha_id,
                "category": data.field.category
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
    //私信弹出
    $('#message').on('click',function(){
        layer.open({
        type: 0,
        title:"发送私信",
        offset: '100px',
        btn:['发送'],
  

        content: '<textarea id="sixin" name="sixin" placeholder="请输入内容" class="layui-textarea"></textarea>' ,//这里content是一个普通的String
              yes:function(){
                layer.msg("cc")
               $.ajax({
                async:false,
                url:"/user/message",
                data:{
                userb:$("#userb").val(),    
                content:$("#sixin").val()
                },
                type:'POST',
                success:function(text){
                    if(text.msg=='success'){
                        layer.msg("发送成功")
                    }
                }
                
               })
        }
    });
    })
    




    exports('content', {});
});
