package view

import (
	"embed"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

type D map[string]interface{}

var TplFS embed.FS

// Render 渲染通用视图
func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// 渲染简单视图
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// Render 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()
	data["Users"], _ = user.All()
	data["Categories"], _ = category.All()

	allFiles := getTemplateFiles(tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFS(TplFS, allFiles...)
	logger.LogError(err)

	//渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
	// 设置模板相对路径
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 将 articles.show 更正为 articles/show
	layoutFiles, err := fs.Glob(TplFS, viewDir+"layouts/*.gohtml")
	logger.LogError(err)

	return append(layoutFiles, tplFiles...)
}
