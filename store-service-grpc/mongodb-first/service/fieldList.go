package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongodb-first/common"
	"mongodb-first/dao"
)

// GetColumnNameList 获取数据表属性信息
func GetColumnNameList(client *mongo.Client) []common.Table {

	var tableList []common.Table

	ctx := context.Background()
	collections, err := client.Database("article").ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Printf("err => %s", err)
	}

	for _, collectionName := range collections {
		var table common.Table
		table.SourceName = "mongodb1"                                                          //数据源名
		table.DatabaseName = "article"                                                         //数据库名
		table.TableName = collectionName                                                       //集合名
		table.RelateFlag, table.ColumnList, err = getFields(client, "article", collectionName) //字段列表
		if err != nil {
			log.Printf("err => %s", err)
		}

		//添加到数据表列表
		tableList = append(tableList, table)
		fmt.Printf("Database: %s, Collection: %s, Fields: %+v\n", "article", collectionName, table.ColumnList)
	}

	return tableList
}

// 获取集合中的字段
func getFields(client *mongo.Client, dbName, collectionName string) (string, []string, error) {
	cursor, err := client.Database(dbName).Collection(collectionName).Aggregate(context.Background(), []bson.M{{"$sample": bson.M{"size": 1}}})
	if err != nil {
		return "", nil, err
	}
	defer cursor.Close(context.Background())

	var doc bson.M
	if cursor.Next(context.Background()) {
		err := cursor.Decode(&doc)
		if err != nil {
			return "", nil, err
		}
	} else {
		return "", nil, fmt.Errorf("no documents found in collection '%s'", collectionName)
	}

	var fid string
	var fields []string
	for k := range doc {
		if k != "_id" && k[len(k)-2:] == "id" {
			fid = k
		}
		fields = append(fields, k)
	}
	return fid, fields, nil
}

func GetColumnData(table common.Table) []map[string]interface{} {
	var client *mongo.Client
	if table.DatabaseName == "article" {
		client = dao.Article
	}

	collection := client.Database(table.DatabaseName).Collection(table.TableName)

	var result []map[string]interface{}
	filter := bson.M{"_id": 0}

	for _, v := range table.ColumnList {
		filter[v] = 1
	}
	log.Printf("filter => %+v", filter)

	cur, err := collection.Find(context.Background(), bson.M{}, options.Find().SetProjection(filter))
	if err != nil {
		log.Printf("err => %s", err)
	}
	err = cur.All(context.Background(), &result)
	if err != nil {
		log.Printf("err => %s", err)
	}
	log.Printf("result => %v", result)

	return result
}
