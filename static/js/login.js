layui.define(['layer','form'],function(exports){
	var layer=layui.layer
	var form=layui.form()
	var $ = layui.jquery

	form.verify({
		password:[/(.+){6,12}$/, '密码必须6到12位'],
	});

	form.on('submit(logingo)',function(data){
		
		$.ajax({
			async:false,
			url:"/user/login",
			data:{"email":data.field.email,
				  "password":data.field.password,
				  "vercode":data.field.vercode,
				  "captcha_id":data.field.captcha_id},
			type:'POST',
			success:function(text){
				if(text.msg=='success'){
					layer.msg('欢迎登录')
					setTimeout(function (){location.href='/'}, 3000);
					
				}else if(text.code!=0){
					layer.msg(text.msg)
				}
			}

		});
		
		return false;

	});

	



	exports('login',{});
});