package redis_test

import (
	"context"
	"errors"
	"keep-remind-app/businesses/redis"
	_redisMock "keep-remind-app/businesses/redis/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	errCase = errors.New("error_case")
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSet(t *testing.T) {
	t.Run("Test Set | Valid", func(t *testing.T) {
		var (
			redisRepository _redisMock.RedisRepository
			redisUsecase    redis.RedisUsecase
		)
		redisUsecase = redis.NewRedisUsecase(&redisRepository)
		redisRepository.Mock.On("Set", context.Background(), "key", "value", time.Minute).Return(nil)
		err := redisUsecase.Set(context.Background(), "key", "value", time.Minute)
		assert.Nil(t, err)
	})

	t.Run("Test Set | InValid", func(t *testing.T) {
		var (
			redisRepository _redisMock.RedisRepository
			redisUsecase    redis.RedisUsecase
		)
		redisUsecase = redis.NewRedisUsecase(&redisRepository)
		redisRepository.Mock.On("Set", context.Background(), "key", "value", time.Minute).Return(errCase)
		err := redisUsecase.Set(context.Background(), "key", "value", time.Minute)
		assert.NotNil(t, err)
	})
}
func TestGet(t *testing.T) {
	t.Run("Test Get | Valid", func(t *testing.T) {
		var (
			redisRepository _redisMock.RedisRepository
			redisUsecase    redis.RedisUsecase
		)
		redisUsecase = redis.NewRedisUsecase(&redisRepository)
		redisRepository.Mock.On("Get", context.Background(), "key").Return("value", nil)
		result, err := redisUsecase.Get(context.Background(), "key")
		assert.Nil(t, err)
		assert.Equal(t, "value", result)
	})
	t.Run("Test Get | InValid", func(t *testing.T) {
		var (
			redisRepository _redisMock.RedisRepository
			redisUsecase    redis.RedisUsecase
		)
		redisUsecase = redis.NewRedisUsecase(&redisRepository)
		redisRepository.Mock.On("Get", context.Background(), "key").Return("", errCase)
		_, err := redisUsecase.Get(context.Background(), "key")
		assert.NotNil(t, err)
	})
}
func TestDel(t *testing.T) {
	t.Run("Test Del | Valid", func(t *testing.T) {
		var (
			redisRepository _redisMock.RedisRepository
			redisUsecase    redis.RedisUsecase
		)
		redisUsecase = redis.NewRedisUsecase(&redisRepository)
		redisRepository.Mock.On("Del", context.Background(), "key").Return(nil)
		err := redisUsecase.Del(context.Background(), "key")
		assert.Nil(t, err)
	})
	t.Run("Test Del | InValid", func(t *testing.T) {
		var (
			redisRepository _redisMock.RedisRepository
			redisUsecase    redis.RedisUsecase
		)
		redisUsecase = redis.NewRedisUsecase(&redisRepository)
		redisRepository.Mock.On("Del", context.Background(), "key").Return(errCase)
		err := redisUsecase.Del(context.Background(), "key")
		assert.NotNil(t, err)
	})
}
