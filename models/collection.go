package models

import (
	"fmt"
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
func DelCollection(collection *Collection) int64 {

	o := orm.NewOrm()
	id, err := o.Delete(collection)
	o.Commit()
	fmt.Println(id)
	if err != nil {
		panic(err)
	}
	return id
}
