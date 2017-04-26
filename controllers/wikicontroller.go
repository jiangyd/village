package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang-commonmark/markdown"
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

var msg []*Node
var subtree []*Node

func SetNodeArray(doc *Node, node []*Node) {
	//判断当前树是否顶级树
	if doc.Pid == 0 {
		//直接把树加到数组中
		msg = append(msg, &Node{Name: doc.Name, Id: doc.Id, Pid: doc.Pid, Children: []*Node{}})

		return
	} else {
		//不是顶级树
		//遍历顶级树，查找当前树的父级
		for _, v := range node {
			//如果是当前树的父级,把当前树加到父级的子树中
			if doc.Pid == v.Id {
				v.Children = append(v.Children, &Node{Name: doc.Name, Id: doc.Id, Pid: doc.Pid, Children: []*Node{}})

				return
			} else {
				//递归遍历
				SetNodeArray(doc, v.Children)
			}
		}
	}
	return
}

func (self *WiKi) WiKiList() {
	self.TplName = "wikilist.html"
	self.Data["root"] = admin.GetRootDoc()
}

func (self *WiKi) WiKiDetial() {
	id := self.Ctx.Input.Param(":id")
	fmt.Println(id, "iiiiiiii")
	root := self.Ctx.Input.Param(":root")
	nid, _ := strconv.Atoi(id)
	doc := admin.GetDocById(nid)
	self.Layout = "wikitree.html"
	self.TplName = "wikidetial.html"
	self.Data["id"] = root
	self.Data["selectnode"] = id
	self.Data["doc"] = doc
	md := markdown.New(markdown.HTML(true))
	cc := md.RenderToString([]byte(doc.Content))
	fmt.Println(cc)
	self.Data["content"] = cc

}

func (self *WiKi) WiKiRoot() {
	id := self.Ctx.Input.Param(":id")
	self.Data["id"] = id
	self.TplName = "wikitree.html"
}

func (self *WiKi) WiKiRootTree() {
	id := self.Ctx.Input.Param(":id")
	pid, _ := strconv.Atoi(id)

	doc := admin.GetDocById(pid)
	subtree = []*Node{}
	msg = []*Node{}
	subtree = append(subtree, &Node{Name: doc.Title, Id: doc.Id, Pid: doc.Pid, Children: []*Node{}})

	//通过父id查出的数据
	rootdoc := admin.GetDocByPid(pid)
	GetRootTree(rootdoc)
	// for _, s := range subtree {
	// 	fmt.Println(s.Name)
	// }
	for _, k := range subtree {
		SetNodeArray(k, msg)
	}

	self.Data["json"] = &msg
	self.ServeJSON()
}

func GetRootTree(doc []*admin.Document) {
	//获取子树
	for _, d := range doc {
		fmt.Println(d.Title)
		subtree = append(subtree, &Node{Name: d.Title, Id: d.Id, Pid: d.Pid, Children: []*Node{}})
		if len(admin.GetDocByPid(d.Id)) > 0 {
			fmt.Println(d.Title, "yyyyy")
			GetRootTree(admin.GetDocByPid(d.Id))
		}
	}

}
