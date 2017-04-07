layui.define(['layer','form','element'],function(exports){
	var layer=layui.layer;
    var elements=layui.element();
	var $ = layui.jquery;
	var form=layui.form();

	//添加菜单form
	form.on('submit(addmenu)', function(data) {
        $.ajax({
            async: false,
            url: "/menu/add",
            data: {
                "key": data.field.key,
                "title": data.field.title,
            },
            type: 'POST',
            success: function(text) {
                if (text.code == 0) {
                    location.href = '/menu/'

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }
        });
        return false;
    });

    //添加子菜单form
        //添加菜单form
    form.on('submit(addsubmenu)', function(data) {
        $.ajax({
            async: false,
            url: "/submenu/add",
            data: {
                "parent":data.field.parent,
                "key": data.field.key,
                "title": data.field.title,
                "url":data.field.url
            },
            type: 'POST',
            success: function(text) {
                if (text.code == 0) {
                    location.href = '/menu/'

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }
        });
        return false;
    });

    //添加菜单弹出
    $("#addmenu").on('click',function(){
    	layer.open({
    		type:1,
    		title:"添加菜单",
    		area: ['500px', '250px'],
    		content:'<form class="layui-form" method="post" style="margin:20px 20px 0px"><div class="layui-form-item"><label class="layui-form-label">菜单key</label><div class="layui-input-block"><input type="text" name="key" required  lay-verify="required" placeholder="请输入key" autocomplete="off" class="layui-input"></div></div><div class="layui-form-item"><label class="layui-form-label">菜单标题</label><div class="layui-input-block"><input type="text" name="title" required  lay-verify="required" placeholder="请输入菜单标题" autocomplete="off" class="layui-input"></div></div><button class="layui-btn layui-btn-normal" lay-filter="addmenu" lay-submit type="button" style="float: right;">确认</button></form>'
    	});
    });

    //添加子菜单弹出
    $("#addsubmenu").on('click',function(){
    	layer.open({
    		type:1,
    		title:"添加子菜单",
    		area: ['500px', '400px'],
    		content:$("#div_addsubmenu")
    	});
        form.render('select');
    });


	exports('admin',{});
});

//修改子菜单
    function modifysubmenu(key){
        $.get("/getsubmenuinfo?key="+key,function(data){
            alert(data.data.key)
        })
        layer.open({
            type:1,
            title:"修改子菜单",
            area: ['500px', '400px'],
            content:$("#div_addsubmenu")
        });
    }