package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"village/models"
)

type FirendController struct {
	beego.Controller
}

func (self *FirendController) Firend() {
	useraid := self.GetSession("uid")
	userbid, _ := strconv.Atoi(self.Input().Get("uid"))
	if useraid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		if useraid.(int) == userbid {
			msg := map[string]interface{}{"code": 1, "msg": "不能关注自己"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}
		firend := models.Firend{UserA: &models.User{Id: useraid.(int)}, UserB: &models.User{Id: userbid}}
		models.FirendAdd(&firend)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
}
