package session

import (
	"github.com/gorilla/sessions"
	"github.com/lonli7/goblog/pkg/logger"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("33446a9dcf9ea060a0a6532b166da32f304af0de"))

var Session *sessions.Session

var Request *http.Request

var Response http.ResponseWriter

// 初始化会话
func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error

	Session, err = Store.Get(r, "goblog-session")
	logger.LogError(err)

	Request = r
	Response = w
}

// Put 写入键值对应的会话数据
func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

// Get 获取会话数据，获取数据时做类型检测
func Get(key string) interface{} {
	return Session.Values[key]
}

// Forget 删除某个会话
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// Flush 删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

// Save 保持会话
func Save() {
	err := Session.Save(Request, Response)
	logger.LogError(err)
}
