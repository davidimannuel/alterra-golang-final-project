package label

import (
	"keep-remind-app/businesses/label"
	"time"
)

type LabelResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *label.LabelDomain) LabelResponse {
	return LabelResponse{
		ID:        domain.ID,
		UserID:    domain.UserID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomains(domains []label.LabelDomain) (res []LabelResponse) {
	for i := range domains {
		res = append(res, FromDomain(&domains[i]))
	}
	return res
}
