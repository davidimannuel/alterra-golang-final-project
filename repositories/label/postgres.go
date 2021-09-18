package label

import (
	"context"
	"keep-remind-app/businesses/label"

	"gorm.io/gorm"
)

type labelRepository struct {
	DB *gorm.DB
}

func NewLabelRepository(db *gorm.DB) label.LabelRepository {
	return &labelRepository{
		DB: db,
	}
}

func (repo *labelRepository) buildParameter(ctx context.Context, param *label.LabelParameter) (query *gorm.DB) {
	query = repo.DB
	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.UserID != 0 {
		query = query.Where("user_id = ?", param.UserID)
	}
	if param.Name != "" {
		query = query.Where("name = ?", param.Name)
	}
	return query
}

func (repo *labelRepository) FindOne(ctx context.Context, param *label.LabelParameter) (res label.LabelDomain, err error) {
	query := repo.buildParameter(ctx, param)
	model := LabelModel{}
	if err = query.First(&model).Error; err != nil {
		return res, err
	}
	return model.toDomain(), err
}

func (repo *labelRepository) FindAllPagination(ctx context.Context, param *label.LabelParameter) (res []label.LabelDomain, count int, err error) {
	query := repo.buildParameter(ctx, param)
	models := []LabelModel{}
	if err = query.Offset(param.GetOffset()).Limit(param.GetPerPage()).Find(&models).Error; err != nil {
		return res, count, err
	}
	var totalData int64
	if err = query.Model(&LabelModel{}).Count(&totalData).Error; err != nil {
		return res, count, err
	}
	return toDomains(models), int(totalData), err
}

func (repo *labelRepository) FindAll(ctx context.Context, param *label.LabelParameter) (res []label.LabelDomain, err error) {
	query := repo.buildParameter(ctx, param)
	models := []LabelModel{}
	if err = query.Offset(param.GetOffset()).Limit(param.PerPage).Find(&models).Error; err != nil {
		return res, err
	}
	return toDomains(models), err
}

func (repo *labelRepository) Add(ctx context.Context, data *label.LabelDomain) (res int, err error) {
	model := fromDomain(data)
	if err = repo.DB.Create(&model).Error; err != nil {
		return 0, err
	}
	return int(model.ID), err
}

func (repo *labelRepository) Edit(ctx context.Context, data *label.LabelDomain) (err error) {
	model := fromDomain(data)
	if err = repo.DB.Model(&model).Updates(LabelModel{Name: model.Name}).Error; err != nil {
		return err
	}
	return err
}

func (repo *labelRepository) Delete(ctx context.Context, id int) (err error) {
	if err = repo.DB.Delete(&LabelModel{}, id).Error; err != nil {
		return err
	}
	return err
}
