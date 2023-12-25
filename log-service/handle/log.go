package handle

import (
	"commons/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"log-service/model"
	"net/http"
)

/*
日志服务创建结构体
处理有关log的全部代码
*/

type UserName struct {
	UserName string `json:"user_name"`
}

// LogHandler /*
type LogHandler struct {
}

func New() *LogHandler {
	return &LogHandler{}
}

func (*LogHandler) GetLogInfo(ctx *gin.Context) {
	r := &result.Result{}

	ctx.JSON(http.StatusOK, r.Success("查看日志成功！"))
}

/*
插入操作
*/

// Logging 接受用户操作/*
func (*LogHandler) Logging(ctx *gin.Context) {
	res := &result.Result{}
	var req map[string]interface{}

	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	switch req["data_type"] {
	case "actionlogs":
		userName, ok := req["user_name"].(string)
		Content, ok := req["content"].(string)
		//json解析的数据是float64,需要二次转换
		Status, ok := req["status"].(string)
		if ok {
			fmt.Println("断言成功")
		} else {
			// 处理类型断言失败的情况
			fmt.Println("无法转换为 string 类型")
			ctx.JSON(200, res.Fail(400, "断言失败"))
			return
		}
		err := model.RecordUserActionLog(userName, Content, Status)
		if err != nil {
			fmt.Printf("将日志数据插入到数据库失败: %s\n", err.Error())
			ctx.JSON(200, res.Fail(4001, "将日志数据插入到数据库失败"))
			return
		}
		ctx.JSON(200, res.Success("日志记录成功"))
	case "tasklogs":
		userName, ok := req["user_name"].(string)
		TaskName, ok := req["task_name"].(string)
		Status, ok := req["status"].(string)
		Takingtime, ok := req["taking_time"].(string)
		if ok {
			fmt.Println("断言成功")
		} else {
			// 处理类型断言失败的情况
			fmt.Println("无法转换为 string 类型")
			ctx.JSON(200, res.Fail(400, "断言失败"))
			return
		}
		err := model.RecordUserTackLog(userName, TaskName, Status, Takingtime)
		if err != nil {
			fmt.Printf("将日志数据插入到数据库失败: %s\n", err.Error())
			ctx.JSON(200, res.Fail(4001, "将日志数据插入到数据库失败"))
			return
		}
		ctx.JSON(200, res.Success("日志记录成功"))
	case "invokelogs":
		//userName, ok := req["user_name"].(string)
		ApiUrl, ok := req["api_url"].(string)
		Invokeip, ok := req["invoke_ip"].(string)
		Invokelog, ok := req["invoke_log"].(string)
		//json解析的数据是float64,需要二次转换
		Status, ok := req["status"].(string)
		User, ok := req["user_name"].(string)
		if ok {
			fmt.Println("断言成功")
		} else {
			// 处理类型断言失败的情况
			fmt.Println("无法转换为 string 类型")
			ctx.JSON(200, res.Fail(400, "断言失败"))
			return
		}
		err := model.RecordUserInvokeLog(ApiUrl, User, Invokeip, Invokelog, Status)
		if err != nil {
			fmt.Printf("将日志数据插入到数据库失败: %s\n", err.Error())
			ctx.JSON(200, res.Fail(4001, "将日志数据插入到数据库失败"))
			return
		}
		ctx.JSON(200, res.Success("日志记录成功"))
	default:
		// 处理未知的 data_type 值
		ctx.JSON(200, res.Fail(400, "未知的 data_type"))
	}
}

/*
删除操作
*/

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
	UserId, ok := req["table_name"].(int)
	Table, ok := req["table"].(string)
	if ok {
		fmt.Println("断言成功")
	} else {
		fmt.Println("断言失败")
		ctx.JSON(200, res.Fail(400, "断言失败"))
	}
	err := model.DeleteUserLogs(UserId, Table)
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

	UserId, ok := req["user_id"].(int)
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
