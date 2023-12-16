package handle

import (
	auth "commons/api/auth/gen"
	"commons/result"
	"encoding/json"
	"errors"
	"gateway-service/client"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

type Info struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Authority int    `json:"authority"`
}

func ProxyHandler(ctx *gin.Context) {
	r := &result.Result{}
	var info Info
	//根据路径匹配路由
	split := strings.Split(ctx.Request.URL.String(), "/")
	proxyUrl, _ := url.Parse("http://" + hosts[split[1]])

	if _, ok := admits[split[1]]; !ok {
		//鉴权
		token := ctx.Request.Header.Get("token")
		//log.Printf("token => %s", token)

		res, err := client.AuthClient.VerifyToken(ctx, &auth.Req{Token: token})
		if err != nil {
			//鉴权失败
			log.Printf("err => %s", err)
			if errors.Is(err, jwt.ErrInvalidKey) {
				ctx.JSON(http.StatusOK, r.Fail(400, "token 无效！"))
				return
			} else if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.JSON(http.StatusOK, r.Fail(400, "token 过期！"))
				return
			}
		}
		//鉴权成功
		//log.Printf("res => %s", res)
		err = json.Unmarshal(res.Info, &info)
		if err != nil {
			ctx.JSON(http.StatusOK, r.Fail(400, "用户JSON数据解析错误！"))
			return
		}

		log.Printf("用户 Id:%v Username:%s 访问数据中台", info.Id, info.Username)
		ctx.Request.Header.Set("id", strconv.Itoa(info.Id))
		ctx.Request.Header.Set("username", info.Username)
		ctx.Request.Header.Set("role", info.Role)
		ctx.Request.Header.Set("authority", strconv.Itoa(info.Authority))
		//ctx.JSON(200, "成功")
	}

	//请求放行
	ctx.Request.URL.Path = ctx.Param("path") //将路径替换
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
