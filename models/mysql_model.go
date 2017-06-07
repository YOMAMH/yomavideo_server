package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

type User struct {
	Id int64
	UserName string
	PassWord string
}

func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@/go_db?charset=utf8",30)
	// 注册定义的 model
	orm.RegisterModel(new(User))
}

// 创建用户
func CreateUser(userStr string, passStr string) int64 {
	o := orm.NewOrm()
	user := User{UserName: userStr, PassWord: passStr}

	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {log.Fatal(err)}
	return id

}

// 查找用户
func FindUserByName(userStr string, password string) User {

	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("userName", userStr).Filter("passWord", password).One(&user)
	if err != nil {log.Fatal(err)}
	return user

}