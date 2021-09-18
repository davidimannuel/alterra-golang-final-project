package auth

import (
	"context"
	userDomain "keep-remind-app/businesses/user"
	"keep-remind-app/helpers/encrypt"
	"keep-remind-app/server/middlewares"
)

type authUsecase struct {
	userUc  userDomain.UserUsecase
	jwtAuth *middlewares.ConfigJWT
}

func NewAuthUsecase(userUc userDomain.UserUsecase, jwtAuth *middlewares.ConfigJWT) AuthUsecase {
	return &authUsecase{
		userUc:  userUc,
		jwtAuth: jwtAuth,
	}
}

func (uc *authUsecase) Register(ctx context.Context, data *AuthDomain) (jwtToken string, err error) {
	user, err := uc.userUc.Add(ctx, &userDomain.UserDomain{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		return "", err
	}
	return uc.jwtAuth.GenerateToken(user.ID), nil
}

func (uc *authUsecase) Login(ctx context.Context, data *AuthDomain) (jwtToken string, err error) {
	user, err := uc.userUc.FindOne(ctx, &userDomain.UserParameter{Email: data.Email})
	if err != nil {
		return "", err
	}
	if !encrypt.CheckPasswordHash(data.Password, user.Password) {
		return "", ErrInvalidPassword
	}

	return uc.jwtAuth.GenerateToken(user.ID), nil
}
