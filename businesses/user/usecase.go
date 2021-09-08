package user

import (
	"context"
	"keep-remind-app/businesses"
	"keep-remind-app/helpers/encrypt"
)

type userUsecase struct {
	repository Repository
}

func NewUsecase(repository Repository) Usecase {
	return &userUsecase{
		repository: repository,
	}
}

func (uc userUsecase) Add(ctx context.Context, data *Domain) (res Domain, err error) {
	data.Password, err = encrypt.HashPassword(data.Password)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}
	result, err := uc.repository.Add(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	return result, err
}
