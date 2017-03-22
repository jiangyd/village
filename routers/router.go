package routers

import (
	"github.com/astaxie/beego"
	"village/controllers"
)

func init() {
	beego.Router("/", &controllers.ContentController{}, "GET:Index")
	beego.Router("/user/login", &controllers.UserController{}, "GET:LoginPage")
	beego.Router("/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/user/logout", &controllers.UserController{}, "GET:Logout")
	beego.Router("/user/forget", &controllers.UserController{}, "GET:Forget")
	beego.Router("/user/set", &controllers.UserController{}, "GET:Set")      //设置信息页面
	beego.Router("/user/set", &controllers.UserController{}, "POST:SetInfo") //修改基本信息
	beego.Router("/user/message", &controllers.UserController{}, "GET:Message")
	beego.Router("/user/topic", &controllers.UserController{}, "GET:UserTopic")
	beego.Router("/user/collection", &controllers.UserController{}, "GET:Collection")
	beego.Router("/collection", &controllers.CollecController{}, "POST:Collec")
	beego.Router("/user/follow", &controllers.UserController{}, "GET:FollowPage")
	beego.Router("/user/firend", &controllers.FirendController{}, "POST:Firend")
	beego.Router("/user/register", &controllers.UserController{}, "GET:RegisterPage")
	beego.Router("/user/register", &controllers.UserController{}, "POST:Register")
	beego.Router("/user/imgupload", &controllers.UploadImg{}, "POST:Upload")
	beego.Router("/user/updatepwd", &controllers.UserController{}, "POST:UpdatePwd")
	beego.Router("/user/detial/:uid([0-9]+)", &controllers.UserController{}, "GET:UserDetial")
	beego.Router("/capt", &controllers.Capt{})
	beego.Router("/topic/:id([0-9]+)", &controllers.TopicController{}, "GET:TopicDetial")
	beego.Router("/topic/create", &controllers.TopicController{}, "GET:CreatePage")
	beego.Router("/topic/create", &controllers.TopicController{}, "POST:CreateTopic")
	beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "GET:EditPage")
	beego.Router("/topic/edit/", &controllers.TopicController{}, "POST:EditTopic")
	beego.Router("/topic/reply", &controllers.TopicController{}, "POST:ReplyTopic")
	beego.Router("/dz", &controllers.DzController{}, "POST:Dz")
	beego.Router("/topic/adopt", &controllers.TopicController{}, "POST:Adopt")
	beego.Router("/topic/:type([a-z]+)", &controllers.ContentController{}, "GET:TopicList")
	// beego.Router("/topic/waitreply", &controllers.ContentController{}, "GET:WaitReply")
	// beego.Router("/topic/newreply", &controllers.ContentController{}, "GET:NewReply")
	// beego.Router("/topic/up", &controllers.ContentController{}, "GET:UpTopicList")
	// beego.Router("/topic/adopt", &controllers.ContentController{}, "GET:AdoptTopicList")

	beego.Router("/site", &controllers.SiteController{}, "GET:SitePage")

	//beego.Router("/wd/index", &controllers.ContentController{}, "GET:Index")
}
