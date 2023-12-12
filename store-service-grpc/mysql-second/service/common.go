package service

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

// TableInfo 数据表属性
type TableInfo struct {
	Field   string `gorm:"column:Field"`
	Type    string `gorm:"column:Type"`
	Null    string `gorm:"column:Null"`
	Key     string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra   string `gorm:"column:Extra"`
}
