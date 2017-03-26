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

//是否存在点赞

func IsDz(t string, typeid int, uid *User) bool {
	o := orm.NewOrm()
	var dz Dz
	return o.QueryTable(dz).Filter("Type", t).Filter("TypeId", typeid).Filter("Uid", uid).Exist()
}

func AddDz(dz *Dz) int64 {

	o := orm.NewOrm()
	id, err := o.Insert(dz)
	if err != nil {
		return -1
	}
	return id
}

//取消点赞
func DelDz(dz *Dz) int64 {
	o := orm.NewOrm()
	id, _ := o.Delete(dz)
	return id
}
