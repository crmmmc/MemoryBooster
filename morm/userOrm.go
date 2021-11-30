package morm

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"mb/models"
	"mb/mtool"
	"time"
)

//新增,返回uid
func InsertUserNewUUID(u *models.User) string {

	var us models.User

	us.Uid = mtool.GetMyUUID()
	us.Email = u.Email
	us.Password = u.Password
	us.Nickname = u.Nickname
	us.Country = u.Country
	us.Sign_date = time.Now().Unix()
	us.Last_log_date = time.Now().Unix()
	us.State = 1

	//新建词典
	us.Current_dict = AutoInsertUserdict(&us)

	var oi = orm.NewOrm()
	id, err := oi.Insert(&us)
	if err == nil {
		fmt.Println("us insert success: id= ", id)
	} else {
		fmt.Println(err)
	}

	//新建Userdata
	userdata := models.Userdata{
		Tid:        us.Uid,
		Vocabulary: 0,
		Check_day:  0,
		Token:      "notoken",
	}

	oi.Insert(&userdata)

	return us.Uid
}

//read
func ReadUser(str string) *models.User {

	oi := orm.NewOrm()
	us := new(models.User)
	us.Uid = str

	err := oi.Read(us)

	if err != nil {
		fmt.Println("ReadUser() err")
	}

	return us

}

//update
func UpdateUser(u *models.User) error {
	oi := orm.NewOrm()
	_, err := oi.Update(u)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

//指定更新password,nickname,country
func UpdateUserInfoFromClient(u *models.User) error {

	oi := orm.NewOrm()
	_, err := oi.Update(u, "Password", "Nickname", "Country")
	if err != nil {
		fmt.Println(err)
	}

	return err
}

//指定更新词书
func UpdateCurrentDictFromClient(u *models.User) {
	oi := orm.NewOrm()
	_, err := oi.Update(u, "Current_dict")
	if err != nil {
		fmt.Println(err)
	}
}

//建立连接
func init() {
	mtool.GetDbConn()
}
