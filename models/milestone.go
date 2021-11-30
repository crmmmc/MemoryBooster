package models

import "github.com/beego/beego/v2/adapter/orm"

type Milestone struct {
	Id   int `orm:"pk"`
	Tid  string
	Name string
}

func (m *Milestone) TableName() string {
	return "milestone"
}

func init() {
	orm.RegisterModel(new(Milestone))
}
