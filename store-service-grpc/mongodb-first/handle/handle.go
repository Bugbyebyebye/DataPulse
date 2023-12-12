package handle

import "commons/result"

type StoreHandle struct {
}

func New() *StoreHandle {
	return &StoreHandle{}
}

// res 引入统一返回值
var res result.Result
