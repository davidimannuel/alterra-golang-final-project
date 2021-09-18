package telegramUser_test

import (
	"context"
	"errors"
	"keep-remind-app/businesses"
	redisDomain "keep-remind-app/businesses/redis"
	_redisMock "keep-remind-app/businesses/redis/mocks"
	"keep-remind-app/businesses/telegramUser"
	_telegramUserMock "keep-remind-app/businesses/telegramUser/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	errCase = errors.New("error_case")
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("Test Find All | Valid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := []telegramUser.TelegramUserDomain{
			{
				ID:       1,
				UserID:   1,
				Username: "test",
			},
		}
		param := telegramUser.TelegramUserParameter{}
		telegramUserRepository.Mock.On("FindAll", context.Background(), &param).Return(data, nil).Once()
		result, err := telegramUserUsecase.FindAll(context.Background(), &param)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})
	t.Run("Test Find All | InValid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := []telegramUser.TelegramUserDomain{
			{
				ID:       1,
				UserID:   1,
				Username: "test",
			},
		}
		param := telegramUser.TelegramUserParameter{}
		telegramUserRepository.Mock.On("FindAll", context.Background(), &param).Return(data, errCase).Once()
		_, err := telegramUserUsecase.FindAll(context.Background(), &param)
		assert.NotNil(t, err)
	})
}

func TestFindAllPagination(t *testing.T) {
	t.Run("Test Find All Pagination | Valid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := []telegramUser.TelegramUserDomain{
			{
				ID:       1,
				UserID:   1,
				Username: "test",
			},
		}
		param := telegramUser.TelegramUserParameter{}
		telegramUserRepository.Mock.On("FindAllPagination", context.Background(), &param).Return(data, 1, nil).Once()
		result, _, err := telegramUserUsecase.FindAllPagination(context.Background(), &param)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Test Find All Pagination | InValid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := []telegramUser.TelegramUserDomain{
			{
				ID:       1,
				UserID:   1,
				Username: "test",
			},
		}
		param := telegramUser.TelegramUserParameter{}
		telegramUserRepository.Mock.On("FindAllPagination", context.Background(), &param).Return(data, 0, errCase).Once()
		_, _, err := telegramUserUsecase.FindAllPagination(context.Background(), &param)
		assert.NotNil(t, err)
	})
}

