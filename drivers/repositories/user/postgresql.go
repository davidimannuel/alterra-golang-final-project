package user

import (
	"context"
	"keep-remind-app/businesses/user"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) user.Repository {
	return &userRepository{
		DB: db,
	}
}

func (repo *userRepository) Add(ctx context.Context, data *user.Domain) (res user.Domain, err error) {
	model := fromDomain(data)
	result := repo.DB.Create(&model)
	if result.Error != nil {
		return res, result.Error
	}
	return toDomain(model), err
}
