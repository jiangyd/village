package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"village/models/admin"
)

type WiKi struct {
	beego.Controller
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

func (self *WiKi) WiKiPage() {
	self.TplName = "wiki.html"
	self.Data["root"] = admin.GetRootDoc()
}

func (self *WiKi) WiKiDoc() {
	id := self.Ctx.Input.Param(":id")
	nid, _ := strconv.Atoi(id)
	doc := admin.GetDocById(nid)
	self.Layout = "wiki.html"
	self.TplName = "doc.html"
	self.Data["doc"] = doc
}
