package common

type Table struct {
	RelateFlag   string   `json:"relate_flag"`
	SourceName   string   `json:"source_name"`
	DatabaseName string   `json:"database_name"`
	TableName    string   `json:"table_name"`
	ColumnList   []string `json:"column_list"`
}

// TableInfo 数据表属性
type TableInfo struct {
	Field   string `gorm:"column:Field"`
	Type    string `gorm:"column:Type"`
	Null    string `gorm:"column:Null"`
	Key     string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra   string `gorm:"column:Extra"`
}
