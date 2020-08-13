package common

import (
	"user_system/app/dao"
)

//IsExist :判断是否存在用户
func IsExist(name string) bool {

	//如果长度为0说明尚未有用户注册
	if len(dao.Slice) == 0 {
		return false
	}

	//遍历切片,dao.Slice;
	for _, v := range dao.Slice {
		return v.Name == name //此时只能和第一个比较，所以第一个之后全为false
		// if v.Name == user {
		// 	return true
		// }
	}

	return false
}

//IsRight :判断密码是否正确
func IsRight(user string, passwd string) bool {
	for _, v := range dao.Slice {
		if v.Name == user {
			//先确认姓名一致，密码相同返回true
			return v.Passwd == passwd
		}
	}
	return false
}
