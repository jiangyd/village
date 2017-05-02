package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"village/models"
	"village/models/admin"
	_ "village/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/village?charset=utf8&loc=Asia%2FShanghai")
	orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(
		new(models.User),
		new(models.Topic),
		new(models.Categorys),
		new(models.Reply),
		new(models.Message),
		new(models.Dz),
		new(models.Collection),
		new(models.Firend),
		new(models.Sites),
		new(models.QiNiuFile),
		new(admin.Menu),
		new(admin.SubMenu),
		new(admin.Document),
	)
	orm.RunSyncdb("default", false, true)
}
func main() {
	/*开启session设置*/
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "SessionId"
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "root:123456@tcp(127.0.0.1:3306)/village?charset=utf8&loc=Asia%2FShanghai"
	/*end*/
	beego.Run()

}
