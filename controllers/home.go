package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

type ResDate struct {
	Code int
	Msg string
	Date interface{}
}

func (h *HomeController) URLMapping() {
	h.Mapping("carouses", h.Carouses)
}

// @router /carouses/ [get]
func (this *HomeController) Carouses() {
	this.Ctx.Output.Header("Content-Type", "application/json")
	
}