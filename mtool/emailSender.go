package mtool

import "github.com/beego/beego/v2/adapter/utils"

func SendEmailByAddress(email string, vc string) string {

	con := `{"username":"ruiming_chen@yeah.net","password":"XMAGJSFZATOPANJQ","Host":"smtp.yeah.net","Port":25}`
	ema := utils.NewEMail(con)
	ema.To = []string{email, "ruiming_chen@yeah.net"}
	ema.From = "ruiming_chen@yeah.net"
	ema.Subject = "AD: MemoryBooster 验证邮件"
	ema.Text = "您的验证码为: " + vc
	//ema.HTML = "<h1>test</h1>"
	//ema.AttachFile("1.jpg")
	//ema.AttachFile("1.jpg","1")
	err := ema.Send()
	if err != nil {
		return "send email err:\n" + err.Error()
	} else {
		return "send email success"
	}

}

func SendEmailByAddressYandex(email string, vc string) string {

	con := `{"username":"MemoryBooster@yandex.com","password":"ksoioxdblfltaimz","Host":"smtp.yandex.com","Port":465}`
	ema := utils.NewEMail(con)
	ema.To = []string{email, "MemoryBooster@yandex.com"}
	ema.From = "MemoryBooster@yandex.com"
	ema.Subject = "MemoryBooster 验证邮件"
	ema.Text = "您的验证码为: " + vc
	//ema.HTML = "<h1>test</h1>"
	//ema.AttachFile("1.jpg")
	//ema.AttachFile("1.jpg","1")
	err := ema.Send()
	if err != nil {
		return "send email err:\n" + err.Error()
	} else {
		return "send email success"
	}

}

func SendEmailByAddressFzu(email string, vc string) string {

	con := `{"username":"831904215@fzu.edu.cn","password":"FGFQf4arIbwSUN4k","Host":"smtp.fzu.edu.cn","Port":25}`
	ema := utils.NewEMail(con)
	ema.To = []string{email, "831904215@fzu.edu.cn"}
	ema.From = "831904215@fzu.edu.cn"
	ema.Subject = "MemoryBooster 验证邮件"
	ema.Text = "您的验证码为: " + vc
	//ema.HTML = "<h1>test</h1>"
	//ema.AttachFile("1.jpg")
	//ema.AttachFile("1.jpg","1")
	err := ema.Send()
	if err != nil {
		return "send email err:\n" + err.Error()
	} else {
		return "send email success"
	}

}

func SendEmailByAddress1110(email string, vc string) string {

	con := `{"username":"831904215@fzu.edu.cn","password":"FGFQf4arIbwSUN4k","Host":"smtp.fzu.edu.cn","Port":25}`
	ema := utils.NewEMail(con)
	ema.To = []string{email, "831904215@fzu.edu.cn"}
	ema.From = "831904215@fzu.edu.cn"
	ema.Subject = "MemoryBooster 验证邮件"
	ema.Text = "您的验证码为: " + vc
	//ema.HTML = "<h1>test</h1>"
	//ema.AttachFile("1.jpg")
	//ema.AttachFile("1.jpg","1")
	err := ema.Send()
	if err != nil {
		return "send email err:\n" + err.Error()
	} else {
		return "send email success"
	}

}

func SendEmailByAddressQQ(email string, vc string) string {

	con := `{"username":"778706440@qq.com","password":"qozyzxqndldpbccf","Host":"smtp.qq.com","Port":587}`
	ema := utils.NewEMail(con)
	ema.To = []string{email, "778706440@qq.com"}
	ema.From = "778706440@qq.com"
	ema.Subject = "MemoryBooster 验证邮件"
	ema.Text = "您的验证码为: " + vc
	//ema.HTML = "<h1>test</h1>"
	//ema.AttachFile("1.jpg")
	//ema.AttachFile("1.jpg","1")
	err := ema.Send()
	if err != nil {
		return "send email err:\n" + err.Error()
	} else {
		return "send email success"
	}

}
