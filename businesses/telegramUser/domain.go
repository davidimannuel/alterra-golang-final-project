package telegramUser

import (
	"context"
	"keep-remind-app/businesses"
	"time"
)

var (
	redisKeyActivatedTelegram = "activatedTelegram-"
)

type TelegramUserDomain struct {
	ID        int
	UserID    uint
	Username  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TelegramUserParameter struct {
	UserID   uint
	Username string
	Status   string
	businesses.BaseParameter
}

type TelegramUserRepository interface {
	FindAllPagination(ctx context.Context, param *TelegramUserParameter) ([]TelegramUserDomain, int, error)
	FindAll(ctx context.Context, param *TelegramUserParameter) ([]TelegramUserDomain, error)
	FindOne(ctx context.Context, param *TelegramUserParameter) (TelegramUserDomain, error)
	Add(ctx context.Context, data *TelegramUserDomain) (int, error)
	EditStatus(ctx context.Context, data *TelegramUserDomain) error
	Delete(ctx context.Context, data *TelegramUserDomain) error
}

type TelegramUserUsecase interface {
	FindAllPagination(ctx context.Context, param *TelegramUserParameter) ([]TelegramUserDomain, businesses.Pagination, error)
	FindAll(ctx context.Context, param *TelegramUserParameter) ([]TelegramUserDomain, error)
	FindOne(ctx context.Context, param *TelegramUserParameter) (TelegramUserDomain, error)
	Add(ctx context.Context, data *TelegramUserDomain) (int, error)
	Activated(ctx context.Context, data *TelegramUserDomain) error
	GenerateActivatedOTP(ctx context.Context, id int) (otp string, err error)
	Delete(ctx context.Context, data *TelegramUserDomain) error
}
