package morm

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"mb/models"
	"mb/mtool"
)

//插入一个signuser并且返回验证码
func UpdateOrCreateSignUserAndReturnVerificationCode(email string, password string) string {

	//读取
	su := ReadOrCreateSignUserByEmail(email)

	//upadte
	oi := orm.NewOrm()
	su.Uid = mtool.GetMyUUID()
	su.Email = email
	su.Password = password

	oi.Update(su)

	//获取验证码
	vc := su.Uid[len(su.Uid)-4 : len(su.Uid)]

	return vc
}

//用email read
func ReadOrCreateSignUserByEmail(email string) *models.Signuser {

	oi := orm.NewOrm()
	su := new(models.Signuser)
	su.Email = email

	err := oi.Read(su, "Email")

	if err != nil {
		fmt.Println("ReadSignUser() err :", err.Error())
		fmt.Println("create new:")

		//新建
		su.Uid = "1000000000"
		su.State = 101
		_, errr := oi.Insert(su)
		if errr != nil {
			fmt.Println(errr.Error())
		}
		return su
	}

	//返回
	return su

}

//验证函数，用email读取
func ReadByEmail(email string) (*models.Signuser, error) {

	oi := orm.NewOrm()
	su := new(models.Signuser)
	su.Email = email

	err := oi.Read(su, "Email")

	if err != nil {
		return nil, err
	}

	//返回
	return su, nil
}

//删除
func DeleteSignUserByUid(email string) string {
	oi := orm.NewOrm()
	su := new(models.Signuser)
	su.Email = email

	_, err := oi.Delete(su)
	if err != nil {
		fmt.Println("delete signuser err :", err.Error())
		return err.Error()
	} else {
		return "success"
	}
}

//建立连接
func init() {
	mtool.GetDbConn()
}
