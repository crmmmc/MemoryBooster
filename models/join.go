package models

/*
对去orm化对测试
*/

/*
完整的用户数据
包含：
user
user_data
milestone
*/
type Profile struct {
	User      *User
	UserData  []*Userdata
	Milestone []*Milestone
}
