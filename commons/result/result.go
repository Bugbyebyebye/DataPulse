package result

import (
	"commons/util"
)

//封装了 返回值形式

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Time    string `json:"time"`
}

// Success 成功响应
func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Message = "请求成功！"
	r.Data = data
	r.Time = util.GetResTime()
	return r
}

// Fail 失败响应
func (r *Result) Fail(code int, msg string) *Result {
	r.Code = code
	r.Message = msg
	r.Time = util.GetResTime()
	return r
}
