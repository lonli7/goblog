package logger

import "log"

// 记录错误日志
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}