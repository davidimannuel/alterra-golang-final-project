package note

import (
	"context"
	"keep-remind-app/businesses"
	"time"
)

type NoteDomain struct {
	ID         int
	UserID     int
	Title      string
	Note       string
	Labels     []LabelDomain
	ReminderAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type LabelDomain struct {
	ID     int
	UserID int
	Name   string
}

type NoteParameter struct {
	UserID     int
	Title      string
	Note       string
	ReminderAt string
	businesses.BaseParameter
}

type NoteRepository interface {
	FindAllPagination(ctx context.Context, param *NoteParameter) ([]NoteDomain, int, error)
	FindAll(ctx context.Context, param *NoteParameter) ([]NoteDomain, error)
	FindOne(ctx context.Context, param *NoteParameter) (NoteDomain, error)
	Add(ctx context.Context, data *NoteDomain) (int, error)
	Edit(ctx context.Context, data *NoteDomain) error
	Delete(ctx context.Context, data *NoteDomain) error
}

type NoteUsecase interface {
	FindAllPagination(ctx context.Context, param *NoteParameter) ([]NoteDomain, businesses.Pagination, error)
	FindAll(ctx context.Context, param *NoteParameter) ([]NoteDomain, error)
	FindOne(ctx context.Context, param *NoteParameter) (NoteDomain, error)
	Add(ctx context.Context, data *NoteDomain) (int, error)
	AddWithImageBytes(ctx context.Context, title string, imageBytes []byte) (int, error)
	Edit(ctx context.Context, data *NoteDomain) error
	Delete(ctx context.Context, id int) error
}
