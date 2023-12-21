package handle

import (
	"commons/result"
	"gateway-service/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

// 转发服务地址
var hosts = map[string]string{
	"auth":  "localhost:8081",
	"log":   "localhost:8082",
	"store": "localhost:8083",
	"task":  "localhost:8084",
}

// 允许跳过鉴权
var admits = map[string]string{
	"auth": "localhost:8081",
}

func ProxyHandler(ctx *gin.Context) {
	r := &result.Result{}
	//根据路径匹配路由
	split := strings.Split(ctx.Request.URL.String(), "/")
	proxyUrl, _ := url.Parse("http://" + hosts[split[1]])

	if _, ok := admits[split[1]]; !ok {
		//鉴权
		token := ctx.Request.Header.Get("token")
		//log.Printf("token => %s", token)

		res, err := client.VerifyToken(ctx, token)
		if err != nil {
			//鉴权失败
			log.Printf("err => %s", err)
		}
		//鉴权成功
		log.Printf("res => %+v", res)
		code := res.Code
		if code == 4001 {
			log.Printf("token 无效")
			ctx.JSON(http.StatusOK, r.Fail(4001, "token 无效"))
			return
		} else if code == 4002 {
			log.Printf("token 过期")
			ctx.JSON(http.StatusOK, r.Fail(4002, "token 过期"))
			return
		}

		data := res.Data.(map[string]interface{})
		log.Printf("用户 Id:%v Username:%s 访问数据中台", data["id"], data["username"])
		ctx.Request.Header.Set("id", strconv.Itoa(int(data["id"].(float64))))
		ctx.Request.Header.Set("username", data["username"].(string))
		ctx.Request.Header.Set("role", data["role"].(string))
		ctx.Request.Header.Set("authority", strconv.Itoa(int(data["authority"].(float64))))
	}

	//请求放行
	ctx.Request.URL.Path = ctx.Param("path") //将路径替换
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
