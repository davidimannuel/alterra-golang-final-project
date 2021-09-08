package note

import (
	"keep-remind-app/businesses/note"
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title      string
	Note       string
	ReminderAt *time.Time
}

func fromDomain(domain *note.Domain) *Note {
	return &Note{
		Model: gorm.Model{
			ID:        uint(domain.Id),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Title:      domain.Note,
		Note:       domain.Note,
		ReminderAt: domain.ReminderAt,
	}
}

func toDomain(model *Note) note.Domain {
	return note.Domain{
		Id:         int(model.ID),
		Title:      model.Title,
		Note:       model.Note,
		ReminderAt: model.ReminderAt,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
