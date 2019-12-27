package main

import "github.com/gin-gonic/gin"

func main() {
	//redis 测试
	//test.Redisgo()
	//session 测试
	r := gin.Default()
	
	r.Run(":" + "3000")

}
