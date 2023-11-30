package handle

import (
	"auth-service/util"
	"commons/result"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
