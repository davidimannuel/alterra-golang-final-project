package note

import (
	"keep-remind-app/businesses/note"
	"keep-remind-app/repositories/label"
	"time"

	"gorm.io/gorm"
)

var tableName string = "notes"

type NoteModel struct {
	gorm.Model
	UserID     uint
	Title      string
	Note       string
	ReminderAt *time.Time
	Labels     []*label.LabelModel `gorm:"many2many:note_labels;"`
}

func (model *NoteModel) TableName() string {
	return tableName
}

func fromDomain(domain *note.NoteDomain) *NoteModel {
	return &NoteModel{
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

func (model *NoteModel) toDomain() note.NoteDomain {
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

func toDomains(models []NoteModel) (res []note.NoteDomain) {
	for i := range models {
		res = append(res, models[i].toDomain())
	}
	return res
}
