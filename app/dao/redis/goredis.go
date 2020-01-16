package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"user_system/config"

	"github.com/gomodule/redigo/redis"
)

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

//Set ：兼容老版本redis，集群使用时需要考虑数据脏读的情况
func Set(key string, data interface{}, time int) error {
	conn := RedisClient.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

//GetStringValue :
func GetStringValue(k string) string {
	c := RedisClient.Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}

//SetKeyExpire ：设置到期时间
func SetKeyExpire(k string, ex int) {
	c := RedisClient.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

//CheckKey :
func CheckKey(k string) bool {
	c := RedisClient.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return exist

}

//DelKey :
func DelKey(k string) error {
	c := RedisClient.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//SetJSON 	:
func SetJSON(k string, data interface{}) error {
	c := RedisClient.Get()
	defer c.Close()
	value, _ := json.Marshal(data)
	n, _ := c.Do("SETNX", k, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}
