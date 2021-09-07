package response

import (
	"alterra-golang-final-project/businesses/note"
	"time"
)

type Note struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Note       string     `json:"note"`
	ReminderAt *time.Time `json:"reminder_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (res *Note) FromDomain(domain *note.Domain) Note {
	return Note{
		Id:         domain.Id,
		Title:      domain.Title,
		Note:       domain.Note,
		ReminderAt: domain.ReminderAt,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
