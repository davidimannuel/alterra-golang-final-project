package note

import (
	"context"
	"keep-remind-app/businesses"
	"keep-remind-app/businesses/ocr"
)

type noteUsecase struct {
	repository NoteRepository
	ocrUsecase ocr.OCRUsecase
}

func NewNoteUsecase(repository NoteRepository, ocrUsecase ocr.OCRUsecase) NoteUsecase {
	return &noteUsecase{
		repository: repository,
		ocrUsecase: ocrUsecase,
	}
}

func (uc noteUsecase) Add(ctx context.Context, data *NoteDomain) (res NoteDomain, err error) {
	data.UserID = ctx.Value("user_id").(int)
	res.ID, err = uc.repository.Add(ctx, data)
	if err != nil {
		return NoteDomain{}, err
	}
	return res, err
}

func (uc noteUsecase) AddWithImageBytes(ctx context.Context, title string, imageBytes []byte) (res NoteDomain, err error) {
	text, err := uc.ocrUsecase.GetImageTextFromImageBytes(ctx, imageBytes)
	if err != nil {
		return res, err
	}
	res.ID, err = uc.repository.Add(ctx, &NoteDomain{
		UserID: ctx.Value("user_id").(int),
		Title:  title,
		Note:   text,
	})
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc noteUsecase) FindAll(ctx context.Context, param *NoteParameter) (res []NoteDomain, err error) {
	res, err = uc.repository.FindAll(ctx, param)
	if err != nil {
		return res, businesses.ErrInternalServer
	}
	return res, err
}

func (uc noteUsecase) FindByID(ctx context.Context, param *NoteParameter) (res NoteDomain, err error) {
	panic("implement me")
}
