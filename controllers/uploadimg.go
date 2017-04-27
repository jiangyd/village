package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"village/models"
)

type UploadImg struct {
	beego.Controller
}

//七牛回调函数
func (self *UploadImg) QiNiuCallBack() {
	var ob models.QiNiuFile
	json.Unmarshal(self.Ctx.Input.RequestBody, &ob)
	models.SaveQiNiuFile(&ob)
	msg := map[string]interface{}{"hash": ob.Hash, "key": ob.Key, "filesize": ob.FileSize}
	self.Data["json"] = &msg
	self.ServeJSON()

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
	fmt.Println(sessionid, "sessionid")
	if sessionid == nil {
		msg := map[string]interface{}{"code": 1, "message": "need loging"}
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
		body := UpQiNiu("."+path, "/static/images/"+timestr+"_"+h.Filename)
		msg := map[string]interface{}{"code": 0, "msg": "", "message": "", "url": "http://file.testwd.cn/" + body.Key, "data": map[string]interface{}{"src": "http://file.testwd.cn/" + body.Key, "title": ""}}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
}
