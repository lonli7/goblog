package main

import (
	"github.com/gorilla/mux"
	"github.com/lonli7/goblog/app/http/middlewares"
	"github.com/lonli7/goblog/bootstrap"
	"net/http"
)

var router *mux.Router



func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe("localhost:8080", middlewares.RemoveTrailingSlash(router))
}
