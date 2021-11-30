package models

import "github.com/beego/beego/v2/adapter/orm"

type Accuracy struct {
	Id    int
	Tid   string
	Rate  int
	Group int
}

func (a *Accuracy) TableName() string {
	return "accuracy"
}

func init() {
	orm.RegisterModel(new(Accuracy))
}
