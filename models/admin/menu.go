package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

//菜单表
type Menu struct {
	Key   string `orm:"pk"`
	Title string `orm:"unique"`
}

//二级菜单表
type SubMenu struct {
	Key    string `orm:"pk"`
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

//更新菜单
func UpdateMenu(menu *Menu) int64 {
	o := orm.NewOrm()
	id, err := o.Update(menu)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

//删除菜单
func DelMenu(menu *Menu) int64 {
	o := orm.NewOrm()
	id, _ := o.Delete(menu)
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

//通过key获取菜单
func GetMenuByKey(key string) Menu {
	o := orm.NewOrm()
	var menu Menu
	o.QueryTable(menu).Filter("Key", key).One(&menu)
	return menu
}

//当前菜单是否存在子菜单
func IsHasSubMenu(menu *Menu) bool {
	o := orm.NewOrm()
	var submenu SubMenu
	return o.QueryTable(submenu).Filter("Parent", menu).Exist()
}

//添加子菜单
func AddSubMenu(menu *SubMenu) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(menu)
	return id
}

//更新子菜单信息
func UpdateSubMenu(menu *SubMenu) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(menu)
	return id
}

//删除子菜单
func DelSubMenu(menu *SubMenu) int64 {
	o := orm.NewOrm()
	id, _ := o.Delete(menu)
	return id
}

//获取所有子菜单
func GetAllSubMenu() []*SubMenu {
	o := orm.NewOrm()
	var submenu SubMenu
	var submenus []*SubMenu
	o.QueryTable(submenu).RelatedSel().All(&submenus)
	return submenus
}

//通过key获取子菜单信息
func GetSubMenuByKey(key string) SubMenu {
	o := orm.NewOrm()
	var submenu SubMenu
	o.QueryTable(submenu).Filter("Key", key).RelatedSel().One(&submenu)
	return submenu
}

//获取子菜单
func GetSubMenuByM(menu *Menu) []*SubMenu {
	o := orm.NewOrm()
	var submenu SubMenu
	var submenus []*SubMenu
	o.QueryTable(submenu).Filter("Parent", menu).RelatedSel().All(&submenus)
	return submenus
}
