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
	UserID int `json:"user_id"`
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
		//get := ctx.Request.Header.Get("id")
		//UserId, err := strconv.Atoi(get)
		//if err != nil {
		//	// 转换失败，处理错误
		//	fmt.Println("转换失败:", err)
		//}
		UserIdCache, ok := req["user_id"].(float64)
		var UserId int
		Content, ok := req["content"].(string)
		//json解析的数据是float64,需要二次转换
		Status, ok := req["status"].(string)
		if ok {
			UserId = int(UserIdCache)
			fmt.Println("断言成功")
		} else {
			// 处理类型断言失败的情况
			fmt.Println("无法转换为 string 类型")
			ctx.JSON(200, res.Fail(400, "断言失败"))
			return
		}
		err := model.RecordUserActionLog(UserId, Content, Status)
		if err != nil {
			fmt.Printf("将日志数据插入到数据库失败: %s\n", err.Error())
			ctx.JSON(200, res.Fail(4001, "将日志数据插入到数据库失败"))
			return
		}
		ctx.JSON(200, res.Success("日志记录成功"))
	case "tasklogs":
		//get := ctx.Request.Header.Get("id")
		//UserId, err := strconv.Atoi(get)
		//if err != nil {
		//	// 转换失败，处理错误
		//	fmt.Println("转换失败:", err)
		//}
		UserIdCache, ok := req["user_id"].(float64)
		var UserId int
		TaskName, ok := req["task_name"].(string)
		Status, ok := req["status"].(string)
		Takingtime, ok := req["taking_time"].(string)
		if ok {
			UserId = int(UserIdCache)
			fmt.Println("断言成功")
		} else {
			// 处理类型断言失败的情况
			fmt.Println("无法转换为 string 类型")
			ctx.JSON(200, res.Fail(400, "断言失败"))
			return
		}
		err := model.RecordUserTackLog(UserId, TaskName, Status, Takingtime)
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
		//get := ctx.Request.Header.Get("id")
		//UserId, err := strconv.Atoi(get)
		//if err != nil {
		//	// 转换失败，处理错误
		//	fmt.Println("转换失败:", err)
		//}
		UserIdCache, ok := req["user_id"].(float64)
		var UserId int
		if ok {
			UserId = int(UserIdCache)
			fmt.Println("断言成功")
		} else {
			// 处理类型断言失败的情况
			fmt.Println("无法转换为 string 类型")
			ctx.JSON(200, res.Fail(400, "断言失败"))
			return
		}
		err := model.RecordUserInvokeLog(ApiUrl, UserId, Invokeip, Invokelog, Status)
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
