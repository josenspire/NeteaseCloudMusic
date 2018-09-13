package main

import (
	_ "NeteaseCloudMusic/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/image", "./static/img")	// default will setup static folder, need to setup static second directory
	beego.Run()
}
