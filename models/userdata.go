package models

import "github.com/beego/beego/v2/adapter/orm"

type Userdata struct {
	Id         int `orm:"pk"`
	Tid        string
	Vocabulary int
	Check_day  int
	Token      string
}

func (u *Userdata) TableName() string {
	return "user_data"
}

func init() {
	orm.RegisterModel(new(Userdata))
}
