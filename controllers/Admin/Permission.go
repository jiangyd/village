package admin

import (
	"github.com/astaxie/beego"
	"village/models/admin"
)

type RolePermission struct {
	beego.Controller
}

func (self *RolePermission) RoleManagelist() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.TplName = "admin/role.html"
	self.Layout = "admin/nav.html"
	self.Data["role"] = admin.GetAllRole()

}

//菜单操作
func (self *Admin) RoleAction() {
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
		fmt.Println(menu, "menu")
		admin.UpdateMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "修改成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		menu := admin.GetMenuByKey(key)
		if admin.IsHasSubMenu(&menu) {
			msg := map[string]interface{}{"code": 1, "msg": "该菜单下包含子菜单,不能删除!"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			admin.DelMenu(&menu)
			msg := map[string]interface{}{"code": 0, "msg": "删除成功"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}

	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}
