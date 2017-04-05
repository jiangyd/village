package admin

import (
	"github.com/astaxie/beego/orm"
)

//菜单表
type Menu struct {
	Key   string `orm:"pk;"`
	Title string `orm:"unique"`
}

//二级菜单表
type SubMenu struct {
	Key    string `orm:"pk;"`
	Title  string `orm:"unique"`
	Url    string
	Parent *Menu `orm:"rel(fk)"`
}

//添加菜单
func AddMenu(menu *Menu) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(menu)
	return id
}

//添加子菜单
func AddSubMenu(menu *SubMenu) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(menu)
	return id
}

//获取菜单
func GetAllMenu() []*Menu {
	o := orm.NewOrm()
	var menu Menu
	var menus []*Menu
	o.QueryTable(menu).All(&menus)
	return menus
}

//获取所有子菜单
func GetAllSubMenu() []*SubMenu {
	o := orm.NewOrm()
	var submenu SubMenu
	var submenus []*SubMenu
	o.QueryTable(submenu).All(&submenus)
	return submenus
}

//获取子菜单
func GetSubMenuByM(menu *Menu) []*SubMenu {
	o := orm.NewOrm()
	var submenu SubMenu
	var submenus []*SubMenu
	o.QueryTable(submenu).Filter("Parent", menu).All(&submenus)
	return submenus
}
