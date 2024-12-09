package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"resource-server/internal/constants"

	"time"
)

type UrlUtil struct {
	BaseUrl   string
	Redisutil *RedisUtil
}

func NewUrlUtils(redis *RedisUtil, baseurl string) *UrlUtil {
	return &UrlUtil{
		BaseUrl:   baseurl,
		Redisutil: redis,
	}

}

func (u *UrlUtil) CreateShortLink(sourceUrl string) (string, error) {
	code := u.linkcodecreate(sourceUrl, 6)
	key := constants.BuildUrlKey(code)
	//可以增加重试机制
	err := u.Redisutil.CreateJsonCache(key, sourceUrl, 10*time.Hour)
	if err != nil {
		fmt.Println("生成短链失败")
		return "", err
	}
	return fmt.Sprintf("%s%s", u.BaseUrl, code), nil

}

func (u *UrlUtil) GetSourceUrl(code string) (string, error) {
	key := constants.BuildUrlKey(code)
	var sourceurl string
	err := u.Redisutil.GetJsonDataByKey(key, &sourceurl)
	//TODO 加入数据持久化
	if err != nil {
		fmt.Println("短链接失效")
		return "", err
	}
	return sourceurl, nil
}

func (u *UrlUtil) linkcodecreate(url string, length int) string {
	// 1. 对 URL 进行 SHA256 哈希
	hash := sha256.Sum256([]byte(url))

	// 2. 将哈希结果转换为 BASE64 URL 安全编码
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	// 3. 截取前 N 个字符作为短码
	if length > len(encoded) {
		length = len(encoded)
	}
	return encoded[:length]
}
