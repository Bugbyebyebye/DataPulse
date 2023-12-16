package common

type Database struct {
	DatabaseName string  `json:"database_name"`
	TableList    []Table `json:"table_list"`
}

type Table struct {
	TableName  string   `json:"table_name"`
	ColumnList []string `json:"column_list"`
}
