package controllers

import (
	"fmt"
	"github.com/lonli7/goblog/pkg/flash"
	"github.com/lonli7/goblog/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

// BaseController 基础控制器
type BaseController struct {

}

// ResponseForSQLError 处理SQL错误并返回
func (bc BaseController) ResponseForSQLError(w http.ResponseWriter, err error) {
	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprint(w, "404 数据未找到")
		logger.LogError(err)
	} else {
		bc.ResponseForServerError(w, err)
	}
}

func (bc BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("未授权操作!")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (bc BaseController) ResponseForServerError(w http.ResponseWriter, err error) {
	logger.LogError(err)
	w.WriteHeader(http.StatusInternalServerError)
	_, err = fmt.Fprint(w, "500 服务器内部错误")
	logger.LogError(err)
}
