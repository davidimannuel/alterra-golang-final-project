package redis

import (
	"context"
	redisDomain "keep-remind-app/businesses/redis"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	Client *redis.Client
}

func NewRedisRepository(client *redis.Client) redisDomain.RedisRepository {
	return &redisRepository{
		Client: client,
	}
}

func (repo *redisRepository) Set(ctx context.Context, key string, value interface{}, expTime time.Duration) error {
	return repo.Client.Set(ctx, key, value, expTime).Err()
}

func (repo *redisRepository) Get(ctx context.Context, key string) (string, error) {
	return repo.Client.Get(ctx, key).Result()
}
