package handle

import (
	"commons/result"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"task-service/auto/api"
	"time"
)

var res result.Result

// RandomString 生成随机字符串
func RandomString(length int) (string, error) {
	const lettersAndDigits = "abcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, length)
	max := big.NewInt(int64(len(lettersAndDigits)))

	for i := range bytes {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		bytes[i] = lettersAndDigits[n.Int64()]
	}

	return string(bytes), nil
}

func (*TaskHandle) RunDocker(ctx *gin.Context) {
	port := ctx.Query("port")
	name := ctx.Query("name")
	log.Printf("port %s name %s", port, name)
	tunnel := make(chan string)
	namestr, err := RandomString(8)
	if err != nil {
		ctx.JSON(200, res.Fail(400, "生成随机字符串失败"))
	}
	go func() {
		err, _ := api.RunDocker(port, namestr)
		if err != nil {
			tunnel <- "服务生成出错，请检查日志"
		}
		close(tunnel) // 发送完数据后关闭通道
	}()
	select {
	case response := <-tunnel:
		ctx.JSON(200, res.Fail(400, response))
	case <-time.After(5 * time.Second): // 等待5秒
		ctx.JSON(200, res.Success("生成api的链接为:"+namestr+"emotionalbug.top"))
	}
}

func (*TaskHandle) RestartDocker(ctx *gin.Context) {
	name := ctx.Query("port")
	tunnel := make(chan string)

	go func() {
		err := api.RestartDocker(name)
		if err != nil {
			tunnel <- "服务重启出错，请检查日志"
		}
		close(tunnel) // 发送完数据后关闭通道
	}()
	select {
	case response := <-tunnel:
		ctx.JSON(200, res.Fail(400, response))
	case <-time.After(5 * time.Second): // 等待5秒
		ctx.JSON(200, res.Success("服务重启成功"))
	}
}

func (*TaskHandle) StopDocker(ctx *gin.Context) {
	name := ctx.Query("port")
	tunnel := make(chan string)

	go func() {
		err := api.StopDocker(name)
		if err != nil {
			tunnel <- "服务删除出错，请检查日志"
		}
		tunnel <- "服务删除成功"
		close(tunnel) // 发送完数据后关闭通道
	}()
	select {
	case response := <-tunnel:
		ctx.JSON(200, res.Success(response))
	case <-time.After(5 * time.Second): // 等待5秒
		ctx.JSON(200, res.Success("服务Loop.请尽快联系管理员"))
	}
}
