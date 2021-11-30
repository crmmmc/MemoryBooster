package controllers

import (
	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
	"mb/models"
	"mb/mtool"
)

type LogController struct {
	beego.Controller
}

//登录
func (l *LogController) Post() {

	email := l.GetString("email")
	password := l.GetString("password")
	oi := orm.NewOrm()

	//验证
	//用户为空
	if email == "" || password == "" {
		l.Ctx.WriteString("log err: email or password can not be nil")
		return
	}

	//用户不存在
	us := new(models.User)
	us.Email = email
	err := oi.Read(us, "Email") //读取用户的全部信息
	if err != nil {
		l.Ctx.WriteString("log err: not have the user")
		return
	}

	//密码错误
	if password != us.Password {
		l.Ctx.WriteString("log err: wrong password")
		return
	} else {
		//登录成功
		//l.Ctx.WriteString("log success")
		//设置token
		var tokenUs = mtool.TUser{
			Uid:   us.Uid,
			Email: us.Email,
		} //tokenUser
		var token string

		token, terr := mtool.GenerateToken(&tokenUs, 0)
		if terr != nil {
			l.Ctx.WriteString("set err: set token error")
			return
		} else {
			userdata := new(models.Userdata)
			userdata.Tid = us.Uid
			oi.Read(userdata, "Tid")
			userdata.Token = token
			oi.Update(userdata, "Token")

			type sjson struct {
				Uid   string
				Token string
			}

			sj := sjson{
				Uid:   us.Uid,
				Token: token,
			}

			l.Data["json"] = sj
			l.ServeJSON()
		}
	}

}
