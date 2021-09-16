package telegramUser

import (
	"context"
	"fmt"
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
	if ctx.Value("user_id") == 0 {
		return ErrInvalidUser
	}
	telUser := TelegramUserDomain{}
	if data.ID != 0 {
		telUser, err = uc.telegramUserRepository.FindOne(ctx, &TelegramUserParameter{Username: data.Username})
		if err != nil {
			return err
		}
	}
	data.UserID = uint(ctx.Value("user_id").(int))
	if telUser.Username != data.Username && data.Username != "" {
		exist, _ := uc.telegramUserRepository.FindOne(ctx, &TelegramUserParameter{Username: data.Username})
		if exist.ID != 0 {
			return ErrUsernameExist
		}
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
	fmt.Println(RedisKeyActivatedTelegram + username + otp)
	res, err := uc.redisUsecase.Get(ctx, RedisKeyActivatedTelegram+username+otp)
	if err != nil {
		return ErrDataNotFound
	}
	id, _ := strconv.Atoi(res)
	param := new(TelegramUserParameter)
	param.Status = "true"
	exist, _ := uc.telegramUserRepository.FindOne(ctx, param)
	fmt.Println(exist)
	if exist.ID != 0 {
		exist.IsActive = false
		return uc.telegramUserRepository.EditStatus(ctx, &exist)
	}
	err = uc.telegramUserRepository.EditStatus(ctx, &TelegramUserDomain{ID: id, IsActive: true})
	if err != nil {
		return businesses.ErrInternalServer
	}
	return uc.redisUsecase.Del(ctx, RedisKeyActivatedTelegram+username+otp)
}

func (uc *telegramUserUsecase) Delete(ctx context.Context, data *TelegramUserDomain) error {
	return uc.telegramUserRepository.Delete(ctx, data)
}
