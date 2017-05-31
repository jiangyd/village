package controllers

import (
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
	if uid == nil {
		msg := map[string]interface{}{"code": 1, "msg": "need login"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
	if models.IsCollecExit(t, tpid, &models.User{Id: uid.(int)}) {
		//收藏存在就删除,collection表做了多字段关联唯一键
		models.DelCollection(t, tpid, &models.User{Id: uid.(int)})
		msg := map[string]interface{}{"code": 0, "msg": "取消收藏成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		//不存在就添加
		collec := models.Collection{Type: t, TypeId: tpid, Uid: &models.User{Id: uid.(int)}}
		models.AddCollection(&collec)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()

	}

}
