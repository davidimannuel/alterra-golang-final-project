package user

import "keep-remind-app/businesses/user"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromDomain(domain *user.Domain) User {
	return User{
		ID:    domain.ID,
		Name:  domain.Name,
		Email: domain.Email,
	}
}
