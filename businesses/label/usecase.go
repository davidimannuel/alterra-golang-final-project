package label

import (
	"context"
	"errors"
	"keep-remind-app/businesses"
)

type labelUsecase struct {
	labelRepository LabelRepository
}

func NewLabelUsecase(labelRepository LabelRepository) LabelUsecase {
	return &labelUsecase{
		labelRepository: labelRepository,
	}
}

func (uc *labelUsecase) validate(ctx context.Context, data *LabelDomain) error {
	if ctx.Value("user_id").(int) == 0 {
		return errors.New("invalid user")
	}
	data.UserID = ctx.Value("user_id").(int)
	return nil
}

func (uc *labelUsecase) FindAllPagination(ctx context.Context, param *LabelParameter) (res []LabelDomain, p businesses.Pagination, err error) {
	res, count, err := uc.labelRepository.FindAllPagination(ctx, param)
	if err != nil {
		return
	}
	return res, param.GetPageInfo(count), err
}

func (uc *labelUsecase) FindAll(ctx context.Context, param *LabelParameter) (res []LabelDomain, err error) {
	res, err = uc.labelRepository.FindAll(ctx, param)
	if err != nil {
		return
	}
	return
}

func (uc *labelUsecase) FindOne(ctx context.Context, param *LabelParameter) (res LabelDomain, err error) {
	res, err = uc.labelRepository.FindOne(ctx, param)
	if err != nil {
		return
	}
	return
}

func (uc *labelUsecase) Add(ctx context.Context, data *LabelDomain) (res int, err error) {
	err = uc.validate(ctx, data)
	if err != nil {
		return
	}
	res, err = uc.labelRepository.Add(ctx, data)
	if err != nil {
		return
	}
	return
}

func (uc *labelUsecase) Edit(ctx context.Context, data *LabelDomain) (err error) {
	err = uc.validate(ctx, data)
	if err != nil {
		return
	}
	return uc.labelRepository.Edit(ctx, data)
}

func (uc *labelUsecase) Delete(ctx context.Context, id int) error {
	return uc.labelRepository.Delete(ctx, id)
}
