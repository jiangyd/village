package admin

import (
	"github.com/astaxie/beego/orm"
)

type Document struct {
	Id        int
	Title     string
	Pid       int    `orm:"null"` //父ID
	ListOrder int    //排序设置
	Content   string `orm:"type(text)"`
}

//添加节点
func AddDoc(doc *Document) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(doc)
	return id
}

//获取所有节点
func GetDoc() []*Document {
	o := orm.NewOrm()
	var doc Document
	var docs []*Document
	o.QueryTable(doc).All(&docs)
	return docs
}

//查找节点信息
func GetDocById(id int) Document {
	o := orm.NewOrm()
	var doc Document
	o.QueryTable(doc).Filter("Id", id).One(&doc)
	return doc
}

//获取根目录
func GetRootDoc() []*Document {
	o := orm.NewOrm()
	var doc Document
	var docs []*Document
	o.QueryTable(doc).Filter("Pid", 0).All(&docs)
	return docs
}

//更新节点
func UpdateDoc(doc *Document) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(doc)
	return id
}

//查看当前节点下是否存在子节点
func IsExitSubDoc(pid int) bool {
	o := orm.NewOrm()
	var doc Document
	return o.QueryTable(doc).Filter("Pid", pid).Exist()
}

//删除节点
func DelDoc(id int) {
	o := orm.NewOrm()
	var doc Document
	o.QueryTable(doc).Filter("Id", id).Delete()
}
