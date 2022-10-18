package view

import (
	"github.com/lonli7/goblog/pkg/logger"
	"github.com/lonli7/goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

// Render 渲染通用视图
func Render(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// 渲染简单视图
func RenderSimple(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// Render 渲染视图
func RenderTemplate(w io.Writer, name string, data interface{}, tplFiles ...string) {
	// 设置模板相对路径
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 将 articles.show 更正为 articles/show
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	allFiles := append(layoutFiles, tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	//渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}
