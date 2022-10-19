package controllers

import (
	"github.com/lonli7/goblog/app/models/article"
	"github.com/lonli7/goblog/app/models/category"
	"github.com/lonli7/goblog/app/requests"
	"github.com/lonli7/goblog/pkg/flash"
	"github.com/lonli7/goblog/pkg/route"
	"github.com/lonli7/goblog/pkg/view"
	"net/http"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Create(w http.ResponseWriter, _ *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

func (c *CategoryController) Store(w http.ResponseWriter, r *http.Request) {
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	errors := requests.ValidateCategoryForm(_category)

	if len(errors) == 0 {
		err := _category.Create()
		if _category.ID > 0 {
			flash.Success("分类创建成功")
			indexURL := route.Name2URL("articles.index")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			c.ResponseForServerError(w, err)
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors": errors,
		}, "categories.create")
	}
}

func (c *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_category, err := category.Get(id)
	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)
	if err != nil {
		c.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Articles": articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}
