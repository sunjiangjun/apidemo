package main

import (
	_ "apidemo/routers"

	"github.com/astaxie/beego"
	"apidemo/utils"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	utils.Debug(beego.AppConfig.String("tag"))
	beego.Run()
}
