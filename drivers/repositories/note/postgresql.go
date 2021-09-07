package note

import (
	"alterra-golang-final-project/businesses/note"
	"context"

	"gorm.io/gorm"
)

type noteRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) note.Repository {
	return &noteRepository{
		DB: db,
	}
}

func (repo *noteRepository) Add(ctx context.Context, data *note.Domain) (res note.Domain, err error) {
	model := fromDomain(data)
	result := repo.DB.Create(&model)
	if result.Error != nil {
		return res, result.Error
	}
	return toDomain(model), err
}
