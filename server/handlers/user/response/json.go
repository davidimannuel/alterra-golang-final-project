package response

import "keep-remind-app/businesses/user"

type JSON struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

func (res *JSON) FromDomain(domain *user.Domain) JSON {
	return JSON{
		Id:          domain.Id,
		Username:    domain.Username,
		CountryCode: domain.CountryCode,
		Phone:       domain.Phone,
		Email:       domain.Email,
	}
}
