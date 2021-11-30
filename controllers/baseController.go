package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"mb/mtool"
)

type baseController struct {
	beego.Controller
	isLogin bool
}

//验证是否登录
func (b *baseController) Prepare() {
	tokenString := b.GetString("token")
	tokenUs, err := mtool.ValidateToken(tokenString)

	//无token
	if tokenString == "" {
		b.isLogin = false
		//b.Ctx.WriteString("err: No token")
		b.check()
	}

	if err != nil {
		//验证失败
		b.isLogin = false
		b.Ctx.WriteString("err: Token valid false")
	} else {
		//成功
		b.isLogin = true
		fmt.Println(tokenUs)
	}
	//check()
	b.check()
}

func (b *baseController) check() {
	if !b.isLogin {
		b.StopRun()
	}
}
