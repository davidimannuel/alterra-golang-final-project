package telegramUser

import (
	"context"
	"keep-remind-app/businesses/telegramUser"
	telegramuser "keep-remind-app/businesses/telegramUser"
	"strconv"

	"gorm.io/gorm"
)

type telegramUserRepository struct {
	DB *gorm.DB
}

func NewTelegramUserRepository(db *gorm.DB) telegramUser.TelegramUserRepository {
	return &telegramUserRepository{
		DB: db,
	}
}

func (repo *telegramUserRepository) buildParameter(ctx context.Context, param *telegramuser.TelegramUserParameter) (query *gorm.DB) {
	query = repo.DB
	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.UserID != 0 {
		query = query.Where("user_id = ?", param.UserID)
	}
	if param.Username != "" {
		query = query.Where("username = ?", param.Username)
	}
	if param.Status != "" {
		status, _ := strconv.ParseBool(param.Status)
		query = query.Where("status = ?", status)
	}
	return query
}

func (repo telegramUserRepository) FindAllPagination(ctx context.Context, param *telegramuser.TelegramUserParameter) (res []telegramuser.TelegramUserDomain, count int, err error) {
	query := repo.buildParameter(ctx, param)
	models := []TelegramUserModel{}
	if err = query.Offset(param.GetOffset()).Limit(param.PerPage).Find(&models).Error; err != nil {
		return res, count, err
	}
	var totalData int64
	if err = query.Count(&totalData).Error; err != nil {
		return res, count, err
	}
	return toDomains(models), int(totalData), err
}

func (repo telegramUserRepository) FindAll(ctx context.Context, param *telegramuser.TelegramUserParameter) (res []telegramuser.TelegramUserDomain, err error) {
	query := repo.buildParameter(ctx, param)
	models := []TelegramUserModel{}
	if err = query.Offset(param.GetOffset()).Limit(param.PerPage).Find(&models).Error; err != nil {
		return res, err
	}
	return toDomains(models), err
}

func (repo telegramUserRepository) FindOne(ctx context.Context, param *telegramuser.TelegramUserParameter) (res telegramuser.TelegramUserDomain, err error) {
	query := repo.buildParameter(ctx, param)
	model := TelegramUserModel{}
	if err = query.First(&model).Error; err != nil {
		return res, err
	}
	return model.toDomain(), err
}

func (repo telegramUserRepository) Add(ctx context.Context, data *telegramuser.TelegramUserDomain) (res int, err error) {
	model := fromDomain(data)
	if err = repo.DB.Create(&model).Error; err != nil {
		return 0, err
	}
	return int(model.ID), err
}

func (repo telegramUserRepository) EditStatus(ctx context.Context, data *telegramuser.TelegramUserDomain) (err error) {
	if err = repo.DB.Model(&TelegramUserModel{}).Where("id = ?", data.ID).Update("is_active", data.IsActive).Error; err != nil {
		return err
	}
	return err
}

func (repo telegramUserRepository) Delete(ctx context.Context, data *telegramuser.TelegramUserDomain) (err error) {
	model := fromDomain(data)
	if err = repo.DB.Delete(&model).Error; err != nil {
		return err
	}
	return err
}
