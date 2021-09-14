package auth

import "keep-remind-app/businesses/auth"

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterUserRequest) ToDomain() *auth.AuthDomain {
	return &auth.AuthDomain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *Login) ToDomain() *auth.AuthDomain {
	return &auth.AuthDomain{
		Email:    req.Email,
		Password: req.Password,
	}
}
