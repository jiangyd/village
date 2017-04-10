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
//屏蔽帖子
function disabletopic(id){
    layer.confirm("屏蔽此贴后将不可见,是否确认？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/topic/disable",
                data: {
                   "id":id
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/topicmanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}
//取消屏蔽帖子
function enabletopic(id){
    layer.confirm("取消屏蔽此贴后将可见,是否确认？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/topic/enable",
                data: {
                   "id":id
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/topicmanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}
//屏蔽评论
function disablereply(id){
    layer.confirm("屏蔽此评论后将不可见,是否确认？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/reply/disable",
                data: {
                   "id":id
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/replymanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}
//取消屏蔽评论
function enablereply(id){
    layer.confirm("取消屏蔽此评论后将可见,是否确认？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/reply/enable",
                data: {
                   "id":id
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/replymanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}

//禁用用户
function disableuser(id){
    layer.confirm("禁用此用户将不能登录,是否确认？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/user/disable",
                data: {
                   "id":id
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/usermanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}
//启用用户
function enableuser(id){
    layer.confirm("启用此用户将恢复登录,是否确认？", {
        btn: ["是", "否"],
        yes: function() {
            $.ajax({
                type: "post",
                url: "/user/enable",
                data: {
                   "id":id
                },
                success: function(data) {
                    if (data.code == 0) {
                        location.href = '/admin/usermanagelist';

                    } else if (data.code != 0) {
                        layer.msg(data.msg);
                    };
                }
            })
        }
    })
}

