package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	// "log"
)

type Collection struct {
	Id     int
	Type   string //收藏类型  tid:帖子收藏
	TypeId int
	Uid    *User `orm:"rel(fk)"`
}

//添加收藏
func AddCollection(collection *Collection) {
	o := orm.NewOrm()
	_, err := o.Insert(collection)
	if err != nil {
		// log.Fatalln(err.Error())
		fmt.Println("执行失败")
	}
}

//通过用户id查找收藏

func FindCollecByUid(Type string, uid *User) []*Collection {
	o := orm.NewOrm()
	var collection Collection
	var collections []*Collection
	o.QueryTable(collection).Filter("Uid", uid).Filter("Type", Type).RelatedSel().All(&collections, "TypeId")
	return collections

}

//删除收藏
func DelCollection(t string, typeid int, uid *User) {

	o := orm.NewOrm()
	var collection Collection
	o.QueryTable(collection).Filter("Type", t).Filter("TypeId", typeid).Filter("Uid", uid).Delete()
}

//查找收藏
func IsCollecExit(t string, typeid int, uid *User) bool {
	fmt.Println("type:", t)
	fmt.Println("typeid:", typeid)
	o := orm.NewOrm()
	var collection Collection
	return o.QueryTable(collection).Filter("Type", t).Filter("TypeId", typeid).Filter("Uid", uid).Exist()

}
