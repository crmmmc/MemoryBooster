package morm

import (
	"github.com/beego/beego/v2/adapter/orm"
	"mb/models"
	"mb/mtool"
)

func InsertMilestoneByUid(mi *models.Milestone) string {

	oi := orm.NewOrm()
	_, err := oi.Insert(mi)

	if err != nil {
		return "insert err"
	} else {
		return "insert success"
	}

}

//建立连接
func init() {
	mtool.GetDbConn()
}
