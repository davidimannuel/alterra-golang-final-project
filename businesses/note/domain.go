package note

import (
	"context"
	"time"
)

type NoteDomain struct {
	ID         int
	UserID     int
	Title      string
	Note       string
	ReminderAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type NoteParameter struct {
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

type NoteRepository interface {
	FindAllPagination(ctx context.Context, param *NoteParameter) ([]NoteDomain, int, error)
	FindAll(ctx context.Context, param *NoteParameter) ([]NoteDomain, error)
	FindOne(ctx context.Context, param *NoteParameter) (NoteDomain, error)
	Add(ctx context.Context, data *NoteDomain) (int, error)
	Edit(ctx context.Context, data *NoteDomain) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type NoteUsecase interface {
	// FindAllPagination(ctx context.Context, param *NoteParameter) ([]NoteDomain, int, error)
	FindAll(ctx context.Context, param *NoteParameter) ([]NoteDomain, error)
	FindByID(ctx context.Context, param *NoteParameter) (NoteDomain, error)
	// FindByTitle(ctx context.Context, param *NoteParameter) (NoteDomain, error)
	// FindByTitleOrNote(ctx context.Context, param *NoteParameter) (NoteDomain, error)
	Add(ctx context.Context, data *NoteDomain) (NoteDomain, error)
	AddWithImageBytes(ctx context.Context, title string, imageBytes []byte) (NoteDomain, error)
	// Edit(ctx context.Context, data *NoteDomain) (NoteDomain, error)
	// Delete(ctx context.Context, id int) (NoteDomain, error)
}
