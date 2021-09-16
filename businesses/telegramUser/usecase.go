package telegramUser

import (
	"context"
	"errors"
	"keep-remind-app/businesses"
	redisDomain "keep-remind-app/businesses/redis"
	"keep-remind-app/helpers/str"
	"strconv"
	"time"
)

var (
	RedisKeyRegisteredTelegram = "registeredTelegram-"
	LastChatTelegram           = "lastChatTelegram-"
	LastActionTelegram         = "lastActionTelegram-"
	RedisKeyActivatedTelegram  = "activatedTelegram-"
	RedisKeyTelegramUser       = "telegramUser-"

	LatestActionOtp = "otp"
)

type telegramUserUsecase struct {
	telegramUserRepository TelegramUserRepository
	redisUsecase           redisDomain.RedisUsecase
}

func NewTelegramUserUsecase(telegramUserRepository TelegramUserRepository, redisUsecase redisDomain.RedisUsecase) TelegramUserUsecase {
	return &telegramUserUsecase{
		telegramUserRepository: telegramUserRepository,
		redisUsecase:           redisUsecase,
	}
}

func (uc *telegramUserUsecase) validate(ctx context.Context, data *TelegramUserDomain) (err error) {
	userID := uint(ctx.Value("user_id").(int))
	exist, _ := uc.telegramUserRepository.FindOne(ctx, &TelegramUserParameter{UserID: userID})
	if exist.ID != 0 {
		return errors.New("User only have 1 account")
	}
	data.UserID = userID
	exist, _ = uc.telegramUserRepository.FindOne(ctx, &TelegramUserParameter{Username: data.Username})
	if exist.ID != 0 {
		return ErrUsernameExist
	}

	return
}

func (uc *telegramUserUsecase) FindAllPagination(ctx context.Context, param *TelegramUserParameter) (res []TelegramUserDomain, p businesses.Pagination, err error) {
	res, count, err := uc.telegramUserRepository.FindAllPagination(ctx, param)
	if err != nil {
		return
	}
	return res, param.GetPageInfo(count), err
}

func (uc *telegramUserUsecase) FindAll(ctx context.Context, param *TelegramUserParameter) (res []TelegramUserDomain, err error) {
	res, err = uc.telegramUserRepository.FindAll(ctx, param)
	if err != nil {
		return
	}
	return
}

func (uc *telegramUserUsecase) FindOne(ctx context.Context, param *TelegramUserParameter) (res TelegramUserDomain, err error) {
	res, err = uc.telegramUserRepository.FindOne(ctx, param)
	if err != nil {
		return
	}
	return
}

func (uc *telegramUserUsecase) Add(ctx context.Context, data *TelegramUserDomain) (res int, err error) {
	err = uc.validate(ctx, data)
	if err != nil {
		return
	}
	res, err = uc.telegramUserRepository.Add(ctx, data)
	if err != nil {
		return
	}
	return
}

func (uc *telegramUserUsecase) GenerateActivatedOTP(ctx context.Context, id int) (otp string, err error) {
	param := &TelegramUserParameter{}
	param.ID = id
	exist, err := uc.telegramUserRepository.FindOne(ctx, &TelegramUserParameter{})
	if err != nil {
		return "", ErrDataNotFound
	}
	otp = str.RandomNumberString(6)
	key := RedisKeyActivatedTelegram + exist.Username + otp
	err = uc.redisUsecase.Set(ctx, key, exist.ID, time.Minute*5)
	if err != nil {
		return "", businesses.ErrInternalServer
	}
	err = uc.redisUsecase.Set(ctx, LastActionTelegram+exist.Username, LatestActionOtp, 0)
	return otp, err
}

func (uc *telegramUserUsecase) Activated(ctx context.Context, username string, otp string) error {
	res, err := uc.redisUsecase.Get(ctx, RedisKeyActivatedTelegram+username+otp)
	if err != nil {
		return ErrDataNotFound
	}
	id, _ := strconv.Atoi(res)
	param := new(TelegramUserParameter)
	param.Username = username
	param.ID = id
	exist, err := uc.telegramUserRepository.FindOne(ctx, param)
	if err != nil {
		return businesses.ErrInternalServer
	}
	// set userID telegram
	err = uc.redisUsecase.Set(ctx, RedisKeyTelegramUser+username, exist.UserID, 0)
	if err != nil {
		return businesses.ErrInternalServer
	}
	err = uc.redisUsecase.Del(ctx, RedisKeyActivatedTelegram+username+otp)
	if err != nil {
		return businesses.ErrInternalServer
	}
	return uc.redisUsecase.Del(ctx, LastActionTelegram+username)
}

func (uc *telegramUserUsecase) Delete(ctx context.Context, data *TelegramUserDomain) error {
	return uc.telegramUserRepository.Delete(ctx, data)
}
