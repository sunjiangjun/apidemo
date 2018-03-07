// message
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sunjiangjun/apidemo/models"
)

type MessageController struct {
	beego.Controller
}

func (message *MessageController) Get() {
	message.Data["json"] = models.GetMessage()
	message.ServeJSON()
}
