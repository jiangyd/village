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
	sessionid := self.GetSession("uid")
	if tid >= 0 {
		fmt.Println(self.GetSession("SessionId"))
		topic := models.FindTopicById(tid)
		author := topic.Author.Id
		//作者更多的文章
		self.Data["other_topic"] = models.FindTopicByUid(topic.Author)
		models.IncrView(&topic)
		self.Data["topic"] = topic
		self.Data["replyinfo"] = models.FindReplyByTid(&models.Topic{Id: tid})
		self.Data["upper_topic"] = models.FindTopicById(tid - 1)
		self.Data["lower_topic"] = models.FindTopicById(tid + 1)

		self.TplName = "topic/topicdetial.html"
		if sessionid == nil {
			self.Data["islogin"] = false
			self.Data["isdz"] = false
			self.Data["isself"] = false
			self.Data["isfirend"] = false
			self.Data["isfollow"] = false
		} else {
			//判断是否访问自己的详情页
			if sessionid.(int) == author {
				self.Data["isself"] = true
			} else {
				self.Data["isself"] = false
				//判断是否已关注用户
				if models.IsFirend(&models.User{Id: sessionid.(int)}, &models.User{Id: author}) {
					self.Data["isfirend"] = true
				} else {
					self.Data["isfirend"] = false
				}
			}
			self.Data["islogin"] = true
			self.Data["userinfo"] = models.FindUserDetialById(sessionid.(int))
			self.Data["collection"] = models.FindCollec("tid", tid, &models.User{Id: sessionid.(int)})
			self.Data["isdz"] = models.IsDz("tid", tid, &models.User{Id: sessionid.(int)})
		}

	}

}

func (self *TopicController) CreatePage() {
	uid := self.GetSession("uid")
	if uid == nil {
		self.Ctx.Redirect(302, "/user/login")
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(uid.(int))
		self.Data["categorys"] = models.GetTopicCategory()
		self.TplName = "topic/createpage.html"
	}
}

func (self *TopicController) CreateTopic() {
	title, content, vercode, captcha_id, category := self.Input().Get("title"), self.Input().Get("content"), self.Input().Get("vercode"), self.Input().Get("captcha_id"), self.Input().Get("category")
	category_id, _ := strconv.Atoi(category)
	if !CheckCode(vercode, captcha_id) {
		msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		uid := self.GetSession("uid")
		fmt.Println(uid)
		topic := models.Topic{Title: title, Content: content, Author: &models.User{Id: uid.(int)}, Category: &models.Categorys{Id: category_id}}
		tid := models.SaveTopic(&topic)
		msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

func (self *TopicController) EditPage() {
	uid := self.GetSession("uid")
	tid, _ := strconv.Atoi(self.Ctx.Input.Param(":id"))
	if uid == nil {
		self.Ctx.Redirect(302, "/user/login")
	} else {
		self.Data["islogin"] = true
		self.Data["userinfo"] = models.FindUserDetialById(uid.(int))
		self.Data["topic"] = models.FindTopicById(tid)
		self.Data["categorys"] = models.GetTopicCategory()
		self.TplName = "topic/edittopic.html"
	}
}

func (self *TopicController) EditTopic() {
	topic_id, title, content, vercode, captcha_id, category := self.Input().Get("topic_id"), self.Input().Get("title"), self.Input().Get("content"), self.Input().Get("vercode"), self.Input().Get("captcha_id"), self.Input().Get("category")
	category_id, _ := strconv.Atoi(category)
	tid, _ := strconv.Atoi(topic_id)
	uid := self.GetSession("uid")
	if uid == nil {
		self.Ctx.Redirect(302, "/user/login")
	} else {
		if !CheckCode(vercode, captcha_id) {
			msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			self.Data["islogin"] = true
			self.Data["userinfo"] = models.FindUserDetialById(uid.(int))
			topic := models.FindTopicById(tid)
			topic.Content = content
			topic.Title = title
			topic.Category = &models.Categorys{Id: category_id}
			models.UpdateTopic(&topic)
			msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
			self.Data["json"] = &msg
			self.ServeJSON()
		}

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

//设为满意答案

func (self *TopicController) Adopt() {
	uid := self.GetSession("uid")
	if uid == nil {
		self.Ctx.Redirect(302, "/")
	} else {
		t_id, r_id := self.Input().Get("tid"), self.Input().Get("rid")
		tid, _ := strconv.Atoi(t_id)
		rid, _ := strconv.Atoi(r_id)
		topic := models.FindTopicById(tid)
		//判断是否是作者
		if topic.Author.Id == uid.(int) {
			//更新满意答案了，同样会更新该topic记录的更新时间
			topic.Adopt = &models.Reply{Id: rid}
			models.UpdateTopic(&topic)
			msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			msg := map[string]interface{}{"code": 1, "msg": "无操作权限", "tid": tid}
			self.Data["json"] = &msg
			self.ServeJSON()
		}

	}

}
