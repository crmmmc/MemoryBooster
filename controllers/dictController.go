package controllers

import (
	"mb/models"
	"mb/mtool"
)

type DictController struct {
	baseController
}

//获取
func (d *DictController) Get() {

	dict := new(models.Userdict)
	uid := d.GetString("dict_uid")

	if uid == "" {
		d.Ctx.WriteString("err: no uid")
		return
	}

	dict.Uid = uid
	j := mtool.DictJsonReader(dict.Uid)
	d.Data["json"] = j
	d.ServeJSON()

}
