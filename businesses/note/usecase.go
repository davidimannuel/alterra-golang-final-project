package note

import (
	"context"
	"keep-remind-app/businesses"
	"keep-remind-app/businesses/ocr"
)

type noteUsecase struct {
	repository Repository
	ocrUsecase ocr.Usecase
}

func NewUsecase(repository Repository, ocrUsecase ocr.Usecase) Usecase {
	return &noteUsecase{
		repository: repository,
		ocrUsecase: ocrUsecase,
	}
}

func (uc noteUsecase) Add(ctx context.Context, data *Domain) (res Domain, err error) {
	data.UserID = ctx.Value("user_id").(int)
	res.ID, err = uc.repository.Add(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	return res, err
}

func (uc noteUsecase) AddWithImageBytes(ctx context.Context, title string, imageBytes []byte) (res Domain, err error) {
	text, err := uc.ocrUsecase.GetImageTextFromImageBytes(ctx, imageBytes)
	if err != nil {
		return res, err
	}
	res.ID, err = uc.repository.Add(ctx, &Domain{
		UserID: ctx.Value("user_id").(int),
		Title:  title,
		Note:   text,
	})
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc noteUsecase) FindAll(ctx context.Context, parameter Parameter) (res []Domain, err error) {
	res, err = uc.repository.FindAll(ctx, parameter)
	if err != nil {
		return res, businesses.ErrInternalServer
	}
	return res, err
}

func (uc noteUsecase) FindByID(ctx context.Context, parameter Parameter) (res Domain, err error) {
	panic("implement me")
}
