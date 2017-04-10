package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"village/models"
	"village/models/admin"
)

type Admin struct {
	beego.Controller
}

func (self *Admin) Login() {
	self.TplName = "admin/menu.html"
	self.Data["menu"] = admin.GetAllMenu()
	fmt.Println(admin.GetAllSubMenu())
	self.Data["submenu"] = admin.GetAllSubMenu()
	for a, b := range admin.GetAllSubMenu() {
		fmt.Println(a, "a")
		fmt.Println(b.Parent.Key, "key")
		fmt.Println(b.Parent.Title, "title")
	}
}

//菜单页面
func (self *Admin) MenuManageList() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/menu.html"
}

//子菜单页面
func (self *Admin) SubMenuManageList() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/submenu.html"
}

//用户页面
func (self *Admin) UserManageList() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/user.html"
	self.Data["users"] = models.GetAllUser()
}

//评论页面
func (self *Admin) ReplyManageList() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/reply.html"
	self.Data["replys"] = models.GetAllReply()
}

//分类页面
func (self *Admin) CategoryManageList() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/category.html"
	self.Data["categorys"] = models.GetAllCategory()
}

//帖子页面
func (self *Admin) TopicManageList() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/topic.html"
	self.Data["topics"] = models.GetAllTopic()
}

//文档页面
func (self *Admin) DocumentInfo() {
	self.Layout = "admin/nav.html"
	self.TplName = "admin/document.html"
}

//编辑子菜单页面
func (self *Admin) GetSubMenuInfo() {
	key := self.Input().Get("key")
	submenu := admin.GetSubMenuByKey(key)
	self.TplName = "admin/editsubmenu.html"
	self.Data["submenu"] = submenu
	self.Data["menu"] = admin.GetAllMenu()
	// msg := map[string]interface{}{"code": 0, "msg": "", "data": map[string]interface{}{"key": submenu.Key, "title": submenu.Title, "parent": submenu.Parent, "url": submenu.Url}}
	// self.Data["json"] = &msg
	// self.ServeJSON()
}

//编辑菜单页面
func (self *Admin) GetMenuInfo() {
	key := self.Input().Get("key")
	self.Data["menu"] = admin.GetMenuByKey(key)
	self.TplName = "admin/editmenu.html"

	// msg := map[string]interface{}{"code": 0, "msg": "", "data": map[string]interface{}{"key": submenu.Key, "title": submenu.Title, "parent": submenu.Parent, "url": submenu.Url}}
	// self.Data["json"] = &msg
	// self.ServeJSON()
}

//编辑分类页面
func (self *Admin) GetCategoryInfo() {
	categoryid := self.Input().Get("categoryid")
	id, _ := strconv.Atoi(categoryid)
	self.TplName = "admin/editcategory.html"
	self.Data["category"] = models.FindCategory(id)
}

//菜单操作
func (self *Admin) MenuAction() {
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

//子菜单操作
func (self *Admin) SubMenuAction() {
	action := self.Ctx.Input.Param(":action")
	fmt.Println(action, "action")
	parent, key, title, url := self.Input().Get("parent"), self.Input().Get("key"), self.Input().Get("title"), self.Input().Get("url")
	switch action {
	case "add":
		submenu := admin.SubMenu{Key: key, Title: title, Url: url, Parent: &admin.Menu{Key: parent}}
		admin.AddSubMenu(&submenu)
		msg := map[string]interface{}{"code": 0, "msg": "添加成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "modify":
		menu := admin.GetSubMenuByKey(key)
		menu.Key = key
		menu.Title = title
		menu.Parent = &admin.Menu{Key: parent}
		menu.Url = url
		fmt.Println(menu, "aaaa")
		admin.UpdateSubMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "修改成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		menu := admin.GetSubMenuByKey(key)
		admin.DelSubMenu(&menu)
		msg := map[string]interface{}{"code": 0, "msg": "删除成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

//分类操作
func (self *Admin) CategoryAction() {
	action := self.Ctx.Input.Param(":action")
	switch action {
	case "add":
		category, categorytype := self.Input().Get("category"), self.Input().Get("categorytype")
		categorys := models.Categorys{Category: category, CategoryType: categorytype}
		models.AddCategory(&categorys)
		msg := map[string]interface{}{"code": 0, "msg": "添加成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "modify":
		categoryid, category, categorytype := self.Input().Get("categoryid"), self.Input().Get("category"), self.Input().Get("categorytype")
		id, _ := strconv.Atoi(categoryid)
		categorys := models.FindCategory(id)
		categorys.Category = category
		categorys.CategoryType = categorytype
		fmt.Println(categorys)
		models.UpdateCategory(&categorys)
		msg := map[string]interface{}{"code": 0, "msg": "修改成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		categoryid := self.Input().Get("categoryid")
		id, _ := strconv.Atoi(categoryid)
		categorys := models.FindCategory(id)
		models.DelCategory(&categorys)
		msg := map[string]interface{}{"code": 0, "msg": "删除成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

//帖子操作
func (self *Admin) TopicAction() {
	action := self.Ctx.Input.Param(":action")
	switch action {
	case "disable":
		id := self.Input().Get("id")
		tid, _ := strconv.Atoi(id)
		topic := models.FindTopicById(tid)
		topic.Disable = true
		models.UpdateTopic(&topic)
		msg := map[string]interface{}{"code": 0, "msg": "屏蔽帖子成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "enable":
		id := self.Input().Get("id")
		tid, _ := strconv.Atoi(id)
		topic := models.FindTopicById(tid)
		topic.Disable = false
		models.UpdateTopic(&topic)
		msg := map[string]interface{}{"code": 0, "msg": "取消屏蔽帖子成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

//评论操作
func (self *Admin) ReplyAction() {
	action := self.Ctx.Input.Param(":action")
	switch action {
	case "disable":
		id := self.Input().Get("id")
		rid, _ := strconv.Atoi(id)
		reply := models.FindReplyByRid(rid)
		reply.Disable = true
		models.UpdateReply(&reply)
		msg := map[string]interface{}{"code": 0, "msg": "屏蔽评论成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "enable":
		id := self.Input().Get("id")
		rid, _ := strconv.Atoi(id)
		reply := models.FindReplyByRid(rid)
		reply.Disable = false
		models.UpdateReply(&reply)
		msg := map[string]interface{}{"code": 0, "msg": "屏蔽评论成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}

//用户操作
func (self *Admin) UserAction() {
	action := self.Ctx.Input.Param(":action")
	switch action {
	case "disable":
		id := self.Input().Get("id")
		uid, _ := strconv.Atoi(id)
		user := models.FindUserDetialById(uid)
		user.Status = 1
		models.UpdateUser(&user)
		msg := map[string]interface{}{"code": 0, "msg": "用户禁用成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "enable":
		id := self.Input().Get("id")
		uid, _ := strconv.Atoi(id)
		user := models.FindUserDetialById(uid)
		user.Status = 0
		models.UpdateUser(&user)
		msg := map[string]interface{}{"code": 0, "msg": "用户启用成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}

}
