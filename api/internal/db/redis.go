package db

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
	"time"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedis(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})
	cxt := context.Background()
	if err := rdb.Ping(cxt).Err(); err != nil {
		panic(err)
	}
	return &RedisClient{
		Client: rdb,
		Ctx:    cxt,
	}
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	var v interface{}
	switch val := value.(type) {
	case string, []byte, int, int64, float64, bool:
		v = val
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return err
		}
		v = string(b)
	}
	return r.Client.Set(r.Ctx, key, v, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *RedisClient) Del(keys ...string) error {
	return r.Client.Del(r.Ctx, keys...).Err()
}

func (r *RedisClient) Exists(key string) (bool, error) {
	cnt, err := r.Client.Exists(r.Ctx, key).Result()
	return cnt > 0, err
}
