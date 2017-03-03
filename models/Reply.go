package models

import (
	"github.com/astaxie/beego/orm"
)

type Reply struct {
	Id      int
	Topic   *Topic `orm:"rel(fk)"`
	Content string `orm:"type(text)"`
	User    *User  `orm:"rel(fk)"`
	Up      int    `orm:"default(0)"`
}

func SaveReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(reply)
	return id
}

func DeleteReply(reply *Reply) {
	o := orm.NewOrm()
	o.Delete(reply)
}

func UpReply(reply *Reply) {
	o := orm.NewOrm()
	reply.Up = reply.Up + 1
	o.Update(reply, "Up")
}
