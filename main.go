package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "mb/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {

		//开启session
		beego.BConfig.WebConfig.Session.SessionOn = true
		beego.BConfig.WebConfig.Session.SessionName = "mbsession"
		beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 36000
		beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 36000

		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	}
	beego.Run()
}
