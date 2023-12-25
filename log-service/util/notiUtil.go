package util

import (
	"bytes"
	"commons/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发送通知相关

func SendTest(*gin.Context) {
	url := config.EnvConf.WEBHOOK.SendUrl
	payload := []byte(fmt.Sprintf(`{"title": "%s", "content": "%s", "importance": %d}`, "Title", "Content", "Importance"))

	// 创建一个请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头部信息
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应结果
	fmt.Println("Response Status:", resp.Status)
}
func SendEmail(Title, Content string, Importance int) {

}
func SendWebHook(Title, Content string, Importance int) {
	url := config.EnvConf.WEBHOOK.SendUrl
	payload := []byte(fmt.Sprintf(`{"title": "%s", "content": "%s", "importance": %d}`, Title, Content, Importance))

	// 创建一个请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头部信息
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应结果
	fmt.Println("Response Status:", resp.Status)
}
