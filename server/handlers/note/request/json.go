package request

import (
	"alterra-golang-final-project/businesses/note"
	"alterra-golang-final-project/helpers/str"
	"time"
)

type Note struct {
	Title      string  `json:"title"`
	Note       string  `json:"note"`
	ReminderAt *string `json:"reminder_at"`
}

func (req *Note) ToDomain() *note.Domain {
	return &note.Domain{
		Title:      req.Title,
		Note:       req.Note,
		ReminderAt: str.ToNilTime(req.ReminderAt, time.RFC3339),
	}
}
