package note

import "context"

type Entity struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	Note       string  `json:"note"`
	ReminderAt string  `json:"reminder_at"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
}

type Parameter struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	Note       string  `json:"note"`
	ReminderAt string  `json:"reminder_at"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
	OrderBy    string  `json:"order_by"`
	Sort       string  `json:"sort"`
	Offset     int     `json:"offset"`
	Limit      int     `json:"limit"`
}

type Repository interface {
	FindAll(ctx context.Context, parameter Parameter) (data []Entity, count int, err error)
	SelectAll(ctx context.Context, parameter Parameter) (data []Entity, err error)
	FindByID(ctx context.Context, parameter Parameter) (data Entity, err error)
	FindByTitle(ctx context.Context, parameter Parameter) (data Entity, err error)
	FindByTitleOrNote(ctx context.Context, parameter Parameter) (data Entity, err error)
	Create(ctx context.Context, data *Entity) (lastInsertId int, err error)
	Update(ctx context.Context, data *Entity) (lastUpdateId int, err error)
	Delete(ctx context.Context, id int) (lastUpdateId int, err error)
}

type Usecase interface {
	FindAll(ctx context.Context, parameter Parameter) (data []Entity, count int, err error)
	SelectAll(ctx context.Context, parameter Parameter) (data []Entity, err error)
	FindByID(ctx context.Context, parameter Parameter) (data Entity, err error)
	FindByTitle(ctx context.Context, parameter Parameter) (data Entity, err error)
	FindByTitleOrNote(ctx context.Context, parameter Parameter) (data Entity, err error)
	Create(ctx context.Context, data *Entity) (lastInsertId int, err error)
	Update(ctx context.Context, data *Entity) (lastUpdateId int, err error)
	Delete(ctx context.Context, id int) (lastUpdateId int, err error)
}
