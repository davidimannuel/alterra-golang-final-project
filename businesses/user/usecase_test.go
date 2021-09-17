package user_test

import (
	"context"
	"errors"
	"keep-remind-app/businesses/user"
	_userMock "keep-remind-app/businesses/user/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errCase = errors.New("error_case")
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFindOne(t *testing.T) {
	t.Run("Find One | Valid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			ID:   1,
			Name: "Label 1",
		}
		param := user.UserParameter{}
		userRepository.On("FindOne", context.Background(), &param).Return(userData, nil).Once()
		result, err := userUsecase.FindOne(context.Background(), &param)
		assert.Nil(t, err)
		assert.Equal(t, 1, result.ID)
	})

	t.Run("Find One | InValid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			ID:   1,
			Name: "Label 1",
		}
		param := user.UserParameter{}
		userRepository.On("FindOne", context.Background(), &param).Return(userData, errCase).Once()
		_, err := userUsecase.FindOne(context.Background(), &param)
		assert.NotNil(t, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add | Valid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			Name:     "Label 1",
			Password: "123",
		}
		param := user.UserParameter{ID: userData.ID}
		userRepository.On("Add", context.Background(), &userData).Return(user.UserDomain{}, nil).Once()
		userRepository.On("FindOne", context.Background(), &param).Return(user.UserDomain{}, nil).Once()
		_, err := userUsecase.Add(context.Background(), &userData)
		assert.Nil(t, err)
	})

	t.Run("Add | InValid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			Name:     "Label 1",
			Password: "",
		}
		param := user.UserParameter{ID: userData.ID}
		userRepository.On("Add", context.Background(), &userData).Return(user.UserDomain{}, errCase).Once()
		userRepository.On("FindOne", context.Background(), &param).Return(user.UserDomain{}, nil).Once()
		_, err := userUsecase.Add(context.Background(), &userData)
		assert.NotNil(t, err)
	})

	t.Run("Add | InValid Email", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			Email:    "test@gmail.com",
			Name:     "Label 1",
			Password: "",
		}
		param := user.UserParameter{Email: "test@gmail.com"}
		userRepository.On("Add", context.Background(), &userData).Return(user.UserDomain{}, nil).Once()
		userRepository.On("FindOne", context.Background(), &param).Return(user.UserDomain{ID: 1}, nil).Once()
		_, err := userUsecase.Add(context.Background(), &userData)
		assert.NotNil(t, err)
	})
}
func TestEdit(t *testing.T) {
	t.Run("Edit | Valid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			ID:       1,
			Name:     "Label 1",
			Password: "123",
		}
		param := user.UserParameter{ID: userData.ID}
		userRepository.On("Edit", context.Background(), &userData).Return(nil).Once()
		userRepository.On("FindOne", context.Background(), &param).Return(user.UserDomain{}, nil).Once()
		err := userUsecase.Edit(context.Background(), &userData)
		assert.Nil(t, err)
	})

	t.Run("Edit | InValid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			ID:       1,
			Name:     "Label 1",
			Password: "123",
		}
		param := user.UserParameter{ID: userData.ID}
		userRepository.On("Edit", context.Background(), &userData).Return(nil).Once()
		userRepository.On("FindOne", context.Background(), &param).Return(user.UserDomain{}, errCase).Once()
		err := userUsecase.Edit(context.Background(), &userData)
		assert.NotNil(t, err)
	})

	t.Run("Edit | InValid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userData := user.UserDomain{
			ID:       1,
			Name:     "Label 1",
			Password: "123",
		}
		param := user.UserParameter{ID: userData.ID}
		userRepository.On("Edit", context.Background(), &userData).Return(errCase).Once()
		userRepository.On("FindOne", context.Background(), &param).Return(user.UserDomain{}, nil).Once()
		err := userUsecase.Edit(context.Background(), &userData)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userRepository.On("Delete", context.Background(), 1).Return(nil).Once()
		err := userUsecase.Delete(context.Background(), 1)
		assert.Nil(t, err)
	})
	t.Run("Delete | InValid", func(t *testing.T) {
		var (
			userRepository _userMock.UserRepository
			userUsecase    user.UserUsecase
		)
		userUsecase = user.NewUserUsecase(&userRepository)
		userRepository.On("Delete", context.Background(), 1).Return(errCase).Once()
		err := userUsecase.Delete(context.Background(), 1)
		assert.NotNil(t, err)
	})
}
