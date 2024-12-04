package utils

import (
	"context"
	"fmt"
	"resource-server/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUtil struct {
	mongoClient  *mongo.Client
	DataBaseName string
}

func NewMongoUtil(logger *logrus.Logger, MongoConfig *config.MongoConfig) *MongoUtil {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Errorf("mongodb连接错误 %v", err)
		return nil
	}

	// 等待连接成功
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Errorf("mongodb连接错误 %v", err)
		return nil
	}

	fmt.Println("Connected to MongoDB!")
	return &MongoUtil{
		mongoClient:  client,
		DataBaseName: MongoConfig.DataBaseName,
	}
}

// 插入文档
func (u *MongoUtil) InsertDocument(collectionName string, document interface{}) error {
	collection := u.mongoClient.Database(u.DataBaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return fmt.Errorf("插入文档失败: %v", err)
	}
	return nil
}

// 删除文档
func (u *MongoUtil) DeleteDocument(collectionName string, filter interface{}) (int64, error) {
	collection := u.mongoClient.Database(u.DataBaseName).Collection(collectionName)
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, fmt.Errorf("删除文档失败: %v", err)
	}
	return result.DeletedCount, nil
}

// 查询文档
func (u *MongoUtil) SearchDocument(collectionName string, filter interface{}) ([]interface{}, error) {
	collection := u.mongoClient.Database(u.DataBaseName).Collection(collectionName)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("查询文档失败: %v", err)
	}
	defer cursor.Close(context.Background())

	var results []interface{}
	for cursor.Next(context.Background()) {
		var result interface{}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("解码文档失败: %v", err)
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("游标错误: %v", err)
	}
	return results, nil
}

// 更新文档
func (u *MongoUtil) UpdateDocument(collectionName string, filter interface{}, update interface{}) (int64, error) {
	collection := u.mongoClient.Database("your_database_name").Collection(collectionName)
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, fmt.Errorf("更新文档失败: %v", err)
	}
	return result.ModifiedCount, nil
}
