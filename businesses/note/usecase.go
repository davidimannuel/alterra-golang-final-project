package note

import (
	"context"
	"keep-remind-app/businesses"
	"keep-remind-app/businesses/ocr"
)

type noteUsecase struct {
	noteRepository NoteRepository
	ocrUsecase     ocr.OCRUsecase
}

func NewNoteUsecase(noteRepository NoteRepository, ocrUsecase ocr.OCRUsecase) NoteUsecase {
	return &noteUsecase{
		noteRepository: noteRepository,
		ocrUsecase:     ocrUsecase,
	}
}

func (uc noteUsecase) FindAllPagination(ctx context.Context, param *NoteParameter) (res []NoteDomain, p businesses.Pagination, err error) {
	res, count, err := uc.noteRepository.FindAllPagination(ctx, param)
	if err != nil {
		return res, param.GetPageInfo(count), err
	}
	return
}

func (uc noteUsecase) FindAll(ctx context.Context, param *NoteParameter) (res []NoteDomain, err error) {
	res, err = uc.noteRepository.FindAll(ctx, param)
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc noteUsecase) FindOne(ctx context.Context, param *NoteParameter) (res NoteDomain, err error) {
	res, err = uc.noteRepository.FindOne(ctx, param)
	if err != nil {
		return
	}
	return
}

func (uc noteUsecase) Add(ctx context.Context, data *NoteDomain) (res int, err error) {
	data.UserID = ctx.Value("user_id").(int)
	res, err = uc.noteRepository.Add(ctx, data)
	if err != nil {
		return
	}
	return res, err
}

func (uc noteUsecase) AddWithImageBytes(ctx context.Context, title string, imageBytes []byte) (res int, err error) {
	text, err := uc.ocrUsecase.GetImageTextFromImageBytes(ctx, imageBytes)
	if err != nil {
		return res, err
	}
	res, err = uc.noteRepository.Add(ctx, &NoteDomain{
		UserID: ctx.Value("user_id").(int),
		Title:  title,
		Note:   text,
	})
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc noteUsecase) Edit(ctx context.Context, data *NoteDomain) error {
	panic("impl")
}
func (uc noteUsecase) Delete(ctx context.Context, id int) error {
	panic("impl")
}
