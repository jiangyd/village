package admin

import (
	"github.com/astaxie/beego"
)

type Admin struct {
	beego.Controller
}

func (self *Admin) Login() {
	self.TplName = "admin/login.html"
}
