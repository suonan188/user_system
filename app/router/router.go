package router

import (
	v1 "user_system/app/controller/v1"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//InitRouter :
func InitRouter() *gin.Engine {

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	//store := sessions.NewCookieStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("gsessionid", store))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/login", v1.Login)
		apiv1.POST("/register", v1.Register)
		apiv1.GET("/verify", v1.Verify)
		apiv1.GET("/indexpage", v1.Index)
	}

	return r
}
