package config

//log-service 服务自己的配置
//数据库配置等

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// Cache redis 配置
type Cache interface {
	Put(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

// Mysqldb mysql 配置
const (
	Mysqldb = "root:maojiukeai1412@tcp(222.186.50.126:20134)/df_system?charset=utf8"
)

// Filepath 文件保存路径
const Filepath = "./images"

// QinuConfig 七牛云配置
var QinuConfig = map[string]string{
	"AccessKey": "89VRcOeJLPI3HAuf_lbdT0qU6KrazC7_KdtCZsxT",
	"SecretKey": "UBMz3-03Cw-z-qkIfbMqbvFygbP_72nZCCodlSkg",
	"Bucket":    "hjbsport",
	"Url":       "cdn.emotionalbug.top",
}
