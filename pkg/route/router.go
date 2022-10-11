package route

import (
	"github.com/gorilla/mux"
	"github.com/lonli7/goblog/pkg/logger"
	"net/http"
)

var Router *mux.Router


// 初始化路由
func Initialize() {
	Router = mux.NewRouter()
}

// 通过路由名称来获取URL
func RouteName2URL(routerName string, pairs ...string) string {
	URL, err := Router.Get(routerName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return URL.String()
}

// 获取URL参数
func GetRouterVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}