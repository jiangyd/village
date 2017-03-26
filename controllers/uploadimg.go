package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"village/models"
)

type UploadImg struct {
	beego.Controller
}

func (self *UploadImg) Upload() {
	//获取头像文件
	sessionid := self.GetSession("uid")
	if sessionid == nil {
		self.Data["islogin"] = false
		self.Ctx.Redirect(302, "/")
	} else {

		user := models.FindUserDetialById(sessionid.(int))
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

//发帖上传的附件图片
func (self *UploadImg) TopicUpload() {
	sessionid := self.GetSession("uid")
	if sessionid == nil {
		msg := map[string]interface{}{"code": 1, "msg": "need loging"}
		self.Data["json"] = &msg
		self.ServeJSON()
	} else {
		timestamp := time.Now().Unix()
		f, h, _ := self.GetFile("file")
		timestr := strconv.Itoa(int(timestamp))
		path := "/static/images/" + timestr + h.Filename
		defer f.Close()
		self.SaveToFile("file", "."+path)
		//传入文件路径，及七牛保存的文件名
		UpQiNiu("."+path, "/static/images/"+timestr+"_"+h.Filename)
		msg := map[string]interface{}{"code": 0, "msg": "", "data": map[string]interface{}{"src": path, "title": ""}}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
}
