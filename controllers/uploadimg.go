package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UploadImg struct {
	beego.Controller
}

func (self *UploadImg) Upload() {
	//获取头像文件
	f, h, _ := self.GetFile("file")
	path := "./static/images/" + h.Filename
	defer f.Close()

	err := self.SaveToFile("file", path)
	fmt.Println(err)
	msg := map[string]interface{}{"code": 0, "msg": "success"}
	self.Data["json"] = &msg
	self.ServeJSON()

}
