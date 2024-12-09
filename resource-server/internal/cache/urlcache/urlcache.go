package urlcache

import (
	"fmt"
	"resource-server/internal/constants"
	"resource-server/internal/dao"
	"resource-server/internal/models"
	"resource-server/utils/redisutils"

	"time"
)

type UrlCache struct {
	redisClient *redisutils.RedisUtil
	UrlDao      dao.UrlDao
}

func NewUrlCache(redisClient *redisutils.RedisUtil, urlDao dao.UrlDao) *UrlCache {
	return &UrlCache{
		redisClient: redisClient,
		UrlDao:      urlDao,
	}
}

func (c *UrlCache) Get(code string) (*string, error) {
	var sourceURL = new(string) // 初始化指针

	fmt.Println(constants.BuildUrlKey(code))
	err := c.redisClient.GetJsonDataByKey(constants.BuildUrlKey(code), sourceURL)
	if err != nil {
		fmt.Println(err)
		sourceURL, err = c.LoadCache(code)
		if err != nil {
			return nil, err
		}
		return sourceURL, nil
	}
	return sourceURL, nil
}

// Set 设置缓存
func (c *UrlCache) Set(value *models.Url, expiration time.Duration) error {
	return c.redisClient.CreateJsonCache(constants.BuildUrlKey(value.Code), value.SourceURL, expiration) // 假设此方法已实现
}

func (c *UrlCache) LoadCache(code string) (*string, error) {
	url, err := c.UrlDao.GetUrlByCode(code)
	if err != nil { //数据库查询不到就写入一个空缓存防止击穿
		err = c.Set(url, 0) //永不过期
		if err != nil {
			return nil, err
		}
	}
	err = c.Set(url, 0)
	if err != nil {
		fmt.Println("写入缓存失败")
	}
	return &url.SourceURL, nil
}
