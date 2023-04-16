package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configDir = "config"
)

var v = viper.New()

func init() {
	v.SetConfigName("db")
	v.SetConfigType("json")
	v.AddConfigPath(configDir)

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func Get(key string) interface{} {
	return v.Get(key)
}

func Getstring(key string) string {
	value, ok := v.Get(key).(string)
	if !ok {
		return ""
	}

	return value
}
