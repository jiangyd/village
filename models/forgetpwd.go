package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ForGetPwd struct {
	Id    int
	Name  *User `orm:"rel(fk);unique"`
	Uuid  string
	Etime time.Time
}

func CheckForGet(uuid string, t time.Time) bool {
	o := orm.NewOrm()
	var forgetpwd ForGetPwd
	return o.QueryTable(forgetpwd).Filter("UUID", uuid).Filter("Etime__gte", t).Exist()
}

func FindForGetPwdByUuid(uuid string) ForGetPwd {
	o := orm.NewOrm()
	var forget ForGetPwd
	o.QueryTable(forget).Filter("UUID", uuid).RelatedSel().One(&forget)
	return forget
}

//通过用户查找

func FindForGetPwdByuser(uid int) ForGetPwd {
	o := orm.NewOrm()
	var forgetpwd ForGetPwd
	o.QueryTable(forgetpwd).Filter("Name", uid).RelatedSel().One(&forgetpwd)
	return forgetpwd
}

//是否存在
func IsExitForGetPwdByuser(uid int) bool {
	o := orm.NewOrm()
	var forgetpwd ForGetPwd
	return o.QueryTable(forgetpwd).Filter("Name", uid).Exist()

}

//添加找回密码
func AddForGetPwd(forgetpwd *ForGetPwd) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(forgetpwd)
	return id
}

//更新uuid，时间

func UpdateForGetPwd(forgetpwd *ForGetPwd) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(forgetpwd)
	return id
}
