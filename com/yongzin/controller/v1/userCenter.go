package v1

import (
	"user_system/com/yongzin/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Login :登录
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{"state": 3, "text": "账号或密码不能为空"})
		return
	}
	state := service.Login(username, password)

	session := sessions.Default(c)

	if session.Get("hello") != "world" {
		session.Set("hello", "world")

		session.Save()

	}
	//session.Delete("gsessionid")
	//session.Clear()
	c.JSON(200, gin.H{"hello": session.Get("hello"), "state": state})

}

// Register : 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(200, gin.H{"state": 3, "text": "账号或密码不能为空"})
		return
	}
	state := service.Register(username, password)
	c.JSON(200, state)
}
