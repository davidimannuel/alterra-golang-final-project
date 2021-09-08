package request

import (
	"keep-remind-app/businesses/note"
	"keep-remind-app/helpers/str"
	"time"
)

type JSON struct {
	Title      string  `json:"title"`
	Note       string  `json:"note"`
	ReminderAt *string `json:"reminder_at"`
}

func (req *JSON) ToDomain() *note.Domain {
	return &note.Domain{
		Title:      req.Title,
		Note:       req.Note,
		ReminderAt: str.ToNilTime(req.ReminderAt, time.RFC3339),
	}
}
