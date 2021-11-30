package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"mb/models"
	"mb/morm"
)

type UserController struct {
	baseController
}

//更新部分字段
func (u *UserController) Patch() {

	uid := u.GetString("uid")
	password := u.GetString("password")
	nickname := u.GetString("nickname")
	country := u.GetString("country")
	nu := u.GetString("nu")

	if uid == "" {
		u.Ctx.WriteString("err: no uid")
		return
	}

	oi := orm.NewOrm()
	us := new(models.User)
	us.Uid = uid
	oi.Read(us)

	if password != nu {
		us.Password = password
	}

	if nickname != nu {
		us.Nickname = nickname
	}

	if country != nu {
		us.Country = country
	}

	fmt.Println(*us)

	err := morm.UpdateUserInfoFromClient(us)

	if err != nil {
		str := err.Error()
		str = "err: " + str
		u.Ctx.WriteString(str)
	} else {
		u.Ctx.WriteString("update profile success")
	}

	//var user models.User
	//var err error
	//if err = json.Unmarshal(u.Ctx.Input.RequestBody, &user); err == nil {
	//	morm.UpdateUserInfoFromClient(&user)
	//	u.Ctx.WriteString("update user info success")
	//} else {
	//	fmt.Println("userController patch err")
	//	fmt.Println(err.Error())
	//	u.Ctx.WriteString("userController patch err")
	//}

}
