package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"village/models"
)

type UploadImg struct {
	beego.Controller
}

func (self *UploadImg) Upload() {
	//获取头像文件
	uid := self.GetSession("uid")
	if uid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {

		user := models.FindUserDetialById(uid.(int))
		f, h, _ := self.GetFile("file")
		path := "/static/images/" + h.Filename
		defer f.Close()

		err := self.SaveToFile("file", "."+path)
		fmt.Println(err)
		user.Avatar = path
		models.UpdateUser(&user)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()

	}

}
