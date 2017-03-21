package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Cites struct {
	Id       int
	Url      string
	Ctime    time.Time `orm:"auto_now_add;type(datetime)"`
	User     *User     `orm:"rel(fk)"`
	Title    string
	Content  string
	Img      string
	Up       int
	Category *Categorys `orm:"rel(fk)"`
}

func GetAllCite() []*Cites {
	o := orm.NewOrm()
	var cite Cites
	var cites []*Cites
	o.QueryTable(cite).RelatedSel().All(&cites)
	return cites
}
