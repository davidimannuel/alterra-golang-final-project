package telegramuser

import "keep-remind-app/businesses/telegramUser"

type TelegramUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func fromDomain(domain telegramUser.TelegramUserDomain) TelegramUserResponse {
	return TelegramUserResponse{
		ID:       domain.ID,
		Username: domain.Username,
		IsActive: domain.IsActive,
	}
}

func fromDomains(domains []telegramUser.TelegramUserDomain) (res []TelegramUserResponse) {
	for i := range domains {
		res = append(res, fromDomain(domains[i]))
	}
	return
}
