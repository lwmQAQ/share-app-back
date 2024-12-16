package models

import "time"

// Resource 定义了资源的结构体
type Resource struct {
	ID          StringObjectID `bson:"_id,omitempty" json:"id"`
	Title       string         `bson:"title" json:"title"`                   // 标题
	Tags        []string       `bson:"tags,omitempty" json:"tags,omitempty"` // 标签列表
	PublishTime time.Time      `bson:"updated_at" json:"updated_at"`         // 发布时间

}
