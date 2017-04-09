layui.define(['layer', 'form', 'element'], function(exports) {
    var layer = layui.layer;
    var elements = layui.element();
    var $ = layui.jquery;
    var form = layui.form();

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
                    location.href = '/admin/menumanagelist'

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }
        });
        return false;
    });

    //添加子菜单form
    form.on('submit(addsubmenu)', function(data) {
        $.ajax({
            async: false,
            url: "/submenu/add",
            data: {
                "parent": data.field.parent,
                "key": data.field.key,
                "title": data.field.title,
                "url": data.field.url
            },
            type: 'POST',
            success: function(text) {
                if (text.code == 0) {
                    location.href = '/admin/submenumanagelist'

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }
        });
        return false;
    });
    //添加分类
        form.on('submit(addcategory)', function(data) {
        $.ajax({
            async: false,
            url: "/category/add",
            data: {
                "categorytype": data.field.categorytype,
                "category": data.field.category
            },
            type: 'POST',
            success: function(text) {
                if (text.code == 0) {
                    location.href = '/admin/categorymanagelist'

                } else if (text.code != 0) {
                    layer.msg(text.msg)
                }
            }
        });
        return false;
    });

    //添加菜单弹出
    $("#addmenu").on('click', function() {
        layer.open({
            type: 1,
            title: "添加菜单",
            area: ['500px', '250px'],
            content: '<form class="layui-form" method="post" style="margin:20px 20px 0px"><div class="layui-form-item"><label class="layui-form-label">菜单key</label><div class="layui-input-block"><input type="text" name="key" required  lay-verify="required" placeholder="请输入key" autocomplete="off" class="layui-input"></div></div><div class="layui-form-item"><label class="layui-form-label">菜单标题</label><div class="layui-input-block"><input type="text" name="title" required  lay-verify="required" placeholder="请输入菜单标题" autocomplete="off" class="layui-input"></div></div><button class="layui-btn layui-btn-normal" lay-filter="addmenu" lay-submit type="button" style="float: right;">确认</button></form>'
        });
    });

    //添加子菜单弹出
    $("#addsubmenu").on('click', function() {
        layer.open({
            type: 1,
            title: "添加子菜单",
            area: ['500px', '400px'],
            content: $("#div_addsubmenu")
        });
        form.render('select');
    });

    //添加分类单弹出
    $("#addcategory").on('click', function() {
        layer.open({
            type: 1,
            title: "添加分类",
            area: ['500px', '260px'],
            content: $("#div_addcategory")
        });
        form.render('select');
    });


    exports('admin', {});
});


//修改菜单
function modifymenu(key) {
    layer.open({
        type: 2,
        title: "修改菜单",
        resize: false,
        area: ['500px', '260px'],
        content: ["/getmenuinfo?key=" + key]
    });
}
//删除菜单
function delmenu(key) {
    layer.confirm("是否确认删除？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/menu/del",
                data: {
                    key: key
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/menumanagelist/';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}

//修改子菜单
function modifysubmenu(key) {
    layer.open({
        type: 2,
        title: "修改子菜单",
        resize: false,
        area: ['500px', '360px'],
        content: ["/getsubmenuinfo?key=" + key]
    });
}
//删除子菜单
function delsubmenu(key) {
    layer.confirm("是否确认删除？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/submenu/del",
                data: {
                    "key": key
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/submenumanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}

//修改分类
function modifycategory(id) {
    layer.open({
        type: 2,
        title: "修改分类",
        resize: false,
        area: ['500px', '260px'],
        content: ["/getcategoryinfo?categoryid="+id]
    });
}
//删除分类
function delcategory(id) {
    layer.confirm("是否确认删除？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/category/del",
                data: {
                   "categoryid":id
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/categorymanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}
