package config

import "time"

//Conf :全局配置结构体
type Conf struct {
	DbType     string `yaml:"DB_TYPE"`
	DbUser     string `yaml:"DB_USER"`
	DbPassword string `yaml:"DB_PASSWORD"`
	DbHost     string `yaml:"DB_HOST"`
	DbName     string `yaml:"DB_NAME"`
	//redis
	RedisMaxIdle     int           `yaml:"RedisMaxIdle"`
	RedisMaxActive   int           `yaml:"RedisMaxActive"`
	RedisIdleTimeout time.Duration `yaml:"RedisIdleTimeout"`
	RedisHost        string        `yaml:"RedisHost"`
	RedisPassword    string        `yaml:"RedisPassword"`
}

// Config 全局配置
var (
	ENV    string
	Config *Conf
)
