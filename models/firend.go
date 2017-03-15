package models

import (
	"github.com/astaxie/beego/orm"
)

type Firend struct {
	Id    int
	UserA *User `orm:"rel(fk)"` //关注人
	UserB *User `orm:"rel(fk)"` //被关注的人
}

//通过关注人查找
func FindFirendByUserA(uid *User) []*Firend {
	o := orm.NewOrm()
	var firend Firend
	var firends []*Firend
	o.QueryTable(firend).Filter("UserA", uid).RelatedSel().All(&firends)
	return firends
}

//通过被关注人查找
func FindFirendByUserB(uid *User) []*Firend {
	o := orm.NewOrm()
	var firend Firend
	var firends []*Firend
	o.QueryTable(firend).Filter("UserB", uid).RelatedSel().All(&firends)
	return firends
}