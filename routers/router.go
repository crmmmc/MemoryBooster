package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mb/controllers"
)

func init() {

	/*
		固定路由
	*/

	beego.Router("/test", &controllers.TestController{}) //测试

	beego.Router("/user/profile", &controllers.ProfileController{}) //个人信息
	beego.Router("/user/info", &controllers.UserController{})
	beego.Router("/dict/userdict", &controllers.DictController{})
	beego.Router("/dict/userdict/group", &controllers.GroupController{})
	beego.Router("/dict/userdict/group/tool", &controllers.GroupToolController{})
	beego.Router("/user/current_dict", &controllers.UserdictController{})
	beego.Router("/user/milestone", &controllers.MilestoneController{}) //里程碑
	beego.Router("/sign/info", &controllers.SignController{})           //注册
	beego.Router("/log", &controllers.LogController{})                  //登录

	//正则路由

}
