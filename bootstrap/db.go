package bootstrap

import (
	"github.com/lonli7/goblog/pkg/config"
	"github.com/lonli7/goblog/pkg/model"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
}
