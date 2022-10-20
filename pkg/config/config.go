package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"goblog/pkg/logger"
)

// Viper viper 实例
var Viper *viper.Viper

type StrMap map[string]interface{}

func init() {
	Viper = viper.New()
	Viper.SetConfigName(".env")
	Viper.SetConfigType("env")
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()
	logger.LogError(err)

	Viper.SetEnvPrefix("appenv")
	Viper.AutomaticEnv()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}

	return Get(envName)
}

// 新增配置项
func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

// 获取配置项
func Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
