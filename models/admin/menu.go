package admin

import (
	""
)


//菜单表
type Menu struct{
	Id int
	Key string
	title string
}

//二级菜单表
type SubMenu struct{
	Id int
	Key string
	title string
	Url string
	Parent *Menu
}