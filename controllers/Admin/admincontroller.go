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
	self.TplName = "admin/menu.html"
	self.Data["menu"] = admin.GetAllMenu()
	fmt.Println(admin.GetAllSubMenu())
	self.Data["submenu"] = admin.GetAllSubMenu()
}

func (self *Admin) Menu() {
	self.TplName = "admin/menu.html"
}
