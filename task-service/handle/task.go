package handle

import (
	"commons/result"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
	"task-service/auto/api"
	"task-service/model"
	"time"
)

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
	res := &result.Result{}
	var req map[string]interface{}
	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	get := ctx.Request.Header.Get("id")
	// 尝试将字符串转换为整数
	UserId, err := strconv.Atoi(get)
	if err != nil {
		// 转换失败，处理错误
		fmt.Println("转换失败:", err)
	}
	APIName, ok := req["api_name"].(string)
	APIDesc, ok := req["api_desc"].(string)
	if ok {
		fmt.Println("断言成功")
	}
	//port := ctx.Query("port")
	//name := ctx.Query("name")
	//log.Printf("port %s name %s", port, name)
	tunnel := make(chan string)
	namestr, err := RandomString(8)
	if err != nil {
		ctx.JSON(200, res.Fail(400, "生成随机字符串失败"))
	}
	//将字符串拼接为链接
	APIUrl := fmt.Sprintf("%s.emotionalbug.top", namestr)
	go func() {
		err := api.RunDocker(namestr)
		if err != nil {
			tunnel <- "服务生成出错，请检查日志"
		}
		close(tunnel) // 发送完数据后关闭通道
	}()
	select {
	case response := <-tunnel:
		ctx.JSON(200, res.Fail(400, response))
	case <-time.After(5 * time.Second): // 等待5秒
		err := model.InsetAPIList(APIName, APIUrl, APIDesc, UserId)
		if err != nil {
			fmt.Println("插入出错")
		}
		ctx.JSON(200, res.Success(namestr+".emotionalbug.top"))
	}
	//err = logsmodel.InsetAPIList(APIName, APIUrl, APIDesc, UserId)
	//if err != nil {
	//	fmt.Println("插入出错")
	//}

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
