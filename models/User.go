package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int
	Nickname string `orm:"unique"`
	Password string
	Email    string `orm:"unique"`
	Avatar   string
	Ctime    time.Time `orm:"auto_now_add;type(datetime)"`
	Sign     string    `orm:"size(255)"`
	Sex      int       //0:男,1:女，2:保密
	City     string
	Status   int `orm:"default(0)"` //0:启用 1:禁用
	Secret   string
	Mfa      bool `orm:"default(false)"`
}

//获取所有用户
func GetAllUser() []*User {
	o := orm.NewOrm()
	var user User
	var users []*User
	o.QueryTable(user).OrderBy("-Ctime").All(&users)
	return users
}

//检查用户名密码
func CheckLogin(email, password string) User {
	o := orm.NewOrm()
	var user User
	o.QueryTable(user).Filter("Email", email).Filter("Password", password).RelatedSel().One(&user)
	return user
}

//查询用户是否存在
func IsUserExit(uid *User) bool {
	o := orm.NewOrm()
	var user User
	return o.QueryTable(user).Filter("Id", uid).Exist()

}

func IsUserExitByEmail(email string) bool {
	o := orm.NewOrm()
	var user User
	return o.QueryTable(user).Filter("Email", email).Exist()
}

//查看用户详细信息
func FindUserDetialById(id int) User {
	o := orm.NewOrm()
	var user User
	o.QueryTable(user).Filter("Id", id).One(&user)
	return user
}

func AddUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

//更新用户信息
func UpdateUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(user)
	return id
}

func FindUserByEmail(email string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Email", email).One(&user)
	return err != orm.ErrNoRows, user
}

func NewUser() []*User {
	o := orm.NewOrm()
	var user User
	var users []*User
	o.QueryTable(user).OrderBy("-Ctime").Limit(12).All(&users)
	return users
}
