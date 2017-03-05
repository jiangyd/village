layui.define(['layer','form','element'],function(exports){
	var layer=layui.layer
	var form=layui.form()
	var elements=layui.element()
	var $ = layui.jquery

	elements.use('element',function(){
		var element = layui.element();
		element.on('user',function(data){
			element.tabChange('user',data.index);
		});
		
	});

	form.on('submit()',function(data){

	})
	



	exports('login',{});
});