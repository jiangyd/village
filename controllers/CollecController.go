package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"village/models"
)

//收藏
type CollecController struct {
	beego.Controller
}

func (self *CollecController) Collec() {
	t, typeid := self.Input().Get("type"), self.Input().Get("typeid")
	tpid, _ := strconv.Atoi(typeid)
	uid := self.GetSession("uid")
	collec := models.Collection{Type: t, TypeId: tpid, Uid: &models.User{Id: uid.(int)}}
	if models.AddCollection(&collec) >= 0 {
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		fmt.Println(collec)
		models.DelCollection(&collec)
		msg := map[string]interface{}{"code": 0, "msg": "取消收藏成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}
