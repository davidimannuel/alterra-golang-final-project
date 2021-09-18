package telegramUser

import "errors"

var (
	ErrInvalidUser   = errors.New("Invalid User")
	ErrUsernameExist = errors.New("Username Exist")
	ErrDataNotFound  = errors.New("Data Not Found")
)
