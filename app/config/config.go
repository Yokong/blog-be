package config

import (
	"os"

	"github.com/spf13/viper"
)

var c Config

// Config 配置
type Config struct {
	Db
	QiNiu
	ServerAddr string
	Mode       string
}

// Db 数据库配置
type Db struct {
	Addr string
}

// QiNiu 七牛配置
type QiNiu struct {
	Ak     string
	Sk     string
	Bucket string
	Domain string
}

func InitConfig() error {
	name := os.Getenv("CONFIG_NAME")
	tp := os.Getenv("CONFIG_TYPE")
	path := os.Getenv("CONFIG_PATH")
	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType(tp)
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return c
}
