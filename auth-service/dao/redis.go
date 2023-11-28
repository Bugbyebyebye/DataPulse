package dao

//具体的redis数据库操作

import (
	"commons/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	client *redis.Client
}

func init() {
	c := redis.NewClient(&redis.Options{
		Addr:     config.Conf.RC.Host,
		Password: config.Conf.RC.Password,
		DB:       config.Conf.RC.Db,
	})
	Rc = &RedisCache{
		client: c,
	}
}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.client.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.client.Get(ctx, key).Result()
	return result, err
}
