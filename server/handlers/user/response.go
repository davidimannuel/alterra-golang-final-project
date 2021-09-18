package user

import "keep-remind-app/businesses/user"

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromDomain(domain *user.UserDomain) UserResponse {
	return UserResponse{
		ID:    domain.ID,
		Name:  domain.Name,
		Email: domain.Email,
	}
}
