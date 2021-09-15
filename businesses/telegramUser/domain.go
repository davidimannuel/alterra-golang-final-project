package telegramuser

import (
	"keep-remind-app/businesses"
	"time"
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
	businesses.BaseParameter
}
