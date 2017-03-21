package controllers

import (
	"github.com/astaxie/beego"
	"village/models"
)

type CiteController struct {
	beego.Controller
}

func (self *CiteController) CitePage() {
	self.Data["cites"] = models.GetAllCite()
	self.TplName = "cite/cite.html"
}
