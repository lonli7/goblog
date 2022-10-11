package model

import (
	"github.com/lonli7/goblog/pkg/logger"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	DB, err = gorm.Open(config, &gorm.Config{})
	logger.LogError(err)

	return DB
}