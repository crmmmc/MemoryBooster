package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

type User struct {
	Uid           string `orm:"pk" json:"uid"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Nickname      string `json:"nickname"`
	Country       string `json:"country"`
	Sign_date     int64  `json:"sign_date"`
	Last_log_date int64  `json:"last_log_date"`
	State         int    `json:"state"`
	Current_dict  string `json:"current_dict"` //当前字典的id
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "user"
}
