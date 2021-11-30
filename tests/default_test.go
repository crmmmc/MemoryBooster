package test

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
	"mb/models"
	"mb/morm"
	"mb/mtool"
	_ "mb/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	mtool.GetDbConn()
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

//对字典增加测试
func Test1(t *testing.T) {

	for i := 1; i < 100; i++ {
		stri := strconv.Itoa(i)
		testzhen := models.Srcdstdict{
			Src:         "单词l3" + stri,
			Pinyin:      "běn tǔ huò" + stri,
			Dst:         "wordl3" + stri,
			Example_src: "单词是单词l3" + stri,
			Example_dst: "word is a wordl3" + stri,
			State:       0,
			Level:       3,
			Group:       0,
		}
		fmt.Println(testzhen)
		morm.InsertSrcdstdictAndAddGroup(&testzhen)
	}

}

//按level读取并生成group测试
func Test2(t *testing.T) {

	var zp = morm.ReadSrcdstdictByLevel(3)
	var z []models.Srcdstdict

	for i := 0; i < len(zp); i++ {
		var temp = *zp[i]
		z = append(z, temp)
		fmt.Println(temp)
	}

	fmt.Println(z)
	fmt.Println(len(zp))

}

//json生成示范测试
func Test3(t *testing.T) {
	var zp = morm.ReadSrcdstdictByLevel(3)
	var u models.Userdict
	u.Uid = mtool.GetMyUUID()
	mtool.DictJsonWriter(u.Uid, zp)

}

//read json test
func Test4(t *testing.T) {
	n := mtool.DictJsonReader("16343852921241940001789212216")
	for i := 0; i < len(n); i++ {
		fmt.Println(*n[i])
	}
}

//test
func Test5(t *testing.T) {
	stri := " test"
	testzhen := models.Srcdstdict{
		Src:         "单词l3" + stri,
		Dst:         "wordl3" + stri,
		Example_src: "单词是单词l3" + stri,
		Example_dst: "word is a wordl3" + stri,
		State:       0,
		Level:       3,
		Group:       0,
	}
	fmt.Println(testzhen)
	morm.InsertSrcdstdictAndAddGroup(&testzhen)
}

//dictormtest
func Test6(t *testing.T) {
	var v string
	v = "010"
	fmt.Println(morm.ReadSrcdstdict(v))
}

func Test7(t *testing.T) {

	oi := orm.NewOrm()
	var u models.User
	u.Uid = "010"
	oi.Read(&u)

	nanamiProfile := morm.GetProfile(&u)

	j, _ := json.Marshal(nanamiProfile)

	fmt.Printf("%s", j)

}

//重要：用户初始化代码
func Test8(t *testing.T) {
	var u models.User
	uid := morm.InsertUserNewUUID(&u)
	fmt.Println("new user tid is", uid)

	ua := *morm.ReadUser(uid)

	fmt.Println("new user is", ua)
}

func Test9(t *testing.T) {
	u := new(models.User)
	u.Email = "010@qq.com"
	u.Password = "password010"
	u.Nickname = "testman"
	u.State = 1
	morm.InsertUserNewUUID(u)
}

func Test10(t *testing.T) {
	u := new(models.User)
	u.Email = "aaa@qq.com"
	u.Password = "passwordchangeaaaa"
	u.Nickname = "changeaaa"
	u.Country = "china"
	u.State = 5

	u.Uid = "16350807344213070003379621674"

	morm.UpdateUserInfoFromClient(u)

}

func Test11(t *testing.T) {
	u := new(models.User)
	u.Email = "aa"
	u.Password = "p"
	u.Nickname = "changeaaa"
	u.Country = "china"
	u.State = 5

	u.Current_dict = "111"

	u.Uid = "16350807344213070003379621674"

	morm.UpdateCurrentDictFromClient(u)
}

func Test12(t *testing.T) {
	thetimestr := time.Now()
	thetime := time.Now().Unix()
	fmt.Println(thetimestr, "   ", thetime)
}

func Test13(t *testing.T) {
	s := morm.ReadByGroup("16353198837453720002197025905", 3)
	var n int = 0
	for i := 0; i < len(s); i++ {
		fmt.Println(*s[i])
		n++
	}
	fmt.Println(n)
}

func Test14(t *testing.T) {
	uid := "16353198837443390003757245891"
	list := morm.ReadByTid(uid)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func Test15(t *testing.T) {
	email := "2268333096@qq.com"
	str := mtool.SendEmailByAddressFzu(email, "test")
	fmt.Println(str)
}

func Test16(t *testing.T) {
	email := "111@qq.com"
	m := morm.ReadOrCreateSignUserByEmail(email)
	fmt.Println(*m)
}

func Test17(t *testing.T) {
	email := "111@qq.com"
	password := "75761"
	m := morm.UpdateOrCreateSignUserAndReturnVerificationCode(email, password)
	fmt.Println(m)
}

func Test18(t *testing.T) {
	//j:=mtool.DictJsonReader("1635923586579257000919975936")
	//m:=mtool.DictJsonUpdater(j)

}

func Test19(t *testing.T) {
	var u models.User
	u.Nickname = "soy123"
	u.Email = "soy@yy.com"
	u.Password = "soy611"
	uid := morm.InsertUserNewUUID(&u)
	fmt.Println("new user tid is", uid)

	ua := *morm.ReadUser(uid)

	fmt.Println("new user is", ua)
}

func Test20(t *testing.T) {
	user := new(models.User)

	//获取json
	uid := "010"
	user.Uid = uid
	profile := morm.GetProfile(user)
	fmt.Println(*profile.User)

}

func Test21(t *testing.T) {
	email := "ruymi@foxmail.com"
	str := mtool.SendEmailByAddressQQ(email, "test")
	fmt.Println(str)

}
