package controllers

import (
	"mb/models"
	"mb/morm"
)

type MilestoneController struct {
	baseController
}

//创建一个新的milestone
func (m MilestoneController) Post() {
	uuid := m.GetString("user_uid")
	name := m.GetString("name")

	if uuid == "" {
		m.Ctx.WriteString("err: no user_uid")
		return
	} else if name == "" {
		m.Ctx.WriteString("err: no name")
		return
	}

	mi := new(models.Milestone)
	mi.Tid = uuid
	mi.Name = name
	sta := morm.InsertMilestoneByUid(mi)
	m.Ctx.WriteString(sta)
}
