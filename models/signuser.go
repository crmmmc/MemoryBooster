package models

import "github.com/beego/beego/v2/adapter/orm"

//注册时的user缓存
type Signuser struct {
	Email    string `orm:"pk"`
	Uid      string
	Password string
	State    int
}

func init() {
	orm.RegisterModel(new(Signuser))
}

//重写
func (s Signuser) TableName() string {
	return "sign_user"
}
