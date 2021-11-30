package mtool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mb/models"
)

//json创建器，n为对应dict的uid,也可以用来更新整个json文件
func DictJsonWriter(n string, z []*models.Srcdstdict) error {

	j, err1 := json.Marshal(z)

	if err1 != nil {
		fmt.Println("err: write json err")
		return err1
	}

	err2 := ioutil.WriteFile("static/json/"+n+".json", j, 0644)
	if err2 != nil {
		panic(err2)
		return err2
	}

	return nil
}

//json读取器,n为对应dict的uid
func DictJsonReader(n string) []*models.Srcdstdict {

	j, err := ioutil.ReadFile("static/json/" + n + ".json")

	if err != nil {
		fmt.Println("read json err")
	}

	var z []*models.Srcdstdict
	json.Unmarshal(j, &z)

	return z
}

//json更新器
func DictJsonUpdater(dict_uid string, l []*models.Srcdstdict) error {

	var fl []*models.Srcdstdict

	sjlist := DictJsonReader(dict_uid)

	//slice转map
	sjlistmap := make(map[string]models.Srcdstdict)

	for i := 0; i < len(sjlist); i++ {
		sjlistmap[sjlist[i].Uid] = *sjlist[i]
	}

	//按key更新
	for i := 0; i < len(l); i++ {
		_, ok := sjlistmap[l[i].Uid]
		if ok {
			sjlistmap[l[i].Uid] = *l[i]
		}
	}

	//map转slice
	for key := range sjlistmap {
		temp := sjlistmap[key]
		fl = append(fl, &temp)
	}

	//fmt.Println(fl)

	err := DictJsonWriter(dict_uid, fl)

	return err

}

//测试，实际运行中不要用这个函数！
func AllJsonCreator(arg interface{}) []byte {
	j, err := json.Marshal(arg)
	if err != nil {
		fmt.Println("AllJsonCreator() err")
	}
	return j
}
