package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"village/models/admin"
)

type RolePermission struct {
	beego.Controller
}

func (self *RolePermission) RoleManagelist() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/role.html"
	self.Data["role"] = admin.GetAllRole()
	self.Data["permission"] = admin.GetAllPermission()

}

type RoleNode struct {
	Name     string      `json:"name"`
	Id       string      `json:"id"`
	Children []*RoleNode `json:"children"`
}

var rolemsg []*RoleNode

func SetRoleNodeArray(role *admin.Permission, node []*RoleNode) {
	//判断当前树是否顶级树
	if role.Pname == "" {
		//直接把树加到数组中
		rolemsg = append(rolemsg, &RoleNode{Name: role.Description, Id: role.Name, Children: []*RoleNode{}})
		return
	} else {
		//不是顶级树
		//遍历顶级树，查找当前树的父级

		for _, v := range node {
			//如果是当前树的父级,把当前树加到父级的子树中
			if role.Pname == v.Id {
				v.Children = append(v.Children, &RoleNode{Name: role.Description, Id: role.Name, Children: []*RoleNode{}})

				return
			} else {
				//递归遍历
				SetRoleNodeArray(role, v.Children)
			}
		}
	}
	return
}

func (self *RolePermission) GetRoleDocNodes() {
	nodes := admin.GetAllPermission()
	rolemsg = []*RoleNode{}
	for _, i := range nodes {
		SetRoleNodeArray(i, rolemsg)
	}
	self.Data["json"] = &rolemsg
	self.ServeJSON()
}

//权限操作
func (self *RolePermission) RoleAction() {
	action := self.Ctx.Input.Param(":action")
	fmt.Println(action, "action")
	key, title := self.Input().Get("key"), self.Input().Get("title")
	switch action {
	case "add":
		menu := admin.Menu{Key: key, Title: title}
		admin.AddMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "添加成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "modify":
		menu := admin.GetMenuByKey(key)
		menu.Key = key
		menu.Title = title
		fmt.Println(menu, "menu")
		admin.UpdateMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "修改成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		menu := admin.GetMenuByKey(key)
		if admin.IsHasSubMenu(&menu) {
			msg := map[string]interface{}{"code": 1, "msg": "该菜单下包含子菜单,不能删除!"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			admin.DelMenu(&menu)
			msg := map[string]interface{}{"code": 0, "msg": "删除成功"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}

	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}
