package handle

import (
	"auth-service/dao"
	"auth-service/util"
	"commons/result"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

// UploadPic 上传图片 返回图片链接到前端
func (*AuthHandler) UploadPic(ctx *gin.Context) {
	r := &result.Result{}

	file, err := ctx.FormFile("avatar")
	if err != nil {
		log.Printf("文件上传失败！err => %s", err)
		ctx.JSON(http.StatusOK, r.Fail(4001, "图片上传失败！"))
		return
	}
	//将图片存在本地
	//filepath := path.Join(config.Filepath, file.Filename)
	//err = ctx.SaveUploadedFile(file, filepath)
	//if err != nil {
	//	log.Printf("文件 【%s】 保存失败！ err => %s", file.Filename, err)
	//	ctx.JSON(http.StatusOK, r.Fail(4001, "图片保存失败！"))
	//	return
	//}
	//调用上传图片方法
	open, _ := file.Open()
	url, err := util.UploadToQiNiu(open, file.Size)
	if err != nil {
		log.Printf("图片上传七牛云失败！ err => %s", err)
		ctx.JSON(http.StatusOK, r.Fail(4001, "图片上传七牛云失败！"))
		return
	}

	ctx.JSON(http.StatusOK, r.Success(url))
}

// PostEmailCode 发送邮箱验证码接口
func (*AuthHandler) PostEmailCode(ctx *gin.Context) {
	res := &result.Result{}

	email := ctx.Query("email")
	log.Printf("email => %s", email)
	if email == "" {
		ctx.JSON(200, res.Fail(4001, "邮箱不能为空！"))
		return
	}

	code, err := util.FormEmail(email)
	if err != nil {
		log.Printf("post email error => %s", err)
		return
	}
	//将验证码存入redis 有效时间3min
	dao.Rc.Put(context.Background(), "DATAPULSE"+email, strconv.Itoa(code), 3*time.Minute)
	ctx.JSON(200, res.Success("验证码发送成功！"))
}

// VerifyToken 校验token
func (*AuthHandler) VerifyToken(ctx *gin.Context) {
	var token string
	err := ctx.BindJSON(&token)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("token => %s", token)

	claims, err := util.ParseToken(token)
	if err != nil {
		//log.Printf("err => %s", err)
		if errors.Is(err, jwt.ErrInvalidKey) {
			ctx.JSON(http.StatusOK, res.Fail(4001, "token 无效！"))
			return
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			ctx.JSON(http.StatusOK, res.Fail(4002, "token 过期！"))
			return
		}
	}

	ctx.JSON(http.StatusOK, res.Success(claims))
}
