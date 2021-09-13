package auth

import (
	"context"
	"keep-remind-app/businesses"
	userDomain "keep-remind-app/businesses/user"
	"keep-remind-app/helpers/encrypt"
	"keep-remind-app/server/middlewares"
)

type authUsecase struct {
	userUc  userDomain.Usecase
	jwtAuth *middlewares.ConfigJWT
}

func NewUsecase(userUc userDomain.Usecase, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &authUsecase{
		userUc:  userUc,
		jwtAuth: jwtAuth,
	}
}

func (uc *authUsecase) Register(ctx context.Context, data *Domain) (res Domain, err error) {
	user, _ := uc.userUc.FindByEmail(ctx, userDomain.Parameter{Email: data.Email})
	if user.ID != 0 {
		return res, businesses.ErrDuplicateData
	}
	user, err = uc.userUc.Add(ctx, &userDomain.Domain{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		return res, err
	}
	res.JWTToken = uc.jwtAuth.GenerateToken(user.ID)
	res.Name = data.Name
	res.Email = data.Email
	res.Password = user.Password
	return
}

func (uc *authUsecase) Login(ctx context.Context, data *Domain) (res Domain, err error) {
	user, err := uc.userUc.FindByEmail(ctx, userDomain.Parameter{Email: data.Email})
	if err != nil {
		return
	}
	if !encrypt.CheckPasswordHash(data.Password, user.Password) {
		return res, ErrInvalidPassword
	}

	res.JWTToken = uc.jwtAuth.GenerateToken(user.ID)
	res.Name = user.Name
	res.Email = user.Email
	return
}
