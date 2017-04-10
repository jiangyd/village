package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Reply struct {
	Id      int
	Topic   *Topic    `orm:"rel(fk)"`
	Content string    `orm:"type(text)"`
	User    *User     `orm:"rel(fk)"`
	Up      int       `orm:"default(0)"`
	Ctime   time.Time `orm:"auto_now_add;type(datetime)"`
	Disable bool      `orm:"default(false)"` //是否屏蔽
}

func SaveReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(reply)
	return id
}

//查找活跃用户
func FindHotUser() []orm.Params {
	o := orm.NewOrm()
	var replys []orm.Params
	o.Raw("select count(r.user_id) as num,u.nickname,r.user_id ,u.avatar from reply r,user u where r.user_id=u.id  group by r.user_id  order by num  DESC limit 12").Values(&replys, "num", "nickname", "user_id", "avatar")
	return replys
}

//通过用户id查找评论
func FindReplyByUid(uid *User) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replys []*Reply
	o.QueryTable(reply).Filter("User", uid).OrderBy("-Ctime").RelatedSel().All(&replys)
	return replys

}

//通过主题id,查找评论
func FindReplyByTid(tid *Topic) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replys []*Reply
	o.QueryTable(reply).Filter("Topic", tid).OrderBy("-Up", "-Ctime").RelatedSel().All(&replys)
	return replys
}

//更新评论
func UpdateReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(reply)
	return id
}

//删除评论
func DeleteReply(reply *Reply) {
	o := orm.NewOrm()
	o.Delete(reply)
}

//更新点赞数
func UpReply(reply *Reply) {
	o := orm.NewOrm()
	reply.Up = reply.Up + 1
	o.Update(reply, "Up")
}

func FindReplyByRid(id int) Reply {
	o := orm.NewOrm()
	var reply Reply
	o.QueryTable(reply).Filter("Id", id).RelatedSel().One(&reply)
	return reply
}

//获取所有评论
func GetAllReply() []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replys []*Reply
	o.QueryTable(reply).OrderBy("-Ctime").RelatedSel().All(&replys)
	return replys
}
