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
	"strings"
)

// 转发服务地址
var hosts = map[string]string{
	"auth": "localhost:8081",
	"log":  "localhost:8082",
}

// 允许跳过鉴权
var admits = map[string]string{
	"auth": "localhost:8081",
}

type Info struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
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
		v := Info{}
		err = json.Unmarshal(res.Info, &v)
		if err != nil {
			ctx.JSON(http.StatusOK, r.Fail(400, "用户JSON数据解析错误！"))
			return
		}

		log.Printf("用户 Id:%v Username:%s 登录数据中台", v.Id, v.Username)
		ctx.Set("id", v.Id)
		ctx.Set("username", v.Username)
		//ctx.JSON(200, "成功")
	}

	//请求放行
	ctx.Request.URL.Path = ctx.Param("path") //将路径替换
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
