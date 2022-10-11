package route

import (
	"github.com/gorilla/mux"
	"github.com/lonli7/goblog/pkg/logger"
)

var route *mux.Router


func SetRoute(r *mux.Router) {
	route = r
}


func Name2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return url.String()
}