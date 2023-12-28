package handle

import (
	"commons/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"log-service/model"
	"net/http"
	"strconv"
)

// DeleteOneLogs 通过id删除日志
func (*LogHandler) DeleteOneLogs(ctx *gin.Context) {
	res := &result.Result{}
	var req map[string]interface{}
	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	Table, ok := req["table_name"].(string)
	IDCache, ok := req["log_id"].(float64)
	var ID int
	if ok {
		ID = int(IDCache)
		fmt.Println("断言成功")
	} else {
		// 处理类型断言失败的情况
		fmt.Println("无法转换")
		ctx.JSON(200, res.Fail(400, "断言失败"))
		return
	}
	err := model.DeleteOneLogs(ID, Table)
	if err != nil {
		fmt.Printf("删除日志失败: %s\n", err.Error())
		ctx.JSON(200, res.Fail(4001, "删除日志失败"))
		return
	}
	ctx.JSON(200, res.Success("删除日志成功"))

}

// DeleteUserLogs 删除当前用户的所有日志
func (*LogHandler) DeleteUserLogs(ctx *gin.Context) {
	res := &result.Result{}
	var req map[string]interface{}
	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	get := ctx.Request.Header.Get("id")
	UserId, err := strconv.Atoi(get)
	if err != nil {
		// 转换失败，处理错误
		fmt.Println("转换失败:", err)
	}
	Table, ok := req["table"].(string)
	if ok {
		fmt.Println("断言成功")
	} else {
		fmt.Println("断言失败")
		ctx.JSON(200, res.Fail(400, "断言失败"))
	}
	err = model.DeleteUserLogs(UserId, Table)
	if err != nil {
		fmt.Printf("删除日志失败: %s\n", err.Error())
		ctx.JSON(200, res.Fail(4001, "删除日志失败"))
		return
	}
	ctx.JSON(200, res.Success("删除日志成功"))
}

/*
查询操作
*/

// GetUserLogs 获取用户日志
func (*LogHandler) GetUserLogs(ctx *gin.Context) {
	res := &result.Result{}
	var req map[string]interface{}

	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}

	get := ctx.Request.Header.Get("id")
	UserId, err := strconv.Atoi(get)
	if err != nil {
		// 转换失败，处理错误
		fmt.Println("转换失败:", err)
	}
	fmt.Printf("%s", UserId)
	Table, ok := req["table_name"].(string)
	//Table := "sad"
	if ok {
		fmt.Println("断言成功")
	} else {
		// 处理类型断言失败的情况
		fmt.Println("无法转换为 string 类型")
		ctx.JSON(200, res.Fail(400, "断言失败"))
		return
	}
	logs, err := model.SearchUserLogs(UserId, Table) // 这里用你的用户名调用了之前的函数
	if err != nil {
		fmt.Println("获取不到日志，model异常", err)
		ctx.JSON(200, res.Fail(400, "获取不到日志，请查看日志调用model是否异常"))
		return
	}
	//// 如果 logs 是一个包含单个元素的数组，直接发送单个元素到响应
	//if len(logs) == 1 {
	//	ctx.JSON(200, logs[0])
	//	return
	//}
	ctx.JSON(200, logs)
}
