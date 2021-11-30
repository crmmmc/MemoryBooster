package controllers

import (
	"fmt"
	"mb/morm"
	"mb/mtool"
)

type GroupToolController struct {
	baseController
}

//请求背过的组
func (gt *GroupToolController) Get() {
	uid := gt.GetString("uid")

	if uid == "" {
		gt.Ctx.WriteString("err: no uid")
		return
	}

	var s []int
	//读取dict
	l := mtool.DictJsonReader(uid)
	for i := 0; i < len(l); i++ {
		//1-99代表背过
		if l[i].State > 0 && l[i].State < 100 {
			var g = l[i].Group
			var num int = 0
			for i1 := 0; i1 < len(s); i1++ {
				if s[i1] == g {
					num++
				}
			}
			if num == 0 {
				s = append(s, g)
			}
		}
	}
	gt.Data["json"] = s
	gt.ServeJSON()
}

//前端请求一个新的组，服务器把这个组的第一个单词设为背过
func (gt *GroupToolController) Put() {

	uid := gt.GetString("uid")

	if uid == "" {
		gt.Ctx.WriteString("err: no uid")
		return
	}

	var n int = 0
	var g int = 1
	var b bool
	l := mtool.DictJsonReader(uid)

	for i := 0; i < len(l); i++ {

		//计数
		if l[i].Group == g {
			if l[i].State == 0 {
				n++
			}
		} else {
			g = l[i].Group
			n = 0
			if l[i].State == 0 {
				n++
			}
		}

		//判断,返回下一个group,update
		if n == 10 {

			l[i-9].State = 1

			//update
			mtool.DictJsonWriter(uid, l)

			//return
			rg := morm.ReadByGroup(uid, g)
			gt.Data["json"] = rg
			gt.ServeJSON()

			fmt.Println(g, n)

			b = true
			break
		}

	}

	if !b {
		gt.Ctx.WriteString("no group can be return")
	}

}
