package user

import (
	"context"
	"keep-remind-app/businesses"
	"time"
)

type UserDomain struct {
	ID        int
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserParameter struct {
	ID        int
	Name      string
	Password  string
	Email     string
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
	businesses.BaseParameter
}

type UserRepository interface {
	FindOne(ctx context.Context, param *UserParameter) (UserDomain, error)
	Add(ctx context.Context, data *UserDomain) (UserDomain, error)
	Edit(ctx context.Context, data *UserDomain) error
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	FindOne(ctx context.Context, param *UserParameter) (UserDomain, error)
	Add(ctx context.Context, data *UserDomain) (UserDomain, error)
	Edit(ctx context.Context, data *UserDomain) error
	Delete(ctx context.Context, id int) error
}
