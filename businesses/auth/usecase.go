package auth

import (
	"context"
	"keep-remind-app/businesses/user"
	"keep-remind-app/helpers/encrypt"
	"keep-remind-app/server/middlewares"
)

type authUsecase struct {
	userUc  user.Usecase
	jwtAuth *middlewares.ConfigJWT
}

func NewUsecase(userUc user.Usecase, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &authUsecase{
		userUc:  userUc,
		jwtAuth: jwtAuth,
	}
}

func (uc *authUsecase) Register(ctx context.Context, data *Domain) (res Domain, err error) {
	user, err := uc.userUc.Add(ctx, &user.Domain{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		return res, err
	}
	res.JWTToken = uc.jwtAuth.GenerateToken(user.Id)
	res.Name = data.Name
	res.Email = data.Email
	res.Password = user.Password
	return
}

func (uc *authUsecase) Login(ctx context.Context, data *Domain) (res Domain, err error) {
	user, err := uc.userUc.FindByEmail(ctx, user.Parameter{Email: data.Email})
	if err != nil {
		return
	}
	if !encrypt.CheckPasswordHash(data.Password, user.Password) {
		return res, ErrInvalidPassword
	}

	res.JWTToken = uc.jwtAuth.GenerateToken(user.Id)
	res.Name = user.Name
	res.Email = user.Email
	return
}
