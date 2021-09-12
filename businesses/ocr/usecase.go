package ocr

import (
	"context"

	"github.com/otiai10/gosseract/v2"
)

type ocrUsecase struct {
	ocr *gosseract.Client
}

func NewUsecase(ocr *gosseract.Client) Usecase {
	return &ocrUsecase{
		ocr: ocr,
	}
}

func (uc ocrUsecase) GetImageTextFromImagePath(ctx context.Context, path string) (res Domain, err error) {
	uc.ocr.SetImage(path)
	text, err := uc.ocr.Text()
	if err != nil {
		return Domain{}, err
	}
	return Domain{ImageText: text}, nil
}

func (uc ocrUsecase) GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (res Domain, err error) {
	uc.ocr.SetImageFromBytes(bytes)
	text, err := uc.ocr.Text()
	if err != nil {
		return Domain{}, err
	}
	return Domain{ImageText: text}, nil
}
