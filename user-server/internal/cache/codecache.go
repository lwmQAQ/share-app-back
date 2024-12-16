package cache

import (
	"time"
	"user-server/internal/constant"
	"user-server/utils"
)

type CodeCache struct {
	redisClient *utils.RedisUtil
}

func NewCodeCache(redisClient *utils.RedisUtil) *CodeCache {
	return &CodeCache{
		redisClient: redisClient,
	}
}

func (c *CodeCache) Get(email string) (string, error) {
	var token string
	err := c.redisClient.GetJsonDataByKey(constant.BuildCodeKey(email), &token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *CodeCache) Put(code string, email string) error {
	return c.redisClient.CreateJsonCache(constant.BuildCodeKey(email), code, 10*time.Minute)
}
