package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mongodb-first/common"
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
		table.TableName = collectionName
		table.ColumnList, err = getFields(client, "article", collectionName)
		if err != nil {
			log.Printf("err => %s", err)
		}
		tableList = append(tableList, table)
		fmt.Printf("Database: %s, Collection: %s, Fields: %+v\n", "article", collectionName, table.ColumnList)
	}

	return tableList
}

func getFields(client *mongo.Client, dbName, collectionName string) ([]string, error) {
	cursor, err := client.Database(dbName).Collection(collectionName).Aggregate(context.Background(), []bson.M{{"$sample": bson.M{"size": 1}}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var doc bson.M
	if cursor.Next(context.Background()) {
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no documents found in collection '%s'", collectionName)
	}

	var fields []string
	for k := range doc {
		fields = append(fields, k)
	}
	return fields, nil
}
