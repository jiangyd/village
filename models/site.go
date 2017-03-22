package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Sites struct {
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

func GetAllSite() []*Sites {
	o := orm.NewOrm()
	var site Sites
	var sites []*Sites
	o.QueryTable(site).RelatedSel().All(&sites)
	return sites
}
