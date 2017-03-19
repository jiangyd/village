package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"village/models"
)

type UserController struct {
	beego.Controller
}

//登录页面
func (self *UserController) LoginPage() {
	self.TplName = "user/login.html"
}

//登录接口
func (self *UserController) Login() {
	email, password, vercode, captcha_id := self.Input().Get("email"), self.Input().Get("password"), self.Input().Get("vercode"), self.Input().Get("captcha_id")
	// if !CheckCode(vercode, captcha_id) {
	// 	msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
	// 	self.Data["json"] = &msg
	// 	self.ServeJSON()
	// 	return
	// }
	fmt.Println(vercode, captcha_id)
	userinfo := models.CheckLogin(email, password)
	if len(userinfo.Email) > 0 {
		self.SetSession("uid", userinfo.Id)
		self.SetSession("nickname", userinfo.Nickname)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		msg := map[string]interface{}{"code": 1, "msg": "用户名或密码错误!"}
		self.Data["json"] = &msg
		self.ServeJSON()
		return

	}

}

func (self *UserController) Logout() {
	self.DelSession("uid")
	self.Data["islogin"] = false
	self.Ctx.Redirect(302, "/")
}

func (self *UserController) UpdatePwd() {
	oldpassword, newpassword := self.Input().Get("oldpassword"), self.Input().Get("newpassword")
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {

		user := models.FindUserDetialById(uid.(int))
		if oldpassword == user.Password {
			user.Password = newpassword
			models.UpdateUser(&user)
			self.DelSession("uid")
			msg := map[string]interface{}{"code": 0, "msg": "success"}
			self.Data["json"] = &msg
			self.ServeJSON()

		} else {
			msg := map[string]interface{}{"code": 1, "msg": "原密码错误"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}
	}

}

func (self *UserController) Forget() {
	self.TplName = "user/forget.html"
}

func (self *UserController) UserDetial() {
	fmt.Println("aaaaaaaaaaaaaaa")
	uid := self.Ctx.Input.Param(":uid")
	userid, _ := strconv.Atoi(uid)
	//判断用户是否存在
	if models.IsUserExit(&models.User{Id: userid}) {
		user := models.FindUserDetialById(userid)
		self.Data["detial_userinfo"] = user
		self.Data["MyTopic"] = models.FindTopicByUid(&models.User{Id: userid})
		self.Data["MyReply"] = models.FindReplyByUid(&models.User{Id: userid})
		sessionuid := self.GetSession("uid")
		//判断是否登陆状态
		if sessionuid == nil {
			self.Data["isfollow"] = false
			self.Data["islogin"] = false

		} else {
			self.Data["islogin"] = true
			self.Data["userinfo"] = models.FindUserDetialById(sessionuid.(int))
			//判断是否访问自己的详情页
			if sessionuid.(int) == userid {

				self.Data["isself"] = true
				//判断是否已关注用户
				if models.IsFirend(&models.User{Id: sessionuid.(int)}, &models.User{Id: userid}) {
					self.Data["isfollow"] = true
				} else {
					self.Data["isfollow"] = false
				}
			} else {
				self.Data["isself"] = false
			}

		}
		self.TplName = "user/detial.html"
	} else {
		self.Ctx.Redirect(302, "/")
	}
}

func (self *UserController) Set() {

	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {

		user := models.FindUserDetialById(uid.(int))
		self.Data["islogin"] = true
		self.Data["userinfo"] = user
		self.Data["IsSeting"] = true
		self.TplName = "user/setinfo.html"
	}

}

func (self *UserController) SetInfo() {

	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		fmt.Println("mmm")
		nickname, sex, city, sign := self.Input().Get("nickname"), self.Input().Get("sex"), self.Input().Get("city"), self.Input().Get("sign")
		sexv, _ := strconv.Atoi(sex)
		user := models.FindUserDetialById(uid.(int))
		user.Nickname = nickname
		user.Sex = sexv
		user.City = city
		user.Sign = sign
		models.UpdateUser(&user)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

func (self *UserController) Message() {
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		user := models.FindUserDetialById(uid.(int))
		self.Data["islogin"] = true
		self.Data["userinfo"] = user
		self.Data["IsMessage"] = true
		self.TplName = "user/message.html"
	}

}

func (self *UserController) UserTopic() {
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		user := models.FindUserDetialById(uid.(int))
		self.Data["islogin"] = true
		self.Data["userinfo"] = user
		self.Data["IsMyTopic"] = true
		self.Data["MyTopic"] = models.FindTopicByUid(&models.User{Id: uid.(int)})
		self.Data["MyReply"] = models.FindReplyByUid(&models.User{Id: uid.(int)})
		self.TplName = "user/usertopic.html"

	}

}

func (self *UserController) Collection() {

	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		user := models.FindUserDetialById(uid.(int))
		self.Data["islogin"] = true
		self.Data["userinfo"] = user
		self.Data["IsCollection"] = true

		ids := models.FindCollecByUid("tid", &models.User{Id: uid.(int)})
		//判断是否有收藏数据
		if len(ids) > 0 {
			var tids []int
			for _, item := range ids {
				tids = append(tids, item.TypeId)
			}
			self.Data["MyCollecTopic"] = models.FindTopicByIds(tids)
		} else {
			self.Data["MyCollecTopic"] = []int{}
		}
		self.TplName = "user/collection.html"
	}
}

func (self *UserController) FollowPage() {
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		user := models.FindUserDetialById(uid.(int))
		self.Data["islogin"] = true
		self.Data["userinfo"] = user
		self.Data["IsFollow"] = true
		self.TplName = "user/follow.html"
	}

}

func (self *UserController) RegisterPage() {

	self.TplName = "user/register.html"
}

func (self *UserController) Register() {
	flash := beego.NewFlash()
	email, nickname, password, vercode, captcha_id := self.Input().Get("email"), self.Input().Get("nickname"), self.Input().Get("password"), self.Input().Get("vercode"), self.Input().Get("captcha_id")
	if !CheckCode(vercode, captcha_id) {
		msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

	if len(password) == 0 || len(nickname) == 0 {
		flash.Error("用户名或密码不能为空")
		flash.Store(&self.Controller)
		self.Redirect("/user/register", 302)
	} else if len(email) == 0 {
		flash.Error("邮箱不能为空")
		flash.Store(&self.Controller)
		self.Redirect("/user/register", 302)
	} else if flag, _ := models.FindUserByEmail(email); flag {
		flash.Error("该邮箱已注册过，请使用其它邮箱")
		flash.Store(&self.Controller)
		self.Redirect("/user/register", 302)
	} else {
		user := models.User{Nickname: nickname, Email: email, Password: password}
		uid := models.AddUser(&user)
		self.SetSession("uid", int(uid))
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
}
