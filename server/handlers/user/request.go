package user

import "keep-remind-app/businesses/user"

type CreateUserRequest struct {
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *CreateUserRequest) ToDomain() *user.UserDomain {
	return &user.UserDomain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
