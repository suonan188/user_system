package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql驱动
	"log"
	"user_system/config"
)

var userdb *gorm.DB

//GetUserDb :
func GetUserDb() *gorm.DB {
	return userdb
}

//InitMsql :mysql初始化
func InitMsql() {
	var err error
	userdb, err = gorm.Open(config.Config.DbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.DbUser,
		config.Config.DbPassword,
		config.Config.DbHost,
		config.Config.DbName))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	userdb.SingularTable(true)

	userdb.DB().SetMaxIdleConns(10)
	userdb.DB().SetMaxOpenConns(100)
	fmt.Printf(config.Config.DbUser, " 连接成功\n\n")

}
