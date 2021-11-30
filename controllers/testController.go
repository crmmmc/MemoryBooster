package controllers

type TestController struct {
	baseController
}

//测试token
func (t *TestController) Get() {
	str := t.GetString("str")
	str = "The message is " + str
	t.Ctx.WriteString(str)
}
