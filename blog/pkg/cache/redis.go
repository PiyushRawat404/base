package cache

import (
	"blog/pkg/config"
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func LoadRedis(cfg config.Config) (*redis.Client, error) {
	dbIndex, err := strconv.Atoi(cfg.RedisDB)
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_DB: %w", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       dbIndex,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		_ = client.Close()
		return nil, err
	}

	fmt.Println("Connected to redis")
	return client, nil
}
