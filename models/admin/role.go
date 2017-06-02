package admin

import (
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id       int
	Rolename string
	Role     *SubMenu `orm:"rel(fk)"`
}

//添加角色

func AddRole(role *Role) {
	o := orm.NewOrm()
	o.Insert(role)
}
