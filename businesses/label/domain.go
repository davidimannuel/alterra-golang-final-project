package label

import (
	"context"
	"keep-remind-app/businesses"
	"time"
)

type LabelDomain struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LabelParameter struct {
	UserID int
	Name   string
	businesses.BaseParameter
}

type LabelRepository interface {
	FindAllPagination(ctx context.Context, param *LabelParameter) ([]LabelDomain, int, error)
	FindAll(ctx context.Context, param *LabelParameter) ([]LabelDomain, error)
	FindOne(ctx context.Context, param *LabelParameter) (LabelDomain, error)
	Add(ctx context.Context, data *LabelDomain) (int, error)
	Edit(ctx context.Context, data *LabelDomain) error
	Delete(ctx context.Context, id int) error
}

type LabelUsecase interface {
	FindAllPagination(ctx context.Context, param *LabelParameter) ([]LabelDomain, businesses.Pagination, error)
	FindAll(ctx context.Context, param *LabelParameter) ([]LabelDomain, error)
	FindOne(ctx context.Context, param *LabelParameter) (LabelDomain, error)
	Add(ctx context.Context, data *LabelDomain) (int, error)
	Edit(ctx context.Context, data *LabelDomain) error
	Delete(ctx context.Context, id int) error
}
