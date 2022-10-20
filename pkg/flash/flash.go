package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

type Flashes map[string]interface{}

var flashKey = "_flashes"

func init() {
	// 在 gorilla/sessions 中存储 map 和 struct 数据
	// 需要提前注册 gob，方便后续 gob 序列化编码、解码
	gob.Register(Flashes{})
}

// Info 添加Info类型的消息提示
func Info(message string) {
	addFlash("info", message)
}

func Warning(message string) {
	addFlash("warning", message)
}

func Success(message string) {
	addFlash("success", message)
}

func Danger(message string) {
	addFlash("danger", message)
}

func All() Flashes {
	val := session.Get(flashKey)
	// val.(Flashes) 类型检查 val 是否为 Flashes 类型
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	// 读取即销毁
	session.Forget(flashKey)
	return flashMessages
}

func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
