package user

import (
	"context"
	"time"
)

type Domain struct {
	Id          int
	Username    string
	CountryCode string
	Phone       string
	Password    string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type Parameter struct {
	Id          int
	Username    string
	CountryCode string
	Phone       string
	Email       string
	Password    string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   *string
	OrderBy     string
	Sort        string
	Offset      int
	Limit       int
}

type Repository interface {
	// FindAll(ctx context.Context, parameter Parameter) ([]Domain, int, error)
	// SelectAll(ctx context.Context, parameter Parameter) ([]Domain, error)
	// FindByID(ctx context.Context, parameter Parameter) (Domain, error)
	// FindByTitle(ctx context.Context, parameter Parameter) (Domain, error)
	// FindByTitleOrNote(ctx context.Context, parameter Parameter) (Domain, error)
	Add(ctx context.Context, data *Domain) (Domain, error)
	// Edit(ctx context.Context, data *Domain) (Domain, error)
	// Delete(ctx context.Context, id int) (Domain, error)
}

type Usecase interface {
	// FindAll(ctx context.Context, parameter Parameter) ([]Domain, int, error)
	// SelectAll(ctx context.Context, parameter Parameter) ([]Domain, error)
	// FindByID(ctx context.Context, parameter Parameter) (Domain, error)
	// FindByTitle(ctx context.Context, parameter Parameter) (Domain, error)
	// FindByTitleOrNote(ctx context.Context, parameter Parameter) (Domain, error)
	Add(ctx context.Context, data *Domain) (Domain, error)
	// Edit(ctx context.Context, data *Domain) (Domain, error)
	// Delete(ctx context.Context, id int) (Domain, error)
}
