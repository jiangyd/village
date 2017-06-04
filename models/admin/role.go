package admin

import (
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id          int
	Rolename    string        //权限名称
	Permissions []*Permission `orm:"rel(m2m)"`
	Disable     bool          `orm:"default(false)"`
}

//添加角色
func AddRole(role *Role) {
	o := orm.NewOrm()
	o.Insert(role)

}

//修改角色
func UpdateRole(role *Role) {
	o := orm.NewOrm()
	o.Update(role)
}

//删除角色
func DelRole(role *Role) {
	o := orm.NewOrm()
	o.Delete(role)
}

//通过角色名称查找
func FindRoleByName(rolename string) Role {
	o := orm.NewOrm()
	var role Role
	o.QueryTable(role).Filter("Rolename", rolename).One(&role)
	return role
}

//获取所有角色
func GetAllRole() []*Role {
	o := orm.NewOrm()
	var role Role
	var roles []*Role
	o.QueryTable(role).All(&roles)
	return roles
}
