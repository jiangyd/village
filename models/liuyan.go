package models

import (
	"github.com/astaxie/beego/orm"
)

type Liuyan struct {
	Id      int
	Email   string
	Content string
}

func AddLiuYan(liuyan *Liuyan) {
	o := orm.NewOrm()
	o.Insert(liuyan)
}

func GetLiuYan() []*Liuyan {
	o := orm.NewOrm()
	var liuyan Liuyan
	var liuyans []*Liuyan
	o.QueryTable(liuyan).All(&liuyans)
	return liuyans
}
