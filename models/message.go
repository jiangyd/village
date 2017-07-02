package models

import (
	"fmt"
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
func GetMyMsg(recv *User, msgtype string) []orm.Params {
	o := orm.NewOrm()
	var msg []orm.Params
	o.Raw("select a.id,a.send_id,a.content,a.ctime,a.read,a.msgtype,u.nickname from message a , user u  where a.recv_id=u.id and a.recv_id=? and a.msgtype=?", recv, msgtype).Values(&msg)
	return msg
}
