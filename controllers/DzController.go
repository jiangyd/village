package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"village/models"
)

type DzController struct {
	beego.Controller
}

func (self *DzController) Dz() {
	t, typeid := self.Input().Get("type"), self.Input().Get("typeid")
	tpid, _ := strconv.Atoi(typeid)
	uid := self.GetSession("uid")
	dz := models.Dz{Type: t, TypeId: tpid, Uid: &models.User{Id: uid.(int)}}
	if models.Adddz(&dz) >= 0 {
		switch t {
		case "tid":
			//查找主题id,给主题点赞数加1
			topic := models.FindTopicById(tpid)
			models.UpTopic(&topic)
		case "rid":
			//查找评论id,给评论点赞数加1
			reply := models.FindReplyByRid(tpid)
			models.UpReply(&reply)

		}
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		msg := map[string]interface{}{"code": 1, "msg": "不能重复点赞"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}
