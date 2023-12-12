package handle

import "commons/result"

type StoreHandle struct {
}

func New() *StoreHandle {
	return &StoreHandle{}
}

// res 引入统一返回值
var res result.Result

// DataSource 数据源信息
type DataSource struct {
	SaveName     string      `json:"save_name,omitempty"`
	FromName     string      `json:"source_name"`
	DatabaseList []Database  `json:"database_list,omitempty"`
	Databases    interface{} `json:"databases,omitempty"`
}

type Database struct {
	DatabaseName string  `json:"database_name"`
	TableList    []Table `json:"table_list"`
}

type Table struct {
	TableName  string   `json:"table_name"`
	ColumnList []string `json:"column_list"`
}
