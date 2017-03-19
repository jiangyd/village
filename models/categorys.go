package models

import (
	"github.com/astaxie/beego/orm"
)

//分类表
type Categorys struct {
	Id       int
	Category string `orm:"unique"`
}

//文章对应分类表
// type Topic_Category struct {
// 	Id         int
// 	TopicId    *Topic     `orm:"rel(fk)"`
// 	CategoryId *Categorys `orm:"rel(fk)"`
// }

// //添加文章关联分类
// func AddTopic_Category(t *Topic_Category) int64 {
// 	o := orm.NewOrm()
// 	id, _ := o.Insert(t)
// 	return id
// }

func GetAllCategory() []*Categorys {
	o := orm.NewOrm()
	var category Categorys
	var categorys []*Categorys
	o.QueryTable(category).All(&categorys)
	return categorys

}
