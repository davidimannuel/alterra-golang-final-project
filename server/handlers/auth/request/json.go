package request

import "keep-remind-app/businesses/auth"

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *Register) ToDomain() *auth.Domain {
	return &auth.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *Login) ToDomain() *auth.Domain {
	return &auth.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
