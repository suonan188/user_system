package main

import (
	"os"
	dao "user_system/app/dao"
	"user_system/app/dao/redis"
	"user_system/config"
	"user_system/app/router"
)

func main() {
	config.InitConfig()
	dao.InitMsql()
	redis.InitRedis()

	r := router.InitRouter()
	port := os.Getenv("prot")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)

}
