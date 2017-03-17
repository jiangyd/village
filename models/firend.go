package models

import (
	"github.com/astaxie/beego/orm"
)

type Firend struct {
	Id    int
	UserA *User `orm:"rel(fk)"` //关注人
	UserB *User `orm:"rel(fk)"` //被关注的人
}

//添加关注
func FirendAdd(firend *Firend) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(firend)
	return id
}

//是否关注关系
func IsFirend(usera, userb *User) bool {
	o := orm.NewOrm()
	var firend Firend
	return o.QueryTable(firend).Filter("UserA", usera).Filter("UserB", userb).Exist()

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
