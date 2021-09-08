package user

import (
	"keep-remind-app/businesses/user"

	"gorm.io/gorm"
)

var tableName string = "users"

type Model struct {
	gorm.Model
	Username    string
	CountryCode string
	Phone       string
	Email       string
	Password    string
}

func (model *Model) TableName() string {
	return tableName
}

func fromDomain(domain *user.Domain) *Model {
	return &Model{
		Model: gorm.Model{
			ID:        uint(domain.Id),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Username:    domain.Username,
		CountryCode: domain.CountryCode,
		Phone:       domain.Phone,
		Email:       domain.Email,
		Password:    domain.Password,
	}
}

func toDomain(model *Model) user.Domain {
	return user.Domain{
		Id:          int(model.ID),
		Username:    model.Username,
		CountryCode: model.CountryCode,
		Phone:       model.Phone,
		Email:       model.Email,
		Password:    model.Password,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
