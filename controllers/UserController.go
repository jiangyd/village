package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "strconv"
	"tester/models"
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
	fmt.Println(email, password)
	fmt.Println(vercode, captcha_id)
	userinfo := models.CheckLogin(email, password)
	if len(userinfo.Email) > 0 {
		self.SetSession("uid", userinfo.Id)
		self.SetSession("nickname", userinfo.Nickname)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
		self.Ctx.Redirect(302, "/")

	} else {
		msg := map[string]interface{}{"code": 1, "msg": "用户名或密码错误!"}
		self.Data["json"] = &msg
		self.ServeJSON()
		return

	}

}

func (self *UserController) Forget() {
	self.TplName = "user/forget.html"
}

func (self *UserController) Set() {
	self.TplName = "user/set.html"
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
		models.AddUser(&user)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
}

//用户详情
func (self *UserController) Detial() {
	// id := self.Ctx.Input.Param(":uid")
	// uid, _ := strconv.Atoi(id)
	// models.FindUserDetialById(uid)
	self.TplName = "user/detial.html"
}
