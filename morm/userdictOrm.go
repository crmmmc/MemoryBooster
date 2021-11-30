package morm

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"mb/models"
	"mb/mtool"
	"strconv"
)

/*新建user自动生成不同level的dict
规则：自动搜索level=1，2，3的词条生成三个dict
还没写完
*/
func InsertUserdict(l int, u *models.User) string {

	//初始化
	var userdict models.Userdict

	strl := strconv.Itoa(l)
	userdict.Tid = u.Uid //绑定
	userdict.Level = l
	userdict.Name = "level " + strl + " dict"

	//生成json
	var zp = ReadSrcdstdictByLevel(l)
	userdict.Uid = mtool.GetMyUUID()
	mtool.DictJsonWriter(userdict.Uid, zp)
	userdict.Dict_json = "static/json/" + userdict.Uid + ".json" //写入地址

	var oi = orm.NewOrm()
	id, err := oi.Insert(&userdict)
	if err == nil {
		fmt.Println("user insert success: id= ", id)
	} else {
		fmt.Println(err)
	}

	return userdict.Uid

}

//自动生成器1,2,3
func AutoInsertUserdict(u *models.User) string {

	var current string

	for i := 1; i <= 3; i++ {
		n := InsertUserdict(i, u)
		if i == 1 {
			current = n
		}
	}

	return current
}

//按group和uid读取组，返回组
func ReadByGroup(uid string, group int) []*models.Srcdstdict {

	l := mtool.DictJsonReader(uid)
	var g []*models.Srcdstdict

	for i := 0; i < len(l); i++ {
		if l[i].Group == group {
			g = append(g, l[i])
		}
	}

	return g

}

//输入tid返回所有dict
func ReadByTid(tid string) []*models.Userdict {
	var dictList []*models.Userdict
	oi := orm.NewOrm()
	oi.QueryTable("user_dict").Filter("tid", tid).All(&dictList)
	return dictList
}

//建立连接
func init() {
	mtool.GetDbConn()
}
