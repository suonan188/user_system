package v1

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	comm "user_system/app/common"
	"user_system/app/service"
)

var (
	//Secret :
	Secret = "dong_tech" // 加盐
	//ExpireTime :
	ExpireTime = 3600 // token有效期
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

	claims := &comm.JWTClaims{
		UserID:      1,
		Username:    username,
		Password:    password,
		FullName:    username,
		Permissions: []string{},
	}

	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := comm.GetToken(claims)

	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	//c.String(http.StatusOK, signedToken)
	session := sessions.Default(c)

	if session.Get("hello") != "world" {
		session.Set("hello", "world")

		session.Save()

	}
	//session.Delete("gsessionid")
	//session.Clear()
	c.JSON(200, gin.H{"hello": session.Get("hello"), "state": state, "token": signedToken})

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

//Verify :
func Verify(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token, "-----token------")
	claim, err := comm.VerifyAction(token)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, claim.Username)
}
