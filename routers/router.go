package routers

import (
	"github.com/astaxie/beego"
	"village/controllers"
	"village/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.ContentController{}, "GET:Index")
	beego.Router("/user/login", &controllers.UserController{}, "GET:LoginPage")
	beego.Router("/user/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/user/logout", &controllers.UserController{}, "GET:Logout")
	beego.Router("/user/forget", &controllers.UserController{}, "GET:Forget")
	beego.Router("/user/mfa", &controllers.UserController{}, "GET:MFAPage") //mfa页面
	beego.Router("/user/setmfa", &controllers.UserController{}, "POST:SetMfa")
	beego.Router("/user/closemfa", &controllers.UserController{}, "POST:CloseMfa") //关闭mfa
	beego.Router("/user/mfaverify", &controllers.UserController{}, "GET:MfaVerifyPage")
	beego.Router("/user/mfaverify", &controllers.UserController{}, "POST:MfaVerify")
	beego.Router("/user/set", &controllers.UserController{}, "GET:Set")      //设置信息页面
	beego.Router("/user/set", &controllers.UserController{}, "POST:SetInfo") //修改基本信息
	beego.Router("/user/forgetpwd", &controllers.UserController{}, "POST:ForGetPwd")
	beego.Router("/forgetpwd", &controllers.UserController{}, "GET:ForGetPwdPage")
	beego.Router("/user/setnewpwd", &controllers.UserController{}, "POST:SetNewPwd") //重设密码
	beego.Router("/user/message", &controllers.UserController{}, "GET:Message")
	beego.Router("/user/message", &controllers.UserController{}, "POST:SendMsg") //发送私信
	beego.Router("/user/topic", &controllers.UserController{}, "GET:UserTopic")
	beego.Router("/user/collection", &controllers.UserController{}, "GET:Collection")
	beego.Router("/collection", &controllers.CollecController{}, "POST:Collec") //收藏
	beego.Router("/user/follow", &controllers.UserController{}, "GET:FollowPage")
	beego.Router("/user/firend", &controllers.FirendController{}, "POST:Firend")
	beego.Router("/user/register", &controllers.UserController{}, "GET:RegisterPage")
	beego.Router("/user/register", &controllers.UserController{}, "POST:Register")
	beego.Router("/user/imgupload", &controllers.UploadImg{}, "POST:Upload")
	beego.Router("/user/updatepwd", &controllers.UserController{}, "POST:UpdatePwd")
	beego.Router("/user/detial/:uid([0-9]+)", &controllers.UserController{}, "GET:UserDetial")
	beego.Router("/capt", &controllers.Capt{})                                        //获取验证码
	beego.Router("/topic/search/", &controllers.TopicController{}, "GET:TopicSearch") //通过分类查询帖子
	beego.Router("/topic/:id([0-9]+)", &controllers.TopicController{}, "GET:TopicDetial")
	beego.Router("/topic/create", &controllers.TopicController{}, "GET:CreatePage")
	beego.Router("/topic/create", &controllers.TopicController{}, "POST:CreateTopic")
	beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "GET:EditPage")
	beego.Router("/topic/edit/", &controllers.TopicController{}, "POST:EditTopic")
	beego.Router("/topic/reply", &controllers.TopicController{}, "POST:ReplyTopic")
	beego.Router("/dz", &controllers.DzController{}, "POST:Dz")
	beego.Router("/topic/adopt", &controllers.TopicController{}, "POST:Adopt") //采纳
	beego.Router("/topic/:type([a-z]+)", &controllers.ContentController{}, "GET:TopicList")
	beego.Router("/qiniucallback", &controllers.UploadImg{}, "POST:QiNiuCallBack")
	beego.Router("/topicupload", &controllers.UploadImg{}, "POST:TopicUpload")
	beego.Router("/site", &controllers.SiteController{}, "GET:SitePage")
	beego.Router("/wiki", &controllers.WiKi{}, "GET:WiKiList")
	beego.Router("/about", &controllers.ContentController{}, "GET:About")
	beego.Router("/wiki/:id([0-9]+)", &controllers.WiKi{}, "GET:WiKiRoot")
	beego.Router("/wiki/detial/:root([0-9]+)/:id([0-9]+)", &controllers.WiKi{}, "GET:WiKiDetial")
	beego.Router("/wiki/tree/:id([0-9]+)", &controllers.WiKi{}, "GET:WiKiRootTree")
	//admin
	beego.Router("/admin", &admin.Admin{}, "GET:Login")
	beego.Router("/admin", &admin.Admin{}, "GET:AdminPost")
	beego.Router("/menu/:action([a-z]+)", &admin.Admin{}, "POST:MenuAction")
	beego.Router("/submenu/:action([a-z]+)", &admin.Admin{}, "POST:SubMenuAction")
	beego.Router("/category/:action([a-z]+)", &admin.Admin{}, "POST:CategoryAction")
	beego.Router("/topic/:action([a-z]+)", &admin.Admin{}, "POST:TopicAction")
	beego.Router("/reply/:action([a-z]+)", &admin.Admin{}, "POST:ReplyAction")
	beego.Router("/user/:action([a-z]+)", &admin.Admin{}, "POST:UserAction")
	beego.Router("/site/:action([a-z]+)", &admin.Admin{}, "POST:SiteAction")
	beego.Router("/admin/menumanagelist", &admin.Admin{}, "GET:MenuManageList")
	beego.Router("/admin/submenumanagelist", &admin.Admin{}, "GET:SubMenuManageList")
	beego.Router("/admin/usermanagelist", &admin.Admin{}, "GET:UserManageList")
	beego.Router("/admin/topicmanagelist", &admin.Admin{}, "GET:TopicManageList")
	beego.Router("/admin/replymanagelist", &admin.Admin{}, "GET:ReplyManageList")
	beego.Router("/admin/categorymanagelist", &admin.Admin{}, "GET:CategoryManageList")
	beego.Router("/admin/sitemanagelist", &admin.Admin{}, "GET:SiteManageList")
	beego.Router("/admin/rolemanagelist", &admin.RolePermission{}, "GET:RoleManagelist")

	beego.Router("/getsubmenuinfo", &admin.Admin{}, "GET:GetSubMenuInfo")
	beego.Router("/getmenuinfo", &admin.Admin{}, "GET:GetMenuInfo")
	beego.Router("/getcategoryinfo", &admin.Admin{}, "GET:GetCategoryInfo")
	beego.Router("/getsiteinfo", &admin.Admin{}, "GET:GetSiteInfo")
	beego.Router("/document", &admin.Admin{}, "GET:DocumentInfo")
	beego.Router("/adddocument/:pid([0-9]{0,}?)", &admin.Admin{}, "GET:AddDocumentPage")
	beego.Router("/editdocument/:id([0-9]+)", &admin.Admin{}, "GET:EditDocumentPage")
	beego.Router("/getdocnodes", &admin.Admin{}, "GET:GetDocNodes")
	beego.Router("/docnode/:action([a-z]+)", &admin.Admin{}, "POST:DocNodeAction")

	beego.Router("/getrolenodes", &admin.RolePermission{}, "GET:GetRoleDocNodes")
}
