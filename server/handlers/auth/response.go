package auth

import "keep-remind-app/businesses/auth"

type JWT struct {
	Token string `json:"token"`
}

func (res *JWT) FromDomain(domain *auth.Domain) JWT {
	return JWT{
		Token: domain.JWTToken,
	}
}
