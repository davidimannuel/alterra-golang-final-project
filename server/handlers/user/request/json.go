package request

import "keep-remind-app/businesses/user"

type JSON struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (req *JSON) ToDomain() *user.Domain {
	return &user.Domain{
		Username:    req.Username,
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
		Email:       req.Email,
		Password:    req.Password,
	}
}
