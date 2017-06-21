package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Topic struct {
	Id            int
	Title         string
	Content       string     `orm:"type(text)"`
	Author        *User      `orm:"rel(fk)"`
	Ctime         time.Time  `orm:"auto_now_add;type(datetime)"`
	Utime         time.Time  `orm:"auto_now;null"`
	View          int        `orm:"default(0)"`
	ReplyCount    int        `orm:"default(0)"`
	LastReplyUser *User      `orm:"rel(fk);null"`
	LastReplyTime time.Time  `orm:"null;type(datetime);"`
	Up            int        `orm:"default(0)"`
	Category      *Categorys `orm:"rel(fk)"`
	Disable       bool       `orm:"default(false)"`
}

type QiNiuFile struct {
	Id       int
	Hash     string
	Key      string
	FileSize int
}

func SaveQiNiuFile(qiniufile *QiNiuFile) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(qiniufile)
	return id
}

func SaveTopic(topic *Topic) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(topic)
	if err != nil {
		panic(err)
	}
	return id
}

//所有主题
func GetAllTopic() []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).OrderBy("-Ctime").RelatedSel().All(&topics)
	return topics
}

//更新主题
func UpdateTopic(topic *Topic) int64 {
	o := orm.NewOrm()
	id, err := o.Update(topic)
	if err != nil {
		panic(err)
	}
	return id
}

//通过用户id查找主题
func FindTopicByUid(uid *User) []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).Filter("Author", uid).OrderBy("-Ctime").RelatedSel().All(&topics)
	return topics
}

//增加点赞数
func UpTopic(topic *Topic) {
	o := orm.NewOrm()
	topic.Up = topic.Up + 1
	o.Update(topic, "Up")
}

//通过Id确认是否存在
func IsTopicExit(id int) bool {
	o := orm.NewOrm()
	var topic Topic
	return o.QueryTable(topic).Filter("Id", id).Exist()
}

//主题详情
func FindTopicById(id int) Topic {
	o := orm.NewOrm()
	var topic Topic
	o.QueryTable(topic).Filter("Id", id).RelatedSel().One(&topic)
	return topic

}

//演示
// func FindTopicById(id interface{}) []orm.Params {
// 	o := orm.NewOrm()
// 	var topic []orm.Params
// 	o.Raw("select id, title, content, author_id, ctime, utime, view, reply_count, last_reply_user_id, last_reply_time, up, category_id, adopt_id, disable from topic  where id="+id.(string)).Values(&topic, "id", "title", "content", "author_id", "ctime", "utime", "view", "reply_count", "last_reply_user_id", "last_reply_time", "up", "category_id", "adopt_id", "disable")
// 	return topic
// }

//通过一组主题id查询
func FindTopicByIds(ids []int) []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).Filter("Id__in", ids).RelatedSel().All(&topics)
	return topics
}

//通过分类查询
func FindTopicByCategory(category *Categorys) []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).Filter("Category", category).RelatedSel().All(&topics)
	return topics

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

//主题评论数增加
func IncrReplyCount(topic *Topic) {
	o := orm.NewOrm()
	topic.ReplyCount = topic.ReplyCount + 1
	o.Update(topic, "ReplyCount")

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
	//点赞数量大于等于1
	o.QueryTable(topic).Filter("Up__gte", 1).OrderBy("-Ctime").RelatedSel().All(&topics)
	return topics
}

//已采纳贴
func AdoptTopicList() []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).Filter("Adopt__isnull", false).OrderBy("-Ctime").RelatedSel().All(&topics)
	return topics
}
