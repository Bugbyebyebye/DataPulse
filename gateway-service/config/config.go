package config

//log-service 服务自己的配置
//数据库配置等

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"runtime/debug"
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

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
				log.Printf("Panic info is: %s", debug.Stack())
			}
		}()

		c.Next()
	}
}
