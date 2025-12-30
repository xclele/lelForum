package redis

import (
	"context"
	"fmt"
	"lelForum/settings"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	_, err = client.Ping(context.Background()).Result()
	return
}

func Close() {
	_ = client.Close()
}
