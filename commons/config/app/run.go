package app

// 封装了启停服务方法

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run 用于启停服务 r gin的服务上下文 serverName 服务名称
func Run(r *gin.Engine, serverName string, addr string, stop func()) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		log.Printf("%s running in %s \n", serverName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down project %s app... \n", serverName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//grpc 的停止服务
	if stop != nil {
		stop()
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s app Shudown, case by: %s \n", serverName, err)
	}
	select {
	case <-ctx.Done():
		log.Printf("closed timeout ... \n")
	}
	log.Printf("%s app stop success ...", serverName)
}
