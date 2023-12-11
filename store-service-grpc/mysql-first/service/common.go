package mysql_first_service

type ClientReq struct {
	Message string      `json:"message"`
	Target  string      `json:"target"`
	Param   interface{} `json:"param"`
}

type ServerRes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
