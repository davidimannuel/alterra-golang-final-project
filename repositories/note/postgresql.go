package note

import (
	"context"
	"keep-remind-app/businesses/note"

	"gorm.io/gorm"
)

type noteRepository struct {
	DB *gorm.DB
}

func NewNoteRepository(db *gorm.DB) note.NoteRepository {
	return &noteRepository{
		DB: db,
	}
}

func (repo *noteRepository) buildParameter(ctx context.Context, param *note.NoteParameter) (query *gorm.DB) {
	query = repo.DB
	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.UserID != 0 {
		query = query.Where("user_id = ?", param.UserID)
	}
	if param.Title != "" {
		query = query.Where("name LIKE ?", param.LikeChar(param.Title))
	}
	if param.Note != "" {
		query = query.Where("note LIKE ?", param.LikeChar(param.Note))
	}
	return query
}

func (repo *noteRepository) FindAllPagination(ctx context.Context, param *note.NoteParameter) (res []note.NoteDomain, total int, err error) {
	query := repo.buildParameter(ctx, param)
	models := []NoteModel{}
	if err = query.Offset(param.GetOffset()).Limit(param.GetPerPage()).Preload("Labels").Find(&models).Error; err != nil {
		return res, total, err
	}
	var totalData int64
	err = repo.DB.Model(&NoteModel{}).Count(&totalData).Error
	if err != nil {
		return res, total, err
	}
	return toDomains(models), int(totalData), err
}
func (repo *noteRepository) FindAll(ctx context.Context, param *note.NoteParameter) (res []note.NoteDomain, err error) {
	query := repo.buildParameter(ctx, param)
	models := []NoteModel{}
	if err = query.Preload("Labels").Find(&models).Error; err != nil {
		return res, err
	}
	return toDomains(models), err
}
func (repo *noteRepository) FindOne(ctx context.Context, param *note.NoteParameter) (res note.NoteDomain, err error) {
	query := repo.buildParameter(ctx, param)
	model := NoteModel{}
	if err = query.Offset(param.GetOffset()).Limit(param.PerPage).Find(&model).Error; err != nil {
		return res, err
	}
	return model.toDomain(), err
}

func (repo *noteRepository) Add(ctx context.Context, data *note.NoteDomain) (res int, err error) {
	model := fromDomain(data)
	userId := ctx.Value("user_id").(int)
	for i := range model.Labels {
		model.Labels[i].UserID = uint(userId)
	}
	if err = repo.DB.Create(&model).Error; err != nil {
		return res, err
	}
	return int(model.ID), err
}

func (repo *noteRepository) Edit(ctx context.Context, data *note.NoteDomain) (err error) {
	model := fromDomain(data)
	if err = repo.DB.Save(&model).Error; err != nil {
		return err
	}
	return err
}
func (repo *noteRepository) Delete(ctx context.Context, data *note.NoteDomain) (err error) {
	model := fromDomain(data)
	if err = repo.DB.Delete(&model).Error; err != nil {
		return err
	}
	return err
}
