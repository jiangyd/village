package controllers

import (
	// "fmt"
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
	Pid      int     `json:"pid"`
	Children []*Node `json:"children"`
}

var msg *Node
var subtree []*Node

func SetNodeArray(node []*Node, p []*Node) {
	for _, n := range node {
		if n.Pid == 0 {
			msg = &Node{Name: n.Name, Id: n.Id, Pid: n.Pid, Children: []*Node{}}
		} else {
			for _, k := range node {
				if n.Pid == k.Id {
					p = append(k.Children, &Node{Name: n.Name, Id: n.Id, Pid: n.Pid, Children: []*Node{}})
				}
			}
		}
	}
	msg.Children = p
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

func (self *WiKi) WiKiRootDoc() {
	id := self.Ctx.Input.Param(":id")
	pid, _ := strconv.Atoi(id)
	doc := admin.GetDocById(pid)
	subtree = []*Node{}
	var p []*Node
	p = []*Node{}
	subtree = append(subtree, &Node{Name: doc.Title, Id: doc.Id, Pid: doc.Pid, Children: []*Node{}})

	//通过父id查出的数据
	rootdoc := admin.GetDocByPid(pid)
	GetRootTree(rootdoc)

	self.TplName = "wiki.html"
	SetNodeArray(subtree, p)
	self.Data["json"] = &msg
	self.ServeJSON()
}

func GetRootTree(doc []*admin.Document) {
	//获取子树
	for _, d := range doc {
		subtree = append(subtree, &Node{Name: d.Title, Id: d.Id, Pid: d.Pid, Children: []*Node{}})
		if len(admin.GetDocByPid(d.Id)) > 0 {
			GetRootTree(admin.GetDocByPid(d.Id))
		} else {
			return
		}
	}
}
