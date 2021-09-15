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

func (repo *userRepository) buildParameter(ctx context.Context, param *user.UserParameter) (query *gorm.DB) {
	query = repo.DB
	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.Email != "" {
		query = query.Where("email = ?", param.Email)
	}
	return query
}

func (repo *userRepository) FindOne(ctx context.Context, param *user.UserParameter) (res user.UserDomain, err error) {
	query := repo.buildParameter(ctx, param)
	model := UserModel{}
	err = query.First(&model).Error
	if err != nil {
		return res, err
	}
	return model.toDomain(), nil
}

func (repo *userRepository) Add(ctx context.Context, data *user.UserDomain) (res user.UserDomain, err error) {
	model := fromDomain(data)
	if err = repo.DB.Create(&model).Error; err != nil {
		return res, err
	}
	return model.toDomain(), err
}

func (repo *userRepository) Edit(ctx context.Context, data *user.UserDomain) (err error) {
	model := fromDomain(data)
	if err = repo.DB.Save(&model).Error; err != nil {
		return err
	}
	return err
}
