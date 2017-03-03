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
}

//检查用户名密码
func CheckLogin(email, password string) User {
	o := orm.NewOrm()
	var user User
	o.QueryTable(user).Filter("Email", email).Filter("Password", password).RelatedSel().One(&user)
	return user
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
