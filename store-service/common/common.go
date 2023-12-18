package common

type CreateReq struct {
	TargetDatabase string  `json:"target_database"`
	TargetTable    string  `json:"target_table"`
	TargetType     int     `json:"target_type"`
	TableList      []Table `json:"table_list"`
}

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

// GetUniqueColumnList 字段列表去重
func GetUniqueColumnList(tableList []Table) []string {
	var columns []string
	uniqueMap := make(map[string]bool)
	for _, t := range tableList {
		for _, v := range t.ColumnList {
			if _, ok := uniqueMap[v]; !ok {
				uniqueMap[v] = true
				columns = append(columns, v)
			}
		}
	}
	return columns
}
