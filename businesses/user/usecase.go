package user

import (
	"context"
	"keep-remind-app/businesses"
	"keep-remind-app/helpers/encrypt"
)

type userUsecase struct {
	repository UserRepository
}

func NewUserUsecase(repository UserRepository) UserUsecase {
	return &userUsecase{
		repository: repository,
	}
}

func (uc userUsecase) Add(ctx context.Context, data *UserDomain) (res UserDomain, err error) {
	data.Password, err = encrypt.HashPassword(data.Password)
	if err != nil {
		return res, businesses.ErrInternalServer
	}
	res, err = uc.repository.Add(ctx, data)
	if err != nil {
		return res, businesses.ErrInternalServer
	}
	return
}

func (uc userUsecase) FindByEmail(ctx context.Context, param *UserParameter) (res UserDomain, err error) {
	res, err = uc.repository.FindByEmail(ctx, param)
	if err != nil {
		return res, businesses.ErrNotFound
	}
	return
}

func (uc userUsecase) FindByID(ctx context.Context, param *UserParameter) (res UserDomain, err error) {
	res, err = uc.repository.FindByID(ctx, param)
	if err != nil {
		return res, businesses.ErrNotFound
	}
	return
}
