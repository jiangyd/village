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
	sessionid := self.GetSession("uid")
	if sessionid == nil {
		msg := map[string]interface{}{"code": 1, "msg": "need loging"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		userbid, _ := strconv.Atoi(self.Input().Get("uid"))
		if sessionid.(int) == userbid {
			msg := map[string]interface{}{"code": 1, "msg": "不能关注自己"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			firend := models.Firend{UserA: &models.User{Id: sessionid.(int)}, UserB: &models.User{Id: userbid}}
			//判断是否关注关系,如果是就删除关系，不是就添加关系
			if models.IsFirend(&models.User{Id: sessionid.(int)}, &models.User{Id: userbid}) {
				models.FirendDel(&models.User{Id: sessionid.(int)}, &models.User{Id: userbid})
				msg := map[string]interface{}{"code": 0, "msg": "取消关注成功"}
				self.Data["json"] = &msg
				self.ServeJSON()
			} else {
				models.FirendAdd(&firend)
				msg := map[string]interface{}{"code": 0, "msg": "关注成功"}
				self.Data["json"] = &msg
				self.ServeJSON()
			}

		}

	}
}
