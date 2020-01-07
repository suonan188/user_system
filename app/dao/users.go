package dao

//User :用于存储用户信息的结构体，Id,Name,Passwd
type User struct {
	ID     string `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:username" json:"username"`
	Passwd string `gorm:"column:password" json:"password"`
}

//TableName :获取表名
func (User) TableName() string {
	return "client_user"
}

//Slice :用于存储用户的切片
var Slice []User

//State :用于临时存储用户登录信息的Map
var State = make(map[string]interface{})
