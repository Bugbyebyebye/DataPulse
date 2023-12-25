package dao

import (
	"log"
	"store-service/config"
	"store-service/model"
)

func DeleteWarehouseTable(id int, tableId int) error {

	err := model.DeleteUserTableRelate(id, tableId)
	if err != nil {
		log.Printf("err => %s", err)
	}

	err = model.DeleteTable(tableId)
	if err != nil {
		log.Printf("err => %s", err)
	}

	table, err := model.GetTableInfo(tableId)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("table => %+v", table)
	log.Printf("name => %+v", table["database_name"])
	log.Printf("name => %+v", table["table_name"])

	db := config.GetDbByDatabaseName(table["database_name"].(string))
	err = db.Exec("DROP TABLE IF EXISTS " + table["table_name"].(string)).Error
	if err != nil {
		log.Printf("err => %s", err)
	}
	return err
}
