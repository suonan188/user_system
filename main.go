package main

import (
	"os"
	dao "user_system/com/yongzin/dao"
	"user_system/com/yongzin/dao/redis"
	"user_system/com/yongzin/router"
	"user_system/config"
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
