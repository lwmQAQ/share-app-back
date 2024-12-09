package models

type Url struct {
	Code      string `gorm:"type:varchar(255);primaryKey" json:"code"`     // 主键
	SourceURL string `gorm:"type:varchar(255);not null" json:"source_url"` // 普通字段
}
