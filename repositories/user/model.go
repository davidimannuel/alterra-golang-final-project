package user

import (
	"keep-remind-app/businesses/user"
	"keep-remind-app/repositories/label"
	"keep-remind-app/repositories/note"
	"keep-remind-app/repositories/telegram_user"

	"gorm.io/gorm"
)

var tableName string = "users"

type UserModel struct {
	gorm.Model
	Name          string
	Email         string
	Password      string
	Notes         []note.NoteModel                  `gorm:"foreignKey:UserID"`
	Labels        []label.LabelModel                `gorm:"foreignKey:UserID"`
	TelegramUsers []telegram_user.TelegramUserModel `gorm:"foreignKey:UserID"`
}

func (model *UserModel) TableName() string {
	return tableName
}

func fromDomain(domain *user.UserDomain) *UserModel {
	return &UserModel{
		Model: gorm.Model{
			ID:        uint(domain.ID),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}

func (model *UserModel) toDomain() user.UserDomain {
	return user.UserDomain{
		ID:        int(model.ID),
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
