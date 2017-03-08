package models

import (
	"github.com/astaxie/beego/orm"
)

type Collection struct {
	Id     int
	Type   string //收藏类型  tid:帖子收藏
	TypeId int
	Uid    *User `orm:"rel(fk)"`
}

//添加收藏
func AddCollection(collection *Collection) int64 {

	o := orm.NewOrm()
	id, err := o.Insert(collection)
	if err != nil {
		return -1
	}
	return id
}

//删除收藏
func DelCollection(t string, typeid int, uid *User) {

	o := orm.NewOrm()
	var collection Collection
	// id, err := o.Delete(collection)
	// o.Commit()
	// if err != nil {
	// 	panic(err)
	// }
	// return id
	o.QueryTable(collection).Filter("Type", t).Filter("TypeId", typeid).Filter("Uid", uid).Delete()
}

//查找收藏
func FindCollec(t string, typeid int, uid *User) Collection {
	o := orm.NewOrm()
	var collection Collection
	o.QueryTable(collection).Filter("Type", t).Filter("TypeId", typeid).Filter("Uid", uid).One(&collection)
	return collection
}
