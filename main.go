package main

import (
	"github.com/lonli7/goblog/app/http/middlewares"
	"github.com/lonli7/goblog/bootstrap"
	"github.com/lonli7/goblog/pkg/logger"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe("localhost:8080", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
