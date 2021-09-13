package note

import (
	"context"
	"time"
)

type Domain struct {
	ID         int
	UserID     int
	Title      string
	Note       string
	ReminderAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type Parameter struct {
	ID         int
	UserID     int
	Title      string
	Note       string
	ReminderAt string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  *string
	OrderBy    string
	Sort       string
	Offset     int
	Limit      int
}

type Repository interface {
	FindAllPagination(ctx context.Context, parameter Parameter) ([]Domain, int, error)
	FindAll(ctx context.Context, parameter Parameter) ([]Domain, error)
	FindOne(ctx context.Context, parameter Parameter) (Domain, error)
	Add(ctx context.Context, data *Domain) (int, error)
	Edit(ctx context.Context, data *Domain) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type Usecase interface {
	// FindAllPagination(ctx context.Context, parameter Parameter) ([]Domain, int, error)
	FindAll(ctx context.Context, parameter Parameter) ([]Domain, error)
	FindByID(ctx context.Context, parameter Parameter) (Domain, error)
	// FindByTitle(ctx context.Context, parameter Parameter) (Domain, error)
	// FindByTitleOrNote(ctx context.Context, parameter Parameter) (Domain, error)
	Add(ctx context.Context, data *Domain) (Domain, error)
	AddWithImageBytes(ctx context.Context, title string, imageBytes []byte) (Domain, error)
	// Edit(ctx context.Context, data *Domain) (Domain, error)
	// Delete(ctx context.Context, id int) (Domain, error)
}
