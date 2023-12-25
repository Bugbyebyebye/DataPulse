package handle

import "commons/result"

type TaskHandle struct {
}

var res result.Result

func New() *TaskHandle {
	return &TaskHandle{}
}
