package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id      int
	Send    *User     `orm:"rel(fk)"`
	Recv    *User     `orm:"rel(fk)"`
	Content string    `orm:"size(500)"`
	Ctime   time.Time `orm:"auto_now_add;type(datetime)"`
	Read    int       `orm:"default(0)"` //消息是否已读,0未读,1已读,默认为0
}

func SendMsg(message *Message) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(message)
	return id
}
