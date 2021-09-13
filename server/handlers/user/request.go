package user

import "keep-remind-app/businesses/user"

type UserAdd struct {
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *UserAdd) ToDomain() *user.Domain {
	return &user.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
