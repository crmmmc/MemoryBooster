package morm

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"mb/models"
)

/*
不用导入orm包！
*/

func GetProfile(u *models.User) *models.Profile {

	u = ReadUser(u.Uid)

	var uid = u.Uid
	var dlist []*models.Userdata
	var mlist []*models.Milestone

	oi := orm.NewOrm()

	_, errd := oi.Raw("select * from user_data where tid = ?", uid).QueryRows(&dlist)
	if errd != nil {
		fmt.Println("GetProfile() err")
	}

	_, errm := oi.Raw("select * from milestone where tid = ?", uid).QueryRows(&mlist)
	if errm != nil {
		fmt.Println("GetProfile() err")
	}

	var p models.Profile
	p.User = u
	p.UserData = dlist
	p.Milestone = mlist

	return &p
}
