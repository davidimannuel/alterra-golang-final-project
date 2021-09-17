package user

import (
	"context"
	"errors"
	"keep-remind-app/businesses"
	"keep-remind-app/helpers/encrypt"
)

type userUsecase struct {
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uc *userUsecase) validate(ctx context.Context, data *UserDomain) (err error) {
	user := UserDomain{}
	if data.ID != 0 {
		user, err = uc.userRepository.FindOne(ctx, &UserParameter{ID: data.ID})
		if err != nil {
			return errors.New("invalid user")
		}
	}
	if data.Email != "" && data.Email != user.Email {
		exist, _ := uc.userRepository.FindOne(ctx, &UserParameter{Email: data.Email})
		if exist.ID != 0 {
			return errors.New("email already used")
		}
	}
	return
}

func (uc *userUsecase) FindOne(ctx context.Context, param *UserParameter) (res UserDomain, err error) {
	res, err = uc.userRepository.FindOne(ctx, param)
	if err != nil {
		return res, businesses.ErrNotFound
	}
	return
}

func (uc *userUsecase) Add(ctx context.Context, data *UserDomain) (res UserDomain, err error) {
	err = uc.validate(ctx, data)
	if err != nil {
		return res, err
	}
	data.Password, _ = encrypt.HashPassword(data.Password)

	res, err = uc.userRepository.Add(ctx, data)
	if err != nil {
		return res, businesses.ErrInternalServer
	}
	return
}

func (uc *userUsecase) Edit(ctx context.Context, data *UserDomain) error {
	err := uc.validate(ctx, data)
	if err != nil {
		return err
	}
	return uc.userRepository.Edit(ctx, data)
}

func (uc *userUsecase) Delete(ctx context.Context, id int) error {
	return uc.userRepository.Delete(ctx, id)
}
