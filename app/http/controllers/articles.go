package controllers

import (
	"github.com/gorilla/mux"
	"github.com/lonli7/goblog/app/models/article"
	"github.com/lonli7/goblog/app/policies"
	"github.com/lonli7/goblog/app/requests"
	"github.com/lonli7/goblog/pkg/auth"
	"github.com/lonli7/goblog/pkg/flash"
	"github.com/lonli7/goblog/pkg/logger"
	"github.com/lonli7/goblog/pkg/route"
	"github.com/lonli7/goblog/pkg/view"
	"net/http"
)

type ArticlesController struct {
	BaseController
}

func (a *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 读取对应文章数据
	articles, err := article.Get(id)

	if err != nil {
		a.ResponseForSQLError(w, err)
	} else {
		// 读取成功，显示文章
		view.Render(w, view.D{
			"Article": articles,
			"CanModifyArticle": policies.CanModifyArticle(articles),
		}, "articles.show", "articles._article_meta")
	}
}

func (a *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, pagerData, err := article.GetAll(r, 2)

	if err != nil {
		a.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Articles": articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

func (a *ArticlesController) Create(w http.ResponseWriter, _ *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (a *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	currentUser := auth.User()
	_article := article.Article{
		Title: r.PostFormValue("title"),
		Body: r.PostFormValue("body"),
		UserID: currentUser.ID,
	}

	errors := requests.ValidateArticleForm(_article)

	if len(errors) == 0 {
		err := _article.Create()
		if err != nil {
			logger.LogError(err)
		}

		if _article.ID > 0 {
			indexURL := route.Name2URL("articles.show", "id", _article.GetStringID())
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			a.ResponseForServerError(w, err)
		}
	} else {
		view.Render(w, view.D{
			"Article":  _article,
			"Errors": errors,
		}, "articles.create", "articles._form_field")
	}
}

func (a *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		a.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			a.ResponseForUnauthorized(w, r)
		}
		view.Render(w, view.D{
			"Article": _article,
			"Errors":  view.D{},
		}, "articles.edit", "articles._form_field")
	}
}

// 更新文章
func (a *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		a.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			a.ResponseForUnauthorized(w, r)
		} else {
			// 未出现错误，表单验证
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")

			errors := requests.ValidateArticleForm(_article)

			if len(errors) == 0 {
				// 更新数据
				rowsAffected, err := _article.Update()

				if err != nil {
					a.ResponseForServerError(w, err)
				}

				if rowsAffected > 0 {
					showURL := route.Name2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					flash.Warning("您没有做任何修改!")
				}

			} else {
				view.Render(w, view.D{
					"Article": _article,
					"Errors": errors,
				}, "articles.edit", "articles._form_field")
			}
		}
	}
}

// 删除文章
func (a *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)
	if err != nil {
		a.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			a.ResponseForUnauthorized(w, r)
		} else {
			_, err := _article.Delete()

			if err != nil {
				a.ResponseForServerError(w, err)
			} else {
				indexURL := route.Name2URL("articles.index")
				http.Redirect(w, r, indexURL, http.StatusFound)
			}
		}

	}
}
