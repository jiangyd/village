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
	Read    int       `orm:"default(0)"`       //消息是否已读,0未读,1已读,默认为0
	Msgtype string    `orm:"default(private)"` //消息类型，private私信，system系统信息
}

func SendMsg(message *Message) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(message)
	return id
}

//获取自己收到的私信
func GetMyMsg(recv *User, msgtype string) []*Message {
	o := orm.NewOrm()
	var message Message
	var messages []*Message
	o.QueryTable(message).Filter("Recv", recv).Filter("Msgtype", msgtype).RelatedSel("Send").Filter("Id", recv).All(&messages)
	return messages
}
