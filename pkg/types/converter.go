package types

import (
	"github.com/lonli7/goblog/pkg/logger"
	"strconv"
)

// 将 int64 转换成 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToUint64(num string) uint64 {
	n, err := strconv.ParseUint(num, 10, 64)
	if err != nil {
		logger.LogError(err)
	}

	return n
}

// Uint64ToString 将 uint64 转换为 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}
