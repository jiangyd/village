layui.define(['layer', 'layedit', 'form'], function(exports) {
    var layer = layui.layer
    var layedit = layui.layedit
    var $ = layui.jquery
    var form = layui.form()
    // layedit.set({
    //     uploadImage: {
    //         url: '/topicupload' //接口url
    //             ,
    //         type: 'POST' //默认post
    //     }
    // });







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
