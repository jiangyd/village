package models

import (
	"github.com/astaxie/beego/orm"
)

//分类表
type Categorys struct {
	Id           int
	Category     string `orm:"unique"`
	CategoryType string //topic:主题,cite:站点
}

func GetTopicCategory() []*Categorys {
	o := orm.NewOrm()
	var category Categorys
	var categorys []*Categorys
	o.QueryTable(category).Filter("CategoryType", "topic").All(&categorys)
	return categorys

}
func GetCiteCategory() []*Categorys {
	o := orm.NewOrm()
	var category Categorys
	var categorys []*Categorys
	o.QueryTable(category).Filter("CategoryType", "cite").All(&categorys)
	return categorys

}
