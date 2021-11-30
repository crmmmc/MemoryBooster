package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
	"mb/models"
	"mb/morm"
	"mb/mtool"
	"regexp"
)

type SignController struct {
	beego.Controller
}

//接收邮箱和密码，给前端验证码，并且服务器发送邮件给该邮箱
func (s *SignController) Post() {
	email := s.GetString("email")
	password := s.GetString("password")

	muser := new(models.User)
	muser.Email = email
	oi := orm.NewOrm()
	oierr := oi.Read(muser, "Email")

	fmt.Println(oierr)

	if oierr == nil {
		s.Ctx.WriteString("err: email address <" + email + "> is exist")
		return
	}

	reg, _ := regexp.MatchString("@", email)
	rege, _ := regexp.MatchString("[.]", email)

	if !reg {

		//err
		s.Data["json"] = "email address must have '@'"
		s.ServeJSON()
		return

	} else if !rege {

		//err
		s.Data["json"] = "email address must have '.'"
		s.ServeJSON()
		return

	} else {

		//前端获取验证码
		vc := morm.UpdateOrCreateSignUserAndReturnVerificationCode(email, password)
		s.Data["json"] = vc
		s.ServeJSON()

		//服务器发送邮件给该邮箱
		str := mtool.SendEmailByAddressQQ(email, vc)
		fmt.Println(str)
	}
}

//验证
func (s *SignController) Put() {
	email := s.GetString("email")
	vc := s.GetString("vcode")

	suser, err := morm.ReadByEmail(email)
	if err == nil {
		suvc := suser.Uid[len(suser.Uid)-4 : len(suser.Uid)]
		if vc == suvc {
			user := new(models.User)
			user.Uid = suser.Uid
			user.Email = suser.Email
			user.State = 1

			//插入user表中
			uidstr := morm.InsertUserNewUUID(user)

			//在sign表中删除
			delstr := morm.DeleteSignUserByUid(suser.Email)

			type sjson struct {
				Uid   string
				Token string
			}

			var tuser mtool.TUser
			tuser.Email = email
			tuser.Uid = uidstr
			tok, _ := mtool.GenerateToken(&tuser, 0)

			sj := sjson{
				Uid:   uidstr,
				Token: tok,
			}

			if delstr == "success" {
				s.Data["json"] = sj
				s.ServeJSON()
			}

		} else {
			s.Data["json"] = "100"
			s.ServeJSON()
		}
	} else {
		s.Data["json"] = "101"
		s.ServeJSON()
	}

}

//注册个人信息
func (s *SignController) Patch() {

	uid := s.GetString("uid")
	password := s.GetString("password")
	nickname := s.GetString("nickname")
	country := s.GetString("country")

	if uid == "" {
		s.Ctx.WriteString("err: uid is nil")
		return
	} else if password == "" {
		s.Ctx.WriteString("err: password is nil")
		return
	} else if country == "" {
		s.Ctx.WriteString("err: country is nil")
		return
	} else if nickname == "" {
		s.Ctx.WriteString("err: nickname is nil")
	}

	u := morm.ReadUser(uid)
	u.Password = password
	u.Nickname = nickname
	u.Country = country

	var usToken = mtool.TUser{
		Uid:   uid,
		Email: u.Email,
	}

	token, terr := mtool.GenerateToken(&usToken, 0)

	type sjson struct {
		Uid   string
		Token string
	}

	if terr != nil {
		s.Ctx.WriteString("set err: set token error")
		return
	}

	err := morm.UpdateUser(u)
	if err != nil {
		s.Ctx.WriteString("err to sign")
	} else {
		sj := sjson{
			Uid:   uid,
			Token: token,
		}
		fmt.Println(sj)
		s.Data["json"] = sj
		s.ServeJSON()

	}
}
