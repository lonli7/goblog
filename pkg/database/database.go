package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"goblog/pkg/logger"
	"time"
)

var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error
	config := mysql.Config{
		User:                    "root",
		Passwd:                  "root",
		Addr:                    "127.0.0.1:3306",
		Net:                     "tcp",
		DBName:                  "goblog",
		AllowCleartextPasswords: true,
		AllowNativePasswords:    true,
	}

	// 准备数据库连接池
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大链接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲链接数
	DB.SetMaxIdleConns(25)
	// 设置每个链接过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id BIGINT(20) AUTO_INCREMENT NOT NULL,
		title VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci,
		PRIMARY KEY (id)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
