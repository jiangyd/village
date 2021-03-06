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

func (self *TopicController) TopicSearch() {
	category := self.Input().Get("category")
	categoryob := models.GetCategoryByName(category)
	fmt.Println(category)
	self.Data["cuscategory"] = category
	self.Data["topic"] = models.FindTopicByCategory(&models.Categorys{Id: categoryob.Id})
	self.Data["Categorys"] = models.GetTopicCategory()
	self.TplName = "categorytopic.html"
}

func (self *TopicController) DaShang() {
	//给用户打赏
	userb := self.Input().Get("userb")
	money := self.Input().Get("money")
	moneyint, _ := strconv.Atoi(money)
	userbint, _ := strconv.Atoi(userb)
	sessionid := self.GetSession("uid")
	if sessionid == nil {
		msg := map[string]interface{}{"code": 2, "msg": "need login"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		if sessionid.(int) == userbint {
			msg := map[string]interface{}{"code": 3, "msg": "不能给自己打赏"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			//被打赏的人金钱要增加
			userbinfo := models.FindUserDetialById(userbint)
			userbinfo.Money += moneyint
			models.UpdateUser(&userbinfo)
			//打赏的人金钱则要减少
			userainfo := models.FindUserDetialById(sessionid.(int))
			userainfo.Money -= moneyint
			models.UpdateUser(&userainfo)
			msg := map[string]interface{}{"code": 0, "msg": "打赏成功"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}

	}
}

func (self *TopicController) TopicDetial() {
	id := self.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	sessionid := self.GetSession("uid")
	if models.IsTopicExit(tid) {
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
			self.Data["collection"] = models.IsCollecExit("tid", tid, &models.User{Id: sessionid.(int)})
			self.Data["isdz"] = models.IsDz("tid", tid, &models.User{Id: sessionid.(int)})
		}

	} else {
		self.Redirect("/", 404)
	}

}

func (self *TopicController) TopicDT() {
	id := self.Input().Get("id")
	topdata := models.GetTopicDT(id)
	self.Data["topic"] = topdata[len(topdata)-1]
	self.Data["islogin"] = true
	self.TplName = "topic/ystopicdetial.html"
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
	fmt.Println(vercode, captcha_id)
	// if !CheckCode(vercode, captcha_id) {
	// 	msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
	// 	self.Data["json"] = &msg
	// 	self.ServeJSON()
	// } else {
	uid := self.GetSession("uid")
	fmt.Println(uid)
	topic := models.Topic{Title: title, Content: content, Author: &models.User{Id: uid.(int)}, Category: &models.Categorys{Id: category_id}, Disable: false}
	tid := models.SaveTopic(&topic)
	msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
	self.Data["json"] = &msg
	self.ServeJSON()
	// }

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
		fmt.Println(vercode, captcha_id)
		// if !CheckCode(vercode, captcha_id) {
		// 	msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
		// 	self.Data["json"] = &msg
		// 	self.ServeJSON()
		// } else {
		if models.IsTopicExit(tid) {
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
		} else {
			msg := map[string]interface{}{"code": 2, "msg": "帖子不存在"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}

		// }

	}
}

func (self *TopicController) ReplyTopic() {
	topic_id, content, vercode, captcha_id := self.Input().Get("topic_id"), self.Input().Get("content"), self.Input().Get("vercode"), self.Input().Get("captcha_id")
	uid := self.GetSession("uid")
	if uid == nil {
		msg := map[string]interface{}{"code": 2, "msg": "need login"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		fmt.Println(vercode, captcha_id)
		// if !CheckCode(vercode, captcha_id) {
		// 	msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
		// 	self.Data["json"] = &msg
		// 	self.ServeJSON()
		// } else {
		tid, _ := strconv.Atoi(topic_id)
		reply := models.Reply{Topic: &models.Topic{Id: tid}, Content: content, User: &models.User{Id: uid.(int)}}
		models.SaveReply(&reply)
		topic := models.FindTopicById(tid)
		models.IncrReplyCount(&topic)
		msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
		self.Data["json"] = &msg
		self.ServeJSON()
		// }
	}

}

//设为满意答案

func (self *TopicController) Adopt() {
	uid := self.GetSession("uid")
	if uid == nil {
		msg := map[string]interface{}{"code": 2, "msg": "need login"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		t_id, r_id := self.Input().Get("tid"), self.Input().Get("rid")
		fmt.Println("tid:", t_id, "###rid:", r_id)
		tid, err := strconv.Atoi(t_id)
		if err != nil {
			fmt.Println("tid转换:", err)
		}
		rid, err := strconv.Atoi(r_id)
		if err != nil {
			fmt.Println("tid转换:", err)
		}
		topic := models.FindTopicById(tid)
		//判断是否是作者
		if topic.Author.Id == uid.(int) {
			//更新满意答案了，同样会更新该topic记录的更新时间
			//当前帖子ID,回复id,是否是满意答案
			if models.IsAdoptReply(&models.Topic{Id: tid}, rid) {
				//有则删除
				reply := models.FindReplyByRid(rid)
				reply.Adopt = false
				models.UpdateReply(&reply)
				msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
				self.Data["json"] = &msg
				self.ServeJSON()
			} else {
				//检测当前帖子有满意答案
				if models.IsAdoptReplyByTid(&models.Topic{Id: tid}) {
					//有则不添加
					msg := map[string]interface{}{"code": 1, "msg": "已有满意评论,请先取消!再重新采纳,", "tid": tid}
					self.Data["json"] = &msg
					self.ServeJSON()
				} else {
					//没有则添加
					reply := models.FindReplyByRid(rid)
					reply.Adopt = true
					models.UpdateReply(&reply)
					msg := map[string]interface{}{"code": 0, "msg": "success", "tid": tid}
					self.Data["json"] = &msg
					self.ServeJSON()
				}

			}

		} else {
			msg := map[string]interface{}{"code": 1, "msg": "无操作权限", "tid": tid}
			self.Data["json"] = &msg
			self.ServeJSON()
		}

	}

}
