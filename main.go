package main

import (
	_ "app/models"
	_ "app/routers"
	"github.com/astaxie/beego"
	"log"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)



	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
