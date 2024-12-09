package models

type Url struct {
	Code      string `json:"code" gorm:"primaryKey"` // 设置 Code 为主键
	SourceURL string `json:"SourceUrl"`              // 修正拼写为 SourceURL
}
