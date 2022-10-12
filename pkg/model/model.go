package model

import (
	"fmt"
	"github.com/lonli7/goblog/pkg/config"
	"github.com/lonli7/goblog/pkg/logger"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	gormConfig := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
			config.GetString("database.mysql.username"),
			config.GetString("database.mysql.password"),
			config.GetString("database.mysql.host"),
			config.GetString("database.mysql.port"),
			config.GetString("database.mysql.database"),
			config.GetString("database.mysql.charset")),
	})

	var level gormlogger.LogLevel
	if config.GetBool("app.debug") {
		level = gormlogger.Warn
	} else {
		level = gormlogger.Error
	}

	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})

	logger.LogError(err)

	return DB
}