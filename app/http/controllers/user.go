package controllers

import (
	"fmt"
	"github.com/lonli7/goblog/app/models/article"
	"github.com/lonli7/goblog/app/models/user"
	"github.com/lonli7/goblog/pkg/logger"
	"github.com/lonli7/goblog/pkg/route"
	"github.com/lonli7/goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {

}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_user, err := user.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 用户未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		}
	} else {
		articles, err := article.GetByUserID(_user.GetStringID())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
