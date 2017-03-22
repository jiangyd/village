package controllers

import (
	"github.com/astaxie/beego"
	"village/models"
)

type SiteController struct {
	beego.Controller
}

func (self *SiteController) SitePage() {
	self.Data["sites"] = models.GetAllSite()
	self.TplName = "site/site.html"
}
