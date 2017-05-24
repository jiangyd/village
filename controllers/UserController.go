package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"village/models"
)

type UserController struct {
	beego.Controller
}

//登录页面
func (self *UserController) LoginPage() {
	self.TplName = "user/login.html"
}

func (self *UserController) ForGetPwdPage() {
	//用户点击重置密码链接,需要把uuid回传
	uuid := self.Input().Get("uuid")
	self.Data["uuid"] = uuid
	self.TplName = "user/forgetpwd.html"

}

func (self *UserController) SetNewPwd() {
	uuid, password := self.Input().Get("uuid"), self.Input().Get("password")
	now := time.Now()
	//检测uuid是否有效,有效便更新密码,否则直接返回
	if models.CheckForGet(uuid, now) {
		//通过uuid查找对应要修改密码的用户
		u := models.FindForGetPwdByUuid(uuid)
		user := models.FindUserDetialById(u.Name.Id)
		user.Password = password
		models.UpdateUser(&user)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
	msg := map[string]interface{}{"code": 1, "msg": "invalid token"}
	self.Data["json"] = &msg
	self.ServeJSON()

}

func (self *UserController) ForGetPwd() {
	email, vercode, captcha_id := self.Input().Get("email"), self.Input().Get("vercode"), self.Input().Get("captcha_id")
	if !CheckCode(vercode, captcha_id) {
		msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
		self.Data["json"] = &msg
		self.ServeJSON()
		return
	}
	//通过邮箱判断用户是否存在
	if models.IsUserExitByEmail(email) {
		_, user := models.FindUserByEmail(email)
		uuid := Encrypt(email + Getuuid())
		//当前时间
		now := time.Now()
		//设置过期时间,这里设置1小时后过期
		h, _ := time.ParseDuration("1h")
		//添加时间
		m := now.Add(h)

		//是否第一次找回密码,是则更新表记录的uuid,过期时间,否则添加
		if models.IsExitForGetPwdByuser(user.Id) {
			forgetpwd := models.FindForGetPwdByuser(user.Id)
			forgetpwd.Uuid = uuid
			forgetpwd.Etime = m
			models.UpdateForGetPwd(&forgetpwd)
		} else {
			forgetpwd := models.ForGetPwd{Uuid: uuid, Name: &models.User{Id: user.Id}, Etime: m}
			models.AddForGetPwd(&forgetpwd)
		}
		//发送找回密码邮件
		url := "http://192.168.1.12:8080/forgetpwd/?uuid=" + uuid
		SendMail(email, "<h2>请点击以下链接重置密码,如非本人操作请忽略:</h2><p><a href="+url+">"+url+"</a>", "重置密码")
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		msg := map[string]interface{}{"code": 1, "msg": "邮箱不存在"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
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
		self.SetSession("email", userinfo.Email)
		if userinfo.Mfa {
			msg := map[string]interface{}{"code": 0, "mfa": true}
			self.Data["json"] = &msg
			self.ServeJSON()
			return
		}
		self.SetSession("uid", userinfo.Id)
		self.SetSession("nickname", userinfo.Nickname)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
		return
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

func (self *UserController) MfaVerifyPage() {
	self.TplName = "user/mfapwd.html"
}
func (self *UserController) MfaVerify() {
	email := self.GetSession("email")
	code := self.Input().Get("code")
	if email == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		_, user := models.FindUserByEmail(email.(string))
		if code == Totp(user.Secret, 0) {
			self.SetSession("uid", user.Id)
			msg := map[string]interface{}{"code": 0, "msg": "success"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			msg := map[string]interface{}{"code": 1, "msg": "invalid code"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}
	}

}

func (self *UserController) MFAPage() {
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		user := models.FindUserDetialById(uid.(int))
		secret := GetSecret()
		// secret := "vbj6je5hx7nttlh6"
		qrdata := Getotpauth(user.Nickname, secret, "测试村")
		//把邮箱md5加密成字符串，当作二维码文件名,这样文件名应该是每个用户只有一个
		//不会因为用户多次刷新而生成不必要的文件,以防造成空间浪费
		md5ctx := md5.New()
		md5ctx.Write([]byte(user.Email))
		filename := fmt.Sprintf("%x", md5ctx.Sum(nil))
		fmt.Println(filename, "pppp")
		self.Data["qrimg"] = GetQrCode(qrdata, filename)
		if user.Mfa != true {
			user.Secret = secret
			models.UpdateUser(&user)
			self.Data["secret"] = secret
		} else {
			self.Data["secret"] = "****************"
		}

		self.Data["islogin"] = true
		self.Data["userinfo"] = user
		self.Data["IsMFA"] = true
		self.TplName = "user/mfa.html"
	}
}

func (self *UserController) SetMfa() {
	code1, code2 := self.Input().Get("code1"), self.Input().Get("code2")
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		user := models.FindUserDetialById(uid.(int))
		fmt.Println(code1, code2)
		fmt.Println(Totp(user.Secret, 30), Totp(user.Secret, 0))
		if code1 == Totp(user.Secret, 30) && code2 == Totp(user.Secret, 0) {
			user.Mfa = true
			models.UpdateUser(&user)
			msg := map[string]interface{}{"code": 0, "msg": "success"}
			self.Data["json"] = &msg
			self.ServeJSON()

		} else {
			msg := map[string]interface{}{"code": 1, "msg": "无效密码"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}
	}
}

func (self *UserController) CloseMfa() {
	code1, code2 := self.Input().Get("code1"), self.Input().Get("code2")
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {
		user := models.FindUserDetialById(uid.(int))
		fmt.Println(code1, code2)
		fmt.Println(Totp(user.Secret, 30), Totp(user.Secret, 0))
		if code1 == Totp(user.Secret, 30) && code2 == Totp(user.Secret, 0) {
			user.Mfa = false
			models.UpdateUser(&user)
			msg := map[string]interface{}{"code": 0, "msg": "success"}
			self.Data["json"] = &msg
			self.ServeJSON()

		} else {
			msg := map[string]interface{}{"code": 1, "msg": "无效密码"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}
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

func (self *UserController) SendMsg() {
	sessionid := self.GetSession("uid")
	if sessionid == nil {
		msg := map[string]interface{}{"code": 1, "msg": "need loging"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		userb, msgcontent := self.Input().Get("userb"), self.Input().Get("content")
		userbid, _ := strconv.Atoi(userb)
		message := models.Message{Send: &models.User{Id: sessionid.(int)}, Recv: &models.User{Id: userbid}, Content: msgcontent}
		models.SendMsg(&message)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}
