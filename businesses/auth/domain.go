package auth

import "context"

type AuthDomain struct {
	Name     string
	Email    string
	Username string
	Password string
	JWTToken string
}

type AuthUsecase interface {
	Register(ctx context.Context, data *AuthDomain) (jwtToken string, err error)
	Login(ctx context.Context, data *AuthDomain) (jwtToken string, err error)
}
