package controllers

import (
	"github.com/astaxie/beego"
	"tester/models"
)

type ContentController struct {
	beego.Controller
}

func (self *ContentController) Get() {
	self.TplName = "index.html"

	sess_uid := self.GetSession("uid")
	if sess_uid == nil {
		self.Data["islogin"] = false
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(sess_uid.(int))
	}
	self.Data["newstopics"] = models.NewTopic()
	self.Data["waitreplys"] = models.WaitReply()
	self.Data["newreplys"] = models.NewReply()
	self.Data["Uptopics"] = models.UpTopicList()
	self.Data["Categorys"] = models.GetAllCategory()
	self.Data["NewUser"] = models.NewUser()
}
