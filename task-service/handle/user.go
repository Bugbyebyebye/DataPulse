package handle

import (
	"commons/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"task-service/model"
)

// DeleteApi 实现删除api的操作
func (TaskHandle) DeleteApi(ctx *gin.Context) {
	res := result.Result{}
	var req map[string]interface{}
	// todo 后续替换为当前用户
	UserId := 1
	APID, ok := req["api_id"].(int)
	if ok {
		fmt.Println("断言成功")
	}
	err := model.DeleteAPI(UserId, APID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("删除失败")
		ctx.JSON(200, res.Fail(400, "删除失败失败"))
	}
	ctx.JSON(200, res.Success("删除成功"))
}

// SearchAPIList 实现拉取api列表的功能
func (TaskHandle) SearchAPIList(ctx *gin.Context) {
	res := result.Result{}
	//var req map[string]interface{}

	// 假设当前用户id为1
	UserID := 1 // todo 后续替换为当前用户

	// 调用SearchAPIList函数，并传入UserID
	apiList, err := model.SearchAPIList(UserID)
	if err != nil {
		// 根据需要处理错误
		fmt.Println(err)
		ctx.JSON(200, res.Fail(400, "查询出错"))
		return
	}

	ctx.JSON(200, res.Success(apiList))
}
