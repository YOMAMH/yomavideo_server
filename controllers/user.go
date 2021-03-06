package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"log"
	"yomavideo_server/models"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Id int64
	UserName string
	PassWord string
}

type Success struct {
	Code int
	Date []User
}

func (u *UserController) URLMapping()  {
	u.Mapping("singoIn", u.Create)
	u.Mapping("login", u.Login)
}

// @router /user/singoIn/ [post]
func (this *UserController) Create() {
	this.Ctx.Output.Header("Content-Type", "application/json")
	userName := this.GetString("user")
	passWord := this.GetString("password")

	var userId int64
	userId = models.CreateUser(userName, passWord)
	user := User{userId,userName,passWord}
	dateSet := []User{user}
	res, err := json.Marshal(Success{200,dateSet})

	if err != nil {log.Fatal(err)}
	this.Ctx.Output.Body(res)

}

// @router /user/login/ [post]
func (this *UserController) Login() {
	this.Ctx.Output.Header("Content-Type","application/json")
	userName := this.GetString("userName")
	passWord := this.GetString("passWord")

	user := models.FindUserByName(userName, passWord)
	userObj := User{user.Id,user.UserName,user.PassWord}
	dateSet := []User{userObj}
	res, err := json.Marshal(Success{200,dateSet})
	if err != nil {log.Fatal(err)}
	this.Ctx.Output.Body(res)
}




