// 关注
function firendadd(uid) {
    $.ajax({
        url: "/user/firend",
        data: {
            "uid": uid
        },
        type: "POST",
        success: function(text) {
            if (text.msg == "关注成功") {
                $("#firend").text("取消关注")
            }else if (text.msg == "取消关注成功"){
            	$("#firend").text("关注")
            }else if (text.msg="need loging") {
                location.href = '/user/login'
            }
        }
    })
}

//点赞
function dz(tp, tpid) {
    $.ajax({
        url: "/dz",
        data: {
            "type": tp,
            "typeid": tpid
        },
        type: "POST",
        success: function(text) {
            if (text.msg == "success") {
                $("#dz").text("已点赞")
            } else if (text.code != 0) {
                layer.msg(text.msg)
            }
        }
    })
};
//收藏
function Collec(tp, tpid) {
    $.ajax({
        url: "/collection",
        data: {
            "type": tp,
            "typeid": tpid
        },
        type: "POST",
        success: function(text) {
            if (text.msg == "success") {
                $("#collection").text("取消收藏")
            } else if (text.msg == "取消收藏成功") {
                $("#collection").text("收藏")
            } else if (text.code != 0) {
                layer.msg(text.msg)
            }
        }
    })
};
