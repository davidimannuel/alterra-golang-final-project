package request

import "keep-remind-app/businesses/user"

type Add struct {
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *Add) ToDomain() *user.Domain {
	return &user.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
