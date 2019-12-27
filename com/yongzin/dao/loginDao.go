package dao

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

//UserIsExist :
func UserIsExist(name string) (users User) {
	GetUserDb().Find(&users, "username = ?", name)
	return users
}

//UserIsRight :
func UserIsRight(password, pwd string) bool {
	return password == pwd
}

//AddStruct :添加用户
func AddStruct(users User) {
	fmt.Println(users)
	u1 := uuid.NewV4()
	s := strings.Replace(u1.String(), "-", "", -1)
	users.ID = s
	f := GetUserDb().Save(users).Error
	fmt.Println(f)

}
