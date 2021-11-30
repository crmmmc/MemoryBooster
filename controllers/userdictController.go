package controllers

import (
	"github.com/beego/beego/v2/adapter/orm"
	"mb/models"
	"mb/morm"
)

type UserdictController struct {
	baseController
}

/*更新一个用户的当前dict
user uid，dict uid
*/
func (u *UserdictController) Patch() {

	userUid := u.GetString("user_uid")
	dictUid := u.GetString("dict_uid")
	user := models.User{Uid: userUid, Current_dict: dictUid}
	oi := orm.NewOrm()

	if userUid == "" {
		u.Ctx.WriteString("err: no user_uid")
		return
	} else if dictUid == "" {
		u.Ctx.WriteString("err: no dict_uid")
		return
	}

	var have bool
	dictList := morm.ReadByTid(userUid)
	for i := 0; i < len(dictList); i++ {
		if dictList[i].Uid == dictUid {
			have = true
		}
	}

	if have {
		//限制更新current_dict字段
		_, err := oi.Update(&user, "Current_dict")
		if err != nil {
			u.Ctx.WriteString("update err")
		} else {
			u.Ctx.WriteString("update success")
		}
	} else {
		u.Ctx.WriteString("the user not have this dict")
	}

}
