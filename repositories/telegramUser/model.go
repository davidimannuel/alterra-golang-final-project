package telegramUser

import (
	telegramuser "keep-remind-app/businesses/telegramUser"

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

func fromDomain(domain *telegramuser.TelegramUserDomain) *TelegramUserModel {
	return &TelegramUserModel{
		Model: gorm.Model{
			ID:        uint(domain.ID),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		UserID:   uint(domain.UserID),
		Username: domain.Username,
	}
}

func (model *TelegramUserModel) toDomain() telegramuser.TelegramUserDomain {
	return telegramuser.TelegramUserDomain{
		ID:        int(model.ID),
		UserID:    model.UserID,
		Username:  model.Username,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func toDomains(models []TelegramUserModel) (domains []telegramuser.TelegramUserDomain) {
	for i := range models {
		domains = append(domains, models[i].toDomain())
	}
	return domains
}
