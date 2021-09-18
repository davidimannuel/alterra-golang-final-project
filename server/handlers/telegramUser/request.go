package telegramuser

import "keep-remind-app/businesses/telegramUser"

type AddTelegramUserRequest struct {
	Username string `json:"username"`
}

func (req *AddTelegramUserRequest) toDomain() *telegramUser.TelegramUserDomain {
	return &telegramUser.TelegramUserDomain{
		Username: req.Username,
	}
}
