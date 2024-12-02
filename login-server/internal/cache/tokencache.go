package cache

import (
	"fmt"
	"login-server/internal/constant"
	"login-server/utils"
	"time"
)

type UserTokenCache struct {
	redisClient *utils.RedisUtil
}

func NewUserTokenCache(redisClient *utils.RedisUtil) *UserTokenCache {
	return &UserTokenCache{
		redisClient: redisClient,
	}
}

// Get 根据用户ID获取用户信息
func (c *UserTokenCache) Get(id uint64) (string, error) {
	var token string
	err := c.redisClient.GetJsonDataByKey(constant.BuildTokenKey(id), &token) // 从 Redis 获取用户信息
	if err != nil {
		return token, err
	}
	return token, nil
}

// GetBatch 根据多个用户ID批量获取用户信息
func (c *UserTokenCache) GetBatch(ids []uint64) (map[string]string, error) {
	result := make(map[string]string)
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = constant.BuildTokenKey(id) // 将 uint64 转换为字符串
	}
	interfaceResult := make(map[string]interface{})
	err := c.redisClient.BatchGetJsonDataByKey(keys, interfaceResult)
	if err != nil {
		return nil, err
	}
	for key, value := range interfaceResult {
		if token, ok := value.(string); ok {
			result[key] = token
		} else {
			return nil, fmt.Errorf("类型转换失败，key: %s, value: %v", key, value)
		}
	}

	return result, nil
}

// Set 设置缓存
func (c *UserTokenCache) Set(id uint64, value string, expiration time.Duration) error {
	return c.redisClient.CreateJsonCache(constant.BuildTokenKey(id), value, expiration) // 假设此方法已实现
}

// SetBatch 批量设置缓存
func (c *UserTokenCache) SetBatch(values map[uint64]string, expiration time.Duration) error {
	for key, value := range values {
		err := c.Set(key, value, expiration)
		if err != nil {
			return err // 如果遇到错误，返回
		}
	}
	return nil
}

// Delete 删除缓存
func (c *UserTokenCache) Delete(id uint64) error {
	return c.redisClient.DeleteKey(constant.BuildTokenKey(id)) // 假设此方法已实现
}

// DeleteBatch 批量删除缓存
func (c *UserTokenCache) DeleteBatch(reqs []uint64) error {
	for _, req := range reqs {
		err := c.Delete(req)
		if err != nil {
			return err // 如果遇到错误，返回
		}
	}
	return nil
}

func (c *UserTokenCache) FlushedEx(id uint64) error {

	return c.redisClient.FlushedEx(constant.BuildTokenKey(id), constant.USER_TOKEN_EX)
}
