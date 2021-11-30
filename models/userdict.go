package models

import "github.com/beego/beego/v2/adapter/orm"

//用户词典
type Userdict struct {
	Uid       string `orm:"pk"`
	Tid       string
	Name      string
	Level     int
	Dict_json string
}

func init() {
	orm.RegisterModel(new(Userdict))
}

//重写
func (u Userdict) TableName() string {
	return "user_dict"
}
