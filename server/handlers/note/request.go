package note

import (
	"keep-remind-app/businesses/note"
	"keep-remind-app/helpers/str"
	"time"
)

type AddNoteRequest struct {
	Title      string  `json:"title"`
	Note       string  `json:"note"`
	ReminderAt *string `json:"reminder_at"`
}

func (req *AddNoteRequest) ToDomain() *note.NoteDomain {
	return &note.NoteDomain{
		Title:      req.Title,
		Note:       req.Note,
		ReminderAt: str.ToNilTime(req.ReminderAt, time.RFC3339),
	}
}
