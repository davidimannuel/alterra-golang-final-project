package auth

import (
	"context"
	"fmt"
	"keep-remind-app/businesses/jwt"
	"keep-remind-app/businesses/user"
	"keep-remind-app/helpers/encrypt"
)

type authUsecase struct {
	userUc user.Usecase
	jwtUc  jwt.Usecase
}

func NewUsecase(userUc user.Usecase, jwtUc jwt.Usecase) Usecase {
	return &authUsecase{
		userUc: userUc,
		jwtUc:  jwtUc,
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
	res.JWTToken, err = uc.jwtUc.GenerateToken(ctx, user.Id)
	if err != nil {
		return res, err
	}
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
	fmt.Println(user)
	if !encrypt.CheckPasswordHash(data.Password, user.Password) {
		return res, ErrInvalidPassword
	}
	res.JWTToken, err = uc.jwtUc.GenerateToken(ctx, user.Id)
	if err != nil {
		return
	}
	res.Name = user.Name
	res.Email = user.Email
	return
}
