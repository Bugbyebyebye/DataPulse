package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"store-service/dao"
	"store-service/model"
)

// CreateTable 创建数据仓库数据表
func (*StoreHandle) CreateTable(ctx *gin.Context) {
	id := ctx.GetInt("id")

	var jsonParam DataSource
	ctx.BindJSON(&jsonParam)
	tableName := jsonParam.SaveName
	columns := jsonParam.DatabaseList[0].TableList[0].ColumnList

	//log.Printf("param %+v", columns)
	//req := Req{Target: "getColumnData", Param: jsonParam.DatabaseList}
	//param, _ := json.Marshal(req)

	//
	////TODO 检查是否有重名表
	//table := dao.Warehouse.HasTable(jsonParam.SaveName)
	//if table {
	//	ctx.JSON(http.StatusOK, res.Fail(4001, "表名已经存在"))
	//}

	//创建新数据库
	dao.CreateTableBySQL(jsonParam.SaveName, columns)
	//创建用户数据库关联
	tableId, err := model.InitTable(jsonParam.SaveName, len(columns))
	if err != nil {
		log.Printf("err => %s", err)
	}
	err = model.InitTableUser(tableId, id)
	if err != nil {
		log.Printf("err => %s", err)
	}

	//data = append(data, clientData.Data)
	insertSql := CreateInsertSql(tableName, columns)
	log.Printf("insertSql => %s", insertSql)
	//for _, value := range clientData.Data.([]interface{}) {
	//	valueMap, ok := value.(map[string]interface{})
	//	if !ok {
	//		continue
	//	}
	//	values := make([]interface{}, 0, len(valueMap))
	//	for _, val := range valueMap {
	//		values = append(values, val)
	//	}
	//	log.Printf("values => %+v", values)
	//	err := dao.Warehouse.Exec(insertSql, values...).Error
	//	if err != nil {
	//		log.Printf("err => %s", err)
	//	}
	//}

	//log.Printf("bottom => %+v", data)
	//log.Printf("DatabaseList %+v", param.DatabaseList[0].TableList[0].ColumnList)
	//ctx.JSON(http.StatusOK, res.Success(data))
}

// AlertTable 向数据仓库数据表中追加字段
func (*StoreHandle) AlertTable(ctx *gin.Context) {
	var jsonParam DataSource
	ctx.BindJSON(&jsonParam)

	log.Printf("jsonParam %+v", jsonParam)
	println(jsonParam.SaveName)
	println(jsonParam.FromName)

	columns := jsonParam.DatabaseList[0].TableList[0].ColumnList

	dao.AlertTableBySQL(jsonParam.SaveName, columns)

	ctx.JSON(http.StatusOK, res.Success(200))
}

// CreateInsertSql 根据表名和字段列表生成动态插入的sql语句
func CreateInsertSql(tableName string, fields []string) string {
	var insertSql string
	for i, v := range fields {
		if i == 0 {
			insertSql = "insert into " + tableName + "(" + v
		} else {
			insertSql += "," + v
		}
	}
	insertSql += ")"
	for i, _ := range fields {
		if i == 0 {
			insertSql += " values(?"
		} else {
			insertSql += ",?"
		}
	}
	insertSql += ");"
	return insertSql
}
