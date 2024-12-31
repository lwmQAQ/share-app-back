package models

import "time"

type Apps struct {
	ID          string    `json:"id" gorm:"primaryKey;type:char(36);not null"` // 唯一标识，使用 CHAR(36) 存储 UUID
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`      // app 名字
	Type        int       `json:"type" gorm:"type:int;not null"`               // 类型1网页应用 2桌面应用 3博客
	Version     string    `json:"version" gorm:"type:varchar(50);not null"`    // 版本
	Author      string    `json:"author" gorm:"type:varchar(255);not null"`    // 作者
	Released    time.Time `json:"released" gorm:"type:datetime;not null"`      // 发布时间
	Readme      string    `json:"readme" gorm:"type:varchar(255)"`             // 说明文档
	Description string    `json:"description" gorm:"type:text"`                // 描述
	Icon        string    `json:"icon" gorm:"type:varchar(255)"`               // 图标 URL
	Url         string    `json:"url" gorm:"type:varchar(255)"`                // 下载地址
}
