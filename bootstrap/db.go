package bootstrap

import (
	"github.com/lonli7/goblog/pkg/model"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