func TestFindOne(t *testing.T) {
	t.Run("Test Find One | Valid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := telegramUser.TelegramUserDomain{
			ID:       1,
			UserID:   1,
			Username: "test",
		}

		param := telegramUser.TelegramUserParameter{}
		telegramUserRepository.Mock.On("FindOne", context.Background(), &param).Return(data, nil).Once()
		_, err := telegramUserUsecase.FindOne(context.Background(), &param)
		assert.Nil(t, err)
	})

	t.Run("Test Find One | InValid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := telegramUser.TelegramUserDomain{
			ID:       1,
			UserID:   1,
			Username: "test",
		}

		param := telegramUser.TelegramUserParameter{}
		telegramUserRepository.Mock.On("FindOne", context.Background(), &param).Return(data, errCase).Once()
		_, err := telegramUserUsecase.FindOne(context.Background(), &param)
		assert.NotNil(t, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Test Add | Valid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := telegramUser.TelegramUserDomain{
			UserID:   1,
			Username: "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		// mock function called twice with different return values each call
		telegramUserRepository.Mock.On("FindOne", ctx, &telegramUser.TelegramUserParameter{UserID: 1}).
			Return(telegramUser.TelegramUserDomain{}, nil).Once()

		telegramUserRepository.Mock.On("FindOne", ctx, &telegramUser.TelegramUserParameter{Username: "test"}).
			Return(telegramUser.TelegramUserDomain{}, nil).Once()

		telegramUserRepository.Mock.On("Add", ctx, &data).Return(1, nil).Once()
		_, err := telegramUserUsecase.Add(ctx, &data)
		assert.Nil(t, err)
	})

	t.Run("Test Add | InValid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := telegramUser.TelegramUserDomain{
			ID:       1,
			UserID:   1,
			Username: "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		param := telegramUser.TelegramUserParameter{UserID: 1}
		telegramUserRepository.Mock.On("FindOne", ctx, &param).Return(data, nil).Once()
		telegramUserRepository.Mock.On("Add", ctx, &data).Return(1, nil).Once()
		_, err := telegramUserUsecase.Add(ctx, &data)
		assert.NotNil(t, err)
	})

	t.Run("Test Add | InValid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := telegramUser.TelegramUserDomain{
			ID:       1,
			UserID:   1,
			Username: "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		// param := telegramUser.TelegramUserParameter{UserID: 1}
		telegramUserRepository.Mock.On("FindOne", ctx, mock.AnythingOfType("*telegramUser.TelegramUserParameter")).Return(telegramUser.TelegramUserDomain{}, nil).Once()
		telegramUserRepository.Mock.On("FindOne", ctx, mock.AnythingOfType("*telegramUser.TelegramUserParameter")).Return(telegramUser.TelegramUserDomain{ID: 2}, nil).Once()
		telegramUserRepository.Mock.On("Add", ctx, &data).Return(1, errCase).Once()
		_, err := telegramUserUsecase.Add(ctx, &data)
		assert.NotNil(t, err)
	})
}
func TestGenerateActivatedOTP(t *testing.T) {
	t.Run("Test GenerateActivatedOTP | Valid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := telegramUser.TelegramUserDomain{
			ID:       1,
			UserID:   1,
			Username: "test",
		}
		param := telegramUser.TelegramUserParameter{BaseParameter: businesses.BaseParameter{ID: 1}}
		telegramUserRepository.Mock.On("FindOne", context.Background(), &param).Return(data, nil).Once()

		redisRepository.Mock.On("Set", context.Background(), mock.AnythingOfType("string"), data.ID, time.Minute*5).
			Return(nil).Once()

		redisRepository.Mock.On("Set", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("string"), time.Duration(0)).
			Return(nil).Once()

		_, err := telegramUserUsecase.GenerateActivatedOTP(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func TestActivated(t *testing.T) {
	t.Run("Test Activated | Valid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)
		data := telegramUser.TelegramUserDomain{
			ID:       1,
			UserID:   1,
			Username: "test",
		}
		redisRepository.Mock.On("Get", context.Background(), mock.AnythingOfType("string")).Return("1", nil).Once()

		param := telegramUser.TelegramUserParameter{Username: "test", BaseParameter: businesses.BaseParameter{ID: 1}}
		telegramUserRepository.Mock.On("FindOne", context.Background(), &param).Return(data, nil).Once()

		redisRepository.Mock.On("Set", context.Background(), mock.AnythingOfType("string"), data.UserID, time.Duration(0)).
			Return(nil).Once()

		redisRepository.Mock.On("Del", context.Background(), mock.AnythingOfType("string")).Return(nil).Once()

		redisRepository.Mock.On("Del", context.Background(), mock.AnythingOfType("string")).Return(nil).Once()

		err := telegramUserUsecase.Activated(context.Background(), data.Username, "123456")
		assert.Nil(t, err)
	})
}
func TestDeleted(t *testing.T) {
	t.Run("Test Delete | Valid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)

		telegramUserRepository.Mock.On("Delete", context.Background(), &telegramUser.TelegramUserDomain{}).Return(nil).Once()
		err := telegramUserUsecase.Delete(context.Background(), &telegramUser.TelegramUserDomain{})
		assert.Nil(t, err)
	})

	t.Run("Test Delete | InValid", func(t *testing.T) {
		var (
			redisRepository        _redisMock.RedisRepository
			redisUsecase           redisDomain.RedisUsecase
			telegramUserRepository _telegramUserMock.TelegramUserRepository
		)
		redisUsecase = redisDomain.NewRedisUsecase(&redisRepository)
		telegramUserUsecase := telegramUser.NewTelegramUserUsecase(&telegramUserRepository, redisUsecase)

		telegramUserRepository.Mock.On("Delete", context.Background(), &telegramUser.TelegramUserDomain{}).Return(errCase).Once()
		err := telegramUserUsecase.Delete(context.Background(), &telegramUser.TelegramUserDomain{})
		assert.NotNil(t, err)
	})
}
