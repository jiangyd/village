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

func AddDoc(doc *Document) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(doc)
	return id
}

func GetDoc() []*Document {
	o := orm.NewOrm()
	var doc Document
	var docs []*Document
	o.QueryTable(doc).All(&docs)
	return docs
}
