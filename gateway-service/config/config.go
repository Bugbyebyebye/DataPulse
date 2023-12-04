package config

//log-service 服务自己的配置
//数据库配置等

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
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
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,token")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
				log.Printf("Panic info is: %s", debug.Stack())
				c.AbortWithError(http.StatusInternalServerError, errors.New("Internal server error"))
			}
		}()
		c.Next()
	}
}
