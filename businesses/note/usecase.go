package note

import (
	"context"
	"errors"
	"keep-remind-app/businesses"
	"keep-remind-app/businesses/label"
	"keep-remind-app/businesses/ocr"
)

type noteUsecase struct {
	noteRepository NoteRepository
	ocrUsecase     ocr.OCRUsecase
	labelUsecase   label.LabelUsecase
}

func NewNoteUsecase(noteRepository NoteRepository, ocrUsecase ocr.OCRUsecase, labelUsecase label.LabelUsecase) NoteUsecase {
	return &noteUsecase{
		noteRepository: noteRepository,
		ocrUsecase:     ocrUsecase,
		labelUsecase:   labelUsecase,
	}
}

func (uc *noteUsecase) validate(ctx context.Context, data *NoteDomain) (err error) {
	userID := ctx.Value("user_id").(int)
	if userID == 0 {
		return errors.New("Invalid User")
	}
	data.UserID = userID
	if len(data.Labels) > 0 {
		for i := range data.Labels {
			existLabel, _ := uc.labelUsecase.FindOne(ctx, &label.LabelParameter{Name: data.Labels[i].Name})
			data.Labels[i].ID = existLabel.ID
			data.Labels[i].UserID = userID
		}
	}
	return err
}

func (uc *noteUsecase) FindAllPagination(ctx context.Context, param *NoteParameter) (res []NoteDomain, p businesses.Pagination, err error) {
	res, count, err := uc.noteRepository.FindAllPagination(ctx, param)
	if err != nil {
		return res, param.GetPageInfo(count), err
	}
	return
}

func (uc *noteUsecase) FindAll(ctx context.Context, param *NoteParameter) (res []NoteDomain, err error) {
	res, err = uc.noteRepository.FindAll(ctx, param)
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc *noteUsecase) FindOne(ctx context.Context, param *NoteParameter) (res NoteDomain, err error) {
	res, err = uc.noteRepository.FindOne(ctx, param)
	if err != nil {
		return
	}
	return
}

func (uc *noteUsecase) Add(ctx context.Context, data *NoteDomain) (res int, err error) {
	err = uc.validate(ctx, data)
	if err != nil {
		return
	}
	res, err = uc.noteRepository.Add(ctx, data)
	if err != nil {
		return
	}
	return res, err
}

func (uc *noteUsecase) AddWithImageBytes(ctx context.Context, title string, imageBytes []byte) (res int, err error) {
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

func (uc *noteUsecase) Edit(ctx context.Context, data *NoteDomain) (err error) {
	err = uc.validate(ctx, data)
	if err != nil {
		return
	}
	return uc.noteRepository.Edit(ctx, data)
}
func (uc *noteUsecase) Delete(ctx context.Context, id int) error {
	return uc.noteRepository.Delete(ctx, id)
}
