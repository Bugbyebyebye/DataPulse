package config

//auth-service 服务自己的配置
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
