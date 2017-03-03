package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
	"tester/models"
	_ "tester/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/tester?charset=utf8")
	orm.RegisterModel(
		new(models.User),
		new(models.Topic),
		new(models.Categorys),
		new(models.Reply),
	)
	orm.RunSyncdb("default", false, true)
}
func main() {
	/*开启session设置*/
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "SessionId"
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "root:123456@tcp(127.0.0.1:3306)/tester?charset=utf8"
	/*end*/
	beego.Run()

}
