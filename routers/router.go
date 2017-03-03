package routers

import (
	"github.com/astaxie/beego"
	"village/controllers"
)

func init() {
	beego.Router("/", &controllers.ContentController{})
	beego.Router("/user/login", &controllers.UserController{}, "GET:LoginPage")
	beego.Router("/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/user/forget", &controllers.UserController{}, "GET:Forget")
	beego.Router("/user/:uid([0-9]+),", &controllers.UserController{}, "GET:Detial")
	beego.Router("/user/set", &controllers.UserController{}, "GET:Set")
	beego.Router("/user/register", &controllers.UserController{}, "GET:RegisterPage")
	beego.Router("/user/detial", &controllers.UserController{}, "GET:Detial")
	beego.Router("/user/register", &controllers.UserController{}, "POST:Register")
	beego.Router("/capt", &controllers.Capt{})
	beego.Router("/topic/:id([0-9]+)", &controllers.TopicController{}, "GET:TopicDetial")
	beego.Router("/topic/create", &controllers.TopicController{}, "GET:CreatePage")
	beego.Router("/topic/create", &controllers.TopicController{}, "POST:CreateTopic")
	beego.Router("/topic/reply", &controllers.TopicController{}, "POST:ReplyTopic")
	//beego.Router("/wd/index", &controllers.ContentController{}, "GET:Index")
}
