package user

import (
	"context"
	"keep-remind-app/businesses/user"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (repo *userRepository) Add(ctx context.Context, data *user.UserDomain) (res user.UserDomain, err error) {
	model := fromDomain(data)
	result := repo.DB.Create(&model)
	if result.Error != nil {
		return res, result.Error
	}
	return model.toDomain(), err
}

func (repo *userRepository) FindByEmail(ctx context.Context, param *user.UserParameter) (res user.UserDomain, err error) {
	model := Model{}
	err = repo.DB.Where("email = ?", param.Email).First(&model).Error
	if err != nil {
		return res, err
	}
	return model.toDomain(), nil
}

func (repo *userRepository) FindByID(ctx context.Context, param *user.UserParameter) (res user.UserDomain, err error) {
	model := Model{}
	err = repo.DB.Where("id = ?", param.ID).First(&model).Error
	if err != nil {
		return res, err
	}
	return model.toDomain(), nil
}
