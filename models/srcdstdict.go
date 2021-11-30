package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

//源-译词条
type Srcdstdict struct {
	Uid         string `json:"uid" orm:"pk"`
	Src         string `json:"src"`
	Pinyin      string `json:"pinyin"`
	Dst         string `json:"dst"`
	Example_src string `json:"example___src"`
	Example_dst string `json:"example___dst"`
	State       int    `json:"state"`
	Level       int    `json:"level"`
	Group       int    `json:"group"`
	From        string `json:"from"`
	To          string `json:"to"`
}

//把struct注册到orm
func init() {
	orm.RegisterModel(new(Srcdstdict))
}

/*重写tablename方法
这边为了提高复用性把所有语言都写在同一张表里
*/
func (z *Srcdstdict) TableName() string {
	return "src_dst_dict"
}
