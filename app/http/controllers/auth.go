package controllers

import (
	"fmt"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"

	"goblog/app/models/user"

	"goblog/pkg/view"
)

type AuthController struct {
}

type userForm struct {
	Name            string `valid:"name"`
	Email           string `valid:"email"`
	Password        string `valid:"password"`
	PasswordConfirm string `valid:"password_confirm"`
}

// 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (a *AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 获取表单数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		_user.Create()

		if _user.ID > 0 {
			flash.Success("恭喜您注册成功!")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建用户失败，请联系管理员")
		}
		// 验证不通过 - 重新显示表单
	}
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

func (a *AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	//
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if err := auth.Attempt(email, password); err == nil {
		flash.Success("欢迎回来!")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}
}

func (a *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	flash.Success("您已退出登录")
	http.Redirect(w, r, "/", http.StatusFound)
}
