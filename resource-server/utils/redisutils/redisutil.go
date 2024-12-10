package redisutils

import (
	"encoding/json"
	"fmt"
	"log"
	"resource-server/config"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type RedisUtil struct {
	client *redis.Client
}

func NewRedisUtil(redisConfig *config.RedisConfig) *RedisUtil {
	addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)

	client := redis.NewClient(&redis.Options{
		Addr:         addr,                 // Redis服务器地址
		Password:     redisConfig.Password, // Redis密码，如果没有可以留空
		DB:           redisConfig.DBName,   // Redis数据库编号
		PoolSize:     10,                   // 设置连接池中最大连接数
		MinIdleConns: 3,                    // 设置最小空闲连接数
	})
	if err := client.Ping().Err(); err != nil {
		log.Fatalf("无法连接到Redis: %v", err)
	}
	return &RedisUtil{
		client: client,
	}
}

func (redisUtil *RedisUtil) CreateJsonCache(key string, data interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("数据序列化失败: %v", err)
	}
	// 将数据存储到Redis中，使用Set方法
	err = redisUtil.client.Set(key, jsonData, expiration).Err()
	if err != nil {
		return fmt.Errorf("存储数据到Redis时出错: %v", err)
	}
	return nil
}

func (redisUtil *RedisUtil) GetJsonDataByKey(key string, result interface{}) error {
	// 从Redis中获取键对应的值
	jsonData, err := redisUtil.client.Get(key).Result()
	if err != nil {
		// 如果Redis中不存在该键，返回一个特定的错误提示
		if err == redis.Nil {
			return fmt.Errorf("键 %s 在Redis中不存在", key)
		}
		return fmt.Errorf("从Redis中获取数据时出错: %v", err)
	}
	// 将获取到的JSON数据反序列化到result中
	err = json.Unmarshal([]byte(jsonData), result)

	if err != nil {
		return fmt.Errorf("数据反序列化失败: %v", err)
	}

	return nil
}

func (redisUtil *RedisUtil) BatchGetJsonDataByKey(keys []string, result map[string]interface{}) error {
	values, err := redisUtil.client.MGet(keys...).Result()
	if err != nil {
		return fmt.Errorf("从Redis中获取数据时出错: %v", err)
	}
	// 将结果转换为 map
	// 将结果转换为 map，使用 interface{} 以便兼容不同类型
	for i, key := range keys {
		result[key] = values[i] // 直接存储 interface{}
	}
	return nil
}

func (r *RedisUtil) DeleteKey(key string) error {
	return r.client.Del(key).Err()
}

func (r *RedisUtil) FlushedEx(key string, expiration time.Duration) error {
	return r.client.Expire(key, expiration).Err()
}
func (redisUtil *RedisUtil) GetIntByKey(key string) (int64, error) {
	// 从Redis中获取键对应的值
	value, err := redisUtil.client.Get(key).Result()
	if err != nil {
		// 如果Redis中不存在该键，返回一个特定的错误提示
		if err == redis.Nil {
			return 0, fmt.Errorf("键 %s 在Redis中不存在", key)
		}
		return 0, fmt.Errorf("从Redis中获取数据时出错: %v", err)
	}
	// 将获取到的字符串数据转换为整数
	intValue, err := strconv.Atoi(value) // 将字符串转换为整数
	if err != nil {
		return 0, fmt.Errorf("数据转换为整数时出错: %v", err)
	}

	return int64(intValue), nil
}
func (redisUtil *RedisUtil) SerIntByKey(key string, count int64, ttl time.Duration) error {
	err := redisUtil.client.Set(key, count, ttl).Err()
	return err
}

func (redisUtil *RedisUtil) GetLastUpdate(key string) (int64, error) {
	key = fmt.Sprintf("%s%s", key, ":last_updated")
	times, err := redisUtil.client.Get(key).Result()
	if err == redis.Nil {
		return time.Now().Unix(), nil
	} else if err != nil {
		log.Printf("获取上次更新时间出错: %v", err)
		return 0, err
	}
	lastUpdatedTime, _ := strconv.Atoi(times)
	return int64(lastUpdatedTime), nil

}
func (r *RedisUtil) Close() {
	r.client.Close()
}
