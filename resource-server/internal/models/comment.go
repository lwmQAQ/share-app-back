package models

import "time"

// User represents the user information stored redundantly in the comment.
type User struct {
	ID       uint64 `bson:"id" json:"id"`             // 用户 ID
	Username string `bson:"username" json:"username"` // 用户名
	Avatar   string `bson:"avatar" json:"avatar"`     // 用户头像 URL
}

// Comment represents a comment on a post with redundant user information.
type Comment struct {
	ID        uint64    `bson:"_id" json:"id"`                                // 评论 ID，MongoDB 会自动生成
	PostID    uint64    `bson:"postId" json:"postId"`                         // 所属帖子 ID
	ParentID  *uint64   `bson:"parentId,omitempty" json:"parentId,omitempty"` // 父评论 ID
	Content   string    `bson:"content" json:"content"`                       // 评论内容
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`                   // 评论时间
	Level     int       `bson:"level" json:"level"`                           // 评论层级（1 表示一级评论，2 表示二级评论，依此类推）
	Path      string    `bson:"path" json:"path"`                             // 用于表示评论的层级路径，例如 "1", "1.2", "1.2.3" 等
	Likes     uint64    `bson:"likes" json:"likes"`                           // 点赞数
	User      User      `bson:"user" json:"user"`                             // 冗余存储的用户信息
}
