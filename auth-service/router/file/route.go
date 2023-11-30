package filerouter

import (
	"auth-service/handle"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {
}

func (*FileRouter) Route(r *gin.Engine) {
	file := handle.New()
	r.POST("/file", file.UploadPic)
}
