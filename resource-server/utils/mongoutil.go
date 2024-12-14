package utils

import (
	"context"
	"fmt"
	"resource-server/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUtil struct {
	mongoClient  *mongo.Client
	DataBaseName string
}

func NewMongoUtil(logger *logrus.Logger, MongoConfig *config.MongoConfig) *MongoUtil {
	// 格式化连接地址，包含用户名和密码
	addr := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		MongoConfig.UserName, MongoConfig.Password, MongoConfig.Host, MongoConfig.Port)
	fmt.Println(addr)
	clientOptions := options.Client().ApplyURI(addr)
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
func (u *MongoUtil) InsertDocument(collectionName string, document interface{}) (interface{}, error) {
	collection := u.mongoClient.Database(u.DataBaseName).Collection(collectionName)
	result, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return nil, fmt.Errorf("插入文档失败: %v", err)
	}
	return result.InsertedID, nil
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
func (u *MongoUtil) SearchDocument(collectionName string, filter interface{}, resp interface{}) error {
	collection := u.mongoClient.Database(u.DataBaseName).Collection(collectionName)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("查询文档失败: %v", err)
	}
	defer cursor.Close(context.Background())
	//定义切片
	var r []interface{}
	//查询
	for cursor.Next(context.Background()) {
		var result interface{}
		if err := cursor.Decode(&result); err != nil {
			return fmt.Errorf("解码文档失败: %v", err)
		}
		r = append(r, result)
	}
	//赋值返回
	resp = r
	if err := cursor.Err(); err != nil {
		return fmt.Errorf("游标错误: %v", err)
	}
	return nil
}
func (u *MongoUtil) SearchDocumentByID(collectionName string, id interface{}, result interface{}) error {
	collection := u.mongoClient.Database(u.DataBaseName).Collection(collectionName)

	// 构建过滤条件
	filter := bson.M{"_id": id}

	// 执行查询
	err := collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("未找到文档: %v", err)
		}
		return fmt.Errorf("查询文档失败: %v", err)
	}
	return nil
}

// 更新文档
func (u *MongoUtil) UpdateDocument(collectionName string, filter interface{}, update interface{}) (int64, error) {
	collection := u.mongoClient.Database("Resource").Collection(collectionName)
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, fmt.Errorf("更新文档失败: %v", err)
	}
	return result.ModifiedCount, nil
}
