package handle

import (
	"auth-service/dao"
	"auth-service/util"
	"commons/result"
	"context"
	"github.com/gin-gonic/gin"
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

type EmailReq struct {
	Email string `json:"email"`
}

// PostEmailCode 发送邮箱验证码接口
func (*AuthHandler) PostEmailCode(ctx *gin.Context) {
	res := &result.Result{}
	req := &EmailReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		log.Printf("email error => %s ", err)
		return
	}

	code, err := util.FormEmail(req.Email)
	if err != nil {
		log.Printf("post email error => %s", err)
		return
	}
	//将验证码存入redis 有效时间3min
	dao.Rc.Put(context.Background(), "DATAPULSE"+req.Email, strconv.Itoa(code), 60*time.Minute)
	ctx.JSON(200, res.Success("验证码发送成功！"))
}
