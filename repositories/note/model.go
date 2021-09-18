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
	Labels     []label.LabelModel `gorm:"many2many:note_labels;"`
}

func (model *NoteModel) TableName() string {
	return tableName
}

func fromDomain(domain *note.NoteDomain) *NoteModel {
	domainLabels := fromLabelsDomain(domain.Labels)
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
		Labels:     domainLabels,
	}
}

func fromLabelDomain(labelDomain note.LabelDomain) *label.LabelModel {
	return &label.LabelModel{
		Model: gorm.Model{
			ID: uint(labelDomain.ID),
		},
		UserID: uint(labelDomain.UserID),
		Name:   labelDomain.Name,
	}
}

func fromLabelsDomain(labelsDomain []note.LabelDomain) (res []label.LabelModel) {
	for i := range labelsDomain {
		res = append(res, *fromLabelDomain(labelsDomain[i]))
	}
	return res
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
		Labels:     toLabelsDomain(model.Labels),
	}

}

func toDomains(models []NoteModel) (res []note.NoteDomain) {
	for i := range models {
		res = append(res, models[i].toDomain())
	}
	return res
}

func toLabelsDomain(models []label.LabelModel) (domains []note.LabelDomain) {
	for i := range models {
		domains = append(domains, note.LabelDomain{
			ID:     int(models[i].ID),
			UserID: int(models[i].UserID),
			Name:   models[i].Name,
		})
	}
	return domains
}
