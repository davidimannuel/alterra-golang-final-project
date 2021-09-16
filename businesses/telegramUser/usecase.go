package telegramUser

import (
	"context"
	"keep-remind-app/businesses"
)

type telegramUserUsecase struct {
	telegramUserRepository TelegramUserRepository
}

func NewTelegramUserUsecase(telegramUserRepository TelegramUserRepository) TelegramUserUsecase {
	return &telegramUserUsecase{
		telegramUserRepository: telegramUserRepository,
	}
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
	res, err = uc.telegramUserRepository.Add(ctx, data)
	if err != nil {
		return
	}
	return
}

func (uc *telegramUserUsecase) Edit(ctx context.Context, data *TelegramUserDomain) error {
	return uc.telegramUserRepository.Edit(ctx, data)
}

func (uc *telegramUserUsecase) Delete(ctx context.Context, data *TelegramUserDomain) error {
	return uc.telegramUserRepository.Delete(ctx, data)
}
