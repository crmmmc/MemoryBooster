package mtool

import (
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
)

var OrmObject orm.Ormer

func init() {

	//默认数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//本地数据库
	//orm.RegisterDataBase("default", "mysql", "root:crm123456@/mbdb?charset=utf8")
	//服务器数据库
	orm.RegisterDataBase("default", "mysql", "mbdb:crm123456@tcp(47.99.163.120:3306)/mbdb?charset=utf8&loc=Local")
}

func GetDbConn() orm.Ormer {
	return OrmObject
}
