package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StringObjectID string

// 自定义 MarshalBSON，自动将 ObjectID 转换为 string
func (id *StringObjectID) MarshalBSON() ([]byte, error) {
	objectID, err := primitive.ObjectIDFromHex(string(*id))
	if err != nil {
		return nil, err
	}
	return bson.Marshal(objectID)
}

// 自定义 UnmarshalBSON，自动将 ObjectID 转换为 string
func (id *StringObjectID) UnmarshalBSON(data []byte) error {
	var objectID primitive.ObjectID
	if err := bson.Unmarshal(data, &objectID); err != nil {
		return err
	}
	*id = StringObjectID(objectID.Hex())
	return nil
}

// Post 表示一个帖子
type Post struct {
	ID          StringObjectID `bson:"_id,omitempty" json:"id"`
	Title       string         `bson:"title" json:"title"`                                   // 帖子标题
	Content     string         `bson:"content" json:"content"`                               // 富文本内容，HTML 格式
	Author      User           `bson:"author" json:"author"`                                 // 帖子的作者信息
	LikesCount  int64          `bson:"likes_count" json:"likes_count"`                       // 点赞数
	CommentsIDs []uint64       `bson:"comments_ids,omitempty" json:"comments_ids,omitempty"` // 评论的 ObjectID 列表
	Type        int            `bson:"type" json:"type,omitempty"`                           // 资源类型  1、是bt文件 2、是网盘资源
	Link        string         `bson:"link" json:"link,omitempty"`                           // 资源连接 如果是种子文件前端会上传文件然后返回一个url
	Tags        []string       `bson:"tags,omitempty" json:"tags,omitempty"`                 // 帖子标签列表
	CreatedAt   time.Time      `bson:"created_at" json:"created_at"`                         // 创建时间
	UpdatedAt   time.Time      `bson:"updated_at" json:"updated_at"`                         // 更新时间
}
