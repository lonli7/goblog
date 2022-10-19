package middlewares

import (
	"github.com/lonli7/goblog/pkg/auth"
	"net/http"
)

// Auth 登录用户才能访问
func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.Check() {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}
