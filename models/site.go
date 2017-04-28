package models

import (
	"fmt"
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

//获取所有的站点
func GetAllSite() []*Sites {
	o := orm.NewOrm()
	var site Sites
	var sites []*Sites
	o.QueryTable(site).RelatedSel().All(&sites)
	return sites
}

//通过siteid查找
func FindSiteById(id int) Sites {
	o := orm.NewOrm()
	var sites Sites
	o.QueryTable(sites).Filter("Id", id).RelatedSel().One(&sites)
	return sites
}

//添加站点
func AddSite(site *Sites) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(site)
	fmt.Println(err)
	return id
}

//修改站点
func UpdateSite(site *Sites) int64 {
	o := orm.NewOrm()
	id, err := o.Update(site)
	fmt.Println(err)
	return id
}

//删除站点
func DelSite(id int) {
	o := orm.NewOrm()
	var site Sites
	o.QueryTable(site).Filter("Id", id).Delete()

}
