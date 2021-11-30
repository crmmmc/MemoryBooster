package controllers

import (
	"mb/models"
	"mb/morm"
)

//主控
type ProfileController struct {
	baseController
}

//通过uid请求profile
func (m *ProfileController) Get() {

	user := new(models.User)

	//获取json
	uid := m.GetString("uid")
	user.Uid = uid
	profile := morm.GetProfile(user)
	m.Data["json"] = profile
	if uid != "" {
		m.ServeJSON()
	} else {
		m.Ctx.WriteString("uid is empty")
	}

}
