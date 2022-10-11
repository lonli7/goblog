package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lonli7/goblog/app/models/article"
	"github.com/lonli7/goblog/pkg/logger"
	"github.com/lonli7/goblog/pkg/route"
	"github.com/lonli7/goblog/pkg/types"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

type ArticlesController struct {

}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 读取对应文章数据
	articles, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 读取成功，显示文章
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"Name2URL": route.Name2URL,
				"Uint64ToString": types.Uint64ToString,
			}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, articles)
		logger.LogError(err)
	}
}