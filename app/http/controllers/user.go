package controllers

import (
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type UserController struct {
	BaseController
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_user, err := user.Get(id)

	if err != nil {
		uc.ResponseForSQLError(w, err)
	} else {
		articles, err := article.GetByUserID(_user.GetStringID())
		if err != nil {
			uc.ResponseForServerError(w, err)
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
