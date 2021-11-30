package morm

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"mb/models"
	"mb/mtool"
)

//新增
func InsertSrcdstdictAndAddGroup(z *models.Srcdstdict) {

	var zhen models.Srcdstdict

	zhen.Uid = mtool.GetMyUUID()
	zhen.Src = z.Src
	zhen.Pinyin = z.Pinyin
	zhen.Dst = z.Dst
	zhen.Example_src = z.Example_src
	zhen.Example_dst = z.Example_dst
	zhen.State = z.State
	zhen.Level = z.Level
	zhen.Group = z.Group
	zhen.From = "zh"
	zhen.To = "en"

	var oi = orm.NewOrm()
	id, err := oi.Insert(&zhen)
	if err == nil {
		fmt.Println("zhendict insert success: id= ", id)
	} else {
		fmt.Println(err)
	}
}

//按uid读取
func ReadSrcdstdict(str string) *models.Srcdstdict {
	var oi = orm.NewOrm()
	var s *models.Srcdstdict
	err := oi.Raw("select * from src_dst_dict where uid = ?", str).QueryRow(&s)
	if err != nil {
		fmt.Println("ReadSrcdstdict() err")
	}
	return s
}

/*select ... where level = l
生成level=l的dict，并且生成group
*/
func ReadSrcdstdictByLevel(l int) []*models.Srcdstdict {

	//取[]
	var zhendictList []*models.Srcdstdict
	var oi = orm.NewOrm()
	oi.QueryTable("src_dst_dict").Filter("level", l).All(&zhendictList)

	//生成group
	for i := 0; i < len(zhendictList); i++ {
		zhendictList[i].Group = i/10 + 1
	}
	return zhendictList

}

//建立连接
func init() {
	mtool.GetDbConn()
}
