package auth

import "context"

type Domain struct {
	Name     string
	Email    string
	Username string
	Password string
	JWTToken string
}

type Repository interface {
	Register(ctx context.Context, data *Domain) (Domain, error)
	Login(ctx context.Context, data *Domain) (Domain, error)
}

type Usecase interface {
	Register(ctx context.Context, data *Domain) (Domain, error)
	Login(ctx context.Context, data *Domain) (Domain, error)
}
