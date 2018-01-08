// message
package controllers

import (
	"github.com/astaxie/beego"
	"apidemo/models"
)

type MessageController struct {
	beego.Controller
}

func (message *MessageController) Get() {
	message.Data["json"] = models.GetMessage()
	message.ServeJSON()
}
