package note

import (
	"keep-remind-app/businesses/note"
	"time"
)

type NoteResponse struct {
	ID         int             `json:"id"`
	Title      string          `json:"title"`
	Note       string          `json:"note"`
	ReminderAt *time.Time      `json:"reminder_at"`
	Labels     []LabelResponse `json:"labels"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

type LabelResponse struct {
	ID     int
	UserID int
	Name   string
}

func FromDomain(domain *note.NoteDomain) NoteResponse {
	return NoteResponse{
		ID:         domain.ID,
		Title:      domain.Title,
		Note:       domain.Note,
		ReminderAt: domain.ReminderAt,
		Labels:     FromLabelsDomais(domain.Labels),
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomains(domains []note.NoteDomain) (res []NoteResponse) {
	for i := range domains {
		res = append(res, FromDomain(&domains[i]))
	}
	return res
}

func FromLabelsDomais(domains []note.LabelDomain) (res []LabelResponse) {
	for i := range domains {
		res = append(res, LabelResponse{
			ID:     domains[i].ID,
			UserID: domains[i].UserID,
			Name:   domains[i].Name,
		})
	}
	return res
}
