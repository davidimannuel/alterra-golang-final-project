package redis

import (
	"context"
	"time"
)

type RedisUsecase interface {
	Set(ctx context.Context, key string, value interface{}, expTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type RedisRepository interface {
	Set(ctx context.Context, key string, value interface{}, expTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
