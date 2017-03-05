package models

import (
	"github.com/astaxie/beego/orm"
)

type Dz struct {
	Id     int
	Type   string //点赞类型  tid:主题,rid:评论
	TypeId int    //对应类型id
	Uid    *User  `orm:"rel(fk)"`
}

func FindDz(t string, typeid int) []*Dz {
	o := orm.NewOrm()
	var dz Dz
	var dzs []*Dz
	o.QueryTable(dz).Filter("Type", t).Filter("TypeId", typeid).RelatedSel().All(&dzs)
	return dzs
}

func Adddz(dz *Dz) int64 {

	o := orm.NewOrm()
	id, err := o.Insert(dz)
	if err != nil {
		return -1
	}
	return id
}
