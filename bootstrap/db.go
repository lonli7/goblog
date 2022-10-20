package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	migration(db)
}

func migration(db *gorm.DB) {
	// DB 自动迁移数据
	err := db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
	logger.LogError(err)
}
