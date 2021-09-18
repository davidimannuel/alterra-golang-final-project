package note

import (
	"keep-remind-app/businesses/note"
	"keep-remind-app/helpers/str"
	"time"
)

type AddNoteRequest struct {
	Title      string         `json:"title"`
	Note       string         `json:"note"`
	ReminderAt *string        `json:"reminder_at"`
	Labels     []LabelRequest `json:"labels"`
}

type LabelRequest struct {
	Name string `json:"name"`
}

func (req *AddNoteRequest) ToDomain() *note.NoteDomain {
	return &note.NoteDomain{
		Title:      req.Title,
		Note:       req.Note,
		ReminderAt: str.ToNilTime(req.ReminderAt, time.RFC3339),
		Labels:     ToLabelsDomain(req.Labels),
	}
}

func (req *LabelRequest) ToDomain() *note.LabelDomain {
	return &note.LabelDomain{
		Name: req.Name,
	}
}

func ToLabelsDomain(labelsReq []LabelRequest) (res []note.LabelDomain) {
	for i := range labelsReq {
		res = append(res, *labelsReq[i].ToDomain())
	}
	return res
}
