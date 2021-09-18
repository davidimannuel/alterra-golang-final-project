package ocr

import (
	"context"
	"keep-remind-app/drivers/ocr"
)

type ocrUsecase struct{}

func NewOCRUsecase() OCRUsecase {
	return &ocrUsecase{}
}

func (uc ocrUsecase) GetImageTextFromImagePath(ctx context.Context, path string) (res string, err error) {
	res, err = ocr.GetImageTextFromImagePath(path)
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc ocrUsecase) GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (res string, err error) {
	res, err = ocr.GetImageTextFromImageBytes(bytes)
	if err != nil {
		return res, err
	}
	return res, err
}
