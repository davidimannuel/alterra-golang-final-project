package redis

import (
	"context"
	"time"
)

type redisUsecase struct {
	redisRepository RedisRepository
}

func NewRedisUsecase(redisRepository RedisRepository) RedisUsecase {
	return &redisUsecase{
		redisRepository: redisRepository,
	}
}

func (repo *redisUsecase) Set(ctx context.Context, key string, value interface{}, expTime time.Duration) error {
	return repo.redisRepository.Set(ctx, key, value, expTime)
}

func (repo *redisUsecase) Get(ctx context.Context, key string) (string, error) {
	return repo.redisRepository.Get(ctx, key)
}

func (repo *redisUsecase) Del(ctx context.Context, key string) error {
	return repo.redisRepository.Del(ctx, key)
}
