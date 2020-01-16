package service

import (
	"regexp"
	dao "user_system/app/dao"
)

//Login 登录
func Login(name, passwd string) map[string]interface{} {

	//先判断用户是否存在，存在再判断密码是否正确
	user := dao.UserIsExist(name)
	if user != (dao.User{}) {

		if passwd == user.Passwd {
			dao.State["state"] = 1
			dao.State["text"] = "登录成功！"

		} else {
			dao.State["state"] = 0
			dao.State["text"] = "密码错误！"
		}
	} else {
		dao.State["state"] = 2
		dao.State["text"] = "登录失败！此用户尚未注册！"
	}

	return dao.State
	//c.String(http.StatusOK, "%v", State)
}

//Register :用户注册
func Register(name, passwd string) map[string]interface{} {
	//先判断用户名是否存在
	user := dao.UserIsExist(name)
	pwdReg := "^[\\w_-]{6,16}$"
	nameRwd := "^[a-zA-Z0-9_-]{4,16}$"
	BoolPwd, _ := regexp.MatchString(pwdReg, passwd)
	Boolname, _ := regexp.MatchString(nameRwd, name)
	if !Boolname {
		//用户名验证
		dao.State["state"] = 1
		dao.State["text"] = "用户名4到16位"
		return dao.State
	}
	if !BoolPwd {
		//密码验证
		dao.State["state"] = 1
		dao.State["text"] = "密码长度6-16"
		return dao.State
	}
	if user != (dao.User{}) {
		//注册状态
		dao.State["state"] = 1
		dao.State["text"] = "此用户已存在！"
		return dao.State
	}
	user.Name = name
	user.Passwd = passwd
	dao.AddStruct(user)
	dao.State["state"] = 2
	dao.State["text"] = "注册成功!"
	return dao.State

}
