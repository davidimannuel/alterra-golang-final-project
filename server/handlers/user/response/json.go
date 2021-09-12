package response

import "keep-remind-app/businesses/user"

type JSON struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (res *JSON) FromDomain(domain *user.Domain) JSON {
	return JSON{
		Id:    domain.Id,
		Name:  domain.Name,
		Email: domain.Email,
	}
}
