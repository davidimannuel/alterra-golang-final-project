package note

import (
	"keep-remind-app/businesses/note"
	"time"
)

type Note struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Note       string     `json:"note"`
	ReminderAt *time.Time `json:"reminder_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func FromDomain(domain *note.Domain) Note {
	return Note{
		ID:         domain.ID,
		Title:      domain.Title,
		Note:       domain.Note,
		ReminderAt: domain.ReminderAt,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomains(domains []note.Domain) (res []Note) {
	for i := range domains {
		res = append(res, FromDomain(&domains[i]))
	}
	return res
}
