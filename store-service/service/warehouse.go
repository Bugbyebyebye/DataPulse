package service

import (
	"log"
	"store-service/common"
	"store-service/config"
	"store-service/dao"
	"store-service/model"
)

// GetWarehousePublicTableInfo 获取数据仓库中公共的数据源
func GetWarehousePublicTableInfo() []common.Table {
	var result []common.Table

	tables, err := model.GetAllPublicTable()
	if err != nil {
		log.Printf("err => %s", err)
	}

	for _, v := range tables {
		//log.Printf("v => %+v", v)
		var table common.Table
		table.SourceName = "mysql"
		table.TableName = v["table_name"].(string)
		table.DatabaseName = v["database_name"].(string)

		//根据数据库名 获取gorm连接指针
		db := config.GetDbByDatabaseName(table.DatabaseName)
		table.ColumnList = dao.GetAllColumnName(table.TableName, db)

		//标出关联标志
		for _, v := range table.ColumnList {
			if v[len(v)-2:] == "id" {
				table.RelateFlag = v
			}
		}

		result = append(result, table)
	}

	return result
}
