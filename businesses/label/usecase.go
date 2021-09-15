package label

import (
	"context"
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
	res, err = uc.labelRepository.Add(ctx, data)
	if err != nil {
		return
	}
	return
}

func (uc *labelUsecase) Edit(ctx context.Context, data *LabelDomain) error {
	return uc.labelRepository.Edit(ctx, data)
}

func (uc *labelUsecase) Delete(ctx context.Context, data *LabelDomain) error {
	return uc.labelRepository.Delete(ctx, data)
}
