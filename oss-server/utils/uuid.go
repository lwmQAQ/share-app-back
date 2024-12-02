package utils

import "github.com/google/uuid"

// NewUUID 生成并返回一个新的 UUID 字符串
func NewUUID() string {
	// 生成新的 UUID
	newUUID := uuid.New()
	return newUUID.String() // 将 UUID 转换为字符串并返回
}
