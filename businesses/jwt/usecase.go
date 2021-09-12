package jwt

import (
	"context"
	"keep-remind-app/drivers/jwt"
)

type jwtUsecase struct {
	config *jwt.ConfigJWT
}

func NewUsecase(config *jwt.ConfigJWT) Usecase {
	return &jwtUsecase{
		config: config,
	}
}

func (uc jwtUsecase) GenerateToken(ctx context.Context, userId int) (res string, err error) {
	return uc.config.GenerateToken(userId)
}
