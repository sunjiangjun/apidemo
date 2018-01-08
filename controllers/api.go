package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

//@router /v1/api/ [get]
func (this *ApiController) name()  {



	this.Data["json"]="hello"
	this.ServeJSON()
}