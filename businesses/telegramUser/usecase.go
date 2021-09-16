package telegramUser

import (
	"context"
	"keep-remind-app/businesses"
	redisDomain "keep-remind-app/businesses/redis"
	"keep-remind-app/helpers/str"
	"strconv"
	"time"
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
	data.UserID = uint(ctx.Value("user_id").(int))
	if data.ID != 0 && data.Username != "" {
		exist, _ := uc.telegramUserRepository.FindOne(ctx, &TelegramUserParameter{Username: data.Username, UserID: data.UserID})
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
	key := redisKeyActivatedTelegram + strconv.Itoa(ctx.Value("user_id").(int)) + otp
	err = uc.redisUsecase.Set(ctx, key, exist.ID, time.Minute*5)
	return otp, err
}

func (uc *telegramUserUsecase) Activated(ctx context.Context, data *TelegramUserDomain) error {
	exist, _ := uc.telegramUserRepository.FindOne(ctx, &TelegramUserParameter{Username: data.Username, UserID: data.UserID})
	if exist.ID != 0 {
		exist.IsActive = false
		return uc.telegramUserRepository.EditStatus(ctx, &exist)
	}
	data.IsActive = true
	return uc.telegramUserRepository.EditStatus(ctx, data)
}

func (uc *telegramUserUsecase) Delete(ctx context.Context, data *TelegramUserDomain) error {
	return uc.telegramUserRepository.Delete(ctx, data)
}
