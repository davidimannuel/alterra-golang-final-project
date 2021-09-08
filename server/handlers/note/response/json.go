package response

import (
	"keep-remind-app/businesses/note"
	"time"
)

type JSON struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Note       string     `json:"note"`
	ReminderAt *time.Time `json:"reminder_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (res *JSON) FromDomain(domain *note.Domain) JSON {
	return JSON{
		Id:         domain.Id,
		Title:      domain.Title,
		Note:       domain.Note,
		ReminderAt: domain.ReminderAt,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
