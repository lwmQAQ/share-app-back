package utils

import (
	"context"
	"fmt"
	"log"
	"resource-server/config"
	"resource-server/internal/models"
	"time"

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

	// 创建 ClientOptions 并设置连接池参数
	clientOpts := options.Client().
		ApplyURI(addr).
		SetMaxPoolSize(50).                  // 最大连接数
		SetMinPoolSize(10).                  // 最小连接数
		SetMaxConnIdleTime(30 * time.Second) // 空闲连接的最大存活时间

	client, err := mongo.Connect(context.TODO(), clientOpts)
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

func (u *MongoUtil) FullSearch(searchTerm string, collectionName string) ([]*models.Resource, error) {
	collection := u.mongoClient.Database("Resource").Collection(collectionName)
	// 设置查询选项，限制返回前10个结果
	findOptions := options.Find().SetLimit(10)
	// 定义投影，指定只返回 "name" 字段 `bson:"updated_at" json:"updated_at"`
	// Post 表示一个帖子

	projection := bson.D{{Key: "_id", Value: 1}, {Key: "title", Value: 1}, {Key: "updated_at", Value: 1}, {Key: "tags", Value: 1}} // 1表示包含此字段
	cursor, err := collection.Find(
		context.Background(),
		bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: searchTerm}}}},
		findOptions.SetProjection(projection), // 设置投影
	)
	if err != nil {
		log.Fatalf("Failed to execute text search: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*models.Resource
	// 迭代查询结果并输出
	for cursor.Next(context.Background()) {
		var result = new(models.Resource)
		if err := cursor.Decode(result); err != nil {
			log.Fatalf("Failed to decode document: %v", err)
			return nil, err
		}
		fmt.Println(result)
		results = append(results, result)
	}

	return results, nil
}
