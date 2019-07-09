package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tianxinbaiyun/goniushop/models"
	_ "github.com/tianxinbaiyun/goniushop/routers"
	"github.com/tianxinbaiyun/goniushop/services"
	_ "github.com/tianxinbaiyun/goniushop/utils"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = "localhost"
	beego.BConfig.Listen.HTTPPort = 8080
	beego.SetLogger("file", `{"filename":"./logs/app.log"}`)
	beego.InsertFilter("/api/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080

}
