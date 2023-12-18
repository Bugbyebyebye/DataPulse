package common

type Table struct {
	RelateFlag   string   `json:"relate_flag"`
	SourceName   string   `json:"source_name"`
	DatabaseName string   `json:"database_name"`
	TableName    string   `json:"table_name"`
	ColumnList   []string `json:"column_list"`
}
