package admin

import (
	"github.com/astaxie/beego/orm"
)

type Permission struct {
	Name        string  `orm:"pk"`
	Description string  `orm:"unique"`
	Pname       string  `orm:"null"` //父权限
	Roles       []*Role `orm:"reverse(many)"`
}

func GetAllPermission() []*Permission {
	o := orm.NewOrm()
	var permission Permission
	var permissions []*Permission
	o.QueryTable(permission).OrderBy("Pname").RelatedSel().All(&permissions)
	return permissions
}
