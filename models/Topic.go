package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Topic struct {
	Id            int
	Title         string
	Content       string    `orm:"type(text)"`
	Author        *User     `orm:"rel(fk)"`
	Ctime         time.Time `orm:"auto_now_add;type(datetime)"`
	Utime         time.Time `orm:"null"`
	View          int       `orm:"default(0)"`
	ReplyCount    int       `orm:"default(0)"`
	LastReplyUser *User     `orm:"rel(fk);null"`
	LastReplyTime time.Time `orm:"null;type(datetime);"`
	Up            int       `orm:"default(0)"`
}

func SaveTopic(topic *Topic) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(topic)
	if err != nil {
		panic(err)
	}
	return id
}

func UpTopic(topic *Topic) {
	o := orm.NewOrm()
	topic.Up = topic.Up + 1
	o.Update(topic, "Up")
}

//主题详情
func FindTopicById(id int) Topic {
	o := orm.NewOrm()
	var topic Topic
	o.QueryTable(topic).Filter("id", id).RelatedSel().One(&topic)
	return topic
}

//最新主题,通过更新时间倒序查询
func NewTopic() []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).OrderBy("-Utime").RelatedSel().All(&topics)
	return topics

}

//主题浏览数增加
func IncrView(topic *Topic) {
	o := orm.NewOrm()
	topic.View = topic.View + 1
	o.Update(topic, "View")

}

//最新回复,通过回复时间倒序查询
func NewReply() []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).OrderBy("-LastReplyTime").RelatedSel().All(&topics)
	return topics
}

//待回复,回复数为0的主题
func WaitReply() []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).Filter("ReplyCount", 0).OrderBy("-Ctime").RelatedSel().All(&topics)
	return topics
}

//点赞帖，最新点赞的主题
func UpTopicList() []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).Filter("Up", 0).OrderBy("-Ctime").RelatedSel().All(&topics)
	return topics
}
