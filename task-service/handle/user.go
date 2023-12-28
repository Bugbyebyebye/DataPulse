package handle

import (
	"commons/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"task-service/model"
)

// DeleteApi 实现删除api的操作
func (TaskHandle) DeleteApi(ctx *gin.Context) {
	res := &result.Result{}
	var req map[string]interface{}

	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(200, res.Fail(400, "json数据错误！"))
		return
	}
	get := ctx.Request.Header.Get("id")
	//尝试将字符串转换为整数
	UserId, err := strconv.Atoi(get)
	if err != nil {
		// 转换失败，处理错误
		fmt.Println("转换失败:", err)
	}
	//UserId := 1 //测试用
	APICache, ok := req["api_id"].(float64)
	var APID int
	if ok {
		APID = int(APICache)
		fmt.Println("断言成功")
	} else {
		fmt.Println(APID)
		ctx.JSON(200, res.Fail(400, "请输入api_id"))
		return
	}
	err = model.DeleteAPI(UserId, APID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("删除失败")
		ctx.JSON(200, res.Fail(400, "失败失败"))
		return
	}
	ctx.JSON(200, res.Success("删除成功"))
}

func (TaskHandle) SearchStatusLables(ctx *gin.Context) {
	res := &result.Result{}
	var req map[string]interface{}
	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(200, res.Fail(400, "json数据错误！"))
		return
	}
	UserIDCache, ok := req["user_id"].(float64)
	var UserID int
	if ok {
		fmt.Println("断言成功，验证传入了ID")
		UserID = int(UserIDCache)
		stateCounts, err := model.CountStates(UserID)
		if err != nil {
			fmt.Println("查询出错")
			fmt.Println(err)
		}
		ctx.JSON(200, stateCounts)
		return
	} else {
		fmt.Println("断言失败，验证没有传入了ID")
		fmt.Println(UserID)
		get := ctx.Request.Header.Get("id")
		// 尝试将字符串转换为整数
		UserID, err := strconv.Atoi(get)
		if err != nil {
			// 转换失败，处理错误
			fmt.Println("转换失败:", err)
		}
		stateCounts, err := model.CountStates(UserID)
		if err != nil {
			fmt.Println("查询出错")
			fmt.Println(err)
		}
		ctx.JSON(200, res.Success(stateCounts))
		return
	}
}

// SearchAPIList 实现拉取api列表的功能
func (TaskHandle) SearchAPIList(ctx *gin.Context) {
	res := &result.Result{}
	//var req map[string]interface{}

	get := ctx.Request.Header.Get("id")
	// 尝试将字符串转换为整数
	UserID, err := strconv.Atoi(get)
	if err != nil {
		// 转换失败，处理错误
		fmt.Println("转换失败:", err)
	}
	//UserID := 1
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
