package urlutils

import (
	"crypto/sha256"
	"encoding/base64"

	"resource-server/internal/cache/urlcache"
	"resource-server/internal/models"
)

type UrlUtil struct {
	BaseUrl  string
	urlCache *urlcache.UrlCache
}

func NewUrlUtils(baseurl string, urlCache *urlcache.UrlCache) *UrlUtil {
	return &UrlUtil{
		BaseUrl:  baseurl,
		urlCache: urlCache,
	}

}

func (u *UrlUtil) CreateShortLink(sourceUrl string) (string, error) {
	code := u.linkcodecreate(sourceUrl, 6)
	err := u.urlCache.Set(&models.Url{
		Code:      code,
		SourceURL: sourceUrl,
	}, 0)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (u *UrlUtil) GetSourceUrl(code string) (string, error) {
	url, err := u.urlCache.Get(code)
	if err != nil {
		return "", err
	}
	return *url, err

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
