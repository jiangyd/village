package admin

import (
	"fmt"
	"github.com/astaxie/beego"
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

func (self *Admin) Menu() {
	self.TplName = "admin/menu.html"
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
