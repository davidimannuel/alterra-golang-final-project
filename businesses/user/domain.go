package user

import (
	"context"
	"keep-remind-app/businesses"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Parameter struct {
	Id        int
	Name      string
	Password  string
	Email     string
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
	businesses.BaseParameter
}

type Repository interface {
	// FindAll(ctx context.Context, parameter Parameter) ([]Domain, int, error)
	// SelectAll(ctx context.Context, parameter Parameter) ([]Domain, error)
	// FindByID(ctx context.Context, parameter Parameter) (Domain, error)
	FindByEmail(ctx context.Context, parameter Parameter) (Domain, error)
	Add(ctx context.Context, data *Domain) (Domain, error)
	// Edit(ctx context.Context, data *Domain) (Domain, error)
	// Delete(ctx context.Context, id int) (Domain, error)
}

type Usecase interface {
	// FindAll(ctx context.Context, parameter Parameter) ([]Domain, int, error)
	// SelectAll(ctx context.Context, parameter Parameter) ([]Domain, error)
	// FindByID(ctx context.Context, parameter Parameter) (Domain, error)
	FindByEmail(ctx context.Context, parameter Parameter) (Domain, error)
	Add(ctx context.Context, data *Domain) (Domain, error)
	// Edit(ctx context.Context, data *Domain) (Domain, error)
	// Delete(ctx context.Context, id int) (Domain, error)
}
