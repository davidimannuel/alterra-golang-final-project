package note

import (
	"keep-remind-app/businesses/note"
	"time"

	"gorm.io/gorm"
)

var tableName string = "notes"

type Model struct {
	gorm.Model
	UserID     uint
	Title      string
	Note       string
	ReminderAt *time.Time
}

func (model *Model) TableName() string {
	return tableName
}

func fromDomain(domain *note.NoteDomain) *Model {
	return &Model{
		Model: gorm.Model{
			ID:        uint(domain.ID),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		UserID:     uint(domain.UserID),
		Title:      domain.Title,
		Note:       domain.Note,
		ReminderAt: domain.ReminderAt,
	}
}

func (model *Model) toDomain() note.NoteDomain {
	return note.NoteDomain{
		ID:         int(model.ID),
		UserID:     int(model.UserID),
		Title:      model.Title,
		Note:       model.Note,
		ReminderAt: model.ReminderAt,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}

func toDomains(models []Model) (res []note.NoteDomain) {
	for i := range models {
		res = append(res, models[i].toDomain())
	}
	return res
}
