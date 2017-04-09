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

//添加分类
func AddCategory(category *Categorys) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(category)
	return id
}

//更新分类
func UpdateCategory(category *Categorys) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(category)
	return id
}

//删除分类
func DelCategory(category *Categorys) int64 {
	o := orm.NewOrm()
	id, _ := o.Delete(category)
	return id
}

//查找分类
func FindCategory(id int) Categorys {
	o := orm.NewOrm()
	var cate Categorys
	o.QueryTable(cate).Filter("Id", id).One(&cate)
	return cate
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
