package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"village/models"
)

type TopicController struct {
	beego.Controller
}

func (self *TopicController) TopicDetial() {
	id := self.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	if tid >= 0 {
		fmt.Println(self.GetSession("SessionId"))
		topic := models.FindTopicById(tid)
		models.IncrView(&topic)
		self.Data["topic"] = topic
		self.Data["replyinfo"] = models.FindReplyByTid(&models.Topic{Id: tid})
		self.Data["upper_topic"] = models.FindTopicById(tid - 1)
		self.Data["lower_topic"] = models.FindTopicById(tid + 1)

		self.TplName = "topic/topicdetial.html"

	}

}

func (self *TopicController) CreatePage() {
	uid := self.GetSession("uid")
	if uid == nil {
		self.Ctx.Redirect(302, "/user/login")
	} else {
		self.TplName = "topic/createpage.html"
	}
}

func (self *TopicController) CreateTopic() {
	title, content, vercode, captcha_id := self.Input().Get("title"), self.Input().Get("content"), self.Input().Get("vercode"), self.Input().Get("captcha_id")
	if !CheckCode(vercode, captcha_id) {
		msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		uid := self.GetSession("uid")
		fmt.Println(uid)
		topic := models.Topic{Title: title, Content: content, Author: &models.User{Id: uid.(int)}}
		tid := models.SaveTopic(&topic)
		msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

func (self *TopicController) ReplyTopic() {
	topic_id, content, vercode, captcha_id := self.Input().Get("topic_id"), self.Input().Get("content"), self.Input().Get("vercode"), self.Input().Get("captcha_id")
	uid := self.GetSession("uid")
	if uid == nil {
		self.Ctx.Redirect(302, "/user/login")
	} else {
		if !CheckCode(vercode, captcha_id) {
			msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			tid, _ := strconv.Atoi(topic_id)
			reply := models.Reply{Topic: &models.Topic{Id: tid}, Content: content, User: &models.User{Id: uid.(int)}}
			models.SaveReply(&reply)
			topic := models.FindTopicById(tid)
			models.IncrReplyCount(&topic)
			msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
			self.Data["json"] = &msg
			self.ServeJSON()
		}
	}

}
