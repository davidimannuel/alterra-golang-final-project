package user

import (
	"keep-remind-app/businesses/user"
	"keep-remind-app/drivers/repositories/note"

	"gorm.io/gorm"
)

var tableName string = "users"

type Model struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Notes    []note.Model `gorm:"foreignKey:UserID"`
}

func (model *Model) TableName() string {
	return tableName
}

func fromDomain(domain *user.Domain) *Model {
	return &Model{
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

func (model *Model) toDomain() user.Domain {
	return user.Domain{
		ID:        int(model.ID),
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
