package view

import (
	"github.com/lonli7/goblog/pkg/logger"
	"github.com/lonli7/goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

// Render 渲染视图
func Render(w io.Writer, name string, data interface{}) {
	// 设置模板相对路径
	viewDir := "resources/views/"

	// 将 articles.show 更正为 articles/show
	name = strings.Replace(name, ".", "/", -1)

	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	newFiles := append(files, viewDir+name+".gohtml")

	tmpl, err := template.New(name + ".gohtml").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(newFiles...)
	logger.LogError(err)

	//渲染模板
	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
