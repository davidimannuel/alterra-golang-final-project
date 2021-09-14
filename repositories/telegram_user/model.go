package telegram_user

import (
	"gorm.io/gorm"
)

var tableName string = "telegram_users"

type TelegramUserModel struct {
	gorm.Model
	UserID   uint
	Username string `gorm:"column:username"`
	IsActive bool
}

func (model *TelegramUserModel) TableName() string {
	return tableName
}
