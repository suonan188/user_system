package main

import (
	"fmt"
	"time"
	"user_system/config"

	"github.com/garyburd/redigo/redis"
)

func main() {
	config.InitConfig()
	InitRedis()
	// c, err := redis.Dial("tcp", "127.0.0.1:6379")
	// if err != nil {
	// 	fmt.Println("Connect to redis error", err)
	// 	return
	// }
	// defer c.Close()
	fmt.Println("redis 连接成功")
}

//RedisClient :
var RedisClient *redis.Pool

//InitRedis redis初始化
func InitRedis() {
	//建立数据库连接池
	fmt.Println("--建立数据库连接池--")
	RedisClient = &redis.Pool{
		MaxIdle:     config.Config.RedisMaxIdle,     //最初的连接数量
		MaxActive:   config.Config.RedisMaxActive,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: config.Config.RedisIdleTimeout, //连接关闭时间 300秒 （300秒不使用自动关闭）

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) { //要连接的redis数据库

			conn, err := redis.Dial("tcp", config.Config.RedisHost)
			if err != nil {
				return nil, err
			}
			if config.Config.RedisPassword != "" {
				if _, err := conn.Do("AUTH", config.Config.RedisPassword); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
	}
	fmt.Println("redis 连接成功")

}
