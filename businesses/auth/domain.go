package auth

import "context"

type Domain struct {
	Name     string
	Email    string
	Username string
	Password string
	JWTToken string
}

type Usecase interface {
	Register(ctx context.Context, data *Domain) (Domain, error)
	Login(ctx context.Context, data *Domain) (Domain, error)
}
