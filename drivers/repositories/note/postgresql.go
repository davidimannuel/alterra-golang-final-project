package note

import (
	"context"
	"keep-remind-app/businesses/note"

	"gorm.io/gorm"
)

type noteRepository struct {
	DB *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) note.Repository {
	return &noteRepository{
		DB: db,
	}
}

func (repo *noteRepository) FindAllPagination(ctx context.Context, parameter note.Parameter) (res []note.Domain, total int, err error) {
	notes := []Model{}
	// offset := (page - 1) * perpage
	err = repo.DB.Where("user_id = ?", ctx.Value("user_id").(int)).Find(&notes).Error
	if err != nil {
		return res, total, err
	}
	var totalData int64
	err = repo.DB.Count(&totalData).Error
	if err != nil {
		return res, total, err
	}
	return toDomains(notes), int(totalData), err
}
func (repo *noteRepository) FindAll(ctx context.Context, parameter note.Parameter) (res []note.Domain, err error) {
	notes := []Model{}
	err = repo.DB.Where("user_id = ?", ctx.Value("user_id").(int)).Find(&notes).Error
	if err != nil {
		return res, err
	}
	return toDomains(notes), err
}
func (repo *noteRepository) FindOne(ctx context.Context, parameter note.Parameter) (note.Domain, error) {
	panic("implement")
}

func (repo *noteRepository) Add(ctx context.Context, data *note.Domain) (res int, err error) {
	model := fromDomain(data)
	result := repo.DB.Create(&model)
	if result.Error != nil {
		return res, result.Error
	}
	return int(model.ID), err
}

func (repo *noteRepository) Edit(ctx context.Context, data *note.Domain) (res int, err error) {
	panic("implement")
}
func (repo *noteRepository) Delete(ctx context.Context, id int) (res int, err error) {
	panic("implement")
}
