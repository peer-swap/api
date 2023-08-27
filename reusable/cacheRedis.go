package reusable

import (
	"github.com/go-redis/redis"
	"time"
)

type CacheRedisConfig struct {
	RedisUrl      string
	RedisPassword string
}

type CacheRedis struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *CacheRedis {
	return &CacheRedis{client: client}
}

func NewCacheWithRedisFromConfig(config CacheRedisConfig) *CacheRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: config.RedisPassword,
	})

	return NewRedisCache(client)
}

func (c CacheRedis) Set(key string, value string, expiry time.Duration) error {
	err := c.client.Set(key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c CacheRedis) Get(key string) (string, error) {
	value, err := c.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c CacheRedis) Forget(key string) error {
	_, err := c.client.Del(key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c CacheRedis) GetOrDefault(key string, fail string) string {
	value, err := c.client.Get(key).Result()
	if err != nil {
		return fail
	}

	return value
}
