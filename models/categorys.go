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

//获取所有分类
func GetAllCategory() []*Categorys {
	o := orm.NewOrm()
	var category Categorys
	var categorys []*Categorys
	o.QueryTable(category).All(&categorys)
	return categorys
}

//获取主题分类
func GetTopicCategory() []*Categorys {
	o := orm.NewOrm()
	var category Categorys
	var categorys []*Categorys
	o.QueryTable(category).Filter("CategoryType", "topic").All(&categorys)
	return categorys
}

//获取站点分类
func GetCiteCategory() []*Categorys {
	o := orm.NewOrm()
	var category Categorys
	var categorys []*Categorys
	o.QueryTable(category).Filter("CategoryType", "cite").All(&categorys)
	return categorys
}
