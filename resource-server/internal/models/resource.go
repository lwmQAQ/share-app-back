package models

import "time"

// Resource 定义了资源的结构体
type Resource struct {
	Title       string    `json:"title"`        // 标题
	Tags        []string  `json:"tags"`         // 标签列表
	PublishTime time.Time `json:"publish_time"` // 发布时间
	Publisher   string    `json:"publisher"`    // 发布者
	LikeCount   int       `json:"like_count"`   // 点击量
}
