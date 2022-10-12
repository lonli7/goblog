package main

import (
	"github.com/lonli7/goblog/app/http/middlewares"
	"github.com/lonli7/goblog/bootstrap"
	"github.com/lonli7/goblog/config"
	c "github.com/lonli7/goblog/pkg/config"
	"github.com/lonli7/goblog/pkg/logger"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
