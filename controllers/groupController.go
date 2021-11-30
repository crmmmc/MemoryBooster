package controllers

import (
	"encoding/json"
	"fmt"
	"mb/models"
	"mb/morm"
	"mb/mtool"
)

type GroupController struct {
	baseController
}

//按组请求
func (g *GroupController) Get() {
	uid := g.GetString("uid")
	group, _ := g.GetInt("group")
	l := morm.ReadByGroup(uid, group)
	g.Data["json"] = l
	g.ServeJSON()
}

//按组更新，还没测完
func (g *GroupController) Post() {

	//dict_uid
	duid := g.GetString("uid")

	if duid == "" {
		g.Ctx.WriteString("err: uid is nil")
	}

	//前端传的json
	var l []*models.Srcdstdict
	err := json.Unmarshal(g.Ctx.Input.RequestBody, &l)

	if err != nil {
		g.Ctx.WriteString("err: create json")
		fmt.Println(err.Error())
		return
	} else {
		//更新json
		derr := mtool.DictJsonUpdater(duid, l)
		if derr != nil {
			g.Ctx.WriteString("err: update json error")
			return
		} else {
			g.Ctx.WriteString("update success")
		}
	}

}
