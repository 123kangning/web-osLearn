package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var dsn = "kangning:9264@tcp(127.0.0.1:3306)/osLearn?charset=utf8mb4&parseTime=True&loc=Local"

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,                                //缓存预编译命令
		SkipDefaultTransaction: true,                                //禁用默认事务操作
		Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句))
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Project{})
	if err != nil {
		panic(err)
	}
}
