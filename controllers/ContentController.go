package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"village/models"
)

type ContentController struct {
	beego.Controller
}

func (self *ContentController) Index() {
	self.TplName = "index.html"
	sess_uid := self.GetSession("uid")
	self.Data["isnewtopic"] = true
	if sess_uid == nil {
		self.Data["islogin"] = false
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(sess_uid.(int))
	}
	self.Data["newstopics"] = models.NewTopic()
	self.Data["Categorys"] = models.GetTopicCategory()
	self.Data["NewUser"] = models.NewUser()
	self.Data["HotUser"] = models.FindHotUser()

}

func (self *ContentController) TopicList() {
	t := self.Ctx.Input.Param(":type")
	switch t {
	case "newtopic":
		fmt.Println("当时发生的发生")
		self.Data["newstopics"] = models.NewTopic()
		self.Data["isnewtopic"] = true
	case "waitreply":
		self.Data["waitreplys"] = models.WaitReply()
		self.Data["iswaitreply"] = true
		// self.Data["isnewtopic"] = false
	case "newreply":
		self.Data["newreplys"] = models.NewReply()
		self.Data["isnewreply"] = true
		// self.Data["isnewtopic"] = false
	case "uptopic":
		self.Data["Uptopics"] = models.UpTopicList()
		self.Data["isuptopic"] = true
		// self.Data["isnewtopic"] = false
	case "adopt":
		self.Data["adopts"] = models.AdoptTopicList()
		self.Data["isadopt"] = true
		// self.Data["isnewtopic"] = false
	}

	self.TplName = "index.html"
	sess_uid := self.GetSession("uid")

	if sess_uid == nil {
		self.Data["islogin"] = false
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(sess_uid.(int))
	}
	self.Data["Categorys"] = models.GetTopicCategory()
	self.Data["NewUser"] = models.NewUser()
	self.Data["HotUser"] = models.FindHotUser()

}

func (self *ContentController) WaitReply() {

	self.TplName = "index.html"
	sess_uid := self.GetSession("uid")

	if sess_uid == nil {
		self.Data["islogin"] = false
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(sess_uid.(int))
	}
	self.Data["waitreplys"] = models.WaitReply()
	self.Data["Categorys"] = models.GetTopicCategory()
	self.Data["NewUser"] = models.NewUser()
	self.Data["HotUser"] = models.FindHotUser()

}

func (self *ContentController) NewReply() {

	self.TplName = "index.html"
	sess_uid := self.GetSession("uid")

	if sess_uid == nil {
		self.Data["islogin"] = false
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(sess_uid.(int))
	}
	self.Data["newreplys"] = models.NewReply()
	self.Data["Categorys"] = models.GetTopicCategory()
	self.Data["NewUser"] = models.NewUser()
	self.Data["HotUser"] = models.FindHotUser()

}

func (self *ContentController) UpTopicList() {

	self.TplName = "index.html"
	sess_uid := self.GetSession("uid")

	if sess_uid == nil {
		self.Data["islogin"] = false
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(sess_uid.(int))
	}

	self.Data["Uptopics"] = models.UpTopicList()
	self.Data["Categorys"] = models.GetTopicCategory()
	self.Data["NewUser"] = models.NewUser()
	self.Data["HotUser"] = models.FindHotUser()

}
func (self *ContentController) AdoptTopicList() {

	self.TplName = "index.html"
	sess_uid := self.GetSession("uid")

	if sess_uid == nil {
		self.Data["islogin"] = false
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(sess_uid.(int))
	}
	self.Data["Categorys"] = models.GetTopicCategory()
	self.Data["NewUser"] = models.NewUser()
	self.Data["HotUser"] = models.FindHotUser()
	self.Data["adopts"] = models.AdoptTopicList()

}
