package logsmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Log struct {
	DataType string `json:"data_type"`
	UserId   int    `json:"user_id"`
	Status   string `json:"status"`
	Content  string `json:"content"`
	TaskName string `json:"task_name"`
}

func PostActionLogs(UserId int, Content, Status string) error {

	// 创建 Log 对象
	log := Log{
		DataType: "actionlogs",
		UserId:   UserId,
		Content:  Content,
		Status:   Status,
	}
	// 编码 JSON
	jsonStr, err := json.Marshal(log)
	if err != nil {
		fmt.Println("编码失败:", err)
		return err
	}
	// 发送 POST 请求
	resp, err := http.Post("http://localhost:8082/recorduserlog", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return err
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	return nil
}

func PostTaskLogs(UserId int, TaskName, Status string) error {

	// 创建 Log 对象
	log := Log{
		DataType: "tasklogs",
		UserId:   UserId,
		TaskName: TaskName,
		Status:   Status,
	}
	// 编码 JSON
	jsonStr, err := json.Marshal(log)
	if err != nil {
		fmt.Println("编码失败:", err)
		return err
	}
	// 发送 POST 请求
	resp, err := http.Post("http://localhost:8082/recorduserlog", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return err
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	return nil
}
