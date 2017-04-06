package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"village/models"
	"village/models/admin"
)

type Admin struct {
	beego.Controller
}

func (self *Admin) Login() {
	self.TplName = "admin/submenu.html"
	self.Data["menu"] = admin.GetAllMenu()
	fmt.Println(admin.GetAllSubMenu())
	self.Data["submenu"] = admin.GetAllSubMenu()
}

//菜单页面
func (self *Admin) MenuManageList() {
	self.TplName = "admin/menu.html"
}

//用户页面
func (self *Admin) UserManageList() {
	self.TplName = "admin/user.html"
	self.Data["users"] = models.GetAllUser()
}

//评论页面
func (self *Admin) ReplyManageList() {
	self.TplName = "admin/reply.html"
	self.Data["replys"] = models.GetAllReply()
}

//分类页面
func (self *Admin) CategoryManageList() {
	self.TplName = "admin/category.html"
	self.Data["categorys"] = models.GetAllCategory()
}

//帖子页面
func (self *Admin) TopicManageList() {
	self.TplName = "admin/topic.html"
	self.Data["topics"] = models.GetAllTopic()
}

//菜单操作
func (self *Admin) MenuAction() {
	action := self.Ctx.Input.Param(":action")
	fmt.Println(action, "action")
	key, title := self.Input().Get("key"), self.Input().Get("title")
	switch action {
	case "add":
		menu := admin.Menu{Key: key, Title: title}
		admin.AddMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "添加成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "modify":
		menu := admin.GetMenuByKey(key)
		menu.Key = key
		menu.Title = title
		admin.UpdateMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "修改成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		menu := admin.GetMenuByKey(key)
		admin.DelMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "删除成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

//子菜单操作
func (self *Admin) SubMenuAction() {
	action := self.Ctx.Input.Param(":action")
	fmt.Println(action, "action")
	parent, key, title, url := self.Input().Get("parent"), self.Input().Get("key"), self.Input().Get("title"), self.Input().Get("url")
	switch action {
	case "add":
		submenu := admin.SubMenu{Key: key, Title: title, Url: url, Parent: &admin.Menu{Key: parent}}
		admin.AddSubMenu(&submenu)
		msg := map[string]interface{}{"code": 0, "msg": "添加成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "modify":
		menu := admin.GetMenuByKey(key)
		menu.Key = key
		menu.Title = title
		admin.UpdateMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "修改成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		menu := admin.GetMenuByKey(key)
		admin.DelMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "删除成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}
