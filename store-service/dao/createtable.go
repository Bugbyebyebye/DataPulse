package dao

import "fmt"

func CreateTableBySQL(tableName string, columns []string) {
	var createTableSQL string
	for i, column := range columns {
		if i == 0 {
			createTableSQL = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s INT AUTO_INCREMENT PRIMARY KEY", tableName, column)
		} else {
			createTableSQL += fmt.Sprintf(", %s VARCHAR(255) NOT NULL", column)
		}
	}

	createTableSQL += ");"

	Warehouse.Exec(createTableSQL)

	fmt.Println("Table created successfully.")
}

func AlertTableBySQL(tableName string, columns []string) {

	var alterTableSQL string
	for i, column := range columns {
		if i == 0 {
			continue // 跳过第一个元素（主键）
		}
		if i == 1 {
			alterTableSQL += fmt.Sprintf(" ADD COLUMN %s VARCHAR(255) NOT NULL", column)
		} else {
			alterTableSQL += fmt.Sprintf(", %s VARCHAR(255) NOT NULL", column)
		}
	}

	alterTableSQL = fmt.Sprintf("ALTER TABLE %s %s;", tableName, alterTableSQL)

	Warehouse.Exec(alterTableSQL)

	fmt.Println("Table altered successfully.")
}
