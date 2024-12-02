package cache

import (
	"fmt"
	"login-server/internal/constant"
	"login-server/internal/dao"
	"login-server/internal/models"
	"login-server/utils"
	"time"
)

type UserInfoCache struct {
	redisClient *utils.RedisUtil
	userDao     dao.UserDao
}

// NewUserInfoCache 创建 UserInfoCache 的实例
func NewUserInfoCache(redisClient *utils.RedisUtil, userDao dao.UserDao) *UserInfoCache {
	return &UserInfoCache{
		redisClient: redisClient,
		userDao:     userDao,
	}
}

// Get 根据用户ID获取用户信息
func (c *UserInfoCache) Get(id uint64) (*models.User, error) {
	var user models.User
	err := c.redisClient.GetJsonDataByKey(constant.BuildInfoKey(id), &user) // 从 Redis 获取用户信息
	if err != nil {
		return &user, err
	}
	return &user, nil
}

// GetBatch 根据多个用户ID批量获取用户信息
func (c *UserInfoCache) GetBatch(ids []uint64) (map[string]*models.User, error) {
	result := make(map[string]*models.User)
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = constant.BuildInfoKey(id) // 将 uint64 转换为字符串
	}
	interfaceResult := make(map[string]interface{})
	err := c.redisClient.BatchGetJsonDataByKey(keys, interfaceResult)
	if err != nil {
		return nil, err
	}
	for key, value := range interfaceResult {
		if user, ok := value.(models.User); ok {
			result[key] = &user
		} else {
			return nil, fmt.Errorf("类型转换失败，key: %s, value: %v", key, value)
		}
	}
	return result, nil
}

// Set 设置缓存
func (c *UserInfoCache) Set(id uint64, value *models.User, expiration time.Duration) error {
	return c.redisClient.CreateJsonCache(constant.BuildInfoKey(id), value, expiration) // 假设此方法已实现
}

// SetBatch 批量设置缓存
func (c *UserInfoCache) SetBatch(values map[uint64]*models.User, expiration time.Duration) error {
	for key, value := range values {
		err := c.Set(key, value, expiration)
		if err != nil {
			return err // 如果遇到错误，返回
		}
	}
	return nil
}

// Delete 删除缓存
func (c *UserInfoCache) Delete(id uint64) error {
	return c.redisClient.DeleteKey(constant.BuildInfoKey(id)) // 假设此方法已实现
}

// DeleteBatch 批量删除缓存
func (c *UserInfoCache) DeleteBatch(ids []uint64) error {
	for _, id := range ids {
		err := c.Delete(id)
		if err != nil {
			return err // 如果遇到错误，返回
		}
	}
	return nil
}

func (c *UserInfoCache) LoadCache(id uint64) (*models.User, error) {
	user, err := c.userDao.SelectUserById(id)
	user.Password = "******" //密码不可看
	if err != nil {          //数据库查询不到就写入一个空缓存防止击穿
		err = c.Set(id, &models.User{}, constant.USER_INFO_EX)
		if err != nil {
			return nil, err
		}
	}
	c.Set(id, user, constant.USER_INFO_EX)
	return user, nil
}
