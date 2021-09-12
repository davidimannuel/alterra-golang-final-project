package note

import (
	"keep-remind-app/businesses/note"
	"time"

	"gorm.io/gorm"
)

var tableName string = "notes"

type Model struct {
	gorm.Model
	UserId     uint
	Title      string
	Note       string
	ReminderAt *time.Time
}

func (model *Model) TableName() string {
	return tableName
}

func fromDomain(domain *note.Domain) *Model {
	return &Model{
		Model: gorm.Model{
			ID:        uint(domain.Id),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		UserId:     uint(domain.UserId),
		Title:      domain.Note,
		Note:       domain.Note,
		ReminderAt: domain.ReminderAt,
	}
}

func (model *Model) toDomain() note.Domain {
	return note.Domain{
		Id:         int(model.ID),
		UserId:     int(model.UserId),
		Title:      model.Title,
		Note:       model.Note,
		ReminderAt: model.ReminderAt,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
