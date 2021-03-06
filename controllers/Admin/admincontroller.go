package admin

import (
	// "encoding/json"
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
	self.TplName = "admin/login.html"
}

func (self *Admin) AdminPost() {
	email, password, vercode, captcha_id := self.Input().Get("email"), self.Input().Get("password"), self.Input().Get("vercode"), self.Input().Get("captcha_id")
	// if !CheckCode(vercode, captcha_id) {
	// 	msg := map[string]interface{}{"code": 1, "msg": "验证码错误"}
	// 	self.Data["json"] = &msg
	// 	self.ServeJSON()
	// 	return
	// }
	fmt.Println(vercode, captcha_id)
	userinfo := models.CheckLogin(email, password, true)

	if len(userinfo.Email) > 0 {
		self.SetSession("email", userinfo.Email)
		if userinfo.Mfa {
			msg := map[string]interface{}{"code": 0, "mfa": true}
			self.Data["json"] = &msg
			self.ServeJSON()
			return
		}
		self.SetSession("uid", userinfo.Id)
		self.SetSession("nickname", userinfo.Nickname)
		msg := map[string]interface{}{"code": 0, "msg": "success"}
		self.Data["json"] = &msg
		self.ServeJSON()
		return
	} else {
		msg := map[string]interface{}{"code": 1, "msg": "用户名或密码错误!"}
		self.Data["json"] = &msg
		self.ServeJSON()
		return

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
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/document.html"
}

//站点页面
func (self *Admin) SiteManageList() {
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Data["site"] = models.GetAllSite()
	self.Data["sitecategory"] = models.GetCiteCategory()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/site.html"
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

//编辑站点页面
func (self *Admin) GetSiteInfo() {
	siteid := self.Input().Get("siteid")
	sid, _ := strconv.Atoi(siteid)
	self.Data["sitecategory"] = models.GetCiteCategory()
	self.Data["site"] = models.FindSiteById(sid)
	self.TplName = "admin/editsite.html"
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

//站点操作
func (self *Admin) SiteAction() {
	action := self.Ctx.Input.Param(":action")
	uid := self.GetSession("uid")
	switch action {
	case "add":
		categoryid, url, title, content, img := self.Input().Get("category"), self.Input().Get("url"), self.Input().Get("title"), self.Input().Get("content"), self.Input().Get("img")
		fmt.Println(img, "img")
		cid, _ := strconv.Atoi(categoryid)
		site := models.Sites{Category: &models.Categorys{Id: cid}, Url: url, Title: title, Content: content, Img: img, User: &models.User{Id: uid.(int)}}
		models.AddSite(&site)
		msg := map[string]interface{}{"code": 0, "msg": "添加成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "modify":
		id, categoryid, url, title, content, img := self.Input().Get("id"), self.Input().Get("category"), self.Input().Get("url"), self.Input().Get("title"), self.Input().Get("content"), self.Input().Get("img")
		cid, _ := strconv.Atoi(categoryid)
		sid, _ := strconv.Atoi(id)
		site := models.FindSiteById(sid)
		site.User = &models.User{Id: uid.(int)}
		site.Category = &models.Categorys{Id: cid}
		site.Url = url
		site.Title = title
		site.Content = content
		site.Img = img
		models.UpdateSite(&site)
		msg := map[string]interface{}{"code": 0, "msg": "修改成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		id := self.Input().Get("id")
		sid, _ := strconv.Atoi(id)
		models.DelSite(sid)
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

//节点操作
func (self *Admin) DocNodeAction() {
	action := self.Ctx.Input.Param(":action")
	switch action {
	case "add":
		id, node, content := self.Input().Get("pid"), self.Input().Get("title"), self.Input().Get("content")
		pid, _ := strconv.Atoi(id)
		doc := admin.Document{Title: node, Pid: pid, Content: content}
		admin.AddDoc(&doc)
		msg := map[string]interface{}{"code": 0, "msg": "节点添加成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "modify":
		id, title, content := self.Input().Get("id"), self.Input().Get("title"), self.Input().Get("content")
		nid, _ := strconv.Atoi(id)
		doc := admin.GetDocById(nid)
		doc.Title = title
		doc.Content = content
		admin.UpdateDoc(&doc)
		msg := map[string]interface{}{"code": 0, "msg": "节点编辑成功"}
		self.Data["json"] = &msg
		self.ServeJSON()
	case "del":
		id := self.Input().Get("id")
		nid, _ := strconv.Atoi(id)
		if admin.IsExitSubDoc(nid) {
			msg := map[string]interface{}{"code": 1, "msg": "当前节点下存在子节点,无法删除!"}
			self.Data["json"] = &msg
			self.ServeJSON()
		} else {
			admin.DelDoc(nid)
			msg := map[string]interface{}{"code": 0, "msg": "节点删除成功"}
			self.Data["json"] = &msg
			self.ServeJSON()
		}
	default:
		msg := map[string]interface{}{"code": 1, "msg": "未找到方法"}
		self.Data["json"] = &msg
		self.ServeJSON()
	}
}

//添加目录页面
func (self *Admin) AddDocumentPage() {
	pid := self.Ctx.Input.Param(":pid")
	self.Data["pid"] = pid
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Layout = "admin/nav.html"
	self.TplName = "admin/adddocument.html"

}

//编辑目录页面
func (self *Admin) EditDocumentPage() {
	id := self.Ctx.Input.Param(":id")
	nid, _ := strconv.Atoi(id)
	self.Data["menu"] = admin.GetAllMenu()
	self.Data["submenu"] = admin.GetAllSubMenu()
	self.Data["node"] = admin.GetDocById(nid)
	self.Layout = "admin/nav.html"
	self.TplName = "admin/editdocument.html"

}

type Node struct {
	Name     string  `json:"name"`
	Id       int     `json:"id"`
	Children []*Node `json:"children"`
}

var msg []*Node

func SetNodeArray(doc *admin.Document, node []*Node) {
	//判断当前树是否顶级树
	if doc.Pid == 0 {
		//直接把树加到数组中
		msg = append(msg, &Node{Name: doc.Title, Id: doc.Id, Children: []*Node{}})

		return
	} else {
		//不是顶级树
		//遍历顶级树，查找当前树的父级
		for _, v := range node {
			//如果是当前树的父级,把当前树加到父级的子树中
			if doc.Pid == v.Id {
				v.Children = append(v.Children, &Node{Name: doc.Title, Id: doc.Id, Children: []*Node{}})

				return
			} else {
				//递归遍历
				SetNodeArray(doc, v.Children)
			}
		}
	}
	return
}

func (self *Admin) GetDocNodes() {
	nodes := admin.GetDoc()
	msg = []*Node{}
	for _, i := range nodes {
		SetNodeArray(i, msg)
	}
	self.Data["json"] = &msg
	self.ServeJSON()
}
