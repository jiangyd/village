package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	//验证码长度
	cpt.ChallengeNums = 4
	//验证码宽度,高度
	cpt.StdHeight = 40
	cpt.StdWidth = 120
}

type Capt struct {
	beego.Controller
}

//校验验证码是否有效
func CheckCode(vercode, captcha_id string) bool {
	yzm := cpt.Verify(captcha_id, vercode)
	return yzm

}

func (this *Capt) Get() {
	this.TplName = "public/captcha.html"
}

func (this *Capt) Post() {
	this.Data["Success"] = cpt.VerifyReq(this.Ctx.Request)
}
